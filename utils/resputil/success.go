package resputil

import (
	"github.com/gofiber/fiber/v2"
)

// Success returns formatted JSON response for success request
func Success(c *fiber.Ctx, statusCode int, systemCode string, data interface{}) error {
	var res = baseResponse{
		Status:     true,
		Code:       systemCode,
		Message:    "",
		StatusCode: statusCode,
		Data:       data,
	}
	
	return c.Status(res.StatusCode).JSON(res)
}
