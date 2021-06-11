// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/ariefrpm/movies/controller/json (interfaces: SearchMovieController)

// Package controller_mock is a generated GoMock package.
package controller_mock

import (
	http "net/http"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockSearchMovieController is a mock of SearchMovieController interface.
type MockSearchMovieController struct {
	ctrl     *gomock.Controller
	recorder *MockSearchMovieControllerMockRecorder
}

// MockSearchMovieControllerMockRecorder is the mock recorder for MockSearchMovieController.
type MockSearchMovieControllerMockRecorder struct {
	mock *MockSearchMovieController
}

// NewMockSearchMovieController creates a new mock instance.
func NewMockSearchMovieController(ctrl *gomock.Controller) *MockSearchMovieController {
	mock := &MockSearchMovieController{ctrl: ctrl}
	mock.recorder = &MockSearchMovieControllerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSearchMovieController) EXPECT() *MockSearchMovieControllerMockRecorder {
	return m.recorder
}

// SearchMovie mocks base method.
func (m *MockSearchMovieController) SearchMovie(arg0 http.ResponseWriter, arg1 *http.Request) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SearchMovie", arg0, arg1)
}

// SearchMovie indicates an expected call of SearchMovie.
func (mr *MockSearchMovieControllerMockRecorder) SearchMovie(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SearchMovie", reflect.TypeOf((*MockSearchMovieController)(nil).SearchMovie), arg0, arg1)
}
