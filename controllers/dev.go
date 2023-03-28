package controllers

import (
	D "docker/database"
	F "docker/database/filters"
	M "docker/models"
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
func PaginationTest(c *fiber.Ctx) error {
	enactments := []M.Law{}
	pagination := new(F.Pagination)
	if err := c.QueryParser(pagination); err != nil {
		U.ResErr(c, err.Error())
	}
	D.DB().Where("type = ?", 3).Scopes(F.Paginate(enactments, pagination)).Find(&enactments)
	return c.JSON(fiber.Map{
		"meta": pagination,
		"data": enactments,
	})
}
