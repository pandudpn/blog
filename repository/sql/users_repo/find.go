package usersrepo

import (
	"context"
	
	"github.com/pandudpn/blog/model"
)

func (ur *usersRepo) FindUserByEmail(ctx context.Context, email string) (*model.User, error) {
	var user model.User
	
	query := `
	SELECT id, name, email, password
	FROM "users"
	WHERE email=$1
	`
	// get data from table `users`
	row := ur.db.QueryRowContext(ctx, query, email)
	
	// decode data into struct of User
	err := row.Scan(
		&user.Id,
		&user.Name,
		&user.Email,
		&user.Password,
	)
	// if query is not found in database, then set error
	if err != nil {
		return nil, err
	}
	
	// sent a single data from database
	return &user, nil
}
