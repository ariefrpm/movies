package grpc

import (
	"fmt"
	"github.com/ariefrpm/movies/proto"
	"github.com/ariefrpm/movies/server"
	"google.golang.org/grpc"
	"log"
	"net"
)

type grpcServer struct {
	searchMovieController proto.SearchMovieServiceServer
	movieDetailController proto.MovieDetailServiceServer
	port               int
	errCh              chan error
}

func NewGrpcServer(port int, searchMovieController proto.SearchMovieServiceServer, movieDetailController proto.MovieDetailServiceServer) server.Server {
	return &grpcServer{
		searchMovieController: searchMovieController,
		movieDetailController: movieDetailController,
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

	proto.RegisterMovieDetailServiceServer(grpcServer, g.movieDetailController)
	proto.RegisterSearchMovieServiceServer(grpcServer, g.searchMovieController)

	err = grpcServer.Serve(lis)
	if err != nil {
		g.errCh <- err
	}

}

func (g *grpcServer) ListenError() <-chan error {
	return g.errCh
}
