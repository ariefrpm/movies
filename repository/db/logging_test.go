package db

import (
	db_mock "github.com/ariefrpm/movies/mock/library/db"
	"github.com/golang/mock/gomock"
	"testing"
)

func Test_dbLogging(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()
	dbMock := db_mock.NewMockDB(ctl)

	type args struct {
		format string
		a      []interface{}
	}
	tests := []struct {
		name   string
		args   args
	}{
		{
			name:   "Success",
			args:   args{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &dbLogging{
				db: dbMock,
			}
			dbMock.EXPECT().Insert(gomock.Any()).Times(2)
			l.Info(tt.name, tt.args)
			l.Error(tt.name, tt.args)
		})
	}
}