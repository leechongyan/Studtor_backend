package time_connector

import (
	"errors"
	"time"

	database_service "github.com/leechongyan/Studtor_backend/database_service/controller"
	"github.com/leechongyan/Studtor_backend/database_service/models"
)

type time_options struct {
	user_id    *int
	course_id  *int
	student_id *int
	tutor_id   *int
	from_time  time.Time
	to_time    time.Time
	isBook     bool
	err        error
}

type Get_time_connector interface {
	SetUserId(user_id int) *time_options
	SetCourseId(course_id int) *time_options
	SetStudentId(student_id int) *time_options
	SetTutorId(tutor_id int) *time_options
	SetFromTime(from_time time.Time) *time_options
	SetToTime(to_time time.Time) *time_options
	SetIsBook(isBook bool) *time_options
	Add() (err error)
	Delete() (err error)
	Get() (times []models.TimeSlot, err error)
}

func Init() *time_options {
	r := time_options{}
	return &r
}

func (c *time_options) SetUserId(user_id int) *time_options {
	c.user_id = &user_id
	return c
}

func (c *time_options) SetCourseId(course_id int) *time_options {
	c.course_id = &course_id
	return c
}

func (c *time_options) SetStudentId(student_id int) *time_options {
	c.student_id = &student_id
	return c
}

func (c *time_options) SetTutorId(tutor_id int) *time_options {
	c.tutor_id = &tutor_id
	return c
}

func (c *time_options) SetFromTime(from_time time.Time) *time_options {
	c.from_time = from_time
	return c
}

func (c *time_options) SetToTime(to_time time.Time) *time_options {
	c.to_time = to_time
	return c
}

func (c *time_options) SetIsBook(isBook bool) *time_options {
	c.isBook = isBook
	return c
}

func (c *time_options) Add() (err error) {
	if c.err != nil {
		return c.err
	}
	if c.from_time.IsZero() || c.to_time.IsZero() {
		return errors.New("Two times have to be provided")
	}
	if c.isBook {
		if c.student_id == nil || c.tutor_id == nil {
			return errors.New("Both student id and tutor id have to be provided")
		}

		if c.course_id == nil {
			return errors.New("Course Id must be provided")
		}
		return database_service.CurrentDatabaseConnector.BookTutorTime(*c.tutor_id, *c.student_id, c.from_time, c.to_time)
	}
	if c.user_id == nil {
		return errors.New("User id has to be provided")
	}
	return database_service.CurrentDatabaseConnector.SaveTutorAvailableTimes(*c.user_id, c.from_time, c.to_time)
}

func (c *time_options) Delete() (err error) {
	if c.err != nil {
		return c.err
	}

	if c.from_time.IsZero() || c.to_time.IsZero() {
		return errors.New("Two times have to be provided")
	}
	if c.isBook {
		if c.student_id == nil || c.tutor_id == nil {
			return errors.New("Both student id and tutor id have to be provided")
		}

		if c.course_id == nil {
			return errors.New("Course Id must be provided")
		}
		return database_service.CurrentDatabaseConnector.UnbookTutorTime(*c.tutor_id, *c.student_id, c.from_time, c.to_time)
	}
	if c.user_id == nil {
		return errors.New("User id has to be provided")
	}
	return database_service.CurrentDatabaseConnector.DeleteTutorAvailableTimes(*c.user_id, c.from_time, c.to_time)
}

func (c *time_options) Get() (times []models.TimeSlot, err error) {
	if c.isBook {
		// for book time
		if !c.from_time.IsZero() && !c.to_time.IsZero() {
			if c.from_time.After(c.to_time) {
				return nil, errors.New("From Time cannot be greater than or equal to To Time")
			}
			return database_service.CurrentDatabaseConnector.GetTimeBookIdFromTo(*c.user_id, c.from_time, c.to_time)
		}
		if !c.from_time.IsZero() {
			return database_service.CurrentDatabaseConnector.GetTimeBookIdFrom(*c.user_id, c.from_time)
		}
		if !c.to_time.IsZero() {
			return database_service.CurrentDatabaseConnector.GetTimeBookIdTo(*c.user_id, c.to_time)
		}
		return database_service.CurrentDatabaseConnector.GetTimeBookId(*c.user_id)
	}

	if !c.from_time.IsZero() && !c.to_time.IsZero() {
		if c.from_time.After(c.to_time) {
			return nil, errors.New("From Time cannot be greater than or equal to To Time")
		}
		return database_service.CurrentDatabaseConnector.GetTimeAvailableIdFromTo(*c.user_id, c.from_time, c.to_time)
	}
	if !c.from_time.IsZero() {
		return database_service.CurrentDatabaseConnector.GetTimeAvailableIdFrom(*c.user_id, c.from_time)
	}
	if !c.to_time.IsZero() {
		return database_service.CurrentDatabaseConnector.GetTimeAvailableIdTo(*c.user_id, c.to_time)
	}
	return database_service.CurrentDatabaseConnector.GetTimeAvailableId(*c.user_id)
}
