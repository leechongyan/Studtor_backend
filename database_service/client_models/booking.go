package client_models

import (
	"fmt"
	"time"

	databaseModel "github.com/leechongyan/Studtor_backend/database_service/database_models"
)

// BookingDetails describes the booking for a particular TutorAvailabilityID
type BookingDetails struct {
	ID           int
	CourseCode   string
	CourseName   string
	TutorID      int
	TutorName    string
	TutorEmail   string
	StudentID    int
	StudentName  string
	StudentEmail string
	Date         time.Time
	TimeSlot     int
}

func ConvertBookingToBookingDetails(book databaseModel.Booking) (booking BookingDetails) {
	// get availability
	availability := book.Availability
	course := book.Course
	booking = BookingDetails{}
	booking.ID = int(book.ID)
	booking.CourseCode = course.CourseCode
	booking.CourseName = course.CourseName
	booking.TutorID = int(availability.TutorID)
	booking.TutorName = availability.Tutor.Name
	booking.TutorEmail = availability.Tutor.Email
	booking.StudentID = int(book.StudentID)
	booking.StudentName = book.Student.Name
	booking.StudentEmail = book.Student.Email
	booking.Date = availability.Date
	booking.TimeSlot = availability.TimeSlot
	fmt.Println("HIT AT CONNECTOR")
	return
}
