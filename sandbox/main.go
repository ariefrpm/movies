package main

import (
	"context"
	"github.com/ariefrpm/movies/proto"
	"google.golang.org/grpc"
	"log"
)

func main()  {
	conn, err := grpc.Dial(":9000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err.Error())
	}
	defer conn.Close()

	searchMovie := proto.NewSearchMovieServiceClient(conn)

	response, err := searchMovie.SearchMovie(
		context.Background(),
		&proto.SearchMovieRequest{Searchword:"Batman", Pagination:1},
	)
	if err != nil {
		log.Fatalf("error search movie: %s", err.Error())
	}
	log.Printf("response search movie: %v ", response.Search)

	detailMovie := proto.NewMovieDetailServiceClient(conn)
	detResponse, err := detailMovie.MovieDetail(context.Background(), &proto.MovieDetailRequest{OmdbID: "tt4116284"})
	if err != nil {
		log.Fatalf("error detail movie: %s", err.Error())
	}
	log.Printf("response detail movie: %v ", detResponse)
}
