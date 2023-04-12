package utils

import (
	"reflect"

	// "github.com/go-playground/locales/en"
	// D "docker/database"
	// "gorm.io/gorm"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/fa"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	fa_translations "github.com/go-playground/validator/v10/translations/fa"
	// en_translations "github.com/go-playground/validator/v10/translations/en"
)

var (
	uni      *ut.UniversalTranslator
	validate *validator.Validate
)
var faTranslation = map[string]string{
	"Name":               "نام",
	"FirstName":          "نام",
	"LastName":           "نام خانوادگی",
	"File":               "فایل",
	"Username":           "نام کاربری",
	"Password":           "رمز عبور",
	"Age":                "سن",
	"Type":               "نوع",
	"Title":              "عنوان",
	"SessionNumber":      "شماره جلسه",
	"SessionDate":        "تاریخ جلسه",
	"NotificationNumber": "شماره ابلاغ",
	"NotificationDate":   "تاریخ ابلاغ",
	"Body":               "بدنه",
	"Tags":               "تگ ها",
	"Image":              "عکس",
	"PersonalCode":       "کد پرسنلی",
	"NationalCode":       "کد ملی",
	"PhoneNumber":        "شماره همراه",
}

func Validate(fields interface{}) map[string]string {

	en := en.New()
	fa := fa.New()
	uni = ut.New(en, fa)

	// this is usually know or extracted from http 'Accept-Language' header
	// also see uni.FindTranslator(...)
	trans, _ := uni.GetTranslator("fa")
	validate = validator.New()
	fa_translations.RegisterDefaultTranslations(validate, trans)
	// ! custom names registration
	validate.RegisterTagNameFunc(func(field reflect.StructField) string {
		return faTranslation[field.Name]
	})
	// ! custom translations
	// ? required
	validate.RegisterTranslation("required", trans, func(ut ut.Translator) error {
		return ut.Add("required", "{0} الزامی است", true) // see universal-translator for details
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("required", fe.Field())
		return t
	})
	// ! possible issues: if fields have another struct in it, getJSONTag
	// ! > won't work properly
	err := validate.Struct(fields)
	if err != nil {
		responseError := make(map[string]string)
		errs := err.(validator.ValidationErrors)
		var jsonTag string
		for _, e := range errs {
			jsonTag = GetJSONTag(fields, e.StructField())
			if jsonTag == "" {
				jsonTag = ToLowerCamel(e.StructField())
			}
			responseError[jsonTag] = e.Translate(trans) // works fine
		}
		return responseError
	}
	return nil
}
