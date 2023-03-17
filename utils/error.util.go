package utils

import "github.com/gofiber/fiber/v2"
// response error, easier way to return a json error
func ResErr(c *fiber.Ctx, err string) error {
	return c.Status(400).JSON(fiber.Map{
		"error": err,
	})
}
