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

// ############################
// ##########    USER   #############
// ############################
// api/admin/users/1

// ! Index User with admin/users route
func IndexUser(c *fiber.Ctx) error {
	user := []M.User{}
	pagination := new(F.Pagination)
	if err := c.QueryParser(pagination); err != nil {
		U.ResErr(c, err.Error())
	}
	D.DB().Where("id > ?", 0).Scopes(
		F.FilterByType(c,
			F.FilterType{QueryName: "name", Operator: "LIKE"},
			F.FilterType{QueryName: "nationalCode", ColumnName: "national_code"},
			F.FilterType{QueryName: "personalCode", ColumnName: "personal_code"}),
		F.Paginate(user, pagination)).Find(&user)
	pass_data := []M.MinUser{}
	for i := 0; i < len(user); i++ {
		pass_data = append(pass_data, M.MinUser{
			ID:           user[i].ID,
			Name:         user[i].Name,
			PhoneNumber:  user[i].PhoneNumber,
			PersonalCode: user[i].PersonalCode,
			NationalCode: user[i].NationalCode,
		})
	}
	return c.JSON(fiber.Map{
		"meta": pagination,
		"data": pass_data,
	})
}
func CheckPasswordHash(password string, hash string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err
}
func EditUser(c *fiber.Ctx) error {

	user := M.User{}
	payload := new(M.EditInput)
	if err := c.BodyParser(payload); err != nil {
		U.ResErr(c, err.Error())
	}
	if errs := U.Validate(payload); errs != nil {
		return c.Status(400).JSON(fiber.Map{"errors": errs})
	}
	result1 := D.DB().Where("id = ?", c.Params("id")).Find(&user)
	if result1.Error != nil {
		return U.DBError(c, result1.Error)
	}
	user.Name = payload.Name
	user.NationalCode = payload.NationalCode
	if payload.Password != "" {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(payload.Password), bcrypt.DefaultCost)
		if err != nil {
			return U.ResErr(c, "خطا در پردازش رمز عبور")
		}
		user.Password = string(hashedPassword)
	}
	user.PhoneNumber = payload.PhoneNumber
	user.PersonalCode = payload.PersonalCode
	result := D.DB().Save(&user)
	if result.Error != nil {
		return U.DBError(c, result.Error)
	}
	return c.JSON(fiber.Map{
		"message": "اطلاعات با موفقیت ادیت شد",
	})
}

// ! user by id with admin/users/{id}
func UserByID(c *fiber.Ctx) error {
	user := M.User{}
	D.DB().Where("id = ?", c.Params("id")).Find(&user)
	if user.Name == "" {
		return U.ResErr(c, "کاربر وجود ندارد")
	}
	minUser := M.MinUser{
		ID:           user.ID,
		Name:         user.Name,
		PhoneNumber:  user.PhoneNumber,
		PersonalCode: user.PersonalCode,
		NationalCode: user.NationalCode,
	}
	return c.JSON(fiber.Map{
		"data": minUser,
	})
}

// ! Delete user with admin/users/{}
func DeleteUser(c *fiber.Ctx) error {
	result := D.DB().Delete(&M.User{}, c.Params("id"))
	if result.Error != nil {
		return U.DBError(c, result.Error)
	}
	if result.RowsAffected == 0 {
		return U.ResErr(c, "کاربر یافت نشد")
	}
	return c.JSON(fiber.Map{
		"message": " کابر با موفقیت حذف شد",
	})
}

// ! create two routes for verification, one for verify, second for unverify
// ! there's no need to convert formValue to int
// ! for verifying, you need to get the user first, if it was verified already, return error
func UserVerification(c *fiber.Ctx) error {
	result := D.DB().Model(&M.User{}).Where("id = ?", c.Params("id")).Update("verified", true)
	if result.Error != nil {
		U.DBError(c, result.Error)
	}
	return c.Status(400).JSON(fiber.Map{
		"message": " بروز رسانی با موفقیت انجام شد",
	})
}
func UserUnVerification(c *fiber.Ctx) error {

	result := D.DB().Model(&M.User{}).Where("id = ?", c.Params("id")).Update("verified", false)
	if result.Error != nil {
		U.DBError(c, result.Error)
	}
	return c.Status(400).JSON(fiber.Map{
		"message": " بروز رسانی با موفقیت انجام شد",
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
		PhoneNumber:  strings.ToLower(payload.PhoneNumber),
		Password:     string(hashedPassword),
		PersonalCode: payload.PersonalCode,
		NationalCode: payload.NationalCode,
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
// user := []M.User{}
//
//	pagination := new(F.Pagination)
//	if err := c.QueryParser(pagination); err != nil {
//		U.ResErr(c, err.Error())
//	}
//	D.DB().Where("id > ?", 0).Scopes(F.Paginate(user, pagination)).Find(&user)
