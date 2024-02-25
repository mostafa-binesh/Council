package models

import "time"

type LawLog struct {
	ID        uint      `json:"id" gorm:"primary_key"`
	LawID     uint      `json:"lawId" gorm:"foreignKey:LawID;constraint:OnUpdate:CASCADE;OnDelete:CSCADE"`
	CreatedAt time.Time `json:"createdAt" gorm:"not null;default:now()"`
}

type UserLog struct {
	ID       uint      `json:"id" gorm:"primary_key"`
	UserID   uint      `json:"userID" gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE;OnDelete:CSCADE"`
	LoginAt  time.Time `json:"loginAt" gorm:"not null;default:now()"`
	LogoutAt time.Time `json:"logoutAt"`
}
