package course_connector

import (
	"errors"

	databaseService "github.com/leechongyan/Studtor_backend/database_service/controller"
	"github.com/leechongyan/Studtor_backend/database_service/models"
)

type course_options struct {
	offset         *int
	size           *int
	courseCode     *string
	schoolCode     *string
	courseId       *int
	tutorId        *int
	course         *models.Course
	courseWithSize *models.CourseWithSize
	err            error
}

type Get_course_connector interface {
	SetSize(size int) *course_options
	SetCourseCode(courseCode string) *course_options
	SetSchoolCode(schoolCode string) *course_options
	SetCourseId(courseId int) *course_options
	SetTutorId(tutorId int) *course_options
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

func (c *course_options) SetOffset(offset int) *course_options {
	// check for offset
	if offset <= 0 {
		c.err = errors.New("Size cannot be 0 or negative")
	}
	c.offset = &offset
	return c
}

func (c *course_options) SetSchoolCode(schoolCode string) *course_options {
	c.schoolCode = &schoolCode
	return c
}

func (c *course_options) SetCourseCode(courseCode string) *course_options {
	c.courseCode = &courseCode
	return c
}

func (c *course_options) SetCourseId(courseId int) *course_options {
	c.courseId = &courseId
	return c
}

func (c *course_options) SetTutorId(tutorId int) *course_options {
	c.tutorId = &tutorId
	return c
}

func (c *course_options) SetCourse(course models.Course) *course_options {
	if c.courseWithSize != nil {
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
	c.courseWithSize = &course
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
	if c.courseId == nil {
		return models.CourseWithSize{}, errors.New("Course id must be provided")
	}
	courseWithoutSize, studentSize, tutorSize, err := databaseService.CurrentDatabaseConnector.GetCourse(*c.courseId)
	if err != nil {
		return models.CourseWithSize{}, err
	}
	// convert to course with size
	course = models.CourseWithSize{}
	course.ID = int(courseWithoutSize.ID)
	course.CourseCode = courseWithoutSize.CourseCode
	course.CourseName = courseWithoutSize.CourseName
	course.StudentSize = studentSize
	course.TutorSize = tutorSize
	return course, nil
}

func (c *course_options) GetAll() (courses []models.CourseWithSize, err error) {
	// check for error first
	if c.err != nil {
		return nil, c.err
	}
	var coursesWithoutSize []models.Course
	var tutorSizes []int
	var studentSizes []int

	if c.size != nil && c.schoolCode != nil && c.offset != nil {
		coursesWithoutSize, tutorSizes, studentSizes, err = databaseService.CurrentDatabaseConnector.GetCoursesForSchoolOfSizeWithOffset(*c.schoolCode, *c.offset, *c.size)
	} else if c.size != nil && c.schoolCode != nil {
		// get for school from start to size x
		coursesWithoutSize, tutorSizes, studentSizes, err = databaseService.CurrentDatabaseConnector.GetCoursesForSchoolOfSize(*c.schoolCode, *c.size)
	} else if c.offset != nil && c.schoolCode != nil {
		// get for school from offset to the end
		coursesWithoutSize, tutorSizes, studentSizes, err = databaseService.CurrentDatabaseConnector.GetCoursesForSchoolWithOffset(*c.schoolCode, *c.offset)
	} else if c.schoolCode != nil {
		// get for school
		coursesWithoutSize, tutorSizes, studentSizes, err = databaseService.CurrentDatabaseConnector.GetCoursesForSchool(*c.schoolCode)
	} else {
		coursesWithoutSize, tutorSizes, studentSizes, err = databaseService.CurrentDatabaseConnector.GetCourses()
	}
	if err != nil {
		return nil, err
	}
	courses = make([]models.CourseWithSize, len(coursesWithoutSize))
	for i, courseWithoutSize := range coursesWithoutSize {
		// convert to course with size
		courses[i] = models.CourseWithSize{}
		courses[i].CourseCode = courseWithoutSize.CourseCode
		courses[i].CourseName = courseWithoutSize.CourseName
		courses[i].ID = int(courseWithoutSize.ID)
		courses[i].StudentSize = studentSizes[i]
		courses[i].TutorSize = tutorSizes[i]
	}
	return courses, nil
}
