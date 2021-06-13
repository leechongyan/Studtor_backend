package booking_connector

import (
	"errors"
	"time"

	clientModel "github.com/leechongyan/Studtor_backend/database_service/client_models"
	databaseService "github.com/leechongyan/Studtor_backend/database_service/controller"
	databaseModel "github.com/leechongyan/Studtor_backend/database_service/database_models"
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
	isTutor        bool
}

type BookingConnector interface {
	SetCourseId(courseId int) *bookingOptions
	SetUserId(userId int) *bookingOptions
	SetAvailabilityId(availabilityId int) *bookingOptions
	SetBookingId(bookingId int) *bookingOptions
	SetDate(date time.Time) *bookingOptions
	SetTimeId(timeId int) *bookingOptions
	SetDays(days int) *bookingOptions
	SetIsTutor(isTutor bool) *bookingOptions
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

func (c *bookingOptions) SetIsTutor(isTutor bool) *bookingOptions {
	c.isTutor = isTutor
	return c
}

func (c *bookingOptions) Add() (id int, err error) {
	if c.err != nil {
		return id, c.err
	}

	if c.userId == nil || c.courseId == nil || c.availabilityId == nil {
		return id, databaseError.ErrNotEnoughParameters
	}

	// check whether availability is valid
	err = isValidBooking(*c.availabilityId, *c.courseId)
	if err != nil {
		return id, err
	}

	return databaseService.CurrentDatabaseConnector.CreateBooking(*c.availabilityId, *c.userId, *c.courseId)
}

func isValidBooking(availabilityId int, courseId int) error {
	avail, err := databaseService.CurrentDatabaseConnector.GetSingleAvailability(availabilityId)
	if err != nil {
		return err
	}
	if avail.Occupied {
		// if is occupied then return error
		return databaseError.ErrInvalidAvailability
	}

	courses := avail.Tutor.Courses
	for _, course := range courses {
		if int(course.ID) == courseId {
			// if tutor teach the course, then return no error
			return nil
		}
	}
	return databaseError.ErrInvalidBooking

}

func (c *bookingOptions) Delete() (err error) {
	if c.err != nil {
		return c.err
	}

	if c.bookingId == nil || c.userId == nil {
		return databaseError.ErrNotEnoughParameters
	}

	err = isValidUnbooking(*c.userId, *c.bookingId)
	if err != nil {
		return err
	}

	return databaseService.CurrentDatabaseConnector.DeleteBookingByID(*c.bookingId)
}

func isValidUnbooking(userId int, bookingId int) error {
	tutorBookings, err := databaseService.CurrentDatabaseConnector.GetBookingsForTutorByID(userId)
	if err != nil {
		return err
	}

	for _, tutorBooking := range tutorBookings {
		if tutorBooking.ID == uint(bookingId) {
			return nil
		}
	}

	studentBookings, err := databaseService.CurrentDatabaseConnector.GetBookingsForStudentByID(userId)
	if err != nil {
		return err
	}

	for _, studentBooking := range studentBookings {
		if studentBooking.ID == uint(bookingId) {
			return nil
		}
	}

	return databaseError.ErrUnauthorizedAccess
}

func (c *bookingOptions) GetAll() (bookings []clientModel.BookingDetails, err error) {
	if c.err != nil {
		return nil, c.err
	}
	if c.userId == nil {
		return nil, databaseError.ErrNotEnoughParameters
	}
	var books []databaseModel.Booking
	if c.date.IsZero() || c.days == nil {
		if c.isTutor {
			books, err = databaseService.CurrentDatabaseConnector.GetBookingsForTutorByID(*c.userId)
		} else {
			books, err = databaseService.CurrentDatabaseConnector.GetBookingsForStudentByID(*c.userId)
		}
	} else {
		if c.isTutor {
			books, err = databaseService.CurrentDatabaseConnector.GetBookingsForTutorByIDFromDateForSize(*c.userId, c.date, *c.days)
		} else {
			books, err = databaseService.CurrentDatabaseConnector.GetBookingsForStudentByIDFromDateForSize(*c.userId, c.date, *c.days)
		}
	}

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
		return booking, err
	}

	return clientModel.ConvertBookingToBookingDetails(book), err
}
