// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/juju/juju/internal/charm/repository (interfaces: CharmHubClient)
//
// Generated by this command:
//
//	mockgen -typed -package mocks -destination mocks/charmhub_client_mock.go github.com/juju/juju/internal/charm/repository CharmHubClient
//

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	url "net/url"
	reflect "reflect"

	charm "github.com/juju/charm/v13"
	charmhub "github.com/juju/juju/internal/charmhub"
	transport "github.com/juju/juju/internal/charmhub/transport"
	gomock "go.uber.org/mock/gomock"
)

// MockCharmHubClient is a mock of CharmHubClient interface.
type MockCharmHubClient struct {
	ctrl     *gomock.Controller
	recorder *MockCharmHubClientMockRecorder
}

// MockCharmHubClientMockRecorder is the mock recorder for MockCharmHubClient.
type MockCharmHubClientMockRecorder struct {
	mock *MockCharmHubClient
}

// NewMockCharmHubClient creates a new mock instance.
func NewMockCharmHubClient(ctrl *gomock.Controller) *MockCharmHubClient {
	mock := &MockCharmHubClient{ctrl: ctrl}
	mock.recorder = &MockCharmHubClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCharmHubClient) EXPECT() *MockCharmHubClientMockRecorder {
	return m.recorder
}

// DownloadAndRead mocks base method.
func (m *MockCharmHubClient) DownloadAndRead(arg0 context.Context, arg1 *url.URL, arg2 string, arg3 ...charmhub.DownloadOption) (*charm.CharmArchive, error) {
	m.ctrl.T.Helper()
	varargs := []any{arg0, arg1, arg2}
	for _, a := range arg3 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DownloadAndRead", varargs...)
	ret0, _ := ret[0].(*charm.CharmArchive)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DownloadAndRead indicates an expected call of DownloadAndRead.
func (mr *MockCharmHubClientMockRecorder) DownloadAndRead(arg0, arg1, arg2 any, arg3 ...any) *MockCharmHubClientDownloadAndReadCall {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0, arg1, arg2}, arg3...)
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DownloadAndRead", reflect.TypeOf((*MockCharmHubClient)(nil).DownloadAndRead), varargs...)
	return &MockCharmHubClientDownloadAndReadCall{Call: call}
}

// MockCharmHubClientDownloadAndReadCall wrap *gomock.Call
type MockCharmHubClientDownloadAndReadCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockCharmHubClientDownloadAndReadCall) Return(arg0 *charm.CharmArchive, arg1 error) *MockCharmHubClientDownloadAndReadCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockCharmHubClientDownloadAndReadCall) Do(f func(context.Context, *url.URL, string, ...charmhub.DownloadOption) (*charm.CharmArchive, error)) *MockCharmHubClientDownloadAndReadCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockCharmHubClientDownloadAndReadCall) DoAndReturn(f func(context.Context, *url.URL, string, ...charmhub.DownloadOption) (*charm.CharmArchive, error)) *MockCharmHubClientDownloadAndReadCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// ListResourceRevisions mocks base method.
func (m *MockCharmHubClient) ListResourceRevisions(arg0 context.Context, arg1, arg2 string) ([]transport.ResourceRevision, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListResourceRevisions", arg0, arg1, arg2)
	ret0, _ := ret[0].([]transport.ResourceRevision)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListResourceRevisions indicates an expected call of ListResourceRevisions.
func (mr *MockCharmHubClientMockRecorder) ListResourceRevisions(arg0, arg1, arg2 any) *MockCharmHubClientListResourceRevisionsCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListResourceRevisions", reflect.TypeOf((*MockCharmHubClient)(nil).ListResourceRevisions), arg0, arg1, arg2)
	return &MockCharmHubClientListResourceRevisionsCall{Call: call}
}

// MockCharmHubClientListResourceRevisionsCall wrap *gomock.Call
type MockCharmHubClientListResourceRevisionsCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockCharmHubClientListResourceRevisionsCall) Return(arg0 []transport.ResourceRevision, arg1 error) *MockCharmHubClientListResourceRevisionsCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockCharmHubClientListResourceRevisionsCall) Do(f func(context.Context, string, string) ([]transport.ResourceRevision, error)) *MockCharmHubClientListResourceRevisionsCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockCharmHubClientListResourceRevisionsCall) DoAndReturn(f func(context.Context, string, string) ([]transport.ResourceRevision, error)) *MockCharmHubClientListResourceRevisionsCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Refresh mocks base method.
func (m *MockCharmHubClient) Refresh(arg0 context.Context, arg1 charmhub.RefreshConfig) ([]transport.RefreshResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Refresh", arg0, arg1)
	ret0, _ := ret[0].([]transport.RefreshResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Refresh indicates an expected call of Refresh.
func (mr *MockCharmHubClientMockRecorder) Refresh(arg0, arg1 any) *MockCharmHubClientRefreshCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Refresh", reflect.TypeOf((*MockCharmHubClient)(nil).Refresh), arg0, arg1)
	return &MockCharmHubClientRefreshCall{Call: call}
}

// MockCharmHubClientRefreshCall wrap *gomock.Call
type MockCharmHubClientRefreshCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockCharmHubClientRefreshCall) Return(arg0 []transport.RefreshResponse, arg1 error) *MockCharmHubClientRefreshCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockCharmHubClientRefreshCall) Do(f func(context.Context, charmhub.RefreshConfig) ([]transport.RefreshResponse, error)) *MockCharmHubClientRefreshCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockCharmHubClientRefreshCall) DoAndReturn(f func(context.Context, charmhub.RefreshConfig) ([]transport.RefreshResponse, error)) *MockCharmHubClientRefreshCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}
