package grpc

import (
	"fmt"
	service "github.com/ariefrpm/movies/controller/proto"
	"github.com/ariefrpm/movies/proto"
	"github.com/ariefrpm/movies/server"
	"github.com/ariefrpm/movies/usecase"
	"google.golang.org/grpc"
	"log"
	"net"
)

type grpcServer struct {
	searchMovieUseCase usecase.SearchMovieUseCase
	movieDetailUseCase usecase.MovieDetailUseCase
	port               int
	errCh              chan error
}

func NewGrpcServer(port int, searchMovieUseCase usecase.SearchMovieUseCase, movieDetailUseCase usecase.MovieDetailUseCase) server.Server {
	return &grpcServer{
		searchMovieUseCase: searchMovieUseCase,
		movieDetailUseCase: movieDetailUseCase,
		port:               port,
	}
}

func (g *grpcServer) Run() {
	log.Printf("start running grpc server on port %d\n", g.port)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", g.port))
	if err != nil {
		g.errCh <- err
	}
	grpcServer := grpc.NewServer()

	movieDetailService := service.NewMovieDetailService(g.movieDetailUseCase)
	searchMovieService := service.NewSearchMovieService(g.searchMovieUseCase)

	proto.RegisterMovieDetailServiceServer(grpcServer, movieDetailService)
	proto.RegisterSearchMovieServiceServer(grpcServer, searchMovieService)

	err = grpcServer.Serve(lis)
	if err != nil {
		g.errCh <- err
	}

}

func (g *grpcServer) ListenError() <-chan error {
	return g.errCh
}
