package entities

import (
	"time"

	"gorm.io/gorm"
)

type OauthAccessToken struct {
	gorm.Model
	Token     string `gorm:"unique;not null"`
	ExpiresAt *time.Time
	UserID    uint `gorm:"index"`
	User      User `gorm:"foreignKey:UserID;references:ID"`
}
