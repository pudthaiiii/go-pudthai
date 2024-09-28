package entities

import (
	"time"

	"gorm.io/gorm"
)

type OauthRefreshToken struct {
	gorm.Model
	Token              string `gorm:"unique;not null"`
	ExpiresAt          *time.Time
	OAuthAccessTokenID uint             `gorm:"index"`
	OAuthAccessToken   OauthAccessToken `gorm:"foreignKey:OAuthAccessTokenID;references:ID"`
}
