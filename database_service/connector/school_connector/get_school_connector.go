package school_connector

import (
	databaseService "github.com/leechongyan/Studtor_backend/database_service/controller"
	databaseModel "github.com/leechongyan/Studtor_backend/database_service/database_models"
)

type schoolOptions struct {
	err error
}

type SchoolConnector interface {
	GetAll() (schools []databaseModel.School, err error)
}

func Init() *schoolOptions {
	r := schoolOptions{}
	return &r
}

func (c *schoolOptions) GetAll() (schools []databaseModel.School, err error) {
	// check for error first
	if c.err != nil {
		return nil, c.err
	}
	// first need to get all the school first
	return databaseService.CurrentDatabaseConnector.GetSchoolsFacultiesCourses()
}
