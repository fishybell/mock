// Code generated by MockGen. DO NOT EDIT.
// Source: input.go

// Package dot_imports is a generated GoMock package.
package dot_imports

import (
	bytes "bytes"
	. "context"
	. "net/http"
	reflect "reflect"

	gomock "github.com/fishybell/mock/gomock"
)

// MockWithDotImports is a mock of WithDotImports interface.
type MockWithDotImports struct {
	ctrl     *gomock.Controller
	recorder *MockWithDotImportsMockRecorder
}

// MockWithDotImportsMockRecorder is the mock recorder for MockWithDotImports.
type MockWithDotImportsMockRecorder struct {
	mock *MockWithDotImports
}

// NewMockWithDotImports creates a new mock instance.
func NewMockWithDotImports(ctrl *gomock.Controller) *MockWithDotImports {
	mock := &MockWithDotImports{ctrl: ctrl}
	mock.recorder = &MockWithDotImportsMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockWithDotImports) EXPECT() *MockWithDotImportsMockRecorder {
	return m.recorder
}

// Method1 mocks base method.
func (m *MockWithDotImports) Method1() Request {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Method1")
	ret0, _ := ret[0].(Request)
	return ret0
}

// Method1 indicates an expected call of Method1.
func (mr *MockWithDotImportsMockRecorder) Method1() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Method1", reflect.TypeOf((*MockWithDotImports)(nil).Method1))
}

// Method2 mocks base method.
func (m *MockWithDotImports) Method2() *bytes.Buffer {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Method2")
	ret0, _ := ret[0].(*bytes.Buffer)
	return ret0
}

// Method2 indicates an expected call of Method2.
func (mr *MockWithDotImportsMockRecorder) Method2() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Method2", reflect.TypeOf((*MockWithDotImports)(nil).Method2))
}

// Method3 mocks base method.
func (m *MockWithDotImports) Method3() Context {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Method3")
	ret0, _ := ret[0].(Context)
	return ret0
}

// Method3 indicates an expected call of Method3.
func (mr *MockWithDotImportsMockRecorder) Method3() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Method3", reflect.TypeOf((*MockWithDotImports)(nil).Method3))
}
