package utils

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func DBError(c *fiber.Ctx, err error) error {
	var errorText string
	if errors.Is(err, gorm.ErrRecordNotFound) {
		errorText = "داده یافت نشد"
	} else if errors.Is(err, gorm.ErrInvalidData) {
		errorText = "داده نامعتبر است"
	} else if errors.Is(err, gorm.ErrDuplicatedKey) {
		errorText = "مقداری از داده ها، قبلا در پایگاه داده وجود دارد"
	} else {
		errorText = "خطای پیش بینی نشده ی پایگاه داده"
	}
	if Env("APP_DEBUG") == "true" {
		return c.Status(400).JSON(fiber.Map{
			"error": errorText,
			"debug": err,
		})
	}
	return c.Status(400).JSON(fiber.Map{
		"error": errorText,
	})
}
