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
