package routes

import (
	C "docker/controllers"
	AC "docker/controllers/admin"

	// U "docker/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"

	// "github.com/gofiber/fiber/v2/middleware/logger"
	// "github.com/gofiber/fiber/v2/middleware/csrf"
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
	admin.Get("/users", AC.IndexUser)
	admin.Get("/users/:id<int>", AC.UserByID)
	admin.Put("/users/:id<int>",AC.EditUser)
	admin.Get("/users/search", AC.UserSearch)
	admin.Post("/users", AC.AddUser)
	admin.Delete("/users/:id<int>", AC.DeleteUser)
	admin.Get("/laws",AC.IndexLaw)
	admin.Get("/laws/search",AC.LawSearch)
	admin.Get("laws/:id<int>", C.LawByID) // get certain law by id
	admin.Post("/laws",AC.CreateLaw)
	admin.Put("/laws/:id<int>", AC.UpdateLaw)
	admin.Delete("/laws/:id<int>", AC.DeleteLaw)
	// ! laws route
	laws := router.Group("/laws")
	laws.Get("/", C.AllLaws)
	laws.Post("/", C.CreateLaw)
	laws.Get("/search", C.LawSearch)
	laws.Get("/regulations", C.LawRegulations)
	laws.Get("/statutes", C.LawStatutes)
	laws.Get("/enactments", C.LawEnactments)
	laws.Get("/:id<int>", C.LawByID) // get certain law by id
	// ! authentication routes
	router.Post("/signup", C.SignUpUser)
	router.Post("/login", C.Login)
	router.Get("/logout", C.Logout)
	// ! dashboard routes
	dashboard := router.Group("/dashboard", C.AuthMiddleware)
	dashboard.Get("/", C.Dashboard)
	// APP_PORT, _ := U.Env("APP_PORT")
	// ! devs route
	dev := router.Group("/devs")
	dev.Get("/autoMigrate", C.AutoMigrate)
	dev.Get("/translation", C.TranslationTest)
	dev.Get("/pagination", C.PaginationTest) // ?: send limit and page in the query
	dev.Get("/allUsers", C.DevAllUsers)      // ?: send limit and page in the query
	dev.Get("/panic", func(c *fiber.Ctx) error { panic("PANIC!") })
	laws.Get("/advancedLawSearch", C.AdvancedLawSearch)

	// ! listen
	router.Listen(":" + "8070")
}
