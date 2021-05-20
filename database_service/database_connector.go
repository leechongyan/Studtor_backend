package database_service

import (
	"strconv"

	auth_model "github.com/leechongyan/Studtor_backend/authentication_service/models"
	"github.com/spf13/viper"
)

var CurrentDatabaseConnector DatabaseConnector

type Timeslots map[string]interface{}

type DatabaseConnector interface {
	// expose all the possible database api
	SaveUser(user auth_model.User) (err error)
	GetUser(email string) (user auth_model.User, err error)

	GetAllCourses(db_options options) (courses []interface{}, err error)
	GetAllTutors(db_options options) (tutors []string, err error)
	SaveTutorAvailableTimes(db_options options) (err error)
	DeleteTutorAvailableTimes(db_options options) (err error)
	GetTutorAvailableTimes(db_options options) (availableTimes Timeslots, err error)
	BookTutorTime(db_options options) (err error)
	UnBookTutorTime(db_options options) (err error)
	GetBookedTimes(db_options options) (bookedTimes Timeslots, err error)
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
	sqlitedb := &SQLiteDB{}
	sqlitedb.Init()
	CurrentDatabaseConnector = sqlitedb
    options := InitOptions()
	options.SetCourse("CZ3002").SetSize(10).GetBookedTimes()
	options.SetCourse("3002")
	
	// Service layer - service decides how much to query 
	// Service function takes in options object 
	options.SetSize(3).GetBookedTimes()
}
