package controllers

import (
	D "docker/database"
	"strconv"

	"github.com/gofiber/fiber/v2"

	// F "docker/database/filters"
	M "docker/models"
)

func IndexUser(c *fiber.Ctx) error {
	user := []M.User{}

	D.DB().Find(&user)

	return c.Status(200).JSON(fiber.Map{
		"user": user,
	})
}

func DeleteUser(c *fiber.Ctx) error {
	result := D.DB().Delete(&M.User{}, c.FormValue("id"))
	if result.Error != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "در حذف کاربر مشکلی پیش آمده است",
		})
	}
	return c.JSON(fiber.Map{
		"message": " کابر با موفقیت حذف شد",
	})
}

func UserVerification(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.FormValue("id"))
	if err != nil {

	}
	var bool = true
	if c.FormValue("verify") == "0" {
		bool = false
	}
	result := D.DB().Model(&M.User{}).Where("id = ?", id).Update("Verified", bool)

	if result.Error != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "مشکلی در بروز رسانی پیش آمده",
		})
	}
	return c.Status(400).JSON(fiber.Map{
		"message": " بروز رسانی با موفقیت انجام شد",
	})
}
