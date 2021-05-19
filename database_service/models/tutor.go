package models

import (
	"gorm.io/gorm"
)

type Tutor struct {
	gorm.Model
	ID         uint
	Tutor_name string
}
