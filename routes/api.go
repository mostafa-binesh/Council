package routes

import (
	C "docker/controllers"
	A "docker/controllers/admin"

	// U "docker/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"

	// "github.com/gofiber/fiber/v2/middleware/logger"
	// "github.com/gofiber/fiber/v2/middleware/csrf"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func RouterInit() {
	router := fiber.New(fiber.Config{
		// Prefork:       true,
		ServerHeader: "Kurox",
		AppName:      "Higher Education Council",
	})
	// ! add middleware
	router.Use(cors.New())
	// router.Use(logger.New())
	router.Use(recover.New())
	// router.Use(csrf.New()) ! setup csrf token on production
	// #######################
	// ########## ROUTES #############
	// #######################
	// ! routes
	router.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"lastJobDone": "New file structure",
		})
	})
	// ! Admin Route
	admin := router.Group("/admin")
	admin.Get("/users", A.IndexUser)
	admin.Get("userbyid", A.UserByID)
	admin.Get("/userSearch", A.UserSearch)
	admin.Post("/user", A.AddUser)
	admin.Post("/user/verify", A.UserVerification)
	admin.Delete("DeleteUser", A.DeleteUser)
	admin.Put("/laws", A.UpdateLaw)
	admin.Delete("/laws", A.DeleteLaw)
	// ! laws route
	laws := router.Group("/laws")
	laws.Get("/", C.AllLaws)
	laws.Post("/", C.CreateLaw)
	laws.Get("/search", C.LawSearch)
	laws.Get("/advancedLawSearch", C.AdvancedLawSearch) // ! only for test proposes
	laws.Get("/regulations", C.LawRegulations)
	laws.Get("/statutes", C.LawStatutes)
	laws.Get("/enactments", C.LawEnactments)
	laws.Get("/:id<int>", C.LawByID) // get certain law by id
	// ! devs route
	dev := router.Group("/devs")
	dev.Get("/autoMigrate", C.AutoMigrate)
	dev.Get("/translation", C.TranslationTest)
	dev.Get("/pagination", C.PaginationTest) // ?: send limit and page in the query
	dev.Get("/monitor", monitor.New())
	dev.Get("/panic", func(c *fiber.Ctx) error { panic("PANIC!") })
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
