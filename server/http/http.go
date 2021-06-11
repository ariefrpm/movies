package http

import (
	"fmt"
	ctrl "github.com/ariefrpm/movies/controller/json"
	"github.com/ariefrpm/movies/library/router"
	"github.com/ariefrpm/movies/server"
	"github.com/ariefrpm/movies/usecase"
	"log"
	"net/http"
)

type httpServer struct {
	searchMovieUseCase usecase.SearchMovieUseCase
	movieDetailUseCase usecase.MovieDetailUseCase
	port               int
	errCh              chan error
}

func NewHttpServer(port int, searchMovieUseCase usecase.SearchMovieUseCase, movieDetailUseCase usecase.MovieDetailUseCase) server.Server {
	return &httpServer{
		searchMovieUseCase: searchMovieUseCase,
		movieDetailUseCase: movieDetailUseCase,
		port:               port,
	}
}

func (h *httpServer) Run() {
	log.Printf("start running http server on port %d\n", h.port)

	route := router.NewDefaultRouter()

	route.GET("/api/search_movie", ctrl.NewSearchMovieController(h.searchMovieUseCase).SearchMovie)
	route.GET("/api/movie_detail", ctrl.NewMovieDetailController(h.movieDetailUseCase).MovieDetail)

	err := http.ListenAndServe(fmt.Sprintf(":%d", h.port), route.Handler())

	if err != nil {
		h.errCh <- err
	}
}

func (h *httpServer) ListenError() <-chan error {
	return h.errCh
}
