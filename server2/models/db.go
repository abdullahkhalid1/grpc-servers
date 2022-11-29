package models

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() error {
	db, err := gorm.Open(sqlite.Open("users.db"), &gorm.Config{})
	if err != nil {
		return err
	}
	err = db.AutoMigrate(&User{})
	if err != nil {
		return err
	}
	DB = db
	return nil
}
