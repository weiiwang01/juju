// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/juju/juju/cmd/juju/model (interfaces: CommitsCommandAPI)
//
// Generated by this command:
//
//	mockgen -typed -package mocks -destination ./mocks/commits_mock.go github.com/juju/juju/cmd/juju/model CommitsCommandAPI
//

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	model "github.com/juju/juju/core/model"
	gomock "go.uber.org/mock/gomock"
)

// MockCommitsCommandAPI is a mock of CommitsCommandAPI interface.
type MockCommitsCommandAPI struct {
	ctrl     *gomock.Controller
	recorder *MockCommitsCommandAPIMockRecorder
}

// MockCommitsCommandAPIMockRecorder is the mock recorder for MockCommitsCommandAPI.
type MockCommitsCommandAPIMockRecorder struct {
	mock *MockCommitsCommandAPI
}

// NewMockCommitsCommandAPI creates a new mock instance.
func NewMockCommitsCommandAPI(ctrl *gomock.Controller) *MockCommitsCommandAPI {
	mock := &MockCommitsCommandAPI{ctrl: ctrl}
	mock.recorder = &MockCommitsCommandAPIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCommitsCommandAPI) EXPECT() *MockCommitsCommandAPIMockRecorder {
	return m.recorder
}

// Close mocks base method.
func (m *MockCommitsCommandAPI) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockCommitsCommandAPIMockRecorder) Close() *MockCommitsCommandAPICloseCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockCommitsCommandAPI)(nil).Close))
	return &MockCommitsCommandAPICloseCall{Call: call}
}

// MockCommitsCommandAPICloseCall wrap *gomock.Call
type MockCommitsCommandAPICloseCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockCommitsCommandAPICloseCall) Return(arg0 error) *MockCommitsCommandAPICloseCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockCommitsCommandAPICloseCall) Do(f func() error) *MockCommitsCommandAPICloseCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockCommitsCommandAPICloseCall) DoAndReturn(f func() error) *MockCommitsCommandAPICloseCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// ListCommits mocks base method.
func (m *MockCommitsCommandAPI) ListCommits() ([]model.GenerationCommit, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListCommits")
	ret0, _ := ret[0].([]model.GenerationCommit)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListCommits indicates an expected call of ListCommits.
func (mr *MockCommitsCommandAPIMockRecorder) ListCommits() *MockCommitsCommandAPIListCommitsCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListCommits", reflect.TypeOf((*MockCommitsCommandAPI)(nil).ListCommits))
	return &MockCommitsCommandAPIListCommitsCall{Call: call}
}

// MockCommitsCommandAPIListCommitsCall wrap *gomock.Call
type MockCommitsCommandAPIListCommitsCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockCommitsCommandAPIListCommitsCall) Return(arg0 []model.GenerationCommit, arg1 error) *MockCommitsCommandAPIListCommitsCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockCommitsCommandAPIListCommitsCall) Do(f func() ([]model.GenerationCommit, error)) *MockCommitsCommandAPIListCommitsCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockCommitsCommandAPIListCommitsCall) DoAndReturn(f func() ([]model.GenerationCommit, error)) *MockCommitsCommandAPIListCommitsCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}
