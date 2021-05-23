package availability_connector

import (
	"errors"
	"time"

	database_service "github.com/leechongyan/Studtor_backend/database_service/controller"
	"github.com/leechongyan/Studtor_backend/database_service/models"
)

type time_options struct {
	tutor_id        *int
	availability_id *int
	from_time       time.Time
	to_time         time.Time
	err             error
}

type Get_availability_connector interface {
	SetTutorId(tutor_id int) *time_options
	SetAvailabilityId(availability_id int) *time_options
	SetFromTime(from_time time.Time) *time_options
	SetToTime(to_time time.Time) *time_options
	Add() (err error)
	Delete() (err error)
	Get() (times []models.Availability, err error)
}

func Init() *time_options {
	r := time_options{}
	return &r
}

func (c *time_options) SetTutorId(tutor_id int) *time_options {
	c.tutor_id = &tutor_id
	return c
}

func (c *time_options) SetAvailabilityId(availability_id int) *time_options {
	c.availability_id = &availability_id
	return c
}

func (c *time_options) SetFromTime(from_time time.Time) *time_options {
	c.from_time = from_time
	return c
}

func (c *time_options) SetToTime(to_time time.Time) *time_options {
	c.to_time = to_time
	return c
}

func (c *time_options) Add() (err error) {
	if c.err != nil {
		return c.err
	}
	if c.from_time.IsZero() || c.to_time.IsZero() {
		return errors.New("Two times have to be provided")
	}
	if c.tutor_id == nil {
		return errors.New("Tutor id has to be provided")
	}
	return database_service.CurrentDatabaseConnector.SaveTutorAvailability(*c.tutor_id, c.from_time, c.to_time)
}

func (c *time_options) Delete() (err error) {
	if c.err != nil {
		return c.err
	}

	if c.availability_id == nil {
		return errors.New("Availability id has to be provided")
	}
	return database_service.CurrentDatabaseConnector.DeleteTutorAvailabilityById(*c.availability_id)
}

func (c *time_options) Get() (times []models.Availability, err error) {
	if c.err != nil {
		return nil, c.err
	}
	if c.tutor_id == nil {
		return nil, errors.New("Tutor id has to be provided")
	}
	if !c.from_time.IsZero() && !c.to_time.IsZero() {
		if c.from_time.After(c.to_time) {
			return nil, errors.New("From Time cannot be greater than or equal to To Time")
		}
		return database_service.CurrentDatabaseConnector.GetAvailabilityByIdFromTo(*c.tutor_id, c.from_time, c.to_time)
	}
	if !c.from_time.IsZero() {
		return database_service.CurrentDatabaseConnector.GetAvailabilityByIdFrom(*c.tutor_id, c.from_time)
	}
	if !c.to_time.IsZero() {
		return database_service.CurrentDatabaseConnector.GetAvailabilityByIdTo(*c.tutor_id, c.to_time)
	}
	return database_service.CurrentDatabaseConnector.GetAvailabilityById(*c.tutor_id)
}
