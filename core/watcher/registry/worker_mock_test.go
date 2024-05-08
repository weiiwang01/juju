// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/juju/worker/v4 (interfaces: Worker)
//
// Generated by this command:
//
//	mockgen -typed -package registry -destination worker_mock_test.go github.com/juju/worker/v4 Worker
//

// Package registry is a generated GoMock package.
package registry

import (
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockWorker is a mock of Worker interface.
type MockWorker struct {
	ctrl     *gomock.Controller
	recorder *MockWorkerMockRecorder
}

// MockWorkerMockRecorder is the mock recorder for MockWorker.
type MockWorkerMockRecorder struct {
	mock *MockWorker
}

// NewMockWorker creates a new mock instance.
func NewMockWorker(ctrl *gomock.Controller) *MockWorker {
	mock := &MockWorker{ctrl: ctrl}
	mock.recorder = &MockWorkerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockWorker) EXPECT() *MockWorkerMockRecorder {
	return m.recorder
}

// Kill mocks base method.
func (m *MockWorker) Kill() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Kill")
}

// Kill indicates an expected call of Kill.
func (mr *MockWorkerMockRecorder) Kill() *MockWorkerKillCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Kill", reflect.TypeOf((*MockWorker)(nil).Kill))
	return &MockWorkerKillCall{Call: call}
}

// MockWorkerKillCall wrap *gomock.Call
type MockWorkerKillCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockWorkerKillCall) Return() *MockWorkerKillCall {
	c.Call = c.Call.Return()
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockWorkerKillCall) Do(f func()) *MockWorkerKillCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockWorkerKillCall) DoAndReturn(f func()) *MockWorkerKillCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Wait mocks base method.
func (m *MockWorker) Wait() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Wait")
	ret0, _ := ret[0].(error)
	return ret0
}

// Wait indicates an expected call of Wait.
func (mr *MockWorkerMockRecorder) Wait() *MockWorkerWaitCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Wait", reflect.TypeOf((*MockWorker)(nil).Wait))
	return &MockWorkerWaitCall{Call: call}
}

// MockWorkerWaitCall wrap *gomock.Call
type MockWorkerWaitCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockWorkerWaitCall) Return(arg0 error) *MockWorkerWaitCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockWorkerWaitCall) Do(f func() error) *MockWorkerWaitCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockWorkerWaitCall) DoAndReturn(f func() error) *MockWorkerWaitCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}
