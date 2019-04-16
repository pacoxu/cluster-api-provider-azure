/*
Copyright 2019 The Kubernetes Authors.

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

// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/Azure/azure-sdk-for-go/services/network/mgmt/2018-12-01/network/networkapi (interfaces: VirtualNetworksClientAPI)

// Package mock_virtualnetworks is a generated GoMock package.
package mock_virtualnetworks

import (
	context "context"
	network "github.com/Azure/azure-sdk-for-go/services/network/mgmt/2018-12-01/network"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockVirtualNetworksClientAPI is a mock of VirtualNetworksClientAPI interface
type MockVirtualNetworksClientAPI struct {
	ctrl     *gomock.Controller
	recorder *MockVirtualNetworksClientAPIMockRecorder
}

// MockVirtualNetworksClientAPIMockRecorder is the mock recorder for MockVirtualNetworksClientAPI
type MockVirtualNetworksClientAPIMockRecorder struct {
	mock *MockVirtualNetworksClientAPI
}

// NewMockVirtualNetworksClientAPI creates a new mock instance
func NewMockVirtualNetworksClientAPI(ctrl *gomock.Controller) *MockVirtualNetworksClientAPI {
	mock := &MockVirtualNetworksClientAPI{ctrl: ctrl}
	mock.recorder = &MockVirtualNetworksClientAPIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockVirtualNetworksClientAPI) EXPECT() *MockVirtualNetworksClientAPIMockRecorder {
	return m.recorder
}

// CheckIPAddressAvailability mocks base method
func (m *MockVirtualNetworksClientAPI) CheckIPAddressAvailability(arg0 context.Context, arg1, arg2, arg3 string) (network.IPAddressAvailabilityResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckIPAddressAvailability", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(network.IPAddressAvailabilityResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CheckIPAddressAvailability indicates an expected call of CheckIPAddressAvailability
func (mr *MockVirtualNetworksClientAPIMockRecorder) CheckIPAddressAvailability(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckIPAddressAvailability", reflect.TypeOf((*MockVirtualNetworksClientAPI)(nil).CheckIPAddressAvailability), arg0, arg1, arg2, arg3)
}

// CreateOrUpdate mocks base method
func (m *MockVirtualNetworksClientAPI) CreateOrUpdate(arg0 context.Context, arg1, arg2 string, arg3 network.VirtualNetwork) (network.VirtualNetworksCreateOrUpdateFuture, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateOrUpdate", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(network.VirtualNetworksCreateOrUpdateFuture)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateOrUpdate indicates an expected call of CreateOrUpdate
func (mr *MockVirtualNetworksClientAPIMockRecorder) CreateOrUpdate(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateOrUpdate", reflect.TypeOf((*MockVirtualNetworksClientAPI)(nil).CreateOrUpdate), arg0, arg1, arg2, arg3)
}

// Delete mocks base method
func (m *MockVirtualNetworksClientAPI) Delete(arg0 context.Context, arg1, arg2 string) (network.VirtualNetworksDeleteFuture, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0, arg1, arg2)
	ret0, _ := ret[0].(network.VirtualNetworksDeleteFuture)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Delete indicates an expected call of Delete
func (mr *MockVirtualNetworksClientAPIMockRecorder) Delete(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockVirtualNetworksClientAPI)(nil).Delete), arg0, arg1, arg2)
}

// Get mocks base method
func (m *MockVirtualNetworksClientAPI) Get(arg0 context.Context, arg1, arg2, arg3 string) (network.VirtualNetwork, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(network.VirtualNetwork)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get
func (mr *MockVirtualNetworksClientAPIMockRecorder) Get(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockVirtualNetworksClientAPI)(nil).Get), arg0, arg1, arg2, arg3)
}

// List mocks base method
func (m *MockVirtualNetworksClientAPI) List(arg0 context.Context, arg1 string) (network.VirtualNetworkListResultPage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0, arg1)
	ret0, _ := ret[0].(network.VirtualNetworkListResultPage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List
func (mr *MockVirtualNetworksClientAPIMockRecorder) List(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockVirtualNetworksClientAPI)(nil).List), arg0, arg1)
}

// ListAll mocks base method
func (m *MockVirtualNetworksClientAPI) ListAll(arg0 context.Context) (network.VirtualNetworkListResultPage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListAll", arg0)
	ret0, _ := ret[0].(network.VirtualNetworkListResultPage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListAll indicates an expected call of ListAll
func (mr *MockVirtualNetworksClientAPIMockRecorder) ListAll(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAll", reflect.TypeOf((*MockVirtualNetworksClientAPI)(nil).ListAll), arg0)
}

// ListUsage mocks base method
func (m *MockVirtualNetworksClientAPI) ListUsage(arg0 context.Context, arg1, arg2 string) (network.VirtualNetworkListUsageResultPage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListUsage", arg0, arg1, arg2)
	ret0, _ := ret[0].(network.VirtualNetworkListUsageResultPage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListUsage indicates an expected call of ListUsage
func (mr *MockVirtualNetworksClientAPIMockRecorder) ListUsage(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListUsage", reflect.TypeOf((*MockVirtualNetworksClientAPI)(nil).ListUsage), arg0, arg1, arg2)
}

// UpdateTags mocks base method
func (m *MockVirtualNetworksClientAPI) UpdateTags(arg0 context.Context, arg1, arg2 string, arg3 network.TagsObject) (network.VirtualNetworksUpdateTagsFuture, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateTags", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(network.VirtualNetworksUpdateTagsFuture)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateTags indicates an expected call of UpdateTags
func (mr *MockVirtualNetworksClientAPIMockRecorder) UpdateTags(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateTags", reflect.TypeOf((*MockVirtualNetworksClientAPI)(nil).UpdateTags), arg0, arg1, arg2, arg3)
}
