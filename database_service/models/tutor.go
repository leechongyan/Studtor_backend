package models

import (
	"gorm.io/gorm"
)

// Follow conventions here: https://gorm.io/docs/models.html#Conventions

// Tutor is the model for tutors ORM
// Each tutor can teach multiple courses
type Tutor struct {
	gorm.Model
	ID        uint   `gorm:"primaryKey"`
	Email     string `gorm:"primaryKey"`
	FirstName string
	LastName  string
}
