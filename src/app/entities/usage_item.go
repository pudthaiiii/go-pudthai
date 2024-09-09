package entities

import "time"

type UsageItem struct {
	ID           int       `json:"id"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	DeletedAt    time.Time `json:"deleted_at"`
	UsagePerGrg  float64   `json:"usage_per_grg"`
	IngredientId int       `json:"ingredient_id"`
	OrgCode      string    `json:"org_code"`
	ProductCode  string    `json:"product_code"`
	UsageMonthId int       `json:"usage_month_id"`
}
