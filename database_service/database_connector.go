package database_service

import (
	"strconv"

	auth_model "github.com/leechongyan/Studtor_backend/authentication_service/models"
	tut_model "github.com/leechongyan/Studtor_backend/tuition_service/models"
	"github.com/spf13/viper"
)

var CurrentDatabaseConnector DatabaseConnector

type Timeslots map[string]interface{}

type DatabaseConnector interface {
	// expose all the possible database api
	SaveUser(user auth_model.User) (err error)
	GetUser(email string) (user auth_model.User, err error)
	GetAllCourses(from string, size int) (courses []string, err error)
	GetAllTutors(from string, size int) (tutors []string, err error)
	SaveTutorAvailableTimes(slot tut_model.Slot_query) (err error)
	GetTutorAvailableTimes(slot tut_model.Slot_query) (availableTimes Timeslots, err error)
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
	// commenting this out until sqlite implement the required methods
	// sqlitedb := &SQLiteDB{}
	// sqlitedb.Init()
	// CurrentDatabaseConnector = sqlitedb
}
