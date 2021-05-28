package controller

import (
	"time"

	databaseError "github.com/leechongyan/Studtor_backend/constants/errors/database_errors"

	db_model "github.com/leechongyan/Studtor_backend/database_service/models"
)

// mock db
// this mock db has to implement all the methods which will be used by DatabaseConnector

type mockdb struct {
	UserCollection
}

type UserCollection map[string]db_model.User

func InitMock() (db mockdb) {
	db = mockdb{}
	db.UserCollection = make(map[string]db_model.User)
	return db
}

func (db mockdb) GetUsers() (users []db_model.User, err error) {
	return
}

// GetUserById retrieves a user model object by the user's id from the database.
func (db mockdb) GetUserByID(userID int) (user db_model.User, err error) {
	return
}

// GetUserByEmail retrieves a user model object by the user's email from the database.
func (db mockdb) GetUserByEmail(email string) (user db_model.User, err error) {
	user, ok := db.UserCollection[email]
	if !ok {
		return db_model.User{}, databaseError.ErrNoEntry
	}
	return
}

// SaveUser saves an auth_model user object into the database.
func (db mockdb) CreateUser(user db_model.User) (err error) {
	db.UserCollection[user.Email] = user
	return
}

// DeleteUserById deletes an auth_model user object by the user's id from the database.
func (db mockdb) DeleteUserByID(userID int) (err error) {
	return nil
}

// DeleteUserByEmail deletes an auth_model user object by the user's email from the database.
func (db mockdb) DeleteUserByEmail(email string) (err error) {
	return nil
}

func (db mockdb) UpdateUser(user db_model.User) (err error) {
	db.UserCollection[user.Email] = user
	return
}

/*
	TutorCourses model
*/

// GetCoursesForTutor retrieves a list of all courses that a tutor is teaching from the database.
func (db mockdb) GetCoursesForTutor(tutorID int) (courses []db_model.Course, nStudents []int, nTutors []int, err error) {
	return
}

// GetTutorsForCourse retrieves a list of all tutors for a particular course from the database.
func (db mockdb) GetTutorsForCourse(courseID int) (tutors []db_model.User, err error) {
	return
}

// GetTutorsForCourseFromIDOfSize retrieves a list of tutors for a particular course from the database,
// starting from tut_ID to tut_ID + size
func (db mockdb) GetTutorsForCourseFromIDOfSize(courseID int, tutorID int, size int) (tutors []db_model.User, err error) {
	return
}

// GetTutorsForCourseFromID retrieves a list of tutors for a particular course from the database,
// starting from tut_ID to the end
func (db mockdb) GetTutorsForCourseFromID(courseID int, tut_ID int) (tutors []db_model.User, err error) {
	return
}

// GetTutorsForCourseFromID retrieves a list of tutors for a particular course from the database,
// starting from 0 to size
func (db mockdb) GetTutorsForCourseOfSize(courseID int, size int) (tutors []db_model.User, err error) {
	return
}

// CreateTutorCourse saves a tutor_course model object into the database.
// This function is called when a tutor registers interest to teach a course.
func (db mockdb) CreateTutorCourse(tutorID int, courseID int) (err error) {
	return
}

// DeleteTutorCourse deletes an tutor course object from the database.
func (db mockdb) DeleteTutorCourse(tutorID int, courseID int) (err error) {
	return
}

/*
	Courses model
*/

// GetCourses retrieves a course along with the number of students enrolled in the course
// and the number of tutors for the course, from the database.
func (db mockdb) GetCourse(courseID int) (course db_model.Course, nStudents int, nTutors int, err error) {
	return
}

// GetCourses retrieves a list of all courses, along with the number of students
// enrolled in the course and the number of tutors for the course, from the database.
// Sorted by course code.
func (db mockdb) GetCourses() (courses []db_model.Course, nStudents []int, nTutors []int, err error) {
	return
}

/*
	Schools model
*/

// GetSchools retrieves the list of schools from the database
func (db mockdb) GetSchools() (schools []db_model.School, err error) {
	return
}

// TODO (Jordan GetCoursesIdForSchool gets a list of course code attached the the school)
func (db mockdb) GetCoursesForSchool(school_id int) (courses db_model.SchoolCoursesDetails, err error) {
	return
}

// GetSchoolByInstitutionAndSchoolCode retrieves a school from the database
func (db mockdb) GetSchoolByInstitutionAndSchoolCode(institution string, schoolCode string) (school db_model.School, err error) {
	return
}

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
func (db mockdb) GetBookingsByID(userID int) (bookings []db_model.BookingDetails, err error) {
	return
}

// GetBookingsByIDFrom retrieves a list of all bookings by a user, as indicated by userID, starting from time fromTime
func (db mockdb) GetBookingsByIDFrom(userID int, fromTime time.Time) (bookings []db_model.BookingDetails, err error) {
	return
}

// GetBookingsByID retrieves a list of all bookings by a user, as indicated by userID, ending with time toTime
func (db mockdb) GetBookingsByIDTo(userID int, toTime time.Time) (bookings []db_model.BookingDetails, err error) {
	return
}

// GetBookingsByID retrieves a list of all bookings by a user, as indicated by userID,
// starting from time fromTime and ending with time toTime
func (db mockdb) GetBookingsByIDFromTo(userID int, fromTime time.Time, toTime time.Time) (bookings []db_model.BookingDetails, err error) {
	return
}

// CreateBooking saves a booking model object into the database
func (db mockdb) CreateBooking(availabilityID int, userID int, courseID int) (err error) {
	return
}

// DeleteBooking deletes a booking model object into the database
func (db mockdb) DeleteBookingByID(bookingID int) (err error) {
	return
}

/*
	TutorAvailability model
*/

// GetAvailabilityByID retrieves a list of all available timeslots for a tutor, with no time constraints
func (db mockdb) GetAvailabilityByID(tutorID int) (availabilities []db_model.Availability, err error) {
	return
}

// GetAvailabilityByID retrieves a list of all available timeslots for a tutor, starting from time fromTime
func (db mockdb) GetAvailabilityByIDFrom(tutorID int, fromTime time.Time) (availabilities []db_model.Availability, err error) {
	return
}

// GetAvailabilityByID retrieves a list of all available timeslots for a tutor, ending with time toTime
func (db mockdb) GetAvailabilityByIDTo(tutorID int, toTime time.Time) (availabilities []db_model.Availability, err error) {
	return
}

// GetAvailabilityByID retrieves a list of all available timeslots for a tutor,
// starting from time fromTime and ending with time toTime
func (db mockdb) GetAvailabilityByIDFromTo(tutorID int, fromTime time.Time, toTime time.Time) (availabilities []db_model.Availability, err error) {
	return
}

// CreateTutorAvailability saves a tutor availability model object into the database
func (db mockdb) CreateTutorAvailability(tutorID int, fromTime time.Time, toTime time.Time) (err error) {
	return
}

// DeleteTutorAvailability deletes a tutor availability model object by ID from the database
func (db mockdb) DeleteTutorAvailabilityByID(availabilityID int) (err error) {
	return
}
