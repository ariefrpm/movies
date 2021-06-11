package db

type DB interface {
	Insert(query string) error
}