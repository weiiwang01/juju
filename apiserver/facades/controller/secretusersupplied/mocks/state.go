// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/juju/juju/apiserver/facades/controller/secretusersupplied (interfaces: SecretsState)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	secrets "github.com/juju/juju/core/secrets"
	state "github.com/juju/juju/state"
	names "github.com/juju/names/v4"
	gomock "go.uber.org/mock/gomock"
)

// MockSecretsState is a mock of SecretsState interface.
type MockSecretsState struct {
	ctrl     *gomock.Controller
	recorder *MockSecretsStateMockRecorder
}

// MockSecretsStateMockRecorder is the mock recorder for MockSecretsState.
type MockSecretsStateMockRecorder struct {
	mock *MockSecretsState
}

// NewMockSecretsState creates a new mock instance.
func NewMockSecretsState(ctrl *gomock.Controller) *MockSecretsState {
	mock := &MockSecretsState{ctrl: ctrl}
	mock.recorder = &MockSecretsStateMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSecretsState) EXPECT() *MockSecretsStateMockRecorder {
	return m.recorder
}

// DeleteSecret mocks base method.
func (m *MockSecretsState) DeleteSecret(arg0 *secrets.URI, arg1 ...int) ([]secrets.ValueRef, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteSecret", varargs...)
	ret0, _ := ret[0].([]secrets.ValueRef)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteSecret indicates an expected call of DeleteSecret.
func (mr *MockSecretsStateMockRecorder) DeleteSecret(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteSecret", reflect.TypeOf((*MockSecretsState)(nil).DeleteSecret), varargs...)
}

// GetSecret mocks base method.
func (m *MockSecretsState) GetSecret(arg0 *secrets.URI) (*secrets.SecretMetadata, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSecret", arg0)
	ret0, _ := ret[0].(*secrets.SecretMetadata)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSecret indicates an expected call of GetSecret.
func (mr *MockSecretsStateMockRecorder) GetSecret(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSecret", reflect.TypeOf((*MockSecretsState)(nil).GetSecret), arg0)
}

// WatchObsoleteRevisionsNeedPrune mocks base method.
func (m *MockSecretsState) WatchObsoleteRevisionsNeedPrune(arg0 []names.Tag) (state.StringsWatcher, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WatchObsoleteRevisionsNeedPrune", arg0)
	ret0, _ := ret[0].(state.StringsWatcher)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// WatchObsoleteRevisionsNeedPrune indicates an expected call of WatchObsoleteRevisionsNeedPrune.
func (mr *MockSecretsStateMockRecorder) WatchObsoleteRevisionsNeedPrune(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WatchObsoleteRevisionsNeedPrune", reflect.TypeOf((*MockSecretsState)(nil).WatchObsoleteRevisionsNeedPrune), arg0)
}
