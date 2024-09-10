package model

import (
	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	Name      string `json:"name"`
	ProductId uint   `json:"product_id"`
}
