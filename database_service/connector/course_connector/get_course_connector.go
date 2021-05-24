package course_connector

import (
	"errors"

	databaseService "github.com/leechongyan/Studtor_backend/database_service/controller"
	"github.com/leechongyan/Studtor_backend/database_service/models"
)

type courseOptions struct {
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

type CourseConnector interface {
	SetSize(size int) *courseOptions
	SetCourseCode(courseCode string) *courseOptions
	SetSchoolCode(schoolCode string) *courseOptions
	SetCourseId(courseId int) *courseOptions
	SetTutorId(tutorId int) *courseOptions
	SetCourse(course models.Course) *courseOptions
	SetCourseWithSize(course models.CourseWithSize) *courseOptions
	Add() (err error)
	Delete() (err error)
	GetAll() (courses []models.CourseWithSize, err error)
	GetSingle() (course models.CourseWithSize, err error)
}

func Init() *courseOptions {
	r := courseOptions{}
	return &r
}

func (c *courseOptions) SetSize(size int) *courseOptions {
	// check for size
	if size <= 0 {
		c.err = errors.New("Size cannot be 0 or negative")
	}
	c.size = &size
	return c
}

func (c *courseOptions) SetOffset(offset int) *courseOptions {
	// check for offset
	if offset <= 0 {
		c.err = errors.New("Size cannot be 0 or negative")
	}
	c.offset = &offset
	return c
}

func (c *courseOptions) SetSchoolCode(schoolCode string) *courseOptions {
	c.schoolCode = &schoolCode
	return c
}

func (c *courseOptions) SetCourseCode(courseCode string) *courseOptions {
	c.courseCode = &courseCode
	return c
}

func (c *courseOptions) SetCourseId(courseId int) *courseOptions {
	c.courseId = &courseId
	return c
}

func (c *courseOptions) SetTutorId(tutorId int) *courseOptions {
	c.tutorId = &tutorId
	return c
}

func (c *courseOptions) SetCourse(course models.Course) *courseOptions {
	if c.courseWithSize != nil {
		c.err = errors.New("Course With Size has already been set")
		return c
	}
	c.course = &course
	return c
}

func (c *courseOptions) SetCourseWithSize(course models.CourseWithSize) *courseOptions {
	if c.course != nil {
		c.err = errors.New("Course has already been set")
		return c
	}
	c.courseWithSize = &course
	return c
}

func (c *courseOptions) Add() (err error) {
	if c.err != nil {
		return c.err
	}
	// add course to database
	return errors.New("Not implemented")
}

func (c *courseOptions) Delete() (err error) {
	if c.err != nil {
		return c.err
	}
	// delete course to database
	return errors.New("Not implemented")
}

func (c *courseOptions) GetSingle() (course models.CourseWithSize, err error) {
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

func (c *courseOptions) GetAll() (courses []models.CourseWithSize, err error) {
	// check for error first
	if c.err != nil {
		return nil, c.err
	}
	var coursesWithoutSize []models.Course
	var tutorSizes []int
	var studentSizes []int
	if c.schoolCode != nil {
		if c.size != nil && c.offset != nil {
			coursesWithoutSize, tutorSizes, studentSizes, err = databaseService.CurrentDatabaseConnector.GetCoursesForSchoolOfSizeWithOffset(*c.schoolCode, *c.offset, *c.size)
		} else if c.size != nil {
			// get for school from start to size x
			coursesWithoutSize, tutorSizes, studentSizes, err = databaseService.CurrentDatabaseConnector.GetCoursesForSchoolOfSize(*c.schoolCode, *c.size)
		} else if c.offset != nil {
			// get for school from offset to the end
			coursesWithoutSize, tutorSizes, studentSizes, err = databaseService.CurrentDatabaseConnector.GetCoursesForSchoolWithOffset(*c.schoolCode, *c.offset)
		} else {
			// get for school
			coursesWithoutSize, tutorSizes, studentSizes, err = databaseService.CurrentDatabaseConnector.GetCoursesForSchool(*c.schoolCode)
		}
	} else {
		if c.size != nil && c.offset != nil {
			coursesWithoutSize, tutorSizes, studentSizes, err = databaseService.CurrentDatabaseConnector.GetCoursesOfSizeWithOffset(*c.offset, *c.size)
		} else if c.size != nil {
			coursesWithoutSize, tutorSizes, studentSizes, err = databaseService.CurrentDatabaseConnector.GetCoursesOfSize(*c.size)
		} else if c.offset != nil {
			coursesWithoutSize, tutorSizes, studentSizes, err = databaseService.CurrentDatabaseConnector.GetCoursesWithOffset(*c.offset)
		} else {
			coursesWithoutSize, tutorSizes, studentSizes, err = databaseService.CurrentDatabaseConnector.GetCourses()
		}
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
