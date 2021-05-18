package database_service

import (
	"github.com/leechongyan/Studtor_backend/authentication_service/models"
	"github.com/leechongyan/Studtor_backend/helpers"
)

// mock db
// this mock db has to implement all the methods which will be used by DatabaseConnector

type Mockdb struct {
	UserCollection
}

type UserCollection map[string]models.User

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

func (db *Mockdb) Init() {
	db.UserCollection = make(map[string]models.User)
}
