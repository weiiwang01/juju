// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/juju/juju/cmd/juju/commands (interfaces: ControllerAPI)
//
// Generated by this command:
//
//	mockgen -typed -package mocks -destination mocks/controller_mock.go github.com/juju/juju/cmd/juju/commands ControllerAPI
//

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	controller "github.com/juju/juju/controller"
	cloudspec "github.com/juju/juju/environs/cloudspec"
	names "github.com/juju/names/v5"
	gomock "go.uber.org/mock/gomock"
)

// MockControllerAPI is a mock of ControllerAPI interface.
type MockControllerAPI struct {
	ctrl     *gomock.Controller
	recorder *MockControllerAPIMockRecorder
}

// MockControllerAPIMockRecorder is the mock recorder for MockControllerAPI.
type MockControllerAPIMockRecorder struct {
	mock *MockControllerAPI
}

// NewMockControllerAPI creates a new mock instance.
func NewMockControllerAPI(ctrl *gomock.Controller) *MockControllerAPI {
	mock := &MockControllerAPI{ctrl: ctrl}
	mock.recorder = &MockControllerAPIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockControllerAPI) EXPECT() *MockControllerAPIMockRecorder {
	return m.recorder
}

// Close mocks base method.
func (m *MockControllerAPI) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockControllerAPIMockRecorder) Close() *MockControllerAPICloseCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockControllerAPI)(nil).Close))
	return &MockControllerAPICloseCall{Call: call}
}

// MockControllerAPICloseCall wrap *gomock.Call
type MockControllerAPICloseCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockControllerAPICloseCall) Return(arg0 error) *MockControllerAPICloseCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockControllerAPICloseCall) Do(f func() error) *MockControllerAPICloseCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockControllerAPICloseCall) DoAndReturn(f func() error) *MockControllerAPICloseCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// CloudSpec mocks base method.
func (m *MockControllerAPI) CloudSpec(arg0 names.ModelTag) (cloudspec.CloudSpec, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CloudSpec", arg0)
	ret0, _ := ret[0].(cloudspec.CloudSpec)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CloudSpec indicates an expected call of CloudSpec.
func (mr *MockControllerAPIMockRecorder) CloudSpec(arg0 any) *MockControllerAPICloudSpecCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CloudSpec", reflect.TypeOf((*MockControllerAPI)(nil).CloudSpec), arg0)
	return &MockControllerAPICloudSpecCall{Call: call}
}

// MockControllerAPICloudSpecCall wrap *gomock.Call
type MockControllerAPICloudSpecCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockControllerAPICloudSpecCall) Return(arg0 cloudspec.CloudSpec, arg1 error) *MockControllerAPICloudSpecCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockControllerAPICloudSpecCall) Do(f func(names.ModelTag) (cloudspec.CloudSpec, error)) *MockControllerAPICloudSpecCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockControllerAPICloudSpecCall) DoAndReturn(f func(names.ModelTag) (cloudspec.CloudSpec, error)) *MockControllerAPICloudSpecCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// ControllerConfig mocks base method.
func (m *MockControllerAPI) ControllerConfig(arg0 context.Context) (controller.Config, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ControllerConfig", arg0)
	ret0, _ := ret[0].(controller.Config)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ControllerConfig indicates an expected call of ControllerConfig.
func (mr *MockControllerAPIMockRecorder) ControllerConfig(arg0 any) *MockControllerAPIControllerConfigCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ControllerConfig", reflect.TypeOf((*MockControllerAPI)(nil).ControllerConfig), arg0)
	return &MockControllerAPIControllerConfigCall{Call: call}
}

// MockControllerAPIControllerConfigCall wrap *gomock.Call
type MockControllerAPIControllerConfigCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockControllerAPIControllerConfigCall) Return(arg0 controller.Config, arg1 error) *MockControllerAPIControllerConfigCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockControllerAPIControllerConfigCall) Do(f func(context.Context) (controller.Config, error)) *MockControllerAPIControllerConfigCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockControllerAPIControllerConfigCall) DoAndReturn(f func(context.Context) (controller.Config, error)) *MockControllerAPIControllerConfigCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// ModelConfig mocks base method.
func (m *MockControllerAPI) ModelConfig() (map[string]any, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ModelConfig")
	ret0, _ := ret[0].(map[string]any)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ModelConfig indicates an expected call of ModelConfig.
func (mr *MockControllerAPIMockRecorder) ModelConfig() *MockControllerAPIModelConfigCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ModelConfig", reflect.TypeOf((*MockControllerAPI)(nil).ModelConfig))
	return &MockControllerAPIModelConfigCall{Call: call}
}

// MockControllerAPIModelConfigCall wrap *gomock.Call
type MockControllerAPIModelConfigCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockControllerAPIModelConfigCall) Return(arg0 map[string]any, arg1 error) *MockControllerAPIModelConfigCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockControllerAPIModelConfigCall) Do(f func() (map[string]any, error)) *MockControllerAPIModelConfigCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockControllerAPIModelConfigCall) DoAndReturn(f func() (map[string]any, error)) *MockControllerAPIModelConfigCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}
