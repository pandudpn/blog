package handler

import (
	"net/http"
	
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/pandudpn/blog/utils/errsutil"
	"github.com/pandudpn/blog/utils/resputil"
)

type loginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

func (l *loginHandler) login(c *fiber.Ctx) error {
	var req loginRequest
	
	if err := c.BodyParser(&req); err != nil {
		return resputil.Error(c, errsutil.Error(err, errsutil.BodyRequired, errsutil.BadRequest))
	}
	
	err := l.validator.Struct(req)
	if err != nil {
		var errorValidation = make([]map[string]interface{}, 0)
		for _, rr := range err.(validator.ValidationErrors) {
			var rs = map[string]interface{}{
				"field":   rr.Field(),
				"message": rr.Error(),
			}
			
			errorValidation = append(errorValidation, rs)
		}
		
		return resputil.Error(c, errsutil.Error(err, errsutil.BodyRequired, errsutil.BadRequest, errorValidation))
	}
	
	res, err := l.loginUc.Login(c.Context(), req.Email, req.Password)
	if err != nil {
		return resputil.Error(c, err)
	}
	
	return resputil.Success(c, http.StatusOK, "10", res)
}
