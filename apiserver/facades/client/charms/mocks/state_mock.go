// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/juju/juju/apiserver/facades/client/charms/interfaces (interfaces: BackendState,BackendModel,Application,Machine,Unit)

// Package mocks is a generated GoMock package.
package mocks

import (
	gomock "github.com/golang/mock/gomock"
	charm "github.com/juju/charm/v8"
	interfaces "github.com/juju/juju/apiserver/facades/client/charms/interfaces"
	controller "github.com/juju/juju/controller"
	charm0 "github.com/juju/juju/core/charm"
	constraints "github.com/juju/juju/core/constraints"
	instance "github.com/juju/juju/core/instance"
	config "github.com/juju/juju/environs/config"
	state "github.com/juju/juju/state"
	names "github.com/juju/names/v4"
	mgo "gopkg.in/mgo.v2"
	reflect "reflect"
)

// MockBackendState is a mock of BackendState interface
type MockBackendState struct {
	ctrl     *gomock.Controller
	recorder *MockBackendStateMockRecorder
}

// MockBackendStateMockRecorder is the mock recorder for MockBackendState
type MockBackendStateMockRecorder struct {
	mock *MockBackendState
}

// NewMockBackendState creates a new mock instance
func NewMockBackendState(ctrl *gomock.Controller) *MockBackendState {
	mock := &MockBackendState{ctrl: ctrl}
	mock.recorder = &MockBackendStateMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockBackendState) EXPECT() *MockBackendStateMockRecorder {
	return m.recorder
}

// AllCharms mocks base method
func (m *MockBackendState) AllCharms() ([]*state.Charm, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AllCharms")
	ret0, _ := ret[0].([]*state.Charm)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AllCharms indicates an expected call of AllCharms
func (mr *MockBackendStateMockRecorder) AllCharms() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AllCharms", reflect.TypeOf((*MockBackendState)(nil).AllCharms))
}

// Application mocks base method
func (m *MockBackendState) Application(arg0 string) (interfaces.Application, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Application", arg0)
	ret0, _ := ret[0].(interfaces.Application)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Application indicates an expected call of Application
func (mr *MockBackendStateMockRecorder) Application(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Application", reflect.TypeOf((*MockBackendState)(nil).Application), arg0)
}

// Charm mocks base method
func (m *MockBackendState) Charm(arg0 *charm.URL) (*state.Charm, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Charm", arg0)
	ret0, _ := ret[0].(*state.Charm)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Charm indicates an expected call of Charm
func (mr *MockBackendStateMockRecorder) Charm(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Charm", reflect.TypeOf((*MockBackendState)(nil).Charm), arg0)
}

// ControllerConfig mocks base method
func (m *MockBackendState) ControllerConfig() (controller.Config, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ControllerConfig")
	ret0, _ := ret[0].(controller.Config)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ControllerConfig indicates an expected call of ControllerConfig
func (mr *MockBackendStateMockRecorder) ControllerConfig() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ControllerConfig", reflect.TypeOf((*MockBackendState)(nil).ControllerConfig))
}

// ControllerTag mocks base method
func (m *MockBackendState) ControllerTag() names.ControllerTag {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ControllerTag")
	ret0, _ := ret[0].(names.ControllerTag)
	return ret0
}

// ControllerTag indicates an expected call of ControllerTag
func (mr *MockBackendStateMockRecorder) ControllerTag() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ControllerTag", reflect.TypeOf((*MockBackendState)(nil).ControllerTag))
}

// Machine mocks base method
func (m *MockBackendState) Machine(arg0 string) (interfaces.Machine, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Machine", arg0)
	ret0, _ := ret[0].(interfaces.Machine)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Machine indicates an expected call of Machine
func (mr *MockBackendStateMockRecorder) Machine(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Machine", reflect.TypeOf((*MockBackendState)(nil).Machine), arg0)
}

// ModelConstraints mocks base method
func (m *MockBackendState) ModelConstraints() (constraints.Value, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ModelConstraints")
	ret0, _ := ret[0].(constraints.Value)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ModelConstraints indicates an expected call of ModelConstraints
func (mr *MockBackendStateMockRecorder) ModelConstraints() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ModelConstraints", reflect.TypeOf((*MockBackendState)(nil).ModelConstraints))
}

// ModelUUID mocks base method
func (m *MockBackendState) ModelUUID() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ModelUUID")
	ret0, _ := ret[0].(string)
	return ret0
}

// ModelUUID indicates an expected call of ModelUUID
func (mr *MockBackendStateMockRecorder) ModelUUID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ModelUUID", reflect.TypeOf((*MockBackendState)(nil).ModelUUID))
}

// MongoSession mocks base method
func (m *MockBackendState) MongoSession() *mgo.Session {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MongoSession")
	ret0, _ := ret[0].(*mgo.Session)
	return ret0
}

// MongoSession indicates an expected call of MongoSession
func (mr *MockBackendStateMockRecorder) MongoSession() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MongoSession", reflect.TypeOf((*MockBackendState)(nil).MongoSession))
}

// PrepareCharmUpload mocks base method
func (m *MockBackendState) PrepareCharmUpload(arg0 *charm.URL) (charm0.StateCharm, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PrepareCharmUpload", arg0)
	ret0, _ := ret[0].(charm0.StateCharm)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PrepareCharmUpload indicates an expected call of PrepareCharmUpload
func (mr *MockBackendStateMockRecorder) PrepareCharmUpload(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PrepareCharmUpload", reflect.TypeOf((*MockBackendState)(nil).PrepareCharmUpload), arg0)
}

// UpdateUploadedCharm mocks base method
func (m *MockBackendState) UpdateUploadedCharm(arg0 state.CharmInfo) (*state.Charm, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUploadedCharm", arg0)
	ret0, _ := ret[0].(*state.Charm)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateUploadedCharm indicates an expected call of UpdateUploadedCharm
func (mr *MockBackendStateMockRecorder) UpdateUploadedCharm(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUploadedCharm", reflect.TypeOf((*MockBackendState)(nil).UpdateUploadedCharm), arg0)
}

// MockBackendModel is a mock of BackendModel interface
type MockBackendModel struct {
	ctrl     *gomock.Controller
	recorder *MockBackendModelMockRecorder
}

// MockBackendModelMockRecorder is the mock recorder for MockBackendModel
type MockBackendModelMockRecorder struct {
	mock *MockBackendModel
}

// NewMockBackendModel creates a new mock instance
func NewMockBackendModel(ctrl *gomock.Controller) *MockBackendModel {
	mock := &MockBackendModel{ctrl: ctrl}
	mock.recorder = &MockBackendModelMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockBackendModel) EXPECT() *MockBackendModelMockRecorder {
	return m.recorder
}

// Config mocks base method
func (m *MockBackendModel) Config() (*config.Config, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Config")
	ret0, _ := ret[0].(*config.Config)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Config indicates an expected call of Config
func (mr *MockBackendModelMockRecorder) Config() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Config", reflect.TypeOf((*MockBackendModel)(nil).Config))
}

// ModelTag mocks base method
func (m *MockBackendModel) ModelTag() names.ModelTag {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ModelTag")
	ret0, _ := ret[0].(names.ModelTag)
	return ret0
}

// ModelTag indicates an expected call of ModelTag
func (mr *MockBackendModelMockRecorder) ModelTag() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ModelTag", reflect.TypeOf((*MockBackendModel)(nil).ModelTag))
}

// MockApplication is a mock of Application interface
type MockApplication struct {
	ctrl     *gomock.Controller
	recorder *MockApplicationMockRecorder
}

// MockApplicationMockRecorder is the mock recorder for MockApplication
type MockApplicationMockRecorder struct {
	mock *MockApplication
}

// NewMockApplication creates a new mock instance
func NewMockApplication(ctrl *gomock.Controller) *MockApplication {
	mock := &MockApplication{ctrl: ctrl}
	mock.recorder = &MockApplicationMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockApplication) EXPECT() *MockApplicationMockRecorder {
	return m.recorder
}

// AllUnits mocks base method
func (m *MockApplication) AllUnits() ([]interfaces.Unit, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AllUnits")
	ret0, _ := ret[0].([]interfaces.Unit)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AllUnits indicates an expected call of AllUnits
func (mr *MockApplicationMockRecorder) AllUnits() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AllUnits", reflect.TypeOf((*MockApplication)(nil).AllUnits))
}

// Constraints mocks base method
func (m *MockApplication) Constraints() (constraints.Value, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Constraints")
	ret0, _ := ret[0].(constraints.Value)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Constraints indicates an expected call of Constraints
func (mr *MockApplicationMockRecorder) Constraints() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Constraints", reflect.TypeOf((*MockApplication)(nil).Constraints))
}

// IsPrincipal mocks base method
func (m *MockApplication) IsPrincipal() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsPrincipal")
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsPrincipal indicates an expected call of IsPrincipal
func (mr *MockApplicationMockRecorder) IsPrincipal() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsPrincipal", reflect.TypeOf((*MockApplication)(nil).IsPrincipal))
}

// MockMachine is a mock of Machine interface
type MockMachine struct {
	ctrl     *gomock.Controller
	recorder *MockMachineMockRecorder
}

// MockMachineMockRecorder is the mock recorder for MockMachine
type MockMachineMockRecorder struct {
	mock *MockMachine
}

// NewMockMachine creates a new mock instance
func NewMockMachine(ctrl *gomock.Controller) *MockMachine {
	mock := &MockMachine{ctrl: ctrl}
	mock.recorder = &MockMachineMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockMachine) EXPECT() *MockMachineMockRecorder {
	return m.recorder
}

// Constraints mocks base method
func (m *MockMachine) Constraints() (constraints.Value, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Constraints")
	ret0, _ := ret[0].(constraints.Value)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Constraints indicates an expected call of Constraints
func (mr *MockMachineMockRecorder) Constraints() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Constraints", reflect.TypeOf((*MockMachine)(nil).Constraints))
}

// HardwareCharacteristics mocks base method
func (m *MockMachine) HardwareCharacteristics() (*instance.HardwareCharacteristics, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HardwareCharacteristics")
	ret0, _ := ret[0].(*instance.HardwareCharacteristics)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// HardwareCharacteristics indicates an expected call of HardwareCharacteristics
func (mr *MockMachineMockRecorder) HardwareCharacteristics() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HardwareCharacteristics", reflect.TypeOf((*MockMachine)(nil).HardwareCharacteristics))
}

// MockUnit is a mock of Unit interface
type MockUnit struct {
	ctrl     *gomock.Controller
	recorder *MockUnitMockRecorder
}

// MockUnitMockRecorder is the mock recorder for MockUnit
type MockUnitMockRecorder struct {
	mock *MockUnit
}

// NewMockUnit creates a new mock instance
func NewMockUnit(ctrl *gomock.Controller) *MockUnit {
	mock := &MockUnit{ctrl: ctrl}
	mock.recorder = &MockUnitMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockUnit) EXPECT() *MockUnitMockRecorder {
	return m.recorder
}

// AssignedMachineId mocks base method
func (m *MockUnit) AssignedMachineId() (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AssignedMachineId")
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AssignedMachineId indicates an expected call of AssignedMachineId
func (mr *MockUnitMockRecorder) AssignedMachineId() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AssignedMachineId", reflect.TypeOf((*MockUnit)(nil).AssignedMachineId))
}
