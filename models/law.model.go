package models

import "time"

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
	Comments           []Comment
	CreatedAt          time.Time `json:"createdAt" gorm:"not null;default:now()"`
	UpdatedAt          time.Time `json:"updatedAt" gorm:"not null;default:now()"`
}
type LawWithMinimalComment struct {
	ID                 uint             `json:"id" gorm:"primary_key"`
	Type               int              `json:"type" gorm:"type:int;not null"`
	Title              string           `json:"title" gorm:"type:varchar(100);not null"`
	SessionNumber      int              `json:"sessionNumber" gorm:"type:int;not null"`
	SessionDate        time.Time        `json:"sessionDate" gorm:"not null;default:now()"`      // ! change default now later
	NotificationDate   time.Time        `json:"notificationDate" gorm:"not null;default:now()"` // ! change default now later
	NotificationNumber string           `json:"notificationNumber" gorm:"not null"`
	Body               string           `json:"body" gorm:"type:text;not null"`
	Image              string           `json:"image" gorm:"type:varchar(255);not null"`
	Comments           []CommentMinimal `json:"comments"`
	CreatedAt          time.Time        `json:"createdAt" gorm:"not null;default:now()"`
	UpdatedAt          time.Time        `json:"updatedAt" gorm:"not null;default:now()"`
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
type Comment struct {
	ID              uint   `json:"id" gorm:"primary_key"`
	Body            string `json:"body" gorm:"type:text;not null"`
	UserID          uint   `json:"userID"`
	User            User   `json:"user" gorm:"foreignKey:UserID"`
	ParentCommentID uint   `json:"parentCommentID" gorm:"foreignKey:UserID"`
	// ParentComment   *Comment   `gorm:"foreignKey:ParentCommentID"`

	LawID     uint      `json:"lawID"`
	ParentLaw *Law      `json:"parentLaw" gorm:"foreignKey:LawID"`
	CreatedAt time.Time `json:"createdAt" gorm:"not null;default:now()"`
	UpdatedAt time.Time `json:"updatedAt" gorm:"not null;default:now()"`
}
type CommentMinimal struct {
	ID              uint   `json:"id"`
	Body            string `json:"body"`
	FullName        string `json:"fullName"`
	ParentCommentID uint   `json:"parentCommentID"`
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
	// ID        *uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	ID        uint   `gorm:"primary_key"`
	Body      string `gorm:"type:varchar(255);not null"`
	UserID    uint
	User      User      `gorm:"foreignKey:UserID;not null"`
	CreatedAt time.Time `gorm:"not null;default:now()"`
	UpdatedAt time.Time `json:"updatedAt" gorm:"not null;default:now()"`
}
type Attachment struct {
	// ID        *uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	ID        uint   `gorm:"primary_key"`
	FileName  string `gorm:"type:varchar(255);not null"`
	LawID     uint
	Law       *Law      `gorm:"foreignKey:LawID"`
	CreatedAt time.Time `gorm:"not null;default:now()"`
	UpdatedAt time.Time `json:"updatedAt" gorm:"not null;default:now()"`
}
type FAQ struct {
	// ID        *uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	ID           uint   `gorm:"primary_key"`
	Question     string `gorm:"type:varchar(255);not null"`
	Answer       string `gorm:"type:varchar(255);not null"`
	QuestionerID uint
	Questioner   *User `gorm:"foreignKey:QuestionerID;not null"`
	AnswererID   uint
	Answerer     *User     `gorm:"foreignKey:AnswererID;not null"`
	CreatedAt    time.Time `gorm:"not null;default:now()"`
	UpdatedAt    time.Time `json:"updatedAt" gorm:"not null;default:now()"`
}

func GetMinimalComment(comments []Comment) []CommentMinimal {
	var minimalComments []CommentMinimal
	for i := 0; i < len(comments); i++ {
		minimalComment := CommentMinimal{
			ID:              comments[i].ID,
			FullName:        comments[i].User.Name,
			ParentCommentID: comments[i].ParentCommentID,
			Body:            comments[i].Body,
		}
		minimalComments = append(minimalComments, minimalComment)

	}
	return minimalComments
}
