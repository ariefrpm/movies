package model

import "github.com/ariefrpm/movies/entity"

type SearchMovieResponse struct {
	Search   []SearchMovieItemResponse `json:"Search"`
	Response string                    `json:"Response"`
}

type SearchMovieItemResponse struct {
	Title  string `json:"Title"`
	Year   string `json:"Year"`
	ImdbID string `json:"ImdbID"`
	Type   string `json:"Type"`
	Poster string `json:"Poster"`
}

func (s *SearchMovieResponse) PopulateResponse(movies []*entity.Movie)  {
	s.Search = []SearchMovieItemResponse{}
	for _, m := range movies {
		s.Search = append(s.Search, SearchMovieItemResponse{
			Title:  m.Title,
			Year:   m.Year,
			ImdbID: m.ImdbID,
			Type:   m.Type,
			Poster: m.Poster,
		})
	}
}