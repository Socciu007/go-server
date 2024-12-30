package models

import (
	"time"

	"gorm.io/gorm"
)

type Students struct {
	gorm.Model
	ID        string    `gorm:"primaryKey"`
	FullName  string    `gorm:"null" json:"full_name" validate:"max=50"`
	ParentContact string `gorm:"null" json:"parent_contact" validate:"max=15"`
	UserID string `gorm:"null" json:"user_id" validate:"required"`
	DateOfBirth time.Time `gorm:"null" json:"date_of_birth" validate:"max=15"`
	Status string `gorm:"null" json:"status" validate:"max=15"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}

func (Students) TableName() string {
	return "students"
}
