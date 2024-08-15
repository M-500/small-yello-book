package startup

import (
	"gin-svc/internal/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {
	db, err := gorm.Open(mysql.Open("root:root@tcp(127.0.0.1:13317)/y_book_test?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	err = models.InitTables(db)
	if err != nil {
		panic(err)
	}
	return db
}
