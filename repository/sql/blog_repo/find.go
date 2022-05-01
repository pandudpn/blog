package blogrepo

import (
	"context"

	"github.com/pandudpn/blog/model"
)

func (b *blogRepository) FindActiveBlog(ctx context.Context, query string) ([]*model.Blog, error) {
	var (
		blogs = make([]*model.Blog, 0)
		err   error
	)
	q := `
	SELECT
		b.id, b.created_by, b.title, b.body, b.image, b.status, b.created_at, b.updated_at,
		u.id, u.name, u.email
	FROM
		"blog" AS b
	INNER JOIN
		"users" AS u
	ON
		u.id = b.created_by
	WHERE b.deleted_at IS NULL AND b.status = 1
	`

	query = q + query

	rows, err := b.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var (
			blog *model.Blog
			user *model.User
		)

		err = rows.Scan(
			&blog.Id, &blog.CreatedBy, &blog.Title,
			&blog.Body, &blog.Image, &blog.Status,
			&blog.CreatedAt, &blog.UpdatedAt, &user.Id,
			&user.Name, &user.Email,
		)
		if err != nil {
			return nil, err
		}

		blog.User = user
		blogs = append(blogs, blog)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return blogs, nil
}
