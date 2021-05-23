package controller

import (
	"errors"
	"time"

	"github.com/leechongyan/Studtor_backend/constants"
	"github.com/leechongyan/Studtor_backend/database_service/models"
)

// mock db
// this mock db has to implement all the methods which will be used by DatabaseConnector

type mockdb struct {
	UserCollection
}

type UserCollection map[string]models.User

func InitMock() (db mockdb) {
	db = mockdb{}
	db.UserCollection = make(map[string]models.User)
	return db
}

// func (db Mockdb) SaveUser(user auth_model.User) (err error) {
// 	// create a unique id for the user
// 	id := 10
// 	user.Id = &id
// 	db.UserCollection[*user.Email] = user
// 	return
// }

// func (db Mockdb) GetUser(email string) (user auth_model.User, err error) {
// 	user, ok := db.UserCollection[email]
// 	if !ok {
// 		err = errors.New(constants.USER_NOT_IN_DATABASE)
// 		return
// 	}
// 	return
// }

// func createCourse() (course data_model.Course) {
// 	course1 := data_model.Course{}
// 	course1.ID = 123
// 	course1.Course_ID = "123"
// 	course1.Course_code = "CZ1003"
// 	course1.Course_name = "Computational Thinking"

// 	return course1
// }

// func createTutor() (course data_model.Tutor) {
// 	tut1 := data_model.Tutor{}
// 	tut1.ID = 123
// 	return tut1
// }

// func createAvailability() (course data_model.Availability) {
// 	tut1 := data_model.Availability{}
// 	tut1.Course_id = 12
// 	tut1.Tutor_name = "Alice"
// 	tut1.Student_name = "Bob"
// 	tut1.From_time = time.Now()
// 	tut1.To_time = time.Now()
// 	return tut1
// }

// func createBooking() (course data_model.Booking) {
// 	tut1 := data_model.Booking{}
// 	tut1.Course_id = 12
// 	tut1.Tutor_name = "Alice"
// 	tut1.Student_name = "Bob"
// 	tut1.From_time = time.Now()
// 	tut1.To_time = time.Now()
// 	return tut1
// }

// for courses
func (db mockdb) GetUsers() (user []models.User, err error) {
	return nil, nil
}

// GetUserById retrieves a user model object by the user's id from the database.
func (db mockdb) GetUserById(user_id int) (user models.User, err error) {
	return models.User{}, nil
}

// GetUserByEmail retrieves a user model object by the user's email from the database.
func (db mockdb) GetUserByEmail(email string) (user models.User, err error) {
	user, ok := db.UserCollection[email]
	if !ok {
		err = errors.New(constants.USER_NOT_IN_DATABASE)
		return
	}
	return
}

// SaveUser saves an auth_model user object into the database.
func (db mockdb) SaveUser(user models.User) (err error) {
	db.UserCollection[user.Email] = user
	return
}

// DeleteUserById deletes an auth_model user object by the user's id from the database.
func (db mockdb) DeleteUserById(user_id int) (err error) {
	return nil
}

// DeleteUserByEmail deletes an auth_model user object by the user's email from the database.
func (db mockdb) DeleteUserByEmail(email string) (err error) {
	return nil
}

// TODO: Chong Yan, please check if the methods commented below may be deleted.
// GetCoursesIdSize(id int, size int) (courses []models.Course, err error)
// GetCoursesId(id int) (courses []models.Course, err error)
// GetCoursesSize(size int) (courses []models.Course, err error)
// jordan reference
// for courses
// GetACourse()
// database.course + size
// tuition.course(size)
// return
// GetACourse() (course, size, err)
// GetCourses() (courses []models.Course, sizes []int, size, err error)

/*
	TutorCourses model
*/

// GetCoursesForTutor retrieves a list of all courses that a tutor is teaching from the database.
// no need for pagination
func (db mockdb) GetCoursesForTutor(tutor_id int) (course []models.Course, err error) {
	return nil, nil
}

// GetTutorsForCourse retrieves a list of all tutors for a particular course from the database.
func (db mockdb) GetTutorsForCourse(course_id int) (tutor []models.Tutor, err error) {
	return nil, nil
}
func (db mockdb) GetTutorsForCourseFromIdOfSize(course_id, tut_id int, size int) (tutor []models.Tutor, err error) {
	return nil, nil
}
func (db mockdb) GetTutorsForCourseFromId(course_id, tut_id int) (tutor []models.Tutor, err error) {
	return nil, nil
}
func (db mockdb) GetTutorsForCourseOfSize(course_id, size int) (tutor []models.Tutor, err error) {
	return nil, nil
}

// SaveTutorCourse saves a tutor_course model object into the database.
// This function is called when a tutor registers interest to teach a course.
func (db mockdb) SaveTutorCourse(tutor_id int, course_id int) (err error) {
	return nil
}

/*
	Courses model
*/

// GetCourses retrieves a course along with the number of students enrolled in the course
// and the number of tutors for the course, from the database.
func (db mockdb) GetCourse(course_id int) (course models.Course, n_students int, n_tutors int, err error) {
	return models.Course{}, 1, 2, nil
}

// GetCourses retrieves a list of all courses, along with the number of students
// enrolled in the course and the number of tutors for the course, from the database.
func (db mockdb) GetCourses() (courses []models.Course, n_students []int, n_tutors []int, err error) {
	return nil, nil, nil, nil
}

// get the list of courses from this course code to the end
func (db mockdb) GetCoursesFromId(course_code string) (courses []models.Course, n_students []int, n_tutors []int, err error) {
	return nil, nil, nil, nil
}

// get the list of courses from the start for x size
func (db mockdb) GetCoursesOfSize(size int) (courses []models.Course, n_students []int, n_tutors []int, err error) {
	return nil, nil, nil, nil
}

// get the list of courses from this course code up to x size
func (db mockdb) GetCoursesFromIdOfSize(course_code string, size int) (courses []models.Course, n_students []int, n_tutors []int, err error) {
	return nil, nil, nil, nil
}

/*
	Tutors model
*/

// GetTutors retrieves a list of all tutor model objects from the database.
func (db mockdb) GetTutors() (tutors []models.Tutor, err error) {
	return nil, nil
}

// Get a list of tutors from this tutor id to the end
func (db mockdb) GetTutorsFromId(tutor_id int) (tutors []models.Tutor, err error) {
	return nil, nil
}
func (db mockdb) GetTutorsFromIdOfSize(tutor_id int, size int) (tutors []models.Tutor, err error) {
	return nil, nil
}
func (db mockdb) GetTutorsOfSize(size int) (tutors []models.Tutor, err error) {
	return nil, nil
}

// GetTutorById retrieves a tutor model object by the tutor's id from the database.
func (db mockdb) GetTutorById(tutor_id int) (tutor models.Tutor, err error) {
	return models.Tutor{}, nil
}

// GetTutorByEmail retrieves a tutor model object by the tutor's email from the database.
func (db mockdb) GetTutorByEmail(email string) (tutor models.Tutor, err error) {
	return models.Tutor{}, nil
}

// SaveTutor saves a tutor model object into the database.
func (db mockdb) SaveTutor(tutor models.Tutor) (err error) {
	return nil
}

// DeleteTutorById deletes a tutor model object by the tutor's id from the database.
func (db mockdb) DeleteTutorById(tutor_id int) (err error) {
	return nil
}

// DeleteUserByEmail deletes a tutor model object by the tutor's email from the database.
func (db mockdb) DeleteTutorByEmail(email string) (err error) {
	return nil
}

/*
	Booking model
*/

// GetBookingsById retrieves a list of all bookings by a user, as indicated by user_id, with no time constraints
func (db mockdb) GetBookingsById(user_id int) (bookings []models.BookingDetails, err error) {
	return nil, nil
}

// GetBookingsByIdFrom retrieves a list of all bookings by a user, as indicated by user_id, starting from time from_time
func (db mockdb) GetBookingsByIdFrom(user_id int, from_time time.Time) (bookings []models.BookingDetails, err error) {
	return nil, nil
}

// GetBookingsById retrieves a list of all bookings by a user, as indicated by user_id, ending with time to_time
func (db mockdb) GetBookingsByIdTo(user_id int, to_time time.Time) (bookings []models.BookingDetails, err error) {
	return nil, nil
}

// GetBookingsById retrieves a list of all bookings by a user, as indicated by user_id,
// starting from time from_time and ending with time to_time
func (db mockdb) GetBookingsByIdFromTo(user_id int, from_time time.Time, to_time time.Time) (bookings []models.BookingDetails, err error) {
	return nil, nil
}

// SaveBooking saves a booking model object into the database
func (db mockdb) SaveBooking(availability_id int, user_id int, course_id int) (err error) {
	return nil
}

// DeleteBooking deletes a booking model object into the database
func (db mockdb) DeleteBookingById(booking_id int) (err error) {
	return nil
}

/*
	TutorAvailability model
*/

// GetAvailabilityById retrieves a list of all available timeslots for a tutor, with no time constraints
func (db mockdb) GetAvailabilityById(tutor_id int) (availabilities []models.Availability, err error) {
	return nil, nil
}

// GetAvailabilityById retrieves a list of all available timeslots for a tutor, starting from time from_time
func (db mockdb) GetAvailabilityByIdFrom(tutor_id int, from_time time.Time) (availabilities []models.Availability, err error) {
	return nil, nil
}

// GetAvailabilityById retrieves a list of all available timeslots for a tutor, ending with time to_time
func (db mockdb) GetAvailabilityByIdTo(tutor_id int, to_time time.Time) (availabilities []models.Availability, err error) {
	return nil, nil
}

// GetAvailabilityById retrieves a list of all available timeslots for a tutor,
// starting from time from_time and ending with time to_time
func (db mockdb) GetAvailabilityByIdFromTo(tutor_id int, from_time time.Time, to_time time.Time) (availabilities []models.Availability, err error) {
	return nil, nil
}

// SaveTutorAvailability saves a tutor availability model object into the database
func (db mockdb) SaveTutorAvailability(tutor_id int, from_time time.Time, to_time time.Time) (err error) {
	return nil
}

// DeleteTutorAvailability deletes a tutor availability model object by ID from the database
func (db mockdb) DeleteTutorAvailabilityById(availability_id int) (err error) {
	return nil
}
