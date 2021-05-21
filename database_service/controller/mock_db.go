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

func (db Mockdb) SaveUser(user auth_model.User) (err error) {
	// create a unique id for the user
	id := 10
	user.Id = &id
	db.UserCollection[*user.Email] = user
	return
}

func (db Mockdb) GetUser(email string) (user auth_model.User, err error) {
	user, ok := db.UserCollection[email]
	if !ok {
		err = errors.New(constants.USER_NOT_IN_DATABASE)
		return
	}
	return
}

func createCourse() (course data_model.Course) {
	course1 := data_model.Course{}
	course1.ID = 123
	course1.Course_ID = "123"
	course1.Course_code = "CZ1003"
	course1.Course_name = "Computational Thinking"
	course1.Tutor_size = 10

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

// func (db Mockdb) GetAllTutors(db_options DB_options) (tutors []string, err error) {
// 	c := [...]string{"Chin", "Kangyu", "Jordan", "Chongyan"}
// 	// convert from to int
// 	i, _ := strconv.Atoi(*db_options.From_id)
// 	return c[i : i+*db_options.Size], nil
// }

// func (db Mockdb) SaveTutorAvailableTimes(db_options DB_options) (err error) {
// 	return
// }

// func (db Mockdb) DeleteTutorAvailableTimes(db_options DB_options) (err error) {
// 	return
// }

// func (db Mockdb) GetBookedTimes(db_options DB_options) (bookedTimes Timeslots, err error) {
// 	// should have course code for the booked time slots as well
// 	slots := make(map[string][]time.Time)
// 	slots["CZ1003"] = []time.Time{time.Now(), time.Now()}
// 	slots["CZ1004"] = []time.Time{time.Now(), time.Now()}
// 	bookedTimes = make(Timeslots)
// 	bookedTimes["first_name"] = "Jeff"
// 	bookedTimes["last_name"] = "Lee"
// 	bookedTimes["email"] = "clee051@e.ntu.edu.sg"
// 	bookedTimes["time_slots"] = slots
// 	return
// }

// func (db Mockdb) GetTutorAvailableTimes(db_options DB_options) (availableTimes Timeslots, err error) {
// 	// extract from
// 	// query database from and to
// 	// from := slot.From
// 	// to := slot.To

// 	// create 10 timeslots for testing
// 	slots := make([][]time.Time, 10)
// 	for i := range slots {
// 		slots[i] = []time.Time{time.Now(), time.Now()}
// 	}
// 	availableTimes = make(Timeslots)

// 	availableTimes["first_name"] = "Jeff"
// 	availableTimes["last_name"] = "Lee"
// 	availableTimes["email"] = "clee051@e.ntu.edu.sg"
// 	availableTimes["time_slots"] = slots
// 	return
// }

// func (db Mockdb) BookTutorTime(db_options DB_options) (err error) {
// 	return nil
// }

// func (db Mockdb) UnBookTutorTime(db_options DB_options) (err error) {
// 	return nil
// }

// for courses
func (db Mockdb) GetCourses() (courses []models.Course, err error) {
	course1 := createCourse()
	course2 := createCourse()
	c := make([]models.Course, 2)
	c[0] = course1
	c[1] = course2
	return c, nil
}
func (db Mockdb) GetCoursesIdSize(id int, size int) (courses []models.Course, err error) {
	course1 := createCourse()
	course2 := createCourse()
	c := make([]models.Course, 2)
	c[0] = course1
	c[1] = course2
	return c, nil
}
func (db Mockdb) GetCoursesId(id int) (courses []models.Course, err error) {
	course1 := createCourse()
	course2 := createCourse()
	c := make([]models.Course, 2)
	c[0] = course1
	c[1] = course2
	return c, nil
}
func (db Mockdb) GetCoursesSize(size int) (courses []models.Course, err error) {
	course1 := createCourse()
	course2 := createCourse()
	c := make([]models.Course, 2)
	c[0] = course1
	c[1] = course2
	return c, nil
}

func (db Mockdb) GetTutorsCourse(course_id int) (tutors []models.Tutor, err error) {
	tut1 := createTutor()
	tut2 := createTutor()
	c := make([]models.Tutor, 2)
	c[0] = tut1
	c[1] = tut2
	return c, nil
}
func (db Mockdb) GetTutorsCourseIdSize(course_id int, tut_id int, size int) (tutors []models.Tutor, err error) {
	tut1 := createTutor()
	tut2 := createTutor()
	c := make([]models.Tutor, 2)
	c[0] = tut1
	c[1] = tut2
	return c, nil
}
func (db Mockdb) GetTutorsCourseId(course_id int, tut_id int) (tutors []models.Tutor, err error) {
	tut1 := createTutor()
	tut2 := createTutor()
	c := make([]models.Tutor, 2)
	c[0] = tut1
	c[1] = tut2
	return c, nil
}
func (db Mockdb) GetTutorsCourseSize(course_id int, size int) (tutors []models.Tutor, err error) {
	tut1 := createTutor()
	tut2 := createTutor()
	c := make([]models.Tutor, 2)
	c[0] = tut1
	c[1] = tut2
	return c, nil
}

func (db Mockdb) GetTutors() (tutors []models.Tutor, err error) {
	tut1 := createTutor()
	tut2 := createTutor()
	c := make([]models.Tutor, 2)
	c[0] = tut1
	c[1] = tut2
	return c, nil
}
func (db Mockdb) GetTutorsIdSize(tut_id int, size int) (tutors []models.Tutor, err error) {
	tut1 := createTutor()
	tut2 := createTutor()
	c := make([]models.Tutor, 2)
	c[0] = tut1
	c[1] = tut2
	return c, nil
}
func (db Mockdb) GetTutorsId(tut_id int) (tutors []models.Tutor, err error) {
	tut1 := createTutor()
	tut2 := createTutor()
	c := make([]models.Tutor, 2)
	c[0] = tut1
	c[1] = tut2
	return c, nil
}
func (db Mockdb) GetTutorsSize(size int) (tutors []models.Tutor, err error) {
	tut1 := createTutor()
	tut2 := createTutor()
	c := make([]models.Tutor, 2)
	c[0] = tut1
	c[1] = tut2
	return c, nil
}

func (db Mockdb) GetTimeBookId(user_id int) (times []models.TimeSlot, err error) {
	t1 := createTime()
	t2 := createTime()
	c := make([]models.TimeSlot, 2)
	c[0] = t1
	c[1] = t2
	return c, nil
}
func (db Mockdb) GetTimeBookIdFrom(user_id int, from_time time.Time) (times []models.TimeSlot, err error) {
	t1 := createTime()
	t2 := createTime()
	c := make([]models.TimeSlot, 2)
	c[0] = t1
	c[1] = t2
	return c, nil
}
func (db Mockdb) GetTimeBookIdTo(user_id int, to_time time.Time) (times []models.TimeSlot, err error) {
	t1 := createTime()
	t2 := createTime()
	c := make([]models.TimeSlot, 2)
	c[0] = t1
	c[1] = t2
	return c, nil
}
func (db Mockdb) GetTimeBookIdFromTo(user_id int, from_time time.Time, to_time time.Time) (times []models.TimeSlot, err error) {
	t1 := createTime()
	t2 := createTime()
	c := make([]models.TimeSlot, 2)
	c[0] = t1
	c[1] = t2
	return c, nil
}

func (db Mockdb) GetTimeAvailableId(user_id int) (times []models.TimeSlot, err error) {
	t1 := createTime()
	t2 := createTime()
	c := make([]models.TimeSlot, 2)
	c[0] = t1
	c[1] = t2
	return c, nil
}
func (db Mockdb) GetTimeAvailableIdFrom(user_id int, from_time time.Time) (times []models.TimeSlot, err error) {
	t1 := createTime()
	t2 := createTime()
	c := make([]models.TimeSlot, 2)
	c[0] = t1
	c[1] = t2
	return c, nil
}
func (db Mockdb) GetTimeAvailableIdTo(user_id int, to_time time.Time) (times []models.TimeSlot, err error) {
	t1 := createTime()
	t2 := createTime()
	c := make([]models.TimeSlot, 2)
	c[0] = t1
	c[1] = t2
	return c, nil
}
func (db Mockdb) GetTimeAvailableIdFromTo(user_id int, from_time time.Time, to_time time.Time) (times []models.TimeSlot, err error) {
	t1 := createTime()
	t2 := createTime()
	c := make([]models.TimeSlot, 2)
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
