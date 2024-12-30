package models

import (
	"time"

	"gorm.io/gorm"
)

type Finances struct {
	gorm.Model
	ID        string    `gorm:"primaryKey"`
	UserID string `gorm:"null" json:"user_id" validate:"required"`
	Type string `gorm:"null" json:"type"`
	Amount int `gorm:"null" json:"amount"`
	Status string `gorm:"null" json:"status"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}

func (Finances) TableName() string {
	return "finances"
}
