package config

import (
	"gorm.io/gorm"
	"gorm.io/driver/mysql"
	"os"
)

var (
	db *gorm.DB
)

func Connect() {

	var err error
	dsn := os.Getenv("DB_USERNAME") + ":" + os.Getenv("DB_PASSWORD") + "@tcp(127.0.0.1:3306)/" + os.Getenv("DB_NAME") + "?charset=utf8mb4&parseTime=True&loc=Local"

	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}
}

func GetDb() *gorm.DB {
	return db
}
