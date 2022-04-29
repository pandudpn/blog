package usersrepo

import "database/sql"

type usersRepo struct {
	db *sql.DB
}

// NewUsersRepo is constructor of Package User Repository
// and will return an instance of Users Repository with some method
// to be implemented
func NewUsersRepo(db *sql.DB) *usersRepo {
	return &usersRepo{
		db: db,
	}
}
