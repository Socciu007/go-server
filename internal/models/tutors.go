package models

import (
	"time"

	"gorm.io/gorm"
)

type Tutors struct {
	gorm.Model
	ID        string    `gorm:"primaryKey"`
	Qualification string `gorm:"null" json:"qualification"`
	Experience string `gorm:"null" json:"experience"`
	Subject string `gorm:"null" json:"subject" validate:"max=50"`
	UserID string `gorm:"null" json:"user_id" validate:"required"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}

func (Tutors) TableName() string {
	return "tutors"
}
