package user_connector

import (
	databaseError "github.com/leechongyan/Studtor_backend/constants/errors/database_errors"
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
		return databaseError.ErrNotEnoughParameters
	}
	// check whether should update or add new to database
	_, err = databaseService.CurrentDatabaseConnector.GetUserByEmail(c.user.Email)
	// differentiate between no user and error in database
	// if is database error
	if err == databaseError.ErrDatabaseInternalError {
		return
	}
	// if error is account does not exist then create user
	if err == databaseError.ErrNoEntry {
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
		return databaseService.CurrentDatabaseConnector.DeleteUserByID(*c.userId)
	}
	if c.email != nil {
		return databaseService.CurrentDatabaseConnector.DeleteUserByEmail(*c.email)
	}
	return databaseError.ErrNotEnoughParameters
}

func (c *userOptions) Get() (user models.User, err error) {
	if c.err != nil {
		return models.User{}, c.err
	}
	if c.userId != nil {
		return databaseService.CurrentDatabaseConnector.GetUserByID(*c.userId)
	}
	if c.email != nil {
		return databaseService.CurrentDatabaseConnector.GetUserByEmail(*c.email)
	}
	return models.User{}, databaseError.ErrNotEnoughParameters
}
