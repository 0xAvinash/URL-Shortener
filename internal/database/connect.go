package database

import (
	"fmt"
	"os"
	"url-shortener/internal/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() *gorm.DB {

	var db *gorm.DB

	host := os.Getenv("DB_HOST") // "postgres"

	dsn := fmt.Sprintf(
		"host=%s user=postgres password=postgres dbname=urlservice port=5432 sslmode=disable",
		host,
	)

	var err error

	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database!")
	}

	db.AutoMigrate(&models.URL{})
	db.AutoMigrate(&models.Clicks{})

	DB = db

	return db
}
