package controllers

import (
	D "docker/database"
	M "docker/models"
	U "docker/utils"
	"fmt"
	"strconv"
	"time"

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
func LawSearch(c *fiber.Ctx) error {
	laws := []M.Law{}
	// ! nothing exists
	if c.FormValue("startDate") == "" && c.FormValue("endDate") == "" && c.FormValue("title") == "" {
		U.ResErr(c, "لطفا تاریخ ها یا عنوان را پر کنید")
		// ! only dates exists
	} else if c.FormValue("startDate") != "" &&
		c.FormValue("endDate") != "" &&
		c.FormValue("title") == "" {
		D.DB().Where("notification_date BETWEEN ? AND ?", c.FormValue("startDate"), c.FormValue("endDate")).Find(&laws)
		// ! only title exists
	} else if c.FormValue("startDate") == "" &&
		c.FormValue("endDate") != "" &&
		c.FormValue("title") == "" {
		D.DB().Where("notification_date <= ?", c.FormValue("endDate")).Find(&laws)
		// ! only end date
	} else if c.FormValue("startDate") != "" &&
		c.FormValue("endDate") == "" &&
		c.FormValue("title") == "" {
		D.DB().Where("notification_date >= ?", c.FormValue("startDate")).Find(&laws)
		// ! only start date
	} else if c.FormValue("startDate") == "" &&
		c.FormValue("endDate") == "" &&
		c.FormValue("title") != "" {
		D.DB().Where("title LIKE ?", fmt.Sprintf("%%%s%%", c.FormValue("title"))).Find(&laws)
		// ! only title exists
	} else if c.FormValue("startDate") == "" &&
		c.FormValue("endDate") != "" &&
		c.FormValue("title") != "" {
		D.DB().Where("title LIKE ? AND notification_date <= ?", fmt.Sprintf("%%%s%%", c.FormValue("title")),
			c.FormValue("endDate")).Find(&laws)
		// ! both enddate and title
	} else if c.FormValue("startDate") != "" &&
		c.FormValue("endDate") == "" &&
		c.FormValue("title") != "" {
		D.DB().Where("title LIKE ? AND notification_date >= ?", fmt.Sprintf("%%%s%%", c.FormValue("title")),
			c.FormValue("startDate")).Find(&laws)
		// ! both startDate and title
	} else if c.FormValue("startDate") != "" &&
		c.FormValue("endDate") != "" &&
		c.FormValue("title") != "" {
		D.DB().Where("title LIKE ? AND notification_date BETWEEN ? AND ?", fmt.Sprintf("%%%s%%", c.FormValue("title")),
			c.FormValue("startDate"), c.FormValue("endDate")).Find(&laws)
		// ! evry thing exist
	} else {
		return U.ResErr(c, "")
	}
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
	// else if c.FormValue("startDate") == "" || c.FormValue("endDate") == "" {
	// 	D.DB().Where("title LIKE %?%", c.FormValue("title")).Find(&laws)
	// }

}
func LawByID(c *fiber.Ctx) error {
	law := &M.Law{}
	if err := D.DB().Preload("Comments.User").First(law, c.Params("id")).Error; err != nil {
		return c.SendString(err.Error())
		// return U.DBError(c, err)
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
	// law.Comments = M.GetMinimalComment(GetMinimalComment)
	return c.JSON(fiber.Map{
		"data": lawWithComment,
	})
}

// Type               int       `json:"type" gorm:"type:int;not null"`
//
//	Title              string    `json:"title" gorm:"type:varchar(100);not null"`
//	SessionNumber      int       `json:"sessionNumber" gorm:"type:int;not null"`
//	SessionDate        time.Time `json:"sessionDate" gorm:"not null;default:now()"`      // ! change default now later
//	NotificationDate   time.Time `json:"notificationDate" gorm:"not null;default:now()"` // ! change default now later
//	NotificationNumber string    `json:"notificationNumber" gorm:"not null"`
//	Body               string    `json:"body" gorm:"type:text;not null"`
//	Image
func CreateLaw(c *fiber.Ctx) error {
	if c.FormValue("type") == "" || c.FormValue("sessionNumber") == "" || c.FormValue("sessionDate") == "" ||
		c.FormValue("notificationDate") == "" || c.FormValue("title") == "" ||
		c.FormValue("notificationNumber") == "" || c.FormValue("body") == "" {
		return c.Status(400).JSON(fiber.Map{
			"error": "لطفا تمام فیلد ها را پر کنید",
		})
	}
	tp, err := strconv.Atoi(c.FormValue("type"))
	sNumber, err := strconv.Atoi(c.FormValue("sessionNumber"))
	sDate, err := time.Parse("2006-01-02", c.FormValue("sessionDate"))
	nDate, err := time.Parse("2006-01-02", c.FormValue("notificationDate"))
	if err != nil {
	}
	result := D.DB().Create(&M.Law{
		Type:               tp,
		Title:              c.FormValue("title"),
		SessionNumber:      sNumber,
		SessionDate:        sDate,
		NotificationDate:   nDate,
		NotificationNumber: c.FormValue("notificationNumber"),
		Body:               c.FormValue("body"),
		Image:              "https://s2.uupload.ir/files/placeholder-image_ux76.png",
	})

	if result.Error != nil {
		return c.JSON(fiber.Map{
			"message": "خطایی در اضافه کردن مصوبه پیش آمده است.",
		})
	}
	return c.JSON(fiber.Map{
		"message": "مصوبه با موفقیت اضافه شد",
	})
}
