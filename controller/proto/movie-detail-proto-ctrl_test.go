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

func Test_movieDetailService_MovieDetail(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	type args struct {
		ctx     context.Context
		request *proto.MovieDetailRequest
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
		mock    func() (*proto.MovieDetailResponse, usecase.MovieDetailUseCase)
	}{
		{
			name: "error usecase",
			args: args{
				request: &proto.MovieDetailRequest{OmdbID: "1"},
			},
			wantErr: true,
			mock: func() (*proto.MovieDetailResponse, usecase.MovieDetailUseCase) {
				movieUseCase := usecase_mock.NewMockMovieDetailUseCase(ctl)
				movieUseCase.EXPECT().MovieDetail(gomock.Any()).Return(nil, errors.New("error")).Times(1)
				return nil, movieUseCase
			},
		},
		{
			name: "success",
			args: args{
				request: &proto.MovieDetailRequest{OmdbID: "1"},
			},
			wantErr: false,
			mock: func() (*proto.MovieDetailResponse, usecase.MovieDetailUseCase) {
				movieUseCase := usecase_mock.NewMockMovieDetailUseCase(ctl)
				movie := &entity.MovieDetail{
					Title: "Batman",
				}
				movieUseCase.EXPECT().MovieDetail(gomock.Any()).Return(movie, nil).Times(1)
				result := &proto.MovieDetailResponse{
					Title: movie.Title,
				}
				return result, movieUseCase
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res, movieUseCase := tt.mock()
			m := &movieDetailService{
				movieDetailUseCase: movieUseCase,
			}
			got, err := m.MovieDetail(tt.args.ctx, tt.args.request)
			if (err != nil) != tt.wantErr {
				t.Errorf("MovieDetail() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, res) {
				t.Errorf("MovieDetail() got = %v, want %v", got, res)
			}
		})
	}
}
