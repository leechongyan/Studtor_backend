package tutor_connector

import (
	"errors"

	database_service "github.com/leechongyan/Studtor_backend/database_service/controller"
	"github.com/leechongyan/Studtor_backend/database_service/models"
)

type tutor_options struct {
	course_id *int

	size    *int
	from_id *int
	err     error
}

type Get_tutor_connector interface {
	PutCourse(course_id int) *tutor_options
	PutSize(size int) *tutor_options
	PutFromID(from_id int) *tutor_options
	Exec() (tutors []models.Tutor, err error)
}

func Init() *tutor_options {
	r := tutor_options{}
	return &r
}

func (c *tutor_options) PutCourse(course_id int) *tutor_options {
	c.course_id = &course_id
	return c
}

func (c *tutor_options) PutSize(size int) *tutor_options {
	// check for size
	if size <= 0 {
		c.err = errors.New("Size cannot be 0 or negative")
	}
	c.size = &size
	return c
}

func (c *tutor_options) PutFromID(from_id int) *tutor_options {
	c.from_id = &from_id
	return c
}

func (c *tutor_options) Exec() (tutors []models.Tutor, err error) {
	// switch statement to see which one to execute
	// check for error first
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
