package utils

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func DBError(c *fiber.Ctx, err error) error {
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return c.Status(400).JSON(fiber.Map{
			"error": "داده یافت نشد",
		})
	} else if errors.Is(err, gorm.ErrInvalidData) {
		return c.Status(400).JSON(fiber.Map{
			"error": "داده نامعتبر است",
		})
	} else if errors.Is(err, gorm.ErrDuplicatedKey) {
		return c.Status(400).JSON(fiber.Map{
			"error": "مقداری از داده ها، قبلا در پایگاه داده وجود دارد",
		})
	} else {
		return c.Status(400).JSON(fiber.Map{
			"error":     "خطای پیش بینی نشده ی پایگاه داده",
			// "errorText": err.Error(),
		})
	}
}
