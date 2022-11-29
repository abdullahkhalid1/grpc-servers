package models

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() error {
	db, err := gorm.Open(sqlite.Open("Cars.db"), &gorm.Config{})
	if err != nil {
		return err
	}
	err = db.AutoMigrate(&Car{})
	if err != nil {
		return err
	}
	DB = db
	return nil
}
