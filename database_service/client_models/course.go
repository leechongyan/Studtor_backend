package client_models

import (
	databaseModel "github.com/leechongyan/Studtor_backend/database_service/database_models"
)

type CourseWithSize struct {
	ID          int
	CourseCode  string
	CourseName  string
	TutorSize   int
	StudentSize int
}

func ConvertFromWithoutSizeToWithSize(courseWithoutSize databaseModel.Course, tutorSize int, studentSize int) (courseWithSize CourseWithSize) {
	courseWithSize.CourseCode = courseWithoutSize.CourseCode
	courseWithSize.CourseName = courseWithoutSize.CourseName
	courseWithSize.ID = int(courseWithoutSize.ID)
	courseWithSize.StudentSize = studentSize
	courseWithSize.TutorSize = tutorSize
	return
}
