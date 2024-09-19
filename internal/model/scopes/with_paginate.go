package scopes

import "gorm.io/gorm"

func WithPaginate(page int, perPage int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		offset := (page - 1) * perPage

		return db.Offset(offset).Limit(perPage)
	}
}
