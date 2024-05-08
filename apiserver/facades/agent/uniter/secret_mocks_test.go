// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/juju/juju/apiserver/facades/agent/uniter (interfaces: SecretService)
//
// Generated by this command:
//
//	mockgen -typed -package uniter -destination secret_mocks_test.go github.com/juju/juju/apiserver/facades/agent/uniter SecretService
//

// Package uniter is a generated GoMock package.
package uniter

import (
	context "context"
	reflect "reflect"

	secrets "github.com/juju/juju/core/secrets"
	service "github.com/juju/juju/domain/secret/service"
	gomock "go.uber.org/mock/gomock"
)

// MockSecretService is a mock of SecretService interface.
type MockSecretService struct {
	ctrl     *gomock.Controller
	recorder *MockSecretServiceMockRecorder
}

// MockSecretServiceMockRecorder is the mock recorder for MockSecretService.
type MockSecretServiceMockRecorder struct {
	mock *MockSecretService
}

// NewMockSecretService creates a new mock instance.
func NewMockSecretService(ctrl *gomock.Controller) *MockSecretService {
	mock := &MockSecretService{ctrl: ctrl}
	mock.recorder = &MockSecretServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSecretService) EXPECT() *MockSecretServiceMockRecorder {
	return m.recorder
}

// CreateCharmSecret mocks base method.
func (m *MockSecretService) CreateCharmSecret(arg0 context.Context, arg1 *secrets.URI, arg2 service.CreateCharmSecretParams) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateCharmSecret", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateCharmSecret indicates an expected call of CreateCharmSecret.
func (mr *MockSecretServiceMockRecorder) CreateCharmSecret(arg0, arg1, arg2 any) *MockSecretServiceCreateCharmSecretCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateCharmSecret", reflect.TypeOf((*MockSecretService)(nil).CreateCharmSecret), arg0, arg1, arg2)
	return &MockSecretServiceCreateCharmSecretCall{Call: call}
}

// MockSecretServiceCreateCharmSecretCall wrap *gomock.Call
type MockSecretServiceCreateCharmSecretCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockSecretServiceCreateCharmSecretCall) Return(arg0 error) *MockSecretServiceCreateCharmSecretCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockSecretServiceCreateCharmSecretCall) Do(f func(context.Context, *secrets.URI, service.CreateCharmSecretParams) error) *MockSecretServiceCreateCharmSecretCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockSecretServiceCreateCharmSecretCall) DoAndReturn(f func(context.Context, *secrets.URI, service.CreateCharmSecretParams) error) *MockSecretServiceCreateCharmSecretCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// DeleteSecret mocks base method.
func (m *MockSecretService) DeleteSecret(arg0 context.Context, arg1 *secrets.URI, arg2 service.DeleteSecretParams) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteSecret", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteSecret indicates an expected call of DeleteSecret.
func (mr *MockSecretServiceMockRecorder) DeleteSecret(arg0, arg1, arg2 any) *MockSecretServiceDeleteSecretCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteSecret", reflect.TypeOf((*MockSecretService)(nil).DeleteSecret), arg0, arg1, arg2)
	return &MockSecretServiceDeleteSecretCall{Call: call}
}

// MockSecretServiceDeleteSecretCall wrap *gomock.Call
type MockSecretServiceDeleteSecretCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockSecretServiceDeleteSecretCall) Return(arg0 error) *MockSecretServiceDeleteSecretCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockSecretServiceDeleteSecretCall) Do(f func(context.Context, *secrets.URI, service.DeleteSecretParams) error) *MockSecretServiceDeleteSecretCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockSecretServiceDeleteSecretCall) DoAndReturn(f func(context.Context, *secrets.URI, service.DeleteSecretParams) error) *MockSecretServiceDeleteSecretCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// GetConsumedRevision mocks base method.
func (m *MockSecretService) GetConsumedRevision(arg0 context.Context, arg1 *secrets.URI, arg2 string, arg3, arg4 bool, arg5 *string) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetConsumedRevision", arg0, arg1, arg2, arg3, arg4, arg5)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetConsumedRevision indicates an expected call of GetConsumedRevision.
func (mr *MockSecretServiceMockRecorder) GetConsumedRevision(arg0, arg1, arg2, arg3, arg4, arg5 any) *MockSecretServiceGetConsumedRevisionCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetConsumedRevision", reflect.TypeOf((*MockSecretService)(nil).GetConsumedRevision), arg0, arg1, arg2, arg3, arg4, arg5)
	return &MockSecretServiceGetConsumedRevisionCall{Call: call}
}

// MockSecretServiceGetConsumedRevisionCall wrap *gomock.Call
type MockSecretServiceGetConsumedRevisionCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockSecretServiceGetConsumedRevisionCall) Return(arg0 int, arg1 error) *MockSecretServiceGetConsumedRevisionCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockSecretServiceGetConsumedRevisionCall) Do(f func(context.Context, *secrets.URI, string, bool, bool, *string) (int, error)) *MockSecretServiceGetConsumedRevisionCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockSecretServiceGetConsumedRevisionCall) DoAndReturn(f func(context.Context, *secrets.URI, string, bool, bool, *string) (int, error)) *MockSecretServiceGetConsumedRevisionCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// GetSecretValue mocks base method.
func (m *MockSecretService) GetSecretValue(arg0 context.Context, arg1 *secrets.URI, arg2 int, arg3 service.SecretAccessor) (secrets.SecretValue, *secrets.ValueRef, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSecretValue", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(secrets.SecretValue)
	ret1, _ := ret[1].(*secrets.ValueRef)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetSecretValue indicates an expected call of GetSecretValue.
func (mr *MockSecretServiceMockRecorder) GetSecretValue(arg0, arg1, arg2, arg3 any) *MockSecretServiceGetSecretValueCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSecretValue", reflect.TypeOf((*MockSecretService)(nil).GetSecretValue), arg0, arg1, arg2, arg3)
	return &MockSecretServiceGetSecretValueCall{Call: call}
}

// MockSecretServiceGetSecretValueCall wrap *gomock.Call
type MockSecretServiceGetSecretValueCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockSecretServiceGetSecretValueCall) Return(arg0 secrets.SecretValue, arg1 *secrets.ValueRef, arg2 error) *MockSecretServiceGetSecretValueCall {
	c.Call = c.Call.Return(arg0, arg1, arg2)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockSecretServiceGetSecretValueCall) Do(f func(context.Context, *secrets.URI, int, service.SecretAccessor) (secrets.SecretValue, *secrets.ValueRef, error)) *MockSecretServiceGetSecretValueCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockSecretServiceGetSecretValueCall) DoAndReturn(f func(context.Context, *secrets.URI, int, service.SecretAccessor) (secrets.SecretValue, *secrets.ValueRef, error)) *MockSecretServiceGetSecretValueCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// GrantSecretAccess mocks base method.
func (m *MockSecretService) GrantSecretAccess(arg0 context.Context, arg1 *secrets.URI, arg2 service.SecretAccessParams) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GrantSecretAccess", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// GrantSecretAccess indicates an expected call of GrantSecretAccess.
func (mr *MockSecretServiceMockRecorder) GrantSecretAccess(arg0, arg1, arg2 any) *MockSecretServiceGrantSecretAccessCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GrantSecretAccess", reflect.TypeOf((*MockSecretService)(nil).GrantSecretAccess), arg0, arg1, arg2)
	return &MockSecretServiceGrantSecretAccessCall{Call: call}
}

// MockSecretServiceGrantSecretAccessCall wrap *gomock.Call
type MockSecretServiceGrantSecretAccessCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockSecretServiceGrantSecretAccessCall) Return(arg0 error) *MockSecretServiceGrantSecretAccessCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockSecretServiceGrantSecretAccessCall) Do(f func(context.Context, *secrets.URI, service.SecretAccessParams) error) *MockSecretServiceGrantSecretAccessCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockSecretServiceGrantSecretAccessCall) DoAndReturn(f func(context.Context, *secrets.URI, service.SecretAccessParams) error) *MockSecretServiceGrantSecretAccessCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// RevokeSecretAccess mocks base method.
func (m *MockSecretService) RevokeSecretAccess(arg0 context.Context, arg1 *secrets.URI, arg2 service.SecretAccessParams) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RevokeSecretAccess", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// RevokeSecretAccess indicates an expected call of RevokeSecretAccess.
func (mr *MockSecretServiceMockRecorder) RevokeSecretAccess(arg0, arg1, arg2 any) *MockSecretServiceRevokeSecretAccessCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RevokeSecretAccess", reflect.TypeOf((*MockSecretService)(nil).RevokeSecretAccess), arg0, arg1, arg2)
	return &MockSecretServiceRevokeSecretAccessCall{Call: call}
}

// MockSecretServiceRevokeSecretAccessCall wrap *gomock.Call
type MockSecretServiceRevokeSecretAccessCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockSecretServiceRevokeSecretAccessCall) Return(arg0 error) *MockSecretServiceRevokeSecretAccessCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockSecretServiceRevokeSecretAccessCall) Do(f func(context.Context, *secrets.URI, service.SecretAccessParams) error) *MockSecretServiceRevokeSecretAccessCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockSecretServiceRevokeSecretAccessCall) DoAndReturn(f func(context.Context, *secrets.URI, service.SecretAccessParams) error) *MockSecretServiceRevokeSecretAccessCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// UpdateCharmSecret mocks base method.
func (m *MockSecretService) UpdateCharmSecret(arg0 context.Context, arg1 *secrets.URI, arg2 service.UpdateCharmSecretParams) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateCharmSecret", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateCharmSecret indicates an expected call of UpdateCharmSecret.
func (mr *MockSecretServiceMockRecorder) UpdateCharmSecret(arg0, arg1, arg2 any) *MockSecretServiceUpdateCharmSecretCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateCharmSecret", reflect.TypeOf((*MockSecretService)(nil).UpdateCharmSecret), arg0, arg1, arg2)
	return &MockSecretServiceUpdateCharmSecretCall{Call: call}
}

// MockSecretServiceUpdateCharmSecretCall wrap *gomock.Call
type MockSecretServiceUpdateCharmSecretCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockSecretServiceUpdateCharmSecretCall) Return(arg0 error) *MockSecretServiceUpdateCharmSecretCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockSecretServiceUpdateCharmSecretCall) Do(f func(context.Context, *secrets.URI, service.UpdateCharmSecretParams) error) *MockSecretServiceUpdateCharmSecretCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockSecretServiceUpdateCharmSecretCall) DoAndReturn(f func(context.Context, *secrets.URI, service.UpdateCharmSecretParams) error) *MockSecretServiceUpdateCharmSecretCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}
