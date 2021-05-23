package controller

import (
	"strconv"
	"time"

	// auth_model "github.com/leechongyan/Studtor_backend/authentication_service/models"
	"github.com/leechongyan/Studtor_backend/database_service/models"
	"github.com/spf13/viper"
)

var CurrentDatabaseConnector DatabaseConnector

type Timeslots map[string]interface{}

// DatabaseConnector provides the methods to interact with the models in the database
// Refer to `diagrams/studtor.drawio`, entity relationship diagram, for definitions of models.
type DatabaseConnector interface {
	/*
		Users model
	*/

	// GetUsers retrieves a list of all user model objects from the database.
	GetUsers() (user []models.User, err error)
	// GetUserById retrieves a user model object by the user's id from the database.
	GetUserById(user_id int) (user models.User, err error)
	// GetUserByEmail retrieves a user model object by the user's email from the database.
	GetUserByEmail(email string) (user models.User, err error)
	// SaveUser saves an auth_model user object into the database.
	SaveUser(user models.User) (err error)
	// DeleteUserById deletes an auth_model user object by the user's id from the database.
	DeleteUserById(user_id int) (err error)
	// DeleteUserByEmail deletes an auth_model user object by the user's email from the database.
	DeleteUserByEmail(email string) (err error)

	/*
		TutorCourses model
	*/

	// GetCoursesForTutor retrieves a list of all courses that a tutor is teaching from the database.
	GetCoursesForTutor(tutor_id int) (course []models.Course, err error)
	// GetTutorsForCourse retrieves a list of all tutors for a particular course from the database.
	GetTutorsForCourse(course_id int) (tutor []models.Tutor, err error)
	// GetTutorsForCourseFromIdOfSize retrieves a list of tutors for a particular course from the database,
	// starting from tut_id to tut_id + size
	GetTutorsForCourseFromIdOfSize(course_id, tut_id int, size int) (tutor []models.Tutor, err error)
	// GetTutorsForCourseFromId retrieves a list of tutors for a particular course from the database,
	// starting from tut_id to the end
	GetTutorsForCourseFromId(course_id, tut_id int) (tutor []models.Tutor, err error)
	// GetTutorsForCourseFromId retrieves a list of tutors for a particular course from the database,
	// starting from 0 to size
	GetTutorsForCourseOfSize(course_id, size int) (tutor []models.Tutor, err error)
	// SaveTutorCourse saves a tutor_course model object into the database.
	// This function is called when a tutor registers interest to teach a course.
	SaveTutorCourse(tutor_id int, course_id int) (err error)

	/*
		Courses model
	*/

	// GetCourses retrieves a course along with the number of students enrolled in the course
	// and the number of tutors for the course, from the database.
	GetCourse(course_id int) (course models.Course, n_students int, n_tutors int, err error)
	// GetCourses retrieves a list of all courses, along with the number of students
	// enrolled in the course and the number of tutors for the course, from the database.
	// Sorted by course code.
	GetCourses() (courses []models.Course, n_students []int, n_tutors []int, err error)
	// GetCoursesFromId retrieves the list of courses from this course code to the end.
	// Sorted by course code.
	GetCoursesFromId(course_code string) (courses []models.Course, n_students []int, n_tutors []int, err error)
	// GetCoursesOfSize retrieves the list of courses from the start for size size
	// Sorted by course code.
	GetCoursesOfSize(size int) (courses []models.Course, n_students []int, n_tutors []int, err error)
	// GetCoursesFromIdOfSize retrieves the list of courses from this course code up to x size/
	// Sorted by course code
	GetCoursesFromIdOfSize(course_code string, size int) (courses []models.Course, n_students []int, n_tutors []int, err error)

	/*
		Tutors model
	*/

	// GetTutors retrieves a list of all tutor model objects from the database.
	GetTutors() (tutors []models.Tutor, err error)
	// GetTutorsFromId retrieves a list of tutors from this tutor id to the end
	GetTutorsFromId(tutor_id int) (tutors []models.Tutor, err error)
	// GetTutorsFromIdOfSize retrieves a list of tutors from this tutor_id to tutor_id + size
	GetTutorsFromIdOfSize(tutor_id int, size int) (tutors []models.Tutor, err error)
	// GetTutorsOfSize retrieves a list of tutors from start to size
	GetTutorsOfSize(size int) (tutors []models.Tutor, err error)
	// GetTutorById retrieves a tutor model object by the tutor's id from the database.
	GetTutorById(tutor_id int) (tutor models.Tutor, err error)
	// GetTutorByEmail retrieves a tutor model object by the tutor's email from the database.
	GetTutorByEmail(email string) (tutor models.Tutor, err error)
	// SaveTutor saves a tutor model object into the database.
	SaveTutor(tutor models.User) (err error)
	// DeleteTutorById deletes a tutor model object by the tutor's id from the database.
	DeleteTutorById(tutor_id int) (err error)
	// DeleteUserByEmail deletes a tutor model object by the tutor's email from the database.
	DeleteTutorByEmail(email string) (err error)

	/*
		Booking model
	*/

	// GetBookingsById retrieves a list of all bookings by a user, as indicated by user_id, with no time constraints
	GetBookingsById(user_id int) (bookings []models.BookingDetails, err error)
	// GetBookingsByIdFrom retrieves a list of all bookings by a user, as indicated by user_id, starting from time from_time
	GetBookingsByIdFrom(user_id int, from_time time.Time) (bookings []models.BookingDetails, err error)
	// GetBookingsById retrieves a list of all bookings by a user, as indicated by user_id, ending with time to_time
	GetBookingsByIdTo(user_id int, to_time time.Time) (bookings []models.BookingDetails, err error)
	// GetBookingsById retrieves a list of all bookings by a user, as indicated by user_id,
	// starting from time from_time and ending with time to_time
	GetBookingsByIdFromTo(user_id int, from_time time.Time, to_time time.Time) (bookings []models.BookingDetails, err error)
	// SaveBooking saves a booking model object into the database
	SaveBooking(availability_id int, user_id int, course_id int) (err error)
	// DeleteBooking deletes a booking model object into the database
	DeleteBookingById(booking_id int) (err error)

	/*
		TutorAvailability model
	*/

	// GetAvailabilityById retrieves a list of all available timeslots for a tutor, with no time constraints
	GetAvailabilityById(tutor_id int) (availabilities []models.Availability, err error)
	// GetAvailabilityById retrieves a list of all available timeslots for a tutor, starting from time from_time
	GetAvailabilityByIdFrom(tutor_id int, from_time time.Time) (availabilities []models.Availability, err error)
	// GetAvailabilityById retrieves a list of all available timeslots for a tutor, ending with time to_time
	GetAvailabilityByIdTo(tutor_id int, to_time time.Time) (availabilities []models.Availability, err error)
	// GetAvailabilityById retrieves a list of all available timeslots for a tutor,
	// starting from time from_time and ending with time to_time
	GetAvailabilityByIdFromTo(tutor_id int, from_time time.Time, to_time time.Time) (availabilities []models.Availability, err error)
	// SaveTutorAvailability saves a tutor availability model object into the database
	SaveTutorAvailability(tutor_id int, from_time time.Time, to_time time.Time) (err error)
	// DeleteTutorAvailability deletes a tutor availability model object by ID from the database
	DeleteTutorAvailabilityById(availability_id int) (err error)
}

func InitDatabase() (err error) {
	isMock, _ := strconv.ParseBool(viper.GetString("mock_database"))
	if isMock {
		// TODO: Chong Yan, please change the methods here for your mockdb if you'd still like to test with it
		// CurrentDatabaseConnector = InitMock()
		return
	}
	// place the db that you want to instantiate here
	// commenting this out until sqlite implement the required methods
	CurrentDatabaseConnector = InitSQLite()
	return
}
