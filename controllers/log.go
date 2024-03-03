package controllers

import (
	"github.com/gofiber/fiber/v2"

	M "docker/models"
)

func AdminLogMiddleware(c *fiber.Ctx) error {
	c.Next()
	if !M.GetLog(c) {
		return c.JSON(fiber.Map{
			"error": "این درخواست مشکل دارد. لطفا لحظاتی بعد تلاش کنید",
		})
	}
}
