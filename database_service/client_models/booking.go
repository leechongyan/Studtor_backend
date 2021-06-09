package client_models

import (
	"time"

	databaseModel "github.com/leechongyan/Studtor_backend/database_service/database_models"
)

// BookingDetails describes the booking for a particular TutorAvailabilityID
type BookingDetails struct {
	id           int
	courseCode   string
	courseName   string
	tutorID      int
	tutorName    string
	tutorEmail   string
	studentID    int
	studentName  string
	studentEmail string
	date         time.Time
	timeSlot     int
}

func (booking BookingDetails) CourseCode() string {
	return booking.courseCode
}

func (booking BookingDetails) CourseName() string {
	return booking.courseName
}

func (booking BookingDetails) TutorName() string {
	return booking.tutorName
}

func (booking BookingDetails) TutorEmail() string {
	return booking.tutorEmail
}

func (booking BookingDetails) StudentName() string {
	return booking.studentName
}

func (booking BookingDetails) StudentEmail() string {
	return booking.studentEmail
}

func (booking BookingDetails) Date() time.Time {
	return booking.date
}

func (booking BookingDetails) TimeSlot() int {
	return booking.timeSlot
}

func ConvertBookingToBookingDetails(book databaseModel.Booking) (booking BookingDetails) {
	// get availability
	availability := book.Availability
	course := book.Course
	booking = BookingDetails{}
	booking.id = int(book.ID)
	booking.courseCode = course.CourseCode
	booking.courseName = course.CourseName
	booking.tutorID = int(availability.TutorID)
	booking.tutorName = availability.Tutor.Name
	booking.tutorEmail = availability.Tutor.Email
	booking.studentID = int(book.StudentID)
	booking.studentName = book.Student.Name
	booking.studentEmail = book.Student.Email
	booking.date = availability.Date
	booking.timeSlot = availability.TimeSlot
	return
}
