package model

import (
	"github.com/google/uuid"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type Merchant struct {
	gorm.Model
	Uuid           uuid.UUID      `json:"uuid" gorm:"type:uuid;default:uuid_generate_v4();index"`
	Name           string         `json:"name" gorm:"not null;size:80"`
	Description    string         `json:"description" gorm:"size:255"`
	IsActive       int            `json:"isActive" gorm:"size:1;default:0"`
	FrontendDomain string         `json:"frontendDomain" gorm:"size:255"`
	BackendDomain  string         `json:"backendDomain" gorm:"size:255"`
	Locale         string         `json:"locale" gorm:"size:5"`
	Services       datatypes.JSON `gorm:"type:jsonb"`
	Settings       datatypes.JSON `gorm:"type:jsonb"`
	VerifyStatus   int            `json:"verifyStatus" gorm:"size:1;default:0"`
	SupportContact string         `json:"supportContact" gorm:"size:255"`
	Users          []User         `gorm:"foreignKey:MerchantID"`
	Roles          []Role         `gorm:"foreignKey:MerchantID"`
}
