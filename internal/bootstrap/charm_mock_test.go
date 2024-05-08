// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/juju/juju/core/charm (interfaces: Repository)
//
// Generated by this command:
//
//	mockgen -typed -package bootstrap -destination charm_mock_test.go github.com/juju/juju/core/charm Repository
//

// Package bootstrap is a generated GoMock package.
package bootstrap

import (
	context "context"
	url "net/url"
	reflect "reflect"

	charm "github.com/juju/charm/v13"
	resource "github.com/juju/charm/v13/resource"
	charm0 "github.com/juju/juju/core/charm"
	gomock "go.uber.org/mock/gomock"
)

// MockRepository is a mock of Repository interface.
type MockRepository struct {
	ctrl     *gomock.Controller
	recorder *MockRepositoryMockRecorder
}

// MockRepositoryMockRecorder is the mock recorder for MockRepository.
type MockRepositoryMockRecorder struct {
	mock *MockRepository
}

// NewMockRepository creates a new mock instance.
func NewMockRepository(ctrl *gomock.Controller) *MockRepository {
	mock := &MockRepository{ctrl: ctrl}
	mock.recorder = &MockRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRepository) EXPECT() *MockRepositoryMockRecorder {
	return m.recorder
}

// DownloadCharm mocks base method.
func (m *MockRepository) DownloadCharm(arg0 context.Context, arg1 string, arg2 charm0.Origin, arg3 string) (charm0.CharmArchive, charm0.Origin, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DownloadCharm", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(charm0.CharmArchive)
	ret1, _ := ret[1].(charm0.Origin)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// DownloadCharm indicates an expected call of DownloadCharm.
func (mr *MockRepositoryMockRecorder) DownloadCharm(arg0, arg1, arg2, arg3 any) *MockRepositoryDownloadCharmCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DownloadCharm", reflect.TypeOf((*MockRepository)(nil).DownloadCharm), arg0, arg1, arg2, arg3)
	return &MockRepositoryDownloadCharmCall{Call: call}
}

// MockRepositoryDownloadCharmCall wrap *gomock.Call
type MockRepositoryDownloadCharmCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockRepositoryDownloadCharmCall) Return(arg0 charm0.CharmArchive, arg1 charm0.Origin, arg2 error) *MockRepositoryDownloadCharmCall {
	c.Call = c.Call.Return(arg0, arg1, arg2)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockRepositoryDownloadCharmCall) Do(f func(context.Context, string, charm0.Origin, string) (charm0.CharmArchive, charm0.Origin, error)) *MockRepositoryDownloadCharmCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockRepositoryDownloadCharmCall) DoAndReturn(f func(context.Context, string, charm0.Origin, string) (charm0.CharmArchive, charm0.Origin, error)) *MockRepositoryDownloadCharmCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// GetDownloadURL mocks base method.
func (m *MockRepository) GetDownloadURL(arg0 context.Context, arg1 string, arg2 charm0.Origin) (*url.URL, charm0.Origin, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDownloadURL", arg0, arg1, arg2)
	ret0, _ := ret[0].(*url.URL)
	ret1, _ := ret[1].(charm0.Origin)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetDownloadURL indicates an expected call of GetDownloadURL.
func (mr *MockRepositoryMockRecorder) GetDownloadURL(arg0, arg1, arg2 any) *MockRepositoryGetDownloadURLCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDownloadURL", reflect.TypeOf((*MockRepository)(nil).GetDownloadURL), arg0, arg1, arg2)
	return &MockRepositoryGetDownloadURLCall{Call: call}
}

// MockRepositoryGetDownloadURLCall wrap *gomock.Call
type MockRepositoryGetDownloadURLCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockRepositoryGetDownloadURLCall) Return(arg0 *url.URL, arg1 charm0.Origin, arg2 error) *MockRepositoryGetDownloadURLCall {
	c.Call = c.Call.Return(arg0, arg1, arg2)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockRepositoryGetDownloadURLCall) Do(f func(context.Context, string, charm0.Origin) (*url.URL, charm0.Origin, error)) *MockRepositoryGetDownloadURLCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockRepositoryGetDownloadURLCall) DoAndReturn(f func(context.Context, string, charm0.Origin) (*url.URL, charm0.Origin, error)) *MockRepositoryGetDownloadURLCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// GetEssentialMetadata mocks base method.
func (m *MockRepository) GetEssentialMetadata(arg0 context.Context, arg1 ...charm0.MetadataRequest) ([]charm0.EssentialMetadata, error) {
	m.ctrl.T.Helper()
	varargs := []any{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetEssentialMetadata", varargs...)
	ret0, _ := ret[0].([]charm0.EssentialMetadata)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetEssentialMetadata indicates an expected call of GetEssentialMetadata.
func (mr *MockRepositoryMockRecorder) GetEssentialMetadata(arg0 any, arg1 ...any) *MockRepositoryGetEssentialMetadataCall {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0}, arg1...)
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEssentialMetadata", reflect.TypeOf((*MockRepository)(nil).GetEssentialMetadata), varargs...)
	return &MockRepositoryGetEssentialMetadataCall{Call: call}
}

// MockRepositoryGetEssentialMetadataCall wrap *gomock.Call
type MockRepositoryGetEssentialMetadataCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockRepositoryGetEssentialMetadataCall) Return(arg0 []charm0.EssentialMetadata, arg1 error) *MockRepositoryGetEssentialMetadataCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockRepositoryGetEssentialMetadataCall) Do(f func(context.Context, ...charm0.MetadataRequest) ([]charm0.EssentialMetadata, error)) *MockRepositoryGetEssentialMetadataCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockRepositoryGetEssentialMetadataCall) DoAndReturn(f func(context.Context, ...charm0.MetadataRequest) ([]charm0.EssentialMetadata, error)) *MockRepositoryGetEssentialMetadataCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// ListResources mocks base method.
func (m *MockRepository) ListResources(arg0 context.Context, arg1 string, arg2 charm0.Origin) ([]resource.Resource, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListResources", arg0, arg1, arg2)
	ret0, _ := ret[0].([]resource.Resource)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListResources indicates an expected call of ListResources.
func (mr *MockRepositoryMockRecorder) ListResources(arg0, arg1, arg2 any) *MockRepositoryListResourcesCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListResources", reflect.TypeOf((*MockRepository)(nil).ListResources), arg0, arg1, arg2)
	return &MockRepositoryListResourcesCall{Call: call}
}

// MockRepositoryListResourcesCall wrap *gomock.Call
type MockRepositoryListResourcesCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockRepositoryListResourcesCall) Return(arg0 []resource.Resource, arg1 error) *MockRepositoryListResourcesCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockRepositoryListResourcesCall) Do(f func(context.Context, string, charm0.Origin) ([]resource.Resource, error)) *MockRepositoryListResourcesCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockRepositoryListResourcesCall) DoAndReturn(f func(context.Context, string, charm0.Origin) ([]resource.Resource, error)) *MockRepositoryListResourcesCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// ResolveForDeploy mocks base method.
func (m *MockRepository) ResolveForDeploy(arg0 context.Context, arg1 charm0.CharmID) (charm0.ResolvedDataForDeploy, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ResolveForDeploy", arg0, arg1)
	ret0, _ := ret[0].(charm0.ResolvedDataForDeploy)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ResolveForDeploy indicates an expected call of ResolveForDeploy.
func (mr *MockRepositoryMockRecorder) ResolveForDeploy(arg0, arg1 any) *MockRepositoryResolveForDeployCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ResolveForDeploy", reflect.TypeOf((*MockRepository)(nil).ResolveForDeploy), arg0, arg1)
	return &MockRepositoryResolveForDeployCall{Call: call}
}

// MockRepositoryResolveForDeployCall wrap *gomock.Call
type MockRepositoryResolveForDeployCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockRepositoryResolveForDeployCall) Return(arg0 charm0.ResolvedDataForDeploy, arg1 error) *MockRepositoryResolveForDeployCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockRepositoryResolveForDeployCall) Do(f func(context.Context, charm0.CharmID) (charm0.ResolvedDataForDeploy, error)) *MockRepositoryResolveForDeployCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockRepositoryResolveForDeployCall) DoAndReturn(f func(context.Context, charm0.CharmID) (charm0.ResolvedDataForDeploy, error)) *MockRepositoryResolveForDeployCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// ResolveResources mocks base method.
func (m *MockRepository) ResolveResources(arg0 context.Context, arg1 []resource.Resource, arg2 charm0.CharmID) ([]resource.Resource, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ResolveResources", arg0, arg1, arg2)
	ret0, _ := ret[0].([]resource.Resource)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ResolveResources indicates an expected call of ResolveResources.
func (mr *MockRepositoryMockRecorder) ResolveResources(arg0, arg1, arg2 any) *MockRepositoryResolveResourcesCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ResolveResources", reflect.TypeOf((*MockRepository)(nil).ResolveResources), arg0, arg1, arg2)
	return &MockRepositoryResolveResourcesCall{Call: call}
}

// MockRepositoryResolveResourcesCall wrap *gomock.Call
type MockRepositoryResolveResourcesCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockRepositoryResolveResourcesCall) Return(arg0 []resource.Resource, arg1 error) *MockRepositoryResolveResourcesCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockRepositoryResolveResourcesCall) Do(f func(context.Context, []resource.Resource, charm0.CharmID) ([]resource.Resource, error)) *MockRepositoryResolveResourcesCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockRepositoryResolveResourcesCall) DoAndReturn(f func(context.Context, []resource.Resource, charm0.CharmID) ([]resource.Resource, error)) *MockRepositoryResolveResourcesCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// ResolveWithPreferredChannel mocks base method.
func (m *MockRepository) ResolveWithPreferredChannel(arg0 context.Context, arg1 string, arg2 charm0.Origin) (*charm.URL, charm0.Origin, []charm0.Platform, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ResolveWithPreferredChannel", arg0, arg1, arg2)
	ret0, _ := ret[0].(*charm.URL)
	ret1, _ := ret[1].(charm0.Origin)
	ret2, _ := ret[2].([]charm0.Platform)
	ret3, _ := ret[3].(error)
	return ret0, ret1, ret2, ret3
}

// ResolveWithPreferredChannel indicates an expected call of ResolveWithPreferredChannel.
func (mr *MockRepositoryMockRecorder) ResolveWithPreferredChannel(arg0, arg1, arg2 any) *MockRepositoryResolveWithPreferredChannelCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ResolveWithPreferredChannel", reflect.TypeOf((*MockRepository)(nil).ResolveWithPreferredChannel), arg0, arg1, arg2)
	return &MockRepositoryResolveWithPreferredChannelCall{Call: call}
}

// MockRepositoryResolveWithPreferredChannelCall wrap *gomock.Call
type MockRepositoryResolveWithPreferredChannelCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockRepositoryResolveWithPreferredChannelCall) Return(arg0 *charm.URL, arg1 charm0.Origin, arg2 []charm0.Platform, arg3 error) *MockRepositoryResolveWithPreferredChannelCall {
	c.Call = c.Call.Return(arg0, arg1, arg2, arg3)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockRepositoryResolveWithPreferredChannelCall) Do(f func(context.Context, string, charm0.Origin) (*charm.URL, charm0.Origin, []charm0.Platform, error)) *MockRepositoryResolveWithPreferredChannelCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockRepositoryResolveWithPreferredChannelCall) DoAndReturn(f func(context.Context, string, charm0.Origin) (*charm.URL, charm0.Origin, []charm0.Platform, error)) *MockRepositoryResolveWithPreferredChannelCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}
