package admin

import (
	D "docker/database"
	F "docker/database/filters"
	U "docker/utils"
	"strings"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"

	C "docker/controllers"
	// "strconv"

	// F "docker/database/filters"
	M "docker/models"
)

// TODO
// ! niazi be get ba'd delete nist, delete kon, erroresh ro ba U.DBError handle kon
// ! joda kardan verify va unverify
// ! eshtebah haye hejaii mesle code_persenel > personal_code
// ! update kardan moshkel dare, documentation update e gorm ro bekhun
// ! c.Query("id") shayad vojud nadashte bashe, bayad ba c.param handle she
// ############################
// ##########    USER   #############
// ############################
// api/admin/users/1

// ! Index User with admin/users route
func IndexUser(c *fiber.Ctx) error {
	user := []M.User{}

	D.DB().Find(&user)

	return c.Status(200).JSON(fiber.Map{
		"data": user,
	})
}

// ! user by id with admin/users/{id}
func UserByID(c *fiber.Ctx) error {
	user := M.User{}
	D.DB().Where("id = ?", c.Params("id")).Find(&user)
	if user.Name == "" {
		return U.ResErr(c, "کاربر وجود ندارد")
	}
	return c.JSON(fiber.Map{
		"data": user,
	})
}

// ! Delete user with admin/users/{}
func DeleteUser(c *fiber.Ctx) error {

	result := D.DB().Delete(&M.User{}, c.Params("id"))

	if result.Error != nil {
		return U.DBError(c, result.Error)
	}
	return c.JSON(fiber.Map{
		"message": " کابر با موفقیت حذف شد",
	})
}

// ! create two routes for verification, one for verify, second for unverify
// ! there's no need to convert formValue to int
// ! for verifying, you need to get the user first, if it was verified already, return error
func UserVerification(c *fiber.Ctx) error {

	result := D.DB().Model(&M.User{}).Where("id = ?", c.Params("id")).Update("Verified", true)
	if result.Error != nil {
		U.DBError(c, result.Error)
	}
	return c.Status(400).JSON(fiber.Map{
		"message": " بروز رسانی با موفقیت انجام شد",
	})
}
func UserUnVerification(c *fiber.Ctx) error {

	result := D.DB().Model(&M.User{}).Where("id = ?", c.Params("id")).Update("Verified", false)
	if result.Error != nil {
		U.DBError(c, result.Error)
	}
	return c.Status(400).JSON(fiber.Map{
		"message": " بروز رسانی با موفقیت انجام شد",
	})
}

func UserSearch(c *fiber.Ctx) error {
	user := []M.User{}
	D.DB().Scopes(
		F.FilterByType(c,
			F.FilterType{QueryName: "name", Operator: "LIKE"},
			F.FilterType{QueryName: "national_code", ColumnName: "national_code"},
			F.FilterType{QueryName: "personal_code", ColumnName: "personal_code"})).
		Find(&user)
	return c.JSON(fiber.Map{
		"data": user,
	})
}

func AddUser(c *fiber.Ctx) error {
	payload := new(M.SignUpInput)
	// ! parse body
	res, err := C.BodyParserHandle(c, payload)
	if err != nil {
		return res
	}
	// ! validate request
	// err = C.validate.Struct(payload)
	if err != nil {
		return C.ValidationHandle(c, err)
	}
	// hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(payload.Password), bcrypt.DefaultCost)
	if err != nil {
		U.ResErr(c, err.Error())
	}
	newUser := M.User{
		Name:         payload.Name,
		Email:        strings.ToLower(payload.Email),
		Password:     string(hashedPassword),
		PersonalCode: payload.PersonalCode,
		NationalCode: payload.NationalCode,
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

// ############################
// ##########    LAW   #############
// ############################
func UpdateLaw(c *fiber.Ctx) error {
	law := M.Law{}
	payload := new(M.CreateLawInput)
	// parsing the payload
	if err := c.BodyParser(payload); err != nil {
		U.ResErr(c, err.Error())
	}
	if errs := U.Validate(payload); errs != nil {
		return c.Status(400).JSON(fiber.Map{"errors": errs})
	}
	result1 := D.DB().Where("id = ?", c.Params("id")).Find(&law)
	if result1.Error != nil {
		return U.DBError(c, result1.Error)
	}
	law.Body = payload.Body
	law.Image = payload.Image
	law.NotificationDate = payload.NotificationDate
	law.NotificationNumber = payload.NotificationNumber
	law.SessionDate = payload.SessionDate
	law.SessionNumber = payload.SessionNumber
	law.Title = payload.Title
	law.Type = payload.Type
	result := D.DB().Save(&law)
	if result.Error != nil {
		return U.ResErr(c, "مشکلی در به روز رسانی به وجود آمده")
	}
	return c.JSON(fiber.Map{
		"message": "به روز رسانی با موفقیت انجام شد",
	})
}

func DeleteLaw(c *fiber.Ctx) error {
	result := D.DB().Delete(&M.Law{}, c.Params("id"))
	if result.Error != nil {
		return U.DBError(c, result.Error)
	}
	return c.JSON(fiber.Map{
		"message": "حذف کردن با موفقیت انجام شد",
	})
}
