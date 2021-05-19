package database_service

import (
	"errors"
	"strconv"
	"time"

	auth_model "github.com/leechongyan/Studtor_backend/authentication_service/models"
	tut_model "github.com/leechongyan/Studtor_backend/tuition_service/models"

	"github.com/leechongyan/Studtor_backend/constants"
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

func (db Mockdb) GetAllCourses(from string, size int) (courses []string, err error) {
	c := [...]string{"CZ1001", "CZ2001", "CZ3001", "CZ4001"}
	// convert from to int
	i, _ := strconv.Atoi(from)
	return c[i : i+size], nil
}

func (db Mockdb) GetAllTutors(from string, size int) (tutors []string, err error) {
	c := [...]string{"Chin", "Kangyu", "Jordan", "Chongyan"}
	// convert from to int
	i, _ := strconv.Atoi(from)
	return c[i : i+size], nil
}

func (db Mockdb) GetAllTutorsForACourse(courseID string, from string, size int) (tutors []string, err error) {
	c := [...]string{"Chin", "Kangyu", "Jordan", "Chongyan"}
	// convert from to int
	i, _ := strconv.Atoi(from)
	return c[i : i+size], nil
}

func (db Mockdb) SaveTutorAvailableTimes(slot tut_model.TimeFrame_query) (err error) {
	return
}

func (db Mockdb) DeleteTutorAvailableTimes(slot tut_model.TimeFrame_query) (err error) {
	return
}

func (db Mockdb) GetTutorBookedTimes(slot tut_model.TimeFrame_query) (bookedTimes Timeslots, err error) {
	// should have course code for the booked time slots as well
	slots := make(map[string][]time.Time)
	slots["CZ1003"] = []time.Time{time.Now(), time.Now()}
	slots["CZ1004"] = []time.Time{time.Now(), time.Now()}
	bookedTimes = make(Timeslots)
	bookedTimes["first_name"] = "Jeff"
	bookedTimes["last_name"] = "Lee"
	bookedTimes["email"] = "clee051@e.ntu.edu.sg"
	bookedTimes["time_slots"] = slots
	return
}

func (db Mockdb) GetTutorAvailableTimes(slot tut_model.TimeFrame_query) (availableTimes Timeslots, err error) {
	// extract from
	// query database from and to
	// from := slot.From
	// to := slot.To

	// create 10 timeslots for testing
	slots := make([][]time.Time, 10)
	for i := range slots {
		slots[i] = []time.Time{time.Now(), time.Now()}
	}
	availableTimes = make(Timeslots)

	availableTimes["first_name"] = "Jeff"
	availableTimes["last_name"] = "Lee"
	availableTimes["email"] = "clee051@e.ntu.edu.sg"
	availableTimes["time_slots"] = slots
	return
}

func (db Mockdb) BookTutorTime(student_email string, slot tut_model.TimeFrame_query) (err error) {
	return nil
}

func (db Mockdb) UnBookTutorTime(student_email string, slot tut_model.TimeFrame_query) (err error) {
	return nil
}

func (db Mockdb) GetStudentBookedTimes(slot tut_model.TimeFrame_query) (bookedTimes Timeslots, err error) {
	// should have course code for the booked time slots as well
	slots := make(map[string][]time.Time)
	slots["CZ1003"] = []time.Time{time.Now(), time.Now()}
	slots["CZ1004"] = []time.Time{time.Now(), time.Now()}
	bookedTimes = make(Timeslots)
	bookedTimes["first_name"] = "Jeff"
	bookedTimes["last_name"] = "Lee"
	bookedTimes["email"] = "clee051@e.ntu.edu.sg"
	bookedTimes["time_slots"] = slots
	return
}
