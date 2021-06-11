package controller

import (
	"github.com/ariefrpm/movies/proto"
	"github.com/ariefrpm/movies/usecase"
	"golang.org/x/net/context"
)

type searchMovieService struct {
	searchMovieUseCase usecase.SearchMovieUseCase
}

func NewSearchMovieService(searchMovieUseCase usecase.SearchMovieUseCase) proto.SearchMovieServiceServer {
	return &searchMovieService{
		searchMovieUseCase:searchMovieUseCase,
	}
}

func (s *searchMovieService) SearchMovie(ctx context.Context, request *proto.SearchMovieRequest) (*proto.SearchMovieResponse, error) {
	movies, err := s.searchMovieUseCase.SearchMovie(int(request.Pagination), request.Searchword)
	if err != nil {
		return nil, err
	}

	result := &proto.SearchMovieResponse{
		Search:   []*proto.SearchMovieItemResponse{},
		Response: movies.Response,
	}
	for _, item := range movies.Search {
		result.Search = append(result.Search, &proto.SearchMovieItemResponse{
			Title:  item.Title,
			Year:   item.Year,
			ImdbID: item.ImdbID,
			Type:   item.Type,
			Poster: item.Poster,
		})
	}
	return result, nil
}
