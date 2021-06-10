package tutor_connector

import (
	"errors"

	userModel "github.com/leechongyan/Studtor_backend/database_service/client_models"
	databaseService "github.com/leechongyan/Studtor_backend/database_service/controller"
	databaseModel "github.com/leechongyan/Studtor_backend/database_service/database_models"
	databaseError "github.com/leechongyan/Studtor_backend/database_service/errors"
	"gorm.io/gorm"
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
	GetAll() (tutors []userModel.UserProfile, err error)
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
		c.err = databaseError.ErrInvalidSizeParameter
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
		return databaseError.ErrNotEnoughParameters
	}
	err = databaseService.CurrentDatabaseConnector.CreateTutorCourse(*c.tutorId, *c.courseId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return databaseError.ErrNoRecordFound
		}
		return err
	}
	return
}

func (c *tutorOptions) Delete() (err error) {
	if c.err != nil {
		return c.err
	}
	if c.tutorId == nil || c.courseId == nil {
		return databaseError.ErrNotEnoughParameters
	}
	return databaseService.CurrentDatabaseConnector.DeleteTutorCourse(*c.tutorId, *c.courseId)
}

func (c *tutorOptions) GetAll() (tutors []userModel.UserProfile, err error) {
	if c.err != nil {
		return nil, c.err
	}

	if c.courseId == nil {
		return nil, databaseError.ErrNotEnoughParameters
	}
	var databaseUsers []databaseModel.User

	if c.size != nil && c.tutorId != nil {
		databaseUsers, err = databaseService.CurrentDatabaseConnector.GetTutorsForCourseFromIDOfSize(*c.courseId, *c.tutorId, *c.size)
	} else if c.size != nil {
		// get from the start
		databaseUsers, err = databaseService.CurrentDatabaseConnector.GetTutorsForCourseOfSize(*c.courseId, *c.size)
	} else if c.tutorId != nil {
		// get from from_id to the end
		databaseUsers, err = databaseService.CurrentDatabaseConnector.GetTutorsForCourseFromID(*c.courseId, *c.tutorId)
	} else {
		// get all courses
		databaseUsers, err = databaseService.CurrentDatabaseConnector.GetTutorsForCourse(*c.courseId)
	}

	tutors = make([]userModel.UserProfile, len(databaseUsers))
	for i, databaseUser := range databaseUsers {
		tutors[i] = userModel.ConvertFromDatabaseUserToUserProfile(databaseUser)
	}
	return
}
