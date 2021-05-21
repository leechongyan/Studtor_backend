package course_connector

import (
	"errors"

	database_service "github.com/leechongyan/Studtor_backend/database_service/controller"
	"github.com/leechongyan/Studtor_backend/database_service/models"
)

type course_options struct {
	size    *int
	from_id *int
	err     error
}

type Get_course_connector interface {
	PutSize(size int) *course_options
	PutFromID(from_id int) *course_options
	Exec() (courses []models.Course, err error)
}

func Init() *course_options {
	r := course_options{}
	return &r
}

func (c *course_options) PutSize(size int) *course_options {
	// check for size
	if size <= 0 {
		c.err = errors.New("Size cannot be 0 or negative")
	}
	c.size = &size
	return c
}

func (c *course_options) PutFromID(from_id int) *course_options {
	c.from_id = &from_id
	return c
}

func (c *course_options) Exec() (courses []models.Course, err error) {
	// switch statement to see which one to execute
	// check for error first
	if c.err != nil {
		return nil, c.err
	}
	if c.size != nil && c.from_id != nil {
		return database_service.CurrentDatabaseConnector.GetCoursesIdSize(*c.from_id, *c.size)
	}
	if c.size != nil {
		// get from the start
		return database_service.CurrentDatabaseConnector.GetCoursesSize(*c.size)
	}

	if c.from_id != nil {
		// get from from_id to the end
		return database_service.CurrentDatabaseConnector.GetCoursesId(*c.from_id)
	}
	// get all courses
	return database_service.CurrentDatabaseConnector.GetCourses()
}
