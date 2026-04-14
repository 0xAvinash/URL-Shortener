package models

import "gorm.io/gorm"

type Clicks struct {
	ID        uint64 `json:"id"`
	ShortCode string `json:"short_code"`
	IP        string `json:"ip_address"`
	UserAgent string `json:"user_agent" gorm:"type:text"`
	gorm.Model
}
