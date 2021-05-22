package models

import (
	"database/sql"

	"gorm.io/gorm"
)

// Follow conventions here: https://gorm.io/docs/models.html#Conventions

// User is the model for user ORM
type User struct {
	gorm.Model
	ID            uint   `gorm:"primaryKey"`
	Email         string `gorm:"primaryKey"`
	FirstName     string
	LastName      string
	Password      string
	Token         sql.NullString
	UserType      string
	RefreshToken  sql.NullString
	VKey          sql.NullString
	Verified      int
	UserCreatedAt sql.NullTime
	UserUpdatedAt sql.NullTime
}
