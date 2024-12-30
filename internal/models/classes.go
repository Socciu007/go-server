package models

import (
	"time"

	"gorm.io/gorm"
)

type Classes struct {
	gorm.Model
	ID        string    `gorm:"primaryKey"`
	TutorID string `gorm:"null" json:"tutor_id" validate:"required"`
	CourseID string `gorm:"null" json:"course_id" validate:"required"`
	Schedule string `gorm:"null" json:"schedule"`
	Status string `gorm:"null" json:"status" validate:"max=15"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}

func (Classes) TableName() string {
	return "classes"
}
