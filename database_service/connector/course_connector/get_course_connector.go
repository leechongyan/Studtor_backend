package course_connector

import (
	databaseError "github.com/leechongyan/Studtor_backend/constants/errors/database_errors"
	userModel "github.com/leechongyan/Studtor_backend/database_service/client_models"
	databaseService "github.com/leechongyan/Studtor_backend/database_service/controller"
	databaseModel "github.com/leechongyan/Studtor_backend/database_service/database_models"
)

type courseOptions struct {
	courseId *int
	tutorId  *int
	course   *databaseModel.Course
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

func (c *courseOptions) SetCourse(course databaseModel.Course) *courseOptions {
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
		return userModel.CourseWithSize{}, c.err
	}
	if c.courseId == nil {
		return userModel.CourseWithSize{}, databaseError.ErrNotEnoughParameters
	}
	courseWithoutSize, studentSize, tutorSize, err := databaseService.CurrentDatabaseConnector.GetCourse(*c.courseId)
	if err != nil {
		return userModel.CourseWithSize{}, err
	}
	// convert to course with size
	course = convertFromCourseToCourseWithSize(courseWithoutSize, studentSize, tutorSize)
	return
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
		return nil, err
	}
	courses = make([]userModel.CourseWithSize, len(coursesWithoutSize))
	for i, courseWithoutSize := range coursesWithoutSize {
		// convert to course with size
		courses[i] = convertFromCourseToCourseWithSize(courseWithoutSize, studentSizes[i], tutorSizes[i])
	}
	return
}

func convertFromCourseToCourseWithSize(courseWithoutSize databaseModel.Course, studentSize int, tutorSize int) (courseWithSize userModel.CourseWithSize) {
	courseWithSize = userModel.CourseWithSize{}
	courseWithSize.CourseCode = courseWithoutSize.CourseCode
	courseWithSize.CourseName = courseWithoutSize.CourseName
	courseWithSize.ID = int(courseWithoutSize.ID)
	courseWithSize.StudentSize = studentSize
	courseWithSize.TutorSize = tutorSize
	return
}
