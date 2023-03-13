package controllers
// ! this file has been created for test purposes
import (
	// "github.com/go-playground/locales/en"
	// ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	// en_translations "github.com/go-playground/validator/v10/translations/en"
	"github.com/gofiber/fiber/v2"
	// "fmt"
)

// User contains user information
type User struct {
	FirstName      string     `validate:"required"`
	LastName       string     `validate:"required"`
	Age            uint8      `validate:"gte=0,lte=130"`
	Email          string     `validate:"required,email"`
	FavouriteColor string     `validate:"iscolor"`                // alias for 'hexcolor|rgb|rgba|hsl|hsla'
	Addresses      []*Address `validate:"required,dive,required"` // a person can have a home and cottage...
}

// Address houses a users address information
type Address struct {
	Street string `validate:"required"`
	City   string `validate:"required"`
	Planet string `validate:"required"`
	Phone  string `validate:"required"`
}

// use a single instance , it caches struct info

func Translate(c *fiber.Ctx) error {
	type User struct {
		Username string `validate:"required"`
		Tagline  string `validate:"required,lt=10,numeric"`
		// Tagline2 string `validate:"required,gt=1,numeric"`
		Tagline2 string `validate:"required,min=2,max=4,numeric"`
	}
	type User2 struct {
		Name     string `validate:"required,gorm=unique.users.name" json:"name"`
		Username string `validate:"required" json:"username"`
		Tagline  string `validate:"required,lt=10,numeric" json:"simple tagline"`
		// Tagline2 string `validate:"required,gt=1,numeric"`
		Tagline2 string `validate:"required,min=2,max=4,numeric" json:"tagline 2"`
	}
	type Project struct {
		Title string `db:"title" json:"title" validate:"required,lte=25"` 
		// --> verify that the field exists and is less than or equal to 25 characters
	
		Description string `db:"description" json:"description" validate:"required"`
		// --> verify that the field exists
	
		WebsiteURL string `db:"website_url" json:"website_url" validate:"uri"`
		// --> verify that the field is correct URL string
	
		// Tags []string `db:"tags" json:"tags" validate:"len=3"`
		// --> verify that the field contains exactly three elements
	}
	// user := User{
	// 	Username: "Joeybloggs",
	// 	Tagline:  "This tagline is way too long.",
	// 	// Tagline2: "1",
	// 	Tagline2: "sadassdasdsaas",
	// }
	// user2 := User2{
	// 	// Name: "asdsad", // exists
	// 	Name: "bbvcb",
	// }
	project := Project {
		Title: "sadsa",
		// Description: "sadsadasasd",
		WebsiteURL: "sdsadsa",
		// Tags: [2]string{"saddsada","sadsadssad"},
	}
	// err := validate.Struct(user)
	// err := validate.Struct(user2)
	err := validate.Struct(project)
	if err != nil {

		// translate all error at once
		errs := err.(validator.ValidationErrors)

		// returns a map with key = namespace & value = translated error
		// NOTICE: 2 errors are returned and you'll see something surprising
		// translations are i18n aware!!!!
		// eg. '10 characters' vs '1 character'

		// fmt.Println(errs.Translate(trans))
		c.JSON(fiber.Map{
			"error":            err,
			"translatedErrors": errs.Translate(trans),
			// "ValidatorErrors": ValidatorErrors2(err),
			"ValidatorErrors2": ValidatorErrors(errs.Translate(trans)),
		})
	}
	return nil
}
