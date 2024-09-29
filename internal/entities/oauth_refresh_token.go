package entities

import (
	"time"

	"gorm.io/gorm"
)

type OauthRefreshToken struct {
	gorm.Model
	Token              string `gorm:"unique;not null"`
	ExpiresAt          *time.Time
	OauthAccessTokenID uint             `gorm:"index"`
	OauthAccessToken   OauthAccessToken `gorm:"foreignKey:OauthAccessTokenID;references:ID"`
}
