package admin

import (
	M "docker/models"

	"github.com/gofiber/fiber/v2"
)

func AdminLogMiddleware(c *fiber.Ctx) error {
	// !! todo: do the logging before the handler

	// do the before handler jobs
	err := c.Next()
	// after going through the handler, log the info
	if !M.GetLog(c) {
		return c.JSON(fiber.Map{
			"error": "این درخواست مشکل دارد. لطفا لحظاتی بعد تلاش کنید",
		})
	}
	// and return the response
	return err
}
