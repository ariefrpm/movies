package controller

import (
	"encoding/json"
	"errors"
	"github.com/ariefrpm/movies/model"
	"github.com/ariefrpm/movies/usecase"
	"net/http"
)

type MovieDetailController interface {
	MovieDetail(response http.ResponseWriter, request *http.Request)
}

type movieDetailController struct {
	movieDetailUseCase usecase.MovieDetailUseCase
}

func NewMovieDetailController(movieDetailUC usecase.MovieDetailUseCase) MovieDetailController {
	return &movieDetailController{
		movieDetailUseCase: movieDetailUC,
	}
}

func (m *movieDetailController) MovieDetail(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("Content-Type", "application/json")

	result := model.MovieDetailResponse{}
	omdbID, err := validateDetailRequest(request)
	if err != nil {
		result.Response = err.Error()
		json.NewEncoder(response).Encode(result)
		return
	}

	movie, err := m.movieDetailUseCase.MovieDetail(omdbID)
	if err != nil {
		result.Response = err.Error()
		json.NewEncoder(response).Encode(result)
		return
	}

	result.PopulateResponse(movie)
	json.NewEncoder(response).Encode(result)
}

func validateDetailRequest(request *http.Request) (string, error) {
	id := request.URL.Query().Get("i")

	if id == "" {
		return "", errors.New("imdbID is empty")
	}

	return id, nil
}