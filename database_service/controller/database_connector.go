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

// DatabaseConnector provides the methods to interact with the models in the database
// Refer to `diagrams/studtor.drawio`, entity relationship diagram, for definitions of models.
type DatabaseConnector interface {
	// expose all the possible database api
	SaveUser(user auth_model.User) (err error)
	GetUserById(user_id int) (user auth_model.User, err error)
	GetUserByEmail(email string) (user auth_model.User, err error)
	DeleteUserById(user_id int) (err error)
	DeleteUserByEmail(email string) (err error)

	// for courses
	// GetACourse()
	// TODO: jordan reference
	// database.course + size
	// tuition.course(size)
	// return
	// GetACourse() (course, size, err)

	// GetCourses() (courses []models.Course, sizes []int, size, err error)

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

	// time u want a timeslot
	GetTimeBookId(user_id int) (times []models.Booking, err error)
	GetTimeBookIdFrom(user_id int, from_time time.Time) (times []models.Booking, err error)
	GetTimeBookIdTo(user_id int, to_time time.Time) (times []models.Booking, err error)
	GetTimeBookIdFromTo(user_id int, from_time time.Time, to_time time.Time) (times []models.Booking, err error)

	// this is for tutor
	GetTimeAvailableId(user_id int) (times []models.Availability, err error)
	GetTimeAvailableIdFrom(user_id int, from_time time.Time) (times []models.Availability, err error)
	GetTimeAvailableIdTo(user_id int, to_time time.Time) (times []models.Availability, err error)
	GetTimeAvailableIdFromTo(user_id int, from_time time.Time, to_time time.Time) (times []models.Availability, err error)

	SaveTutorAvailableTimes(user_id int, from_time time.Time, to_time time.Time) (err error)
	DeleteTutorAvailableTimes(user_id int, from_time time.Time, to_time time.Time) (err error)
	BookTutorTime(tutor_id int, student_id int, from_time time.Time, to_time time.Time) (err error)
	UnbookTutorTime(tutor_id int, student_id int, from_time time.Time, to_time time.Time) (err error)
}

func InitDatabase() (err error) {
	isMock, _ := strconv.ParseBool(viper.GetString("mock_database"))
	if isMock {
		// TODO: Chong Yan, please change the methods here for your mockdb if you'd still like to test with it
		CurrentDatabaseConnector = InitMock()
		return
	}
	// place the db that you want to instantiate here
	// commenting this out until sqlite implement the required methods

	// please remove this in future
	// sqlitedb := &SQLiteDB{}
	// sqlitedb.Init()
	// CurrentDatabaseConnector = sqlitedb

	// do this instead
	// CurrentDatabaseConnector = InitSQLite()
	return
}
