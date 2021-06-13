package database_models

import "gorm.io/gorm"

// Follow conventions here: https://gorm.io/docs/models.html#Conventions

// Booking is the model for bookings ORM
type Booking struct {
	gorm.Model
	ID             uint
	Student        User `gorm:"foreignKey:StudentID"`
	StudentID      uint
	TutorID        uint
	Availability   Availability
	AvailabilityID uint
	Course         Course
	CourseID       uint
}
