package entity

type Movie struct {
	Title  string
	Year   string
	ImdbID string
	Type   string
	Poster string
}

type MovieList struct {
	Search       []*Movie
	TotalResults string
	Response     string
}
