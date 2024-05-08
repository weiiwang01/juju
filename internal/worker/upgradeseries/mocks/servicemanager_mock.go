// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/juju/juju/internal/service (interfaces: SystemdServiceManager)
//
// Generated by this command:
//
//	mockgen -typed -package mocks -destination mocks/servicemanager_mock.go github.com/juju/juju/internal/service SystemdServiceManager
//

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	common "github.com/juju/juju/internal/service/common"
	gomock "go.uber.org/mock/gomock"
)

// MockSystemdServiceManager is a mock of SystemdServiceManager interface.
type MockSystemdServiceManager struct {
	ctrl     *gomock.Controller
	recorder *MockSystemdServiceManagerMockRecorder
}

// MockSystemdServiceManagerMockRecorder is the mock recorder for MockSystemdServiceManager.
type MockSystemdServiceManagerMockRecorder struct {
	mock *MockSystemdServiceManager
}

// NewMockSystemdServiceManager creates a new mock instance.
func NewMockSystemdServiceManager(ctrl *gomock.Controller) *MockSystemdServiceManager {
	mock := &MockSystemdServiceManager{ctrl: ctrl}
	mock.recorder = &MockSystemdServiceManagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSystemdServiceManager) EXPECT() *MockSystemdServiceManagerMockRecorder {
	return m.recorder
}

// CreateAgentConf mocks base method.
func (m *MockSystemdServiceManager) CreateAgentConf(arg0, arg1 string) (common.Conf, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateAgentConf", arg0, arg1)
	ret0, _ := ret[0].(common.Conf)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateAgentConf indicates an expected call of CreateAgentConf.
func (mr *MockSystemdServiceManagerMockRecorder) CreateAgentConf(arg0, arg1 any) *MockSystemdServiceManagerCreateAgentConfCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateAgentConf", reflect.TypeOf((*MockSystemdServiceManager)(nil).CreateAgentConf), arg0, arg1)
	return &MockSystemdServiceManagerCreateAgentConfCall{Call: call}
}

// MockSystemdServiceManagerCreateAgentConfCall wrap *gomock.Call
type MockSystemdServiceManagerCreateAgentConfCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockSystemdServiceManagerCreateAgentConfCall) Return(arg0 common.Conf, arg1 error) *MockSystemdServiceManagerCreateAgentConfCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockSystemdServiceManagerCreateAgentConfCall) Do(f func(string, string) (common.Conf, error)) *MockSystemdServiceManagerCreateAgentConfCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockSystemdServiceManagerCreateAgentConfCall) DoAndReturn(f func(string, string) (common.Conf, error)) *MockSystemdServiceManagerCreateAgentConfCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// FindAgents mocks base method.
func (m *MockSystemdServiceManager) FindAgents(arg0 string) (string, []string, []string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindAgents", arg0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].([]string)
	ret2, _ := ret[2].([]string)
	ret3, _ := ret[3].(error)
	return ret0, ret1, ret2, ret3
}

// FindAgents indicates an expected call of FindAgents.
func (mr *MockSystemdServiceManagerMockRecorder) FindAgents(arg0 any) *MockSystemdServiceManagerFindAgentsCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindAgents", reflect.TypeOf((*MockSystemdServiceManager)(nil).FindAgents), arg0)
	return &MockSystemdServiceManagerFindAgentsCall{Call: call}
}

// MockSystemdServiceManagerFindAgentsCall wrap *gomock.Call
type MockSystemdServiceManagerFindAgentsCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockSystemdServiceManagerFindAgentsCall) Return(arg0 string, arg1, arg2 []string, arg3 error) *MockSystemdServiceManagerFindAgentsCall {
	c.Call = c.Call.Return(arg0, arg1, arg2, arg3)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockSystemdServiceManagerFindAgentsCall) Do(f func(string) (string, []string, []string, error)) *MockSystemdServiceManagerFindAgentsCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockSystemdServiceManagerFindAgentsCall) DoAndReturn(f func(string) (string, []string, []string, error)) *MockSystemdServiceManagerFindAgentsCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// WriteServiceFile mocks base method.
func (m *MockSystemdServiceManager) WriteServiceFile() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WriteServiceFile")
	ret0, _ := ret[0].(error)
	return ret0
}

// WriteServiceFile indicates an expected call of WriteServiceFile.
func (mr *MockSystemdServiceManagerMockRecorder) WriteServiceFile() *MockSystemdServiceManagerWriteServiceFileCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WriteServiceFile", reflect.TypeOf((*MockSystemdServiceManager)(nil).WriteServiceFile))
	return &MockSystemdServiceManagerWriteServiceFileCall{Call: call}
}

// MockSystemdServiceManagerWriteServiceFileCall wrap *gomock.Call
type MockSystemdServiceManagerWriteServiceFileCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockSystemdServiceManagerWriteServiceFileCall) Return(arg0 error) *MockSystemdServiceManagerWriteServiceFileCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockSystemdServiceManagerWriteServiceFileCall) Do(f func() error) *MockSystemdServiceManagerWriteServiceFileCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockSystemdServiceManagerWriteServiceFileCall) DoAndReturn(f func() error) *MockSystemdServiceManagerWriteServiceFileCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// WriteSystemdAgent mocks base method.
func (m *MockSystemdServiceManager) WriteSystemdAgent(arg0, arg1, arg2 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WriteSystemdAgent", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// WriteSystemdAgent indicates an expected call of WriteSystemdAgent.
func (mr *MockSystemdServiceManagerMockRecorder) WriteSystemdAgent(arg0, arg1, arg2 any) *MockSystemdServiceManagerWriteSystemdAgentCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WriteSystemdAgent", reflect.TypeOf((*MockSystemdServiceManager)(nil).WriteSystemdAgent), arg0, arg1, arg2)
	return &MockSystemdServiceManagerWriteSystemdAgentCall{Call: call}
}

// MockSystemdServiceManagerWriteSystemdAgentCall wrap *gomock.Call
type MockSystemdServiceManagerWriteSystemdAgentCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockSystemdServiceManagerWriteSystemdAgentCall) Return(arg0 error) *MockSystemdServiceManagerWriteSystemdAgentCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockSystemdServiceManagerWriteSystemdAgentCall) Do(f func(string, string, string) error) *MockSystemdServiceManagerWriteSystemdAgentCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockSystemdServiceManagerWriteSystemdAgentCall) DoAndReturn(f func(string, string, string) error) *MockSystemdServiceManagerWriteSystemdAgentCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}
