package routes

import (
	C "docker/controllers"
	AC "docker/controllers/admin"
	U "docker/utils"
	// U "docker/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	// "github.com/gofiber/fiber/v2/middleware/csrf"
	"github.com/gofiber/fiber/v2/middleware/encryptcookie"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func RouterInit() {
	router := fiber.New(fiber.Config{
		ServerHeader: "Kurox",
		AppName:      "Higher Education Council",
	})
	// ! add middleware
	router.Use(cors.New(cors.Config{
		AllowOrigins: U.Env("APP_ALLOW_ORIGINS"),
	}))
	router.Use(func(c *fiber.Ctx) error {
		U.BaseURL = c.BaseURL()
		return c.Next()
	})
	router.Use(logger.New())
	router.Use(recover.New())
	// #######################
	// ########## ROUTES #############
	// #######################
	router.Static("/public", "./public")
	router.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"msg":        "freeman was here :)",
			"lastChange": "add static files",
		})
	})
	// ! Admin Route
	admin := router.Group("/admin")
	admin.Get("/users", AC.IndexUser)
	admin.Get("/users/:id<int>", AC.UserByID)
	admin.Put("/users/:id<int>", AC.EditUser)
	admin.Post("/users", AC.AddUser)
	admin.Delete("/users/:id<int>", AC.DeleteUser)
	admin.Get("/laws", AC.IndexLaw)
	admin.Get("/laws/search", AC.LawSearch)
	admin.Get("laws/:id<int>", C.LawByID)
	admin.Post("/laws", AC.CreateLaw)
	admin.Put("/laws/:id<int>", AC.UpdateLaw)
	admin.Delete("/laws/:id<int>", AC.DeleteLaw)
	admin.Delete("/laws/:id<int>/files/:fileID<int>", AC.DeleteFile) // ! TODO : file az storage ham bayad paak she
	// ! laws route
	laws := router.Group("/laws")
	laws.Get("/", C.AllLaws)
	laws.Post("/", C.CreateLaw)
	laws.Get("/search", C.LawSearch)
	laws.Get("/regulations", C.LawRegulations)
	laws.Get("/statutes", C.LawStatutes)
	laws.Get("/enactments", C.LawEnactments)
	laws.Get("/:id<int>", C.LawByID)
	// ! authentication routes
	router.Post("/signup", C.SignUpUser)
	router.Post("/login", C.Login)
	router.Get("/logout", C.Logout)
	// ! messaging
	msg := router.Group("correspondence")
	msg.Use(encryptcookie.New(encryptcookie.Config{
		// ! only base64 charasters
		// ! A-Z | a-z | 0-9 | + | /
		Key: "S6e5+xc65+4dfs/nb4/f56+EW+56N4d6",
	}))
	// msg.Post("/register", C.GuestRegister)
	// msg.Get("/messages", C.GuestMessages)
	msg.Post("/messages", C.GuestSendMessage)
	msg.Get("/chats", C.GuestChats)
	msg.Post("/chats", C.CreateGuestChat)
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
	dev.Post("/upload", C.UploadFile)
	dev.Post("/fileExistenaceCheck", C.ExistenceCheck)
	dev.Post("/gormUnique", C.GormG)
	// ! listen
	router.Listen(":" + U.Env("APP_PORT"))
}
