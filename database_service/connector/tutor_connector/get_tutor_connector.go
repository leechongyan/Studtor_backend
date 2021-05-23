package tutor_connector

import (
	"errors"

	database_service "github.com/leechongyan/Studtor_backend/database_service/controller"
	"github.com/leechongyan/Studtor_backend/database_service/models"
)

type tutor_options struct {
	course_id   *int
	size        *int
	tutor_id    *int
	tutor_email *string
	tutor       *models.Tutor
	err         error
}

type Get_tutor_connector interface {
	SetCourse(course_id int) *tutor_options
	SetSize(size int) *tutor_options
	SetTutorId(tutor_id int) *tutor_options
	SetTutorEmail(tutor_email string) *tutor_options
	SetTutor(tutor models.Tutor) *tutor_options
	Add() (err error)
	Delete() (err error)
	GetAll() (tutors []models.Tutor, err error)
	GetSingle() (tutor models.Tutor, err error)
}

func Init() *tutor_options {
	r := tutor_options{}
	return &r
}

func (c *tutor_options) SetCourse(course_id int) *tutor_options {
	c.course_id = &course_id
	return c
}

func (c *tutor_options) SetSize(size int) *tutor_options {
	// check for size
	if size <= 0 {
		c.err = errors.New("Size cannot be 0 or negative")
	}
	c.size = &size
	return c
}

func (c *tutor_options) SetTutorId(tutor_id int) *tutor_options {
	c.tutor_id = &tutor_id
	return c
}

func (c *tutor_options) SetTutorEmail(tutor_email string) *tutor_options {
	c.tutor_email = &tutor_email
	return c
}

func (c *tutor_options) SetTutor(tutor models.Tutor) *tutor_options {
	c.tutor = &tutor
	return c
}

func (c *tutor_options) Add() (err error) {
	if c.err != nil {
		return c.err
	}
	// add the tutor here connect to database
	if c.tutor != nil {
		return database_service.CurrentDatabaseConnector.SaveTutor(*c.tutor)
	}
	return errors.New("Tutor object must be provided")
}

func (c *tutor_options) Delete() (err error) {
	if c.err != nil {
		return c.err
	}
	if c.tutor_email != nil {
		return database_service.CurrentDatabaseConnector.DeleteTutorByEmail(*c.tutor_email)
	}
	if c.tutor_id != nil {
		return database_service.CurrentDatabaseConnector.DeleteTutorById(*c.tutor_id)
	}
	return errors.New("Tutor email or id must be provided")
}

func (c *tutor_options) GetSingle() (tutor models.Tutor, err error) {
	if c.err != nil {
		return models.Tutor{}, c.err
	}
	if c.tutor_email != nil {
		return database_service.CurrentDatabaseConnector.GetTutorByEmail(*c.tutor_email)
	}
	if c.tutor_id != nil {
		return database_service.CurrentDatabaseConnector.GetTutorById(*c.tutor_id)
	}
	return models.Tutor{}, errors.New("Tutor email or id must be provided")
}

func (c *tutor_options) GetAll() (tutors []models.Tutor, err error) {
	if c.err != nil {
		return nil, c.err
	}
	if c.err != nil {
		return nil, c.err
	}
	if c.course_id == nil {
		// get all tutors regardless of the course they teach
		if c.size != nil && c.tutor_id != nil {
			return database_service.CurrentDatabaseConnector.GetTutorsFromIdOfSize(*c.tutor_id, *c.size)
		}
		if c.size != nil {
			// get from the start
			return database_service.CurrentDatabaseConnector.GetTutorsOfSize(*c.size)
		}

		if c.tutor_id != nil {
			// get from from_id to the end
			return database_service.CurrentDatabaseConnector.GetTutorsFromId(*c.tutor_id)
		}
		// get all courses
		return database_service.CurrentDatabaseConnector.GetTutors()
	}
	// get all tutors for a course
	if c.size != nil && c.tutor_id != nil {
		return database_service.CurrentDatabaseConnector.GetTutorsForCourseFromIdOfSize(*c.course_id, *c.tutor_id, *c.size)
	}
	if c.size != nil {
		// get from the start
		return database_service.CurrentDatabaseConnector.GetTutorsForCourseOfSize(*c.course_id, *c.size)
	}

	if c.tutor_id != nil {
		// get from from_id to the end
		return database_service.CurrentDatabaseConnector.GetTutorsForCourseFromId(*c.course_id, *c.tutor_id)
	}
	// get all courses
	return database_service.CurrentDatabaseConnector.GetTutorsForCourse(*c.course_id)
}
