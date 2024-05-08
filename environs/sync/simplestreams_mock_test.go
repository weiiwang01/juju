// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/juju/juju/environs/tools (interfaces: SimplestreamsFetcher)
//
// Generated by this command:
//
//	mockgen -typed -package sync_test -destination simplestreams_mock_test.go github.com/juju/juju/environs/tools SimplestreamsFetcher
//

// Package sync_test is a generated GoMock package.
package sync_test

import (
	context "context"
	reflect "reflect"

	simplestreams "github.com/juju/juju/environs/simplestreams"
	gomock "go.uber.org/mock/gomock"
)

// MockSimplestreamsFetcher is a mock of SimplestreamsFetcher interface.
type MockSimplestreamsFetcher struct {
	ctrl     *gomock.Controller
	recorder *MockSimplestreamsFetcherMockRecorder
}

// MockSimplestreamsFetcherMockRecorder is the mock recorder for MockSimplestreamsFetcher.
type MockSimplestreamsFetcherMockRecorder struct {
	mock *MockSimplestreamsFetcher
}

// NewMockSimplestreamsFetcher creates a new mock instance.
func NewMockSimplestreamsFetcher(ctrl *gomock.Controller) *MockSimplestreamsFetcher {
	mock := &MockSimplestreamsFetcher{ctrl: ctrl}
	mock.recorder = &MockSimplestreamsFetcherMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSimplestreamsFetcher) EXPECT() *MockSimplestreamsFetcherMockRecorder {
	return m.recorder
}

// GetMetadata mocks base method.
func (m *MockSimplestreamsFetcher) GetMetadata(arg0 context.Context, arg1 []simplestreams.DataSource, arg2 simplestreams.GetMetadataParams) ([]any, *simplestreams.ResolveInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMetadata", arg0, arg1, arg2)
	ret0, _ := ret[0].([]any)
	ret1, _ := ret[1].(*simplestreams.ResolveInfo)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetMetadata indicates an expected call of GetMetadata.
func (mr *MockSimplestreamsFetcherMockRecorder) GetMetadata(arg0, arg1, arg2 any) *MockSimplestreamsFetcherGetMetadataCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMetadata", reflect.TypeOf((*MockSimplestreamsFetcher)(nil).GetMetadata), arg0, arg1, arg2)
	return &MockSimplestreamsFetcherGetMetadataCall{Call: call}
}

// MockSimplestreamsFetcherGetMetadataCall wrap *gomock.Call
type MockSimplestreamsFetcherGetMetadataCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockSimplestreamsFetcherGetMetadataCall) Return(arg0 []any, arg1 *simplestreams.ResolveInfo, arg2 error) *MockSimplestreamsFetcherGetMetadataCall {
	c.Call = c.Call.Return(arg0, arg1, arg2)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockSimplestreamsFetcherGetMetadataCall) Do(f func(context.Context, []simplestreams.DataSource, simplestreams.GetMetadataParams) ([]any, *simplestreams.ResolveInfo, error)) *MockSimplestreamsFetcherGetMetadataCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockSimplestreamsFetcherGetMetadataCall) DoAndReturn(f func(context.Context, []simplestreams.DataSource, simplestreams.GetMetadataParams) ([]any, *simplestreams.ResolveInfo, error)) *MockSimplestreamsFetcherGetMetadataCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// NewDataSource mocks base method.
func (m *MockSimplestreamsFetcher) NewDataSource(arg0 simplestreams.Config) simplestreams.DataSource {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewDataSource", arg0)
	ret0, _ := ret[0].(simplestreams.DataSource)
	return ret0
}

// NewDataSource indicates an expected call of NewDataSource.
func (mr *MockSimplestreamsFetcherMockRecorder) NewDataSource(arg0 any) *MockSimplestreamsFetcherNewDataSourceCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewDataSource", reflect.TypeOf((*MockSimplestreamsFetcher)(nil).NewDataSource), arg0)
	return &MockSimplestreamsFetcherNewDataSourceCall{Call: call}
}

// MockSimplestreamsFetcherNewDataSourceCall wrap *gomock.Call
type MockSimplestreamsFetcherNewDataSourceCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockSimplestreamsFetcherNewDataSourceCall) Return(arg0 simplestreams.DataSource) *MockSimplestreamsFetcherNewDataSourceCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockSimplestreamsFetcherNewDataSourceCall) Do(f func(simplestreams.Config) simplestreams.DataSource) *MockSimplestreamsFetcherNewDataSourceCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockSimplestreamsFetcherNewDataSourceCall) DoAndReturn(f func(simplestreams.Config) simplestreams.DataSource) *MockSimplestreamsFetcherNewDataSourceCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}
