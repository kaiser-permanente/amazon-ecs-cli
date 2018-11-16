// Copyright 2015-2017 Amazon.com, Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//     http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/aws/amazon-ecs-cli/ecs-cli/modules/clients/aws/ecs (interfaces: ECSClient)

// Package mock_ecs is a generated GoMock package.
package mock_ecs

import (
	reflect "reflect"

	ecs "github.com/aws/amazon-ecs-cli/ecs-cli/modules/clients/aws/ecs"
	cache "github.com/aws/amazon-ecs-cli/ecs-cli/modules/utils/cache"
	ecs0 "github.com/aws/aws-sdk-go/service/ecs"
	gomock "github.com/golang/mock/gomock"
)

// MockECSClient is a mock of ECSClient interface
type MockECSClient struct {
	ctrl     *gomock.Controller
	recorder *MockECSClientMockRecorder
}

// MockECSClientMockRecorder is the mock recorder for MockECSClient
type MockECSClientMockRecorder struct {
	mock *MockECSClient
}

// NewMockECSClient creates a new mock instance
func NewMockECSClient(ctrl *gomock.Controller) *MockECSClient {
	mock := &MockECSClient{ctrl: ctrl}
	mock.recorder = &MockECSClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockECSClient) EXPECT() *MockECSClientMockRecorder {
	return m.recorder
}

// CreateCluster mocks base method
func (m *MockECSClient) CreateCluster(arg0 string) (string, error) {
	ret := m.ctrl.Call(m, "CreateCluster", arg0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateCluster indicates an expected call of CreateCluster
func (mr *MockECSClientMockRecorder) CreateCluster(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateCluster", reflect.TypeOf((*MockECSClient)(nil).CreateCluster), arg0)
}

// CreateService mocks base method
func (m *MockECSClient) CreateService(arg0 string, arg1 *ecs0.CreateServiceInput) error {
	ret := m.ctrl.Call(m, "CreateService", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateService indicates an expected call of CreateService
func (mr *MockECSClientMockRecorder) CreateService(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateService", reflect.TypeOf((*MockECSClient)(nil).CreateService), arg0, arg1)
}

// DeleteCluster mocks base method
func (m *MockECSClient) DeleteCluster(arg0 string) (string, error) {
	ret := m.ctrl.Call(m, "DeleteCluster", arg0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteCluster indicates an expected call of DeleteCluster
func (mr *MockECSClientMockRecorder) DeleteCluster(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteCluster", reflect.TypeOf((*MockECSClient)(nil).DeleteCluster), arg0)
}

// DeleteService mocks base method
func (m *MockECSClient) DeleteService(arg0 string) error {
	ret := m.ctrl.Call(m, "DeleteService", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteService indicates an expected call of DeleteService
func (mr *MockECSClientMockRecorder) DeleteService(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteService", reflect.TypeOf((*MockECSClient)(nil).DeleteService), arg0)
}

// DescribeService mocks base method
func (m *MockECSClient) DescribeService(arg0 string) (*ecs0.DescribeServicesOutput, error) {
	ret := m.ctrl.Call(m, "DescribeService", arg0)
	ret0, _ := ret[0].(*ecs0.DescribeServicesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeService indicates an expected call of DescribeService
func (mr *MockECSClientMockRecorder) DescribeService(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeService", reflect.TypeOf((*MockECSClient)(nil).DescribeService), arg0)
}

// DescribeTaskDefinition mocks base method
func (m *MockECSClient) DescribeTaskDefinition(arg0 string) (*ecs0.TaskDefinition, error) {
	ret := m.ctrl.Call(m, "DescribeTaskDefinition", arg0)
	ret0, _ := ret[0].(*ecs0.TaskDefinition)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeTaskDefinition indicates an expected call of DescribeTaskDefinition
func (mr *MockECSClientMockRecorder) DescribeTaskDefinition(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeTaskDefinition", reflect.TypeOf((*MockECSClient)(nil).DescribeTaskDefinition), arg0)
}

// DescribeTasks mocks base method
func (m *MockECSClient) DescribeTasks(arg0 []*string) ([]*ecs0.Task, error) {
	ret := m.ctrl.Call(m, "DescribeTasks", arg0)
	ret0, _ := ret[0].([]*ecs0.Task)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeTasks indicates an expected call of DescribeTasks
func (mr *MockECSClientMockRecorder) DescribeTasks(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeTasks", reflect.TypeOf((*MockECSClient)(nil).DescribeTasks), arg0)
}

// GetEC2InstanceIDs mocks base method
func (m *MockECSClient) GetEC2InstanceIDs(arg0 []*string) (map[string]string, error) {
	ret := m.ctrl.Call(m, "GetEC2InstanceIDs", arg0)
	ret0, _ := ret[0].(map[string]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetEC2InstanceIDs indicates an expected call of GetEC2InstanceIDs
func (mr *MockECSClientMockRecorder) GetEC2InstanceIDs(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEC2InstanceIDs", reflect.TypeOf((*MockECSClient)(nil).GetEC2InstanceIDs), arg0)
}

// GetTasksPages mocks base method
func (m *MockECSClient) GetTasksPages(arg0 *ecs0.ListTasksInput, arg1 ecs.ProcessTasksAction) error {
	ret := m.ctrl.Call(m, "GetTasksPages", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// GetTasksPages indicates an expected call of GetTasksPages
func (mr *MockECSClientMockRecorder) GetTasksPages(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTasksPages", reflect.TypeOf((*MockECSClient)(nil).GetTasksPages), arg0, arg1)
}

// IsActiveCluster mocks base method
func (m *MockECSClient) IsActiveCluster(arg0 string) (bool, error) {
	ret := m.ctrl.Call(m, "IsActiveCluster", arg0)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IsActiveCluster indicates an expected call of IsActiveCluster
func (mr *MockECSClientMockRecorder) IsActiveCluster(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsActiveCluster", reflect.TypeOf((*MockECSClient)(nil).IsActiveCluster), arg0)
}

// RegisterTaskDefinition mocks base method
func (m *MockECSClient) RegisterTaskDefinition(arg0 *ecs0.RegisterTaskDefinitionInput) (*ecs0.TaskDefinition, error) {
	ret := m.ctrl.Call(m, "RegisterTaskDefinition", arg0)
	ret0, _ := ret[0].(*ecs0.TaskDefinition)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RegisterTaskDefinition indicates an expected call of RegisterTaskDefinition
func (mr *MockECSClientMockRecorder) RegisterTaskDefinition(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterTaskDefinition", reflect.TypeOf((*MockECSClient)(nil).RegisterTaskDefinition), arg0)
}

// RegisterTaskDefinitionIfNeeded mocks base method
func (m *MockECSClient) RegisterTaskDefinitionIfNeeded(arg0 *ecs0.RegisterTaskDefinitionInput, arg1 cache.Cache) (*ecs0.TaskDefinition, error) {
	ret := m.ctrl.Call(m, "RegisterTaskDefinitionIfNeeded", arg0, arg1)
	ret0, _ := ret[0].(*ecs0.TaskDefinition)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RegisterTaskDefinitionIfNeeded indicates an expected call of RegisterTaskDefinitionIfNeeded
func (mr *MockECSClientMockRecorder) RegisterTaskDefinitionIfNeeded(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterTaskDefinitionIfNeeded", reflect.TypeOf((*MockECSClient)(nil).RegisterTaskDefinitionIfNeeded), arg0, arg1)
}

// RunTask mocks base method
func (m *MockECSClient) RunTask(arg0 *ecs0.RunTaskInput) (*ecs0.RunTaskOutput, error) {
	ret := m.ctrl.Call(m, "RunTask", arg0)
	ret0, _ := ret[0].(*ecs0.RunTaskOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RunTask indicates an expected call of RunTask
func (mr *MockECSClientMockRecorder) RunTask(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RunTask", reflect.TypeOf((*MockECSClient)(nil).RunTask), arg0)
}

// StopTask mocks base method
func (m *MockECSClient) StopTask(arg0 string) error {
	ret := m.ctrl.Call(m, "StopTask", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// StopTask indicates an expected call of StopTask
func (mr *MockECSClientMockRecorder) StopTask(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StopTask", reflect.TypeOf((*MockECSClient)(nil).StopTask), arg0)
}

// UpdateService mocks base method
func (m *MockECSClient) UpdateService(arg0, arg1 string, arg2 int64, arg3 *ecs0.DeploymentConfiguration, arg4 *ecs0.NetworkConfiguration, arg5 *int64, arg6 bool) error {
	ret := m.ctrl.Call(m, "UpdateService", arg0, arg1, arg2, arg3, arg4, arg5, arg6)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateService indicates an expected call of UpdateService
func (mr *MockECSClientMockRecorder) UpdateService(arg0, arg1, arg2, arg3, arg4, arg5, arg6 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateService", reflect.TypeOf((*MockECSClient)(nil).UpdateService), arg0, arg1, arg2, arg3, arg4, arg5, arg6)
}
