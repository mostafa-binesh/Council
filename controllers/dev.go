package controllers

import (
	U "docker/utils"

	"github.com/gofiber/fiber/v2"
)

func TranslationTest(c *fiber.Ctx) error {
	type User struct {
		Username string `validate:"required" json:"username"`
		Password string `validate:"required" json:"wirdpass"`
	}

	user := User{Username: "kurosh79"}
	if errs := U.Validate(user); errs != nil {
		return c.JSON(fiber.Map{"errors": U.Validate(user)})
	}
	return c.JSON(fiber.Map{"msg": "everything is fine"})
}
