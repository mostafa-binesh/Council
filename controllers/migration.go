package controllers

import (
	D "docker/database"
	S "docker/database/seeders"
	M "docker/models"
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

// ! add any migration that you wanna add to the database
func AutoMigrate(c *fiber.Ctx) error {
	// ! drop all tables if 'dropAllTables' field is 1 in the query
	if c.Query("dropAllTables") == "1" {
		fmt.Println("dropping all tables")
		D.DB().Migrator().DropTable(&M.User{}, &M.Law{}, &M.Comment{}, &M.Keyword{}, &M.Comment{})
	}
	// ! migrate tables
	err := D.DB().AutoMigrate(&M.User{}, &M.Law{}, &M.Comment{}, &M.Keyword{})
	if err != nil {
		return c.Status(400).SendString(err.Error())
	}
	var seederRepeatCount int64
	seederRepeatCount = 1
	seedCountQuery := c.Query("seederRepeatCount")
	if seedCountQuery != "" {
		var err error
		seederRepeatCount, err = strconv.ParseInt(seedCountQuery, 10, 64)
		if err != nil {
			panic("seedCount query param. cannot be parsed")
		}
	}
	fmt.Println("Tables migration done...")
	// ! seeders
	fmt.Printf("seeder gonna run for %d loop", seederRepeatCount)
	for i := 0; i < int(seederRepeatCount); i++ {
		S.InitSeeder()
	}
	return c.SendString("migrate completed")
}
