package scopes

import "gorm.io/gorm"

func WithSearchAndPaginate(search string, keySearch []string, page, perPage int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if search != "" {
			for _, key := range keySearch {
				db = db.Or(key+" LIKE ?", "%"+search+"%")
			}
		}

		return db.Scopes(WithPaginate(page, perPage))
	}
}
