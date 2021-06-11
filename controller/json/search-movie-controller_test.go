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

func Test_searchMovieController_SearchMovie(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	tests := []struct {
		name   string
		mock func() (*http.Request, usecase.SearchMovieUseCase)
	}{
		{
			name: "validation error",
			mock: func() (request *http.Request, useCase usecase.SearchMovieUseCase) {
				movieUseCase := usecase_mock.NewMockSearchMovieUseCase(ctl)
				req := httptest.NewRequest("GET", "http://example.com/foo", nil)
				return req, movieUseCase
			},
		},
		{
			name: "usecase error",
			mock: func() (request *http.Request, useCase usecase.SearchMovieUseCase) {
				movieUseCase := usecase_mock.NewMockSearchMovieUseCase(ctl)
				req := httptest.NewRequest("GET", "http://example.com/foo?pagination=1&searchword=batman", nil)
				movieUseCase.EXPECT().SearchMovie(gomock.Any(), gomock.Any()).Return(nil, errors.New("error")).Times(1)
				return req, movieUseCase
			},
		},
		{
			name: "success",
			mock: func() (request *http.Request, useCase usecase.SearchMovieUseCase) {
				movieUseCase := usecase_mock.NewMockSearchMovieUseCase(ctl)
				movies := &entity.MovieList{
					Response:     "True",
				}
				req := httptest.NewRequest("GET", "http://example.com/foo?pagination=1&searchword=batman", nil)
				movieUseCase.EXPECT().SearchMovie(gomock.Any(), gomock.Any()).Return(movies, nil).Times(1)
				return req, movieUseCase
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, movieUseCase := tt.mock()
			s := &searchMovieController{
				searchMoveUseCase: movieUseCase,
			}
			w := httptest.NewRecorder()
			s.SearchMovie(w, req)
		})
	}
}