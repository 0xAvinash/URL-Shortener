package database

import (
	"url-shortener/internal/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func connectDB() *gorm.DB {

	var db *gorm.DB
	dsn := "host=localhost user=postgres password=postgres dbname=urlservice port=5432 sslmode=disable TimeZone=Asia/Kolkata"
	var err error

	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database!")
	}

	db.AutoMigrate(&models.URL{})
	db.AutoMigrate(&models.Clicks{})

	return db
}
