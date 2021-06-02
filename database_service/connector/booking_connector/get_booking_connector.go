package booking_connector

import (
	"time"

	databaseError "github.com/leechongyan/Studtor_backend/constants/errors/database_errors"
	httpError "github.com/leechongyan/Studtor_backend/constants/errors/http_errors"
	databaseService "github.com/leechongyan/Studtor_backend/database_service/controller"
	databaseModel "github.com/leechongyan/Studtor_backend/database_service/database_models"
)

type bookingOptions struct {
	courseId       *int
	userId         *int
	availabilityId *int
	bookingId      *int
	date           time.Time
	days           *int
	err            error
}

type BookingConnector interface {
	SetCourseId(courseId int) *bookingOptions
	SetUserId(userId int) *bookingOptions
	SetAvailabilityId(availabilityId int) *bookingOptions
	SetBookingId(bookingId int) *bookingOptions
	SetDate(date time.Time) *bookingOptions
	SetTimeId(timeId int) *bookingOptions
	SetDays(days int) *bookingOptions
	Add() (err error)
	Delete() (err error)
	GetAll() (bookings []databaseModel.BookingDetails, err error)
	GetSingle() (booking databaseModel.BookingDetails, err error)
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

func (c *bookingOptions) SetDate(date time.Time) *bookingOptions {
	c.date = date
	return c
}

func (c *bookingOptions) SetDays(days int) *bookingOptions {
	c.days = &days
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

func (c *bookingOptions) GetAll() (bookings []databaseModel.BookingDetails, err error) {
	if c.err != nil {
		return nil, c.err
	}
	if c.userId == nil || c.date.IsZero() || c.days == nil {
		return nil, databaseError.ErrNotEnoughParameters
	}

	return databaseService.CurrentDatabaseConnector.GetBookingsByIDFromDateForSize(*c.userId, c.date, *c.days)

}

func (c *bookingOptions) GetSingle() (booking databaseModel.BookingDetails, err error) {
	if c.err != nil {
		return databaseModel.BookingDetails{}, c.err
	}
	if c.bookingId == nil {
		return databaseModel.BookingDetails{}, databaseError.ErrNotEnoughParameters
	}
	return databaseService.CurrentDatabaseConnector.GetSingleBooking(*c.bookingId)
}
