package models

import(
	"time"
)

type Role struct {
    ID           uint        `gorm:"primaryKey"`
    Name         string      `gorm:"type:varchar(100);not null"`
    CreatedAt    *time.Time  `gorm:"not null;default:now()"`
    UpdatedAt    *time.Time  `gorm:"not null;default:now()"`
}
