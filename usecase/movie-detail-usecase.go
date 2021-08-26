package usecase

import (
	"github.com/ariefrpm/movies/entity"
	"github.com/ariefrpm/movies/repository/api"
	"github.com/ariefrpm/movies/repository/db"
)

type MovieDetailHandler func(MovieDetailUseCase) MovieDetailUseCase

type MovieDetailUseCase interface {
	MovieDetail(omdbID string) (*entity.MovieDetail, error)
}

type movieDetailUseCase struct {
	movieDetailRepo api.MovieRepo
	loggingRepo db.Logging
}

func NewMovieDetailUseCase(movieRepo api.MovieRepo, loggingRepo db.Logging) MovieDetailUseCase  {
	return &movieDetailUseCase{
		movieDetailRepo: movieRepo,
		loggingRepo: loggingRepo,
	}
}

func (m *movieDetailUseCase) MovieDetail(omdbID string) (*entity.MovieDetail, error) {
	movie, err := m.movieDetailRepo.MovieDetail(omdbID)
	if err != nil {
		m.loggingRepo.Error("%s", err.Error())
	} else {
		m.loggingRepo.Info("omdbID: %s, Movie title: %s", omdbID, movie.Title)
	}
	return movie, err
}


