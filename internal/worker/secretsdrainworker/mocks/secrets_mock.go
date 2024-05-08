// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/juju/juju/internal/secrets (interfaces: BackendsClient)
//
// Generated by this command:
//
//	mockgen -typed -package mocks -destination mocks/secrets_mock.go github.com/juju/juju/internal/secrets BackendsClient
//

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	secrets "github.com/juju/juju/core/secrets"
	provider "github.com/juju/juju/internal/secrets/provider"
	gomock "go.uber.org/mock/gomock"
)

// MockBackendsClient is a mock of BackendsClient interface.
type MockBackendsClient struct {
	ctrl     *gomock.Controller
	recorder *MockBackendsClientMockRecorder
}

// MockBackendsClientMockRecorder is the mock recorder for MockBackendsClient.
type MockBackendsClientMockRecorder struct {
	mock *MockBackendsClient
}

// NewMockBackendsClient creates a new mock instance.
func NewMockBackendsClient(ctrl *gomock.Controller) *MockBackendsClient {
	mock := &MockBackendsClient{ctrl: ctrl}
	mock.recorder = &MockBackendsClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBackendsClient) EXPECT() *MockBackendsClientMockRecorder {
	return m.recorder
}

// DeleteContent mocks base method.
func (m *MockBackendsClient) DeleteContent(arg0 *secrets.URI, arg1 int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteContent", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteContent indicates an expected call of DeleteContent.
func (mr *MockBackendsClientMockRecorder) DeleteContent(arg0, arg1 any) *MockBackendsClientDeleteContentCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteContent", reflect.TypeOf((*MockBackendsClient)(nil).DeleteContent), arg0, arg1)
	return &MockBackendsClientDeleteContentCall{Call: call}
}

// MockBackendsClientDeleteContentCall wrap *gomock.Call
type MockBackendsClientDeleteContentCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockBackendsClientDeleteContentCall) Return(arg0 error) *MockBackendsClientDeleteContentCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockBackendsClientDeleteContentCall) Do(f func(*secrets.URI, int) error) *MockBackendsClientDeleteContentCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockBackendsClientDeleteContentCall) DoAndReturn(f func(*secrets.URI, int) error) *MockBackendsClientDeleteContentCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// DeleteExternalContent mocks base method.
func (m *MockBackendsClient) DeleteExternalContent(arg0 secrets.ValueRef) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteExternalContent", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteExternalContent indicates an expected call of DeleteExternalContent.
func (mr *MockBackendsClientMockRecorder) DeleteExternalContent(arg0 any) *MockBackendsClientDeleteExternalContentCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteExternalContent", reflect.TypeOf((*MockBackendsClient)(nil).DeleteExternalContent), arg0)
	return &MockBackendsClientDeleteExternalContentCall{Call: call}
}

// MockBackendsClientDeleteExternalContentCall wrap *gomock.Call
type MockBackendsClientDeleteExternalContentCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockBackendsClientDeleteExternalContentCall) Return(arg0 error) *MockBackendsClientDeleteExternalContentCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockBackendsClientDeleteExternalContentCall) Do(f func(secrets.ValueRef) error) *MockBackendsClientDeleteExternalContentCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockBackendsClientDeleteExternalContentCall) DoAndReturn(f func(secrets.ValueRef) error) *MockBackendsClientDeleteExternalContentCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// GetBackend mocks base method.
func (m *MockBackendsClient) GetBackend(arg0 *string, arg1 bool) (provider.SecretsBackend, string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBackend", arg0, arg1)
	ret0, _ := ret[0].(provider.SecretsBackend)
	ret1, _ := ret[1].(string)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetBackend indicates an expected call of GetBackend.
func (mr *MockBackendsClientMockRecorder) GetBackend(arg0, arg1 any) *MockBackendsClientGetBackendCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBackend", reflect.TypeOf((*MockBackendsClient)(nil).GetBackend), arg0, arg1)
	return &MockBackendsClientGetBackendCall{Call: call}
}

// MockBackendsClientGetBackendCall wrap *gomock.Call
type MockBackendsClientGetBackendCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockBackendsClientGetBackendCall) Return(arg0 provider.SecretsBackend, arg1 string, arg2 error) *MockBackendsClientGetBackendCall {
	c.Call = c.Call.Return(arg0, arg1, arg2)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockBackendsClientGetBackendCall) Do(f func(*string, bool) (provider.SecretsBackend, string, error)) *MockBackendsClientGetBackendCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockBackendsClientGetBackendCall) DoAndReturn(f func(*string, bool) (provider.SecretsBackend, string, error)) *MockBackendsClientGetBackendCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// GetContent mocks base method.
func (m *MockBackendsClient) GetContent(arg0 *secrets.URI, arg1 string, arg2, arg3 bool) (secrets.SecretValue, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetContent", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(secrets.SecretValue)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetContent indicates an expected call of GetContent.
func (mr *MockBackendsClientMockRecorder) GetContent(arg0, arg1, arg2, arg3 any) *MockBackendsClientGetContentCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetContent", reflect.TypeOf((*MockBackendsClient)(nil).GetContent), arg0, arg1, arg2, arg3)
	return &MockBackendsClientGetContentCall{Call: call}
}

// MockBackendsClientGetContentCall wrap *gomock.Call
type MockBackendsClientGetContentCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockBackendsClientGetContentCall) Return(arg0 secrets.SecretValue, arg1 error) *MockBackendsClientGetContentCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockBackendsClientGetContentCall) Do(f func(*secrets.URI, string, bool, bool) (secrets.SecretValue, error)) *MockBackendsClientGetContentCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockBackendsClientGetContentCall) DoAndReturn(f func(*secrets.URI, string, bool, bool) (secrets.SecretValue, error)) *MockBackendsClientGetContentCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// GetRevisionContent mocks base method.
func (m *MockBackendsClient) GetRevisionContent(arg0 *secrets.URI, arg1 int) (secrets.SecretValue, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRevisionContent", arg0, arg1)
	ret0, _ := ret[0].(secrets.SecretValue)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRevisionContent indicates an expected call of GetRevisionContent.
func (mr *MockBackendsClientMockRecorder) GetRevisionContent(arg0, arg1 any) *MockBackendsClientGetRevisionContentCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRevisionContent", reflect.TypeOf((*MockBackendsClient)(nil).GetRevisionContent), arg0, arg1)
	return &MockBackendsClientGetRevisionContentCall{Call: call}
}

// MockBackendsClientGetRevisionContentCall wrap *gomock.Call
type MockBackendsClientGetRevisionContentCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockBackendsClientGetRevisionContentCall) Return(arg0 secrets.SecretValue, arg1 error) *MockBackendsClientGetRevisionContentCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockBackendsClientGetRevisionContentCall) Do(f func(*secrets.URI, int) (secrets.SecretValue, error)) *MockBackendsClientGetRevisionContentCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockBackendsClientGetRevisionContentCall) DoAndReturn(f func(*secrets.URI, int) (secrets.SecretValue, error)) *MockBackendsClientGetRevisionContentCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// SaveContent mocks base method.
func (m *MockBackendsClient) SaveContent(arg0 *secrets.URI, arg1 int, arg2 secrets.SecretValue) (secrets.ValueRef, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveContent", arg0, arg1, arg2)
	ret0, _ := ret[0].(secrets.ValueRef)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SaveContent indicates an expected call of SaveContent.
func (mr *MockBackendsClientMockRecorder) SaveContent(arg0, arg1, arg2 any) *MockBackendsClientSaveContentCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveContent", reflect.TypeOf((*MockBackendsClient)(nil).SaveContent), arg0, arg1, arg2)
	return &MockBackendsClientSaveContentCall{Call: call}
}

// MockBackendsClientSaveContentCall wrap *gomock.Call
type MockBackendsClientSaveContentCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockBackendsClientSaveContentCall) Return(arg0 secrets.ValueRef, arg1 error) *MockBackendsClientSaveContentCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockBackendsClientSaveContentCall) Do(f func(*secrets.URI, int, secrets.SecretValue) (secrets.ValueRef, error)) *MockBackendsClientSaveContentCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockBackendsClientSaveContentCall) DoAndReturn(f func(*secrets.URI, int, secrets.SecretValue) (secrets.ValueRef, error)) *MockBackendsClientSaveContentCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}
