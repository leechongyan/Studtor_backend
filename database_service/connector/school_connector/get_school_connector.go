package school_connector

import (
	databaseService "github.com/leechongyan/Studtor_backend/database_service/controller"
	"github.com/leechongyan/Studtor_backend/database_service/models"
)

type schoolOptions struct {
	err error
}

type SchoolConnector interface {
	GetAll() (courses []models.CourseWithSize, err error)
}

func Init() *schoolOptions {
	r := schoolOptions{}
	return &r
}

func (c *schoolOptions) GetAll() (schools []models.SchoolWithCourses, err error) {
	// check for error first
	if c.err != nil {
		return nil, c.err
	}
	// first need to get all the school first
	schs, err := databaseService.CurrentDatabaseConnector.GetSchools()
	if err != nil {
		return
	}
	schools = make([]models.SchoolWithCourses, len(schs))
	for i, sch := range schs {
		schools[i].CourseCodes, err = databaseService.CurrentDatabaseConnector.GetCoursesCodeForSchool(sch.SchoolCode)
		if err != nil {
			return
		}
	}
	if err != nil {
		return
	}
	return
}
