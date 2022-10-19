package config

import (
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
)

func Connect() {
	connection, err := gorm.Open(sqlite.Open("database/database.sqlite"), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		panic("Failed to connect to databse due to previous error")
	}
	DB = connection
}

func GetDB() *gorm.DB {
	return DB
}
