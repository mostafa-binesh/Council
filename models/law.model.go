package models

import (
	D "docker/database"
	U "docker/utils"
	"time"
)

type Law struct {
	ID                 uint      `json:"id" gorm:"primary_key"`
	Type               int       `json:"type" gorm:"type:int;not null"`
	Title              string    `json:"title" gorm:"type:varchar(100);not null"`
	SessionNumber      int       `json:"sessionNumber" gorm:"type:int;not null"`
	SessionDate        time.Time `json:"sessionDate" gorm:"not null;default:now()"`      // ! change default now later
	NotificationDate   time.Time `json:"notificationDate" gorm:"not null;default:now()"` // ! change default now later
	NotificationNumber string    `json:"notificationNumber" gorm:"not null"`
	Body               string    `json:"body" gorm:"type:text;not null"`
	Image              string    `json:"image" gorm:"type:varchar(255);not null"`
	Comments           []Comment `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Files              []File    `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	NumberItems        int       `json:"NumberItems" gorm:"type:int;not null"`
	NumberNotes        int       `json:"NumberNotes" gorm:"type:int;not null"`
	Recommender        string    `json:"Recommender" gorm:"not null"`
	CreatedAt          time.Time `json:"createdAt" gorm:"not null;default:now()"`
	UpdatedAt          time.Time `json:"updatedAt" gorm:"not null;default:now()"`
}
type LawByID struct {
	ID                 uint             `json:"id"`
	Type               int              `json:"type"`
	Title              string           `json:"title"`
	SessionNumber      int              `json:"sessionNumber"`
	SessionDate        time.Time        `json:"sessionDate"`
	NotificationDate   time.Time        `json:"notificationDate"`
	NotificationNumber string           `json:"notificationNumber"`
	Body               string           `json:"body"`
	Image              string           `json:"image"`
	Comments           []CommentMinimal `json:"comments"`
	SeenCount          int64            `json:"seenCount"`
	Files              []FileMinimal    `json:"files"`
	NumberItems        int              `json:"NumberItems"`
	NumberNotes        int              `json:"NumberNotes"`
	Recommender        string           `json:"Recommender"`
}
type LawMain struct {
	ID                 uint             `json:"id"`
	Type               int              `json:"type"`
	Title              string           `json:"title"`
	NotificationDate   time.Time        `json:"notificationDate"`
	NotificationNumber string           `json:"notificationNumber"`
	Body               string           `json:"body"`
	Image              string           `json:"image"`
	SeenCount          int64            `json:"seenCount"`
	Comments           []CommentMinimal `json:"comments"`
	CreatedAt          time.Time        `json:"created_at"`
	UpdatedAt          time.Time        `json:"updated_at"`
}
type LawMinimal struct {
	ID               uint      `json:"id"`
	Title            string    `json:"title"`
	Image            string    `json:"image"`
	NotificationDate time.Time `json:"date"`
}
type LawMinimal_min struct {
	ID    uint   `json:"id"`
	Title string `json:"title"`
	Image string `json:"image"`
}
type LawStatutesMinimal struct {
	ID               uint      `json:"id"`
	Title            string    `json:"title"`
	Image            string    `json:"image"`
	SessionNumber    int       `json:"sessionNumber"`
	NotificationDate time.Time `json:"date"`
}
type CreateLawInput struct {
	Type               int       `json:"type" validate:"required"`
	Title              string    `json:"title"  validate:"required"`
	SessionNumber      int       `json:"sessionNumber"`
	SessionDate        time.Time `json:"sessionDate" validate:"required"`      // ! change default now later
	NotificationDate   time.Time `json:"notificationDate" validate:"required"` // ! change default now later
	NotificationNumber string    `json:"notificationNumber" validate:"required"`
	Body               string    `json:"body" validate:"required"`
	NumberItems        int       `json:"numberItems"`
	NumberNotes        int       `json:"numberNotes"`
	Recommender        string    `json:"recommender"`
	Tags               string    `json:"tags" validate:"required"`
}
type EditLawInput struct {
	Type               int       `json:"type" validate:"required"`
	Title              string    `json:"title"  validate:"required"`
	SessionNumber      int       `json:"sessionNumber"`
	SessionDate        time.Time `json:"sessionDate" validate:"required"`      // ! change default now later
	NotificationDate   time.Time `json:"notificationDate" validate:"required"` // ! change default now later
	NotificationNumber string    `json:"notificationNumber" validate:"required"`
	Body               string    `json:"body" validate:"required"`
	NumberItems        int       `json:"numberItems"`
	NumberNotes        int       `json:"numberNotes"`
	Recommender        string    `json:"recommender"`
	Tags               string    `json:"tags" validate:"required"`
	AttachmentsId      []uint64  `json:"attachmentsId" validate:"required"`
}
type UpdatedLaws struct {
	LastOnline time.Time `json:"lastOnline" validate:"required"` // ! change default now later
}
type Comment struct {
	ID              uint      `json:"id" gorm:"primary_key"`
	Body            string    `json:"body" gorm:"type:text;not null"`
	FullName        string    `json:"fullName" gorm:"type:varchar(100)"`
	Email           string    `json:"email" gorm:"type:varchar(100)"`
	ParentCommentID uint      `json:"parentCommentID" gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE;OnDelete:CSCADE"`
	LawID           uint      `gorm:"foreignKey:LawID;constraint:OnUpdate:CASCADE;OnDelete:CSCADE"`
	Status          bool      `json:"status" gorm:"boolean"`
	ParentLaw       *Law      `json:"parentLaw" gorm:"foreignKey:LawID;constraint:OnUpdate:CASCADE;OnDelete:CSCADE"`
	CreatedAt       time.Time `json:"createdAt" gorm:"not null;default:now()"`
	UpdatedAt       time.Time `json:"updatedAt" gorm:"not null;default:now()"`
}
type CommentMinimal struct {
	ID       uint   `json:"id"`
	Body     string `json:"body" validate:"required"`
	FullName string `json:"fullName"`
	Email    string `json:"email"`
	Status   bool   `json:"status"`
	LawID    uint   `json:"law_id" validate:"required"`
}
type OfflineLaws struct {
}

// type UserMigration struct {
// 	// ID        *uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
// 	ID          uint       `gorm:"primary_key"`
// 	Name        string     `gorm:"type:varchar(255);not null"`
// 	LastName    string     `gorm:"type:varchar(255);not null"`
// 	Username    string     `gorm:"type:varchar(255);not null"`
// 	PhoneNumber string     `gorm:"type:varchar(255);not null"`
// 	Email       string     `gorm:"type:varchar(255);not null"`
// 	Password    string     `gorm:"type:varchar(255);not null"`
// 	CreatedAt   time.Time `gorm:"not null;default:now()"`
// 	UpdatedAt   time.Time `json:"updatedAt" gorm:"not null;default:now()"`
// }

type Keyword struct {
	ID        uint   `gorm:"primary_key"`
	Keyword   string `gorm:"type:varchar(70)"`
	LawID     uint
	Law       *Law      `gorm:"foreignKey:LawID;constraint:OnUpdate:CASCADE;OnDelete:CSCADE"`
	CreatedAt time.Time `gorm:"not null;default:now()"`
	UpdatedAt time.Time `json:"updatedAt" gorm:"not null;default:now()"`
}
type Attachment struct {
	ID        uint   `gorm:"primary_key"`
	FileName  string `gorm:"type:varchar(255);not null"`
	LawID     uint
	Type      int       `gorm:"type:int;not null"`
	Law       *Law      `gorm:"foreignKey:LawID;constraint:OnUpdate:CASCADE;OnDelete:CSCADE"`
	CreatedAt time.Time `gorm:"not null;default:now()"`
	UpdatedAt time.Time `json:"updatedAt" gorm:"not null;default:now()"`
}
type FAQ struct {
	ID           uint   `gorm:"primary_key"`
	Question     string `gorm:"type:varchar(255);not null"`
	Answer       string `gorm:"type:varchar(255);not null"`
	QuestionerID uint
	Questioner   *User `gorm:"foreignKey:QuestionerID;not null"`
	AnswererID   uint
	Answerer     *User     `gorm:"foreignKey:AnswererID;constraint:OnUpdate:CASCADE;OnDelete:CSCADE;not null"`
	CreatedAt    time.Time `gorm:"not null;default:now()"`
	UpdatedAt    time.Time `json:"updatedAt" gorm:"not null;default:now()"`
}

func GetMinimalComment(lawID uint, flag bool) []CommentMinimal {
	comments := []Comment{}
	if !flag {
		D.DB().Where("law_id = ?", lawID).Where("status = true").Find(&comments)
	}
	if flag {
		D.DB().Where("law_id = ?", lawID).Find(&comments)
	}
	var minimalComments []CommentMinimal
	for i := 0; i < len(comments); i++ {
		minimalComment := CommentMinimal{
			ID:       comments[i].ID,
			Email:    comments[i].Email,
			FullName: comments[i].FullName,
			Body:     comments[i].Body,
			Status:   comments[i].Status,
		}
		minimalComments = append(minimalComments, minimalComment)

	}
	return minimalComments
}
func getSeenCount(lawID uint) int64 {
	var count int64
	D.DB().Model(&LawLog{}).Where("law_id = ?", lawID).Count(&count)
	return count
}
func LawToLawByID(law *Law) *LawMain {
	return &LawMain{
		ID:                 law.ID,
		Type:               law.Type,
		Title:              law.Title,
		NotificationDate:   law.NotificationDate,
		NotificationNumber: law.NotificationNumber,
		Body:               law.Body,
		Image:              U.BaseURL + "/public/uploads/" + law.Image,
		Comments:           GetMinimalComment(law.ID, false),
		SeenCount:          getSeenCount(law.ID),
	}
}
func getFilesMini(lawID uint) []FileMinimal {
	files := []File{}
	D.DB().Where("law_id = ?", lawID).Find(&files)
	var fileArray []FileMinimal
	for i := 0; i < len(files); i++ {
		file := FileMinimal{
			ID:   files[i].ID,
			Type: files[i].Type,
			URL:  U.BaseURL + "/public/uploads/" + files[i].Name,
		}
		fileArray = append(fileArray, file)

	}
	return fileArray
}
func LawToSeenAdmin(law *Law) *LawByID {
	return &LawByID{
		ID:                 law.ID,
		Type:               law.Type,
		Title:              law.Title,
		SessionNumber:      law.SessionNumber,
		SessionDate:        law.SessionDate,
		NotificationDate:   law.NotificationDate,
		NotificationNumber: law.NotificationNumber,
		Body:               law.Body,
		Image:              U.BaseURL + "/public/uploads/" + law.Image,
		// Comments:           GetMinimalComment(law.ID, true),
		SeenCount:          getSeenCount(law.ID),
		Files:              getFilesMini(law.ID),
		NumberItems:        law.NumberItems,
		NumberNotes:        law.NumberNotes,
		Recommender:        law.Recommender,
	}
}
