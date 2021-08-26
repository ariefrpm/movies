package main

import (
	controller "github.com/ariefrpm/movies/controller/json"
	client "github.com/ariefrpm/movies/library/http-client"
	service "github.com/ariefrpm/movies/controller/proto"
	"github.com/ariefrpm/movies/repository/api"
	"github.com/ariefrpm/movies/repository/db"
	"github.com/ariefrpm/movies/server/grpc"
	"github.com/ariefrpm/movies/server/http"
	"github.com/ariefrpm/movies/usecase"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	movieApiRepo := api.NewMovieRepo(client.NewDefaultHttpClient())
	loggingRepo := db.NewLogging()

	var movieDetailUseCase usecase.MovieDetailUseCase
	movieDetailUseCase = usecase.NewMovieDetailUseCase(movieApiRepo, loggingRepo)
	searchMovieUseCase := usecase.NewSearchMovieUseCase(movieApiRepo, loggingRepo)

	movieDetailController := controller.NewMovieDetailController(movieDetailUseCase)
	searchMovieController := controller.NewSearchMovieController(searchMovieUseCase)
	movieDetailService := service.NewMovieDetailService(movieDetailUseCase)
	searchMovieService := service.NewSearchMovieService(searchMovieUseCase)

	httpServer := http.NewHttpServer(8080, searchMovieController, movieDetailController)
	grpcServer := grpc.NewGrpcServer(9000, searchMovieService, movieDetailService)

	go httpServer.Run()
	go grpcServer.Run()

	term := make(chan os.Signal)
	signal.Notify(term, os.Interrupt, syscall.SIGTERM)

	select {
	case o := <-term:
		log.Printf("exiting gracefully %s", o.String())
	case er := <-httpServer.ListenError():
		log.Printf("error starting http server, exiting gracefully %s", er.Error())
	case er := <-grpcServer.ListenError():
		log.Printf("error starting grpc server, exiting gracefully %s", er.Error())
	}

}


