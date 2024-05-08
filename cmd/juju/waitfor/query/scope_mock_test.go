// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/juju/juju/cmd/juju/waitfor/query (interfaces: FuncScope,Scope)
//
// Generated by this command:
//
//	mockgen -typed -package query -destination scope_mock_test.go github.com/juju/juju/cmd/juju/waitfor/query FuncScope,Scope
//

// Package query is a generated GoMock package.
package query

import (
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockFuncScope is a mock of FuncScope interface.
type MockFuncScope struct {
	ctrl     *gomock.Controller
	recorder *MockFuncScopeMockRecorder
}

// MockFuncScopeMockRecorder is the mock recorder for MockFuncScope.
type MockFuncScopeMockRecorder struct {
	mock *MockFuncScope
}

// NewMockFuncScope creates a new mock instance.
func NewMockFuncScope(ctrl *gomock.Controller) *MockFuncScope {
	mock := &MockFuncScope{ctrl: ctrl}
	mock.recorder = &MockFuncScopeMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockFuncScope) EXPECT() *MockFuncScopeMockRecorder {
	return m.recorder
}

// Add mocks base method.
func (m *MockFuncScope) Add(arg0 string, arg1 any) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Add", arg0, arg1)
}

// Add indicates an expected call of Add.
func (mr *MockFuncScopeMockRecorder) Add(arg0, arg1 any) *MockFuncScopeAddCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Add", reflect.TypeOf((*MockFuncScope)(nil).Add), arg0, arg1)
	return &MockFuncScopeAddCall{Call: call}
}

// MockFuncScopeAddCall wrap *gomock.Call
type MockFuncScopeAddCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockFuncScopeAddCall) Return() *MockFuncScopeAddCall {
	c.Call = c.Call.Return()
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockFuncScopeAddCall) Do(f func(string, any)) *MockFuncScopeAddCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockFuncScopeAddCall) DoAndReturn(f func(string, any)) *MockFuncScopeAddCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Call mocks base method.
func (m *MockFuncScope) Call(arg0 *Identifier, arg1 []Box) (any, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Call", arg0, arg1)
	ret0, _ := ret[0].(any)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Call indicates an expected call of Call.
func (mr *MockFuncScopeMockRecorder) Call(arg0, arg1 any) *MockFuncScopeCallCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Call", reflect.TypeOf((*MockFuncScope)(nil).Call), arg0, arg1)
	return &MockFuncScopeCallCall{Call: call}
}

// MockFuncScopeCallCall wrap *gomock.Call
type MockFuncScopeCallCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockFuncScopeCallCall) Return(arg0 any, arg1 error) *MockFuncScopeCallCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockFuncScopeCallCall) Do(f func(*Identifier, []Box) (any, error)) *MockFuncScopeCallCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockFuncScopeCallCall) DoAndReturn(f func(*Identifier, []Box) (any, error)) *MockFuncScopeCallCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// MockScope is a mock of Scope interface.
type MockScope struct {
	ctrl     *gomock.Controller
	recorder *MockScopeMockRecorder
}

// MockScopeMockRecorder is the mock recorder for MockScope.
type MockScopeMockRecorder struct {
	mock *MockScope
}

// NewMockScope creates a new mock instance.
func NewMockScope(ctrl *gomock.Controller) *MockScope {
	mock := &MockScope{ctrl: ctrl}
	mock.recorder = &MockScopeMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockScope) EXPECT() *MockScopeMockRecorder {
	return m.recorder
}

// GetIdentValue mocks base method.
func (m *MockScope) GetIdentValue(arg0 string) (Box, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetIdentValue", arg0)
	ret0, _ := ret[0].(Box)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetIdentValue indicates an expected call of GetIdentValue.
func (mr *MockScopeMockRecorder) GetIdentValue(arg0 any) *MockScopeGetIdentValueCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetIdentValue", reflect.TypeOf((*MockScope)(nil).GetIdentValue), arg0)
	return &MockScopeGetIdentValueCall{Call: call}
}

// MockScopeGetIdentValueCall wrap *gomock.Call
type MockScopeGetIdentValueCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockScopeGetIdentValueCall) Return(arg0 Box, arg1 error) *MockScopeGetIdentValueCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockScopeGetIdentValueCall) Do(f func(string) (Box, error)) *MockScopeGetIdentValueCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockScopeGetIdentValueCall) DoAndReturn(f func(string) (Box, error)) *MockScopeGetIdentValueCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// GetIdents mocks base method.
func (m *MockScope) GetIdents() []string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetIdents")
	ret0, _ := ret[0].([]string)
	return ret0
}

// GetIdents indicates an expected call of GetIdents.
func (mr *MockScopeMockRecorder) GetIdents() *MockScopeGetIdentsCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetIdents", reflect.TypeOf((*MockScope)(nil).GetIdents))
	return &MockScopeGetIdentsCall{Call: call}
}

// MockScopeGetIdentsCall wrap *gomock.Call
type MockScopeGetIdentsCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockScopeGetIdentsCall) Return(arg0 []string) *MockScopeGetIdentsCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockScopeGetIdentsCall) Do(f func() []string) *MockScopeGetIdentsCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockScopeGetIdentsCall) DoAndReturn(f func() []string) *MockScopeGetIdentsCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}
