package blogrepo

import "database/sql"

type blogRepository struct {
	db *sql.DB
}

// NewBlogRepository is constructor of package Blog Repository
// and will return an instance of Blog Repository
// with value of Database Connection
func NewBlogRepository(db *sql.DB) *blogRepository {
	return &blogRepository{
		db: db,
	}
}
