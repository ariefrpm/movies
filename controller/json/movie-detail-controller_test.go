package controller

import (
	"errors"
	"github.com/ariefrpm/movies/entity"
	usecase_mock "github.com/ariefrpm/movies/mock/usecase"
	"github.com/ariefrpm/movies/usecase"
	"github.com/golang/mock/gomock"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_movieDetailController_MovieDetail(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	tests := []struct {
		name   string
		mock func() (*http.Request, usecase.MovieDetailUseCase)
	}{
		{
			name: "validation error",
			mock: func() (*http.Request, usecase.MovieDetailUseCase) {
				movieUseCase := usecase_mock.NewMockMovieDetailUseCase(ctl)
				req := httptest.NewRequest("GET", "http://example.com/foo?a=123", nil)
				return req, movieUseCase
			},
		},
		{
			name: "usecase error",
			mock: func() (*http.Request, usecase.MovieDetailUseCase) {
				movieUseCase := usecase_mock.NewMockMovieDetailUseCase(ctl)
				req := httptest.NewRequest("GET", "http://example.com/foo?i=123", nil)
				movieUseCase.EXPECT().MovieDetail(gomock.Any()).Return(nil, errors.New("error")).Times(1)
				return req, movieUseCase
			},
		},
		{
			name: "success",
			mock: func() (*http.Request, usecase.MovieDetailUseCase) {
				movieUseCase := usecase_mock.NewMockMovieDetailUseCase(ctl)
				movie := &entity.MovieDetail{
					Title:  "title",
				}
				req := httptest.NewRequest("GET", "http://example.com/foo?i=123", nil)
				movieUseCase.EXPECT().MovieDetail(gomock.Any()).Return(movie, nil).Times(1)
				return req, movieUseCase
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, movieUseCase := tt.mock()
			m := &movieDetailController{
				movieDetailUseCase: movieUseCase,
			}
			w := httptest.NewRecorder()
			m.MovieDetail(w, req)
		})
	}
}