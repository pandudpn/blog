package resputil

import (
	"net/http"
	
	"github.com/gofiber/fiber/v2"
	"github.com/pandudpn/blog/utils/errsutil"
)

// Error returns formatted JSON response for failed request
func Error(c *fiber.Ctx, err error) error {
	var res = baseResponse{
		Status:     false,
		Code:       "99",
		Message:    "Internal server error",
		StatusCode: http.StatusInternalServerError,
	}
	
	if err, ok := err.(*errsutil.ErrorResponse); ok {
		res.Code = err.SystemCode()
		res.Message = err.Message()
		res.StatusCode = err.StatusCode()
		if err.Validation() != nil {
			if val, found := err.Validation().([]interface{}); found {
				if len(val) > 0 {
					res.Validation = err.Validation()
				}
			}
		}
	}
	
	return c.Status(res.StatusCode).JSON(res)
}
