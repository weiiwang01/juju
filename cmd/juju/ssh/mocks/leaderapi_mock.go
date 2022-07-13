// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/juju/juju/cmd/juju/ssh (interfaces: LeaderAPI)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockLeaderAPI is a mock of LeaderAPI interface.
type MockLeaderAPI struct {
	ctrl     *gomock.Controller
	recorder *MockLeaderAPIMockRecorder
}

// MockLeaderAPIMockRecorder is the mock recorder for MockLeaderAPI.
type MockLeaderAPIMockRecorder struct {
	mock *MockLeaderAPI
}

// NewMockLeaderAPI creates a new mock instance.
func NewMockLeaderAPI(ctrl *gomock.Controller) *MockLeaderAPI {
	mock := &MockLeaderAPI{ctrl: ctrl}
	mock.recorder = &MockLeaderAPIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockLeaderAPI) EXPECT() *MockLeaderAPIMockRecorder {
	return m.recorder
}

// Close mocks base method.
func (m *MockLeaderAPI) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockLeaderAPIMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockLeaderAPI)(nil).Close))
}

// Leader mocks base method.
func (m *MockLeaderAPI) Leader(arg0 string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Leader", arg0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Leader indicates an expected call of Leader.
func (mr *MockLeaderAPIMockRecorder) Leader(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Leader", reflect.TypeOf((*MockLeaderAPI)(nil).Leader), arg0)
}
