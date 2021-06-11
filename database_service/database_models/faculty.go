package database_models

import (
	"gorm.io/gorm"
)

type Faculty struct {
	gorm.Model
	ID          uint
	FacultyName string
	Courses     []Course
	SchoolID    uint
}
