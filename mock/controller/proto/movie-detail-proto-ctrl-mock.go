// Code generated by MockGen. DO NOT EDIT.
// Source: proto/movie-detail.pb.go

// Package controller_mock is a generated GoMock package.
package controller_mock

import (
	context "context"
	reflect "reflect"

	proto "github.com/ariefrpm/movies/proto"
	gomock "github.com/golang/mock/gomock"
	grpc "google.golang.org/grpc"
)

// MockMovieDetailServiceClient is a mock of MovieDetailServiceClient interface.
type MockMovieDetailServiceClient struct {
	ctrl     *gomock.Controller
	recorder *MockMovieDetailServiceClientMockRecorder
}

// MockMovieDetailServiceClientMockRecorder is the mock recorder for MockMovieDetailServiceClient.
type MockMovieDetailServiceClientMockRecorder struct {
	mock *MockMovieDetailServiceClient
}

// NewMockMovieDetailServiceClient creates a new mock instance.
func NewMockMovieDetailServiceClient(ctrl *gomock.Controller) *MockMovieDetailServiceClient {
	mock := &MockMovieDetailServiceClient{ctrl: ctrl}
	mock.recorder = &MockMovieDetailServiceClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMovieDetailServiceClient) EXPECT() *MockMovieDetailServiceClientMockRecorder {
	return m.recorder
}

// MovieDetail mocks base method.
func (m *MockMovieDetailServiceClient) MovieDetail(ctx context.Context, in *proto.MovieDetailRequest, opts ...grpc.CallOption) (*proto.MovieDetailResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "MovieDetail", varargs...)
	ret0, _ := ret[0].(*proto.MovieDetailResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MovieDetail indicates an expected call of MovieDetail.
func (mr *MockMovieDetailServiceClientMockRecorder) MovieDetail(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MovieDetail", reflect.TypeOf((*MockMovieDetailServiceClient)(nil).MovieDetail), varargs...)
}

// MockMovieDetailServiceServer is a mock of MovieDetailServiceServer interface.
type MockMovieDetailServiceServer struct {
	ctrl     *gomock.Controller
	recorder *MockMovieDetailServiceServerMockRecorder
}

// MockMovieDetailServiceServerMockRecorder is the mock recorder for MockMovieDetailServiceServer.
type MockMovieDetailServiceServerMockRecorder struct {
	mock *MockMovieDetailServiceServer
}

// NewMockMovieDetailServiceServer creates a new mock instance.
func NewMockMovieDetailServiceServer(ctrl *gomock.Controller) *MockMovieDetailServiceServer {
	mock := &MockMovieDetailServiceServer{ctrl: ctrl}
	mock.recorder = &MockMovieDetailServiceServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMovieDetailServiceServer) EXPECT() *MockMovieDetailServiceServerMockRecorder {
	return m.recorder
}

// MovieDetail mocks base method.
func (m *MockMovieDetailServiceServer) MovieDetail(arg0 context.Context, arg1 *proto.MovieDetailRequest) (*proto.MovieDetailResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MovieDetail", arg0, arg1)
	ret0, _ := ret[0].(*proto.MovieDetailResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MovieDetail indicates an expected call of MovieDetail.
func (mr *MockMovieDetailServiceServerMockRecorder) MovieDetail(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MovieDetail", reflect.TypeOf((*MockMovieDetailServiceServer)(nil).MovieDetail), arg0, arg1)
}