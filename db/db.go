package db

import (
	"gin-boiler-plate/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func Init() {
	connectDB()
}

func GetDB() *gorm.DB {
	return db
}

func connectDB() {
	dsn := "root:tkdwlstn@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	db.AutoMigrate(&model.User{})
	if err != nil {
		panic(err)
	}
}
