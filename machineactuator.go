/*
Copyright 2018 The Kubernetes Authors.
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at
    http://www.apache.org/licenses/LICENSE-2.0
Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package azureactuator

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/Azure/azure-sdk-for-go/services/resources/mgmt/2018-02-01/resources"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/autorest/azure/auth"
	azureconfigv1 "github.com/platform9/azure-actuator/azureproviderconfig/v1alpha1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/serializer"
	clusterv1 "sigs.k8s.io/cluster-api/pkg/apis/cluster/v1alpha1"
)

type AzureClient struct {
	SubscriptionID string
	VMPassword     string
	Authorizer     autorest.Authorizer
	kubeadmToken   string
	ctx            context.Context
	//	machineClient  client.MachineInterface
	scheme       *runtime.Scheme
	codecFactory *serializer.CodecFactory
}

type MachineActuatorParams struct {
	KubeadmToken string
	//TODO Add more
}

const (
	templateFile   = "deployment-template.json"
	parametersFile = "deployment-params.json"
)

func NewMachineActuator(params MachineActuatorParams) (*AzureClient, error) {
	scheme, codecFactory, err := azureconfigv1.NewSchemeAndCodecs()
	authorizer, err := auth.NewAuthorizerFromFile(azure.PublicCloud.ResourceManagerEndpoint)
	if err != nil {
		log.Fatalf("Failed to get OAuth config: %v", err)
		return nil, err
	}
	authInfo, err := readJSON(os.Getenv("AZURE_AUTH_LOCATION"))
	if err != nil {
		return nil, err
	}
	return &AzureClient{
		SubscriptionID: (*authInfo)["subscriptionId"].(string),
		VMPassword:     (*authInfo)["clientSecret"].(string), //Do NOT keep - should not reuse secret here
		Authorizer:     authorizer,
		kubeadmToken:   params.KubeadmToken,
		ctx:            context.Background(),
		scheme:         scheme,
		codecFactory:   codecFactory,
	}, nil
}

func (azure *AzureClient) Create(cluster *clusterv1.Cluster, machine *clusterv1.Machine) error {
	_, err := azure.createOrUpdateGroup(cluster)
	if err != nil {
		return err
	}

	_, err = azure.createOrUpdateDeployment(cluster, machine)
	if err != nil {
		return err
	}

	//Get the Login info from the VMs
	/*
		_, _, err = azure.getLogin(cluster, machine)
		if err != nil {
			return err
		}
	*/

	//Set up Kubernetes
	return nil
}

func (azure *AzureClient) Update(cluster *clusterv1.Cluster, goalMachine *clusterv1.Machine) error {
	//Parse in configurations
	var goalMachineConfig azureconfigv1.AzureMachineProviderConfig
	err := azure.decodeMachineProviderConfig(goalMachine.Spec.ProviderConfig, &goalMachineConfig)
	if err != nil {
		return err
	}
	var clusterConfig azureconfigv1.AzureClusterProviderConfig
	err = azure.decodeClusterProviderConfig(cluster.Spec.ProviderConfig, &clusterConfig)
	if err != nil {
		return err
	}
	_, err = azure.vmIfExists(cluster, goalMachine)
	return nil
}

func (azure *AzureClient) Delete(cluster *clusterv1.Cluster, machine *clusterv1.Machine) error {
	//Parse in configurations
	var machineConfig azureconfigv1.AzureMachineProviderConfig
	err := azure.decodeMachineProviderConfig(machine.Spec.ProviderConfig, &machineConfig)
	if err != nil {
		return err
	}
	var clusterConfig azureconfigv1.AzureClusterProviderConfig
	err = azure.decodeClusterProviderConfig(cluster.Spec.ProviderConfig, &clusterConfig)
	if err != nil {
		return err
	}
	//Check if the machine exists
	vm, err := azure.vmIfExists(cluster, machine)
	if vm == nil {
		//Skip deleting if we couldn't find anything to delete
		return nil
	}
	//Delete the resource group, which will automatically delete
	//all associated resources
	groupsClient := resources.NewGroupsClient(azure.SubscriptionID)
	groupsClient.Authorizer = azure.Authorizer
	groupsDeleteFuture, err := groupsClient.Delete(azure.ctx, clusterConfig.ResourceGroup)
	if err != nil {
		return err
	}
	return groupsDeleteFuture.Future.WaitForCompletion(azure.ctx, groupsClient.BaseClient.Client)
}

func (azure *AzureClient) GetKubeConfig(cluster *clusterv1.Cluster, machine *clusterv1.Machine) (string, error) {
	return "", nil
}

func (azure *AzureClient) Exists(cluster *clusterv1.Cluster, machine *clusterv1.Machine) (bool, error) {
	vm, err := azure.vmIfExists(cluster, machine)
	if err != nil {
		return false, err
	}
	return (vm != nil), nil
}

func (azure *AzureClient) decodeMachineProviderConfig(providerConfig clusterv1.ProviderConfig, out runtime.Object) error {
	_, _, err := azure.codecFactory.UniversalDecoder().Decode(providerConfig.Value.Raw, nil, out)
	if err != nil {
		return fmt.Errorf("machine providerconfig decoding failure: %v", err)
	}
	return nil
}

func (azure *AzureClient) decodeClusterProviderConfig(providerConfig clusterv1.ProviderConfig, out runtime.Object) error {
	_, _, err := azure.codecFactory.UniversalDecoder().Decode(providerConfig.Value.Raw, nil, out)
	if err != nil {
		return fmt.Errorf("cluster providerconfig decoding failure: %v", err)
	}
	return nil
}

func readJSON(path string) (*map[string]interface{}, error) {
	fileContents, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	data := make(map[string]interface{})
	err = json.Unmarshal(fileContents, &data)
	if err != nil {
		return nil, err
	}

	return &data, nil
}

func (azure *AzureClient) convertMachineToDeploymentParams(machine *clusterv1.Machine) (*map[string]interface{}, error) {
	var machineConfig azureconfigv1.AzureMachineProviderConfig
	err := azure.decodeMachineProviderConfig(machine.Spec.ProviderConfig, &machineConfig)
	if err != nil {
		return nil, err
	}
	params := map[string]interface{}{
		"virtualNetworks_ClusterAPIVM_vnet_name": map[string]interface{}{
			"value": "ClusterAPIVnet",
		},
		"virtualMachines_ClusterAPIVM_name": map[string]interface{}{
			"value": getVMName(machine),
		},
		"networkInterfaces_ClusterAPI_name": map[string]interface{}{
			"value": getNetworkInterfaceName(machine),
		},
		"publicIPAddresses_ClusterAPI_ip_name": map[string]interface{}{
			"value": getPublicIPName(machine),
		},
		"networkSecurityGroups_ClusterAPIVM_nsg_name": map[string]interface{}{
			"value": "ClusterAPINSG",
		},
		"subnets_default_name": map[string]interface{}{
			"value": "ClusterAPISubnet",
		},
		"securityRules_default_allow_ssh_name": map[string]interface{}{
			"value": "clusterapiuser",
		},
		"image_publisher": map[string]interface{}{
			"value": machineConfig.Image.Publisher,
		},
		"image_offer": map[string]interface{}{
			"value": machineConfig.Image.Offer,
		},
		"image_sku": map[string]interface{}{
			"value": machineConfig.Image.SKU,
		},
		"image_version": map[string]interface{}{
			"value": machineConfig.Image.Version,
		},
		"osDisk_name": map[string]interface{}{
			"value": getOSDiskName(machine),
		},
		"os_type": map[string]interface{}{
			"value": machineConfig.OSDisk.OSType,
		},
		"storage_account_type": map[string]interface{}{
			"value": machineConfig.OSDisk.ManagedDisk.StorageAccountType,
		},
		"disk_size_GB": map[string]interface{}{
			"value": machineConfig.OSDisk.DiskSizeGB,
		},
		"vm_user": map[string]interface{}{
			"value": "ClusterAPI",
		},
		"vm_password": map[string]interface{}{
			"value": "_",
		},
		"vm_size": map[string]interface{}{
			"value": machineConfig.VMSize,
		},
		"location": map[string]interface{}{
			"value": machineConfig.Location,
		},
	}
	return &params, nil
}
