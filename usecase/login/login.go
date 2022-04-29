package login

import (
	"context"
	"log"
	
	"github.com/pandudpn/blog/model"
	"github.com/pandudpn/blog/repository"
	"github.com/pandudpn/blog/utils/errsutil"
	"golang.org/x/crypto/bcrypt"
)

type loginUc struct {
	userRepo  repository.UserRepositoryInterface
	cacheRepo repository.CacheRepositoryInterface
}

// NewLoginUseCase is constructor of package Login UseCase (business logic)
// this is will return an instance of User Repository and Cache Repository
func NewLoginUseCase(ur repository.UserRepositoryInterface, cr repository.CacheRepositoryInterface) *loginUc {
	return &loginUc{
		userRepo:  ur,
		cacheRepo: cr,
	}
}

func (l *loginUc) Login(ctx context.Context, email, password string) (*model.ResponseLogin, error) {
	user, err := l.userRepo.FindUserByEmail(ctx, email)
	if err != nil {
		return nil, errsutil.Error(err, errsutil.UserNotFound, errsutil.NotFound)
	}
	
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, errsutil.Error(err, errsutil.PasswordNotMatch, errsutil.BadRequest)
	}
	
	res, err := l.cacheRepo.SetLoginAccessToken(ctx, user)
	if err != nil {
		log.Println(err)
		return nil, errsutil.Error(err, errsutil.InternalError, errsutil.InternalServer)
	}
	
	return &res, nil
}
