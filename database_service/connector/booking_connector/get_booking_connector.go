package booking_connector

import (
	"time"

	databaseError "github.com/leechongyan/Studtor_backend/constants/errors/database_errors"
	httpError "github.com/leechongyan/Studtor_backend/constants/errors/http_errors"
	databaseService "github.com/leechongyan/Studtor_backend/database_service/controller"
	"github.com/leechongyan/Studtor_backend/database_service/models"
)

type bookingOptions struct {
	courseId       *int
	userId         *int
	availabilityId *int
	bookingId      *int
	fromTime       time.Time
	toTime         time.Time
	err            error
}

type BookingConnector interface {
	SetCourseId(courseId int) *bookingOptions
	SetUserId(userId int) *bookingOptions
	SetAvailabilityId(availabilityId int) *bookingOptions
	SetBookingId(bookingId int) *bookingOptions
	SetFromTime(fromTime time.Time) *bookingOptions
	SetToTime(toTime time.Time) *bookingOptions
	Add() (err error)
	Delete() (err error)
	GetAll() (times []models.BookingDetails, err error)
}

func Init() *bookingOptions {
	r := bookingOptions{}
	return &r
}

func (c *bookingOptions) SetCourseId(courseId int) *bookingOptions {
	c.courseId = &courseId
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

	if c.userId == nil || c.courseId == nil || c.availabilityId == nil {
		return databaseError.ErrNotEnoughParameters
	}

	return databaseService.CurrentDatabaseConnector.CreateBooking(*c.availabilityId, *c.userId, *c.courseId)
}

func (c *bookingOptions) Delete() (err error) {
	if c.err != nil {
		return c.err
	}

	if c.bookingId == nil || c.userId == nil {
		return databaseError.ErrNotEnoughParameters
	}

	tutorBookings, err := databaseService.CurrentDatabaseConnector.GetBookingsByID(*c.userId)
	if err != nil {
		return
	}

	for _, tutorBooking := range tutorBookings {
		if tutorBooking.ID == *c.bookingId {
			return databaseService.CurrentDatabaseConnector.DeleteBookingByID(*c.bookingId)
		}
	}

	return httpError.ErrUnauthorizedAccess
}

func (c *bookingOptions) GetAll() (times []models.BookingDetails, err error) {
	if c.err != nil {
		return nil, c.err
	}
	if c.userId == nil {
		return nil, databaseError.ErrNotEnoughParameters
	}
	if !c.fromTime.IsZero() && !c.toTime.IsZero() {
		if c.fromTime.After(c.toTime) {
			return nil, databaseError.ErrInvalidTimes
		}
		return databaseService.CurrentDatabaseConnector.GetBookingsByIDFromTo(*c.userId, c.fromTime, c.toTime)
	}
	if !c.fromTime.IsZero() {
		return databaseService.CurrentDatabaseConnector.GetBookingsByIDFrom(*c.userId, c.fromTime)
	}
	if !c.toTime.IsZero() {
		return databaseService.CurrentDatabaseConnector.GetBookingsByIDTo(*c.userId, c.toTime)
	}
	return databaseService.CurrentDatabaseConnector.GetBookingsByID(*c.userId)
}
