package model

import (
	"time"
)

type Movie struct {
	ID          uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	Title       string    `gorm:"type:varchar(255)" json:"title"`
	Genre       string    `gorm:"type:varchar(100)" json:"genre"`
	Duration    int       `gorm:"type:int" json:"duration"`      // Duration in minutes
	ReleaseDate time.Time `gorm:"type:date" json:"release_date"` // Release date
	Description string    `gorm:"type:text" json:"description"`
	CreatedDate time.Time `gorm:"autoCreateTime" json:"created_date"` // Automatically set on create
	CreatedBy   string    `gorm:"type:varchar(100)" json:"created_by"`
}
