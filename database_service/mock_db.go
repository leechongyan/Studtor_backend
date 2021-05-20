package database_service

import (
	"errors"
	"fmt"
	"strconv"
	"time"

	auth_model "github.com/leechongyan/Studtor_backend/authentication_service/models"

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

func (db Mockdb) GetAllCourses(db_options DB_options) (courses []interface{}, err error) {
	fmt.Print("START DEBUGGING")
	course1 := make(map[string]interface{})
	course1["code"] = "CZ1003"
	course1["title"] = "Computational Thinking"
	course1["students"] = 10
	course1["tutors"] = 15
	fmt.Print(course1)
	course2 := make(map[string]interface{})
	course2["code"] = "CZ3003"
	course2["title"] = "Object Thinking"
	course2["students"] = 4
	course2["tutors"] = 20
	fmt.Print("START DEBUGGING2")
	c := make([]interface{}, 2)
	fmt.Print("START DEBUGGING3")
	c[0] = course1
	c[1] = course2
	// convert from to int
	return c, nil
}

func (db Mockdb) GetAllTutors(db_options DB_options) (tutors []string, err error) {
	c := [...]string{"Chin", "Kangyu", "Jordan", "Chongyan"}
	// convert from to int
	i, _ := strconv.Atoi(*db_options.From_id)
	return c[i : i+*db_options.Size], nil
}

func (db Mockdb) SaveTutorAvailableTimes(db_options DB_options) (err error) {
	return
}

func (db Mockdb) DeleteTutorAvailableTimes(db_options DB_options) (err error) {
	return
}

func (db Mockdb) GetBookedTimes(db_options DB_options) (bookedTimes Timeslots, err error) {
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

func (db Mockdb) GetTutorAvailableTimes(db_options DB_options) (availableTimes Timeslots, err error) {
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

func (db Mockdb) BookTutorTime(db_options DB_options) (err error) {
	return nil
}

func (db Mockdb) UnBookTutorTime(db_options DB_options) (err error) {
	return nil
}
