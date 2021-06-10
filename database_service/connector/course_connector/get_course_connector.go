package course_connector

import (
	"errors"

	userModel "github.com/leechongyan/Studtor_backend/database_service/client_models"
	databaseService "github.com/leechongyan/Studtor_backend/database_service/controller"
	databaseModel "github.com/leechongyan/Studtor_backend/database_service/database_models"
	databaseError "github.com/leechongyan/Studtor_backend/database_service/errors"
	"gorm.io/gorm"
)

// user can only access userModel

type courseOptions struct {
	courseId *int
	tutorId  *int
	course   *userModel.CourseWithSize
	err      error
}

type CourseConnector interface {
	SetCourseId(courseId int) *courseOptions
	SetTutorId(tutorId int) *courseOptions
	SetCourse(course databaseModel.Course) *courseOptions
	Add() (err error)
	Delete() (err error)
	GetAll() (courses []userModel.CourseWithSize, err error)
	GetSingle() (course userModel.CourseWithSize, err error)
}

func Init() *courseOptions {
	r := courseOptions{}
	return &r
}

func (c *courseOptions) SetCourseId(courseId int) *courseOptions {
	c.courseId = &courseId
	return c
}

func (c *courseOptions) SetTutorId(tutorId int) *courseOptions {
	c.tutorId = &tutorId
	return c
}

func (c *courseOptions) SetCourse(course userModel.CourseWithSize) *courseOptions {
	c.course = &course
	return c
}

func (c *courseOptions) Add() (err error) {
	if c.err != nil {
		return c.err
	}
	// add course to database
	return databaseError.ErrMethodNotImplemented
}

func (c *courseOptions) Delete() (err error) {
	if c.err != nil {
		return c.err
	}
	// delete course to database
	return databaseError.ErrMethodNotImplemented
}

func (c *courseOptions) GetSingle() (course userModel.CourseWithSize, err error) {
	if c.err != nil {
		return course, c.err
	}
	if c.courseId == nil {
		return course, databaseError.ErrNotEnoughParameters
	}
	courseWithoutSize, studentSize, tutorSize, err := databaseService.CurrentDatabaseConnector.GetCourse(*c.courseId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return course, databaseError.ErrNoRecordFound
		}
		return course, err
	}
	// convert to course with size
	return userModel.ConvertFromWithoutSizeToWithSize(courseWithoutSize, tutorSize, studentSize), err
}

func (c *courseOptions) GetAll() (courses []userModel.CourseWithSize, err error) {
	// check for error first
	if c.err != nil {
		return nil, c.err
	}
	var coursesWithoutSize []databaseModel.Course
	var tutorSizes []int
	var studentSizes []int

	if c.tutorId != nil {
		coursesWithoutSize, tutorSizes, studentSizes, err = databaseService.CurrentDatabaseConnector.GetCoursesForTutor(*c.tutorId)
	} else {
		coursesWithoutSize, tutorSizes, studentSizes, err = databaseService.CurrentDatabaseConnector.GetCourses()
	}
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return courses, databaseError.ErrNoRecordFound
		}
		return courses, err
	}
	courses = make([]userModel.CourseWithSize, len(coursesWithoutSize))
	for i, courseWithoutSize := range coursesWithoutSize {
		// convert to course with size
		courses[i] = userModel.ConvertFromWithoutSizeToWithSize(courseWithoutSize, tutorSizes[i], studentSizes[i])
	}
	return
}
