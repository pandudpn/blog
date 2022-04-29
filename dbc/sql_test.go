package dbc_test

import (
	"database/sql"
	"fmt"
	"os"
	"testing"
	
	"github.com/lib/pq"
	"github.com/pandudpn/blog/dbc"
	"github.com/stretchr/testify/assert"
)

func TestNewConnectionSql_Error1(t *testing.T) {
	t.Run("Test Case #1 Error Open Connection", func(t *testing.T) {
		var expectedResult *sql.DB
		expectedErr := fmt.Errorf("sql: unknown driver %q (forgotten import?)", os.Getenv("DB_DRIVER"))
		db, err := dbc.NewConnectionSql()
		
		assert.Equal(t, expectedResult, db)
		assert.Equal(t, expectedErr, err)
	})
}

func TestNewConnectionSql_Error2(t *testing.T) {
	t.Run("Test Case #2 Error Ping Database", func(t *testing.T) {
		os.Setenv("DB_DRIVER", "postgres")
		os.Setenv("DB_SSL_MODE", "")
		
		var expectedResult *sql.DB
		db, err := dbc.NewConnectionSql()
		
		assert.Equal(t, expectedResult, db)
		assert.Equal(t, pq.ErrSSLNotSupported, err)
	})
}

func TestNewConnectionSql_Success(t *testing.T) {
	t.Run("Test Case #3 Success open connection", func(t *testing.T) {
		os.Setenv("DB_HOST", "127.0.0.1")
		os.Setenv("DB_PORT", "5432")
		os.Setenv("DB_NAME", "blog")
		os.Setenv("DB_USERNAME", "pandu")
		os.Setenv("DB_PASSWORD", "123456")
		os.Setenv("DB_DRIVER", "postgres")
		os.Setenv("DB_SSL_MODE", "disable")
		os.Setenv("DB_TIMEZONE", "Asia/Jakarta")
		
		_, err := dbc.NewConnectionSql()
		
		assert.Equal(t, nil, err)
	})
}
