package handler

import (
	_ "embed"

	defaultError "github.com/cthiel77/server-demo/internal/error"
	"github.com/cthiel77/server-demo/server/handler/response"
	"github.com/gofiber/fiber/v2"
)

// Api404 the api 404 handdler
func Api404(c *fiber.Ctx) error {
	return c.Status(fiber.StatusNotFound).JSON(response.ErrorData{
		Code: defaultError.StatusNotFound,
		Msg:  "api page found",
	})
}
