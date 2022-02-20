// Code generated by MockGen. DO NOT EDIT.
// Source: blog.go

// Package mock_usecase is a generated GoMock package.
package mock_usecase

import (
	post "blog/internal/post"
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockBlog is a mock of Blog interface.
type MockBlog struct {
	ctrl     *gomock.Controller
	recorder *MockBlogMockRecorder
}

// MockBlogMockRecorder is the mock recorder for MockBlog.
type MockBlogMockRecorder struct {
	mock *MockBlog
}

// NewMockBlog creates a new mock instance.
func NewMockBlog(ctrl *gomock.Controller) *MockBlog {
	mock := &MockBlog{ctrl: ctrl}
	mock.recorder = &MockBlogMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBlog) EXPECT() *MockBlogMockRecorder {
	return m.recorder
}

// ListAllBlogs mocks base method.
func (m *MockBlog) ListAllBlogs(ctx context.Context, limit, offset int) ([]*post.Blog, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListAllBlogs", ctx, limit, offset)
	ret0, _ := ret[0].([]*post.Blog)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListAllBlogs indicates an expected call of ListAllBlogs.
func (mr *MockBlogMockRecorder) ListAllBlogs(ctx, limit, offset interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAllBlogs", reflect.TypeOf((*MockBlog)(nil).ListAllBlogs), ctx, limit, offset)
}