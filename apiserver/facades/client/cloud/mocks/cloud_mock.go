// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/juju/juju/apiserver/facades/client/cloud (interfaces: UserService,User,ModelCredentialService,CredentialService,CloudService,CloudPermissionService)
//
// Generated by this command:
//
//	mockgen -package mocks -destination mocks/cloud_mock.go github.com/juju/juju/apiserver/facades/client/cloud UserService,User,ModelCredentialService,CredentialService,CloudService,CloudPermissionService
//

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	cloud "github.com/juju/juju/apiserver/facades/client/cloud"
	cloud0 "github.com/juju/juju/cloud"
	credential "github.com/juju/juju/core/credential"
	permission "github.com/juju/juju/core/permission"
	watcher "github.com/juju/juju/core/watcher"
	service "github.com/juju/juju/domain/credential/service"
	names "github.com/juju/names/v5"
	gomock "go.uber.org/mock/gomock"
)

// MockUserService is a mock of UserService interface.
type MockUserService struct {
	ctrl     *gomock.Controller
	recorder *MockUserServiceMockRecorder
}

// MockUserServiceMockRecorder is the mock recorder for MockUserService.
type MockUserServiceMockRecorder struct {
	mock *MockUserService
}

// NewMockUserService creates a new mock instance.
func NewMockUserService(ctrl *gomock.Controller) *MockUserService {
	mock := &MockUserService{ctrl: ctrl}
	mock.recorder = &MockUserServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserService) EXPECT() *MockUserServiceMockRecorder {
	return m.recorder
}

// User mocks base method.
func (m *MockUserService) User(arg0 names.UserTag) (cloud.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "User", arg0)
	ret0, _ := ret[0].(cloud.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// User indicates an expected call of User.
func (mr *MockUserServiceMockRecorder) User(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "User", reflect.TypeOf((*MockUserService)(nil).User), arg0)
}

// MockUser is a mock of User interface.
type MockUser struct {
	ctrl     *gomock.Controller
	recorder *MockUserMockRecorder
}

// MockUserMockRecorder is the mock recorder for MockUser.
type MockUserMockRecorder struct {
	mock *MockUser
}

// NewMockUser creates a new mock instance.
func NewMockUser(ctrl *gomock.Controller) *MockUser {
	mock := &MockUser{ctrl: ctrl}
	mock.recorder = &MockUserMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUser) EXPECT() *MockUserMockRecorder {
	return m.recorder
}

// DisplayName mocks base method.
func (m *MockUser) DisplayName() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DisplayName")
	ret0, _ := ret[0].(string)
	return ret0
}

// DisplayName indicates an expected call of DisplayName.
func (mr *MockUserMockRecorder) DisplayName() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DisplayName", reflect.TypeOf((*MockUser)(nil).DisplayName))
}

// MockModelCredentialService is a mock of ModelCredentialService interface.
type MockModelCredentialService struct {
	ctrl     *gomock.Controller
	recorder *MockModelCredentialServiceMockRecorder
}

// MockModelCredentialServiceMockRecorder is the mock recorder for MockModelCredentialService.
type MockModelCredentialServiceMockRecorder struct {
	mock *MockModelCredentialService
}

// NewMockModelCredentialService creates a new mock instance.
func NewMockModelCredentialService(ctrl *gomock.Controller) *MockModelCredentialService {
	mock := &MockModelCredentialService{ctrl: ctrl}
	mock.recorder = &MockModelCredentialServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockModelCredentialService) EXPECT() *MockModelCredentialServiceMockRecorder {
	return m.recorder
}

// CredentialModelsAndOwnerAccess mocks base method.
func (m *MockModelCredentialService) CredentialModelsAndOwnerAccess(arg0 names.CloudCredentialTag) ([]cloud0.CredentialOwnerModelAccess, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CredentialModelsAndOwnerAccess", arg0)
	ret0, _ := ret[0].([]cloud0.CredentialOwnerModelAccess)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CredentialModelsAndOwnerAccess indicates an expected call of CredentialModelsAndOwnerAccess.
func (mr *MockModelCredentialServiceMockRecorder) CredentialModelsAndOwnerAccess(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CredentialModelsAndOwnerAccess", reflect.TypeOf((*MockModelCredentialService)(nil).CredentialModelsAndOwnerAccess), arg0)
}

// MockCredentialService is a mock of CredentialService interface.
type MockCredentialService struct {
	ctrl     *gomock.Controller
	recorder *MockCredentialServiceMockRecorder
}

// MockCredentialServiceMockRecorder is the mock recorder for MockCredentialService.
type MockCredentialServiceMockRecorder struct {
	mock *MockCredentialService
}

// NewMockCredentialService creates a new mock instance.
func NewMockCredentialService(ctrl *gomock.Controller) *MockCredentialService {
	mock := &MockCredentialService{ctrl: ctrl}
	mock.recorder = &MockCredentialServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCredentialService) EXPECT() *MockCredentialServiceMockRecorder {
	return m.recorder
}

// AllCloudCredentialsForOwner mocks base method.
func (m *MockCredentialService) AllCloudCredentialsForOwner(arg0 context.Context, arg1 string) (map[credential.ID]cloud0.Credential, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AllCloudCredentialsForOwner", arg0, arg1)
	ret0, _ := ret[0].(map[credential.ID]cloud0.Credential)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AllCloudCredentialsForOwner indicates an expected call of AllCloudCredentialsForOwner.
func (mr *MockCredentialServiceMockRecorder) AllCloudCredentialsForOwner(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AllCloudCredentialsForOwner", reflect.TypeOf((*MockCredentialService)(nil).AllCloudCredentialsForOwner), arg0, arg1)
}

// CheckAndRevokeCredential mocks base method.
func (m *MockCredentialService) CheckAndRevokeCredential(arg0 context.Context, arg1 credential.ID, arg2 bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckAndRevokeCredential", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// CheckAndRevokeCredential indicates an expected call of CheckAndRevokeCredential.
func (mr *MockCredentialServiceMockRecorder) CheckAndRevokeCredential(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckAndRevokeCredential", reflect.TypeOf((*MockCredentialService)(nil).CheckAndRevokeCredential), arg0, arg1, arg2)
}

// CheckAndUpdateCredential mocks base method.
func (m *MockCredentialService) CheckAndUpdateCredential(arg0 context.Context, arg1 credential.ID, arg2 cloud0.Credential, arg3 bool) ([]service.UpdateCredentialModelResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckAndUpdateCredential", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].([]service.UpdateCredentialModelResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CheckAndUpdateCredential indicates an expected call of CheckAndUpdateCredential.
func (mr *MockCredentialServiceMockRecorder) CheckAndUpdateCredential(arg0, arg1, arg2, arg3 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckAndUpdateCredential", reflect.TypeOf((*MockCredentialService)(nil).CheckAndUpdateCredential), arg0, arg1, arg2, arg3)
}

// CloudCredential mocks base method.
func (m *MockCredentialService) CloudCredential(arg0 context.Context, arg1 credential.ID) (cloud0.Credential, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CloudCredential", arg0, arg1)
	ret0, _ := ret[0].(cloud0.Credential)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CloudCredential indicates an expected call of CloudCredential.
func (mr *MockCredentialServiceMockRecorder) CloudCredential(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CloudCredential", reflect.TypeOf((*MockCredentialService)(nil).CloudCredential), arg0, arg1)
}

// CloudCredentialsForOwner mocks base method.
func (m *MockCredentialService) CloudCredentialsForOwner(arg0 context.Context, arg1, arg2 string) (map[string]cloud0.Credential, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CloudCredentialsForOwner", arg0, arg1, arg2)
	ret0, _ := ret[0].(map[string]cloud0.Credential)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CloudCredentialsForOwner indicates an expected call of CloudCredentialsForOwner.
func (mr *MockCredentialServiceMockRecorder) CloudCredentialsForOwner(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CloudCredentialsForOwner", reflect.TypeOf((*MockCredentialService)(nil).CloudCredentialsForOwner), arg0, arg1, arg2)
}

// RemoveCloudCredential mocks base method.
func (m *MockCredentialService) RemoveCloudCredential(arg0 context.Context, arg1 credential.ID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveCloudCredential", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveCloudCredential indicates an expected call of RemoveCloudCredential.
func (mr *MockCredentialServiceMockRecorder) RemoveCloudCredential(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveCloudCredential", reflect.TypeOf((*MockCredentialService)(nil).RemoveCloudCredential), arg0, arg1)
}

// UpdateCloudCredential mocks base method.
func (m *MockCredentialService) UpdateCloudCredential(arg0 context.Context, arg1 credential.ID, arg2 cloud0.Credential) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateCloudCredential", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateCloudCredential indicates an expected call of UpdateCloudCredential.
func (mr *MockCredentialServiceMockRecorder) UpdateCloudCredential(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateCloudCredential", reflect.TypeOf((*MockCredentialService)(nil).UpdateCloudCredential), arg0, arg1, arg2)
}

// WatchCredential mocks base method.
func (m *MockCredentialService) WatchCredential(arg0 context.Context, arg1 credential.ID) (watcher.Watcher[struct{}], error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WatchCredential", arg0, arg1)
	ret0, _ := ret[0].(watcher.Watcher[struct{}])
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// WatchCredential indicates an expected call of WatchCredential.
func (mr *MockCredentialServiceMockRecorder) WatchCredential(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WatchCredential", reflect.TypeOf((*MockCredentialService)(nil).WatchCredential), arg0, arg1)
}

// MockCloudService is a mock of CloudService interface.
type MockCloudService struct {
	ctrl     *gomock.Controller
	recorder *MockCloudServiceMockRecorder
}

// MockCloudServiceMockRecorder is the mock recorder for MockCloudService.
type MockCloudServiceMockRecorder struct {
	mock *MockCloudService
}

// NewMockCloudService creates a new mock instance.
func NewMockCloudService(ctrl *gomock.Controller) *MockCloudService {
	mock := &MockCloudService{ctrl: ctrl}
	mock.recorder = &MockCloudServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCloudService) EXPECT() *MockCloudServiceMockRecorder {
	return m.recorder
}

// Cloud mocks base method.
func (m *MockCloudService) Cloud(arg0 context.Context, arg1 string) (*cloud0.Cloud, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Cloud", arg0, arg1)
	ret0, _ := ret[0].(*cloud0.Cloud)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Cloud indicates an expected call of Cloud.
func (mr *MockCloudServiceMockRecorder) Cloud(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Cloud", reflect.TypeOf((*MockCloudService)(nil).Cloud), arg0, arg1)
}

// DeleteCloud mocks base method.
func (m *MockCloudService) DeleteCloud(arg0 context.Context, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteCloud", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteCloud indicates an expected call of DeleteCloud.
func (mr *MockCloudServiceMockRecorder) DeleteCloud(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteCloud", reflect.TypeOf((*MockCloudService)(nil).DeleteCloud), arg0, arg1)
}

// ListAll mocks base method.
func (m *MockCloudService) ListAll(arg0 context.Context) ([]cloud0.Cloud, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListAll", arg0)
	ret0, _ := ret[0].([]cloud0.Cloud)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListAll indicates an expected call of ListAll.
func (mr *MockCloudServiceMockRecorder) ListAll(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAll", reflect.TypeOf((*MockCloudService)(nil).ListAll), arg0)
}

// UpsertCloud mocks base method.
func (m *MockCloudService) UpsertCloud(arg0 context.Context, arg1 cloud0.Cloud) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpsertCloud", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpsertCloud indicates an expected call of UpsertCloud.
func (mr *MockCloudServiceMockRecorder) UpsertCloud(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpsertCloud", reflect.TypeOf((*MockCloudService)(nil).UpsertCloud), arg0, arg1)
}

// MockCloudPermissionService is a mock of CloudPermissionService interface.
type MockCloudPermissionService struct {
	ctrl     *gomock.Controller
	recorder *MockCloudPermissionServiceMockRecorder
}

// MockCloudPermissionServiceMockRecorder is the mock recorder for MockCloudPermissionService.
type MockCloudPermissionServiceMockRecorder struct {
	mock *MockCloudPermissionService
}

// NewMockCloudPermissionService creates a new mock instance.
func NewMockCloudPermissionService(ctrl *gomock.Controller) *MockCloudPermissionService {
	mock := &MockCloudPermissionService{ctrl: ctrl}
	mock.recorder = &MockCloudPermissionServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCloudPermissionService) EXPECT() *MockCloudPermissionServiceMockRecorder {
	return m.recorder
}

// CloudsForUser mocks base method.
func (m *MockCloudPermissionService) CloudsForUser(arg0 names.UserTag) ([]cloud0.CloudAccess, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CloudsForUser", arg0)
	ret0, _ := ret[0].([]cloud0.CloudAccess)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CloudsForUser indicates an expected call of CloudsForUser.
func (mr *MockCloudPermissionServiceMockRecorder) CloudsForUser(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CloudsForUser", reflect.TypeOf((*MockCloudPermissionService)(nil).CloudsForUser), arg0)
}

// CreateCloudAccess mocks base method.
func (m *MockCloudPermissionService) CreateCloudAccess(arg0 string, arg1 names.UserTag, arg2 permission.Access) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateCloudAccess", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateCloudAccess indicates an expected call of CreateCloudAccess.
func (mr *MockCloudPermissionServiceMockRecorder) CreateCloudAccess(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateCloudAccess", reflect.TypeOf((*MockCloudPermissionService)(nil).CreateCloudAccess), arg0, arg1, arg2)
}

// GetCloudAccess mocks base method.
func (m *MockCloudPermissionService) GetCloudAccess(arg0 string, arg1 names.UserTag) (permission.Access, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCloudAccess", arg0, arg1)
	ret0, _ := ret[0].(permission.Access)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCloudAccess indicates an expected call of GetCloudAccess.
func (mr *MockCloudPermissionServiceMockRecorder) GetCloudAccess(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCloudAccess", reflect.TypeOf((*MockCloudPermissionService)(nil).GetCloudAccess), arg0, arg1)
}

// GetCloudUsers mocks base method.
func (m *MockCloudPermissionService) GetCloudUsers(arg0 string) (map[string]permission.Access, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCloudUsers", arg0)
	ret0, _ := ret[0].(map[string]permission.Access)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCloudUsers indicates an expected call of GetCloudUsers.
func (mr *MockCloudPermissionServiceMockRecorder) GetCloudUsers(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCloudUsers", reflect.TypeOf((*MockCloudPermissionService)(nil).GetCloudUsers), arg0)
}

// RemoveCloudAccess mocks base method.
func (m *MockCloudPermissionService) RemoveCloudAccess(arg0 string, arg1 names.UserTag) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveCloudAccess", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveCloudAccess indicates an expected call of RemoveCloudAccess.
func (mr *MockCloudPermissionServiceMockRecorder) RemoveCloudAccess(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveCloudAccess", reflect.TypeOf((*MockCloudPermissionService)(nil).RemoveCloudAccess), arg0, arg1)
}

// UpdateCloudAccess mocks base method.
func (m *MockCloudPermissionService) UpdateCloudAccess(arg0 string, arg1 names.UserTag, arg2 permission.Access) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateCloudAccess", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateCloudAccess indicates an expected call of UpdateCloudAccess.
func (mr *MockCloudPermissionServiceMockRecorder) UpdateCloudAccess(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateCloudAccess", reflect.TypeOf((*MockCloudPermissionService)(nil).UpdateCloudAccess), arg0, arg1, arg2)
}
