package models

import "gorm.io/gorm"

type URL struct {
	ID        uint64   `json:"id" gorm:"primaryKey"`
	ShortCode string `json:"short_code" gorm:"uniqueIndex;size:10"`
	LongURL   string `json:"long_url" gorm:"type:text"`
	gorm.Model
}
