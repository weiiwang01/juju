// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/juju/juju/cmd/juju/application/deployer (interfaces: Resolver,Bundle)
//
// Generated by this command:
//
//	mockgen -typed -package mocks -destination mocks/resolver_mock.go github.com/juju/juju/cmd/juju/application/deployer Resolver,Bundle
//

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	charm "github.com/juju/charm/v13"
	charm0 "github.com/juju/juju/api/common/charm"
	base "github.com/juju/juju/core/base"
	gomock "go.uber.org/mock/gomock"
)

// MockResolver is a mock of Resolver interface.
type MockResolver struct {
	ctrl     *gomock.Controller
	recorder *MockResolverMockRecorder
}

// MockResolverMockRecorder is the mock recorder for MockResolver.
type MockResolverMockRecorder struct {
	mock *MockResolver
}

// NewMockResolver creates a new mock instance.
func NewMockResolver(ctrl *gomock.Controller) *MockResolver {
	mock := &MockResolver{ctrl: ctrl}
	mock.recorder = &MockResolverMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockResolver) EXPECT() *MockResolverMockRecorder {
	return m.recorder
}

// GetBundle mocks base method.
func (m *MockResolver) GetBundle(arg0 context.Context, arg1 *charm.URL, arg2 charm0.Origin, arg3 string) (charm.Bundle, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBundle", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(charm.Bundle)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBundle indicates an expected call of GetBundle.
func (mr *MockResolverMockRecorder) GetBundle(arg0, arg1, arg2, arg3 any) *MockResolverGetBundleCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBundle", reflect.TypeOf((*MockResolver)(nil).GetBundle), arg0, arg1, arg2, arg3)
	return &MockResolverGetBundleCall{Call: call}
}

// MockResolverGetBundleCall wrap *gomock.Call
type MockResolverGetBundleCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockResolverGetBundleCall) Return(arg0 charm.Bundle, arg1 error) *MockResolverGetBundleCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockResolverGetBundleCall) Do(f func(context.Context, *charm.URL, charm0.Origin, string) (charm.Bundle, error)) *MockResolverGetBundleCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockResolverGetBundleCall) DoAndReturn(f func(context.Context, *charm.URL, charm0.Origin, string) (charm.Bundle, error)) *MockResolverGetBundleCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// ResolveBundleURL mocks base method.
func (m *MockResolver) ResolveBundleURL(arg0 *charm.URL, arg1 charm0.Origin) (*charm.URL, charm0.Origin, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ResolveBundleURL", arg0, arg1)
	ret0, _ := ret[0].(*charm.URL)
	ret1, _ := ret[1].(charm0.Origin)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ResolveBundleURL indicates an expected call of ResolveBundleURL.
func (mr *MockResolverMockRecorder) ResolveBundleURL(arg0, arg1 any) *MockResolverResolveBundleURLCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ResolveBundleURL", reflect.TypeOf((*MockResolver)(nil).ResolveBundleURL), arg0, arg1)
	return &MockResolverResolveBundleURLCall{Call: call}
}

// MockResolverResolveBundleURLCall wrap *gomock.Call
type MockResolverResolveBundleURLCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockResolverResolveBundleURLCall) Return(arg0 *charm.URL, arg1 charm0.Origin, arg2 error) *MockResolverResolveBundleURLCall {
	c.Call = c.Call.Return(arg0, arg1, arg2)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockResolverResolveBundleURLCall) Do(f func(*charm.URL, charm0.Origin) (*charm.URL, charm0.Origin, error)) *MockResolverResolveBundleURLCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockResolverResolveBundleURLCall) DoAndReturn(f func(*charm.URL, charm0.Origin) (*charm.URL, charm0.Origin, error)) *MockResolverResolveBundleURLCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// ResolveCharm mocks base method.
func (m *MockResolver) ResolveCharm(arg0 *charm.URL, arg1 charm0.Origin, arg2 bool) (*charm.URL, charm0.Origin, []base.Base, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ResolveCharm", arg0, arg1, arg2)
	ret0, _ := ret[0].(*charm.URL)
	ret1, _ := ret[1].(charm0.Origin)
	ret2, _ := ret[2].([]base.Base)
	ret3, _ := ret[3].(error)
	return ret0, ret1, ret2, ret3
}

// ResolveCharm indicates an expected call of ResolveCharm.
func (mr *MockResolverMockRecorder) ResolveCharm(arg0, arg1, arg2 any) *MockResolverResolveCharmCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ResolveCharm", reflect.TypeOf((*MockResolver)(nil).ResolveCharm), arg0, arg1, arg2)
	return &MockResolverResolveCharmCall{Call: call}
}

// MockResolverResolveCharmCall wrap *gomock.Call
type MockResolverResolveCharmCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockResolverResolveCharmCall) Return(arg0 *charm.URL, arg1 charm0.Origin, arg2 []base.Base, arg3 error) *MockResolverResolveCharmCall {
	c.Call = c.Call.Return(arg0, arg1, arg2, arg3)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockResolverResolveCharmCall) Do(f func(*charm.URL, charm0.Origin, bool) (*charm.URL, charm0.Origin, []base.Base, error)) *MockResolverResolveCharmCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockResolverResolveCharmCall) DoAndReturn(f func(*charm.URL, charm0.Origin, bool) (*charm.URL, charm0.Origin, []base.Base, error)) *MockResolverResolveCharmCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// MockBundle is a mock of Bundle interface.
type MockBundle struct {
	ctrl     *gomock.Controller
	recorder *MockBundleMockRecorder
}

// MockBundleMockRecorder is the mock recorder for MockBundle.
type MockBundleMockRecorder struct {
	mock *MockBundle
}

// NewMockBundle creates a new mock instance.
func NewMockBundle(ctrl *gomock.Controller) *MockBundle {
	mock := &MockBundle{ctrl: ctrl}
	mock.recorder = &MockBundleMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBundle) EXPECT() *MockBundleMockRecorder {
	return m.recorder
}

// ContainsOverlays mocks base method.
func (m *MockBundle) ContainsOverlays() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ContainsOverlays")
	ret0, _ := ret[0].(bool)
	return ret0
}

// ContainsOverlays indicates an expected call of ContainsOverlays.
func (mr *MockBundleMockRecorder) ContainsOverlays() *MockBundleContainsOverlaysCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ContainsOverlays", reflect.TypeOf((*MockBundle)(nil).ContainsOverlays))
	return &MockBundleContainsOverlaysCall{Call: call}
}

// MockBundleContainsOverlaysCall wrap *gomock.Call
type MockBundleContainsOverlaysCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockBundleContainsOverlaysCall) Return(arg0 bool) *MockBundleContainsOverlaysCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockBundleContainsOverlaysCall) Do(f func() bool) *MockBundleContainsOverlaysCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockBundleContainsOverlaysCall) DoAndReturn(f func() bool) *MockBundleContainsOverlaysCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Data mocks base method.
func (m *MockBundle) Data() *charm.BundleData {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Data")
	ret0, _ := ret[0].(*charm.BundleData)
	return ret0
}

// Data indicates an expected call of Data.
func (mr *MockBundleMockRecorder) Data() *MockBundleDataCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Data", reflect.TypeOf((*MockBundle)(nil).Data))
	return &MockBundleDataCall{Call: call}
}

// MockBundleDataCall wrap *gomock.Call
type MockBundleDataCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockBundleDataCall) Return(arg0 *charm.BundleData) *MockBundleDataCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockBundleDataCall) Do(f func() *charm.BundleData) *MockBundleDataCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockBundleDataCall) DoAndReturn(f func() *charm.BundleData) *MockBundleDataCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// ReadMe mocks base method.
func (m *MockBundle) ReadMe() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadMe")
	ret0, _ := ret[0].(string)
	return ret0
}

// ReadMe indicates an expected call of ReadMe.
func (mr *MockBundleMockRecorder) ReadMe() *MockBundleReadMeCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadMe", reflect.TypeOf((*MockBundle)(nil).ReadMe))
	return &MockBundleReadMeCall{Call: call}
}

// MockBundleReadMeCall wrap *gomock.Call
type MockBundleReadMeCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockBundleReadMeCall) Return(arg0 string) *MockBundleReadMeCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockBundleReadMeCall) Do(f func() string) *MockBundleReadMeCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockBundleReadMeCall) DoAndReturn(f func() string) *MockBundleReadMeCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}
