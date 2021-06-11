package db

import "log"

//dummy db
type mysql struct {}

func NewMysqlConnection() DB {
	return &mysql{}
}

func (*mysql) Insert(query string) error {
	log.Printf("Inserting %s\n", query)
	return nil
}

