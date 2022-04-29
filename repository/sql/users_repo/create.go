package usersrepo

import (
	"context"
	"fmt"
	
	"github.com/pandudpn/blog/model"
)

func (ur *usersRepo) CreateUser(ctx context.Context, user *model.User) error {
	query := `
	INSERT INTO "users" (name, email, password, created_at)
	VALUES ($1, $2, $3, $4)
	`
	
	stmt, err := ur.db.PrepareContext(ctx, query)
	if err != nil {
		return err
	}
	// close for inactive statement query
	defer stmt.Close()
	// exec query before
	res, err := stmt.Exec(
		user.Name,
		user.Email,
		user.Password,
		user.CreatedAt,
	)
	// check if execution data has error
	if err != nil {
		return err
	}
	
	lastInsertId, err := res.LastInsertId()
	if err == nil {
		// set new id into data User
		user.Id = int(lastInsertId)
	}
	
	rowsAffected, _ := res.RowsAffected()
	
	if rowsAffected > 0 {
		return nil
	}
	
	err = fmt.Errorf("no one created")
	return err
}
