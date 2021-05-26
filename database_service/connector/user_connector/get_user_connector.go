package user_connector

import (
	"errors"

	"github.com/leechongyan/Studtor_backend/constants"
	databaseService "github.com/leechongyan/Studtor_backend/database_service/controller"
	"github.com/leechongyan/Studtor_backend/database_service/models"
)

type userOptions struct {
	userId *int
	email  *string
	err    error
	user   *models.User
}

type UserConnector interface {
	SetUserId(userId int) *userOptions
	SetUserEmail(email string) *userOptions
	SetUser(user models.User) *userOptions
	Add() (err error)
	Delete() (err error)
	Get() (user models.User, err error)
}

func Init() *userOptions {
	r := userOptions{}
	return &r
}

func (c *userOptions) SetUserId(userId int) *userOptions {
	c.userId = &userId
	return c
}

func (c *userOptions) SetUserEmail(email string) *userOptions {
	c.email = &email
	return c
}

func (c *userOptions) SetUser(user models.User) *userOptions {
	c.user = &user
	return c
}

func (c *userOptions) Add() (err error) {
	if c.err != nil {
		return c.err
	}
	if c.user == nil {
		return errors.New("User object has to be provided")
	}
	// check whether should update or add new to database
	_, err = databaseService.CurrentDatabaseConnector.GetUserByEmail(c.user.Email)
	// differentiate between no user and error in database
	// if is network error
	if err.Error() == constants.DATABASE_ERROR {
		return
	}
	// if error is account does not exist then create user
	if err.Error() == constants.NONEXISTENT_ACCOUNT {
		return databaseService.CurrentDatabaseConnector.CreateUser(*c.user)
	}
	// user exists
	return databaseService.CurrentDatabaseConnector.UpdateUser(*c.user)
}

func (c *userOptions) Delete() (err error) {
	if c.err != nil {
		return c.err
	}
	if c.userId != nil {
		return databaseService.CurrentDatabaseConnector.DeleteUserById(*c.userId)
	}
	if c.email != nil {
		return databaseService.CurrentDatabaseConnector.DeleteUserByEmail(*c.email)
	}
	return errors.New("User id or User email has to be provided")
}

func (c *userOptions) Get() (user models.User, err error) {
	if c.err != nil {
		return models.User{}, c.err
	}
	if c.userId != nil {
		user, err = databaseService.CurrentDatabaseConnector.GetUserById(*c.userId)
		if err != nil {
			return models.User{}, err
		}
		return user, err
	}
	if c.email != nil {
		user, err = databaseService.CurrentDatabaseConnector.GetUserByEmail(*c.email)
		if err != nil {
			return models.User{}, err
		}
		return user, err
	}
	return models.User{}, errors.New("User id or User email has to be provided")
}
