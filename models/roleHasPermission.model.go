package models

import (
	"time"
)

type RoleHasPermission struct {
	ID           uint       `gorm:"primaryKey"`
	PermissionID uint       `gorm:"not null"`
	Permission   Permission `gorm:"foreignKey:PermissionID"`
	RoleID       uint       `gorm:"not null"`
	Role         Role       `gorm:"foreignKey:RoleID"`
	CreatedAt    *time.Time `gorm:"not null;default:now()"`
	UpdatedAt    *time.Time `gorm:"not null;default:now()"`
}
