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
// Source: ./cloudstack/SystemVMService.go
//
// Generated by this command:
//
//	mockgen -destination=./cloudstack/SystemVMService_mock.go -package=cloudstack -copyright_file=header.txt -source=./cloudstack/SystemVMService.go
//

// Package cloudstack is a generated GoMock package.
package cloudstack

import (
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockSystemVMServiceIface is a mock of SystemVMServiceIface interface.
type MockSystemVMServiceIface struct {
	ctrl     *gomock.Controller
	recorder *MockSystemVMServiceIfaceMockRecorder
	isgomock struct{}
}

// MockSystemVMServiceIfaceMockRecorder is the mock recorder for MockSystemVMServiceIface.
type MockSystemVMServiceIfaceMockRecorder struct {
	mock *MockSystemVMServiceIface
}

// NewMockSystemVMServiceIface creates a new mock instance.
func NewMockSystemVMServiceIface(ctrl *gomock.Controller) *MockSystemVMServiceIface {
	mock := &MockSystemVMServiceIface{ctrl: ctrl}
	mock.recorder = &MockSystemVMServiceIfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSystemVMServiceIface) EXPECT() *MockSystemVMServiceIfaceMockRecorder {
	return m.recorder
}

// ChangeServiceForSystemVm mocks base method.
func (m *MockSystemVMServiceIface) ChangeServiceForSystemVm(p *ChangeServiceForSystemVmParams) (*ChangeServiceForSystemVmResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ChangeServiceForSystemVm", p)
	ret0, _ := ret[0].(*ChangeServiceForSystemVmResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ChangeServiceForSystemVm indicates an expected call of ChangeServiceForSystemVm.
func (mr *MockSystemVMServiceIfaceMockRecorder) ChangeServiceForSystemVm(p any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ChangeServiceForSystemVm", reflect.TypeOf((*MockSystemVMServiceIface)(nil).ChangeServiceForSystemVm), p)
}

// DestroySystemVm mocks base method.
func (m *MockSystemVMServiceIface) DestroySystemVm(p *DestroySystemVmParams) (*DestroySystemVmResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DestroySystemVm", p)
	ret0, _ := ret[0].(*DestroySystemVmResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DestroySystemVm indicates an expected call of DestroySystemVm.
func (mr *MockSystemVMServiceIfaceMockRecorder) DestroySystemVm(p any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DestroySystemVm", reflect.TypeOf((*MockSystemVMServiceIface)(nil).DestroySystemVm), p)
}

// GetSystemVmByID mocks base method.
func (m *MockSystemVMServiceIface) GetSystemVmByID(id string, opts ...OptionFunc) (*SystemVm, int, error) {
	m.ctrl.T.Helper()
	varargs := []any{id}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetSystemVmByID", varargs...)
	ret0, _ := ret[0].(*SystemVm)
	ret1, _ := ret[1].(int)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetSystemVmByID indicates an expected call of GetSystemVmByID.
func (mr *MockSystemVMServiceIfaceMockRecorder) GetSystemVmByID(id any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{id}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSystemVmByID", reflect.TypeOf((*MockSystemVMServiceIface)(nil).GetSystemVmByID), varargs...)
}

// GetSystemVmByName mocks base method.
func (m *MockSystemVMServiceIface) GetSystemVmByName(name string, opts ...OptionFunc) (*SystemVm, int, error) {
	m.ctrl.T.Helper()
	varargs := []any{name}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetSystemVmByName", varargs...)
	ret0, _ := ret[0].(*SystemVm)
	ret1, _ := ret[1].(int)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetSystemVmByName indicates an expected call of GetSystemVmByName.
func (mr *MockSystemVMServiceIfaceMockRecorder) GetSystemVmByName(name any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{name}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSystemVmByName", reflect.TypeOf((*MockSystemVMServiceIface)(nil).GetSystemVmByName), varargs...)
}

// GetSystemVmID mocks base method.
func (m *MockSystemVMServiceIface) GetSystemVmID(name string, opts ...OptionFunc) (string, int, error) {
	m.ctrl.T.Helper()
	varargs := []any{name}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetSystemVmID", varargs...)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(int)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetSystemVmID indicates an expected call of GetSystemVmID.
func (mr *MockSystemVMServiceIfaceMockRecorder) GetSystemVmID(name any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{name}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSystemVmID", reflect.TypeOf((*MockSystemVMServiceIface)(nil).GetSystemVmID), varargs...)
}

// GetSystemVmsUsageHistoryByID mocks base method.
func (m *MockSystemVMServiceIface) GetSystemVmsUsageHistoryByID(id string, opts ...OptionFunc) (*SystemVmsUsageHistory, int, error) {
	m.ctrl.T.Helper()
	varargs := []any{id}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetSystemVmsUsageHistoryByID", varargs...)
	ret0, _ := ret[0].(*SystemVmsUsageHistory)
	ret1, _ := ret[1].(int)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetSystemVmsUsageHistoryByID indicates an expected call of GetSystemVmsUsageHistoryByID.
func (mr *MockSystemVMServiceIfaceMockRecorder) GetSystemVmsUsageHistoryByID(id any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{id}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSystemVmsUsageHistoryByID", reflect.TypeOf((*MockSystemVMServiceIface)(nil).GetSystemVmsUsageHistoryByID), varargs...)
}

// GetSystemVmsUsageHistoryByName mocks base method.
func (m *MockSystemVMServiceIface) GetSystemVmsUsageHistoryByName(name string, opts ...OptionFunc) (*SystemVmsUsageHistory, int, error) {
	m.ctrl.T.Helper()
	varargs := []any{name}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetSystemVmsUsageHistoryByName", varargs...)
	ret0, _ := ret[0].(*SystemVmsUsageHistory)
	ret1, _ := ret[1].(int)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetSystemVmsUsageHistoryByName indicates an expected call of GetSystemVmsUsageHistoryByName.
func (mr *MockSystemVMServiceIfaceMockRecorder) GetSystemVmsUsageHistoryByName(name any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{name}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSystemVmsUsageHistoryByName", reflect.TypeOf((*MockSystemVMServiceIface)(nil).GetSystemVmsUsageHistoryByName), varargs...)
}

// GetSystemVmsUsageHistoryID mocks base method.
func (m *MockSystemVMServiceIface) GetSystemVmsUsageHistoryID(name string, opts ...OptionFunc) (string, int, error) {
	m.ctrl.T.Helper()
	varargs := []any{name}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetSystemVmsUsageHistoryID", varargs...)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(int)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetSystemVmsUsageHistoryID indicates an expected call of GetSystemVmsUsageHistoryID.
func (mr *MockSystemVMServiceIfaceMockRecorder) GetSystemVmsUsageHistoryID(name any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{name}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSystemVmsUsageHistoryID", reflect.TypeOf((*MockSystemVMServiceIface)(nil).GetSystemVmsUsageHistoryID), varargs...)
}

// ListSystemVms mocks base method.
func (m *MockSystemVMServiceIface) ListSystemVms(p *ListSystemVmsParams) (*ListSystemVmsResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListSystemVms", p)
	ret0, _ := ret[0].(*ListSystemVmsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListSystemVms indicates an expected call of ListSystemVms.
func (mr *MockSystemVMServiceIfaceMockRecorder) ListSystemVms(p any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListSystemVms", reflect.TypeOf((*MockSystemVMServiceIface)(nil).ListSystemVms), p)
}

// ListSystemVmsUsageHistory mocks base method.
func (m *MockSystemVMServiceIface) ListSystemVmsUsageHistory(p *ListSystemVmsUsageHistoryParams) (*ListSystemVmsUsageHistoryResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListSystemVmsUsageHistory", p)
	ret0, _ := ret[0].(*ListSystemVmsUsageHistoryResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListSystemVmsUsageHistory indicates an expected call of ListSystemVmsUsageHistory.
func (mr *MockSystemVMServiceIfaceMockRecorder) ListSystemVmsUsageHistory(p any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListSystemVmsUsageHistory", reflect.TypeOf((*MockSystemVMServiceIface)(nil).ListSystemVmsUsageHistory), p)
}

// MigrateSystemVm mocks base method.
func (m *MockSystemVMServiceIface) MigrateSystemVm(p *MigrateSystemVmParams) (*MigrateSystemVmResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MigrateSystemVm", p)
	ret0, _ := ret[0].(*MigrateSystemVmResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MigrateSystemVm indicates an expected call of MigrateSystemVm.
func (mr *MockSystemVMServiceIfaceMockRecorder) MigrateSystemVm(p any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MigrateSystemVm", reflect.TypeOf((*MockSystemVMServiceIface)(nil).MigrateSystemVm), p)
}

// NewChangeServiceForSystemVmParams mocks base method.
func (m *MockSystemVMServiceIface) NewChangeServiceForSystemVmParams(id, serviceofferingid string) *ChangeServiceForSystemVmParams {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewChangeServiceForSystemVmParams", id, serviceofferingid)
	ret0, _ := ret[0].(*ChangeServiceForSystemVmParams)
	return ret0
}

// NewChangeServiceForSystemVmParams indicates an expected call of NewChangeServiceForSystemVmParams.
func (mr *MockSystemVMServiceIfaceMockRecorder) NewChangeServiceForSystemVmParams(id, serviceofferingid any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewChangeServiceForSystemVmParams", reflect.TypeOf((*MockSystemVMServiceIface)(nil).NewChangeServiceForSystemVmParams), id, serviceofferingid)
}

// NewDestroySystemVmParams mocks base method.
func (m *MockSystemVMServiceIface) NewDestroySystemVmParams(id string) *DestroySystemVmParams {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewDestroySystemVmParams", id)
	ret0, _ := ret[0].(*DestroySystemVmParams)
	return ret0
}

// NewDestroySystemVmParams indicates an expected call of NewDestroySystemVmParams.
func (mr *MockSystemVMServiceIfaceMockRecorder) NewDestroySystemVmParams(id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewDestroySystemVmParams", reflect.TypeOf((*MockSystemVMServiceIface)(nil).NewDestroySystemVmParams), id)
}

// NewListSystemVmsParams mocks base method.
func (m *MockSystemVMServiceIface) NewListSystemVmsParams() *ListSystemVmsParams {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewListSystemVmsParams")
	ret0, _ := ret[0].(*ListSystemVmsParams)
	return ret0
}

// NewListSystemVmsParams indicates an expected call of NewListSystemVmsParams.
func (mr *MockSystemVMServiceIfaceMockRecorder) NewListSystemVmsParams() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewListSystemVmsParams", reflect.TypeOf((*MockSystemVMServiceIface)(nil).NewListSystemVmsParams))
}

// NewListSystemVmsUsageHistoryParams mocks base method.
func (m *MockSystemVMServiceIface) NewListSystemVmsUsageHistoryParams() *ListSystemVmsUsageHistoryParams {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewListSystemVmsUsageHistoryParams")
	ret0, _ := ret[0].(*ListSystemVmsUsageHistoryParams)
	return ret0
}

// NewListSystemVmsUsageHistoryParams indicates an expected call of NewListSystemVmsUsageHistoryParams.
func (mr *MockSystemVMServiceIfaceMockRecorder) NewListSystemVmsUsageHistoryParams() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewListSystemVmsUsageHistoryParams", reflect.TypeOf((*MockSystemVMServiceIface)(nil).NewListSystemVmsUsageHistoryParams))
}

// NewMigrateSystemVmParams mocks base method.
func (m *MockSystemVMServiceIface) NewMigrateSystemVmParams(virtualmachineid string) *MigrateSystemVmParams {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewMigrateSystemVmParams", virtualmachineid)
	ret0, _ := ret[0].(*MigrateSystemVmParams)
	return ret0
}

// NewMigrateSystemVmParams indicates an expected call of NewMigrateSystemVmParams.
func (mr *MockSystemVMServiceIfaceMockRecorder) NewMigrateSystemVmParams(virtualmachineid any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewMigrateSystemVmParams", reflect.TypeOf((*MockSystemVMServiceIface)(nil).NewMigrateSystemVmParams), virtualmachineid)
}

// NewPatchSystemVmParams mocks base method.
func (m *MockSystemVMServiceIface) NewPatchSystemVmParams() *PatchSystemVmParams {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewPatchSystemVmParams")
	ret0, _ := ret[0].(*PatchSystemVmParams)
	return ret0
}

// NewPatchSystemVmParams indicates an expected call of NewPatchSystemVmParams.
func (mr *MockSystemVMServiceIfaceMockRecorder) NewPatchSystemVmParams() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewPatchSystemVmParams", reflect.TypeOf((*MockSystemVMServiceIface)(nil).NewPatchSystemVmParams))
}

// NewRebootSystemVmParams mocks base method.
func (m *MockSystemVMServiceIface) NewRebootSystemVmParams(id string) *RebootSystemVmParams {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewRebootSystemVmParams", id)
	ret0, _ := ret[0].(*RebootSystemVmParams)
	return ret0
}

// NewRebootSystemVmParams indicates an expected call of NewRebootSystemVmParams.
func (mr *MockSystemVMServiceIfaceMockRecorder) NewRebootSystemVmParams(id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewRebootSystemVmParams", reflect.TypeOf((*MockSystemVMServiceIface)(nil).NewRebootSystemVmParams), id)
}

// NewScaleSystemVmParams mocks base method.
func (m *MockSystemVMServiceIface) NewScaleSystemVmParams(id, serviceofferingid string) *ScaleSystemVmParams {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewScaleSystemVmParams", id, serviceofferingid)
	ret0, _ := ret[0].(*ScaleSystemVmParams)
	return ret0
}

// NewScaleSystemVmParams indicates an expected call of NewScaleSystemVmParams.
func (mr *MockSystemVMServiceIfaceMockRecorder) NewScaleSystemVmParams(id, serviceofferingid any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewScaleSystemVmParams", reflect.TypeOf((*MockSystemVMServiceIface)(nil).NewScaleSystemVmParams), id, serviceofferingid)
}

// NewStartSystemVmParams mocks base method.
func (m *MockSystemVMServiceIface) NewStartSystemVmParams(id string) *StartSystemVmParams {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewStartSystemVmParams", id)
	ret0, _ := ret[0].(*StartSystemVmParams)
	return ret0
}

// NewStartSystemVmParams indicates an expected call of NewStartSystemVmParams.
func (mr *MockSystemVMServiceIfaceMockRecorder) NewStartSystemVmParams(id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewStartSystemVmParams", reflect.TypeOf((*MockSystemVMServiceIface)(nil).NewStartSystemVmParams), id)
}

// NewStopSystemVmParams mocks base method.
func (m *MockSystemVMServiceIface) NewStopSystemVmParams(id string) *StopSystemVmParams {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewStopSystemVmParams", id)
	ret0, _ := ret[0].(*StopSystemVmParams)
	return ret0
}

// NewStopSystemVmParams indicates an expected call of NewStopSystemVmParams.
func (mr *MockSystemVMServiceIfaceMockRecorder) NewStopSystemVmParams(id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewStopSystemVmParams", reflect.TypeOf((*MockSystemVMServiceIface)(nil).NewStopSystemVmParams), id)
}

// PatchSystemVm mocks base method.
func (m *MockSystemVMServiceIface) PatchSystemVm(p *PatchSystemVmParams) (*PatchSystemVmResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PatchSystemVm", p)
	ret0, _ := ret[0].(*PatchSystemVmResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PatchSystemVm indicates an expected call of PatchSystemVm.
func (mr *MockSystemVMServiceIfaceMockRecorder) PatchSystemVm(p any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PatchSystemVm", reflect.TypeOf((*MockSystemVMServiceIface)(nil).PatchSystemVm), p)
}

// RebootSystemVm mocks base method.
func (m *MockSystemVMServiceIface) RebootSystemVm(p *RebootSystemVmParams) (*RebootSystemVmResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RebootSystemVm", p)
	ret0, _ := ret[0].(*RebootSystemVmResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RebootSystemVm indicates an expected call of RebootSystemVm.
func (mr *MockSystemVMServiceIfaceMockRecorder) RebootSystemVm(p any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RebootSystemVm", reflect.TypeOf((*MockSystemVMServiceIface)(nil).RebootSystemVm), p)
}

// ScaleSystemVm mocks base method.
func (m *MockSystemVMServiceIface) ScaleSystemVm(p *ScaleSystemVmParams) (*ScaleSystemVmResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ScaleSystemVm", p)
	ret0, _ := ret[0].(*ScaleSystemVmResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ScaleSystemVm indicates an expected call of ScaleSystemVm.
func (mr *MockSystemVMServiceIfaceMockRecorder) ScaleSystemVm(p any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ScaleSystemVm", reflect.TypeOf((*MockSystemVMServiceIface)(nil).ScaleSystemVm), p)
}

// StartSystemVm mocks base method.
func (m *MockSystemVMServiceIface) StartSystemVm(p *StartSystemVmParams) (*StartSystemVmResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StartSystemVm", p)
	ret0, _ := ret[0].(*StartSystemVmResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StartSystemVm indicates an expected call of StartSystemVm.
func (mr *MockSystemVMServiceIfaceMockRecorder) StartSystemVm(p any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StartSystemVm", reflect.TypeOf((*MockSystemVMServiceIface)(nil).StartSystemVm), p)
}

// StopSystemVm mocks base method.
func (m *MockSystemVMServiceIface) StopSystemVm(p *StopSystemVmParams) (*StopSystemVmResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StopSystemVm", p)
	ret0, _ := ret[0].(*StopSystemVmResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StopSystemVm indicates an expected call of StopSystemVm.
func (mr *MockSystemVMServiceIfaceMockRecorder) StopSystemVm(p any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StopSystemVm", reflect.TypeOf((*MockSystemVMServiceIface)(nil).StopSystemVm), p)
}
