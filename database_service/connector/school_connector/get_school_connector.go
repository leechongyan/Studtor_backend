package school_connector

import (
	databaseService "github.com/leechongyan/Studtor_backend/database_service/controller"
	"github.com/leechongyan/Studtor_backend/database_service/models"
)

type schoolOptions struct {
	err error
}

type SchoolConnector interface {
	GetAll() (schools []models.CoursesForSchool, err error)
}

func Init() *schoolOptions {
	r := schoolOptions{}
	return &r
}

func (c *schoolOptions) GetAll() (schools []models.CoursesForSchool, err error) {
	// check for error first
	if c.err != nil {
		return nil, c.err
	}
	// first need to get all the school first
	schs, err := databaseService.CurrentDatabaseConnector.GetSchools()
	if err != nil {
		return
	}
	schools = make([]models.CoursesForSchool, len(schs))
	for i, sch := range schs {
		schools[i], err = databaseService.CurrentDatabaseConnector.GetCoursesForSchool(int(sch.ID))
		if err != nil {
			return
		}
	}
	return
}
