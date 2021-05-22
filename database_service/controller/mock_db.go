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

type Mockdb struct {
	UserCollection
}

type UserCollection map[string]auth_model.User

func (db *Mockdb) Init() {
	db.UserCollection = make(map[string]auth_model.User)
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

func (db Mockdb) SaveUser(user auth_model.User) (err error) {
	id := 10
	user.Id = &id
	db.UserCollection[*user.Email] = user
	return
}
func (db Mockdb) GetUserById(user_id int) (user auth_model.User, err error) {
	return auth_model.User{}, nil
}
func (db Mockdb) GetUserByEmail(email string) (user auth_model.User, err error) {
	user, ok := db.UserCollection[email]
	if !ok {
		err = errors.New(constants.USER_NOT_IN_DATABASE)
		return
	}
	return
}
func (db Mockdb) DeleteUserById(user_id int) (err error) {
	return nil
}
func (db Mockdb) DeleteUserByEmail(email string) (err error) {
	return nil
}

func createCourse() (course data_model.Course) {
	course1 := data_model.Course{}
	course1.ID = 123
	course1.Course_ID = "123"
	course1.Course_code = "CZ1003"
	course1.Course_name = "Computational Thinking"

	return course1
}

func createTutor() (course data_model.Tutor) {
	tut1 := data_model.Tutor{}
	tut1.ID = 123
	return tut1
}

func createTime() (course data_model.TimeSlot) {
	tut1 := data_model.TimeSlot{}
	tut1.Course_id = 12
	tut1.Tutor_name = "Alice"
	tut1.Student_name = "Bob"
	tut1.From_time = time.Now()
	tut1.To_time = time.Now()
	return tut1
}

// for courses
func (db Mockdb) GetCourses() (courses []data_model.Course, err error) {
	course1 := createCourse()
	course2 := createCourse()
	c := make([]data_model.Course, 2)
	c[0] = course1
	c[1] = course2
	return c, nil
}
func (db Mockdb) GetCoursesIdSize(id int, size int) (courses []data_model.Course, err error) {
	course1 := createCourse()
	course2 := createCourse()
	c := make([]data_model.Course, 2)
	c[0] = course1
	c[1] = course2
	return c, nil
}
func (db Mockdb) GetCoursesId(id int) (courses []data_model.Course, err error) {
	course1 := createCourse()
	course2 := createCourse()
	c := make([]data_model.Course, 2)
	c[0] = course1
	c[1] = course2
	return c, nil
}
func (db Mockdb) GetCoursesSize(size int) (courses []data_model.Course, err error) {
	course1 := createCourse()
	course2 := createCourse()
	c := make([]data_model.Course, 2)
	c[0] = course1
	c[1] = course2
	return c, nil
}

func (db Mockdb) GetTutorsCourse(course_id int) (tutors []data_model.Tutor, err error) {
	tut1 := createTutor()
	tut2 := createTutor()
	c := make([]data_model.Tutor, 2)
	c[0] = tut1
	c[1] = tut2
	return c, nil
}
func (db Mockdb) GetTutorsCourseIdSize(course_id int, tut_id int, size int) (tutors []data_model.Tutor, err error) {
	tut1 := createTutor()
	tut2 := createTutor()
	c := make([]data_model.Tutor, 2)
	c[0] = tut1
	c[1] = tut2
	return c, nil
}
func (db Mockdb) GetTutorsCourseId(course_id int, tut_id int) (tutors []data_model.Tutor, err error) {
	tut1 := createTutor()
	tut2 := createTutor()
	c := make([]data_model.Tutor, 2)
	c[0] = tut1
	c[1] = tut2
	return c, nil
}
func (db Mockdb) GetTutorsCourseSize(course_id int, size int) (tutors []data_model.Tutor, err error) {
	tut1 := createTutor()
	tut2 := createTutor()
	c := make([]data_model.Tutor, 2)
	c[0] = tut1
	c[1] = tut2
	return c, nil
}

func (db Mockdb) GetTutors() (tutors []models.Tutor, err error) {
	tut1 := createTutor()
	tut2 := createTutor()
	c := make([]data_model.Tutor, 2)
	c[0] = tut1
	c[1] = tut2
	return c, nil
}
func (db Mockdb) GetTutorsIdSize(tut_id int, size int) (tutors []data_model.Tutor, err error) {
	tut1 := createTutor()
	tut2 := createTutor()
	c := make([]data_model.Tutor, 2)
	c[0] = tut1
	c[1] = tut2
	return c, nil
}
func (db Mockdb) GetTutorsId(tut_id int) (tutors []data_model.Tutor, err error) {
	tut1 := createTutor()
	tut2 := createTutor()
	c := make([]data_model.Tutor, 2)
	c[0] = tut1
	c[1] = tut2
	return c, nil
}
func (db Mockdb) GetTutorsSize(size int) (tutors []data_model.Tutor, err error) {
	tut1 := createTutor()
	tut2 := createTutor()
	c := make([]data_model.Tutor, 2)
	c[0] = tut1
	c[1] = tut2
	return c, nil
}

func (db Mockdb) GetTimeBookId(user_id int) (times []data_model.TimeSlot, err error) {
	t1 := createTime()
	t2 := createTime()
	c := make([]data_model.TimeSlot, 2)
	c[0] = t1
	c[1] = t2
	return c, nil
}
func (db Mockdb) GetTimeBookIdFrom(user_id int, from_time time.Time) (times []data_model.TimeSlot, err error) {
	t1 := createTime()
	t2 := createTime()
	c := make([]data_model.TimeSlot, 2)
	c[0] = t1
	c[1] = t2
	return c, nil
}
func (db Mockdb) GetTimeBookIdTo(user_id int, to_time time.Time) (times []data_model.TimeSlot, err error) {
	t1 := createTime()
	t2 := createTime()
	c := make([]data_model.TimeSlot, 2)
	c[0] = t1
	c[1] = t2
	return c, nil
}
func (db Mockdb) GetTimeBookIdFromTo(user_id int, from_time time.Time, to_time time.Time) (times []data_model.TimeSlot, err error) {
	t1 := createTime()
	t2 := createTime()
	c := make([]data_model.TimeSlot, 2)
	c[0] = t1
	c[1] = t2
	return c, nil
}

func (db Mockdb) GetTimeAvailableId(user_id int) (times []data_model.TimeSlot, err error) {
	t1 := createTime()
	t2 := createTime()
	c := make([]data_model.TimeSlot, 2)
	c[0] = t1
	c[1] = t2
	return c, nil
}
func (db Mockdb) GetTimeAvailableIdFrom(user_id int, from_time time.Time) (times []data_model.TimeSlot, err error) {
	t1 := createTime()
	t2 := createTime()
	c := make([]data_model.TimeSlot, 2)
	c[0] = t1
	c[1] = t2
	return c, nil
}
func (db Mockdb) GetTimeAvailableIdTo(user_id int, to_time time.Time) (times []data_model.TimeSlot, err error) {
	t1 := createTime()
	t2 := createTime()
	c := make([]data_model.TimeSlot, 2)
	c[0] = t1
	c[1] = t2
	return c, nil
}
func (db Mockdb) GetTimeAvailableIdFromTo(user_id int, from_time time.Time, to_time time.Time) (times []data_model.TimeSlot, err error) {
	t1 := createTime()
	t2 := createTime()
	c := make([]data_model.TimeSlot, 2)
	c[0] = t1
	c[1] = t2
	return c, nil
}

func (db Mockdb) SaveTutorAvailableTimes(user_id int, from_time time.Time, to_time time.Time) (err error) {
	return nil
}
func (db Mockdb) DeleteTutorAvailableTimes(user_id int, from_time time.Time, to_time time.Time) (err error) {
	return nil
}
func (db Mockdb) BookTutorTime(tutor_id int, student_id int, from_time time.Time, to_time time.Time) (err error) {
	return nil
}
func (db Mockdb) UnbookTutorTime(tutor_id int, student_id int, from_time time.Time, to_time time.Time) (err error) {
	return nil
}
