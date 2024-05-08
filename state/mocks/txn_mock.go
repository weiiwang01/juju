// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/juju/txn/v3 (interfaces: Runner)
//
// Generated by this command:
//
//	mockgen -typed -package mocks -destination mocks/txn_mock.go github.com/juju/txn/v3 Runner
//

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	txn "github.com/juju/txn/v3"
	gomock "go.uber.org/mock/gomock"
)

// MockRunner is a mock of Runner interface.
type MockRunner struct {
	ctrl     *gomock.Controller
	recorder *MockRunnerMockRecorder
}

// MockRunnerMockRecorder is the mock recorder for MockRunner.
type MockRunnerMockRecorder struct {
	mock *MockRunner
}

// NewMockRunner creates a new mock instance.
func NewMockRunner(ctrl *gomock.Controller) *MockRunner {
	mock := &MockRunner{ctrl: ctrl}
	mock.recorder = &MockRunnerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRunner) EXPECT() *MockRunnerMockRecorder {
	return m.recorder
}

// MaybePruneTransactions mocks base method.
func (m *MockRunner) MaybePruneTransactions(arg0 txn.PruneOptions) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MaybePruneTransactions", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// MaybePruneTransactions indicates an expected call of MaybePruneTransactions.
func (mr *MockRunnerMockRecorder) MaybePruneTransactions(arg0 any) *MockRunnerMaybePruneTransactionsCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MaybePruneTransactions", reflect.TypeOf((*MockRunner)(nil).MaybePruneTransactions), arg0)
	return &MockRunnerMaybePruneTransactionsCall{Call: call}
}

// MockRunnerMaybePruneTransactionsCall wrap *gomock.Call
type MockRunnerMaybePruneTransactionsCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockRunnerMaybePruneTransactionsCall) Return(arg0 error) *MockRunnerMaybePruneTransactionsCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockRunnerMaybePruneTransactionsCall) Do(f func(txn.PruneOptions) error) *MockRunnerMaybePruneTransactionsCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockRunnerMaybePruneTransactionsCall) DoAndReturn(f func(txn.PruneOptions) error) *MockRunnerMaybePruneTransactionsCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// ResumeTransactions mocks base method.
func (m *MockRunner) ResumeTransactions() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ResumeTransactions")
	ret0, _ := ret[0].(error)
	return ret0
}

// ResumeTransactions indicates an expected call of ResumeTransactions.
func (mr *MockRunnerMockRecorder) ResumeTransactions() *MockRunnerResumeTransactionsCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ResumeTransactions", reflect.TypeOf((*MockRunner)(nil).ResumeTransactions))
	return &MockRunnerResumeTransactionsCall{Call: call}
}

// MockRunnerResumeTransactionsCall wrap *gomock.Call
type MockRunnerResumeTransactionsCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockRunnerResumeTransactionsCall) Return(arg0 error) *MockRunnerResumeTransactionsCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockRunnerResumeTransactionsCall) Do(f func() error) *MockRunnerResumeTransactionsCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockRunnerResumeTransactionsCall) DoAndReturn(f func() error) *MockRunnerResumeTransactionsCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Run mocks base method.
func (m *MockRunner) Run(arg0 txn.TransactionSource) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Run", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Run indicates an expected call of Run.
func (mr *MockRunnerMockRecorder) Run(arg0 any) *MockRunnerRunCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Run", reflect.TypeOf((*MockRunner)(nil).Run), arg0)
	return &MockRunnerRunCall{Call: call}
}

// MockRunnerRunCall wrap *gomock.Call
type MockRunnerRunCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockRunnerRunCall) Return(arg0 error) *MockRunnerRunCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockRunnerRunCall) Do(f func(txn.TransactionSource) error) *MockRunnerRunCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockRunnerRunCall) DoAndReturn(f func(txn.TransactionSource) error) *MockRunnerRunCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// RunTransaction mocks base method.
func (m *MockRunner) RunTransaction(arg0 *txn.Transaction) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RunTransaction", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// RunTransaction indicates an expected call of RunTransaction.
func (mr *MockRunnerMockRecorder) RunTransaction(arg0 any) *MockRunnerRunTransactionCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RunTransaction", reflect.TypeOf((*MockRunner)(nil).RunTransaction), arg0)
	return &MockRunnerRunTransactionCall{Call: call}
}

// MockRunnerRunTransactionCall wrap *gomock.Call
type MockRunnerRunTransactionCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockRunnerRunTransactionCall) Return(arg0 error) *MockRunnerRunTransactionCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockRunnerRunTransactionCall) Do(f func(*txn.Transaction) error) *MockRunnerRunTransactionCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockRunnerRunTransactionCall) DoAndReturn(f func(*txn.Transaction) error) *MockRunnerRunTransactionCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}
