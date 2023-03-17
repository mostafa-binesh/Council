package admin

import (
	D "docker/database"
	U "docker/utils"
	"github.com/gofiber/fiber/v2"
	"strconv"

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
		return U.ResErr(c, "در حذف کاربر مشکلی پیش آمده است")
	}
	return c.JSON(fiber.Map{
		"message": " کابر با موفقیت حذف شد",
	})
}

// ! create two routes for verification, one for verify, second for unverify
// ! there's no need to convert formValue to int
// ! for verifying, you need to get the user first, if it was verified already, return error
func UserVerification(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.FormValue("id"))
	if err != nil {
		return U.ResErr(c, err.Error())
	}
	var bool = true
	if c.FormValue("verify") == "0" {
		bool = false
	}
	result := D.DB().Model(&M.User{}).Where("id = ?", id).Update("Verified", bool)
	if result.Error != nil {
		U.DBError(c, result.Error)
	}
	return c.Status(400).JSON(fiber.Map{
		"message": " بروز رسانی با موفقیت انجام شد",
	})
}
