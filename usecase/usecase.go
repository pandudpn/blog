package usecase

import (
	"context"
	
	"github.com/pandudpn/blog/model"
)

type LoginUseCaseInterface interface {
	Login(ctx context.Context, email, password string) (*model.ResponseLogin, error)
}
