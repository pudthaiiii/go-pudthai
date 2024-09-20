package entities

import "time"

type OAuthAccessToken struct {
	ID        uint   `gorm:"primaryKey"`
	Token     string `gorm:"unique;not null"`
	ExpiresAt *time.Time
	UserID    uint      `gorm:"index"`
	User      User      `gorm:"foreignKey:UserID;references:ID"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}
