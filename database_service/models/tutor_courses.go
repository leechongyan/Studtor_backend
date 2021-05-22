package models

import (
	"gorm.io/gorm"
)

// Follow conventions here: https://gorm.io/docs/models.html#Conventions

// TutorCourses is the model for the 'tutor_courses' is the join table
type TutorCourses struct {
	gorm.Model
	ID       uint `gorm:"primaryKey"`
	TutorID  uint
	CourseID uint
}
