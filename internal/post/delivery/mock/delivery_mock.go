// Code generated by MockGen. DO NOT EDIT.
// Source: http.go

// Package mock_delivery is a generated GoMock package.
package mock_delivery

import (
	http "net/http"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockBlogHandlers is a mock of BlogHandlers interface.
type MockBlogHandlers struct {
	ctrl     *gomock.Controller
	recorder *MockBlogHandlersMockRecorder
}

// MockBlogHandlersMockRecorder is the mock recorder for MockBlogHandlers.
type MockBlogHandlersMockRecorder struct {
	mock *MockBlogHandlers
}

// NewMockBlogHandlers creates a new mock instance.
func NewMockBlogHandlers(ctrl *gomock.Controller) *MockBlogHandlers {
	mock := &MockBlogHandlers{ctrl: ctrl}
	mock.recorder = &MockBlogHandlersMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBlogHandlers) EXPECT() *MockBlogHandlersMockRecorder {
	return m.recorder
}

// AllPosts mocks base method.
func (m *MockBlogHandlers) AllPosts() http.HandlerFunc {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AllPosts")
	ret0, _ := ret[0].(http.HandlerFunc)
	return ret0
}

// AllPosts indicates an expected call of AllPosts.
func (mr *MockBlogHandlersMockRecorder) AllPosts() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AllPosts", reflect.TypeOf((*MockBlogHandlers)(nil).AllPosts))
}

// CreatePost mocks base method.
func (m *MockBlogHandlers) CreatePost() http.HandlerFunc {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreatePost")
	ret0, _ := ret[0].(http.HandlerFunc)
	return ret0
}

// CreatePost indicates an expected call of CreatePost.
func (mr *MockBlogHandlersMockRecorder) CreatePost() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreatePost", reflect.TypeOf((*MockBlogHandlers)(nil).CreatePost))
}
