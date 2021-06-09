package database_models

import (
	"gorm.io/gorm"
)

type Faculty struct {
	gorm.Model
	ID         uint
	SchoolName string
	Courses    []Course
	SchoolID   uint
}
