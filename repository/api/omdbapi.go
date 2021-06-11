package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/ariefrpm/movies/entity"
	client "github.com/ariefrpm/movies/library/http-client"
)

type MovieRepo interface {
	SearchMovie(page int, word string) (*entity.MovieList, error)
	MovieDetail(imdbID string) (*entity.MovieDetail, error)
}

type movieRepo struct {
	client client.HttpClient
}

func NewMovieRepo(client client.HttpClient) MovieRepo {
	return &movieRepo{
		client: client,
	}
}

const (
	BaseUrl = "http://www.omdbapi.com"
	OmdbKey = "faf7e5bb"
)

func (m *movieRepo) SearchMovie(page int, word string) (*entity.MovieList, error) {
	url := fmt.Sprintf("%s/?apikey=%s&s=%s&page=%d", BaseUrl, OmdbKey, word, page)

	data, err := m.client.GET(url)
	if err != nil {
		return nil, err
	}

	var res entity.MovieList
	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}

	if res.Response != "True" {
		return nil, errors.New(fmt.Sprintf("response error: %s", res.Response))
	}

	return &res, nil
}

func (m *movieRepo) MovieDetail(imdbID string) (*entity.MovieDetail, error) {
	url := fmt.Sprintf("%s/?apikey=%s&i=%s", BaseUrl, OmdbKey, imdbID)

	data, err := m.client.GET(url)
	if err != nil {
		return nil, err
	}

	var res entity.MovieDetail
	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}

	if res.Response != "True" {
		return nil, errors.New(fmt.Sprintf("response error: %s", res.Response))
	}

	return &res, nil
}
