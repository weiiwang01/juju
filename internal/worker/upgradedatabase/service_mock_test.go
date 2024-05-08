// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/juju/juju/internal/worker/upgradedatabase (interfaces: UpgradeService,ModelService)
//
// Generated by this command:
//
//	mockgen -typed -package upgradedatabase -destination service_mock_test.go github.com/juju/juju/internal/worker/upgradedatabase UpgradeService,ModelService
//

// Package upgradedatabase is a generated GoMock package.
package upgradedatabase

import (
	context "context"
	reflect "reflect"

	model "github.com/juju/juju/core/model"
	upgrade "github.com/juju/juju/core/upgrade"
	watcher "github.com/juju/juju/core/watcher"
	upgrade0 "github.com/juju/juju/domain/upgrade"
	version "github.com/juju/version/v2"
	gomock "go.uber.org/mock/gomock"
)

// MockUpgradeService is a mock of UpgradeService interface.
type MockUpgradeService struct {
	ctrl     *gomock.Controller
	recorder *MockUpgradeServiceMockRecorder
}

// MockUpgradeServiceMockRecorder is the mock recorder for MockUpgradeService.
type MockUpgradeServiceMockRecorder struct {
	mock *MockUpgradeService
}

// NewMockUpgradeService creates a new mock instance.
func NewMockUpgradeService(ctrl *gomock.Controller) *MockUpgradeService {
	mock := &MockUpgradeService{ctrl: ctrl}
	mock.recorder = &MockUpgradeServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUpgradeService) EXPECT() *MockUpgradeServiceMockRecorder {
	return m.recorder
}

// ActiveUpgrade mocks base method.
func (m *MockUpgradeService) ActiveUpgrade(arg0 context.Context) (upgrade0.UUID, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ActiveUpgrade", arg0)
	ret0, _ := ret[0].(upgrade0.UUID)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ActiveUpgrade indicates an expected call of ActiveUpgrade.
func (mr *MockUpgradeServiceMockRecorder) ActiveUpgrade(arg0 any) *MockUpgradeServiceActiveUpgradeCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ActiveUpgrade", reflect.TypeOf((*MockUpgradeService)(nil).ActiveUpgrade), arg0)
	return &MockUpgradeServiceActiveUpgradeCall{Call: call}
}

// MockUpgradeServiceActiveUpgradeCall wrap *gomock.Call
type MockUpgradeServiceActiveUpgradeCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockUpgradeServiceActiveUpgradeCall) Return(arg0 upgrade0.UUID, arg1 error) *MockUpgradeServiceActiveUpgradeCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockUpgradeServiceActiveUpgradeCall) Do(f func(context.Context) (upgrade0.UUID, error)) *MockUpgradeServiceActiveUpgradeCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockUpgradeServiceActiveUpgradeCall) DoAndReturn(f func(context.Context) (upgrade0.UUID, error)) *MockUpgradeServiceActiveUpgradeCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// CreateUpgrade mocks base method.
func (m *MockUpgradeService) CreateUpgrade(arg0 context.Context, arg1, arg2 version.Number) (upgrade0.UUID, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUpgrade", arg0, arg1, arg2)
	ret0, _ := ret[0].(upgrade0.UUID)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateUpgrade indicates an expected call of CreateUpgrade.
func (mr *MockUpgradeServiceMockRecorder) CreateUpgrade(arg0, arg1, arg2 any) *MockUpgradeServiceCreateUpgradeCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUpgrade", reflect.TypeOf((*MockUpgradeService)(nil).CreateUpgrade), arg0, arg1, arg2)
	return &MockUpgradeServiceCreateUpgradeCall{Call: call}
}

// MockUpgradeServiceCreateUpgradeCall wrap *gomock.Call
type MockUpgradeServiceCreateUpgradeCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockUpgradeServiceCreateUpgradeCall) Return(arg0 upgrade0.UUID, arg1 error) *MockUpgradeServiceCreateUpgradeCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockUpgradeServiceCreateUpgradeCall) Do(f func(context.Context, version.Number, version.Number) (upgrade0.UUID, error)) *MockUpgradeServiceCreateUpgradeCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockUpgradeServiceCreateUpgradeCall) DoAndReturn(f func(context.Context, version.Number, version.Number) (upgrade0.UUID, error)) *MockUpgradeServiceCreateUpgradeCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// SetControllerReady mocks base method.
func (m *MockUpgradeService) SetControllerReady(arg0 context.Context, arg1 upgrade0.UUID, arg2 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetControllerReady", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetControllerReady indicates an expected call of SetControllerReady.
func (mr *MockUpgradeServiceMockRecorder) SetControllerReady(arg0, arg1, arg2 any) *MockUpgradeServiceSetControllerReadyCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetControllerReady", reflect.TypeOf((*MockUpgradeService)(nil).SetControllerReady), arg0, arg1, arg2)
	return &MockUpgradeServiceSetControllerReadyCall{Call: call}
}

// MockUpgradeServiceSetControllerReadyCall wrap *gomock.Call
type MockUpgradeServiceSetControllerReadyCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockUpgradeServiceSetControllerReadyCall) Return(arg0 error) *MockUpgradeServiceSetControllerReadyCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockUpgradeServiceSetControllerReadyCall) Do(f func(context.Context, upgrade0.UUID, string) error) *MockUpgradeServiceSetControllerReadyCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockUpgradeServiceSetControllerReadyCall) DoAndReturn(f func(context.Context, upgrade0.UUID, string) error) *MockUpgradeServiceSetControllerReadyCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// SetDBUpgradeCompleted mocks base method.
func (m *MockUpgradeService) SetDBUpgradeCompleted(arg0 context.Context, arg1 upgrade0.UUID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetDBUpgradeCompleted", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetDBUpgradeCompleted indicates an expected call of SetDBUpgradeCompleted.
func (mr *MockUpgradeServiceMockRecorder) SetDBUpgradeCompleted(arg0, arg1 any) *MockUpgradeServiceSetDBUpgradeCompletedCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetDBUpgradeCompleted", reflect.TypeOf((*MockUpgradeService)(nil).SetDBUpgradeCompleted), arg0, arg1)
	return &MockUpgradeServiceSetDBUpgradeCompletedCall{Call: call}
}

// MockUpgradeServiceSetDBUpgradeCompletedCall wrap *gomock.Call
type MockUpgradeServiceSetDBUpgradeCompletedCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockUpgradeServiceSetDBUpgradeCompletedCall) Return(arg0 error) *MockUpgradeServiceSetDBUpgradeCompletedCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockUpgradeServiceSetDBUpgradeCompletedCall) Do(f func(context.Context, upgrade0.UUID) error) *MockUpgradeServiceSetDBUpgradeCompletedCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockUpgradeServiceSetDBUpgradeCompletedCall) DoAndReturn(f func(context.Context, upgrade0.UUID) error) *MockUpgradeServiceSetDBUpgradeCompletedCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// SetDBUpgradeFailed mocks base method.
func (m *MockUpgradeService) SetDBUpgradeFailed(arg0 context.Context, arg1 upgrade0.UUID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetDBUpgradeFailed", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetDBUpgradeFailed indicates an expected call of SetDBUpgradeFailed.
func (mr *MockUpgradeServiceMockRecorder) SetDBUpgradeFailed(arg0, arg1 any) *MockUpgradeServiceSetDBUpgradeFailedCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetDBUpgradeFailed", reflect.TypeOf((*MockUpgradeService)(nil).SetDBUpgradeFailed), arg0, arg1)
	return &MockUpgradeServiceSetDBUpgradeFailedCall{Call: call}
}

// MockUpgradeServiceSetDBUpgradeFailedCall wrap *gomock.Call
type MockUpgradeServiceSetDBUpgradeFailedCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockUpgradeServiceSetDBUpgradeFailedCall) Return(arg0 error) *MockUpgradeServiceSetDBUpgradeFailedCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockUpgradeServiceSetDBUpgradeFailedCall) Do(f func(context.Context, upgrade0.UUID) error) *MockUpgradeServiceSetDBUpgradeFailedCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockUpgradeServiceSetDBUpgradeFailedCall) DoAndReturn(f func(context.Context, upgrade0.UUID) error) *MockUpgradeServiceSetDBUpgradeFailedCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// StartUpgrade mocks base method.
func (m *MockUpgradeService) StartUpgrade(arg0 context.Context, arg1 upgrade0.UUID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StartUpgrade", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// StartUpgrade indicates an expected call of StartUpgrade.
func (mr *MockUpgradeServiceMockRecorder) StartUpgrade(arg0, arg1 any) *MockUpgradeServiceStartUpgradeCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StartUpgrade", reflect.TypeOf((*MockUpgradeService)(nil).StartUpgrade), arg0, arg1)
	return &MockUpgradeServiceStartUpgradeCall{Call: call}
}

// MockUpgradeServiceStartUpgradeCall wrap *gomock.Call
type MockUpgradeServiceStartUpgradeCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockUpgradeServiceStartUpgradeCall) Return(arg0 error) *MockUpgradeServiceStartUpgradeCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockUpgradeServiceStartUpgradeCall) Do(f func(context.Context, upgrade0.UUID) error) *MockUpgradeServiceStartUpgradeCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockUpgradeServiceStartUpgradeCall) DoAndReturn(f func(context.Context, upgrade0.UUID) error) *MockUpgradeServiceStartUpgradeCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// UpgradeInfo mocks base method.
func (m *MockUpgradeService) UpgradeInfo(arg0 context.Context, arg1 upgrade0.UUID) (upgrade.Info, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpgradeInfo", arg0, arg1)
	ret0, _ := ret[0].(upgrade.Info)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpgradeInfo indicates an expected call of UpgradeInfo.
func (mr *MockUpgradeServiceMockRecorder) UpgradeInfo(arg0, arg1 any) *MockUpgradeServiceUpgradeInfoCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpgradeInfo", reflect.TypeOf((*MockUpgradeService)(nil).UpgradeInfo), arg0, arg1)
	return &MockUpgradeServiceUpgradeInfoCall{Call: call}
}

// MockUpgradeServiceUpgradeInfoCall wrap *gomock.Call
type MockUpgradeServiceUpgradeInfoCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockUpgradeServiceUpgradeInfoCall) Return(arg0 upgrade.Info, arg1 error) *MockUpgradeServiceUpgradeInfoCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockUpgradeServiceUpgradeInfoCall) Do(f func(context.Context, upgrade0.UUID) (upgrade.Info, error)) *MockUpgradeServiceUpgradeInfoCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockUpgradeServiceUpgradeInfoCall) DoAndReturn(f func(context.Context, upgrade0.UUID) (upgrade.Info, error)) *MockUpgradeServiceUpgradeInfoCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// WatchForUpgradeReady mocks base method.
func (m *MockUpgradeService) WatchForUpgradeReady(arg0 context.Context, arg1 upgrade0.UUID) (watcher.Watcher[struct{}], error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WatchForUpgradeReady", arg0, arg1)
	ret0, _ := ret[0].(watcher.Watcher[struct{}])
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// WatchForUpgradeReady indicates an expected call of WatchForUpgradeReady.
func (mr *MockUpgradeServiceMockRecorder) WatchForUpgradeReady(arg0, arg1 any) *MockUpgradeServiceWatchForUpgradeReadyCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WatchForUpgradeReady", reflect.TypeOf((*MockUpgradeService)(nil).WatchForUpgradeReady), arg0, arg1)
	return &MockUpgradeServiceWatchForUpgradeReadyCall{Call: call}
}

// MockUpgradeServiceWatchForUpgradeReadyCall wrap *gomock.Call
type MockUpgradeServiceWatchForUpgradeReadyCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockUpgradeServiceWatchForUpgradeReadyCall) Return(arg0 watcher.Watcher[struct{}], arg1 error) *MockUpgradeServiceWatchForUpgradeReadyCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockUpgradeServiceWatchForUpgradeReadyCall) Do(f func(context.Context, upgrade0.UUID) (watcher.Watcher[struct{}], error)) *MockUpgradeServiceWatchForUpgradeReadyCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockUpgradeServiceWatchForUpgradeReadyCall) DoAndReturn(f func(context.Context, upgrade0.UUID) (watcher.Watcher[struct{}], error)) *MockUpgradeServiceWatchForUpgradeReadyCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// WatchForUpgradeState mocks base method.
func (m *MockUpgradeService) WatchForUpgradeState(arg0 context.Context, arg1 upgrade0.UUID, arg2 upgrade.State) (watcher.Watcher[struct{}], error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WatchForUpgradeState", arg0, arg1, arg2)
	ret0, _ := ret[0].(watcher.Watcher[struct{}])
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// WatchForUpgradeState indicates an expected call of WatchForUpgradeState.
func (mr *MockUpgradeServiceMockRecorder) WatchForUpgradeState(arg0, arg1, arg2 any) *MockUpgradeServiceWatchForUpgradeStateCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WatchForUpgradeState", reflect.TypeOf((*MockUpgradeService)(nil).WatchForUpgradeState), arg0, arg1, arg2)
	return &MockUpgradeServiceWatchForUpgradeStateCall{Call: call}
}

// MockUpgradeServiceWatchForUpgradeStateCall wrap *gomock.Call
type MockUpgradeServiceWatchForUpgradeStateCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockUpgradeServiceWatchForUpgradeStateCall) Return(arg0 watcher.Watcher[struct{}], arg1 error) *MockUpgradeServiceWatchForUpgradeStateCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockUpgradeServiceWatchForUpgradeStateCall) Do(f func(context.Context, upgrade0.UUID, upgrade.State) (watcher.Watcher[struct{}], error)) *MockUpgradeServiceWatchForUpgradeStateCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockUpgradeServiceWatchForUpgradeStateCall) DoAndReturn(f func(context.Context, upgrade0.UUID, upgrade.State) (watcher.Watcher[struct{}], error)) *MockUpgradeServiceWatchForUpgradeStateCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

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

// ListModelIDs mocks base method.
func (m *MockModelService) ListModelIDs(arg0 context.Context) ([]model.UUID, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListModelIDs", arg0)
	ret0, _ := ret[0].([]model.UUID)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListModelIDs indicates an expected call of ListModelIDs.
func (mr *MockModelServiceMockRecorder) ListModelIDs(arg0 any) *MockModelServiceListModelIDsCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListModelIDs", reflect.TypeOf((*MockModelService)(nil).ListModelIDs), arg0)
	return &MockModelServiceListModelIDsCall{Call: call}
}

// MockModelServiceListModelIDsCall wrap *gomock.Call
type MockModelServiceListModelIDsCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockModelServiceListModelIDsCall) Return(arg0 []model.UUID, arg1 error) *MockModelServiceListModelIDsCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockModelServiceListModelIDsCall) Do(f func(context.Context) ([]model.UUID, error)) *MockModelServiceListModelIDsCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockModelServiceListModelIDsCall) DoAndReturn(f func(context.Context) ([]model.UUID, error)) *MockModelServiceListModelIDsCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}
