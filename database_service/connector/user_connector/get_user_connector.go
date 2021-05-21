package user_connector

import (
	"errors"

	auth_model "github.com/leechongyan/Studtor_backend/authentication_service/models"
	database_operation "github.com/leechongyan/Studtor_backend/database_service/constants"
	database_service "github.com/leechongyan/Studtor_backend/database_service/controller"
)

type user_options struct {
	user_id *int
	email   *string
	err     error
	user    *auth_model.User
	op      database_operation.Operation
}

type Get_time_connector interface {
	PutUserId(user_id int) *user_options
	PutUserEmail(email string) *user_options
	PutUser(user auth_model.User) *user_options
	SetOperation(op database_operation.Operation) *user_options
	Exec() (user *auth_model.User, err error)
}

func Init() *user_options {
	r := user_options{}
	return &r
}

func (c *user_options) SetOperation(op database_operation.Operation) *user_options {
	c.op = op
	return c
}

func (c *user_options) PutUserId(user_id int) *user_options {
	c.user_id = &user_id
	return c
}

func (c *user_options) PutUserEmail(email string) *user_options {
	c.email = &email
	return c
}

func (c *user_options) PutUser(user auth_model.User) *user_options {
	c.user = &user
	return c
}

func (c *user_options) Exec() (user *auth_model.User, err error) {
	// switch statement to see which one to execute
	// check for error first
	if c.err != nil {
		return nil, c.err
	}

	switch c.op {
	case database_operation.Get:
		{
			if c.user_id != nil {
				user, err := database_service.CurrentDatabaseConnector.GetUserById(*c.user_id)
				return &user, err
			}
			if c.email != nil {
				user, err := database_service.CurrentDatabaseConnector.GetUserByEmail(*c.email)
				return &user, err
			}
			return nil, errors.New("User id or User email has to be provided")
		}
	case database_operation.Delete:
		{
			if c.user_id != nil {
				err := database_service.CurrentDatabaseConnector.DeleteUserById(*c.user_id)
				return nil, err
			}
			if c.email != nil {
				err := database_service.CurrentDatabaseConnector.DeleteUserByEmail(*c.email)
				return nil, err
			}
			return nil, errors.New("User id or User email has to be provided")
		}
	case database_operation.Add:
		{
			if c.user == nil {
				return nil, errors.New("User object has to be provided")
			}
			err := database_service.CurrentDatabaseConnector.SaveUser(*c.user)
			return nil, err
		}
	default:
		{
			return nil, errors.New("No operation specified")
		}
	}
}
