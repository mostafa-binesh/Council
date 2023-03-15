package utils

import "github.com/gofiber/fiber/v2"

func ResErr(c *fiber.Ctx, err string) error {
	return c.Status(400).JSON(fiber.Map{
		"error": err,
	})
}
