package models

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// DB gorm define
var DB *gorm.DB

// ConnectDB is connectiong database
func ConnectDB() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&User{})

	DB = db
}
