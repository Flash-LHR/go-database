// Code generated by MockGen. DO NOT EDIT.
// Source: client.go

// Package mockmysql is a generated GoMock package.
package mockmysql

import (
	context "context"
	sql "database/sql"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	sqlx "github.com/jmoiron/sqlx"
	mysql "trpc.group/trpc-go/trpc-database/mysql"
)

// MockClient is a mock of Client interface.
type MockClient struct {
	ctrl     *gomock.Controller
	recorder *MockClientMockRecorder
}

// MockClientMockRecorder is the mock recorder for MockClient.
type MockClientMockRecorder struct {
	mock *MockClient
}

// NewMockClient creates a new mock instance.
func NewMockClient(ctrl *gomock.Controller) *MockClient {
	mock := &MockClient{ctrl: ctrl}
	mock.recorder = &MockClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockClient) EXPECT() *MockClientMockRecorder {
	return m.recorder
}

// Exec mocks base method.
func (m *MockClient) Exec(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, query}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Exec", varargs...)
	ret0, _ := ret[0].(sql.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Exec indicates an expected call of Exec.
func (mr *MockClientMockRecorder) Exec(ctx, query interface{}, args ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, query}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Exec", reflect.TypeOf((*MockClient)(nil).Exec), varargs...)
}

// Get mocks base method.
func (m *MockClient) Get(ctx context.Context, dest interface{}, query string, args ...interface{}) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, dest, query}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Get", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// Get indicates an expected call of Get.
func (mr *MockClientMockRecorder) Get(ctx, dest, query interface{}, args ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, dest, query}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockClient)(nil).Get), varargs...)
}

// NamedExec mocks base method.
func (m *MockClient) NamedExec(ctx context.Context, query string, args interface{}) (sql.Result, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NamedExec", ctx, query, args)
	ret0, _ := ret[0].(sql.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NamedExec indicates an expected call of NamedExec.
func (mr *MockClientMockRecorder) NamedExec(ctx, query, args interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NamedExec", reflect.TypeOf((*MockClient)(nil).NamedExec), ctx, query, args)
}

// NamedQuery mocks base method.
func (m *MockClient) NamedQuery(ctx context.Context, query string, args interface{}) (*sqlx.Rows, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NamedQuery", ctx, query, args)
	ret0, _ := ret[0].(*sqlx.Rows)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NamedQuery indicates an expected call of NamedQuery.
func (mr *MockClientMockRecorder) NamedQuery(ctx, query, args interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NamedQuery", reflect.TypeOf((*MockClient)(nil).NamedQuery), ctx, query, args)
}

// Query mocks base method.
func (m *MockClient) Query(ctx context.Context, next mysql.NextFunc, query string, args ...interface{}) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, next, query}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Query", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// Query indicates an expected call of Query.
func (mr *MockClientMockRecorder) Query(ctx, next, query interface{}, args ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, next, query}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Query", reflect.TypeOf((*MockClient)(nil).Query), varargs...)
}

// QueryRow mocks base method.
func (m *MockClient) QueryRow(ctx context.Context, dest []interface{}, query string, args ...interface{}) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, dest, query}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "QueryRow", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// QueryRow indicates an expected call of QueryRow.
func (mr *MockClientMockRecorder) QueryRow(ctx, dest, query interface{}, args ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, dest, query}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueryRow", reflect.TypeOf((*MockClient)(nil).QueryRow), varargs...)
}

// QueryToStruct mocks base method.
func (m *MockClient) QueryToStruct(ctx context.Context, dst interface{}, query string, args ...interface{}) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, dst, query}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "QueryToStruct", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// QueryToStruct indicates an expected call of QueryToStruct.
func (mr *MockClientMockRecorder) QueryToStruct(ctx, dst, query interface{}, args ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, dst, query}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueryToStruct", reflect.TypeOf((*MockClient)(nil).QueryToStruct), varargs...)
}

// QueryToStructs mocks base method.
func (m *MockClient) QueryToStructs(ctx context.Context, dst interface{}, query string, args ...interface{}) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, dst, query}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "QueryToStructs", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// QueryToStructs indicates an expected call of QueryToStructs.
func (mr *MockClientMockRecorder) QueryToStructs(ctx, dst, query interface{}, args ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, dst, query}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueryToStructs", reflect.TypeOf((*MockClient)(nil).QueryToStructs), varargs...)
}

// Select mocks base method.
func (m *MockClient) Select(ctx context.Context, dest interface{}, query string, args ...interface{}) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, dest, query}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Select", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// Select indicates an expected call of Select.
func (mr *MockClientMockRecorder) Select(ctx, dest, query interface{}, args ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, dest, query}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Select", reflect.TypeOf((*MockClient)(nil).Select), varargs...)
}

// Transaction mocks base method.
func (m *MockClient) Transaction(ctx context.Context, fn mysql.TxFunc, opts ...mysql.TxOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, fn}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Transaction", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// Transaction indicates an expected call of Transaction.
func (mr *MockClientMockRecorder) Transaction(ctx, fn interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, fn}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Transaction", reflect.TypeOf((*MockClient)(nil).Transaction), varargs...)
}

// Transactionx mocks base method.
func (m *MockClient) Transactionx(ctx context.Context, fn mysql.TxxFunc, opts ...mysql.TxOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, fn}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Transactionx", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// Transactionx indicates an expected call of Transactionx.
func (mr *MockClientMockRecorder) Transactionx(ctx, fn interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, fn}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Transactionx", reflect.TypeOf((*MockClient)(nil).Transactionx), varargs...)
}
