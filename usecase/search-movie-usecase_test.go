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

func Test_searchMovieUseCase_SearchMovie(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	type args struct {
		page int
		word string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
		mock    func() (api.MovieRepo, db.Logging, *entity.MovieList)
	}{
		{
			name:    "error",
			args:    args{},
			wantErr: true,
			mock: func() (api.MovieRepo, db.Logging, *entity.MovieList) {
				mockApiRepo := api_mock.NewMockMovieRepo(ctl)
				mockDBRepo := db_mock.NewMockLogging(ctl)
				mockApiRepo.EXPECT().SearchMovie(gomock.Any(), gomock.Any()).Return(nil, errors.New("error")).Times(1)
				mockDBRepo.EXPECT().Error(gomock.Any(), gomock.Any()).Times(1)

				return mockApiRepo, mockDBRepo, nil
			},
		},
		{
			name:    "success",
			args:    args{},
			wantErr: false,
			mock: func() (api.MovieRepo, db.Logging, *entity.MovieList) {
				movies := &entity.MovieList{
					Response: "True",
				}
				mockApiRepo := api_mock.NewMockMovieRepo(ctl)
				mockDBRepo := db_mock.NewMockLogging(ctl)
				mockApiRepo.EXPECT().SearchMovie(gomock.Any(), gomock.Any()).Return(movies, nil).Times(1)
				mockDBRepo.EXPECT().Info(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Times(1)

				return mockApiRepo, mockDBRepo, movies
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			movieRepo, loggingRepo, movie := tt.mock()

			s := &searchMovieUseCase{
				searchMovieRepo: movieRepo,
				loggingRepo:     loggingRepo,
			}
			got, err := s.SearchMovie(tt.args.page, tt.args.word)
			if (err != nil) != tt.wantErr {
				t.Errorf("SearchMovie() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, movie) {
				t.Errorf("SearchMovie() got = %v, want %v", got, movie)
			}
		})
	}
}
