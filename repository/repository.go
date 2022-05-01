package repository

import (
	"context"

	"github.com/pandudpn/blog/model"
)

type UserRepositoryInterface interface {
	CreateUser(ctx context.Context, user *model.User) error
	FindUserByEmail(ctx context.Context, email string) (*model.User, error)
}

type CacheRepositoryInterface interface {
	GetDataFromAccessToken(ctx context.Context, accessToken string) (int, error)
	SetLoginAccessToken(ctx context.Context, user *model.User) (model.ResponseLogin, error)
}

type BlogRepositoryInterface interface {
	FindActiveBlog(ctx context.Context, query string) ([]*model.Blog, error)
}
