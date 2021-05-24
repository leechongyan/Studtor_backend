package tutor_connector

import (
	"errors"

	databaseService "github.com/leechongyan/Studtor_backend/database_service/controller"
	"github.com/leechongyan/Studtor_backend/database_service/models"
)

type tutorOptions struct {
	courseId   *int
	size       *int
	tutorId    *int
	tutorEmail *string
	tutor      *models.Tutor
	err        error
}

type TutorConnector interface {
	SetCourse(courseId int) *tutorOptions
	SetSize(size int) *tutorOptions
	SetTutorId(tutorId int) *tutorOptions
	SetTutorEmail(tutorEmail string) *tutorOptions
	SetTutor(tutor models.Tutor) *tutorOptions
	Add() (err error)
	Delete() (err error)
	GetAll() (tutors []models.Tutor, err error)
	GetSingle() (tutor models.Tutor, err error)
}

func Init() *tutorOptions {
	r := tutorOptions{}
	return &r
}

func (c *tutorOptions) SetCourse(courseId int) *tutorOptions {
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

func (c *tutorOptions) SetTutorEmail(tutorEmail string) *tutorOptions {
	c.tutorEmail = &tutorEmail
	return c
}

func (c *tutorOptions) SetTutor(tutor models.Tutor) *tutorOptions {
	c.tutor = &tutor
	return c
}

func (c *tutorOptions) Add() (err error) {
	if c.err != nil {
		return c.err
	}
	// add the tutor here connect to database
	if c.tutor != nil {
		return databaseService.CurrentDatabaseConnector.SaveTutor(*c.tutor)
	}
	return errors.New("Tutor object must be provided")
}

func (c *tutorOptions) Delete() (err error) {
	if c.err != nil {
		return c.err
	}
	if c.tutorEmail != nil {
		return databaseService.CurrentDatabaseConnector.DeleteTutorByEmail(*c.tutorEmail)
	}
	if c.tutorId != nil {
		return databaseService.CurrentDatabaseConnector.DeleteTutorById(*c.tutorId)
	}
	return errors.New("Tutor email or id must be provided")
}

func (c *tutorOptions) GetSingle() (tutor models.Tutor, err error) {
	if c.err != nil {
		return models.Tutor{}, c.err
	}
	if c.tutorEmail != nil {
		return databaseService.CurrentDatabaseConnector.GetTutorByEmail(*c.tutorEmail)
	}
	if c.tutorId != nil {
		return databaseService.CurrentDatabaseConnector.GetTutorById(*c.tutorId)
	}
	return models.Tutor{}, errors.New("Tutor email or id must be provided")
}

func (c *tutorOptions) GetAll() (tutors []models.Tutor, err error) {
	if c.err != nil {
		return nil, c.err
	}
	if c.err != nil {
		return nil, c.err
	}
	if c.courseId == nil {
		// get all tutors regardless of the course they teach
		if c.size != nil && c.tutorId != nil {
			return databaseService.CurrentDatabaseConnector.GetTutorsFromIdOfSize(*c.tutorId, *c.size)
		}
		if c.size != nil {
			// get from the start
			return databaseService.CurrentDatabaseConnector.GetTutorsOfSize(*c.size)
		}

		if c.tutorId != nil {
			// get from from_id to the end
			return databaseService.CurrentDatabaseConnector.GetTutorsFromId(*c.tutorId)
		}
		// get all courses
		return databaseService.CurrentDatabaseConnector.GetTutors()
	}
	// get all tutors for a course
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
