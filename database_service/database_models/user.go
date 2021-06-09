package database_models

import (
	"database/sql"
	"time"

	"gorm.io/gorm"
)

// Follow conventions here: https://gorm.io/docs/models.html#Conventions

// User is the model for user ORM
type User struct {
	gorm.Model
	ID             uint
	Name           string
	Email          string
	Password       string
	Token          sql.NullString
	UserType       string
	RefreshToken   sql.NullString
	VKey           sql.NullString
	ProfilePicture sql.NullString
	Verified       bool
	UserCreatedAt  time.Time
	UserUpdatedAt  time.Time
	Availabilities []Availability `gorm:"foreignKey:TutorID"`
	Bookings       []Booking      `gorm:"foreignKey:StudentID"`
	Courses        []Course       `gorm:"many2many:tutor_course;"`
}
