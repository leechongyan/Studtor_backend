package controller

import (
	"strconv"
	"time"

	// "github.com/leechongyan/Studtor_backend/authentication_service/db_model"
	db_model "github.com/leechongyan/Studtor_backend/database_service/models"
	"github.com/spf13/viper"
)

var CurrentDatabaseConnector DatabaseConnector

type Timeslots map[string]interface{}

// DatabaseConnector provIDes the methods to interact with the db_model in the database
// Refer to `diagrams/studtor.drawio`, entity relationship diagram, for definitions of db_model.
type DatabaseConnector interface {
	/*
		Users model
	*/

	// GetUsers retrieves a list of all user model objects from the database.
	GetUsers() (users []db_model.User, err error)
	// GetUserByID retrieves a user model object by the user's ID from the database.
	GetUserByID(userID int) (user db_model.User, err error)
	// GetUserByEmail retrieves a user model object by the user's email from the database.
	GetUserByEmail(email string) (user db_model.User, err error)
	// CreateUser saves an user object into the database.
	CreateUser(user db_model.User) (err error)
	// UpdateUser updates an user object into the database.
	UpdateUser(user db_model.User) (err error)
	// DeleteUserByID deletes an user object by the user's ID from the database.
	DeleteUserByID(userID int) (err error)
	// DeleteUserByEmail deletes an user object by the user's email from the database.
	DeleteUserByEmail(email string) (err error)

	/*
		TutorCourses model
	*/

	// GetCoursesForTutor retrieves a list of all courses that a tutor is teaching from the database.
	GetCoursesForTutor(tutorID int) (courses []db_model.Course, nStudents []int, nTutors []int, err error)
	// GetTutorsForCourse retrieves a list of all tutors for a particular course from the database.
	GetTutorsForCourse(courseID int) (tutors []db_model.User, err error)
	// GetTutorsForCourseFromIDOfSize retrieves a list of tutors for a particular course from the database,
	// starting from tut_ID to tut_ID + size
	GetTutorsForCourseFromIDOfSize(courseID int, tutorID int, size int) (tutors []db_model.User, err error)
	// GetTutorsForCourseFromID retrieves a list of tutors for a particular course from the database,
	// starting from tut_ID to the end
	GetTutorsForCourseFromID(courseID int, tut_ID int) (tutors []db_model.User, err error)
	// GetTutorsForCourseFromID retrieves a list of tutors for a particular course from the database,
	// starting from 0 to size
	GetTutorsForCourseOfSize(courseID int, size int) (tutors []db_model.User, err error)
	// CreateTutorCourse saves a tutor_course model object into the database.
	// This function is called when a tutor registers interest to teach a course.
	CreateTutorCourse(tutorID int, courseID int) (err error)
	// DeleteTutorCourse deletes an tutor course object from the database.
	DeleteTutorCourse(tutorID int, courseID int) (err error)

	/*
		Courses model
	*/

	// GetCourses retrieves a course along with the number of students enrolled in the course
	// and the number of tutors for the course, from the database.
	GetCourse(courseID int) (course db_model.Course, nStudents int, nTutors int, err error)
	// GetCourses retrieves a list of all courses, along with the number of students
	// enrolled in the course and the number of tutors for the course, from the database.
	// Sorted by course code.
	GetCourses() (courses []db_model.Course, nStudents []int, nTutors []int, err error)

	/*
		Schools model
	*/

	// GetSchools retrieves the list of schools from the database
	GetSchools() (schools []db_model.School, err error)
	// TODO (Jordan GetCoursesIdForSchool gets a list of course code attached the the school)
	GetCoursesForSchool(school_id int) (courses db_model.SchoolCoursesDetails, err error)
	// GetSchoolByInstitutionAndSchoolCode retrieves a school from the database
	GetSchoolByInstitutionAndSchoolCode(institution string, schoolCode string) (school db_model.School, err error)

	/*
		SchoolCourses model
	*/

	// GetCoursesForSchoolByID retrieves a list of course codes attached to a school by school courses ID

	// TODO: Jordan issue #27
	// GetCoursesForSchoolByID(schoolCoursesID string) (schoolCourses []db_model.SchoolCoursesDetails, err error)

	/*
		Booking model
	*/

	// GetBookingsByID retrieves a list of all bookings by a user, as indicated by userID, with no time constraints
	GetBookingsByID(userID int) (bookings []db_model.BookingDetails, err error)
	// GetBookingsByIDFrom retrieves a list of all bookings by a user, as indicated by userID, starting from time fromTime
	GetBookingsByIDFrom(userID int, fromTime time.Time) (bookings []db_model.BookingDetails, err error)
	// GetBookingsByID retrieves a list of all bookings by a user, as indicated by userID, ending with time toTime
	GetBookingsByIDTo(userID int, toTime time.Time) (bookings []db_model.BookingDetails, err error)
	// GetBookingsByID retrieves a list of all bookings by a user, as indicated by userID,
	// starting from time fromTime and ending with time toTime
	GetBookingsByIDFromTo(userID int, fromTime time.Time, toTime time.Time) (bookings []db_model.BookingDetails, err error)
	// CreateBooking saves a booking model object into the database
	CreateBooking(availabilityID int, userID int, courseID int) (err error)
	// DeleteBooking deletes a booking model object into the database
	DeleteBookingByID(bookingID int) (err error)

	/*
		TutorAvailability model
	*/

	// GetAvailabilityByID retrieves a list of all available timeslots for a tutor, with no time constraints
	GetAvailabilityByID(tutorID int) (availabilities []db_model.Availability, err error)
	// GetAvailabilityByID retrieves a list of all available timeslots for a tutor, starting from time fromTime
	GetAvailabilityByIDFrom(tutorID int, fromTime time.Time) (availabilities []db_model.Availability, err error)
	// GetAvailabilityByID retrieves a list of all available timeslots for a tutor, ending with time toTime
	GetAvailabilityByIDTo(tutorID int, toTime time.Time) (availabilities []db_model.Availability, err error)
	// GetAvailabilityByID retrieves a list of all available timeslots for a tutor,
	// starting from time fromTime and ending with time toTime
	GetAvailabilityByIDFromTo(tutorID int, fromTime time.Time, toTime time.Time) (availabilities []db_model.Availability, err error)
	// CreateTutorAvailability saves a tutor availability model object into the database
	CreateTutorAvailability(tutorID int, fromTime time.Time, toTime time.Time) (err error)
	// DeleteTutorAvailability deletes a tutor availability model object by ID from the database
	DeleteTutorAvailabilityByID(availabilityID int) (err error)
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
	// sqlitedb := &SQLiteDB{}
	// sqlitedb.Init()
	// CurrentDatabaseConnector = sqlitedb
	return
}
