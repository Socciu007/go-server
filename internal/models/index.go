package models

import (
	"gorm.io/gorm"
)

// InitializeDB initializes the database and migrates all models.
func InitializeDB(db *gorm.DB) {
	// Migrate your models here
	db.AutoMigrate(
		&Users{},
		&Students{},
		&Tutors{},
		&Courses{},
		&Classes{},
		&Finances{},
	)

	// Add associations if needed
	SetupAssociations(db)
}

// setupAssociations sets up associations between models.
func SetupAssociations(db *gorm.DB) {
	// Add associations between models here
	db.Model(&Students{}).Association("Users")
	db.Model(&Tutors{}).Association("Users")
	db.Model(&Classes{}).Association("Tutors")
	db.Model(&Classes{}).Association("Courses")
	db.Model(&Finances{}).Association("Users")
}