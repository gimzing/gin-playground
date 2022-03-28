package models

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	// database, err := gorm.Open("sqlite3", "test.db")
	dsn := "host=localhost user=lukap password=password dbname=gin port=5432 sslmode=disable TimeZone=Australia/Perth"
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database!")
	}

	database.AutoMigrate(&User{})

	DB = database
}
