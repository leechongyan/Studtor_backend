package models

import (
	"gorm.io/gorm"
)

// Follow conventions here: https://gorm.io/docs/models.html#Conventions

// Course is the model for courses ORM
// Each course can be taught by multiple tutors
type School struct {
	gorm.Model
	ID         uint `gorm:"primaryKey"`
	SchoolCode string
}

type SchoolWithSubject struct {
	gorm.Model
	ID          int
	SchoolCode  string
	CourseCodes []int
}
