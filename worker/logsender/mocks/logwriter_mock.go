// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/juju/juju/api/logsender (interfaces: LogWriter)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	params "github.com/juju/juju/rpc/params"
)

// MockLogWriter is a mock of LogWriter interface.
type MockLogWriter struct {
	ctrl     *gomock.Controller
	recorder *MockLogWriterMockRecorder
}

// MockLogWriterMockRecorder is the mock recorder for MockLogWriter.
type MockLogWriterMockRecorder struct {
	mock *MockLogWriter
}

// NewMockLogWriter creates a new mock instance.
func NewMockLogWriter(ctrl *gomock.Controller) *MockLogWriter {
	mock := &MockLogWriter{ctrl: ctrl}
	mock.recorder = &MockLogWriterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockLogWriter) EXPECT() *MockLogWriterMockRecorder {
	return m.recorder
}

// Close mocks base method.
func (m *MockLogWriter) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockLogWriterMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockLogWriter)(nil).Close))
}

// WriteLog mocks base method.
func (m *MockLogWriter) WriteLog(arg0 *params.LogRecord) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WriteLog", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// WriteLog indicates an expected call of WriteLog.
func (mr *MockLogWriterMockRecorder) WriteLog(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WriteLog", reflect.TypeOf((*MockLogWriter)(nil).WriteLog), arg0)
}
