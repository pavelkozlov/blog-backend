// Code generated by MockGen. DO NOT EDIT.
// Source: mocks.go

// Package mock_database is a generated GoMock package.
package mock_database

import (
	database "blog/pkg/database"
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	pgconn "github.com/jackc/pgconn"
	pgproto3 "github.com/jackc/pgproto3/v2"
	pgx "github.com/jackc/pgx/v4"
)

// MockPostgres is a mock of Postgres interface.
type MockPostgres struct {
	ctrl     *gomock.Controller
	recorder *MockPostgresMockRecorder
}

// MockPostgresMockRecorder is the mock recorder for MockPostgres.
type MockPostgresMockRecorder struct {
	mock *MockPostgres
}

// NewMockPostgres creates a new mock instance.
func NewMockPostgres(ctrl *gomock.Controller) *MockPostgres {
	mock := &MockPostgres{ctrl: ctrl}
	mock.recorder = &MockPostgresMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPostgres) EXPECT() *MockPostgresMockRecorder {
	return m.recorder
}

// Acquire mocks base method.
func (m *MockPostgres) Acquire() (database.PgxPoolConn, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Acquire")
	ret0, _ := ret[0].(database.PgxPoolConn)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Acquire indicates an expected call of Acquire.
func (mr *MockPostgresMockRecorder) Acquire() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Acquire", reflect.TypeOf((*MockPostgres)(nil).Acquire))
}

// MockRows is a mock of Rows interface.
type MockRows struct {
	ctrl     *gomock.Controller
	recorder *MockRowsMockRecorder
}

// MockRowsMockRecorder is the mock recorder for MockRows.
type MockRowsMockRecorder struct {
	mock *MockRows
}

// NewMockRows creates a new mock instance.
func NewMockRows(ctrl *gomock.Controller) *MockRows {
	mock := &MockRows{ctrl: ctrl}
	mock.recorder = &MockRowsMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRows) EXPECT() *MockRowsMockRecorder {
	return m.recorder
}

// Close mocks base method.
func (m *MockRows) Close() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Close")
}

// Close indicates an expected call of Close.
func (mr *MockRowsMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockRows)(nil).Close))
}

// CommandTag mocks base method.
func (m *MockRows) CommandTag() pgconn.CommandTag {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CommandTag")
	ret0, _ := ret[0].(pgconn.CommandTag)
	return ret0
}

// CommandTag indicates an expected call of CommandTag.
func (mr *MockRowsMockRecorder) CommandTag() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CommandTag", reflect.TypeOf((*MockRows)(nil).CommandTag))
}

// Err mocks base method.
func (m *MockRows) Err() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Err")
	ret0, _ := ret[0].(error)
	return ret0
}

// Err indicates an expected call of Err.
func (mr *MockRowsMockRecorder) Err() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Err", reflect.TypeOf((*MockRows)(nil).Err))
}

// FieldDescriptions mocks base method.
func (m *MockRows) FieldDescriptions() []pgproto3.FieldDescription {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FieldDescriptions")
	ret0, _ := ret[0].([]pgproto3.FieldDescription)
	return ret0
}

// FieldDescriptions indicates an expected call of FieldDescriptions.
func (mr *MockRowsMockRecorder) FieldDescriptions() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FieldDescriptions", reflect.TypeOf((*MockRows)(nil).FieldDescriptions))
}

// Next mocks base method.
func (m *MockRows) Next() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Next")
	ret0, _ := ret[0].(bool)
	return ret0
}

// Next indicates an expected call of Next.
func (mr *MockRowsMockRecorder) Next() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Next", reflect.TypeOf((*MockRows)(nil).Next))
}

// RawValues mocks base method.
func (m *MockRows) RawValues() [][]byte {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RawValues")
	ret0, _ := ret[0].([][]byte)
	return ret0
}

// RawValues indicates an expected call of RawValues.
func (mr *MockRowsMockRecorder) RawValues() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RawValues", reflect.TypeOf((*MockRows)(nil).RawValues))
}

// Scan mocks base method.
func (m *MockRows) Scan(dest ...interface{}) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range dest {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Scan", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// Scan indicates an expected call of Scan.
func (mr *MockRowsMockRecorder) Scan(dest ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Scan", reflect.TypeOf((*MockRows)(nil).Scan), dest...)
}

// Values mocks base method.
func (m *MockRows) Values() ([]interface{}, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Values")
	ret0, _ := ret[0].([]interface{})
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Values indicates an expected call of Values.
func (mr *MockRowsMockRecorder) Values() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Values", reflect.TypeOf((*MockRows)(nil).Values))
}

// MockPgxPoolConn is a mock of PgxPoolConn interface.
type MockPgxPoolConn struct {
	ctrl     *gomock.Controller
	recorder *MockPgxPoolConnMockRecorder
}

// MockPgxPoolConnMockRecorder is the mock recorder for MockPgxPoolConn.
type MockPgxPoolConnMockRecorder struct {
	mock *MockPgxPoolConn
}

// NewMockPgxPoolConn creates a new mock instance.
func NewMockPgxPoolConn(ctrl *gomock.Controller) *MockPgxPoolConn {
	mock := &MockPgxPoolConn{ctrl: ctrl}
	mock.recorder = &MockPgxPoolConnMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPgxPoolConn) EXPECT() *MockPgxPoolConnMockRecorder {
	return m.recorder
}

// Query mocks base method.
func (m *MockPgxPoolConn) Query(ctx context.Context, sql string, args ...interface{}) (pgx.Rows, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, sql}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Query", varargs...)
	ret0, _ := ret[0].(pgx.Rows)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Query indicates an expected call of Query.
func (mr *MockPgxPoolConnMockRecorder) Query(ctx, sql interface{}, args ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, sql}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Query", reflect.TypeOf((*MockPgxPoolConn)(nil).Query), varargs...)
}

// Release mocks base method.
func (m *MockPgxPoolConn) Release() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Release")
}

// Release indicates an expected call of Release.
func (mr *MockPgxPoolConnMockRecorder) Release() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Release", reflect.TypeOf((*MockPgxPoolConn)(nil).Release))
}