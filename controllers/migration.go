package controllers

import (
	D "docker/database"
	// S "docker/database/seeders"
	M "docker/models"
	"github.com/gofiber/fiber/v2"
)

// ! add any migration that you wanna add to the database
func AutoMigrate(c *fiber.Ctx) error {
	err := D.DB().AutoMigrate(&M.User{}, &M.Law{}, &M.Comment{}, &M.UserMigration{}, &M.Keyword{})
	if err != nil {
		return c.SendString(err.Error())
	}
	// ! seeders
	// S.UserSeeder()
	return c.SendString("migrate completed")
}
