package database_service

import (
	"strconv"

	"github.com/leechongyan/Studtor_backend/authentication_service/models"
	"github.com/leechongyan/Studtor_backend/helpers"
	"github.com/spf13/viper"
)

var CurrentDatabaseConnector DatabaseConnector

type DatabaseConnector interface {
	// expose all the possible database api
	SaveUser(user models.User) (err *helpers.RequestError)
	GetUser(email string) (user models.User, err *helpers.RequestError)
	GetAllCourses(from string, size int) (courses []string, err *helpers.RequestError)
	GetAllTutors(from string, size int) (tutors []string, err *helpers.RequestError)
}

func InitDatabase() {
	isMock, _ := strconv.ParseBool(viper.GetString("mock_database"))
	if isMock {
		mdb := &Mockdb{}
		mdb.Init()
		CurrentDatabaseConnector = mdb
		return
	}
	// place the db that you want to instantiate here
}
