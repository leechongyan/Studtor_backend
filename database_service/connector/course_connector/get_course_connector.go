package course_connector

import (
	"errors"

	database_service "github.com/leechongyan/Studtor_backend/database_service/controller"
	"github.com/leechongyan/Studtor_backend/database_service/models"
)

type course_options struct {
	size             *int
	course_code      *string
	course_id        *int
	tutor_id         *int
	course           *models.Course
	course_with_size *models.CourseWithSize
	err              error
}

type Get_course_connector interface {
	SetSize(size int) *course_options
	SetCourseCode(course_code string) *course_options
	SetCourseId(course_id int) *course_options
	SetTutorId(tutor_id int) *course_options
	SetCourse(course models.Course) *course_options
	SetCourseWithSize(course models.CourseWithSize) *course_options
	Add() (err error)
	Delete() (err error)
	GetAll() (courses []models.CourseWithSize, err error)
	GetSingle() (course models.CourseWithSize, err error)
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

func (c *course_options) SetCourseCode(course_code string) *course_options {
	c.course_code = &course_code
	return c
}

func (c *course_options) SetCourseId(course_id int) *course_options {
	c.course_id = &course_id
	return c
}

func (c *course_options) SetTutorId(tutor_id int) *course_options {
	c.tutor_id = &tutor_id
	return c
}

func (c *course_options) SetCourse(course models.Course) *course_options {
	if c.course_with_size != nil {
		c.err = errors.New("Course With Size has already been set")
		return c
	}
	c.course = &course
	return c
}

func (c *course_options) SetCourseWithSize(course models.CourseWithSize) *course_options {
	if c.course != nil {
		c.err = errors.New("Course has already been set")
		return c
	}
	c.course_with_size = &course
	return c
}

func (c *course_options) Add() (err error) {
	if c.err != nil {
		return c.err
	}
	// add course to database
	return errors.New("Not implemented")
}

func (c *course_options) Delete() (err error) {
	if c.err != nil {
		return c.err
	}
	// delete course to database
	return errors.New("Not implemented")
}

func (c *course_options) GetSingle() (course models.CourseWithSize, err error) {
	if c.err != nil {
		return models.CourseWithSize{}, c.err
	}
	if c.course_id == nil {
		return models.CourseWithSize{}, errors.New("Course id must be provided")
	}
	course_without_size, student_size, tutor_size, err := database_service.CurrentDatabaseConnector.GetCourse(*c.course_id)
	if err != nil {
		return models.CourseWithSize{}, err
	}
	// convert to course with size
	course = models.CourseWithSize{}
	course.ID = int(course_without_size.ID)
	course.CourseCode = course_without_size.CourseCode
	course.CourseName = course_without_size.CourseName
	course.StudentSize = student_size
	course.TutorSize = tutor_size
	return course, nil
}

func (c *course_options) GetAll() (courses []models.CourseWithSize, err error) {
	// check for error first
	if c.err != nil {
		return nil, c.err
	}
	var courses_without_size []models.Course
	var tutor_sizes []int
	var student_sizes []int

	if c.size != nil && c.course_code != nil {
		courses_without_size, tutor_sizes, student_sizes, err = database_service.CurrentDatabaseConnector.GetCoursesFromIdOfSize(*c.course_code, *c.size)
	} else if c.size != nil {
		// get from the start
		courses_without_size, tutor_sizes, student_sizes, err = database_service.CurrentDatabaseConnector.GetCoursesOfSize(*c.size)
	} else if c.course_code != nil {
		// get from from_id to the end
		courses_without_size, tutor_sizes, student_sizes, err = database_service.CurrentDatabaseConnector.GetCoursesFromId(*c.course_code)
	} else {
		courses_without_size, tutor_sizes, student_sizes, err = database_service.CurrentDatabaseConnector.GetCourses()
	}
	if err != nil {
		return nil, err
	}
	courses = make([]models.CourseWithSize, len(courses_without_size))
	for i, cws := range courses_without_size {
		// convert to course with size
		courses[i] = models.CourseWithSize{}
		courses[i].CourseCode = cws.CourseCode
		courses[i].CourseName = cws.CourseName
		courses[i].ID = int(cws.ID)
		courses[i].StudentSize = student_sizes[i]
		courses[i].TutorSize = tutor_sizes[i]
	}
	return courses, nil
}
