package models

import "gorm.io/gorm"

func InitTables(db *gorm.DB) error {
	return db.AutoMigrate([]interface{}{
		&UserModel{},
	})
}
