// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/juju/juju/domain/permission/service (interfaces: State)
//
// Generated by this command:
//
//	mockgen -package service -destination state_mock_test.go github.com/juju/juju/domain/permission/service State
//

// Package service is a generated GoMock package.
package service

import (
	context "context"
	reflect "reflect"

	permission "github.com/juju/juju/core/permission"
	permission0 "github.com/juju/juju/domain/permission"
	uuid "github.com/juju/juju/internal/uuid"
	gomock "go.uber.org/mock/gomock"
)

// MockState is a mock of State interface.
type MockState struct {
	ctrl     *gomock.Controller
	recorder *MockStateMockRecorder
}

// MockStateMockRecorder is the mock recorder for MockState.
type MockStateMockRecorder struct {
	mock *MockState
}

// NewMockState creates a new mock instance.
func NewMockState(ctrl *gomock.Controller) *MockState {
	mock := &MockState{ctrl: ctrl}
	mock.recorder = &MockStateMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockState) EXPECT() *MockStateMockRecorder {
	return m.recorder
}

// CreatePermission mocks base method.
func (m *MockState) CreatePermission(arg0 context.Context, arg1 uuid.UUID, arg2 permission0.UserAccessSpec) (permission.UserAccess, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreatePermission", arg0, arg1, arg2)
	ret0, _ := ret[0].(permission.UserAccess)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreatePermission indicates an expected call of CreatePermission.
func (mr *MockStateMockRecorder) CreatePermission(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreatePermission", reflect.TypeOf((*MockState)(nil).CreatePermission), arg0, arg1, arg2)
}

// DeletePermission mocks base method.
func (m *MockState) DeletePermission(arg0 context.Context, arg1 string, arg2 permission.ID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeletePermission", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeletePermission indicates an expected call of DeletePermission.
func (mr *MockStateMockRecorder) DeletePermission(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeletePermission", reflect.TypeOf((*MockState)(nil).DeletePermission), arg0, arg1, arg2)
}

// ReadAllAccessTypeForUser mocks base method.
func (m *MockState) ReadAllAccessTypeForUser(arg0 context.Context, arg1 string, arg2 permission.ObjectType) ([]permission.UserAccess, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadAllAccessTypeForUser", arg0, arg1, arg2)
	ret0, _ := ret[0].([]permission.UserAccess)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReadAllAccessTypeForUser indicates an expected call of ReadAllAccessTypeForUser.
func (mr *MockStateMockRecorder) ReadAllAccessTypeForUser(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadAllAccessTypeForUser", reflect.TypeOf((*MockState)(nil).ReadAllAccessTypeForUser), arg0, arg1, arg2)
}

// ReadAllUserAccessForTarget mocks base method.
func (m *MockState) ReadAllUserAccessForTarget(arg0 context.Context, arg1 permission.ID) ([]permission.UserAccess, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadAllUserAccessForTarget", arg0, arg1)
	ret0, _ := ret[0].([]permission.UserAccess)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReadAllUserAccessForTarget indicates an expected call of ReadAllUserAccessForTarget.
func (mr *MockStateMockRecorder) ReadAllUserAccessForTarget(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadAllUserAccessForTarget", reflect.TypeOf((*MockState)(nil).ReadAllUserAccessForTarget), arg0, arg1)
}

// ReadAllUserAccessForUser mocks base method.
func (m *MockState) ReadAllUserAccessForUser(arg0 context.Context, arg1 string) ([]permission.UserAccess, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadAllUserAccessForUser", arg0, arg1)
	ret0, _ := ret[0].([]permission.UserAccess)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReadAllUserAccessForUser indicates an expected call of ReadAllUserAccessForUser.
func (mr *MockStateMockRecorder) ReadAllUserAccessForUser(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadAllUserAccessForUser", reflect.TypeOf((*MockState)(nil).ReadAllUserAccessForUser), arg0, arg1)
}

// ReadUserAccessForTarget mocks base method.
func (m *MockState) ReadUserAccessForTarget(arg0 context.Context, arg1 string, arg2 permission.ID) (permission.UserAccess, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadUserAccessForTarget", arg0, arg1, arg2)
	ret0, _ := ret[0].(permission.UserAccess)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReadUserAccessForTarget indicates an expected call of ReadUserAccessForTarget.
func (mr *MockStateMockRecorder) ReadUserAccessForTarget(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadUserAccessForTarget", reflect.TypeOf((*MockState)(nil).ReadUserAccessForTarget), arg0, arg1, arg2)
}

// ReadUserAccessLevelForTarget mocks base method.
func (m *MockState) ReadUserAccessLevelForTarget(arg0 context.Context, arg1 string, arg2 permission.ID) (permission.Access, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadUserAccessLevelForTarget", arg0, arg1, arg2)
	ret0, _ := ret[0].(permission.Access)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReadUserAccessLevelForTarget indicates an expected call of ReadUserAccessLevelForTarget.
func (mr *MockStateMockRecorder) ReadUserAccessLevelForTarget(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadUserAccessLevelForTarget", reflect.TypeOf((*MockState)(nil).ReadUserAccessLevelForTarget), arg0, arg1, arg2)
}

// UpsertPermission mocks base method.
func (m *MockState) UpsertPermission(arg0 context.Context, arg1 permission0.UpsertPermissionArgs) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpsertPermission", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpsertPermission indicates an expected call of UpsertPermission.
func (mr *MockStateMockRecorder) UpsertPermission(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpsertPermission", reflect.TypeOf((*MockState)(nil).UpsertPermission), arg0, arg1)
}
