package tutor_connector

import (
	"errors"

	database_service "github.com/leechongyan/Studtor_backend/database_service/controller"
	"github.com/leechongyan/Studtor_backend/database_service/models"
)

type tutor_options struct {
	course_id *int
	size      *int
	from_id   *int
	err       error
}

type Get_tutor_connector interface {
	SetCourse(course_id int) *tutor_options
	SetSize(size int) *tutor_options
	SetFromID(from_id int) *tutor_options
	Add() (err error)
	Delete() (err error)
	Get() (tutors []models.Tutor, err error)
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

func (c *tutor_options) SetFromID(from_id int) *tutor_options {
	c.from_id = &from_id
	return c
}

func (c *tutor_options) Add() (err error) {
	if c.err != nil {
		return c.err
	}
	// add the tutor here connect to database
	return nil
}

func (c *tutor_options) Delete() (err error) {
	if c.err != nil {
		return c.err
	}
	// delete the tutor here connect to database
	return nil
}

func (c *tutor_options) Get() (tutors []models.Tutor, err error) {
	if c.err != nil {
		return nil, c.err
	}
	if c.err != nil {
		return nil, c.err
	}
	if c.course_id == nil {
		if c.size != nil && c.from_id != nil {
			return database_service.CurrentDatabaseConnector.GetTutorsIdSize(*c.from_id, *c.size)
		}
		if c.size != nil {
			// get from the start
			return database_service.CurrentDatabaseConnector.GetTutorsSize(*c.size)
		}

		if c.from_id != nil {
			// get from from_id to the end
			return database_service.CurrentDatabaseConnector.GetTutorsId(*c.from_id)
		}
		// get all courses
		return database_service.CurrentDatabaseConnector.GetTutors()
	}

	if c.size != nil && c.from_id != nil {
		return database_service.CurrentDatabaseConnector.GetTutorsCourseIdSize(*c.course_id, *c.from_id, *c.size)
	}
	if c.size != nil {
		// get from the start
		return database_service.CurrentDatabaseConnector.GetTutorsCourseSize(*c.course_id, *c.size)
	}

	if c.from_id != nil {
		// get from from_id to the end
		return database_service.CurrentDatabaseConnector.GetTutorsCourseId(*c.course_id, *c.from_id)
	}
	// get all courses
	return database_service.CurrentDatabaseConnector.GetTutorsCourse(*c.course_id)
}
