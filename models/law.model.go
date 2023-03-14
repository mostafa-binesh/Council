package models

import "time"

type Law struct {
	ID                 uint       `json:"id" gorm:"primary_key"`
	Type               int        `json:"type" gorm:"type:int;not null"`
	Title              string     `json:"title" gorm:"type:varchar(100);not null"`
	SessionNumber      int        `json:"sessionNumber" gorm:"type:int;not null"`
	NotificationNumber string     `json:"notificationNumber" gorm:"not null"`
	NotificationDate   *time.Time `json:"notificationDate" gorm:"not null;default:now()"` // ! change default now later
	Body               string     `json:"body" gorm:"type:text;not null"`
	Image              string     `json:"image" gorm:"type:varchar(255);not null"`
	CreatedAt          *time.Time `json:"createdAt" gorm:"not null;default:now()"`
}
type LawMinimal struct {
	ID        uint       `json:"id"`
	Title     string     `json:"title"`
	Image     string     `json:"image"`
	CreatedAt *time.Time `json:"createdAt"`
}
type Comment struct {
	ID              uint   `gorm:"primary_key"`
	Body            string `gorm:"type:text;not null"`
	UserID          uint
	User            UserMigration `gorm:"foreignKey:UserID"`
	ParentCommentID uint          // `gorm:"foreignKey:UserID"`
	ParrentComment  *Comment      `gorm:"foreignKey:ParentCommentID"`
	CreatedAt       *time.Time    `gorm:"not null;default:now()"`
}
type UserMigration struct {
	// ID        *uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	ID          uint       `gorm:"primary_key"`
	Name        string     `gorm:"type:varchar(255);not null"`
	LastName    string     `gorm:"type:varchar(255);not null"`
	Username    string     `gorm:"type:varchar(255);not null"`
	PhoneNumber string     `gorm:"type:varchar(255);not null"`
	Email       string     `gorm:"type:varchar(255);not null"`
	Password    string     `gorm:"type:varchar(255);not null"`
	CreatedAt   *time.Time `gorm:"not null;default:now()"`
}

type Keyword struct {
	// ID        *uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	ID        uint   `gorm:"primary_key"`
	Body      string `gorm:"type:varchar(255);not null"`
	UserID    uint
	User      UserMigration `gorm:"foreignKey:UserID;not null"`
	CreatedAt *time.Time    `gorm:"not null;default:now()"`
}
type Attachment struct {
	// ID        *uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	ID        uint   `gorm:"primary_key"`
	FileName  string `gorm:"type:varchar(255);not null"`
	LawID     uint
	Law       *Law       `gorm:"foreignKey:LawID"`
	CreatedAt *time.Time `gorm:"not null;default:now()"`
}
type FAQ struct {
	// ID        *uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	ID           uint   `gorm:"primary_key"`
	Question     string `gorm:"type:varchar(255);not null"`
	Answer       string `gorm:"type:varchar(255);not null"`
	QuestionerID uint
	Questioner   *User `gorm:"foreignKey:QuestionerID;not null"`
	AnswererID   uint
	Answerer     *User      `gorm:"foreignKey:AnswererID;not null"`
	CreatedAt    *time.Time `gorm:"not null;default:now()"`
}
