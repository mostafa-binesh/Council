package controllers

import (
	D "docker/database"
	F "docker/database/filters"
	M "docker/models"
	U "docker/utils"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func AllLaws(c *fiber.Ctx) error {
	laws := []M.Law{}
	regulations := []M.LawMinimal{}
	statutes := []M.LawStatutesMinimal{}
	enactments := []M.LawMinimal{}
	D.DB().Find(&laws)
	// ! filtering
	for i := 0; i < len(laws); i++ {
		if laws[i].Type == 1 {
			regulations = append(regulations, M.LawMinimal{
				ID:               laws[i].ID,
				Title:            laws[i].Title,
				Image:            laws[i].Image,
				NotificationDate: laws[i].NotificationDate,
			})
		}
	}
	for i := 0; i < len(laws); i++ {
		if laws[i].Type == 2 {
			statutes = append(statutes, M.LawStatutesMinimal{
				ID:               laws[i].ID,
				Title:            laws[i].Title,
				Image:            laws[i].Image,
				SessionNumber:    laws[i].SessionNumber,
				NotificationDate: laws[i].NotificationDate,
			})
		}
	}
	for i := 0; i < len(laws); i++ {
		if laws[i].Type == 3 {
			enactments = append(enactments, M.LawMinimal{
				ID:               laws[i].ID,
				Title:            laws[i].Title,
				Image:            laws[i].Image,
				NotificationDate: laws[i].NotificationDate,
			})
		}
	}
	return c.JSON(fiber.Map{
		"data": fiber.Map{
			"regulations": regulations,
			"statutes":    statutes,
			"enactments":  enactments,
		},
	})
}
func LawEnactments(c *fiber.Ctx) error {
	enactments := []M.Law{}
	D.DB().Where("type = ?", 3).Find(&enactments)
	return c.JSON(fiber.Map{
		"data": enactments,
	})
}
func LawStatutes(c *fiber.Ctx) error {
	statutes := []M.Law{}
	D.DB().Where("type = ?", 2).Find(&statutes)
	return c.JSON(fiber.Map{
		"data": statutes,
	})
}
func LawRegulations(c *fiber.Ctx) error {
	regulations := []M.Law{}
	D.DB().Where("type = ?", 1).Find(&regulations)
	return c.JSON(fiber.Map{
		"data": regulations,
	})
}
func AdvancedLawSearch(c *fiber.Ctx) error {
	laws := []M.Law{}
	D.DB().Scopes(
		F.FilterByType(c,
			F.FilterType{QueryName: "title", Operator: "LIKE"},
			F.FilterType{QueryName: "startDate", ColumnName: "notification_date", Operator: ">="})).
		Find(&laws)
	return c.JSON(fiber.Map{"data": laws})
}

func LawSearch(c *fiber.Ctx) error {
	laws := []M.Law{}
	D.DB().Scopes(
		F.FilterByType(c,
			F.FilterType{QueryName: "title", Operator: "LIKE"},
			F.FilterType{QueryName: "startDate", ColumnName: "notification_date", Operator: ">="},
			F.FilterType{QueryName: "endDate", ColumnName: "notification_date", Operator: "<="})).
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
		"data": pass_data,
	})
}
func LawByID(c *fiber.Ctx) error {
	law := &M.Law{}
	if err := D.DB().Preload("Comments.User").First(law, c.Params("id")).Error; err != nil {
		return U.DBError(c, err)
	}
	lawWithComment := M.LawWithMinimalComment{}
	lawWithComment.ID = law.ID
	lawWithComment.Type = law.Type
	lawWithComment.Title = law.Title
	lawWithComment.SessionNumber = law.SessionNumber
	lawWithComment.SessionDate = law.SessionDate
	lawWithComment.NotificationDate = law.NotificationDate
	lawWithComment.NotificationNumber = law.NotificationNumber
	lawWithComment.Body = law.Body
	lawWithComment.Image = law.Image
	lawWithComment.Comments = M.GetMinimalComment(law.Comments)
	lawWithComment.CreatedAt = law.CreatedAt
	lawWithComment.UpdatedAt = law.UpdatedAt
	return c.JSON(fiber.Map{
		"data": lawWithComment,
	})
}

func CreateLaw(c *fiber.Ctx) error {
	payload := new(M.CreateLawInput)
	// parsing the payload
	if err := c.BodyParser(payload); err != nil {
		U.ResErr(c, err.Error())
	}
	if errs := U.Validate(payload); errs != nil {
		return c.JSON(fiber.Map{"errors": errs})
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
			D.DB().Delete(&M.User{}, law.ID)
			return U.ResErr(c, result.Error.Error())
			// return U.ResErr(c, "خطایی در اضافه کردن تگ ها پیش آمده است.")
		}
	}
	return c.Status(200).JSON(fiber.Map{
		"message": "مصوبه با موفقیت اضافه شد",
	})
}
