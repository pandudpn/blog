// Package dbc is Database Connection
package dbc

import (
	"database/sql"
	"fmt"
	"log"
	"net/url"
	"os"
)

// NewConnectionSql to connect database
func NewConnectionSql() *sql.DB {
	// set host database
	u := url.URL{
		Host:   fmt.Sprintf("%s:%s", os.Getenv("DB_HOST"), os.Getenv("DB_PORT")),
		User:   url.UserPassword(os.Getenv("DB_USERNAME"), os.Getenv("DB_PASSWORD")),
		Path:   os.Getenv("DB_NAME"),
		Scheme: "postgres",
	}
	// make query for set raw query database
	dsn := u.Query()
	// set ssl_mode of database
	dsn.Set("ssl_mode", os.Getenv("DB_SSL_MODE"))
	// set timezone database
	dsn.Set("timezone", os.Getenv("DB_TIMEZONE"))
	
	conn, err := sql.Open("postgres", dsn.Encode())
	if err != nil {
		log.Fatalln("error open connection to database", err)
	}
	
	err = conn.Ping()
	if err != nil {
		log.Fatalln("ping database error", err)
	}
	
	return conn
}
