package database_models

import (
	"gorm.io/gorm"
)

// Follow conventions here: https://gorm.io/docs/models.html#Conventions

// School is the model for schools ORM
type School struct {
	gorm.Model
	ID          uint `gorm:"primaryKey"`
	Institution string
	SchoolCode  string
	SchoolName  string
}
