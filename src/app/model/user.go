package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Uuid            uuid.UUID       `json:"uuid" gorm:"type:uuid;default:uuid_generate_v4();index"`
	Merchant        Merchant        `gorm:"foreignKey:MerchantID;references:ID"`
	MerchantID      uint            `json:"merchant_id" gorm:"not null;index"`
	Type            string          `json:"type" gorm:"size:20;index"`
	FullName        string          `json:"full_name" gorm:"size:80;index"`
	Email           string          `json:"email" gorm:"unique;not null;index;size:80"`
	Password        string          `json:"password" gorm:"not null;index;size:255"`
	Mobile          string          `json:"mobile" gorm:"size:20;index"`
	ProfileImage    string          `json:"profile_image" gorm:"size:255"`
	Company         string          `json:"company" gorm:"size:80"`
	IsActive        int             `json:"is_active" gorm:"size:1;default:0;index"`
	EmailVerifiedAt *gorm.DeletedAt `json:"email_verified_at"`
	Role            Role            `gorm:"foreignKey:RoleID;references:ID"`
	RoleID          uint            `json:"role_id" gorm:"index"`
	IsAllBU         int             `json:"is_all_bu" gorm:"size:1;default:0;index"`
}
