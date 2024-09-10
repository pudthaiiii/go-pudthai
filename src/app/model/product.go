package model

import (
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	ProductName string `json:"ProductName"`
}
