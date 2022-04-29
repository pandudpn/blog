// Package dbc is Database Connection
package dbc

import (
	"database/sql"
	"fmt"
	"net/url"
	"os"
	
	_ "github.com/lib/pq"
)

// NewConnectionSql to connect database
func NewConnectionSql() (*sql.DB, error) {
	// set host database
	dsn := url.URL{
		Host:   fmt.Sprintf("%s:%s", os.Getenv("DB_HOST"), os.Getenv("DB_PORT")),
		User:   url.UserPassword(os.Getenv("DB_USERNAME"), os.Getenv("DB_PASSWORD")),
		Path:   os.Getenv("DB_NAME"),
		Scheme: os.Getenv("DB_DRIVER"),
	}
	// make query for set raw query database
	query := dsn.Query()
	// set ssl_mode of database
	query.Add("sslmode", os.Getenv("DB_SSL_MODE"))
	// set timezone database
	query.Add("TimeZone", os.Getenv("DB_TIMEZONE"))
	
	dsn.RawQuery = query.Encode()
	
	conn, err := sql.Open(dsn.Scheme, dsn.String())
	if err != nil {
		return nil, err
	}
	
	err = conn.Ping()
	if err != nil {
		return nil, err
	}
	
	return conn, nil
}
