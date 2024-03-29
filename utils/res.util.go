package utils

import (
	// F "docker/database/filters"
	"github.com/gofiber/fiber/v2"
)

// response error, easier way to return a json error
// and returns { error: err}
func ResErr(c *fiber.Ctx, err string) error {
	return c.Status(400).JSON(fiber.Map{
		"error": err,
	})
}
func ResValidationErr(c *fiber.Ctx, err map[string]string) error {
	return c.Status(400).JSON(fiber.Map{
		"errors": err,
	})
}
func ResWithPagination(c *fiber.Ctx, data interface{}, pagination Pagination) error {
	return c.Status(200).JSON(fiber.Map{
		"meta": pagination,
		"data": data,
	})
}

func ResMessage(c *fiber.Ctx, msg string) error {
	return c.Status(200).JSON(fiber.Map{
		"msg": msg,
	})
}
func BodyParserErr(c *fiber.Ctx) error {
	return c.Status(400).JSON(fiber.Map{
		"error": "خطای تجزیه ی درخواست",
	})
}
func CreateNewLogError(c *fiber.Ctx) error {
	return ResErr(c, "این درخواست مشکل دارد. لطفا لحظاتی بعد تلاش کنید")
}
