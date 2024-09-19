package model

import "time"

type OAuthRefreshToken struct {
	ID                 uint   `gorm:"primaryKey"`
	Token              string `gorm:"unique;not null"`
	ExpiresAt          *time.Time
	OAuthAccessTokenID uint             `gorm:"index"`
	OAuthAccessToken   OAuthAccessToken `gorm:"foreignKey:OAuthAccessTokenID;references:ID"`
	CreatedAt          time.Time        `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt          time.Time        `json:"updated_at" gorm:"autoUpdateTime"`
}
