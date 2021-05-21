package time_connector

import (
	"errors"
	"time"

	database_operation "github.com/leechongyan/Studtor_backend/database_service/constants"
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
	op         database_operation.Operation
}

type Get_time_connector interface {
	PutUserId(user_id int) *time_options
	PutCourseId(course_id int) *time_options
	PutStudentId(student_id int) *time_options
	PutTutorId(tutor_id int) *time_options
	PutFromTime(from_time time.Time) *time_options
	PutToTime(to_time time.Time) *time_options
	SetIsBook(isBook bool) *time_options
	SetOperation(op database_operation.Operation) *time_options
	Exec() (times []models.TimeSlot, err error)
}

func Init() *time_options {
	r := time_options{}
	return &r
}

func (c *time_options) SetOperation(op database_operation.Operation) *time_options {
	c.op = op
	return c
}

func (c *time_options) PutUserId(user_id int) *time_options {
	c.user_id = &user_id
	return c
}

func (c *time_options) PutCourseId(course_id int) *time_options {
	c.course_id = &course_id
	return c
}

func (c *time_options) PutStudentId(student_id int) *time_options {
	c.student_id = &student_id
	return c
}

func (c *time_options) PutTutorId(tutor_id int) *time_options {
	c.tutor_id = &tutor_id
	return c
}

func (c *time_options) PutFromTime(from_time time.Time) *time_options {
	c.from_time = from_time
	return c
}

func (c *time_options) PutToTime(to_time time.Time) *time_options {
	c.to_time = to_time
	return c
}

func (c *time_options) SetIsBook(isBook bool) *time_options {
	c.isBook = isBook
	return c
}

func (c *time_options) Exec() (times []models.TimeSlot, err error) {
	if c.err != nil {
		return nil, c.err
	}
	switch c.op {
	case database_operation.Get:
		{
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
	case database_operation.Add:
		{
			if c.from_time.IsZero() || c.to_time.IsZero() {
				return nil, errors.New("Two times have to be provided")
			}
			if c.isBook {
				if c.student_id == nil || c.tutor_id == nil {
					return nil, errors.New("Both student id and tutor id have to be provided")
				}

				if c.course_id == nil {
					return nil, errors.New("Course Id must be provided")
				}
				err := database_service.CurrentDatabaseConnector.BookTutorTime(*c.tutor_id, *c.student_id, c.from_time, c.to_time)
				return nil, err
			}
			if c.user_id == nil {
				return nil, errors.New("User id has to be provided")
			}
			err := database_service.CurrentDatabaseConnector.SaveTutorAvailableTimes(*c.user_id, c.from_time, c.to_time)
			return nil, err
		}
	case database_operation.Delete:
		{
			if c.from_time.IsZero() || c.to_time.IsZero() {
				return nil, errors.New("Two times have to be provided")
			}
			if c.isBook {
				if c.student_id == nil || c.tutor_id == nil {
					return nil, errors.New("Both student id and tutor id have to be provided")
				}

				if c.course_id == nil {
					return nil, errors.New("Course Id must be provided")
				}
				err := database_service.CurrentDatabaseConnector.UnbookTutorTime(*c.tutor_id, *c.student_id, c.from_time, c.to_time)
				return nil, err
			}
			if c.user_id == nil {
				return nil, errors.New("User id has to be provided")
			}
			err := database_service.CurrentDatabaseConnector.DeleteTutorAvailableTimes(*c.user_id, c.from_time, c.to_time)
			return nil, err
		}
	default:
		{
			return nil, errors.New("No operation specified")
		}
	}
}
