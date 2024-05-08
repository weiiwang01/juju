// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/juju/juju/internal/worker/undertaker (interfaces: Facade)
//
// Generated by this command:
//
//	mockgen -typed -package undertaker_test -destination facade_mock_test.go github.com/juju/juju/internal/worker/undertaker Facade
//

// Package undertaker_test is a generated GoMock package.
package undertaker_test

import (
	context "context"
	reflect "reflect"

	status "github.com/juju/juju/core/status"
	watcher "github.com/juju/juju/core/watcher"
	cloudspec "github.com/juju/juju/environs/cloudspec"
	config "github.com/juju/juju/environs/config"
	params "github.com/juju/juju/rpc/params"
	gomock "go.uber.org/mock/gomock"
)

// MockFacade is a mock of Facade interface.
type MockFacade struct {
	ctrl     *gomock.Controller
	recorder *MockFacadeMockRecorder
}

// MockFacadeMockRecorder is the mock recorder for MockFacade.
type MockFacadeMockRecorder struct {
	mock *MockFacade
}

// NewMockFacade creates a new mock instance.
func NewMockFacade(ctrl *gomock.Controller) *MockFacade {
	mock := &MockFacade{ctrl: ctrl}
	mock.recorder = &MockFacadeMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockFacade) EXPECT() *MockFacadeMockRecorder {
	return m.recorder
}

// CloudSpec mocks base method.
func (m *MockFacade) CloudSpec(arg0 context.Context) (cloudspec.CloudSpec, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CloudSpec", arg0)
	ret0, _ := ret[0].(cloudspec.CloudSpec)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CloudSpec indicates an expected call of CloudSpec.
func (mr *MockFacadeMockRecorder) CloudSpec(arg0 any) *MockFacadeCloudSpecCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CloudSpec", reflect.TypeOf((*MockFacade)(nil).CloudSpec), arg0)
	return &MockFacadeCloudSpecCall{Call: call}
}

// MockFacadeCloudSpecCall wrap *gomock.Call
type MockFacadeCloudSpecCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockFacadeCloudSpecCall) Return(arg0 cloudspec.CloudSpec, arg1 error) *MockFacadeCloudSpecCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockFacadeCloudSpecCall) Do(f func(context.Context) (cloudspec.CloudSpec, error)) *MockFacadeCloudSpecCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockFacadeCloudSpecCall) DoAndReturn(f func(context.Context) (cloudspec.CloudSpec, error)) *MockFacadeCloudSpecCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// ModelConfig mocks base method.
func (m *MockFacade) ModelConfig(arg0 context.Context) (*config.Config, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ModelConfig", arg0)
	ret0, _ := ret[0].(*config.Config)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ModelConfig indicates an expected call of ModelConfig.
func (mr *MockFacadeMockRecorder) ModelConfig(arg0 any) *MockFacadeModelConfigCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ModelConfig", reflect.TypeOf((*MockFacade)(nil).ModelConfig), arg0)
	return &MockFacadeModelConfigCall{Call: call}
}

// MockFacadeModelConfigCall wrap *gomock.Call
type MockFacadeModelConfigCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockFacadeModelConfigCall) Return(arg0 *config.Config, arg1 error) *MockFacadeModelConfigCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockFacadeModelConfigCall) Do(f func(context.Context) (*config.Config, error)) *MockFacadeModelConfigCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockFacadeModelConfigCall) DoAndReturn(f func(context.Context) (*config.Config, error)) *MockFacadeModelConfigCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// ModelInfo mocks base method.
func (m *MockFacade) ModelInfo() (params.UndertakerModelInfoResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ModelInfo")
	ret0, _ := ret[0].(params.UndertakerModelInfoResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ModelInfo indicates an expected call of ModelInfo.
func (mr *MockFacadeMockRecorder) ModelInfo() *MockFacadeModelInfoCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ModelInfo", reflect.TypeOf((*MockFacade)(nil).ModelInfo))
	return &MockFacadeModelInfoCall{Call: call}
}

// MockFacadeModelInfoCall wrap *gomock.Call
type MockFacadeModelInfoCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockFacadeModelInfoCall) Return(arg0 params.UndertakerModelInfoResult, arg1 error) *MockFacadeModelInfoCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockFacadeModelInfoCall) Do(f func() (params.UndertakerModelInfoResult, error)) *MockFacadeModelInfoCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockFacadeModelInfoCall) DoAndReturn(f func() (params.UndertakerModelInfoResult, error)) *MockFacadeModelInfoCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// ProcessDyingModel mocks base method.
func (m *MockFacade) ProcessDyingModel() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ProcessDyingModel")
	ret0, _ := ret[0].(error)
	return ret0
}

// ProcessDyingModel indicates an expected call of ProcessDyingModel.
func (mr *MockFacadeMockRecorder) ProcessDyingModel() *MockFacadeProcessDyingModelCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProcessDyingModel", reflect.TypeOf((*MockFacade)(nil).ProcessDyingModel))
	return &MockFacadeProcessDyingModelCall{Call: call}
}

// MockFacadeProcessDyingModelCall wrap *gomock.Call
type MockFacadeProcessDyingModelCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockFacadeProcessDyingModelCall) Return(arg0 error) *MockFacadeProcessDyingModelCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockFacadeProcessDyingModelCall) Do(f func() error) *MockFacadeProcessDyingModelCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockFacadeProcessDyingModelCall) DoAndReturn(f func() error) *MockFacadeProcessDyingModelCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// RemoveModel mocks base method.
func (m *MockFacade) RemoveModel() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveModel")
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveModel indicates an expected call of RemoveModel.
func (mr *MockFacadeMockRecorder) RemoveModel() *MockFacadeRemoveModelCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveModel", reflect.TypeOf((*MockFacade)(nil).RemoveModel))
	return &MockFacadeRemoveModelCall{Call: call}
}

// MockFacadeRemoveModelCall wrap *gomock.Call
type MockFacadeRemoveModelCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockFacadeRemoveModelCall) Return(arg0 error) *MockFacadeRemoveModelCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockFacadeRemoveModelCall) Do(f func() error) *MockFacadeRemoveModelCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockFacadeRemoveModelCall) DoAndReturn(f func() error) *MockFacadeRemoveModelCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// SetStatus mocks base method.
func (m *MockFacade) SetStatus(arg0 status.Status, arg1 string, arg2 map[string]any) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetStatus", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetStatus indicates an expected call of SetStatus.
func (mr *MockFacadeMockRecorder) SetStatus(arg0, arg1, arg2 any) *MockFacadeSetStatusCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetStatus", reflect.TypeOf((*MockFacade)(nil).SetStatus), arg0, arg1, arg2)
	return &MockFacadeSetStatusCall{Call: call}
}

// MockFacadeSetStatusCall wrap *gomock.Call
type MockFacadeSetStatusCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockFacadeSetStatusCall) Return(arg0 error) *MockFacadeSetStatusCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockFacadeSetStatusCall) Do(f func(status.Status, string, map[string]any) error) *MockFacadeSetStatusCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockFacadeSetStatusCall) DoAndReturn(f func(status.Status, string, map[string]any) error) *MockFacadeSetStatusCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// WatchModel mocks base method.
func (m *MockFacade) WatchModel() (watcher.Watcher[struct{}], error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WatchModel")
	ret0, _ := ret[0].(watcher.Watcher[struct{}])
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// WatchModel indicates an expected call of WatchModel.
func (mr *MockFacadeMockRecorder) WatchModel() *MockFacadeWatchModelCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WatchModel", reflect.TypeOf((*MockFacade)(nil).WatchModel))
	return &MockFacadeWatchModelCall{Call: call}
}

// MockFacadeWatchModelCall wrap *gomock.Call
type MockFacadeWatchModelCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockFacadeWatchModelCall) Return(arg0 watcher.Watcher[struct{}], arg1 error) *MockFacadeWatchModelCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockFacadeWatchModelCall) Do(f func() (watcher.Watcher[struct{}], error)) *MockFacadeWatchModelCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockFacadeWatchModelCall) DoAndReturn(f func() (watcher.Watcher[struct{}], error)) *MockFacadeWatchModelCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// WatchModelResources mocks base method.
func (m *MockFacade) WatchModelResources() (watcher.Watcher[struct{}], error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WatchModelResources")
	ret0, _ := ret[0].(watcher.Watcher[struct{}])
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// WatchModelResources indicates an expected call of WatchModelResources.
func (mr *MockFacadeMockRecorder) WatchModelResources() *MockFacadeWatchModelResourcesCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WatchModelResources", reflect.TypeOf((*MockFacade)(nil).WatchModelResources))
	return &MockFacadeWatchModelResourcesCall{Call: call}
}

// MockFacadeWatchModelResourcesCall wrap *gomock.Call
type MockFacadeWatchModelResourcesCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockFacadeWatchModelResourcesCall) Return(arg0 watcher.Watcher[struct{}], arg1 error) *MockFacadeWatchModelResourcesCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockFacadeWatchModelResourcesCall) Do(f func() (watcher.Watcher[struct{}], error)) *MockFacadeWatchModelResourcesCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockFacadeWatchModelResourcesCall) DoAndReturn(f func() (watcher.Watcher[struct{}], error)) *MockFacadeWatchModelResourcesCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}
