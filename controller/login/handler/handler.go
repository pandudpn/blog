package handler

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/pandudpn/blog/usecase"
)

type loginHandler struct {
	validator *validator.Validate
	loginUc   usecase.LoginUseCaseInterface
}

// NewHandlerLogin is constructor for Package Handler/Controller Login
// and will return an instance of Login UseCase and Validator struct
func NewHandlerLogin(validator *validator.Validate, l usecase.LoginUseCaseInterface) *loginHandler {
	return &loginHandler{
		validator: validator,
		loginUc:   l,
	}
}

func (l *loginHandler) Route(app *fiber.App) {
	login := app.Group("/v1/login")
	login.Post("", l.login)
}
