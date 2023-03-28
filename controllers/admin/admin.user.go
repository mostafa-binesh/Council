package admin

import (
	D "docker/database"
	F "docker/database/filters"
	U "docker/utils"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
	"strings"

	C "docker/controllers"
	// "strconv"

	// F "docker/database/filters"
	M "docker/models"
)

////////////////////////////////////////////////////////////////////////////////////////////
//######user Admin#####//
///////////////////////////////////////////////////////////////////////////////////////////

// ! Index User with admin/users route
func IndexUser(c *fiber.Ctx) error {
	user := []M.User{}

	D.DB().Find(&user)

	return c.Status(200).JSON(fiber.Map{
		"data": user,
	})
}

func UserByID(c *fiber.Ctx) error {
	user := M.User{}
	D.DB().Where("id=?", c.Query("id")).Find(&user)
	if user.Name == "" {
		return U.ResErr(c, "کاربر وجود ندارد")
	}
	return c.JSON(fiber.Map{
		"data": user,
	})
}
func DeleteUser(c *fiber.Ctx) error {
	user := M.User{}
	D.DB().Where("id=?", c.Query("id")).Find(&user)
	if user.Name == "" {
		return U.ResErr(c, "کاربر وجود ندارد")
	}
	result := D.DB().Delete(&M.User{}, c.Query("id"))
	if result.Error != nil {
		return U.ResErr(c, "در حذف کاربر مشکلی پیش آمده است")
	}
	return c.JSON(fiber.Map{
		"message": " کابر با موفقیت حذف شد",
	})
}

// ! create two routes for verification, one for verify, second for unverify
// ! there's no need to convert formValue to int
// ! for verifying, you need to get the user first, if it was verified already, return error
func UserVerification(c *fiber.Ctx) error {

	var bool = true
	if c.FormValue("verify") == "0" {
		bool = false
	}
	result := D.DB().Model(&M.User{}).Where("id = ?", c.FormValue("id")).Update("Verified", bool)
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
			F.FilterType{QueryName: "national_code", ColumnName: "national_code", Operator: "="},
			F.FilterType{QueryName: "code_persenal", ColumnName: "code_persenal", Operator: "="})).
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
	// ! todo: create a shorter function for the validation, like payload.validate(), validate function can get a T template
	// err = C.validate.Struct(payload)
	if err != nil {
		return C.ValidationHandle(c, err)
	}
	// hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(payload.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": err.Error()})
	}
	newUser := M.User{
		Name:         payload.Name,
		Email:        strings.ToLower(payload.Email), // ! can use fiber toLower function that has better performance
		Password:     string(hashedPassword),
		CodePersonal: payload.CodePersonal,
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

// //////////////////////////////////////////////////////////////////////////////////////////
// ######Law Admin#####//
// /////////////////////////////////////////////////////////////////////////////////////////
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
	D.DB().Where("id=?", c.Query("id")).Find(&law)
	if law.Title == "" {
		return U.ResErr(c, "قانون مد نظر یافت نشد")
	}
	result := D.DB().Model(&law).Where("role = ?", "admin").Updates(M.Law{
		Body: payload.Body,
		Image: payload.Image,
		NotificationDate: payload.NotificationDate,
		NotificationNumber: payload.NotificationNumber,
		SessionNumber: payload.SessionNumber,
		SessionDate: payload.SessionDate,
		Title: payload.Title,
		Type: payload.Type,
	})
	if result.Error !=nil {
		return U.ResErr(c,"مشکلی در به روز رسانی به وجود آمده")
	}
	return c.JSON(fiber.Map{
		"message":"به روز رسانی با موفقیت انجام شد",
	})
}

func DeleteLaw(c *fiber.Ctx)error{
	law := M.Law{}
	D.DB().Where("id=?", c.Query("id")).Find(&law)
	if law.Title == "" {
		return U.ResErr(c, "قانون مد نظر یافت نشد")
	}
	result := D.DB().Delete(&M.Law{} , c.Query("id"))
	if result.Error !=nil {
		return U.ResErr(c,"مشکلی در به حذف کردن  به وجود آمده")
	}
	return c.JSON(fiber.Map{
		"message":"حذف کردن با موفقیت انجام شد",
	})	
}
