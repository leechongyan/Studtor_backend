package booking_connector

import (
	"errors"
	"time"

	database_service "github.com/leechongyan/Studtor_backend/database_service/controller"
	"github.com/leechongyan/Studtor_backend/database_service/models"
)

type booking_options struct {
	course_id       *int
	student_id      *int
	user_id         *int
	availability_id *int
	booking_id      *int
	from_time       time.Time
	to_time         time.Time
	err             error
}

type Get_booking_connector interface {
	SetCourseId(course_id int) *booking_options
	SetStudentId(student_id int) *booking_options
	SetUserId(user_id int) *booking_options
	SetAvailabilityId(availability_id int) *booking_options
	SetBookingId(booking_id int) *booking_options
	SetFromTime(from_time time.Time) *booking_options
	SetToTime(to_time time.Time) *booking_options
	Add() (err error)
	Delete() (err error)
	Get() (times []models.BookingDetails, err error)
}

func Init() *booking_options {
	r := booking_options{}
	return &r
}

func (c *booking_options) SetCourseId(course_id int) *booking_options {
	c.course_id = &course_id
	return c
}

func (c *booking_options) SetStudentId(student_id int) *booking_options {
	c.student_id = &student_id
	return c
}

func (c *booking_options) SetUserId(user_id int) *booking_options {
	c.user_id = &user_id
	return c
}

func (c *booking_options) SetAvailabilityId(availability_id int) *booking_options {
	c.availability_id = &availability_id
	return c
}

func (c *booking_options) SetBookingId(booking_id int) *booking_options {
	c.booking_id = &booking_id
	return c
}

func (c *booking_options) SetFromTime(from_time time.Time) *booking_options {
	c.from_time = from_time
	return c
}

func (c *booking_options) SetToTime(to_time time.Time) *booking_options {
	c.to_time = to_time
	return c
}

func (c *booking_options) Add() (err error) {
	if c.err != nil {
		return c.err
	}

	if c.student_id == nil {
		return errors.New("Student id must be provided")
	}

	if c.course_id == nil {
		return errors.New("Course Id must be provided")
	}

	if c.availability_id == nil {
		return errors.New("Availability Id must be provided")
	}
	return database_service.CurrentDatabaseConnector.SaveBooking(*c.availability_id, *c.student_id, *c.course_id)
}

func (c *booking_options) Delete() (err error) {
	if c.err != nil {
		return c.err
	}

	if c.booking_id == nil {
		return errors.New("Booking Id must be provided")
	}
	return database_service.CurrentDatabaseConnector.DeleteBookingById(*c.booking_id)
}

func (c *booking_options) Get() (times []models.BookingDetails, err error) {
	if c.err != nil {
		return nil, c.err
	}
	if c.user_id == nil {
		return nil, errors.New("User id must be provided")
	}
	if !c.from_time.IsZero() && !c.to_time.IsZero() {
		if c.from_time.After(c.to_time) {
			return nil, errors.New("From Time cannot be greater than or equal to To Time")
		}
		return database_service.CurrentDatabaseConnector.GetBookingsByIdFromTo(*c.user_id, c.from_time, c.to_time)
	}
	if !c.from_time.IsZero() {
		return database_service.CurrentDatabaseConnector.GetBookingsByIdFrom(*c.user_id, c.from_time)
	}
	if !c.to_time.IsZero() {
		return database_service.CurrentDatabaseConnector.GetBookingsByIdTo(*c.user_id, c.to_time)
	}
	return database_service.CurrentDatabaseConnector.GetBookingsById(*c.user_id)
}
