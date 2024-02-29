package controllers

import (
	D "docker/database"
	F "docker/database/filters"
	M "docker/models"
	U "docker/utils"

	// "strings"

	"github.com/gofiber/fiber/v2"
)

func AllLaws(c *fiber.Ctx) error {
	regulations := []M.LawMinimal{}
	statutes := []M.LawStatutesMinimal{}
	enactments := []M.LawMinimal{}
	pagination := U.ParsedPagination(c)

	lawsType1 := []M.Law{}
	D.DB().Where("type = ?", 1).Scopes(U.Paginate(lawsType1, pagination)).Find(&lawsType1).Order("created_at asc")
	// ! filtering
	for i := 0; i < len(lawsType1); i++ {
		if lawsType1[i].Type == 1 {
			regulations = append(regulations, M.LawMinimal{
				ID:               lawsType1[i].ID,
				Title:            lawsType1[i].Title,
				Image:            U.BaseURL + "/public/uploads/" + lawsType1[i].Image,
				NotificationDate: lawsType1[i].NotificationDate,
			})
		}
	}
	lawsType2 := []M.Law{}
	D.DB().Where("type = ?", 2).Scopes(U.Paginate(lawsType2, pagination)).Find(&lawsType2).Order("created_at asc")
	for i := 0; i < len(lawsType2); i++ {
		if lawsType2[i].Type == 2 {
			statutes = append(statutes, M.LawStatutesMinimal{
				ID:               lawsType2[i].ID,
				Title:            lawsType2[i].Title,
				Image:            U.BaseURL + "/public/uploads/" + lawsType2[i].Image,
				SessionNumber:    lawsType2[i].SessionNumber,
				NotificationDate: lawsType2[i].NotificationDate,
			})
		}
	}
	lawsType3 := []M.Law{}
	D.DB().Where("type = ?", 3).Scopes(U.Paginate(lawsType3, pagination)).Find(&lawsType3).Order("created_at asc")
	for i := 0; i < len(lawsType3); i++ {
		if lawsType3[i].Type == 3 {
			enactments = append(enactments, M.LawMinimal{
				ID:               lawsType3[i].ID,
				Title:            lawsType3[i].Title,
				Image:            U.BaseURL + "/public/uploads/" + lawsType3[i].Image,
				NotificationDate: lawsType3[i].NotificationDate,
			})
		}
	}
	return c.JSON(fiber.Map{
		"data": fiber.Map{
			"regulations": regulations,
			"statutes":    statutes,
			"enactments":  enactments,
		},
		"meta": pagination,
	})
}
func LawEnactments(c *fiber.Ctx) error {
	enactments := []M.Law{}
	pagination := U.ParsedPagination(c)
	D.DB().Scopes(U.Paginate(&M.Law{}, pagination)).Where("type = ?", 3).Find(&enactments)
	enactmentsArray := []M.LawMinimal{}
	for i := 0; i < len(enactments); i++ {
		if enactments[i].Type == 3 {
			enactmentsArray = append(enactmentsArray, M.LawMinimal{
				ID:               enactments[i].ID,
				Title:            enactments[i].Title,
				Image:            U.BaseURL + "/public/uploads/" + enactments[i].Image,
				NotificationDate: enactments[i].NotificationDate,
			})
		}
	}
	return c.JSON(fiber.Map{
		"data": enactmentsArray,
		"meta": pagination,
	})
}
func LawStatutes(c *fiber.Ctx) error {
	statutes := []M.Law{}
	pagination := U.ParsedPagination(c)
	D.DB().Scopes(U.Paginate(&M.Law{}, pagination)).Where("type = ?", 2).Find(&statutes)
	statutesArray := []M.LawMinimal{}
	for i := 0; i < len(statutes); i++ {
		if statutes[i].Type == 2 {
			statutesArray = append(statutesArray, M.LawMinimal{
				ID:               statutes[i].ID,
				Title:            statutes[i].Title,
				Image:            U.BaseURL + "/public/uploads/" + statutes[i].Image,
				NotificationDate: statutes[i].NotificationDate,
			})
		}
	}
	return c.JSON(fiber.Map{
		"data": statutesArray,
		"meta": pagination,
	})
}
func LawRegulations(c *fiber.Ctx) error {
	regulations := []M.Law{}
	pagination := U.ParsedPagination(c)
	D.DB().Scopes(U.Paginate(&M.Law{}, pagination)).Where("type = ?", 1).Find(&regulations)
	regulationsArray := []M.LawMinimal{}
	for i := 0; i < len(regulations); i++ {
		if regulations[i].Type == 1 {
			regulationsArray = append(regulationsArray, M.LawMinimal{
				ID:               regulations[i].ID,
				Title:            regulations[i].Title,
				Image:            U.BaseURL + "/public/uploads/" + regulations[i].Image,
				NotificationDate: regulations[i].NotificationDate,
			})
		}
	}
	return c.JSON(fiber.Map{
		"data": regulationsArray,
		"meta": pagination,
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
	pagination := U.ParsedPagination(c)
	D.DB().Scopes(
		F.FilterByType(c,
			F.FilterType{QueryName: "title", ColumnName: "title", Operator: "LIKE"},
			F.FilterType{QueryName: "startDate", ColumnName: "notification_date", Operator: ">="},
			F.FilterType{QueryName: "endDate", ColumnName: "notification_date", Operator: "<="},
			F.FilterType{QueryName: "body", ColumnName: "body", Operator: "LIKE"}),
		U.Paginate(laws, pagination)).Find(&laws)
	pass_data := []M.LawMinimal_min{}
	for i := 0; i < len(laws); i++ {
		pass_data = append(pass_data, M.LawMinimal_min{
			ID:    laws[i].ID,
			Title: laws[i].Title,
			Image: U.BaseURL + "/public/uploads/" + laws[i].Image,
		})
	}
	return c.JSON(fiber.Map{
		"data": pass_data,
		"params": fiber.Map{
			"title":     c.Query("title"),
			"startDate": c.Query("startDate"),
			"endDate":   c.Query("endDate"),
			"body":      c.Query("body"),
		},
		"meta": pagination,
	})
}

// ! CHECK: files ham preload mishe. aya niazi?
func LawByID(c *fiber.Ctx) error {
	// ip := c.IP()
	// return c.JSON(fiber.Map{
	// 	"ip": ip,
	// 	"port": c.Port(),
	// 	"url" : c.BaseURL(),
	// 	"hostname": c.Hostname(),
		
	// })
	law := &M.Law{}
	if err := D.DB().First(law, c.Params("id")).Error; err != nil {
		return U.DBError(c, err)
	}
	LawByID := M.LawToLawByID(law)
	lawLog := &M.LawLog{
		LawID: LawByID.ID,
	}
	if err := D.DB().Create(&lawLog).Error; err != nil {
		return U.DBError(c, err)
	}
	return c.JSON(fiber.Map{
		"data": LawByID,
	})
}

// func CreateLaw(c *fiber.Ctx) error {
// 	payload := new(M.CreateLawInput)
// 	// parsing the payload
// 	if err := c.BodyParser(payload); err != nil {
// 		U.ResErr(c, err.Error())
// 	}
// 	if errs := U.Validate(payload); errs != nil {
// 		return c.Status(400).JSON(fiber.Map{"errors": errs})
// 	}
// 	law := M.Law{
// 		Type:               payload.Type,
// 		Title:              payload.Title,
// 		SessionNumber:      payload.SessionNumber,
// 		SessionDate:        payload.SessionDate,
// 		NotificationDate:   payload.NotificationDate,
// 		NotificationNumber: payload.NotificationNumber,
// 		Body:               payload.Body,
// 		// Image:              payload.Image,
// 	}
// 	result := D.DB().Create(&law)
// 	if result.Error != nil {
// 		return U.ResErr(c, result.Error.Error())
// 	}
// 	var tags = strings.Split(payload.Tags, ",")
// 	for i := 0; i < len(tags); i++ {
// 		result2 := D.DB().Create(&M.Keyword{
// 			Keyword: tags[i],
// 			LawID:   law.ID,
// 		})
// 		if result2.Error != nil {
// 			D.DB().Delete(&M.Law{}, law.ID)
// 			return U.ResErr(c, result.Error.Error())
// 			// return U.ResErr(c, "خطایی در اضافه کردن تگ ها پیش آمده است.")
// 		}
// 	}
// 	return c.Status(200).JSON(fiber.Map{
// 		"message": "مصوبه با موفقیت اضافه شد",
// 	})
// }

// offline one hundered laws
func OfflineLaws(c *fiber.Ctx) error {
	laws := []M.Law{}
	D.DB().Find(&laws)
	responseLaws := []M.LawMain{}
	for i := 0; i < len(laws); i++ {
		responseLaws = append(responseLaws, M.LawMain{
			ID:                 laws[i].ID,
			Type:               laws[i].Type,
			Title:              laws[i].Title,
			NotificationDate:   laws[i].NotificationDate,
			NotificationNumber: laws[i].NotificationNumber,
			Body:               laws[i].Body,
			CreatedAt:          laws[i].CreatedAt,
			UpdatedAt:          laws[i].UpdatedAt,
		})
	}
	return c.JSON(fiber.Map{"data": responseLaws})
}

func AddComment(c *fiber.Ctx) error {
	payload := new(M.CommentMinimal)
	if err := c.BodyParser(payload); err != nil {
		U.ResErr(c, err.Error())
	}
	if errs := U.Validate(payload); errs != nil {
		return c.Status(400).JSON(fiber.Map{"errors": errs})
	}
	comment := M.Comment{
		Body:            payload.Body,
		FullName:        payload.FullName,
		Email:           payload.Email,
		LawID:           payload.LawID,
		ParentCommentID: 0,
		Status:          false,
	}
	result := D.DB().Create(&comment)
	if result.Error != nil {
		return U.ResErr(c, result.Error.Error())
	}
	return c.Status(200).JSON(fiber.Map{
		"message": "کامنت با موفقیت اضافه شد",
	})
}
