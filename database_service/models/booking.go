package models

import (
	"time"
)

// Follow conventions here: https://gorm.io/docs/models.html#Conventions

// Booking is the model for bookings ORM
type Booking struct {
	ID                  uint `gorm:"primaryKey"`
	TutorAvailabilityID int
	UserID              int
	CourseID            int
}

// TODO: Chong Yan, note that this replaces TimeSlot
// BookingDetails describes the booking for a particular TutorAvailabilityID
// Jordan you need to pass me ID as well else I cannot delete
type BookingDetails struct {
	// TODO: Previous names, for reference. Please delete before merge.
	// Course_id    uint
	// Tutor_name   string
	// Student_name string
	// From_time    time.Time
	// To_time      time.Time
	// ID is added by cy
	ID          int
	CourseCode  string
	TutorName   string
	StudentName string
	FromTime    time.Time
	ToTime      time.Time
}
