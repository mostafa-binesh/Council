package routes

import (
	C "docker/controllers"
	AC "docker/controllers/admin"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/encryptcookie"
)

func APIInit(router *fiber.App) {
	router.Get("/migrate", C.AutoMigrate)
	router.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"msg": "freeman was here :)",
		})
	})
	router.Post("/signup", C.SignUpUser)
	router.Post("/login", C.Login)
	// ! laws route
	laws := router.Group("/laws")
	laws.Get("/", C.AllLaws)
	laws.Get("/search", C.LawSearch)
	laws.Get("/regulations", C.LawRegulations)
	laws.Get("/statutes", C.LawStatutes)
	laws.Get("/enactments", C.LawEnactments)
	laws.Get("/:id<int>", C.LawByID)
	laws.Get("/offline", C.OfflineLaws)
	laws.Put("/offilne/update", C.UpdateLawOffline)
	laws.Post("/comment", C.AddComment)

	// ! messaging
	msg := router.Group("correspondence")
	msg.Use(encryptcookie.New(encryptcookie.Config{
		// ! only base64 charasters
		// ! A-Z | a-z | 0-9 | + | /
		Key: "S6e5+xc65+4dfs/nb4/f56+EW+56N4d6",
	}))
	msg.Get("/chats", C.GuestChats)
	msg.Post("/chats", C.CreateGuestChat)
	msg.Post("/messages", C.GuestSendMessage)

	// authentication required endpoints
	router.Post("/login/token/refresh", C.RefreshToken)
	authRequired := router.Group("/", C.JWTAuthentication)

	authRequired.Get("/logout", C.Logout)
	authRequired.Get("/userInfo", C.UserInfo)

	// ! admin Route
	admin := authRequired.Group("/admin")
	admin.Get("/users", AC.IndexUser)
	admin.Get("/users/:id<int>", AC.UserByID)
	admin.Put("/users/:id<int>", AC.EditUser)
	admin.Post("/users", AC.AddUser)
	admin.Put("/users/:id<int>/verify", AC.UserVerification)
	admin.Put("/users/:id<int>/unverify", AC.UserUnVerification)
	admin.Delete("/users/:id<int>", AC.DeleteUser)
	admin.Get("/laws", AC.IndexLaw)
	admin.Get("/laws/search", AC.LawSearch)
	admin.Get("laws/:id<int>", AC.LawByID)
	admin.Post("/laws", AC.CreateLaw)
	admin.Put("/laws/:id<int>", AC.UpdateLaw)
	admin.Delete("/laws/:id<int>", AC.DeleteLaw)
	admin.Get("/laws/offline", C.OfflineLaws)
	admin.Delete("/laws/:id<int>/files/:fileID<int>", AC.DeleteFile) // ! TODO : file az storage ham bayad paak she
	admin.Get("/statics",AC.Statics)
	admin.Post("/uploadFile", AC.UploadFile) 
	admin.Put("/comment/:id<int>/verify",AC.VerifyComment)
	admin.Put("/comment/:id<int>/unverify",AC.UnVerifyComment)
	// ! dashboard routes
	dashboard := authRequired.Group("/dashboard", C.AuthMiddleware)
	dashboard.Get("/", C.Dashboard)

	// ! devs route
	dev := authRequired.Group("/devs")
	dev.Get("/autoMigrate", C.AutoMigrate)
	dev.Get("/changePhotoOfData", C.ChangePhotoofData)
	dev.Get("/translation", C.TranslationTest)
	dev.Get("/pagination", C.PaginationTest) // ?: send limit and page in the query
	dev.Get("/allUsers", C.DevAllUsers)      // ?: send limit and page in the query
	dev.Get("/panic", func(c *fiber.Ctx) error { panic("PANIC!") })
	laws.Get("/advancedLawSearch", C.AdvancedLawSearch)
	dev.Post("/upload", C.UploadFile)
	dev.Post("/fileExistenaceCheck", C.ExistenceCheck)
	dev.Post("/gormUnique", C.GormG)
	router.Get("/contextMemoryAddress", C.FiberContextMemoryAddress)
	devPanel := dev.Group("/admin")
	devPanel.Get("/structInfo", C.StructInfo)
}
