package routes

import (
	C "docker/controllers"
	U "docker/utils"
	"github.com/gofiber/fiber/v2"
)

func RouterInit() {
	router := fiber.New(fiber.Config{
		// Prefork:       true,
		ServerHeader: "Kurox",
		AppName:      "Authentication with Gorm Validation v0.1",
	})
	// routes
	router.Get("/", func(c *fiber.Ctx) error { return c.JSON(fiber.Map{"msg": "hello world"}) })
	// ! devs route
	dev := router.Group("/devs")
	dev.Get("/autoMigrate", C.AutoMigrate)
	// ! authentication routes
	router.Post("/signup", C.SignUpUser)
	router.Post("/login", C.Login)
	router.Get("/logout", C.Logout)
	// ! dashboard routes
	dashboard := router.Group("/dashboard", C.AuthMiddleware)
	dashboard.Get("/", C.Dashboard)
	APP_PORT, _ := U.Env("APP_PORT")
	router.Listen(":" + APP_PORT)
}
