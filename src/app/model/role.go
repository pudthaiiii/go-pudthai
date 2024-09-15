package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Role struct {
	gorm.Model
	Uuid             uuid.UUID `json:"uuid" gorm:"type:uuid;default:uuid_generate_v4();index"`
	Name             string    `json:"name" gorm:"not null;size:80"`
	Description      string    `json:"description" gorm:"null;size:80"`
	IsCorporateAdmin int       `json:"is_corporate_admin" gorm:"size:1;default:0"`
	IsActive         int       `json:"is_active" gorm:"size:1;default:0"`
	MerchantID       uint      `json:"merchant_id"`
	Merchant         Merchant  `gorm:"foreignKey:MerchantID"`
	Users            []User    `gorm:"foreignKey:RoleID"`
}
