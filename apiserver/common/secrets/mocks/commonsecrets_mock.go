// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/juju/juju/apiserver/common/secrets (interfaces: Model,SecretService,SecretBackendService)
//
// Generated by this command:
//
//	mockgen -package mocks -destination mocks/commonsecrets_mock.go github.com/juju/juju/apiserver/common/secrets Model,SecretService,SecretBackendService
//

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	model "github.com/juju/juju/core/model"
	secrets "github.com/juju/juju/core/secrets"
	service "github.com/juju/juju/domain/secret/service"
	service0 "github.com/juju/juju/domain/secretbackend/service"
	config "github.com/juju/juju/environs/config"
	provider "github.com/juju/juju/internal/secrets/provider"
	state "github.com/juju/juju/state"
	names "github.com/juju/names/v5"
	gomock "go.uber.org/mock/gomock"
)

// MockModel is a mock of Model interface.
type MockModel struct {
	ctrl     *gomock.Controller
	recorder *MockModelMockRecorder
}

// MockModelMockRecorder is the mock recorder for MockModel.
type MockModelMockRecorder struct {
	mock *MockModel
}

// NewMockModel creates a new mock instance.
func NewMockModel(ctrl *gomock.Controller) *MockModel {
	mock := &MockModel{ctrl: ctrl}
	mock.recorder = &MockModelMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockModel) EXPECT() *MockModelMockRecorder {
	return m.recorder
}

// CloudCredentialTag mocks base method.
func (m *MockModel) CloudCredentialTag() (names.CloudCredentialTag, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CloudCredentialTag")
	ret0, _ := ret[0].(names.CloudCredentialTag)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// CloudCredentialTag indicates an expected call of CloudCredentialTag.
func (mr *MockModelMockRecorder) CloudCredentialTag() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CloudCredentialTag", reflect.TypeOf((*MockModel)(nil).CloudCredentialTag))
}

// CloudName mocks base method.
func (m *MockModel) CloudName() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CloudName")
	ret0, _ := ret[0].(string)
	return ret0
}

// CloudName indicates an expected call of CloudName.
func (mr *MockModelMockRecorder) CloudName() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CloudName", reflect.TypeOf((*MockModel)(nil).CloudName))
}

// Config mocks base method.
func (m *MockModel) Config() (*config.Config, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Config")
	ret0, _ := ret[0].(*config.Config)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Config indicates an expected call of Config.
func (mr *MockModelMockRecorder) Config() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Config", reflect.TypeOf((*MockModel)(nil).Config))
}

// ControllerUUID mocks base method.
func (m *MockModel) ControllerUUID() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ControllerUUID")
	ret0, _ := ret[0].(string)
	return ret0
}

// ControllerUUID indicates an expected call of ControllerUUID.
func (mr *MockModelMockRecorder) ControllerUUID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ControllerUUID", reflect.TypeOf((*MockModel)(nil).ControllerUUID))
}

// ModelConfig mocks base method.
func (m *MockModel) ModelConfig(arg0 context.Context) (*config.Config, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ModelConfig", arg0)
	ret0, _ := ret[0].(*config.Config)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ModelConfig indicates an expected call of ModelConfig.
func (mr *MockModelMockRecorder) ModelConfig(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ModelConfig", reflect.TypeOf((*MockModel)(nil).ModelConfig), arg0)
}

// Name mocks base method.
func (m *MockModel) Name() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Name")
	ret0, _ := ret[0].(string)
	return ret0
}

// Name indicates an expected call of Name.
func (mr *MockModelMockRecorder) Name() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Name", reflect.TypeOf((*MockModel)(nil).Name))
}

// State mocks base method.
func (m *MockModel) State() *state.State {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "State")
	ret0, _ := ret[0].(*state.State)
	return ret0
}

// State indicates an expected call of State.
func (mr *MockModelMockRecorder) State() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "State", reflect.TypeOf((*MockModel)(nil).State))
}

// Type mocks base method.
func (m *MockModel) Type() state.ModelType {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Type")
	ret0, _ := ret[0].(state.ModelType)
	return ret0
}

// Type indicates an expected call of Type.
func (mr *MockModelMockRecorder) Type() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Type", reflect.TypeOf((*MockModel)(nil).Type))
}

// UUID mocks base method.
func (m *MockModel) UUID() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UUID")
	ret0, _ := ret[0].(string)
	return ret0
}

// UUID indicates an expected call of UUID.
func (mr *MockModelMockRecorder) UUID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UUID", reflect.TypeOf((*MockModel)(nil).UUID))
}

// WatchForModelConfigChanges mocks base method.
func (m *MockModel) WatchForModelConfigChanges() state.NotifyWatcher {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WatchForModelConfigChanges")
	ret0, _ := ret[0].(state.NotifyWatcher)
	return ret0
}

// WatchForModelConfigChanges indicates an expected call of WatchForModelConfigChanges.
func (mr *MockModelMockRecorder) WatchForModelConfigChanges() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WatchForModelConfigChanges", reflect.TypeOf((*MockModel)(nil).WatchForModelConfigChanges))
}

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

// ChangeSecretBackend mocks base method.
func (m *MockSecretService) ChangeSecretBackend(arg0 context.Context, arg1 *secrets.URI, arg2 int, arg3 service.ChangeSecretBackendParams) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ChangeSecretBackend", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// ChangeSecretBackend indicates an expected call of ChangeSecretBackend.
func (mr *MockSecretServiceMockRecorder) ChangeSecretBackend(arg0, arg1, arg2, arg3 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ChangeSecretBackend", reflect.TypeOf((*MockSecretService)(nil).ChangeSecretBackend), arg0, arg1, arg2, arg3)
}

// ListCharmSecrets mocks base method.
func (m *MockSecretService) ListCharmSecrets(arg0 context.Context, arg1 ...service.CharmSecretOwner) ([]*secrets.SecretMetadata, [][]*secrets.SecretRevisionMetadata, error) {
	m.ctrl.T.Helper()
	varargs := []any{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListCharmSecrets", varargs...)
	ret0, _ := ret[0].([]*secrets.SecretMetadata)
	ret1, _ := ret[1].([][]*secrets.SecretRevisionMetadata)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ListCharmSecrets indicates an expected call of ListCharmSecrets.
func (mr *MockSecretServiceMockRecorder) ListCharmSecrets(arg0 any, arg1 ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListCharmSecrets", reflect.TypeOf((*MockSecretService)(nil).ListCharmSecrets), varargs...)
}

// ListGrantedSecrets mocks base method.
func (m *MockSecretService) ListGrantedSecrets(arg0 context.Context, arg1 ...service.SecretAccessor) ([]*secrets.SecretMetadata, [][]*secrets.SecretRevisionMetadata, error) {
	m.ctrl.T.Helper()
	varargs := []any{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListGrantedSecrets", varargs...)
	ret0, _ := ret[0].([]*secrets.SecretMetadata)
	ret1, _ := ret[1].([][]*secrets.SecretRevisionMetadata)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ListGrantedSecrets indicates an expected call of ListGrantedSecrets.
func (mr *MockSecretServiceMockRecorder) ListGrantedSecrets(arg0 any, arg1 ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListGrantedSecrets", reflect.TypeOf((*MockSecretService)(nil).ListGrantedSecrets), varargs...)
}

// ListUserSecrets mocks base method.
func (m *MockSecretService) ListUserSecrets(arg0 context.Context) ([]*secrets.SecretMetadata, [][]*secrets.SecretRevisionMetadata, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListUserSecrets", arg0)
	ret0, _ := ret[0].([]*secrets.SecretMetadata)
	ret1, _ := ret[1].([][]*secrets.SecretRevisionMetadata)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ListUserSecrets indicates an expected call of ListUserSecrets.
func (mr *MockSecretServiceMockRecorder) ListUserSecrets(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListUserSecrets", reflect.TypeOf((*MockSecretService)(nil).ListUserSecrets), arg0)
}

// MockSecretBackendService is a mock of SecretBackendService interface.
type MockSecretBackendService struct {
	ctrl     *gomock.Controller
	recorder *MockSecretBackendServiceMockRecorder
}

// MockSecretBackendServiceMockRecorder is the mock recorder for MockSecretBackendService.
type MockSecretBackendServiceMockRecorder struct {
	mock *MockSecretBackendService
}

// NewMockSecretBackendService creates a new mock instance.
func NewMockSecretBackendService(ctrl *gomock.Controller) *MockSecretBackendService {
	mock := &MockSecretBackendService{ctrl: ctrl}
	mock.recorder = &MockSecretBackendServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSecretBackendService) EXPECT() *MockSecretBackendServiceMockRecorder {
	return m.recorder
}

// GetRevisionsToDrain mocks base method.
func (m *MockSecretBackendService) GetRevisionsToDrain(arg0 context.Context, arg1 model.UUID, arg2 []*secrets.SecretRevisionMetadata) ([]service0.RevisionInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRevisionsToDrain", arg0, arg1, arg2)
	ret0, _ := ret[0].([]service0.RevisionInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRevisionsToDrain indicates an expected call of GetRevisionsToDrain.
func (mr *MockSecretBackendServiceMockRecorder) GetRevisionsToDrain(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRevisionsToDrain", reflect.TypeOf((*MockSecretBackendService)(nil).GetRevisionsToDrain), arg0, arg1, arg2)
}

// GetSecretBackendConfigForAdmin mocks base method.
func (m *MockSecretBackendService) GetSecretBackendConfigForAdmin(arg0 context.Context, arg1 model.UUID) (*provider.ModelBackendConfigInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSecretBackendConfigForAdmin", arg0, arg1)
	ret0, _ := ret[0].(*provider.ModelBackendConfigInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSecretBackendConfigForAdmin indicates an expected call of GetSecretBackendConfigForAdmin.
func (mr *MockSecretBackendServiceMockRecorder) GetSecretBackendConfigForAdmin(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSecretBackendConfigForAdmin", reflect.TypeOf((*MockSecretBackendService)(nil).GetSecretBackendConfigForAdmin), arg0, arg1)
}
