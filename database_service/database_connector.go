package database_service

import (
	"strconv"

	"github.com/leechongyan/Studtor_backend/authentication_service/models"
	"github.com/spf13/viper"
)

var CurrentDatabaseConnector DatabaseConnector

type DatabaseConnector interface {
	// expose all the possible database api
	SaveUser(user models.User) (err error)
	GetUser(email string) (user models.User, err error)
	GetAllCourses(from string, size int) (courses []string, err error)
	GetAllTutors(from string, size int) (tutors []string, err error)
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
	sqlitedb := &SQLiteDB{}
	sqlitedb.Init()
	CurrentDatabaseConnector = sqlitedb
	return
}
