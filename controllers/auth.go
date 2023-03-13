package controllers

import (
	"fmt"
	// "github.com/go-playground/validator/v10"
	D "docker/database"
	M "docker/models"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"golang.org/x/crypto/bcrypt"
	"strings"
)

var (
	Store    *session.Store
	AUTH_KEY string = "authenticated"
	USER_ID  string = "user_id"
)

type Person struct {
	Name string `json:"name" xml:"name" form:"name"`
	Pass string `json:"pass" xml:"pass" form:"pass"`
}

func SignUpUser(c *fiber.Ctx) error {
	payload := new(M.SignUpInput)
	// ! parse body
	res, err := BodyParserHandle(c, payload)
	if err != nil {
		return res
	}
	// ! validate request
	// ! todo: create a shorter function for the validation, like payload.validate(), validate function can get a T template
	err = validate.Struct(payload)
	if err != nil {
		return ValidationHandle(c, err)
	}
	// hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(payload.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": err.Error()})
	}
	newUser := M.User{
		Name:     payload.Name,
		Email:    strings.ToLower(payload.Email), // ! can use fiber toLower function that has better performance
		Password: string(hashedPassword),
		// Photo:    &payload.Photo, // ? don't know why add & in the payload for photo
	}
	// ! add user to the database
	result := D.DB().Create(&newUser)
	// ! if any error exist in the create process, write the error
	if result.Error != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "couldn't create the user"})
	}
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "user has been created successfully"})

}

func Login(c *fiber.Ctx) error {
	payload := new(M.SignInInput)
	// ! parse body
	res, err := BodyParserHandle(c, payload)
	if err != nil {
		fmt.Println("inside error")
		return res
	}
	fmt.Println("body parser pass")
	// ! validate request
	err = validate.Struct(payload)
	if err != nil {
		return ValidationHandle(c, err)
	}
	var user M.User
	result := D.DB().First(&user, "email = ?", strings.ToLower(payload.Email))
	// ! result.error will not be null if no row returned, so i commented it
	// ! i guess the way that i handled it is not the best and i should create a method and check if the error is
	// ! no row found, ignore the error
	// ! maybe can handle it in this way
	// ! https://gorm.io/docs/error_handling.html
	if result.RowsAffected == 0 {
		return ReturnError(c, "Invalid email or password")
	}
	// ! compare the password of payload and returned user from database
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(payload.Password))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Invalid email or Password"})
	}
	sess := GetSess(c)
	sess.Set(USER_ID, user.ID)
	if err := sess.Save(); err != nil {
		return ReturnError(c, "server error", 500)
	}
	return c.SendString("you can now access /dashboard")
}
func Logout(c *fiber.Ctx) error {
	// ! just remove the session
	sess, err := Store.Get(c)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "not authenticated",
		})
	}
	if err := sess.Destroy(); err != nil {
		panic(err)
	}
	return c.SendString("logged out successfully")
}
func Dashboard(c *fiber.Ctx) error {
	// ! has a AuthMiddleware before here
	// ! if session and user exists, client can access here
	user := c.Locals("user").(M.User)
	return c.JSON(fiber.Map{"dashboard": "heres the dashboard", "user": user})
}
func AuthMiddleware(c *fiber.Ctx) error {
	sess := GetSess(c)
	userID := sess.Get(USER_ID)
	if userID == nil {
		return ReturnError(c, "not authenticated", fiber.StatusUnauthorized) // ! notAuthorized is notAuthenticated
	} else {
		c.SendString(fmt.Sprintf("user id is: %s", sess.Get(USER_ID)))
	}
	var user M.User
	result := D.DB().Find(&user, userID)
	if result.Error != nil {
		err := sess.Destroy()
		if err != nil {
			panic(err)
		}
		return ReturnError(c, "cannot authenticate. session removed", 500)
	}
	c.Locals("user", user)
	return c.Next()

}
