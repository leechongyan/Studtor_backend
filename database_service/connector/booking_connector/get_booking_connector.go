package booking_connector

import (
	"errors"
	"time"

	databaseService "github.com/leechongyan/Studtor_backend/database_service/controller"
	"github.com/leechongyan/Studtor_backend/database_service/models"
)

type bookingOptions struct {
	courseId       *int
	studentId      *int
	userId         *int
	availabilityId *int
	bookingId      *int
	fromTime       time.Time
	toTime         time.Time
	err            error
}

type BookingConnector interface {
	SetCourseId(courseId int) *bookingOptions
	SetStudentId(studentId int) *bookingOptions
	SetUserId(userId int) *bookingOptions
	SetAvailabilityId(availabilityId int) *bookingOptions
	SetBookingId(bookingId int) *bookingOptions
	SetFromTime(fromTime time.Time) *bookingOptions
	SetToTime(toTime time.Time) *bookingOptions
	Add() (err error)
	Delete() (err error)
	Get() (times []models.BookingDetails, err error)
}

func Init() *bookingOptions {
	r := bookingOptions{}
	return &r
}

func (c *bookingOptions) SetCourseId(courseId int) *bookingOptions {
	c.courseId = &courseId
	return c
}

func (c *bookingOptions) SetStudentId(studentId int) *bookingOptions {
	c.studentId = &studentId
	return c
}

func (c *bookingOptions) SetUserId(userId int) *bookingOptions {
	c.userId = &userId
	return c
}

func (c *bookingOptions) SetAvailabilityId(availabilityId int) *bookingOptions {
	c.availabilityId = &availabilityId
	return c
}

func (c *bookingOptions) SetBookingId(bookingId int) *bookingOptions {
	c.bookingId = &bookingId
	return c
}

func (c *bookingOptions) SetFromTime(fromTime time.Time) *bookingOptions {
	c.fromTime = fromTime
	return c
}

func (c *bookingOptions) SetToTime(toTime time.Time) *bookingOptions {
	c.toTime = toTime
	return c
}

func (c *bookingOptions) Add() (err error) {
	if c.err != nil {
		return c.err
	}

	if c.studentId == nil {
		return errors.New("Student id must be provided")
	}

	if c.courseId == nil {
		return errors.New("Course Id must be provided")
	}

	if c.availabilityId == nil {
		return errors.New("Availability Id must be provided")
	}
	return databaseService.CurrentDatabaseConnector.SaveBooking(*c.availabilityId, *c.studentId, *c.courseId)
}

func (c *bookingOptions) Delete() (err error) {
	if c.err != nil {
		return c.err
	}

	if c.bookingId == nil {
		return errors.New("Booking Id must be provided")
	}

	if c.studentId == nil && c.userId == nil {
		return errors.New("Student Id or Tutor Id must be provided")
	}

	if c.studentId != nil {
		studentBookings, err := databaseService.CurrentDatabaseConnector.GetBookingsById(*c.studentId)
		if err == nil {
			for _, studentBooking := range studentBookings {
				if studentBooking.ID == *c.bookingId {
					return databaseService.CurrentDatabaseConnector.DeleteBookingById(*c.bookingId)
				}
			}
		}
	}

	if c.userId != nil {
		tutorBookings, err := databaseService.CurrentDatabaseConnector.GetBookingsById(*c.userId)
		if err == nil {
			for _, tutorBooking := range tutorBookings {
				if tutorBooking.ID == *c.bookingId {
					return databaseService.CurrentDatabaseConnector.DeleteBookingById(*c.bookingId)
				}
			}
		}
	}

	return errors.New("Not authorised")
}

func (c *bookingOptions) Get() (times []models.BookingDetails, err error) {
	if c.err != nil {
		return nil, c.err
	}
	if c.userId == nil {
		return nil, errors.New("User id must be provided")
	}
	if !c.fromTime.IsZero() && !c.toTime.IsZero() {
		if c.fromTime.After(c.toTime) {
			return nil, errors.New("From Time cannot be greater than or equal to To Time")
		}
		return databaseService.CurrentDatabaseConnector.GetBookingsByIdFromTo(*c.userId, c.fromTime, c.toTime)
	}
	if !c.fromTime.IsZero() {
		return databaseService.CurrentDatabaseConnector.GetBookingsByIdFrom(*c.userId, c.fromTime)
	}
	if !c.toTime.IsZero() {
		return databaseService.CurrentDatabaseConnector.GetBookingsByIdTo(*c.userId, c.toTime)
	}
	return databaseService.CurrentDatabaseConnector.GetBookingsById(*c.userId)
}
