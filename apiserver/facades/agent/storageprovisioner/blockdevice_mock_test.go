// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/juju/juju/apiserver/facades/agent/storageprovisioner (interfaces: BlockDeviceService)
//
// Generated by this command:
//
//	mockgen -typed -package storageprovisioner_test -destination blockdevice_mock_test.go github.com/juju/juju/apiserver/facades/agent/storageprovisioner BlockDeviceService
//

// Package storageprovisioner_test is a generated GoMock package.
package storageprovisioner_test

import (
	context "context"
	reflect "reflect"

	blockdevice "github.com/juju/juju/core/blockdevice"
	watcher "github.com/juju/juju/core/watcher"
	gomock "go.uber.org/mock/gomock"
)

// MockBlockDeviceService is a mock of BlockDeviceService interface.
type MockBlockDeviceService struct {
	ctrl     *gomock.Controller
	recorder *MockBlockDeviceServiceMockRecorder
}

// MockBlockDeviceServiceMockRecorder is the mock recorder for MockBlockDeviceService.
type MockBlockDeviceServiceMockRecorder struct {
	mock *MockBlockDeviceService
}

// NewMockBlockDeviceService creates a new mock instance.
func NewMockBlockDeviceService(ctrl *gomock.Controller) *MockBlockDeviceService {
	mock := &MockBlockDeviceService{ctrl: ctrl}
	mock.recorder = &MockBlockDeviceServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBlockDeviceService) EXPECT() *MockBlockDeviceServiceMockRecorder {
	return m.recorder
}

// BlockDevices mocks base method.
func (m *MockBlockDeviceService) BlockDevices(arg0 context.Context, arg1 string) ([]blockdevice.BlockDevice, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BlockDevices", arg0, arg1)
	ret0, _ := ret[0].([]blockdevice.BlockDevice)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BlockDevices indicates an expected call of BlockDevices.
func (mr *MockBlockDeviceServiceMockRecorder) BlockDevices(arg0, arg1 any) *MockBlockDeviceServiceBlockDevicesCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BlockDevices", reflect.TypeOf((*MockBlockDeviceService)(nil).BlockDevices), arg0, arg1)
	return &MockBlockDeviceServiceBlockDevicesCall{Call: call}
}

// MockBlockDeviceServiceBlockDevicesCall wrap *gomock.Call
type MockBlockDeviceServiceBlockDevicesCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockBlockDeviceServiceBlockDevicesCall) Return(arg0 []blockdevice.BlockDevice, arg1 error) *MockBlockDeviceServiceBlockDevicesCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockBlockDeviceServiceBlockDevicesCall) Do(f func(context.Context, string) ([]blockdevice.BlockDevice, error)) *MockBlockDeviceServiceBlockDevicesCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockBlockDeviceServiceBlockDevicesCall) DoAndReturn(f func(context.Context, string) ([]blockdevice.BlockDevice, error)) *MockBlockDeviceServiceBlockDevicesCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// WatchBlockDevices mocks base method.
func (m *MockBlockDeviceService) WatchBlockDevices(arg0 context.Context, arg1 string) (watcher.Watcher[struct{}], error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WatchBlockDevices", arg0, arg1)
	ret0, _ := ret[0].(watcher.Watcher[struct{}])
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// WatchBlockDevices indicates an expected call of WatchBlockDevices.
func (mr *MockBlockDeviceServiceMockRecorder) WatchBlockDevices(arg0, arg1 any) *MockBlockDeviceServiceWatchBlockDevicesCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WatchBlockDevices", reflect.TypeOf((*MockBlockDeviceService)(nil).WatchBlockDevices), arg0, arg1)
	return &MockBlockDeviceServiceWatchBlockDevicesCall{Call: call}
}

// MockBlockDeviceServiceWatchBlockDevicesCall wrap *gomock.Call
type MockBlockDeviceServiceWatchBlockDevicesCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockBlockDeviceServiceWatchBlockDevicesCall) Return(arg0 watcher.Watcher[struct{}], arg1 error) *MockBlockDeviceServiceWatchBlockDevicesCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockBlockDeviceServiceWatchBlockDevicesCall) Do(f func(context.Context, string) (watcher.Watcher[struct{}], error)) *MockBlockDeviceServiceWatchBlockDevicesCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockBlockDeviceServiceWatchBlockDevicesCall) DoAndReturn(f func(context.Context, string) (watcher.Watcher[struct{}], error)) *MockBlockDeviceServiceWatchBlockDevicesCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}
