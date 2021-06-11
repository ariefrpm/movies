package api

import (
	"encoding/json"
	"errors"
	"github.com/ariefrpm/movies/entity"
	client "github.com/ariefrpm/movies/library/http-client"
	"github.com/ariefrpm/movies/mock/library/http-client"
	"github.com/golang/mock/gomock"
	"reflect"
	"testing"
)

func Test_movieRepo_SearchMovie(t *testing.T) {
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
		mock    func() (client.HttpClient, *entity.MovieList)
	}{
		{
			name: "http GET error",
			args: args{
				page: 1,
				word: "Batman",
			},
			wantErr: true,
			mock: func() (client.HttpClient, *entity.MovieList) {
				mockClient := http_client_mock.NewMockHttpClient(ctl)
				mockClient.EXPECT().GET(gomock.Any()).Return(nil, errors.New("http GET error")).Times(1)
				return mockClient, nil
			},
		},
		{
			name: "http GET error unmarshal",
			args: args{
				page: 1,
				word: "Batman",
			},
			wantErr: true,
			mock: func() (client.HttpClient, *entity.MovieList) {
				mockClient := http_client_mock.NewMockHttpClient(ctl)
				mockClient.EXPECT().GET(gomock.Any()).Return([]byte{1}, nil).Times(1)
				return mockClient, nil
			},
		},
		{
			name: "http GET response false",
			args: args{
				page: 1,
				word: "Batman",
			},
			wantErr: true,
			mock: func() (client.HttpClient, *entity.MovieList) {
				movie := &entity.MovieList{
					Response: "false",
				}
				bytes, _ := json.Marshal(movie)
				mockClient := http_client_mock.NewMockHttpClient(ctl)
				mockClient.EXPECT().GET(gomock.Any()).Return(bytes, nil).Times(1)
				return mockClient, nil
			},
		},
		{
			name: "Success",
			args: args{
				page: 1,
				word: "Batman",
			},
			wantErr: false,
			mock: func() (client.HttpClient, *entity.MovieList) {
				movie := entity.MovieList{
					Search: []*entity.Movie{{
						Title:  "title",
						Year:   "year",
						ImdbID: "id",
						Type:   "type",
						Poster: "poster",
					}},
					TotalResults: "1",
					Response:     "True",
				}
				bytes, _ := json.Marshal(movie)
				mockClient := http_client_mock.NewMockHttpClient(ctl)
				mockClient.EXPECT().GET(gomock.Any()).Return(bytes, nil).Times(1)
				return mockClient, &movie
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockClient, movies := tt.mock()
			mo := &movieRepo{client: mockClient}
			got, err := mo.SearchMovie(tt.args.page, tt.args.word)
			if (err != nil) != tt.wantErr {
				t.Errorf("SearchMovie() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, movies) {
				t.Errorf("SearchMovie() got = %v, want %v", got, movies)
			}
		})
	}
}

func Test_movieRepo_MovieDetail(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	type args struct {
		imdbID string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
		mock    func() (client.HttpClient, *entity.MovieDetail)
	}{
		{
			name:    "http GET client errors",
			args:    args{},
			wantErr: true,
			mock: func() (client.HttpClient, *entity.MovieDetail) {
				mockClient := http_client_mock.NewMockHttpClient(ctl)
				mockClient.EXPECT().GET(gomock.Any()).Return(nil, errors.New("error")).Times(1)

				return mockClient, nil
			},
		},
		{
			name:    "http GET error unmarshal",
			args:    args{},
			wantErr: true,
			mock: func() (client.HttpClient, *entity.MovieDetail) {
				mockClient := http_client_mock.NewMockHttpClient(ctl)
				mockClient.EXPECT().GET(gomock.Any()).Return([]byte{1}, nil).Times(1)

				return mockClient, nil
			},
		},
		{
			name:    "http GET error response false",
			args:    args{},
			wantErr: true,
			mock: func() (client.HttpClient, *entity.MovieDetail) {
				movie := &entity.MovieDetail{
					Response: "false",
				}
				bytes, _ := json.Marshal(movie)
				mockClient := http_client_mock.NewMockHttpClient(ctl)
				mockClient.EXPECT().GET(gomock.Any()).Return(bytes, nil).Times(1)

				return mockClient, nil
			},
		},
		{
			name:    "Success",
			args:    args{},
			wantErr: false,
			mock: func() (client.HttpClient, *entity.MovieDetail) {
				movie := &entity.MovieDetail{
					Title:    "Title",
					Response: "True",
				}
				bytes, _ := json.Marshal(movie)
				mockClient := http_client_mock.NewMockHttpClient(ctl)
				mockClient.EXPECT().GET(gomock.Any()).Return(bytes, nil).Times(1)

				return mockClient, movie
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockClient, movie := tt.mock()

			m := &movieRepo{
				client: mockClient,
			}
			got, err := m.MovieDetail(tt.args.imdbID)
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
