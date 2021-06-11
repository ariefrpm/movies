// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/ariefrpm/movies/usecase (interfaces: MovieDetailUseCase)

// Package usecase_mock is a generated GoMock package.
package usecase_mock

import (
	reflect "reflect"

	entity "github.com/ariefrpm/movies/entity"
	gomock "github.com/golang/mock/gomock"
)

// MockMovieDetailUseCase is a mock of MovieDetailUseCase interface.
type MockMovieDetailUseCase struct {
	ctrl     *gomock.Controller
	recorder *MockMovieDetailUseCaseMockRecorder
}

// MockMovieDetailUseCaseMockRecorder is the mock recorder for MockMovieDetailUseCase.
type MockMovieDetailUseCaseMockRecorder struct {
	mock *MockMovieDetailUseCase
}

// NewMockMovieDetailUseCase creates a new mock instance.
func NewMockMovieDetailUseCase(ctrl *gomock.Controller) *MockMovieDetailUseCase {
	mock := &MockMovieDetailUseCase{ctrl: ctrl}
	mock.recorder = &MockMovieDetailUseCaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMovieDetailUseCase) EXPECT() *MockMovieDetailUseCaseMockRecorder {
	return m.recorder
}

// MovieDetail mocks base method.
func (m *MockMovieDetailUseCase) MovieDetail(arg0 string) (*entity.MovieDetail, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MovieDetail", arg0)
	ret0, _ := ret[0].(*entity.MovieDetail)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MovieDetail indicates an expected call of MovieDetail.
func (mr *MockMovieDetailUseCaseMockRecorder) MovieDetail(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MovieDetail", reflect.TypeOf((*MockMovieDetailUseCase)(nil).MovieDetail), arg0)
}
