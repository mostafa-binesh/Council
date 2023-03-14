package controllers

import (
	D "docker/database"
	M "docker/models"
	"github.com/gofiber/fiber/v2"
)

func AllLaws(c *fiber.Ctx) error {
	laws := []M.Law{}
	regulations := []M.LawMinimal{}
	statutes := []M.LawMinimal{}
	enactments := []M.LawMinimal{}
	D.DB().Find(&laws)
	// ! filtering
	for i := 0; i < len(laws); i++ {
		if laws[i].Type == 1 {
			regulations = append(regulations, M.LawMinimal{
				ID:        laws[i].ID,
				Title:     laws[i].Title,
				Image:     laws[i].Image,
				CreatedAt: laws[i].CreatedAt,
			})
		}
	}
	for i := 0; i < len(laws); i++ {
		if laws[i].Type == 2 {
			statutes = append(statutes, M.LawMinimal{
				ID:        laws[i].ID,
				Title:     laws[i].Title,
				Image:     laws[i].Image,
				CreatedAt: laws[i].CreatedAt,
			})
		}
	}
	for i := 0; i < len(laws); i++ {
		if laws[i].Type == 3 {
			enactments = append(enactments, M.LawMinimal{
				ID:        laws[i].ID,
				Title:     laws[i].Title,
				Image:     laws[i].Image,
				CreatedAt: laws[i].CreatedAt,
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
func Search(c *fiber.Ctx)error  {
	laws := []M.Law{}
	if c.FormValue("startDate")=="" || c.FormValue("endDate")==""{
		return c.Status(400).JSON(fiber.Map{
			"error": "هر دو تاریخ را پر کنید",
		},)
	}
	D.DB().Where("notification_date BETWEEN ? AND ?" , c.FormValue("startDate"),c.FormValue("endDate")).Find(&laws)
	return c.JSON(fiber.Map{
		"data" : fiber.Map{
			"law": laws,
		},
	})
}