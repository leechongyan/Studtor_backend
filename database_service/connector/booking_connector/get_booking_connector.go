package booking_connector

import (
	"errors"
	"time"

	httpError "github.com/leechongyan/Studtor_backend/constants/errors/http_errors"
	clientModel "github.com/leechongyan/Studtor_backend/database_service/client_models"
	databaseService "github.com/leechongyan/Studtor_backend/database_service/controller"
	databaseError "github.com/leechongyan/Studtor_backend/database_service/errors"
	"gorm.io/gorm"
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
	GetAll() (bookings []clientModel.BookingDetails, err error)
	GetSingle() (booking clientModel.BookingDetails, err error)
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

func (c *bookingOptions) Add() (id int, err error) {
	if c.err != nil {
		return id, c.err
	}

	if c.userId == nil || c.courseId == nil || c.availabilityId == nil {
		return id, databaseError.ErrNotEnoughParameters
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
		if tutorBooking.ID == uint(*c.bookingId) {
			return databaseService.CurrentDatabaseConnector.DeleteBookingByID(*c.bookingId)
		}
	}

	return httpError.ErrUnauthorizedAccess
}

func (c *bookingOptions) GetAll() (bookings []clientModel.BookingDetails, err error) {
	if c.err != nil {
		return nil, c.err
	}
	if c.userId == nil || c.date.IsZero() || c.days == nil {
		return nil, databaseError.ErrNotEnoughParameters
	}

	books, err := databaseService.CurrentDatabaseConnector.GetBookingsByIDFromDateForSize(*c.userId, c.date, *c.days)
	if err != nil {
		return bookings, err
	}
	bookings = make([]clientModel.BookingDetails, len(books))
	for i, book := range books {
		bookings[i] = clientModel.ConvertBookingToBookingDetails(book)
	}
	return
}

func (c *bookingOptions) GetSingle() (booking clientModel.BookingDetails, err error) {
	if c.err != nil {
		return booking, c.err
	}
	if c.bookingId == nil {
		return booking, databaseError.ErrNotEnoughParameters
	}
	book, err := databaseService.CurrentDatabaseConnector.GetSingleBooking(*c.bookingId)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return booking, databaseError.ErrNoRecordFound
		}
		return booking, databaseError.ErrDatabaseInternalError
	}

	return clientModel.ConvertBookingToBookingDetails(book), err
}
