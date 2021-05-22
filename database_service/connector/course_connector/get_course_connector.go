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
	SetSize(size int) *course_options
	SetFromID(from_id int) *course_options
	Add() (err error)
	Delete() (err error)
	Get() (courses []models.Course, err error)
}

func Init() *course_options {
	r := course_options{}
	return &r
}

func (c *course_options) SetSize(size int) *course_options {
	// check for size
	if size <= 0 {
		c.err = errors.New("Size cannot be 0 or negative")
	}
	c.size = &size
	return c
}

func (c *course_options) SetFromID(from_id int) *course_options {
	c.from_id = &from_id
	return c
}

func (c *course_options) Add() (err error) {
	if c.err != nil {
		return c.err
	}
	// add course to database
	return nil
}

func (c *course_options) Delete() (err error) {
	if c.err != nil {
		return c.err
	}
	// delete course to database
	return nil
}

func (c *course_options) Get() (courses []models.Course, err error) {
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
