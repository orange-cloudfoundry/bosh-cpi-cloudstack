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
// Source: ./cloudstack/ResourcetagsService.go
//
// Generated by this command:
//
//	mockgen -destination=./cloudstack/ResourcetagsService_mock.go -package=cloudstack -copyright_file=header.txt -source=./cloudstack/ResourcetagsService.go
//

// Package cloudstack is a generated GoMock package.
package cloudstack

import (
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockResourcetagsServiceIface is a mock of ResourcetagsServiceIface interface.
type MockResourcetagsServiceIface struct {
	ctrl     *gomock.Controller
	recorder *MockResourcetagsServiceIfaceMockRecorder
	isgomock struct{}
}

// MockResourcetagsServiceIfaceMockRecorder is the mock recorder for MockResourcetagsServiceIface.
type MockResourcetagsServiceIfaceMockRecorder struct {
	mock *MockResourcetagsServiceIface
}

// NewMockResourcetagsServiceIface creates a new mock instance.
func NewMockResourcetagsServiceIface(ctrl *gomock.Controller) *MockResourcetagsServiceIface {
	mock := &MockResourcetagsServiceIface{ctrl: ctrl}
	mock.recorder = &MockResourcetagsServiceIfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockResourcetagsServiceIface) EXPECT() *MockResourcetagsServiceIfaceMockRecorder {
	return m.recorder
}

// CreateTags mocks base method.
func (m *MockResourcetagsServiceIface) CreateTags(p *CreateTagsParams) (*CreateTagsResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateTags", p)
	ret0, _ := ret[0].(*CreateTagsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateTags indicates an expected call of CreateTags.
func (mr *MockResourcetagsServiceIfaceMockRecorder) CreateTags(p any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateTags", reflect.TypeOf((*MockResourcetagsServiceIface)(nil).CreateTags), p)
}

// DeleteTags mocks base method.
func (m *MockResourcetagsServiceIface) DeleteTags(p *DeleteTagsParams) (*DeleteTagsResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteTags", p)
	ret0, _ := ret[0].(*DeleteTagsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteTags indicates an expected call of DeleteTags.
func (mr *MockResourcetagsServiceIfaceMockRecorder) DeleteTags(p any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteTags", reflect.TypeOf((*MockResourcetagsServiceIface)(nil).DeleteTags), p)
}

// GetStorageTagID mocks base method.
func (m *MockResourcetagsServiceIface) GetStorageTagID(keyword string, opts ...OptionFunc) (string, int, error) {
	m.ctrl.T.Helper()
	varargs := []any{keyword}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetStorageTagID", varargs...)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(int)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetStorageTagID indicates an expected call of GetStorageTagID.
func (mr *MockResourcetagsServiceIfaceMockRecorder) GetStorageTagID(keyword any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{keyword}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetStorageTagID", reflect.TypeOf((*MockResourcetagsServiceIface)(nil).GetStorageTagID), varargs...)
}

// ListStorageTags mocks base method.
func (m *MockResourcetagsServiceIface) ListStorageTags(p *ListStorageTagsParams) (*ListStorageTagsResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListStorageTags", p)
	ret0, _ := ret[0].(*ListStorageTagsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListStorageTags indicates an expected call of ListStorageTags.
func (mr *MockResourcetagsServiceIfaceMockRecorder) ListStorageTags(p any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListStorageTags", reflect.TypeOf((*MockResourcetagsServiceIface)(nil).ListStorageTags), p)
}

// ListTags mocks base method.
func (m *MockResourcetagsServiceIface) ListTags(p *ListTagsParams) (*ListTagsResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListTags", p)
	ret0, _ := ret[0].(*ListTagsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListTags indicates an expected call of ListTags.
func (mr *MockResourcetagsServiceIfaceMockRecorder) ListTags(p any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListTags", reflect.TypeOf((*MockResourcetagsServiceIface)(nil).ListTags), p)
}

// NewCreateTagsParams mocks base method.
func (m *MockResourcetagsServiceIface) NewCreateTagsParams(resourceids []string, resourcetype string, tags map[string]string) *CreateTagsParams {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewCreateTagsParams", resourceids, resourcetype, tags)
	ret0, _ := ret[0].(*CreateTagsParams)
	return ret0
}

// NewCreateTagsParams indicates an expected call of NewCreateTagsParams.
func (mr *MockResourcetagsServiceIfaceMockRecorder) NewCreateTagsParams(resourceids, resourcetype, tags any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewCreateTagsParams", reflect.TypeOf((*MockResourcetagsServiceIface)(nil).NewCreateTagsParams), resourceids, resourcetype, tags)
}

// NewDeleteTagsParams mocks base method.
func (m *MockResourcetagsServiceIface) NewDeleteTagsParams(resourceids []string, resourcetype string) *DeleteTagsParams {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewDeleteTagsParams", resourceids, resourcetype)
	ret0, _ := ret[0].(*DeleteTagsParams)
	return ret0
}

// NewDeleteTagsParams indicates an expected call of NewDeleteTagsParams.
func (mr *MockResourcetagsServiceIfaceMockRecorder) NewDeleteTagsParams(resourceids, resourcetype any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewDeleteTagsParams", reflect.TypeOf((*MockResourcetagsServiceIface)(nil).NewDeleteTagsParams), resourceids, resourcetype)
}

// NewListStorageTagsParams mocks base method.
func (m *MockResourcetagsServiceIface) NewListStorageTagsParams() *ListStorageTagsParams {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewListStorageTagsParams")
	ret0, _ := ret[0].(*ListStorageTagsParams)
	return ret0
}

// NewListStorageTagsParams indicates an expected call of NewListStorageTagsParams.
func (mr *MockResourcetagsServiceIfaceMockRecorder) NewListStorageTagsParams() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewListStorageTagsParams", reflect.TypeOf((*MockResourcetagsServiceIface)(nil).NewListStorageTagsParams))
}

// NewListTagsParams mocks base method.
func (m *MockResourcetagsServiceIface) NewListTagsParams() *ListTagsParams {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewListTagsParams")
	ret0, _ := ret[0].(*ListTagsParams)
	return ret0
}

// NewListTagsParams indicates an expected call of NewListTagsParams.
func (mr *MockResourcetagsServiceIfaceMockRecorder) NewListTagsParams() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewListTagsParams", reflect.TypeOf((*MockResourcetagsServiceIface)(nil).NewListTagsParams))
}
