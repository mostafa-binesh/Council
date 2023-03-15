package routes

import (
	C "docker/controllers"

	// U "docker/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func RouterInit() {
	router := fiber.New(fiber.Config{
		// Prefork:       true,
		ServerHeader: "Kurox",
		AppName:      "Higher Education Council",
	})
	// ! add middleware
	router.Use(cors.New())
	// routes
	router.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"msg": "WELCOME",
		})
	})
	// ! laws route
	laws := router.Group("/laws")
	laws.Get("/", C.AllLaws)
	laws.Post("/search", C.LawSearch)
	laws.Get("/advancedLawSearch", C.AdvancedLawSearch)
	laws.Get("/regulations", C.LawRegulations)
	laws.Get("/statutes", C.LawStatutes)
	laws.Get("/enactments", C.LawEnactments)
	laws.Get("/:id", C.LawByID) // get certain law by id
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
	// APP_PORT, _ := U.Env("APP_PORT")
	router.Listen(":" + "8070")
}
