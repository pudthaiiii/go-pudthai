package scopes

import "gorm.io/gorm"

func WithSearch(search string, keySearch []string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if search != "" {
			for _, key := range keySearch {
				db = db.Or(key+" LIKE ?", "%"+search+"%")
			}
		}

		return db
	}
}
