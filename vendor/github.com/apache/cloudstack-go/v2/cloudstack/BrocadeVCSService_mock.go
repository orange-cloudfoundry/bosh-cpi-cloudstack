//
// Licensed to the Apache Software Foundation (ASF) under one
// or more contributor license agreements.  See the NOTICE file
// distributed with this work for additional information
// regarding copyright ownership.  The ASF licenses this file
// to you under the Apache License, Version 2.0 (the
// "License"); you may not use this file except in compliance
// with the License.  You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.
//

// Code generated by MockGen. DO NOT EDIT.
// Source: ./cloudstack/BrocadeVCSService.go
//
// Generated by this command:
//
//	mockgen -destination=./cloudstack/BrocadeVCSService_mock.go -package=cloudstack -copyright_file=header.txt -source=./cloudstack/BrocadeVCSService.go
//

// Package cloudstack is a generated GoMock package.
package cloudstack

import (
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockBrocadeVCSServiceIface is a mock of BrocadeVCSServiceIface interface.
type MockBrocadeVCSServiceIface struct {
	ctrl     *gomock.Controller
	recorder *MockBrocadeVCSServiceIfaceMockRecorder
	isgomock struct{}
}

// MockBrocadeVCSServiceIfaceMockRecorder is the mock recorder for MockBrocadeVCSServiceIface.
type MockBrocadeVCSServiceIfaceMockRecorder struct {
	mock *MockBrocadeVCSServiceIface
}

// NewMockBrocadeVCSServiceIface creates a new mock instance.
func NewMockBrocadeVCSServiceIface(ctrl *gomock.Controller) *MockBrocadeVCSServiceIface {
	mock := &MockBrocadeVCSServiceIface{ctrl: ctrl}
	mock.recorder = &MockBrocadeVCSServiceIfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBrocadeVCSServiceIface) EXPECT() *MockBrocadeVCSServiceIfaceMockRecorder {
	return m.recorder
}

// AddBrocadeVcsDevice mocks base method.
func (m *MockBrocadeVCSServiceIface) AddBrocadeVcsDevice(p *AddBrocadeVcsDeviceParams) (*AddBrocadeVcsDeviceResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddBrocadeVcsDevice", p)
	ret0, _ := ret[0].(*AddBrocadeVcsDeviceResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddBrocadeVcsDevice indicates an expected call of AddBrocadeVcsDevice.
func (mr *MockBrocadeVCSServiceIfaceMockRecorder) AddBrocadeVcsDevice(p any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddBrocadeVcsDevice", reflect.TypeOf((*MockBrocadeVCSServiceIface)(nil).AddBrocadeVcsDevice), p)
}

// DeleteBrocadeVcsDevice mocks base method.
func (m *MockBrocadeVCSServiceIface) DeleteBrocadeVcsDevice(p *DeleteBrocadeVcsDeviceParams) (*DeleteBrocadeVcsDeviceResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteBrocadeVcsDevice", p)
	ret0, _ := ret[0].(*DeleteBrocadeVcsDeviceResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteBrocadeVcsDevice indicates an expected call of DeleteBrocadeVcsDevice.
func (mr *MockBrocadeVCSServiceIfaceMockRecorder) DeleteBrocadeVcsDevice(p any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteBrocadeVcsDevice", reflect.TypeOf((*MockBrocadeVCSServiceIface)(nil).DeleteBrocadeVcsDevice), p)
}

// GetBrocadeVcsDeviceNetworkID mocks base method.
func (m *MockBrocadeVCSServiceIface) GetBrocadeVcsDeviceNetworkID(keyword, vcsdeviceid string, opts ...OptionFunc) (string, int, error) {
	m.ctrl.T.Helper()
	varargs := []any{keyword, vcsdeviceid}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetBrocadeVcsDeviceNetworkID", varargs...)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(int)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetBrocadeVcsDeviceNetworkID indicates an expected call of GetBrocadeVcsDeviceNetworkID.
func (mr *MockBrocadeVCSServiceIfaceMockRecorder) GetBrocadeVcsDeviceNetworkID(keyword, vcsdeviceid any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{keyword, vcsdeviceid}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBrocadeVcsDeviceNetworkID", reflect.TypeOf((*MockBrocadeVCSServiceIface)(nil).GetBrocadeVcsDeviceNetworkID), varargs...)
}

// ListBrocadeVcsDeviceNetworks mocks base method.
func (m *MockBrocadeVCSServiceIface) ListBrocadeVcsDeviceNetworks(p *ListBrocadeVcsDeviceNetworksParams) (*ListBrocadeVcsDeviceNetworksResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListBrocadeVcsDeviceNetworks", p)
	ret0, _ := ret[0].(*ListBrocadeVcsDeviceNetworksResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListBrocadeVcsDeviceNetworks indicates an expected call of ListBrocadeVcsDeviceNetworks.
func (mr *MockBrocadeVCSServiceIfaceMockRecorder) ListBrocadeVcsDeviceNetworks(p any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListBrocadeVcsDeviceNetworks", reflect.TypeOf((*MockBrocadeVCSServiceIface)(nil).ListBrocadeVcsDeviceNetworks), p)
}

// ListBrocadeVcsDevices mocks base method.
func (m *MockBrocadeVCSServiceIface) ListBrocadeVcsDevices(p *ListBrocadeVcsDevicesParams) (*ListBrocadeVcsDevicesResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListBrocadeVcsDevices", p)
	ret0, _ := ret[0].(*ListBrocadeVcsDevicesResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListBrocadeVcsDevices indicates an expected call of ListBrocadeVcsDevices.
func (mr *MockBrocadeVCSServiceIfaceMockRecorder) ListBrocadeVcsDevices(p any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListBrocadeVcsDevices", reflect.TypeOf((*MockBrocadeVCSServiceIface)(nil).ListBrocadeVcsDevices), p)
}

// NewAddBrocadeVcsDeviceParams mocks base method.
func (m *MockBrocadeVCSServiceIface) NewAddBrocadeVcsDeviceParams(hostname, password, physicalnetworkid, username string) *AddBrocadeVcsDeviceParams {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewAddBrocadeVcsDeviceParams", hostname, password, physicalnetworkid, username)
	ret0, _ := ret[0].(*AddBrocadeVcsDeviceParams)
	return ret0
}

// NewAddBrocadeVcsDeviceParams indicates an expected call of NewAddBrocadeVcsDeviceParams.
func (mr *MockBrocadeVCSServiceIfaceMockRecorder) NewAddBrocadeVcsDeviceParams(hostname, password, physicalnetworkid, username any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewAddBrocadeVcsDeviceParams", reflect.TypeOf((*MockBrocadeVCSServiceIface)(nil).NewAddBrocadeVcsDeviceParams), hostname, password, physicalnetworkid, username)
}

// NewDeleteBrocadeVcsDeviceParams mocks base method.
func (m *MockBrocadeVCSServiceIface) NewDeleteBrocadeVcsDeviceParams(vcsdeviceid string) *DeleteBrocadeVcsDeviceParams {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewDeleteBrocadeVcsDeviceParams", vcsdeviceid)
	ret0, _ := ret[0].(*DeleteBrocadeVcsDeviceParams)
	return ret0
}

// NewDeleteBrocadeVcsDeviceParams indicates an expected call of NewDeleteBrocadeVcsDeviceParams.
func (mr *MockBrocadeVCSServiceIfaceMockRecorder) NewDeleteBrocadeVcsDeviceParams(vcsdeviceid any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewDeleteBrocadeVcsDeviceParams", reflect.TypeOf((*MockBrocadeVCSServiceIface)(nil).NewDeleteBrocadeVcsDeviceParams), vcsdeviceid)
}

// NewListBrocadeVcsDeviceNetworksParams mocks base method.
func (m *MockBrocadeVCSServiceIface) NewListBrocadeVcsDeviceNetworksParams(vcsdeviceid string) *ListBrocadeVcsDeviceNetworksParams {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewListBrocadeVcsDeviceNetworksParams", vcsdeviceid)
	ret0, _ := ret[0].(*ListBrocadeVcsDeviceNetworksParams)
	return ret0
}

// NewListBrocadeVcsDeviceNetworksParams indicates an expected call of NewListBrocadeVcsDeviceNetworksParams.
func (mr *MockBrocadeVCSServiceIfaceMockRecorder) NewListBrocadeVcsDeviceNetworksParams(vcsdeviceid any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewListBrocadeVcsDeviceNetworksParams", reflect.TypeOf((*MockBrocadeVCSServiceIface)(nil).NewListBrocadeVcsDeviceNetworksParams), vcsdeviceid)
}

// NewListBrocadeVcsDevicesParams mocks base method.
func (m *MockBrocadeVCSServiceIface) NewListBrocadeVcsDevicesParams() *ListBrocadeVcsDevicesParams {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewListBrocadeVcsDevicesParams")
	ret0, _ := ret[0].(*ListBrocadeVcsDevicesParams)
	return ret0
}

// NewListBrocadeVcsDevicesParams indicates an expected call of NewListBrocadeVcsDevicesParams.
func (mr *MockBrocadeVCSServiceIfaceMockRecorder) NewListBrocadeVcsDevicesParams() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewListBrocadeVcsDevicesParams", reflect.TypeOf((*MockBrocadeVCSServiceIface)(nil).NewListBrocadeVcsDevicesParams))
}
