package controllers

import (
	D "docker/database"
	M "docker/models"
	U "docker/utils"
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
				NotificationDate: laws[i].CreatedAt,
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
				NotificationDate: laws[i].CreatedAt,
			})
		}
	}
	for i := 0; i < len(laws); i++ {
		if laws[i].Type == 3 {
			enactments = append(enactments, M.LawMinimal{
				ID:               laws[i].ID,
				Title:            laws[i].Title,
				Image:            laws[i].Image,
				NotificationDate: laws[i].CreatedAt,
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
func LawsEnactments(c *fiber.Ctx)error  {
	enactments := M.Law{}
	D.DB().Where("Type = ?", 3).Find(&enactments)
	return c.JSON(fiber.Map{
		"data" : fiber.Map{
			"enactments":enactments,
		},
	})
}
func LawsStatutes(c *fiber.Ctx)error  {
	statutes := M.Law{}
	D.DB().Where("Type = ?", 2).Find(&statutes)
	return c.JSON(fiber.Map{
		"data" : fiber.Map{
			"statutes":statutes,
		},
	})
}
func LawsRegulations(c *fiber.Ctx)error  {
	regulations := M.Law{}
	D.DB().Where("Type = ?", 1).Find(&regulations)
	return c.JSON(fiber.Map{
		"data" : fiber.Map{
			"regulations":regulations,
		},
	})
}
func LawsSearch(c *fiber.Ctx) error {
	laws := []M.Law{}
	if c.FormValue("startDate") == "" || c.FormValue("endDate") == "" || c.FormValue("title") == "" {
		return c.Status(400).JSON(fiber.Map{
			"error": "لطفا تاریخ ها یا عنوان را پر کنید ",
		})
	}
	if c.FormValue("titel") == "" {
		D.DB().Where("notification_date BETWEEN ? AND ?", c.FormValue("startDate"), c.FormValue("endDate")).Find(&laws)
		return c.JSON(fiber.Map{
			"data": fiber.Map{
				"law": laws,
			},
		})
	}
	if c.FormValue("startDate") == "" || c.FormValue("endDate") == "" {
		D.DB().Where("title = ?", c.FormValue("title")).Find(&laws)
		return c.JSON(fiber.Map{
			"data": fiber.Map{
				"law": laws,
			},
		})
	}

	D.DB().Where("title = ? AND notification_date BETWEEN ? AND ?", c.FormValue("title"),
		c.FormValue("startDate"), c.FormValue("endDate")).Find(&laws)
	return c.JSON(fiber.Map{
		"data": fiber.Map{
			"law": laws,
		},
	})
}
func LawByID(c *fiber.Ctx) error {
	law := &M.Law{}
	if err := D.DB().First(law, c.Params("id")).Error; err != nil {
		return U.DBError(c, err)
	}
	return c.JSON(fiber.Map{
		"data": law,
	})
}
