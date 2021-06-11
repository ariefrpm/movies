package usecase

import (
	"errors"
	"github.com/ariefrpm/movies/entity"
	api_mock "github.com/ariefrpm/movies/mock/repository/api"
	db_mock "github.com/ariefrpm/movies/mock/repository/db"
	"github.com/ariefrpm/movies/repository/api"
	"github.com/ariefrpm/movies/repository/db"
	"github.com/golang/mock/gomock"
	"reflect"
	"testing"
)

func Test_movieDetailUseCase_MovieDetail(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	type args struct {
		omdbID string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
		mock    func() (api.MovieRepo, db.Logging, *entity.MovieDetail)
	}{
		{
			name:    "error",
			args:    args{},
			wantErr: true,
			mock: func() (api.MovieRepo, db.Logging, *entity.MovieDetail) {
				mockApiRepo := api_mock.NewMockMovieRepo(ctl)
				mockDBRepo := db_mock.NewMockLogging(ctl)
				mockApiRepo.EXPECT().MovieDetail(gomock.Any()).Return(nil, errors.New("error")).Times(1)
				mockDBRepo.EXPECT().Error(gomock.Any(), gomock.Any()).Times(1)
				return mockApiRepo, mockDBRepo, nil
			},
		},
		{
			name:    "success",
			args:    args{},
			wantErr: false,
			mock: func() (api.MovieRepo, db.Logging, *entity.MovieDetail) {
				movie := &entity.MovieDetail{
					Title:    "title",
					Response: "True",
				}
				mockApiRepo := api_mock.NewMockMovieRepo(ctl)
				mockDBRepo := db_mock.NewMockLogging(ctl)
				mockApiRepo.EXPECT().MovieDetail(gomock.Any()).Return(movie, nil).Times(1)
				mockDBRepo.EXPECT().Info(gomock.Any(), gomock.Any(), gomock.Any()).Times(1)
				return mockApiRepo, mockDBRepo, movie
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			movieRepo, loggingRepo, movie := tt.mock()
			m := &movieDetailUseCase{
				movieDetailRepo: movieRepo,
				loggingRepo:     loggingRepo,
			}
			got, err := m.MovieDetail(tt.args.omdbID)
			if (err != nil) != tt.wantErr {
				t.Errorf("MovieDetail() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, movie) {
				t.Errorf("MovieDetail() got = %v, want %v", got, movie)
			}
		})
	}
}
