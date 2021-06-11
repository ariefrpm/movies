package usecase

import (
	"github.com/ariefrpm/movies/entity"
	"github.com/ariefrpm/movies/repository/api"
	"github.com/ariefrpm/movies/repository/db"
)

type SearchMovieUseCase interface {
	SearchMovie(page int, word string) (*entity.MovieList, error)
}

type searchMovieUseCase struct {
	searchMovieRepo api.MovieRepo
	loggingRepo db.Logging
}

func NewSearchMovieUseCase(movieRepo api.MovieRepo, logRepo db.Logging) SearchMovieUseCase  {
	return &searchMovieUseCase{
		searchMovieRepo: movieRepo,
		loggingRepo:     logRepo,
	}
}

func (s *searchMovieUseCase) SearchMovie(page int, word string) (*entity.MovieList, error) {
	movies, err := s.searchMovieRepo.SearchMovie(page, word)
	if err != nil {
		s.loggingRepo.Error("%s", err.Error())
	} else {
		s.loggingRepo.Info("Page: %d, Word %s, Total result: %s", page, word, movies.TotalResults)
	}
	return movies, err
}


