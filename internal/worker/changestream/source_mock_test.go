// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/juju/juju/core/changestream (interfaces: EventSource)
//
// Generated by this command:
//
//	mockgen -typed -package changestream -destination source_mock_test.go github.com/juju/juju/core/changestream EventSource
//

// Package changestream is a generated GoMock package.
package changestream

import (
	reflect "reflect"

	changestream "github.com/juju/juju/core/changestream"
	gomock "go.uber.org/mock/gomock"
)

// MockEventSource is a mock of EventSource interface.
type MockEventSource struct {
	ctrl     *gomock.Controller
	recorder *MockEventSourceMockRecorder
}

// MockEventSourceMockRecorder is the mock recorder for MockEventSource.
type MockEventSourceMockRecorder struct {
	mock *MockEventSource
}

// NewMockEventSource creates a new mock instance.
func NewMockEventSource(ctrl *gomock.Controller) *MockEventSource {
	mock := &MockEventSource{ctrl: ctrl}
	mock.recorder = &MockEventSourceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockEventSource) EXPECT() *MockEventSourceMockRecorder {
	return m.recorder
}

// Subscribe mocks base method.
func (m *MockEventSource) Subscribe(arg0 ...changestream.SubscriptionOption) (changestream.Subscription, error) {
	m.ctrl.T.Helper()
	varargs := []any{}
	for _, a := range arg0 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Subscribe", varargs...)
	ret0, _ := ret[0].(changestream.Subscription)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Subscribe indicates an expected call of Subscribe.
func (mr *MockEventSourceMockRecorder) Subscribe(arg0 ...any) *MockEventSourceSubscribeCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Subscribe", reflect.TypeOf((*MockEventSource)(nil).Subscribe), arg0...)
	return &MockEventSourceSubscribeCall{Call: call}
}

// MockEventSourceSubscribeCall wrap *gomock.Call
type MockEventSourceSubscribeCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockEventSourceSubscribeCall) Return(arg0 changestream.Subscription, arg1 error) *MockEventSourceSubscribeCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockEventSourceSubscribeCall) Do(f func(...changestream.SubscriptionOption) (changestream.Subscription, error)) *MockEventSourceSubscribeCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockEventSourceSubscribeCall) DoAndReturn(f func(...changestream.SubscriptionOption) (changestream.Subscription, error)) *MockEventSourceSubscribeCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}
