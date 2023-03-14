package controllers

import (
	D "docker/database"
	S "docker/database/seeders"
	M "docker/models"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

// ! add any migration that you wanna add to the database
func AutoMigrate(c *fiber.Ctx) error {
	// ! drop all tables first
	D.DB().Migrator().DropTable(&M.User{}, &M.Law{}, &M.Comment{}, &M.Keyword{}, &M.Comment{})
	// ! migrate tables
	err := D.DB().AutoMigrate(&M.User{}, &M.Law{}, &M.Comment{}, &M.Keyword{})
	if err != nil {
		return c.Status(400).SendString(err.Error())
	}
	fmt.Println("Tables migration done...")
	// ! seeders
	S.InitSeeder()
	return c.SendString("migrate completed")
}
