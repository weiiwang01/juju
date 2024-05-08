// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/juju/juju/core/leadership (interfaces: Checker,Token)
//
// Generated by this command:
//
//	mockgen -typed -package mocks -destination mocks/leadership_mock.go github.com/juju/juju/core/leadership Checker,Token
//

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	leadership "github.com/juju/juju/core/leadership"
	gomock "go.uber.org/mock/gomock"
)

// MockChecker is a mock of Checker interface.
type MockChecker struct {
	ctrl     *gomock.Controller
	recorder *MockCheckerMockRecorder
}

// MockCheckerMockRecorder is the mock recorder for MockChecker.
type MockCheckerMockRecorder struct {
	mock *MockChecker
}

// NewMockChecker creates a new mock instance.
func NewMockChecker(ctrl *gomock.Controller) *MockChecker {
	mock := &MockChecker{ctrl: ctrl}
	mock.recorder = &MockCheckerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockChecker) EXPECT() *MockCheckerMockRecorder {
	return m.recorder
}

// LeadershipCheck mocks base method.
func (m *MockChecker) LeadershipCheck(arg0, arg1 string) leadership.Token {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LeadershipCheck", arg0, arg1)
	ret0, _ := ret[0].(leadership.Token)
	return ret0
}

// LeadershipCheck indicates an expected call of LeadershipCheck.
func (mr *MockCheckerMockRecorder) LeadershipCheck(arg0, arg1 any) *MockCheckerLeadershipCheckCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LeadershipCheck", reflect.TypeOf((*MockChecker)(nil).LeadershipCheck), arg0, arg1)
	return &MockCheckerLeadershipCheckCall{Call: call}
}

// MockCheckerLeadershipCheckCall wrap *gomock.Call
type MockCheckerLeadershipCheckCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockCheckerLeadershipCheckCall) Return(arg0 leadership.Token) *MockCheckerLeadershipCheckCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockCheckerLeadershipCheckCall) Do(f func(string, string) leadership.Token) *MockCheckerLeadershipCheckCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockCheckerLeadershipCheckCall) DoAndReturn(f func(string, string) leadership.Token) *MockCheckerLeadershipCheckCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// MockToken is a mock of Token interface.
type MockToken struct {
	ctrl     *gomock.Controller
	recorder *MockTokenMockRecorder
}

// MockTokenMockRecorder is the mock recorder for MockToken.
type MockTokenMockRecorder struct {
	mock *MockToken
}

// NewMockToken creates a new mock instance.
func NewMockToken(ctrl *gomock.Controller) *MockToken {
	mock := &MockToken{ctrl: ctrl}
	mock.recorder = &MockTokenMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockToken) EXPECT() *MockTokenMockRecorder {
	return m.recorder
}

// Check mocks base method.
func (m *MockToken) Check() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Check")
	ret0, _ := ret[0].(error)
	return ret0
}

// Check indicates an expected call of Check.
func (mr *MockTokenMockRecorder) Check() *MockTokenCheckCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Check", reflect.TypeOf((*MockToken)(nil).Check))
	return &MockTokenCheckCall{Call: call}
}

// MockTokenCheckCall wrap *gomock.Call
type MockTokenCheckCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockTokenCheckCall) Return(arg0 error) *MockTokenCheckCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockTokenCheckCall) Do(f func() error) *MockTokenCheckCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockTokenCheckCall) DoAndReturn(f func() error) *MockTokenCheckCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}
