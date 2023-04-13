package models

import "time"

type GuestMessage struct {
	ID        int        `json:"id" gorm:"primaryKey"`
	GuestID   int        `json:"guestID" gorm:"uniqueIndex"`
	Guest     Guest      `json:"guest" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Sender    int8       `json:"fullName"` // 1: guest, 2: admin
	Body      string     `json:"body"`
	CreatedAt *time.Time `json:"createdAt" gorm:"not null;default:now()"`
}
