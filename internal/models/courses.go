package models

import (
	"time"

	"gorm.io/gorm"
)

type Courses struct {
	gorm.Model
	ID        string    `gorm:"primaryKey"`
	Name string `gorm:"null" json:"name"`
	Description string `gorm:"null" json:"description"`
	Fee int `gorm:"null" json:"fee"`
	Duration int `gorm:"null" json:"duration"`
	Status string `gorm:"null" json:"status"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}

func (Courses) TableName() string {
	return "courses"
}