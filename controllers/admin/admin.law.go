package admin

import (
	D "docker/database"
	F "docker/database/filters"
	U "docker/utils"
	"strings"

	"github.com/gofiber/fiber/v2"
	// "strconv"

	// F "docker/database/filters"
	M "docker/models"
)

func IndexLaw(c *fiber.Ctx) error {
	laws := []M.Law{}
	pagination := new(F.Pagination)
	if err := c.QueryParser(pagination); err != nil {
		U.ResErr(c, err.Error())
	}
	D.DB().Where("id > ?", 0).Scopes(F.Paginate(laws, pagination)).Find(&laws)
	pass_data := []M.LawMinimal_min{}
	for i := 0; i < len(laws); i++ {
		pass_data = append(pass_data, M.LawMinimal_min{
			ID:    laws[i].ID,
			Title: laws[i].Title,
			Image: laws[i].Image,
		})
	}
	return c.JSON(fiber.Map{
		"meta": pagination,
		"data": pass_data,
	})
}
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
func CreateLaw(c *fiber.Ctx) error {
	payload := new(M.CreateLawInput)
	// parsing the payload
	if err := c.BodyParser(payload); err != nil {
		U.ResErr(c, err.Error())
	}
	if errs := U.Validate(payload); errs != nil {
		return c.Status(400).JSON(fiber.Map{"errors": errs})
	}
	law := M.Law{
		Type:               payload.Type,
		Title:              payload.Title,
		SessionNumber:      payload.SessionNumber,
		SessionDate:        payload.SessionDate,
		NotificationDate:   payload.NotificationDate,
		NotificationNumber: payload.NotificationNumber,
		Body:               payload.Body,
		Image:              payload.Image,
	}
	result := D.DB().Create(&law)
	if result.Error != nil {
		return U.ResErr(c, result.Error.Error())
	}
	var tags = strings.Split(payload.Tags, ",")
	for i := 0; i < len(tags); i++ {
		result2 := D.DB().Create(&M.Keyword{
			Keyword: tags[i],
			LawID:   law.ID,
		})
		if result2.Error != nil {
			D.DB().Delete(&M.Law{}, law.ID)
			return U.ResErr(c, result.Error.Error())
			// return U.ResErr(c, "خطایی در اضافه کردن تگ ها پیش آمده است.")
		}
	}
	return c.Status(200).JSON(fiber.Map{
		"message": "مصوبه با موفقیت اضافه شد",
	})
}
func LawSearch(c *fiber.Ctx) error {
	laws := []M.Law{}
	pagination := new(F.Pagination)
	D.DB().Scopes(
		F.FilterByType(c,
			F.FilterType{QueryName: "title", Operator: "LIKE"},
			F.FilterType{QueryName: "session_number", Operator: "LIKE"},
			F.FilterType{QueryName: "notification_number", Operator: "LIKE"},
			F.FilterType{QueryName: "body", Operator: "LIKE"},
			F.FilterType{QueryName: "recommender", Operator: "LIKE"},
			F.FilterType{QueryName: "type", Operator: "="},
			F.FilterType{QueryName: "notification_startDate", ColumnName: "notification_date", Operator: ">="},
			F.FilterType{QueryName: "notification_endDate", ColumnName: "notification_date", Operator: "<="},
			F.FilterType{QueryName: "session_startDate", ColumnName: "session_date", Operator: ">="},
			F.FilterType{QueryName: "session_endDate", ColumnName: "session_date", Operator: "<="}),
			F.Paginate(laws, pagination)).
		Find(&laws)
	pass_data := []M.LawMinimal_min{}
	for i := 0; i < len(laws); i++ {
		pass_data = append(pass_data, M.LawMinimal_min{
			ID:    laws[i].ID,
			Title: laws[i].Title,
			Image: laws[i].Image,
		})
	}
	return c.JSON(fiber.Map{
		"meta": pagination,
		"data": pass_data,
	})
}
func LawByID(c *fiber.Ctx) error {
	law := &M.Law{}
	if err := D.DB().Preload("Comments.User").First(law, c.Params("id")).Error; err != nil {
		return U.DBError(c, err)
	}
	return c.JSON(fiber.Map{
		"data": law,
	})
}
