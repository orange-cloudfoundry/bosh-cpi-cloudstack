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
// Source: ./cloudstack/ConfigurationService.go
//
// Generated by this command:
//
//	mockgen -destination=./cloudstack/ConfigurationService_mock.go -package=cloudstack -copyright_file=header.txt -source=./cloudstack/ConfigurationService.go
//

// Package cloudstack is a generated GoMock package.
package cloudstack

import (
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockConfigurationServiceIface is a mock of ConfigurationServiceIface interface.
type MockConfigurationServiceIface struct {
	ctrl     *gomock.Controller
	recorder *MockConfigurationServiceIfaceMockRecorder
	isgomock struct{}
}

// MockConfigurationServiceIfaceMockRecorder is the mock recorder for MockConfigurationServiceIface.
type MockConfigurationServiceIfaceMockRecorder struct {
	mock *MockConfigurationServiceIface
}

// NewMockConfigurationServiceIface creates a new mock instance.
func NewMockConfigurationServiceIface(ctrl *gomock.Controller) *MockConfigurationServiceIface {
	mock := &MockConfigurationServiceIface{ctrl: ctrl}
	mock.recorder = &MockConfigurationServiceIfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockConfigurationServiceIface) EXPECT() *MockConfigurationServiceIfaceMockRecorder {
	return m.recorder
}

// ListCapabilities mocks base method.
func (m *MockConfigurationServiceIface) ListCapabilities(p *ListCapabilitiesParams) (*ListCapabilitiesResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListCapabilities", p)
	ret0, _ := ret[0].(*ListCapabilitiesResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListCapabilities indicates an expected call of ListCapabilities.
func (mr *MockConfigurationServiceIfaceMockRecorder) ListCapabilities(p any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListCapabilities", reflect.TypeOf((*MockConfigurationServiceIface)(nil).ListCapabilities), p)
}

// ListConfigurationGroups mocks base method.
func (m *MockConfigurationServiceIface) ListConfigurationGroups(p *ListConfigurationGroupsParams) (*ListConfigurationGroupsResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListConfigurationGroups", p)
	ret0, _ := ret[0].(*ListConfigurationGroupsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListConfigurationGroups indicates an expected call of ListConfigurationGroups.
func (mr *MockConfigurationServiceIfaceMockRecorder) ListConfigurationGroups(p any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListConfigurationGroups", reflect.TypeOf((*MockConfigurationServiceIface)(nil).ListConfigurationGroups), p)
}

// ListConfigurations mocks base method.
func (m *MockConfigurationServiceIface) ListConfigurations(p *ListConfigurationsParams) (*ListConfigurationsResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListConfigurations", p)
	ret0, _ := ret[0].(*ListConfigurationsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListConfigurations indicates an expected call of ListConfigurations.
func (mr *MockConfigurationServiceIfaceMockRecorder) ListConfigurations(p any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListConfigurations", reflect.TypeOf((*MockConfigurationServiceIface)(nil).ListConfigurations), p)
}

// ListDeploymentPlanners mocks base method.
func (m *MockConfigurationServiceIface) ListDeploymentPlanners(p *ListDeploymentPlannersParams) (*ListDeploymentPlannersResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListDeploymentPlanners", p)
	ret0, _ := ret[0].(*ListDeploymentPlannersResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListDeploymentPlanners indicates an expected call of ListDeploymentPlanners.
func (mr *MockConfigurationServiceIfaceMockRecorder) ListDeploymentPlanners(p any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListDeploymentPlanners", reflect.TypeOf((*MockConfigurationServiceIface)(nil).ListDeploymentPlanners), p)
}

// NewListCapabilitiesParams mocks base method.
func (m *MockConfigurationServiceIface) NewListCapabilitiesParams() *ListCapabilitiesParams {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewListCapabilitiesParams")
	ret0, _ := ret[0].(*ListCapabilitiesParams)
	return ret0
}

// NewListCapabilitiesParams indicates an expected call of NewListCapabilitiesParams.
func (mr *MockConfigurationServiceIfaceMockRecorder) NewListCapabilitiesParams() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewListCapabilitiesParams", reflect.TypeOf((*MockConfigurationServiceIface)(nil).NewListCapabilitiesParams))
}

// NewListConfigurationGroupsParams mocks base method.
func (m *MockConfigurationServiceIface) NewListConfigurationGroupsParams() *ListConfigurationGroupsParams {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewListConfigurationGroupsParams")
	ret0, _ := ret[0].(*ListConfigurationGroupsParams)
	return ret0
}

// NewListConfigurationGroupsParams indicates an expected call of NewListConfigurationGroupsParams.
func (mr *MockConfigurationServiceIfaceMockRecorder) NewListConfigurationGroupsParams() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewListConfigurationGroupsParams", reflect.TypeOf((*MockConfigurationServiceIface)(nil).NewListConfigurationGroupsParams))
}

// NewListConfigurationsParams mocks base method.
func (m *MockConfigurationServiceIface) NewListConfigurationsParams() *ListConfigurationsParams {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewListConfigurationsParams")
	ret0, _ := ret[0].(*ListConfigurationsParams)
	return ret0
}

// NewListConfigurationsParams indicates an expected call of NewListConfigurationsParams.
func (mr *MockConfigurationServiceIfaceMockRecorder) NewListConfigurationsParams() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewListConfigurationsParams", reflect.TypeOf((*MockConfigurationServiceIface)(nil).NewListConfigurationsParams))
}

// NewListDeploymentPlannersParams mocks base method.
func (m *MockConfigurationServiceIface) NewListDeploymentPlannersParams() *ListDeploymentPlannersParams {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewListDeploymentPlannersParams")
	ret0, _ := ret[0].(*ListDeploymentPlannersParams)
	return ret0
}

// NewListDeploymentPlannersParams indicates an expected call of NewListDeploymentPlannersParams.
func (mr *MockConfigurationServiceIfaceMockRecorder) NewListDeploymentPlannersParams() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewListDeploymentPlannersParams", reflect.TypeOf((*MockConfigurationServiceIface)(nil).NewListDeploymentPlannersParams))
}

// NewResetConfigurationParams mocks base method.
func (m *MockConfigurationServiceIface) NewResetConfigurationParams(name string) *ResetConfigurationParams {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewResetConfigurationParams", name)
	ret0, _ := ret[0].(*ResetConfigurationParams)
	return ret0
}

// NewResetConfigurationParams indicates an expected call of NewResetConfigurationParams.
func (mr *MockConfigurationServiceIfaceMockRecorder) NewResetConfigurationParams(name any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewResetConfigurationParams", reflect.TypeOf((*MockConfigurationServiceIface)(nil).NewResetConfigurationParams), name)
}

// NewUpdateConfigurationParams mocks base method.
func (m *MockConfigurationServiceIface) NewUpdateConfigurationParams(name string) *UpdateConfigurationParams {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewUpdateConfigurationParams", name)
	ret0, _ := ret[0].(*UpdateConfigurationParams)
	return ret0
}

// NewUpdateConfigurationParams indicates an expected call of NewUpdateConfigurationParams.
func (mr *MockConfigurationServiceIfaceMockRecorder) NewUpdateConfigurationParams(name any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewUpdateConfigurationParams", reflect.TypeOf((*MockConfigurationServiceIface)(nil).NewUpdateConfigurationParams), name)
}

// NewUpdateStorageCapabilitiesParams mocks base method.
func (m *MockConfigurationServiceIface) NewUpdateStorageCapabilitiesParams(id string) *UpdateStorageCapabilitiesParams {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewUpdateStorageCapabilitiesParams", id)
	ret0, _ := ret[0].(*UpdateStorageCapabilitiesParams)
	return ret0
}

// NewUpdateStorageCapabilitiesParams indicates an expected call of NewUpdateStorageCapabilitiesParams.
func (mr *MockConfigurationServiceIfaceMockRecorder) NewUpdateStorageCapabilitiesParams(id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewUpdateStorageCapabilitiesParams", reflect.TypeOf((*MockConfigurationServiceIface)(nil).NewUpdateStorageCapabilitiesParams), id)
}

// ResetConfiguration mocks base method.
func (m *MockConfigurationServiceIface) ResetConfiguration(p *ResetConfigurationParams) (*ResetConfigurationResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ResetConfiguration", p)
	ret0, _ := ret[0].(*ResetConfigurationResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ResetConfiguration indicates an expected call of ResetConfiguration.
func (mr *MockConfigurationServiceIfaceMockRecorder) ResetConfiguration(p any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ResetConfiguration", reflect.TypeOf((*MockConfigurationServiceIface)(nil).ResetConfiguration), p)
}

// UpdateConfiguration mocks base method.
func (m *MockConfigurationServiceIface) UpdateConfiguration(p *UpdateConfigurationParams) (*UpdateConfigurationResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateConfiguration", p)
	ret0, _ := ret[0].(*UpdateConfigurationResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateConfiguration indicates an expected call of UpdateConfiguration.
func (mr *MockConfigurationServiceIfaceMockRecorder) UpdateConfiguration(p any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateConfiguration", reflect.TypeOf((*MockConfigurationServiceIface)(nil).UpdateConfiguration), p)
}

// UpdateStorageCapabilities mocks base method.
func (m *MockConfigurationServiceIface) UpdateStorageCapabilities(p *UpdateStorageCapabilitiesParams) (*UpdateStorageCapabilitiesResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateStorageCapabilities", p)
	ret0, _ := ret[0].(*UpdateStorageCapabilitiesResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateStorageCapabilities indicates an expected call of UpdateStorageCapabilities.
func (mr *MockConfigurationServiceIfaceMockRecorder) UpdateStorageCapabilities(p any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateStorageCapabilities", reflect.TypeOf((*MockConfigurationServiceIface)(nil).UpdateStorageCapabilities), p)
}
