package database_models

import (
	"time"
)

// Follow conventions here: https://gorm.io/docs/models.html#Conventions

// Booking is the model for bookings ORM
type Booking struct {
	ID                  uint `gorm:"primaryKey"`
	TutorAvailabilityID int
	TutorID             int
	UserID              int
	CourseID            int
}

// BookingDetails describes the booking for a particular TutorAvailabilityID
type BookingDetails struct {
	ID          int
	CourseCode  string
	CourseName  string
	TutorID     int
	TutorName   string
	StudentID   int
	StudentName string
	Date        time.Time
	TimeSlot    int
}
