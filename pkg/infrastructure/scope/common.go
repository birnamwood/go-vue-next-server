package scope

import (
	"gorm.io/gorm"
)

// SetLimit 検索件数の設定
func SetLimit(limit int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		// ホーム画面は10件のみ表示
		if limit <= 0 {
			limit = -1
		}
		return db.
			Limit(limit)
	}
}

// SetOffset 検索開始点の設定
func SetOffset(offset int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if offset <= 0 {
			offset = 0
		}
		return db.
			Offset(offset)
	}
}
