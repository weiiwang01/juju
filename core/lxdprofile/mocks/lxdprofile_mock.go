// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/juju/juju/core/lxdprofile (interfaces: LXDProfiler,LXDProfile)
//
// Generated by this command:
//
//	mockgen -typed -package mocks -destination mocks/lxdprofile_mock.go github.com/juju/juju/core/lxdprofile LXDProfiler,LXDProfile
//

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	lxdprofile "github.com/juju/juju/core/lxdprofile"
	gomock "go.uber.org/mock/gomock"
)

// MockLXDProfiler is a mock of LXDProfiler interface.
type MockLXDProfiler struct {
	ctrl     *gomock.Controller
	recorder *MockLXDProfilerMockRecorder
}

// MockLXDProfilerMockRecorder is the mock recorder for MockLXDProfiler.
type MockLXDProfilerMockRecorder struct {
	mock *MockLXDProfiler
}

// NewMockLXDProfiler creates a new mock instance.
func NewMockLXDProfiler(ctrl *gomock.Controller) *MockLXDProfiler {
	mock := &MockLXDProfiler{ctrl: ctrl}
	mock.recorder = &MockLXDProfilerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockLXDProfiler) EXPECT() *MockLXDProfilerMockRecorder {
	return m.recorder
}

// LXDProfile mocks base method.
func (m *MockLXDProfiler) LXDProfile() lxdprofile.LXDProfile {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LXDProfile")
	ret0, _ := ret[0].(lxdprofile.LXDProfile)
	return ret0
}

// LXDProfile indicates an expected call of LXDProfile.
func (mr *MockLXDProfilerMockRecorder) LXDProfile() *MockLXDProfilerLXDProfileCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LXDProfile", reflect.TypeOf((*MockLXDProfiler)(nil).LXDProfile))
	return &MockLXDProfilerLXDProfileCall{Call: call}
}

// MockLXDProfilerLXDProfileCall wrap *gomock.Call
type MockLXDProfilerLXDProfileCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockLXDProfilerLXDProfileCall) Return(arg0 lxdprofile.LXDProfile) *MockLXDProfilerLXDProfileCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockLXDProfilerLXDProfileCall) Do(f func() lxdprofile.LXDProfile) *MockLXDProfilerLXDProfileCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockLXDProfilerLXDProfileCall) DoAndReturn(f func() lxdprofile.LXDProfile) *MockLXDProfilerLXDProfileCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// MockLXDProfile is a mock of LXDProfile interface.
type MockLXDProfile struct {
	ctrl     *gomock.Controller
	recorder *MockLXDProfileMockRecorder
}

// MockLXDProfileMockRecorder is the mock recorder for MockLXDProfile.
type MockLXDProfileMockRecorder struct {
	mock *MockLXDProfile
}

// NewMockLXDProfile creates a new mock instance.
func NewMockLXDProfile(ctrl *gomock.Controller) *MockLXDProfile {
	mock := &MockLXDProfile{ctrl: ctrl}
	mock.recorder = &MockLXDProfileMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockLXDProfile) EXPECT() *MockLXDProfileMockRecorder {
	return m.recorder
}

// Empty mocks base method.
func (m *MockLXDProfile) Empty() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Empty")
	ret0, _ := ret[0].(bool)
	return ret0
}

// Empty indicates an expected call of Empty.
func (mr *MockLXDProfileMockRecorder) Empty() *MockLXDProfileEmptyCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Empty", reflect.TypeOf((*MockLXDProfile)(nil).Empty))
	return &MockLXDProfileEmptyCall{Call: call}
}

// MockLXDProfileEmptyCall wrap *gomock.Call
type MockLXDProfileEmptyCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockLXDProfileEmptyCall) Return(arg0 bool) *MockLXDProfileEmptyCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockLXDProfileEmptyCall) Do(f func() bool) *MockLXDProfileEmptyCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockLXDProfileEmptyCall) DoAndReturn(f func() bool) *MockLXDProfileEmptyCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// ValidateConfigDevices mocks base method.
func (m *MockLXDProfile) ValidateConfigDevices() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ValidateConfigDevices")
	ret0, _ := ret[0].(error)
	return ret0
}

// ValidateConfigDevices indicates an expected call of ValidateConfigDevices.
func (mr *MockLXDProfileMockRecorder) ValidateConfigDevices() *MockLXDProfileValidateConfigDevicesCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ValidateConfigDevices", reflect.TypeOf((*MockLXDProfile)(nil).ValidateConfigDevices))
	return &MockLXDProfileValidateConfigDevicesCall{Call: call}
}

// MockLXDProfileValidateConfigDevicesCall wrap *gomock.Call
type MockLXDProfileValidateConfigDevicesCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockLXDProfileValidateConfigDevicesCall) Return(arg0 error) *MockLXDProfileValidateConfigDevicesCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockLXDProfileValidateConfigDevicesCall) Do(f func() error) *MockLXDProfileValidateConfigDevicesCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockLXDProfileValidateConfigDevicesCall) DoAndReturn(f func() error) *MockLXDProfileValidateConfigDevicesCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}
