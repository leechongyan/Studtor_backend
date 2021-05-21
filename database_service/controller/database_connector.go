package controller

import (
	"strconv"
	"time"

	auth_model "github.com/leechongyan/Studtor_backend/authentication_service/models"
	"github.com/leechongyan/Studtor_backend/database_service/models"
	"github.com/spf13/viper"
)

var CurrentDatabaseConnector DatabaseConnector

type Timeslots map[string]interface{}

type DatabaseConnector interface {
	// expose all the possible database api
	SaveUser(user auth_model.User) (err error)
	GetUser(email string) (user auth_model.User, err error)

	// for courses
	GetCourses() (courses []models.Course, err error)
	GetCoursesIdSize(id int, size int) (courses []models.Course, err error)
	GetCoursesId(id int) (courses []models.Course, err error)
	GetCoursesSize(size int) (courses []models.Course, err error)

	GetTutorsCourse(course_id int) (tutors []models.Tutor, err error)
	GetTutorsCourseIdSize(course_id int, tut_id int, size int) (tutors []models.Tutor, err error)
	GetTutorsCourseId(course_id int, tut_id int) (tutors []models.Tutor, err error)
	GetTutorsCourseSize(course_id int, size int) (tutors []models.Tutor, err error)

	GetTutors() (tutors []models.Tutor, err error)
	GetTutorsIdSize(tut_id int, size int) (tutors []models.Tutor, err error)
	GetTutorsId(tut_id int) (tutors []models.Tutor, err error)
	GetTutorsSize(size int) (tutors []models.Tutor, err error)

	GetTimeBookId(user_id int) (times []models.TimeSlot, err error)
	GetTimeBookIdFrom(user_id int, from_time time.Time) (times []models.TimeSlot, err error)
	GetTimeBookIdTo(user_id int, to_time time.Time) (times []models.TimeSlot, err error)
	GetTimeBookIdFromTo(user_id int, from_time time.Time, to_time time.Time) (times []models.TimeSlot, err error)

	GetTimeAvailableId(user_id int) (times []models.TimeSlot, err error)
	GetTimeAvailableIdFrom(user_id int, from_time time.Time) (times []models.TimeSlot, err error)
	GetTimeAvailableIdTo(user_id int, to_time time.Time) (times []models.TimeSlot, err error)
	GetTimeAvailableIdFromTo(user_id int, from_time time.Time, to_time time.Time) (times []models.TimeSlot, err error)

	SaveTutorAvailableTimes(user_id int, from_time time.Time, to_time time.Time) (err error)
	DeleteTutorAvailableTimes(user_id int, from_time time.Time, to_time time.Time) (err error)
	BookTutorTime(tutor_id int, student_id int, from_time time.Time, to_time time.Time) (err error)
	UnbookTutorTime(tutor_id int, student_id int, from_time time.Time, to_time time.Time) (err error)
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
