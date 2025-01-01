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
// Source: ./cloudstack/CloudIdentifierService.go
//
// Generated by this command:
//
//	mockgen -destination=./cloudstack/CloudIdentifierService_mock.go -package=cloudstack -copyright_file=header.txt -source=./cloudstack/CloudIdentifierService.go
//

// Package cloudstack is a generated GoMock package.
package cloudstack

import (
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockCloudIdentifierServiceIface is a mock of CloudIdentifierServiceIface interface.
type MockCloudIdentifierServiceIface struct {
	ctrl     *gomock.Controller
	recorder *MockCloudIdentifierServiceIfaceMockRecorder
	isgomock struct{}
}

// MockCloudIdentifierServiceIfaceMockRecorder is the mock recorder for MockCloudIdentifierServiceIface.
type MockCloudIdentifierServiceIfaceMockRecorder struct {
	mock *MockCloudIdentifierServiceIface
}

// NewMockCloudIdentifierServiceIface creates a new mock instance.
func NewMockCloudIdentifierServiceIface(ctrl *gomock.Controller) *MockCloudIdentifierServiceIface {
	mock := &MockCloudIdentifierServiceIface{ctrl: ctrl}
	mock.recorder = &MockCloudIdentifierServiceIfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCloudIdentifierServiceIface) EXPECT() *MockCloudIdentifierServiceIfaceMockRecorder {
	return m.recorder
}

// GetCloudIdentifier mocks base method.
func (m *MockCloudIdentifierServiceIface) GetCloudIdentifier(p *GetCloudIdentifierParams) (*GetCloudIdentifierResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCloudIdentifier", p)
	ret0, _ := ret[0].(*GetCloudIdentifierResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCloudIdentifier indicates an expected call of GetCloudIdentifier.
func (mr *MockCloudIdentifierServiceIfaceMockRecorder) GetCloudIdentifier(p any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCloudIdentifier", reflect.TypeOf((*MockCloudIdentifierServiceIface)(nil).GetCloudIdentifier), p)
}

// NewGetCloudIdentifierParams mocks base method.
func (m *MockCloudIdentifierServiceIface) NewGetCloudIdentifierParams(userid string) *GetCloudIdentifierParams {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewGetCloudIdentifierParams", userid)
	ret0, _ := ret[0].(*GetCloudIdentifierParams)
	return ret0
}

// NewGetCloudIdentifierParams indicates an expected call of NewGetCloudIdentifierParams.
func (mr *MockCloudIdentifierServiceIfaceMockRecorder) NewGetCloudIdentifierParams(userid any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewGetCloudIdentifierParams", reflect.TypeOf((*MockCloudIdentifierServiceIface)(nil).NewGetCloudIdentifierParams), userid)
}
