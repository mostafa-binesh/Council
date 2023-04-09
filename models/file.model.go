package models

import (
	"time"
)

type File struct {
	ID   uint   `gorm:"primaryKey"`
	Type string `json:"type"` // like attachment, certificate and etc...
	Name string `json:"name"`
	// relations
	LawID     uint
	Law       Law        `json:"law"`
	CreatedAt *time.Time `json:"createdAt" gorm:"not null;default:now()"`
	UpdatedAt *time.Time `json:"updatedAt" gorm:"not null;default:now()"`
}
