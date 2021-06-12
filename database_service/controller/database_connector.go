package controller

import (
	"time"

	// "github.com/leechongyan/Studtor_backend/authentication_service/databaseModel"
	databaseModel "github.com/leechongyan/Studtor_backend/database_service/database_models"
)

var CurrentDatabaseConnector DatabaseConnector

type Timeslots map[string]interface{}

// DatabaseConnector provIDes the methods to interact with the databaseModel in the database
// Refer to `diagrams/studtor.drawio`, entity relationship diagram, for definitions of databaseModel.
type DatabaseConnector interface {
	/*
		Users model
	*/

	// // GetUsers retrieves a list of all user model objects from the database.
	// GetUsers() (users []databaseModel.User, err error)
	// GetUserByID retrieves a user model object by the user's ID from the database.
	GetUserByID(userID int) (user databaseModel.User, err error)
	// GetUserByEmail retrieves a user model object by the user's email from the database.
	GetUserByEmail(email string) (user databaseModel.User, err error)
	// CreateUser saves an user object into the database.
	CreateUser(user databaseModel.User) (id int, err error)
	// UpdateUser updates an user object into the database.
	UpdateUser(user databaseModel.User) (id int, err error)
	// DeleteUserByID deletes an user object by the user's ID from the database.
	DeleteUserByID(userID int) (err error)
	// DeleteUserByEmail deletes an user object by the user's email from the database.
	DeleteUserByEmail(email string) (err error)

	/*
		TutorCourses model
	*/

	// GetCoursesForTutor retrieves a list of all courses that a tutor is teaching from the database.
	GetCoursesForTutor(tutorID int) (courses []databaseModel.Course, nStudents []int, nTutors []int, err error)
	// GetTutorsForCourse retrieves a list of all tutors for a particular course from the database.
	GetTutorsForCourse(courseID int) (tutors []databaseModel.User, err error)
	// GetTutorsForCourseFromIDOfSize retrieves a list of tutors for a particular course from the database,
	// starting from tut_ID to tut_ID + size
	GetTutorsForCourseFromIDOfSize(courseID int, tutorID int, size int) (tutors []databaseModel.User, err error)
	// GetTutorsForCourseFromID retrieves a list of tutors for a particular course from the database,
	// starting from tut_ID to the end
	GetTutorsForCourseFromID(courseID int, tut_ID int) (tutors []databaseModel.User, err error)
	// GetTutorsForCourseFromID retrieves a list of tutors for a particular course from the database,
	// starting from 0 to size
	GetTutorsForCourseOfSize(courseID int, size int) (tutors []databaseModel.User, err error)
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
	GetCourse(courseID int) (course databaseModel.Course, nStudents int, nTutors int, err error)
	// GetCourses retrieves a list of all courses, along with the number of students
	// enrolled in the course and the number of tutors for the course, from the database.
	// Sorted by course code.
	GetCourses() (courses []databaseModel.Course, nStudents []int, nTutors []int, err error)

	/*
		Schools model
	*/

	// GetSchools retrieves the list of schools from the database
	// GetSchools() (schools []databaseModel.School, err error)
	// GetSchoolByInstitutionAndSchoolCode retrieves a school from the database
	// GetSchoolByInstitutionAndSchoolCode(institution string, schoolCode string) (school databaseModel.School, err error)
	GetSchoolsFacultiesCourses() (schools []databaseModel.School, err error)

	/*
		SchoolCourses model
	*/

	// GetCoursesForSchool retrieves a list of course codes attached to a school by school courses ID
	// GetCoursesForSchool(school_id int) (schoolCoursesDetails databaseModel.SchoolCoursesDetails, err error)

	/*
		Booking model
	*/
	// GetSingleBooking gets a single booking details for a booking ID
	GetSingleBooking(bookingID int) (booking databaseModel.Booking, err error)
	// GetBookingsForStudentByID retrieves a list of all bookings by a student, as indicated by userID, with no time constraints
	GetBookingsForStudentByID(userID int) (bookings []databaseModel.Booking, err error)

	// GetBookingsForStudentByIDFromDateForSize retrieves a list of all bookings for a student from a date up to x days
	GetBookingsForStudentByIDFromDateForSize(userID int, date time.Time, days int) (bookings []databaseModel.Booking, err error)

	// GetBookingsForTutorByID retrieves a list of all bookings by a tutor, as indicated by userID, with no time constraints
	GetBookingsForTutorByID(userID int) (bookings []databaseModel.Booking, err error)

	// GetBookingsForTutorByIDFromDateForSize retrieves a list of all bookings for a tutor from a date up to x days
	GetBookingsForTutorByIDFromDateForSize(userID int, date time.Time, days int) (bookings []databaseModel.Booking, err error)

	// CreateBooking saves a booking model object into the database
	CreateBooking(availabilityID int, userID int, courseID int) (id int, err error)
	// DeleteBooking deletes a booking model object into the database
	DeleteBookingByID(bookingID int) (err error)

	/*
		TutorAvailability model
	*/

	// GetAvailabilityByID retrieves a list of all available timeslots for a tutor, with no time constraints
	GetAvailabilityByID(tutorID int) (availabilities []databaseModel.Availability, err error)

	// GetAvailabilityByID retrieves a list of all available timeslots for a tutor, starting from time fromTime
	// GetAvailabilityByIDFrom(tutorID int, fromTime time.Time) (availabilities []databaseModel.Availability, err error)
	// GetAvailabilityByID retrieves a list of all available timeslots for a tutor, ending with time toTime
	// GetAvailabilityByIDTo(tutorID int, toTime time.Time) (availabilities []databaseModel.Availability, err error)

	// GetSingleAvailability gets an availability information based on the availability ID
	GetSingleAvailability(availabilityID int) (availability databaseModel.Availability, err error)
	// GetAvailabilityByID retrieves a list of all available timeslots for a tutor,
	// starting from time fromTime and ending with time toTime
	// GetAvailabilityByIDFromTo(tutorID int, fromTime time.Time, toTime time.Time) (availabilities []databaseModel.Availability, err error)

	// GetAvailabiltyByIDFromDateForSize retrieves a list of all available timeslots for a tutor from a date up to x days
	GetAvailabilityByIDFromDateForSize(tutorId int, date time.Time, days int) (availabilities []databaseModel.Availability, err error)

	// CreateTutorAvailability saves a tutor availability model object into the database
	CreateTutorAvailability(tutorID int, date time.Time, timeID int) (id int, err error)
	// DeleteTutorAvailability deletes a tutor availability model object by ID from the database
	DeleteTutorAvailabilityByID(availabilityID int) (err error)
}

func InitDatabase(isMock bool, config string) (err error) {
	if isMock {
		// TODO: Chong Yan, please change the methods here for your mockdb if you'd still like to test with it
		// CurrentDatabaseConnector = InitMock()
		return
	}
	// place the db that you want to instantiate here
	// commenting this out until sqlite implement the required methods
	// sqlitedb := &SQLiteDB{}
	// sqlitedb.Init()
	CurrentDatabaseConnector, err = InitPostGres(config)

	return
}
