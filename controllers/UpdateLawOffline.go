package controllers

import (
	D "docker/database"
	M "docker/models"
	U "docker/utils"
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
	D.DB().Where("created_at >= ?",payload.LastOnline).Where("updated_at >= ?",payload.LastOnline).Find(&laws)
	responseLaws := []M.LawMain{}
	for i := 0; i < len(laws); i++ {
		responseLaws = append(responseLaws, M.LawMain{
			ID:                 laws[i].ID,
			Type:               laws[i].Type,
			Title:              laws[i].Title,
			NotificationDate:   laws[i].NotificationDate,
			NotificationNumber: laws[i].NotificationNumber,
			Body:               laws[i].Body,
		})
	}
	return c.JSON(fiber.Map{"data": responseLaws})

}