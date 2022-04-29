package usersrepo_test

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"testing"
	
	"github.com/DATA-DOG/go-sqlmock"
	usersrepo "github.com/pandudpn/blog/repository/sql/users_repo"
	"github.com/stretchr/testify/assert"
)

func TestUsersRepo_CreateUser(t *testing.T) {
	db, mock := NewMockSql()
	
	userRepo := usersrepo.NewUsersRepo(db)
	
	query := `
	INSERT INTO "users" (name, email, password, created_at)
	VALUES ($1, $2, $3, $4)
	`
	prepare := mock.ExpectPrepare(query)
	prepare.
		ExpectExec().
		WithArgs(users.Name, users.Email, users.Password, users.CreatedAt).
		WillReturnResult(sqlmock.NewResult(2, 1))
	
	t.Run("Test Case #1 Success Create", func(t *testing.T) {
		err := userRepo.CreateUser(context.Background(), users)
		
		if err == nil {
			t.Logf("new user_id %d", users.Id)
		}
		
		assert.Equal(t, nil, err)
	})
}

func TestUsersRepo_CreateUserError(t *testing.T) {
	db, mock := NewMockSql()
	
	userRepo := usersrepo.NewUsersRepo(db)
	
	query := `
	INSERT INTO "users" (name, email, password, created_at)
	VALUES ($1, $2, $3, $4)
	`
	prepare := mock.ExpectPrepare(query)
	prepare.
		ExpectExec().
		WithArgs(users.Name, users.Email, users.Password, users.CreatedAt).
		WillReturnResult(sqlmock.NewResult(0, 0))
	
	t.Run("Test Case #2 No Rows Affected", func(t *testing.T) {
		expectedError := errors.New("no one created")
		err := userRepo.CreateUser(context.Background(), users)
		
		assert.Equal(t, expectedError, err)
	})
}

func TestUsersRepo_CreateUserError2(t *testing.T) {
	db, mock := NewMockSql()
	
	userRepo := usersrepo.NewUsersRepo(db)
	
	query := `
	INSERT INTO "user" (name, email, password, created_at)
	VALUES ($1, $2, $3, $4)
	`
	prepare := mock.ExpectPrepare(query)
	prepare.
		ExpectExec().
		WithArgs(users.Name, users.Email, users.Password, users.CreatedAt).
		WillReturnResult(sqlmock.NewResult(0, 0))
	
	t.Run("Test Case #3 Error Prepare Statement", func(t *testing.T) {
		expectedError := fmt.Errorf("Prepare: actual sql: %s", query)
		err := userRepo.CreateUser(context.Background(), users)
		
		assert.Error(t, expectedError, err)
	})
}

func TestUsersRepo_CreateUserError3(t *testing.T) {
	db, mock := NewMockSql()
	
	userRepo := usersrepo.NewUsersRepo(db)
	
	query := `
	INSERT INTO "users" (name, email, password, created_at)
	VALUES ($1, $2, $3, $4)
	`
	prepare := mock.ExpectPrepare(query)
	prepare.
		ExpectExec().
		WillReturnError(sql.ErrConnDone)
	
	t.Run("Test Case #4 Error Execute", func(t *testing.T) {
		expectedError := sql.ErrConnDone
		err := userRepo.CreateUser(context.Background(), users)
		
		assert.Equal(t, expectedError, err)
	})
}
