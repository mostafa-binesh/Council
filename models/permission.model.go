package models

import (
	"time"
)

type Permission struct {
	ID        uint       `gorm:"primaryKey"`
	Name      string     `gorm:"type:varchar(100);not null"`
	RandomID  string     `gorm:"type:varchar(100);not null"`
	CreatedAt *time.Time `gorm:"not null;default:now()"`
	UpdatedAt *time.Time `gorm:"not null;default:now()"`
}

type PermissionGet struct {
    Name string `json:"Name"`
}