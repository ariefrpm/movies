package controller

import (
	"github.com/ariefrpm/movies/proto"
	"github.com/ariefrpm/movies/usecase"
	"golang.org/x/net/context"
)

type movieDetailService struct {
	movieDetailUseCase usecase.MovieDetailUseCase
}

func NewMovieDetailService(movieDetailUseCase usecase.MovieDetailUseCase) proto.MovieDetailServiceServer {
	return &movieDetailService{
		movieDetailUseCase: movieDetailUseCase,
	}
}

func (m *movieDetailService) MovieDetail(ctx context.Context, request *proto.MovieDetailRequest) (*proto.MovieDetailResponse, error) {
	movie, err := m.movieDetailUseCase.MovieDetail(request.OmdbID)
	if err != nil {
		return nil, err
	}
	result := &proto.MovieDetailResponse{
		Title:    movie.Title,
		Year:     movie.Year,
		Rated:    movie.Rated,
		Released: movie.Released,
		Runtime:  movie.Runtime,
		Genre:    movie.Genre,
		Director: movie.Director,
		Writer:   movie.Writer,
		Actors:   movie.Actors,
		Plot:     movie.Plot,
		Language: movie.Language,
		Country:  movie.Country,
		Awards:   movie.Awards,
		Poster:   movie.Poster,
		Response: movie.Response,
	}
	return result, nil
}
