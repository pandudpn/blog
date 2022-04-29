package usersrepo_test

import (
	"database/sql"
	"time"
	
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/pandudpn/blog/model"
)

var users = &model.User{
	Id:        1,
	Name:      "pandu dwi putra nugroho",
	Email:     "pandu@unittest.com",
	Password:  "abc",
	CreatedAt: time.Now(),
	UpdatedAt: sql.NullTime{
		Time:  time.Now(),
		Valid: true,
	},
}

func NewMockSql() (*sql.DB, sqlmock.Sqlmock) {
	db, mock, _ := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	
	return db, mock
}
