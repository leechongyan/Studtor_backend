package models

import (
	"gorm.io/gorm"
)

type Course struct {
	gorm.Model
	ID          uint
	Course_ID   string
	Course_code string
	Course_name string
}

type CourseWithSize struct {
	ID           uint
	Course_ID    string
	Course_code  string
	Course_name  string
	Tutor_size   uint
	Student_size uint
}
