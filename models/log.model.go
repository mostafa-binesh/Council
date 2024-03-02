package models

import (
	D "docker/database"
	"github.com/gofiber/fiber/v2"
	"time"
)

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

type ActionLog struct {
	ID          uint      `json:"id" gorm:"primary_key"`
	IP          string    `json:"ip"`
	HostName    string    `json:"hostName"`
	UserID      uint      `json:"userID" gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE;OnDelete:CSCADE"`
	Port        string    `josn:"port"`
	Url         string    `json:"url"`
	RequestType string    `json:"requestType"`
	RouteName   string    `json:"routeName"`
	ActionTime  time.Time `json:"actionTime" gorm:"not null;default:now()"`
}

func GetLog(c *fiber.Ctx) bool {
	user := c.Locals("user").(User)
	ip := c.Get("X-Real-IP")
	if ip == "" {
		ip = c.Get("X-Forwarded-For")
	}
	if ip == "" {
		ip = c.IP()
	}
	log := ActionLog{
		IP:          ip,
		HostName:    c.Hostname(),
		Port:        c.Port(),
		Url:         c.BaseURL(),
		UserID:      user.ID,
		RequestType: c.Method(),
		RouteName:   c.Route().Path,
	}
	resultLog := D.DB().Create(&log)
	if resultLog.Error != nil {
		return false
	}
	return true
}
