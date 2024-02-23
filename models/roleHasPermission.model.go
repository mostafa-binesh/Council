package models

import (
	"time"
)

type RoleHasPermission struct {
	ID           uint       `gorm:"primaryKey"`
	PermissionID uint       `gorm:"not null"`
	Permission   Permission `gorm:"foreignKey:PermissionID;constraint:OnUpdate:CASCADE;OnDelete:CSCADE"`
	RoleID       uint       `gorm:"not null"`
	Role         Role       `gorm:"foreignKey:RoleID;constraint:OnUpdate:CASCADE;OnDelete:CSCADE"`
	CreatedAt    *time.Time `gorm:"not null;default:now()"`
	UpdatedAt    *time.Time `gorm:"not null;default:now()"`
}
