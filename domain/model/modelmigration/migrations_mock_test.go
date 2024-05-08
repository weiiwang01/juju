// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/juju/juju/domain/model/modelmigration (interfaces: ModelService,ReadOnlyModelService,UserService,ControllerConfigService)
//
// Generated by this command:
//
//	mockgen -typed -package modelmigration -destination migrations_mock_test.go github.com/juju/juju/domain/model/modelmigration ModelService,ReadOnlyModelService,UserService,ControllerConfigService
//

// Package modelmigration is a generated GoMock package.
package modelmigration

import (
	context "context"
	reflect "reflect"

	controller "github.com/juju/juju/controller"
	model "github.com/juju/juju/core/model"
	user "github.com/juju/juju/core/user"
	model0 "github.com/juju/juju/domain/model"
	uuid "github.com/juju/juju/internal/uuid"
	gomock "go.uber.org/mock/gomock"
)

// MockModelService is a mock of ModelService interface.
type MockModelService struct {
	ctrl     *gomock.Controller
	recorder *MockModelServiceMockRecorder
}

// MockModelServiceMockRecorder is the mock recorder for MockModelService.
type MockModelServiceMockRecorder struct {
	mock *MockModelService
}

// NewMockModelService creates a new mock instance.
func NewMockModelService(ctrl *gomock.Controller) *MockModelService {
	mock := &MockModelService{ctrl: ctrl}
	mock.recorder = &MockModelServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockModelService) EXPECT() *MockModelServiceMockRecorder {
	return m.recorder
}

// CreateModel mocks base method.
func (m *MockModelService) CreateModel(arg0 context.Context, arg1 model0.ModelCreationArgs) (func(context.Context) error, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateModel", arg0, arg1)
	ret0, _ := ret[0].(func(context.Context) error)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateModel indicates an expected call of CreateModel.
func (mr *MockModelServiceMockRecorder) CreateModel(arg0, arg1 any) *MockModelServiceCreateModelCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateModel", reflect.TypeOf((*MockModelService)(nil).CreateModel), arg0, arg1)
	return &MockModelServiceCreateModelCall{Call: call}
}

// MockModelServiceCreateModelCall wrap *gomock.Call
type MockModelServiceCreateModelCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockModelServiceCreateModelCall) Return(arg0 func(context.Context) error, arg1 error) *MockModelServiceCreateModelCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockModelServiceCreateModelCall) Do(f func(context.Context, model0.ModelCreationArgs) (func(context.Context) error, error)) *MockModelServiceCreateModelCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockModelServiceCreateModelCall) DoAndReturn(f func(context.Context, model0.ModelCreationArgs) (func(context.Context) error, error)) *MockModelServiceCreateModelCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// DeleteModel mocks base method.
func (m *MockModelService) DeleteModel(arg0 context.Context, arg1 model.UUID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteModel", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteModel indicates an expected call of DeleteModel.
func (mr *MockModelServiceMockRecorder) DeleteModel(arg0, arg1 any) *MockModelServiceDeleteModelCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteModel", reflect.TypeOf((*MockModelService)(nil).DeleteModel), arg0, arg1)
	return &MockModelServiceDeleteModelCall{Call: call}
}

// MockModelServiceDeleteModelCall wrap *gomock.Call
type MockModelServiceDeleteModelCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockModelServiceDeleteModelCall) Return(arg0 error) *MockModelServiceDeleteModelCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockModelServiceDeleteModelCall) Do(f func(context.Context, model.UUID) error) *MockModelServiceDeleteModelCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockModelServiceDeleteModelCall) DoAndReturn(f func(context.Context, model.UUID) error) *MockModelServiceDeleteModelCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// ModelType mocks base method.
func (m *MockModelService) ModelType(arg0 context.Context, arg1 model.UUID) (model.ModelType, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ModelType", arg0, arg1)
	ret0, _ := ret[0].(model.ModelType)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ModelType indicates an expected call of ModelType.
func (mr *MockModelServiceMockRecorder) ModelType(arg0, arg1 any) *MockModelServiceModelTypeCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ModelType", reflect.TypeOf((*MockModelService)(nil).ModelType), arg0, arg1)
	return &MockModelServiceModelTypeCall{Call: call}
}

// MockModelServiceModelTypeCall wrap *gomock.Call
type MockModelServiceModelTypeCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockModelServiceModelTypeCall) Return(arg0 model.ModelType, arg1 error) *MockModelServiceModelTypeCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockModelServiceModelTypeCall) Do(f func(context.Context, model.UUID) (model.ModelType, error)) *MockModelServiceModelTypeCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockModelServiceModelTypeCall) DoAndReturn(f func(context.Context, model.UUID) (model.ModelType, error)) *MockModelServiceModelTypeCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// MockReadOnlyModelService is a mock of ReadOnlyModelService interface.
type MockReadOnlyModelService struct {
	ctrl     *gomock.Controller
	recorder *MockReadOnlyModelServiceMockRecorder
}

// MockReadOnlyModelServiceMockRecorder is the mock recorder for MockReadOnlyModelService.
type MockReadOnlyModelServiceMockRecorder struct {
	mock *MockReadOnlyModelService
}

// NewMockReadOnlyModelService creates a new mock instance.
func NewMockReadOnlyModelService(ctrl *gomock.Controller) *MockReadOnlyModelService {
	mock := &MockReadOnlyModelService{ctrl: ctrl}
	mock.recorder = &MockReadOnlyModelServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockReadOnlyModelService) EXPECT() *MockReadOnlyModelServiceMockRecorder {
	return m.recorder
}

// CreateModel mocks base method.
func (m *MockReadOnlyModelService) CreateModel(arg0 context.Context, arg1 model.UUID, arg2 uuid.UUID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateModel", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateModel indicates an expected call of CreateModel.
func (mr *MockReadOnlyModelServiceMockRecorder) CreateModel(arg0, arg1, arg2 any) *MockReadOnlyModelServiceCreateModelCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateModel", reflect.TypeOf((*MockReadOnlyModelService)(nil).CreateModel), arg0, arg1, arg2)
	return &MockReadOnlyModelServiceCreateModelCall{Call: call}
}

// MockReadOnlyModelServiceCreateModelCall wrap *gomock.Call
type MockReadOnlyModelServiceCreateModelCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockReadOnlyModelServiceCreateModelCall) Return(arg0 error) *MockReadOnlyModelServiceCreateModelCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockReadOnlyModelServiceCreateModelCall) Do(f func(context.Context, model.UUID, uuid.UUID) error) *MockReadOnlyModelServiceCreateModelCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockReadOnlyModelServiceCreateModelCall) DoAndReturn(f func(context.Context, model.UUID, uuid.UUID) error) *MockReadOnlyModelServiceCreateModelCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// DeleteModel mocks base method.
func (m *MockReadOnlyModelService) DeleteModel(arg0 context.Context, arg1 model.UUID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteModel", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteModel indicates an expected call of DeleteModel.
func (mr *MockReadOnlyModelServiceMockRecorder) DeleteModel(arg0, arg1 any) *MockReadOnlyModelServiceDeleteModelCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteModel", reflect.TypeOf((*MockReadOnlyModelService)(nil).DeleteModel), arg0, arg1)
	return &MockReadOnlyModelServiceDeleteModelCall{Call: call}
}

// MockReadOnlyModelServiceDeleteModelCall wrap *gomock.Call
type MockReadOnlyModelServiceDeleteModelCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockReadOnlyModelServiceDeleteModelCall) Return(arg0 error) *MockReadOnlyModelServiceDeleteModelCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockReadOnlyModelServiceDeleteModelCall) Do(f func(context.Context, model.UUID) error) *MockReadOnlyModelServiceDeleteModelCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockReadOnlyModelServiceDeleteModelCall) DoAndReturn(f func(context.Context, model.UUID) error) *MockReadOnlyModelServiceDeleteModelCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

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

// GetUserByName mocks base method.
func (m *MockUserService) GetUserByName(arg0 context.Context, arg1 string) (user.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserByName", arg0, arg1)
	ret0, _ := ret[0].(user.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserByName indicates an expected call of GetUserByName.
func (mr *MockUserServiceMockRecorder) GetUserByName(arg0, arg1 any) *MockUserServiceGetUserByNameCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserByName", reflect.TypeOf((*MockUserService)(nil).GetUserByName), arg0, arg1)
	return &MockUserServiceGetUserByNameCall{Call: call}
}

// MockUserServiceGetUserByNameCall wrap *gomock.Call
type MockUserServiceGetUserByNameCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockUserServiceGetUserByNameCall) Return(arg0 user.User, arg1 error) *MockUserServiceGetUserByNameCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockUserServiceGetUserByNameCall) Do(f func(context.Context, string) (user.User, error)) *MockUserServiceGetUserByNameCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockUserServiceGetUserByNameCall) DoAndReturn(f func(context.Context, string) (user.User, error)) *MockUserServiceGetUserByNameCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// MockControllerConfigService is a mock of ControllerConfigService interface.
type MockControllerConfigService struct {
	ctrl     *gomock.Controller
	recorder *MockControllerConfigServiceMockRecorder
}

// MockControllerConfigServiceMockRecorder is the mock recorder for MockControllerConfigService.
type MockControllerConfigServiceMockRecorder struct {
	mock *MockControllerConfigService
}

// NewMockControllerConfigService creates a new mock instance.
func NewMockControllerConfigService(ctrl *gomock.Controller) *MockControllerConfigService {
	mock := &MockControllerConfigService{ctrl: ctrl}
	mock.recorder = &MockControllerConfigServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockControllerConfigService) EXPECT() *MockControllerConfigServiceMockRecorder {
	return m.recorder
}

// ControllerConfig mocks base method.
func (m *MockControllerConfigService) ControllerConfig(arg0 context.Context) (controller.Config, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ControllerConfig", arg0)
	ret0, _ := ret[0].(controller.Config)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ControllerConfig indicates an expected call of ControllerConfig.
func (mr *MockControllerConfigServiceMockRecorder) ControllerConfig(arg0 any) *MockControllerConfigServiceControllerConfigCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ControllerConfig", reflect.TypeOf((*MockControllerConfigService)(nil).ControllerConfig), arg0)
	return &MockControllerConfigServiceControllerConfigCall{Call: call}
}

// MockControllerConfigServiceControllerConfigCall wrap *gomock.Call
type MockControllerConfigServiceControllerConfigCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockControllerConfigServiceControllerConfigCall) Return(arg0 controller.Config, arg1 error) *MockControllerConfigServiceControllerConfigCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockControllerConfigServiceControllerConfigCall) Do(f func(context.Context) (controller.Config, error)) *MockControllerConfigServiceControllerConfigCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockControllerConfigServiceControllerConfigCall) DoAndReturn(f func(context.Context) (controller.Config, error)) *MockControllerConfigServiceControllerConfigCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}
