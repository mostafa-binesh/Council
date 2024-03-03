package admin

import (
	D "docker/database"
	F "docker/database/filters"
	U "docker/utils"
	"fmt"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	// "strconv"

	// F "docker/database/filters"
	M "docker/models"
)

func IndexLaw(c *fiber.Ctx) error {
	laws := []M.Law{}
	pagination := U.ParsedPagination(c)
	D.DB().Scopes(U.Paginate(laws, pagination)).Find(&laws)
	responseLaws := []M.LawMinimal_min{}
	for i := 0; i < len(laws); i++ {
		responseLaws = append(responseLaws, M.LawMinimal_min{
			ID:    laws[i].ID,
			Title: laws[i].Title,
			Image: U.BaseURL + "/public/uploads/" + laws[i].Image,
		})
	}
	if !M.GetLog(c) {
		return U.CreateNewLogError(c)
	}
	return U.ResWithPagination(c, responseLaws, *pagination)
}
func UpdateLaw(c *fiber.Ctx) error {
	law := M.Law{}
	payload := new(M.EditLawInput)
	// parsing the payload
	if err := c.BodyParser(payload); err != nil {
		U.ResErr(c, err.Error())
	}
	// items := c.Request().PostArgs().PeekMulti("attachmentsId[]")
	if errs := U.Validate(payload); errs != nil {
		return c.Status(400).JSON(fiber.Map{"errors": errs})
	}
	lawId := c.Params("id")
	result1 := D.DB().Where("id = ?", lawId).Find(&law)
	if result1.Error != nil {
		return U.DBError(c, result1.Error)
	}
	law.Body = payload.Body
	// law.Image = payload.Image
	law.NotificationDate = payload.NotificationDate
	law.NotificationNumber = payload.NotificationNumber
	law.SessionDate = payload.SessionDate
	law.SessionNumber = payload.SessionNumber
	law.Title = payload.Title
	law.Type = payload.Type

	imageFile, _ := c.FormFile("image")
	if imageFile != nil {
		fmt.Println("till here")
		// check if imageFile with this name already exists
		if U.FileExistenceCheck(imageFile.Filename, "./public/uploads") {
			return U.ResErr(c, "imageFile already exists")
		}
		// ! imageFile extension check
		// if !(U.HasImageSuffixCheck(imageFile.Filename) || U.HasSuffixCheck(imageFile.Filename, []string{"pdf"})) {
		// 	return c.SendString("imageFile should be image or pdf! please fix it")
		// }
		// Save imageFile to disk
		fileName := U.AddUUIDToString(imageFile.Filename)
		err := c.SaveFile(imageFile, fmt.Sprintf("./public/uploads/%s", fileName))
		if err != nil {
			return U.ResErr(c, err.Error())
		}
		law.Image = fileName
	}
	result := D.DB().Save(&law)
	if result.Error != nil {
		return U.ResErr(c, "مشکلی در به روز رسانی به وجود آمده")
	}

	// ! store ExplanatoryPlan is exists
	file, _ := c.FormFile("explanatoryPlan")
	// if formError != nil {
	// 	return U.ResErr(c, formError.Error())
	// }
	if file != nil {
		fmt.Println("till here")
		// check if file with this name already exists
		if U.FileExistenceCheck(file.Filename, "./public/uploads") {
			return U.ResErr(c, "file already exists")
		}
		// ! file extension check
		// if !(U.HasImageSuffixCheck(file.Filename) || U.HasSuffixCheck(file.Filename, []string{"pdf"})) {
		// 	return c.SendString("file should be image or pdf! please fix it")
		// }
		// Save file to disk
		fileName := U.AddUUIDToString(file.Filename)
		c.SaveFile(file, fmt.Sprintf("./public/uploads/%s", fileName))
		D.DB().Create(&M.File{
			Type:  M.FileTypes["plan"],
			Name:  fileName,
			LawID: law.ID,
		})
	}
	// ! certificate
	file, _ = c.FormFile("certificate")
	if file != nil {
		// check if file with this name already exists
		if U.FileExistenceCheck(file.Filename, "./public/uploads") {
			return U.ResErr(c, "file already exists")
		}
		// ! file extension check
		// if !(U.HasImageSuffixCheck(file.Filename) || U.HasSuffixCheck(file.Filename, []string{"pdf"})) {
		// 	return c.SendString("file should be image or pdf! please fix it")
		// }
		// Save file to disk
		fileName := U.AddUUIDToString(file.Filename)
		c.SaveFile(file, fmt.Sprintf("./public/uploads/%s", fileName))
		D.DB().Create(&M.File{
			Type:  M.FileTypes["certificate"],
			Name:  fileName,
			LawID: law.ID,
		})
	}
	// ! attachments
	// attachments, _ := c.FormFile("explanatoryPlan")
	// form, _ := c.MultipartForm()
	// attachments := form.File["attachment[]"]
	// for _, file := range attachments {
	// 	// check if file with this name already exists
	// 	if U.FileExistenceCheck(file.Filename, "./public/uploads") {
	// 		return U.ResErr(c, "file already exists")
	// 	}
	// 	// ! file extension check
	// 	// if !(U.HasImageSuffixCheck(file.Filename) || U.HasSuffixCheck(file.Filename, []string{"pdf"})) {
	// 	// 	return c.SendString("file should be image or pdf! please fix it")
	// 	// }
	// 	// Save file to disk
	// 	// err = c.SaveFile(file, fmt.Sprintf("./public/uploads/%s", file.Filename))
	// 	fileName := U.AddUUIDToString(file.Filename)
	// 	c.SaveFile(file, fmt.Sprintf("./public/uploads/%s", fileName))
	// 	D.DB().Create(&M.File{
	// 		Type:  M.FileTypes["attachment"],
	// 		Name:  fileName,
	// 		LawID: law.ID,
	// 	})
	// 	// if err != nil {
	// 	// 	return U.ResErr(c, "cannot save")
	// 	// }
	// }

	// var attachmentFilesId []uint64
	// if result := D.DB().Where("law_id = ? AND type = ?", lawId, M.FileTypes["attachment"]).Pluck("id", attachmentFilesId); result.Error != nil {
	// 	return U.DBError(c, result.Error)
	// }
	// shouldRemoveFile := U.Difference(attachmentFilesId, payload.AttachmentsId)
	// if result := D.DB().Where("id IN ?", shouldRemoveFile).Delete(&M.File{}); result.Error != nil {
	// 	return U.DBError(c, result.Error)
	// }

	// todo: handle removing the file from the storage later
	return c.JSON(fiber.Map{
		"message": "به روز رسانی با موفقیت انجام شد",
	})
}
func DeleteLaw(c *fiber.Ctx) error {
	// just to pass the bug
	result := D.DB().Delete(&M.Keyword{}, "law_id = ?", c.Params("id"))
	if result.Error != nil {
		return U.DBError(c, result.Error)
	}
	result = D.DB().Delete(&M.Law{}, c.Params("id"))
	if result.Error != nil {
		return U.DBError(c, result.Error)
	}
	if result.RowsAffected == 0 {
		return U.ResErr(c, "مصوبه یافت نشد")
	}

	return c.JSON(fiber.Map{
		"message": "حذف کردن با موفقیت انجام شد",
	})
}
func CreateLaw(c *fiber.Ctx) error {
	payload := new(M.CreateLawInput)
	// parse payload
	if err := c.BodyParser(payload); err != nil {
		U.ResErr(c, err.Error())
	}
	// validate payload
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
		Recommender:        payload.Recommender,
		NumberItems:        payload.NumberItems,
		NumberNotes:        payload.NumberNotes,
		// Image:              payload.Image,
	}
	// ! store ExplanatoryPlan is exists
	imageFile, _ := c.FormFile("image")

	if imageFile == nil {
		return U.ResValidationErr(c, map[string]string{"image": "آپلود عکس ضروری است"})
	}
	if imageFile != nil {
		// check if imageFile with this name already exists
		if U.FileExistenceCheck(imageFile.Filename, "./public/uploads") {
			return U.ResErr(c, "imageFile already exists")
		}
		// ! imageFile extension check
		// if !(U.HasImageSuffixCheck(imageFile.Filename) || U.HasSuffixCheck(imageFile.Filename, []string{"pdf"})) {
		// 	return c.SendString("imageFile should be image or pdf! please fix it")
		// }
		// Save imageFile to disk
		fileName := U.AddUUIDToString(imageFile.Filename)
		err := c.SaveFile(imageFile, fmt.Sprintf("./public/uploads/%s", fileName))
		if err != nil {
			return U.ResErr(c, err.Error())
		}
		law.Image = fileName
	}
	// store law in the db
	result := D.DB().Create(&law)
	if result.Error != nil {
		return U.ResErr(c, result.Error.Error())
	}
	// parse tags and store them
	var tags = strings.Split(payload.Tags, ",")
	for i := 0; i < len(tags); i++ {
		result2 := D.DB().Create(&M.Keyword{
			Keyword: tags[i],
			LawID:   law.ID,
		})
		if result2.Error != nil {
			D.DB().Delete(&M.Law{}, law.ID)
			return U.ResErr(c, result.Error.Error())
		}
	}

	// ! store ExplanatoryPlan is exists
	file, _ := c.FormFile("explanatoryPlan")
	// if formError != nil {
	// 	return U.ResErr(c, formError.Error())
	// }
	if file != nil {
		fmt.Println("till here")
		// check if file with this name already exists
		if U.FileExistenceCheck(file.Filename, "./public/uploads") {
			return U.ResErr(c, "file already exists")
		}
		// ! file extension check
		// if !(U.HasImageSuffixCheck(file.Filename) || U.HasSuffixCheck(file.Filename, []string{"pdf"})) {
		// 	return c.SendString("file should be image or pdf! please fix it")
		// }
		// Save file to disk
		fileName := U.AddUUIDToString(file.Filename)
		err := c.SaveFile(file, fmt.Sprintf("./public/uploads/%s", fileName))
		if err != nil {
			return U.ResErr(c, err.Error())
		}
		D.DB().Create(&M.File{
			Type:  M.FileTypes["plan"],
			Name:  fileName,
			LawID: law.ID,
		})
	}
	// ! certificate
	file, _ = c.FormFile("certificate")
	if file != nil {
		// check if file with this name already exists
		if U.FileExistenceCheck(file.Filename, "./public/uploads") {
			return U.ResErr(c, "file already exists")
		}
		// ! file extension check
		// if !(U.HasImageSuffixCheck(file.Filename) || U.HasSuffixCheck(file.Filename, []string{"pdf"})) {
		// 	return c.SendString("file should be image or pdf! please fix it")
		// }
		// Save file to disk
		fileName := U.AddUUIDToString(file.Filename)
		err := c.SaveFile(file, fmt.Sprintf("./public/uploads/%s", fileName))
		if err != nil {
			return U.ResErr(c, "خطا در ذخیره ی فایل")
		}
		D.DB().Create(&M.File{
			Type:  M.FileTypes["certificate"],
			Name:  fileName,
			LawID: law.ID,
		})
	}
	// ! attachments
	form, _ := c.MultipartForm()
	attachments := form.File["attachment[]"]
	for _, file := range attachments {
		// check if file with this name already exists
		if U.FileExistenceCheck(file.Filename, "./public/uploads") {
			return U.ResErr(c, "file already exists")
		}
		// ! file extension check
		// if !(U.HasImageSuffixCheck(file.Filename) || U.HasSuffixCheck(file.Filename, []string{"pdf"})) {
		// 	return c.SendString("file should be image or pdf! please fix it")
		// }
		// Save file to disk
		// err = c.SaveFile(file, fmt.Sprintf("./public/uploads/%s", file.Filename))
		fileName := U.AddUUIDToString(file.Filename)
		err := c.SaveFile(file, fmt.Sprintf("./public/uploads/%s", fileName))
		if err != nil {
			return U.ResErr(c, "خطا در ذخیره ی فایل")
		}
		D.DB().Create(&M.File{
			Type:  M.FileTypes["attachment"],
			Name:  fileName,
			LawID: law.ID,
		})
		// if err != nil {
		// 	return U.ResErr(c, "cannot save")
		// }
	}

	// return response
	return c.Status(200).JSON(fiber.Map{
		"message": "مصوبه با موفقیت اضافه شد",
	})
}
func LawSearch(c *fiber.Ctx) error {
	laws := []M.Law{}
	pagination := U.ParsedPagination(c)
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
		U.Paginate(laws, pagination)).
		Find(&laws)
	pass_data := []M.LawMinimal_min{}
	for i := 0; i < len(laws); i++ {
		pass_data = append(pass_data, M.LawMinimal_min{
			ID:    laws[i].ID,
			Title: laws[i].Title,
			// Image: laws[i].Image,
		})
	}
	return c.JSON(fiber.Map{
		"meta": pagination,
		"data": pass_data,
	})
}
func LawByID(c *fiber.Ctx) error {
	law := &M.Law{}
	if err := D.DB().Preload("Files").First(law, c.Params("id")).Error; err != nil {
		return U.DBError(c, err)
	}
	LawByID := M.LawToSeenAdmin(law)

	return c.JSON(fiber.Map{
		"data": LawByID,
	})
}
func CommentsByLawID(c *fiber.Ctx) error {
	comments := []M.Comment{}
	pagination := U.ParsedPagination(c)
	D.DB().Where("law_id = ?", c.Params("id")).Scopes(U.Paginate(comments, pagination)).Find(&comments)
	var minimalComments []M.CommentMinimal
	for i := 0; i < len(comments); i++ {
		minimalComment := M.CommentMinimal{
			ID:       comments[i].ID,
			Email:    comments[i].Email,
			FullName: comments[i].FullName,
			Body:     comments[i].Body,
			Status:   comments[i].Status,
		}
		minimalComments = append(minimalComments, minimalComment)

	}

	return c.JSON(fiber.Map{
		"data": minimalComments,
		"meta": pagination,
	})
}
func DeleteFile(c *fiber.Ctx) error {
	result := D.DB().Delete(&M.File{}, c.Params("fileID"))
	if result.Error != nil {
		return U.DBError(c, result.Error)
	}
	if result.RowsAffected == 0 {
		return U.ResErr(c, "فایل یافت نشد")
	}

	return c.JSON(fiber.Map{
		"message": "فایل حذف شد",
	})
}
func UploadFile(c *fiber.Ctx) error {
	// create payload
	payload := new(M.UploadFile)
	// parse payload
	if err := c.BodyParser(payload); err != nil {
		U.ResErr(c, err.Error())
	}
	// validate payload
	if errs := U.Validate(payload); errs != nil {
		return c.Status(400).JSON(fiber.Map{"errors": errs})
	}
	// check file existence
	file, _ := c.FormFile("file")
	if file == nil {
		return U.ResValidationErr(c, map[string]string{"file": "فایل ضروری است"})
	}
	// check if file with this name already exists
	if U.FileExistenceCheck(file.Filename, "./public/uploads") {
		return U.ResErr(c, "file already exists")
	}
	// Save file to disk
	fileName := U.AddUUIDToString(file.Filename)
	err := c.SaveFile(file, fmt.Sprintf("./public/uploads/%s", fileName))
	if err != nil {
		return U.ResErr(c, err.Error())
	}
	dbFile := M.File{
		Type:  M.FileTypes[payload.Type],
		Name:  fileName,
		LawID: payload.LawId,
	}
	if result := D.DB().Create(&dbFile); result.Error != nil {
		return U.DBError(c, result.Error)
	}

	// return U.ResMessage(c, "فایل آپلود شد")
	return c.JSON(fiber.Map{"id": dbFile.ID})
}
func RemoveFile(c *fiber.Ctx) error {
	// create payload
	payload := new(M.RemoveFileInput)
	// parse payload
	if err := c.BodyParser(payload); err != nil {
		U.ResErr(c, err.Error())
	}
	// validate payload
	if errs := U.Validate(payload); errs != nil {
		return c.Status(400).JSON(fiber.Map{"errors": errs})
	}
	// remove the file from database
	result := D.DB().Where("id = ? AND law_id = ?", payload.FileID, payload.LawID).Delete(&M.File{})
	if result.Error != nil {
		return U.DBError(c, result.Error)
	}
	if result.RowsAffected == 0 {
		return U.ResErr(c, "فایلی یافت نشد")
	}

	// return U.ResMessage(c, "فایل آپلود شد")
	return U.ResMessage(c, "فایل حذف شد")
}

func Statics(c *fiber.Ctx) error {

	var results []struct {
		LawID uint
		Count int64
		Title string
	}
	D.DB().Model(&M.LawLog{}).
		Select("law_logs.law_id, laws.title, count(law_logs.law_id) as count").
		Joins("INNER JOIN laws ON laws.id = law_logs.law_id").
		Where("law_logs.created_at >= ?", time.Now().AddDate(0, 0, -10)).
		Group("law_logs.law_id, laws.title").
		Limit(10).
		Scan(&results)

	return c.JSON(fiber.Map{
		"data": results,
	})
}

func VerifyComment(c *fiber.Ctx) error {
	result := D.DB().Model(&M.Comment{}).Where("id = ?", c.Params("id")).Update("status", true)
	if result.Error != nil {
		U.DBError(c, result.Error)
	}

	return U.ResMessage(c, "کاربر تایید شد")
}
func UnVerifyComment(c *fiber.Ctx) error {
	result := D.DB().Model(&M.Comment{}).Where("id = ?", c.Params("id")).Update("status", false)
	if result.Error != nil {
		U.DBError(c, result.Error)
	}

	return U.ResMessage(c, "کاربر تایید شد")
}
