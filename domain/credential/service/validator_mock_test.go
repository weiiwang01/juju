// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/juju/juju/domain/credential/service (interfaces: CredentialValidator)
//
// Generated by this command:
//
//	mockgen -typed -package service -destination validator_mock_test.go github.com/juju/juju/domain/credential/service CredentialValidator
//

// Package service is a generated GoMock package.
package service

import (
	context "context"
	reflect "reflect"

	cloud "github.com/juju/juju/cloud"
	credential "github.com/juju/juju/core/credential"
	gomock "go.uber.org/mock/gomock"
)

// MockCredentialValidator is a mock of CredentialValidator interface.
type MockCredentialValidator struct {
	ctrl     *gomock.Controller
	recorder *MockCredentialValidatorMockRecorder
}

// MockCredentialValidatorMockRecorder is the mock recorder for MockCredentialValidator.
type MockCredentialValidatorMockRecorder struct {
	mock *MockCredentialValidator
}

// NewMockCredentialValidator creates a new mock instance.
func NewMockCredentialValidator(ctrl *gomock.Controller) *MockCredentialValidator {
	mock := &MockCredentialValidator{ctrl: ctrl}
	mock.recorder = &MockCredentialValidatorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCredentialValidator) EXPECT() *MockCredentialValidatorMockRecorder {
	return m.recorder
}

// Validate mocks base method.
func (m *MockCredentialValidator) Validate(arg0 context.Context, arg1 CredentialValidationContext, arg2 credential.Key, arg3 *cloud.Credential, arg4 bool) ([]error, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Validate", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].([]error)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Validate indicates an expected call of Validate.
func (mr *MockCredentialValidatorMockRecorder) Validate(arg0, arg1, arg2, arg3, arg4 any) *MockCredentialValidatorValidateCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Validate", reflect.TypeOf((*MockCredentialValidator)(nil).Validate), arg0, arg1, arg2, arg3, arg4)
	return &MockCredentialValidatorValidateCall{Call: call}
}

// MockCredentialValidatorValidateCall wrap *gomock.Call
type MockCredentialValidatorValidateCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockCredentialValidatorValidateCall) Return(arg0 []error, arg1 error) *MockCredentialValidatorValidateCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockCredentialValidatorValidateCall) Do(f func(context.Context, CredentialValidationContext, credential.Key, *cloud.Credential, bool) ([]error, error)) *MockCredentialValidatorValidateCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockCredentialValidatorValidateCall) DoAndReturn(f func(context.Context, CredentialValidationContext, credential.Key, *cloud.Credential, bool) ([]error, error)) *MockCredentialValidatorValidateCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}
