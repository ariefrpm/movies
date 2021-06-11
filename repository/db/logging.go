package db

import (
	"fmt"
	"github.com/ariefrpm/movies/library/db"
)

type Logging interface {
	Info(format string, a ...interface{})
	Error(format string, a ...interface{})
}

type dbLogging struct {
	db db.DB
}

func NewLogging() Logging  {
	return &dbLogging{
		db: db.NewMysqlConnection(),
	}
}

func (l *dbLogging) Info(format string, a ...interface{}) {
	//dummy logging
	_ = l.db.Insert(fmt.Sprintf(" INFO: %s\n", fmt.Sprintf(format, a...)))

}

func (l *dbLogging) Error(format string, a ...interface{}) {
	//dummy logging
	_ = l.db.Insert(fmt.Sprintf(" ERROR: %s\n", fmt.Sprintf(format, a ...)))
}
