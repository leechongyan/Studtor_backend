package user_connector

import (
	"errors"

	userModel "github.com/leechongyan/Studtor_backend/database_service/client_models"
	databaseService "github.com/leechongyan/Studtor_backend/database_service/controller"
	databaseModel "github.com/leechongyan/Studtor_backend/database_service/database_models"
	databaseError "github.com/leechongyan/Studtor_backend/database_service/errors"
	"gorm.io/gorm"
)

// user can only access userModel

type userOptions struct {
	userId *int
	email  *string
	err    error
	user   *userModel.User
}

type UserConnector interface {
	SetUserId(userId int) *userOptions
	SetUserEmail(email string) *userOptions
	SetUser(user userModel.User) *userOptions
	Add() (id int, err error)
	Delete() (err error)
	GetUser() (user userModel.User, err error)
	GetProfile() (user userModel.UserProfile, err error)
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

func (c *userOptions) SetUser(user userModel.User) *userOptions {
	c.user = &user
	return c
}

func (c *userOptions) Add() (id int, err error) {
	if c.err != nil {
		return id, c.err
	}
	if c.user == nil {
		return id, databaseError.ErrNotEnoughParameters
	}
	// check whether should update or add new to database
	_, err = databaseService.CurrentDatabaseConnector.GetUserByEmail(*c.user.Email())
	// differentiate between no user and error in database

	// if error is account does not exist then create user
	databaseUser := userModel.ConvertFromAuthUserToDatabaseUser(*c.user)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return databaseService.CurrentDatabaseConnector.CreateUser(databaseUser)
	}
	if err != nil {
		// internal error
		return id, err
	}
	// user exists
	return databaseService.CurrentDatabaseConnector.UpdateUser(databaseUser)
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

func (c *userOptions) GetUser() (user userModel.User, err error) {
	if c.err != nil {
		return userModel.User{}, c.err
	}

	var databaseUser databaseModel.User
	if c.userId != nil {
		databaseUser, err = databaseService.CurrentDatabaseConnector.GetUserByID(*c.userId)
	} else if c.email != nil {
		databaseUser, err = databaseService.CurrentDatabaseConnector.GetUserByEmail(*c.email)
	} else {
		return userModel.User{}, databaseError.ErrNotEnoughParameters
	}
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return userModel.User{}, databaseError.ErrNoRecordFound
		}
		return userModel.User{}, err
	}
	user = userModel.ConvertFromToDatabaseUserToAuthUser(databaseUser)
	return
}

func (c *userOptions) GetProfile() (user userModel.UserProfile, err error) {
	if c.err != nil {
		return userModel.UserProfile{}, c.err
	}

	var databaseUser databaseModel.User
	if c.userId != nil {
		databaseUser, err = databaseService.CurrentDatabaseConnector.GetUserByID(*c.userId)
	} else if c.email != nil {
		databaseUser, err = databaseService.CurrentDatabaseConnector.GetUserByEmail(*c.email)
	} else {
		return userModel.UserProfile{}, databaseError.ErrNotEnoughParameters
	}
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return userModel.UserProfile{}, databaseError.ErrNoRecordFound
		}
		return userModel.UserProfile{}, err
	}
	user = userModel.ConvertFromDatabaseUserToUserProfile(databaseUser)
	return
}
