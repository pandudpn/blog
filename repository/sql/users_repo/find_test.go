package usersrepo_test

import (
	"context"
	"database/sql"
	"testing"
	
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/pandudpn/blog/model"
	usersrepo "github.com/pandudpn/blog/repository/sql/users_repo"
	"github.com/stretchr/testify/assert"
)

func TestUsersRepo_FindUserByEmail(t *testing.T) {
	db, mock := NewMockSql()
	
	userRepo := usersrepo.NewUsersRepo(db)
	
	query := `
	SELECT id, name, email, password
	FROM "users"
	WHERE email=$1
	`
	
	rows := sqlmock.NewRows([]string{"id", "name", "email", "password"}).
		AddRow(users.Id, users.Name, users.Email, users.Password)
	
	mock.ExpectQuery(query).WithArgs(users.Email).WillReturnRows(rows)
	
	t.Run("Test Case #1 Find User By Email", func(t *testing.T) {
		expectedResult := &model.User{
			Id:       users.Id,
			Name:     users.Name,
			Email:    users.Email,
			Password: users.Password,
		}
		
		u, err := userRepo.FindUserByEmail(context.Background(), users.Email)
		
		assert.Equal(t, expectedResult, u)
		assert.Equal(t, nil, err)
	})
}

func TestUsersRepo_FindUserByEmailError(t *testing.T) {
	db, mock := NewMockSql()
	
	userRepo := usersrepo.NewUsersRepo(db)
	
	query := `
	SELECT id, name, email, password
	FROM "users"
	WHERE email=$1
	`
	
	rows := sqlmock.NewRows([]string{"id", "name", "email", "password"})
	
	mock.ExpectQuery(query).WithArgs(users.Email).WillReturnRows(rows)
	
	t.Run("Test Case #2 Error data no rows", func(t *testing.T) {
		var us *model.User
		u, err := userRepo.FindUserByEmail(context.Background(), users.Email)
		
		assert.Equal(t, us, u)
		assert.Equal(t, sql.ErrNoRows, err)
	})
}
