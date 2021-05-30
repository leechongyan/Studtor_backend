package database_models

import (
	"gorm.io/gorm"
)

// Follow conventions here: https://gorm.io/docs/models.html#Conventions

// Course is the model for courses ORM
// Each course can be taught by multiple tutors
type Course struct {
	gorm.Model
	ID         uint `gorm:"primaryKey"`
	CourseCode string
	CourseName string
}
