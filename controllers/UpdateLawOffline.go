package controllers

import (
	D "docker/database"
	M "docker/models"
	U "docker/utils"
	"strings"

	// "time"
	"github.com/gofiber/fiber/v2"
)

func UpdateLawOffline(c *fiber.Ctx) error {
	payload := new(M.UpdatedLaws)
	if err := c.BodyParser(payload); err != nil {
		U.ResErr(c, err.Error())
	}
	if errs := U.Validate(payload); errs != nil {
		return c.Status(400).JSON(fiber.Map{"errors": errs})
	}
	laws := []M.Law{}
	D.DB().Where("created_at >= ?", payload.LastOnline).Find(&laws)
	responseLawsCreate := []M.LawMain{}
	for i := 0; i < len(laws); i++ {
		responseLawsCreate = append(responseLawsCreate, M.LawMain{
			ID:                 laws[i].ID,
			Type:               laws[i].Type,
			Title:              laws[i].Title,
			NotificationDate:   laws[i].NotificationDate,
			NotificationNumber: laws[i].NotificationNumber,
			Body:               laws[i].Body,
		})
	}
	D.DB().Where("updated_at >= ?", payload.LastOnline).Find(&laws)
	responseLawsUpdate := []M.LawMain{}
	for i := 0; i < len(laws); i++ {
		responseLawsUpdate = append(responseLawsUpdate, M.LawMain{
			ID:                 laws[i].ID,
			Type:               laws[i].Type,
			Title:              laws[i].Title,
			NotificationDate:   laws[i].NotificationDate,
			NotificationNumber: laws[i].NotificationNumber,
			Body:               laws[i].Body,
		})
	}
	var lawDelID []string
	var lawDel []M.ActionLog
	// اجرای کوئری و جمع‌آوری نتایج
	D.DB().Where("request_type = ? AND route_name LIKE ?", "DELETE", "%laws%").Where("action_time >= ?", payload.LastOnline).Find(&lawDel)
	// پیمایش نتایج و استخراج قسمت مورد نیاز
	for i := 0; i < len(lawDel); i++ {
		parts := strings.Split(lawDel[i].RouteName, "/")
		if len(parts) >= 3 {
			lawDelID = append(lawDelID, parts[len(parts)-1])
		}
	}
	return c.JSON(fiber.Map{
		"deleteLaw": lawDelID,
		"updateLaw": responseLawsUpdate,
		"createLaw": responseLawsCreate,
	})
}
