package models

import "gorm.io/gorm"

type URL struct {
	ID        uint   `json:"id" gorm:"primaryKey"`
	ShortCode string `json:"short_code" gorm:"uniqueIndex"`
	LongURL   string `json:"long_url"`
	gorm.Model
}
