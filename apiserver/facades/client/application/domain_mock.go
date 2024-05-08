// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/juju/juju/apiserver/facades/client/application (interfaces: MachineService,ApplicationService,StoragePoolGetter,NetworkService,ExternalControllerService)
//
// Generated by this command:
//
//	mockgen -typed -package application -destination domain_mock.go github.com/juju/juju/apiserver/facades/client/application MachineService,ApplicationService,StoragePoolGetter,NetworkService,ExternalControllerService
//

// Package application is a generated GoMock package.
package application

import (
	context "context"
	reflect "reflect"

	crossmodel "github.com/juju/juju/core/crossmodel"
	network "github.com/juju/juju/core/network"
	service "github.com/juju/juju/domain/application/service"
	storage "github.com/juju/juju/internal/storage"
	gomock "go.uber.org/mock/gomock"
)

// MockMachineService is a mock of MachineService interface.
type MockMachineService struct {
	ctrl     *gomock.Controller
	recorder *MockMachineServiceMockRecorder
}

// MockMachineServiceMockRecorder is the mock recorder for MockMachineService.
type MockMachineServiceMockRecorder struct {
	mock *MockMachineService
}

// NewMockMachineService creates a new mock instance.
func NewMockMachineService(ctrl *gomock.Controller) *MockMachineService {
	mock := &MockMachineService{ctrl: ctrl}
	mock.recorder = &MockMachineServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMachineService) EXPECT() *MockMachineServiceMockRecorder {
	return m.recorder
}

// CreateMachine mocks base method.
func (m *MockMachineService) CreateMachine(arg0 context.Context, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateMachine", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateMachine indicates an expected call of CreateMachine.
func (mr *MockMachineServiceMockRecorder) CreateMachine(arg0, arg1 any) *MockMachineServiceCreateMachineCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateMachine", reflect.TypeOf((*MockMachineService)(nil).CreateMachine), arg0, arg1)
	return &MockMachineServiceCreateMachineCall{Call: call}
}

// MockMachineServiceCreateMachineCall wrap *gomock.Call
type MockMachineServiceCreateMachineCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockMachineServiceCreateMachineCall) Return(arg0 error) *MockMachineServiceCreateMachineCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockMachineServiceCreateMachineCall) Do(f func(context.Context, string) error) *MockMachineServiceCreateMachineCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockMachineServiceCreateMachineCall) DoAndReturn(f func(context.Context, string) error) *MockMachineServiceCreateMachineCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// MockApplicationService is a mock of ApplicationService interface.
type MockApplicationService struct {
	ctrl     *gomock.Controller
	recorder *MockApplicationServiceMockRecorder
}

// MockApplicationServiceMockRecorder is the mock recorder for MockApplicationService.
type MockApplicationServiceMockRecorder struct {
	mock *MockApplicationService
}

// NewMockApplicationService creates a new mock instance.
func NewMockApplicationService(ctrl *gomock.Controller) *MockApplicationService {
	mock := &MockApplicationService{ctrl: ctrl}
	mock.recorder = &MockApplicationServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockApplicationService) EXPECT() *MockApplicationServiceMockRecorder {
	return m.recorder
}

// AddUnits mocks base method.
func (m *MockApplicationService) AddUnits(arg0 context.Context, arg1 string, arg2 ...service.AddUnitParams) error {
	m.ctrl.T.Helper()
	varargs := []any{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "AddUnits", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddUnits indicates an expected call of AddUnits.
func (mr *MockApplicationServiceMockRecorder) AddUnits(arg0, arg1 any, arg2 ...any) *MockApplicationServiceAddUnitsCall {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0, arg1}, arg2...)
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddUnits", reflect.TypeOf((*MockApplicationService)(nil).AddUnits), varargs...)
	return &MockApplicationServiceAddUnitsCall{Call: call}
}

// MockApplicationServiceAddUnitsCall wrap *gomock.Call
type MockApplicationServiceAddUnitsCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockApplicationServiceAddUnitsCall) Return(arg0 error) *MockApplicationServiceAddUnitsCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockApplicationServiceAddUnitsCall) Do(f func(context.Context, string, ...service.AddUnitParams) error) *MockApplicationServiceAddUnitsCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockApplicationServiceAddUnitsCall) DoAndReturn(f func(context.Context, string, ...service.AddUnitParams) error) *MockApplicationServiceAddUnitsCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// CreateApplication mocks base method.
func (m *MockApplicationService) CreateApplication(arg0 context.Context, arg1 string, arg2 service.AddApplicationParams, arg3 ...service.AddUnitParams) error {
	m.ctrl.T.Helper()
	varargs := []any{arg0, arg1, arg2}
	for _, a := range arg3 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateApplication", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateApplication indicates an expected call of CreateApplication.
func (mr *MockApplicationServiceMockRecorder) CreateApplication(arg0, arg1, arg2 any, arg3 ...any) *MockApplicationServiceCreateApplicationCall {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0, arg1, arg2}, arg3...)
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateApplication", reflect.TypeOf((*MockApplicationService)(nil).CreateApplication), varargs...)
	return &MockApplicationServiceCreateApplicationCall{Call: call}
}

// MockApplicationServiceCreateApplicationCall wrap *gomock.Call
type MockApplicationServiceCreateApplicationCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockApplicationServiceCreateApplicationCall) Return(arg0 error) *MockApplicationServiceCreateApplicationCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockApplicationServiceCreateApplicationCall) Do(f func(context.Context, string, service.AddApplicationParams, ...service.AddUnitParams) error) *MockApplicationServiceCreateApplicationCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockApplicationServiceCreateApplicationCall) DoAndReturn(f func(context.Context, string, service.AddApplicationParams, ...service.AddUnitParams) error) *MockApplicationServiceCreateApplicationCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// UpdateApplicationCharm mocks base method.
func (m *MockApplicationService) UpdateApplicationCharm(arg0 context.Context, arg1 string, arg2 service.UpdateCharmParams) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateApplicationCharm", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateApplicationCharm indicates an expected call of UpdateApplicationCharm.
func (mr *MockApplicationServiceMockRecorder) UpdateApplicationCharm(arg0, arg1, arg2 any) *MockApplicationServiceUpdateApplicationCharmCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateApplicationCharm", reflect.TypeOf((*MockApplicationService)(nil).UpdateApplicationCharm), arg0, arg1, arg2)
	return &MockApplicationServiceUpdateApplicationCharmCall{Call: call}
}

// MockApplicationServiceUpdateApplicationCharmCall wrap *gomock.Call
type MockApplicationServiceUpdateApplicationCharmCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockApplicationServiceUpdateApplicationCharmCall) Return(arg0 error) *MockApplicationServiceUpdateApplicationCharmCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockApplicationServiceUpdateApplicationCharmCall) Do(f func(context.Context, string, service.UpdateCharmParams) error) *MockApplicationServiceUpdateApplicationCharmCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockApplicationServiceUpdateApplicationCharmCall) DoAndReturn(f func(context.Context, string, service.UpdateCharmParams) error) *MockApplicationServiceUpdateApplicationCharmCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// MockStoragePoolGetter is a mock of StoragePoolGetter interface.
type MockStoragePoolGetter struct {
	ctrl     *gomock.Controller
	recorder *MockStoragePoolGetterMockRecorder
}

// MockStoragePoolGetterMockRecorder is the mock recorder for MockStoragePoolGetter.
type MockStoragePoolGetterMockRecorder struct {
	mock *MockStoragePoolGetter
}

// NewMockStoragePoolGetter creates a new mock instance.
func NewMockStoragePoolGetter(ctrl *gomock.Controller) *MockStoragePoolGetter {
	mock := &MockStoragePoolGetter{ctrl: ctrl}
	mock.recorder = &MockStoragePoolGetterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStoragePoolGetter) EXPECT() *MockStoragePoolGetterMockRecorder {
	return m.recorder
}

// GetStoragePoolByName mocks base method.
func (m *MockStoragePoolGetter) GetStoragePoolByName(arg0 context.Context, arg1 string) (*storage.Config, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetStoragePoolByName", arg0, arg1)
	ret0, _ := ret[0].(*storage.Config)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetStoragePoolByName indicates an expected call of GetStoragePoolByName.
func (mr *MockStoragePoolGetterMockRecorder) GetStoragePoolByName(arg0, arg1 any) *MockStoragePoolGetterGetStoragePoolByNameCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetStoragePoolByName", reflect.TypeOf((*MockStoragePoolGetter)(nil).GetStoragePoolByName), arg0, arg1)
	return &MockStoragePoolGetterGetStoragePoolByNameCall{Call: call}
}

// MockStoragePoolGetterGetStoragePoolByNameCall wrap *gomock.Call
type MockStoragePoolGetterGetStoragePoolByNameCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockStoragePoolGetterGetStoragePoolByNameCall) Return(arg0 *storage.Config, arg1 error) *MockStoragePoolGetterGetStoragePoolByNameCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockStoragePoolGetterGetStoragePoolByNameCall) Do(f func(context.Context, string) (*storage.Config, error)) *MockStoragePoolGetterGetStoragePoolByNameCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockStoragePoolGetterGetStoragePoolByNameCall) DoAndReturn(f func(context.Context, string) (*storage.Config, error)) *MockStoragePoolGetterGetStoragePoolByNameCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// MockNetworkService is a mock of NetworkService interface.
type MockNetworkService struct {
	ctrl     *gomock.Controller
	recorder *MockNetworkServiceMockRecorder
}

// MockNetworkServiceMockRecorder is the mock recorder for MockNetworkService.
type MockNetworkServiceMockRecorder struct {
	mock *MockNetworkService
}

// NewMockNetworkService creates a new mock instance.
func NewMockNetworkService(ctrl *gomock.Controller) *MockNetworkService {
	mock := &MockNetworkService{ctrl: ctrl}
	mock.recorder = &MockNetworkServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockNetworkService) EXPECT() *MockNetworkServiceMockRecorder {
	return m.recorder
}

// GetAllSpaces mocks base method.
func (m *MockNetworkService) GetAllSpaces(arg0 context.Context) (network.SpaceInfos, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllSpaces", arg0)
	ret0, _ := ret[0].(network.SpaceInfos)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllSpaces indicates an expected call of GetAllSpaces.
func (mr *MockNetworkServiceMockRecorder) GetAllSpaces(arg0 any) *MockNetworkServiceGetAllSpacesCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllSpaces", reflect.TypeOf((*MockNetworkService)(nil).GetAllSpaces), arg0)
	return &MockNetworkServiceGetAllSpacesCall{Call: call}
}

// MockNetworkServiceGetAllSpacesCall wrap *gomock.Call
type MockNetworkServiceGetAllSpacesCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockNetworkServiceGetAllSpacesCall) Return(arg0 network.SpaceInfos, arg1 error) *MockNetworkServiceGetAllSpacesCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockNetworkServiceGetAllSpacesCall) Do(f func(context.Context) (network.SpaceInfos, error)) *MockNetworkServiceGetAllSpacesCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockNetworkServiceGetAllSpacesCall) DoAndReturn(f func(context.Context) (network.SpaceInfos, error)) *MockNetworkServiceGetAllSpacesCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Space mocks base method.
func (m *MockNetworkService) Space(arg0 context.Context, arg1 string) (*network.SpaceInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Space", arg0, arg1)
	ret0, _ := ret[0].(*network.SpaceInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Space indicates an expected call of Space.
func (mr *MockNetworkServiceMockRecorder) Space(arg0, arg1 any) *MockNetworkServiceSpaceCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Space", reflect.TypeOf((*MockNetworkService)(nil).Space), arg0, arg1)
	return &MockNetworkServiceSpaceCall{Call: call}
}

// MockNetworkServiceSpaceCall wrap *gomock.Call
type MockNetworkServiceSpaceCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockNetworkServiceSpaceCall) Return(arg0 *network.SpaceInfo, arg1 error) *MockNetworkServiceSpaceCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockNetworkServiceSpaceCall) Do(f func(context.Context, string) (*network.SpaceInfo, error)) *MockNetworkServiceSpaceCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockNetworkServiceSpaceCall) DoAndReturn(f func(context.Context, string) (*network.SpaceInfo, error)) *MockNetworkServiceSpaceCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// SpaceByName mocks base method.
func (m *MockNetworkService) SpaceByName(arg0 context.Context, arg1 string) (*network.SpaceInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SpaceByName", arg0, arg1)
	ret0, _ := ret[0].(*network.SpaceInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SpaceByName indicates an expected call of SpaceByName.
func (mr *MockNetworkServiceMockRecorder) SpaceByName(arg0, arg1 any) *MockNetworkServiceSpaceByNameCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SpaceByName", reflect.TypeOf((*MockNetworkService)(nil).SpaceByName), arg0, arg1)
	return &MockNetworkServiceSpaceByNameCall{Call: call}
}

// MockNetworkServiceSpaceByNameCall wrap *gomock.Call
type MockNetworkServiceSpaceByNameCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockNetworkServiceSpaceByNameCall) Return(arg0 *network.SpaceInfo, arg1 error) *MockNetworkServiceSpaceByNameCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockNetworkServiceSpaceByNameCall) Do(f func(context.Context, string) (*network.SpaceInfo, error)) *MockNetworkServiceSpaceByNameCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockNetworkServiceSpaceByNameCall) DoAndReturn(f func(context.Context, string) (*network.SpaceInfo, error)) *MockNetworkServiceSpaceByNameCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// MockExternalControllerService is a mock of ExternalControllerService interface.
type MockExternalControllerService struct {
	ctrl     *gomock.Controller
	recorder *MockExternalControllerServiceMockRecorder
}

// MockExternalControllerServiceMockRecorder is the mock recorder for MockExternalControllerService.
type MockExternalControllerServiceMockRecorder struct {
	mock *MockExternalControllerService
}

// NewMockExternalControllerService creates a new mock instance.
func NewMockExternalControllerService(ctrl *gomock.Controller) *MockExternalControllerService {
	mock := &MockExternalControllerService{ctrl: ctrl}
	mock.recorder = &MockExternalControllerServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockExternalControllerService) EXPECT() *MockExternalControllerServiceMockRecorder {
	return m.recorder
}

// UpdateExternalController mocks base method.
func (m *MockExternalControllerService) UpdateExternalController(arg0 context.Context, arg1 crossmodel.ControllerInfo) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateExternalController", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateExternalController indicates an expected call of UpdateExternalController.
func (mr *MockExternalControllerServiceMockRecorder) UpdateExternalController(arg0, arg1 any) *MockExternalControllerServiceUpdateExternalControllerCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateExternalController", reflect.TypeOf((*MockExternalControllerService)(nil).UpdateExternalController), arg0, arg1)
	return &MockExternalControllerServiceUpdateExternalControllerCall{Call: call}
}

// MockExternalControllerServiceUpdateExternalControllerCall wrap *gomock.Call
type MockExternalControllerServiceUpdateExternalControllerCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockExternalControllerServiceUpdateExternalControllerCall) Return(arg0 error) *MockExternalControllerServiceUpdateExternalControllerCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockExternalControllerServiceUpdateExternalControllerCall) Do(f func(context.Context, crossmodel.ControllerInfo) error) *MockExternalControllerServiceUpdateExternalControllerCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockExternalControllerServiceUpdateExternalControllerCall) DoAndReturn(f func(context.Context, crossmodel.ControllerInfo) error) *MockExternalControllerServiceUpdateExternalControllerCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}
