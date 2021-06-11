package controller

import (
	"encoding/json"
	"errors"
	"github.com/ariefrpm/movies/model"
	"github.com/ariefrpm/movies/usecase"
	"net/http"
	"strconv"
)

type SearchMovieController interface {
	SearchMovie(response http.ResponseWriter, request *http.Request)
}

type searchMovieController struct {
	searchMoveUseCase usecase.SearchMovieUseCase
}

func NewSearchMovieController(searchMoveUC usecase.SearchMovieUseCase) SearchMovieController {
	return &searchMovieController{
		searchMoveUseCase : searchMoveUC,
	}
}

func (s *searchMovieController) SearchMovie(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("Content-Type", "application/json")
	
	result := model.SearchMovieResponse{}
	page, word, err := validateRequest(request)
	if err != nil {
		result.Response = err.Error()
		json.NewEncoder(response).Encode(result)
		return
	}

	movies, err := s.searchMoveUseCase.SearchMovie(page, word)
	if err != nil {
		result.Response = err.Error()
		json.NewEncoder(response).Encode(result)
		return
	}

	result.PopulateResponse(movies.Search)
	result.Response = movies.Response
	json.NewEncoder(response).Encode(result)
}

func validateRequest(request *http.Request) (int, string, error) {
	page := request.URL.Query().Get("pagination")
	word := request.URL.Query().Get("searchword")

	if page == "" {
		return 0, "", errors.New("pagination is empty")
	}
	if word == "" {
		return 0, "", errors.New("word is empty")
	}

	pageInt, err := strconv.Atoi(page)
	if err != nil {
		return 0, "", errors.New("pagination is not a number")
	}
	
	return pageInt, word, nil
}