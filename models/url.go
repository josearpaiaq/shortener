package models

import "gorm.io/gorm"

type URL struct {
	gorm.Model

	OriginalURL string `gorm:"type:varchar(255);uniqueIndex" json:"original_url"`
	ShortURL    string `gorm:"type:varchar(255);uniqueIndex" json:"short_url"`
	Clicks      int    `gorm:"default:0" json:"clicks"`
}