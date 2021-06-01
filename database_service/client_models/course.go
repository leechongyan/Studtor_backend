package client_models

import (
	databaseModel "github.com/leechongyan/Studtor_backend/database_service/database_models"
)

type CourseWithSize struct {
	id          int
	courseCode  string
	courseName  string
	tutorSize   int
	studentSize int
}

func (course *CourseWithSize) ID() int {
	return course.id
}

func (course *CourseWithSize) CourseName() string {
	return course.courseName
}

func ConvertFromWithoutSizeToWithSize(courseWithoutSize databaseModel.Course, tutorSize int, studentSize int) (coureWithSize CourseWithSize) {
	courseWithSize := CourseWithSize{}
	courseWithSize.courseCode = courseWithoutSize.CourseCode
	courseWithSize.courseName = courseWithoutSize.CourseName
	courseWithSize.id = int(courseWithoutSize.ID)
	courseWithSize.studentSize = studentSize
	courseWithSize.tutorSize = tutorSize
	return
}
