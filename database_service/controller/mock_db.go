package controller

import (
	"errors"
	"time"

	auth_model "github.com/leechongyan/Studtor_backend/authentication_service/models"
	"github.com/leechongyan/Studtor_backend/constants"
	"github.com/leechongyan/Studtor_backend/database_service/models"
	data_model "github.com/leechongyan/Studtor_backend/database_service/models"
)

// mock db
// this mock db has to implement all the methods which will be used by DatabaseConnector

type mockdb struct {
	UserCollection
}

type UserCollection map[string]auth_model.User

func InitMock() (db mockdb) {
	db = mockdb{}
	db.UserCollection = make(map[string]auth_model.User)
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

func (db mockdb) SaveUser(user auth_model.User) (err error) {
	id := 10
	user.Id = &id
	db.UserCollection[*user.Email] = user
	return
}
func (db mockdb) GetUserById(user_id int) (user auth_model.User, err error) {
	return auth_model.User{}, nil
}
func (db mockdb) GetUserByEmail(email string) (user auth_model.User, err error) {
	user, ok := db.UserCollection[email]
	if !ok {
		err = errors.New(constants.USER_NOT_IN_DATABASE)
		return
	}
	return
}
func (db mockdb) DeleteUserById(user_id int) (err error) {
	return nil
}
func (db mockdb) DeleteUserByEmail(email string) (err error) {
	return nil
}

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
func (db mockdb) GetCourses() (courses []data_model.Course, err error) {
	return nil, nil
}
func (db mockdb) GetCoursesIdSize(id int, size int) (courses []data_model.Course, err error) {

	return nil, nil
}
func (db mockdb) GetCoursesId(id int) (courses []data_model.Course, err error) {

	return nil, nil
}
func (db mockdb) GetCoursesSize(size int) (courses []data_model.Course, err error) {

	return nil, nil
}

func (db mockdb) GetTutorsCourse(course_id int) (tutors []data_model.Tutor, err error) {

	return nil, nil
}
func (db mockdb) GetTutorsCourseIdSize(course_id int, tut_id int, size int) (tutors []data_model.Tutor, err error) {

	return nil, nil
}
func (db mockdb) GetTutorsCourseId(course_id int, tut_id int) (tutors []data_model.Tutor, err error) {

	return nil, nil
}
func (db mockdb) GetTutorsCourseSize(course_id int, size int) (tutors []data_model.Tutor, err error) {

	return nil, nil
}

func (db mockdb) GetTutors() (tutors []models.Tutor, err error) {

	return nil, nil
}
func (db mockdb) GetTutorsIdSize(tut_id int, size int) (tutors []data_model.Tutor, err error) {

	return nil, nil
}
func (db mockdb) GetTutorsId(tut_id int) (tutors []data_model.Tutor, err error) {

	return nil, nil
}
func (db mockdb) GetTutorsSize(size int) (tutors []data_model.Tutor, err error) {

	return nil, nil
}

func (db mockdb) GetTimeBookId(user_id int) (times []data_model.Booking, err error) {

	return nil, nil
}
func (db mockdb) GetTimeBookIdFrom(user_id int, from_time time.Time) (times []data_model.Booking, err error) {

	return nil, nil
}
func (db mockdb) GetTimeBookIdTo(user_id int, to_time time.Time) (times []data_model.Booking, err error) {

	return nil, nil
}
func (db mockdb) GetTimeBookIdFromTo(user_id int, from_time time.Time, to_time time.Time) (times []data_model.Booking, err error) {

	return nil, nil
}

func (db mockdb) GetTimeAvailableId(user_id int) (times []data_model.Availability, err error) {

	return nil, nil
}
func (db mockdb) GetTimeAvailableIdFrom(user_id int, from_time time.Time) (times []data_model.Availability, err error) {

	return nil, nil
}
func (db mockdb) GetTimeAvailableIdTo(user_id int, to_time time.Time) (times []data_model.Availability, err error) {

	return nil, nil
}
func (db mockdb) GetTimeAvailableIdFromTo(user_id int, from_time time.Time, to_time time.Time) (times []data_model.Availability, err error) {
	return nil, nil
}

func (db mockdb) SaveTutorAvailableTimes(user_id int, from_time time.Time, to_time time.Time) (err error) {
	return nil
}
func (db mockdb) DeleteTutorAvailableTimes(user_id int, from_time time.Time, to_time time.Time) (err error) {
	return nil
}
func (db mockdb) BookTutorTime(tutor_id int, student_id int, from_time time.Time, to_time time.Time) (err error) {
	return nil
}
func (db mockdb) UnbookTutorTime(tutor_id int, student_id int, from_time time.Time, to_time time.Time) (err error) {
	return nil
}
