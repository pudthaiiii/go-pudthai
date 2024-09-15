package scopes

import "gorm.io/gorm"

func WithIsActive(isActive *int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if isActive != nil {
			db = db.Where("is_active = ?", *isActive)
		}

		return db
	}
}
