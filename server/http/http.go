package http

import (
	"fmt"
	"github.com/ariefrpm/movies/library/router"
	"github.com/ariefrpm/movies/server"
	"github.com/ariefrpm/movies/controller/json"
	"log"
	"net/http"
)

type httpServer struct {
	searchMovieController controller.SearchMovieController
	movieDetailController controller.MovieDetailController
	port               int
	errCh              chan error
}

func NewHttpServer(port int, searchMovieController controller.SearchMovieController, movieDetailController controller.MovieDetailController) server.Server {
	return &httpServer{
		searchMovieController: searchMovieController,
		movieDetailController: movieDetailController,
		port:               port,
	}
}

func (h *httpServer) Run() {
	log.Printf("start running http server on port %d\n", h.port)

	route := router.NewDefaultRouter()

	route.GET("/api/search_movie", h.searchMovieController.SearchMovie)
	route.GET("/api/movie_detail", h.movieDetailController.MovieDetail)

	err := http.ListenAndServe(fmt.Sprintf(":%d", h.port), route.Handler())

	if err != nil {
		h.errCh <- err
	}
}

func (h *httpServer) ListenError() <-chan error {
	return h.errCh
}
