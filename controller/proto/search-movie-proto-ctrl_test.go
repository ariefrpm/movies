package controller

import (
	"errors"
	"github.com/ariefrpm/movies/entity"
	usecase_mock "github.com/ariefrpm/movies/mock/usecase"
	"github.com/ariefrpm/movies/proto"
	"github.com/ariefrpm/movies/usecase"
	"github.com/golang/mock/gomock"
	"golang.org/x/net/context"
	"reflect"
	"testing"
)

func Test_searchMovieService_SearchMovie(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	type args struct {
		ctx     context.Context
		request *proto.SearchMovieRequest
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
		mock func() (*proto.SearchMovieResponse, usecase.SearchMovieUseCase)
	}{
		{
			name:    "error",
			args:    args{
				request: &proto.SearchMovieRequest{
					Pagination: 1,
					Searchword: "batman",
				},
			},
			wantErr: true,
			mock: func() (*proto.SearchMovieResponse, usecase.SearchMovieUseCase) {
				movieUseCase := usecase_mock.NewMockSearchMovieUseCase(ctl)
				movieUseCase.EXPECT().SearchMovie(gomock.Any(), gomock.Any()).Return(nil, errors.New("error")).Times(1)
				return nil, movieUseCase
			},
		},
		{
			name:    "success",
			args:    args{
				request: &proto.SearchMovieRequest{
					Pagination: 1,
					Searchword: "batman",
				},
			},
			wantErr: false,
			mock: func() (*proto.SearchMovieResponse, usecase.SearchMovieUseCase) {
				movieUseCase := usecase_mock.NewMockSearchMovieUseCase(ctl)
				movies := &entity.MovieList{
					Search:       []*entity.Movie{
						{
							Title:  "batman",
						},
					},
					TotalResults: "1",
					Response:     "True",
				}
				movieUseCase.EXPECT().SearchMovie(gomock.Any(), gomock.Any()).Return(movies, nil).Times(1)
				result := &proto.SearchMovieResponse{
					Search:   []*proto.SearchMovieItemResponse{},
					Response: movies.Response,
				}
				for _, item := range movies.Search {
					result.Search = append(result.Search, &proto.SearchMovieItemResponse{
						Title:  item.Title,
					})
				}
				return result, movieUseCase
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res, movieUseCase := tt.mock()
			s := &searchMovieService{
				searchMovieUseCase: movieUseCase,
			}
			got, err := s.SearchMovie(tt.args.ctx, tt.args.request)
			if (err != nil) != tt.wantErr {
				t.Errorf("SearchMovie() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, res) {
				t.Errorf("SearchMovie() got = %v, want %v", got, res)
			}
		})
	}
}