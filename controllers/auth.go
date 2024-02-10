package controllers

import (
	"fmt"
	// "github.com/go-playground/validator/v10"
	D "docker/database"
	M "docker/models"
	U "docker/utils"
	"github.com/gofiber/fiber/v2"
	// "github.com/golang-jwt/jwt"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

type Person struct {
	Name string `json:"name" xml:"name" form:"name"`
	Pass string `json:"pass" xml:"pass" form:"pass"`
}

func SignUpUser(c *fiber.Ctx) error {
	payload := new(M.SignUpInput)
	// payload := new(M.SignInInput)
	// ! parse payload
	if err := c.BodyParser(payload); err != nil {
		U.ResErr(c, err.Error())
	}
	// ! validate request
	if errs := U.Validate(payload); errs != nil {
		return U.ResValidationErr(c, errs)
	}
	// hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(payload.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": err.Error()})
	}
	newUser := M.User{
		Name:         payload.Name,
		PhoneNumber:  strings.ToLower(payload.PhoneNumber), // ! can use fiber toLower function that has better performance
		Password:     string(hashedPassword),
		PersonalCode: payload.PersonalCode,
		NationalCode: payload.NationalCode,
		RoleID:       3,
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
	// ! parse payload
	if err := c.BodyParser(payload); err != nil {
		U.ResErr(c, err.Error())
	}
	// ! validate request
	if errs := U.Validate(payload); errs != nil {
		return U.ResValidationErr(c, errs)
	}
	var user M.User
	result := D.DB().First(&user, "personal_code = ?", strings.ToLower(payload.PersonalCode))
	// ! the reason we didn't handle the error first,
	// ! - is because not found return error option is disabled
	if result.RowsAffected == 0 {
		// return ReturnError(c, "ایمیل یا رمز عبور اشتباه است")
		return U.ResErr(c, "کد پرسنلی یا رمز عبور اشتباه است")
	}
	if result.Error != nil {
		return U.DBError(c, result.Error)
	}
	// ! compare the password of payload and returned user from database
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(payload.Password))
	if err != nil {
		return U.ResErr(c, "کد پرسنلی یا رمز عبور اشتباه است")
	}
	if !user.Verified {
		return c.JSON(fiber.Map{
			"Error": "این فرد هنوز تایید نشده است",
		})
	}
	sess := U.Session(c)
	sess.Set(U.USER_ID, user.ID)
	if err := sess.Save(); err != nil {
		return U.ResErr(c, "خطا در ورود")
	}
	token, err := createToken(user)
	if err != nil {
		// در صورت بروز خطا در ایجاد توکن، پاسخ مناسب را به مشتری ارسال کنید
		return U.ResErr(c, "خطا در ایجاد توکن")
	}
	return c.JSON(fiber.Map{
		"Name":         user.Name,
		"PersonalCode": user.PersonalCode,
		"PhoneNumber":  user.PhoneNumber,
		"NationalCode": user.NationalCode,
		"Permissions":  getPermissions(user),
		"token": token,
	})

}
func Logout(c *fiber.Ctx) error {
	// ! just removing the session
	sess, err := U.Store.Get(c)
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
	sess := U.Session(c)
	userID := sess.Get(U.USER_ID)
	if userID == nil {
		return ReturnError(c, "not authenticated", fiber.StatusUnauthorized) // ! notAuthorized is notAuthenticated
	} else {
		c.SendString(fmt.Sprintf("user id is: %s", sess.Get(U.USER_ID)))
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

// Assuming jwtSecretKey and jwtRefreshSecretKey are defined and hold the secret keys for signing JWT tokens

// TokenDetails struct and CreateToken function should be defined as per your existing code

// Authenticate is a middleware for validating access tokens in the Authorization header
func JWTAuthentication(c *fiber.Ctx) error {
	authHeader := c.Get("Authorization")
	if authHeader == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "No authorization token provided"})
	}

	// Extract the token from the Authorization header
	splitToken := strings.Split(authHeader, "Bearer ")
	if len(splitToken) != 2 {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid authorization token format"})
	}
	tokenString := splitToken[1]

	// Parse the token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Ensure the token's algorithm matches "HS256"
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fiber.NewError(fiber.StatusUnauthorized, "Unexpected signing method in auth token")
		}
		return U.JWTSecretKey, nil
	})

	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid authorization token"})
	}

	// Validate the token
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// You can add additional checks on the claims here if needed

		// For example, extracting and setting the user ID in the Fiber context
		userID := claims["user_id"]
		c.Locals("userID", userID)

		// Proceed to the next middleware/handler
		return c.Next()
	}

	return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid authorization token"})
}
