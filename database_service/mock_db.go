package database_service

import (
	"strconv"

	"github.com/leechongyan/Studtor_backend/authentication_service/models"
	"github.com/leechongyan/Studtor_backend/helpers"
)

// mock db
// this mock db has to implement all the methods which will be used by DatabaseConnector

type Mockdb struct {
	UserCollection
}

type UserCollection map[string]models.User

func (db *Mockdb) Init() {
	db.UserCollection = make(map[string]models.User)
}

func (db Mockdb) SaveUser(user models.User) (err *helpers.RequestError) {
	db.UserCollection[*user.Email] = user
	return
}

func (db Mockdb) GetUser(email string) (user models.User, err *helpers.RequestError) {
	user, ok := db.UserCollection[email]
	if !ok {
		err = helpers.RaiseUserNotInDatabase()
		return
	}
	return
}

func (db Mockdb) GetAllCourses(from string, size int) (courses []string, err *helpers.RequestError) {
	c := [...]string{"CZ1001", "CZ2001", "CZ3001", "CZ4001"}
	// convert from to int
	i, _ := strconv.Atoi(from)
	return c[i : i+size], nil
}

func (db Mockdb) GetAllTutors(from string, size int) (tutors []string, err *helpers.RequestError) {
	c := [...]string{"Chin", "Kangyu", "Jordan", "Chongyan"}
	// convert from to int
	i, _ := strconv.Atoi(from)
	return c[i : i+size], nil
}
