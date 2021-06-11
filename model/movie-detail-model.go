package model

import "github.com/ariefrpm/movies/entity"

type MovieDetailResponse struct {
	Title    string	`json:"title"`
	Year     string	`json:"year"`
	Rated    string	`json:"rated"`
	Released string	`json:"released"`
	Runtime  string	`json:"runtime"`
	Genre    string	`json:"genre"`
	Director string	`json:"director"`
	Writer   string	`json:"writer"`
	Actors   string	`json:"actors"`
	Plot     string	`json:"plot"`
	Language string	`json:"language"`
	Country  string	`json:"country"`
	Awards   string	`json:"awards"`
	Poster   string	`json:"poster"`
	Response string	`json:"response"`
}

func (m *MovieDetailResponse) PopulateResponse(detail *entity.MovieDetail) {
	m.Title = detail.Title
	m.Year = detail.Year
	m.Rated = detail.Rated
	m.Released = detail.Released
	m.Runtime = detail.Runtime
	m.Genre = detail.Genre
	m.Director = detail.Director
	m.Writer = detail.Writer
	m.Actors = detail.Actors
	m.Plot = detail.Plot
	m.Language = detail.Language
	m.Country = detail.Country
	m.Awards = detail.Awards
	m.Poster = detail.Poster
	m.Response = detail.Response
}

