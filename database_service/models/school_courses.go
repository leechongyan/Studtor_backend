package models

import (
	"gorm.io/gorm"
)

// Follow conventions here: https://gorm.io/docs/models.html#Conventions

type SchoolCourses struct {
	gorm.Model
	ID       int
	SchoolID string
	CourseID string
}

type CoursesForSchool struct {
	ID          int
	Institution string
	SchoolCode  string
	SchoolName  string
	CourseCodes []string
	CourseNames []string
}
