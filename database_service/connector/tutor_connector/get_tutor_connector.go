package tutor_connector

import (
	"errors"

	databaseService "github.com/leechongyan/Studtor_backend/database_service/controller"
	"github.com/leechongyan/Studtor_backend/database_service/models"
)

type tutorOptions struct {
	courseId *int
	size     *int
	tutorId  *int
	err      error
}

type TutorConnector interface {
	SetCourseId(courseId int) *tutorOptions
	SetSize(size int) *tutorOptions
	SetTutorId(tutorId int) *tutorOptions
	Add() (err error)
	Delete() (err error)
	GetAll() (tutors []models.User, err error)
}

func Init() *tutorOptions {
	r := tutorOptions{}
	return &r
}

func (c *tutorOptions) SetCourseId(courseId int) *tutorOptions {
	c.courseId = &courseId
	return c
}

func (c *tutorOptions) SetSize(size int) *tutorOptions {
	// check for size
	if size <= 0 {
		c.err = errors.New("Size cannot be 0 or negative")
	}
	c.size = &size
	return c
}

func (c *tutorOptions) SetTutorId(tutorId int) *tutorOptions {
	c.tutorId = &tutorId
	return c
}

func (c *tutorOptions) Add() (err error) {
	if c.err != nil {
		return c.err
	}
	if c.tutorId == nil || c.courseId == nil {
		return errors.New("Tutor ID and Course ID must be provided")
	}
	return databaseService.CurrentDatabaseConnector.SaveTutorCourse(*c.tutorId, *c.courseId)
}

func (c *tutorOptions) Delete() (err error) {
	if c.err != nil {
		return c.err
	}
	if c.tutorId == nil || c.courseId == nil {
		return errors.New("Tutor ID and Course ID must be provided")
	}
	return databaseService.CurrentDatabaseConnector.DeleteTutorCourse(*c.tutorId, *c.courseId)
}

func (c *tutorOptions) GetAll() (tutors []models.User, err error) {
	if c.err != nil {
		return nil, c.err
	}

	if c.courseId == nil {
		return nil, errors.New("Course ID must be provided")
	}
	if c.size != nil && c.tutorId != nil {
		return databaseService.CurrentDatabaseConnector.GetTutorsForCourseFromIdOfSize(*c.courseId, *c.tutorId, *c.size)
	}
	if c.size != nil {
		// get from the start
		return databaseService.CurrentDatabaseConnector.GetTutorsForCourseOfSize(*c.courseId, *c.size)
	}

	if c.tutorId != nil {
		// get from from_id to the end
		return databaseService.CurrentDatabaseConnector.GetTutorsForCourseFromId(*c.courseId, *c.tutorId)
	}
	// get all courses
	return databaseService.CurrentDatabaseConnector.GetTutorsForCourse(*c.courseId)
}
