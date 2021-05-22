package user_connector

import (
	"errors"

	auth_model "github.com/leechongyan/Studtor_backend/authentication_service/models"
	database_service "github.com/leechongyan/Studtor_backend/database_service/controller"
)

type user_options struct {
	user_id *int
	email   *string
	err     error
	user    *auth_model.User
}

type Get_user_connector interface {
	SetUserId(user_id int) *user_options
	SetUserEmail(email string) *user_options
	SetUser(user auth_model.User) *user_options
	Add() (err error)
	Delete() (err error)
	Get() (user auth_model.User, err error)
}

func Init() *user_options {
	r := user_options{}
	return &r
}

func (c *user_options) SetUserId(user_id int) *user_options {
	c.user_id = &user_id
	return c
}

func (c *user_options) SetUserEmail(email string) *user_options {
	c.email = &email
	return c
}

func (c *user_options) SetUser(user auth_model.User) *user_options {
	c.user = &user
	return c
}

func (c *user_options) Add() (err error) {
	if c.err != nil {
		return c.err
	}
	if c.user == nil {
		return errors.New("User object has to be provided")
	}
	return database_service.CurrentDatabaseConnector.SaveUser(*c.user)
}

func (c *user_options) Delete() (err error) {
	if c.err != nil {
		return c.err
	}
	if c.user_id != nil {
		return database_service.CurrentDatabaseConnector.DeleteUserById(*c.user_id)
	}
	if c.email != nil {
		return database_service.CurrentDatabaseConnector.DeleteUserByEmail(*c.email)
	}
	return errors.New("User id or User email has to be provided")
}

func (c *user_options) Get() (user auth_model.User, err error) {
	if c.err != nil {
		return auth_model.User{}, c.err
	}
	if c.user_id != nil {
		user, err := database_service.CurrentDatabaseConnector.GetUserById(*c.user_id)
		if err != nil {
			return auth_model.User{}, err
		}
		return user, err
	}
	if c.email != nil {
		user, err := database_service.CurrentDatabaseConnector.GetUserByEmail(*c.email)
		if err != nil {
			return auth_model.User{}, err
		}
		return user, err
	}
	return auth_model.User{}, errors.New("User id or User email has to be provided")
}
