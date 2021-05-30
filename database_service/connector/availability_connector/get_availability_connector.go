package availability_connector

import (
	"time"

	databaseError "github.com/leechongyan/Studtor_backend/constants/errors/database_errors"
	httpError "github.com/leechongyan/Studtor_backend/constants/errors/http_errors"
	databaseService "github.com/leechongyan/Studtor_backend/database_service/controller"
	databaseModel "github.com/leechongyan/Studtor_backend/database_service/database_models"
)

type availabilityOptions struct {
	tutorId        *int
	availabilityId *int
	fromTime       time.Time
	toTime         time.Time
	err            error
}

type AvailabilityConnector interface {
	SetTutorId(tutorId int) *availabilityOptions
	SetAvailabilityId(availabilityId int) *availabilityOptions
	SetFromTime(fromTime time.Time) *availabilityOptions
	SetToTime(toTime time.Time) *availabilityOptions
	Add() (err error)
	Delete() (err error)
	GetAll() (times []databaseModel.Availability, err error)
}

func Init() *availabilityOptions {
	r := availabilityOptions{}
	return &r
}

func (c *availabilityOptions) SetTutorId(tutorId int) *availabilityOptions {
	c.tutorId = &tutorId
	return c
}

func (c *availabilityOptions) SetAvailabilityId(availabilityId int) *availabilityOptions {
	c.availabilityId = &availabilityId
	return c
}

func (c *availabilityOptions) SetFromTime(fromTime time.Time) *availabilityOptions {
	c.fromTime = fromTime
	return c
}

func (c *availabilityOptions) SetToTime(toTime time.Time) *availabilityOptions {
	c.toTime = toTime
	return c
}

func (c *availabilityOptions) Add() (err error) {
	if c.err != nil {
		return c.err
	}
	if c.fromTime.IsZero() || c.toTime.IsZero() {
		return databaseError.ErrNotEnoughParameters
	}
	if c.tutorId == nil {
		return databaseError.ErrNotEnoughParameters
	}
	return databaseService.CurrentDatabaseConnector.CreateTutorAvailability(*c.tutorId, c.fromTime, c.toTime)
}

func (c *availabilityOptions) Delete() (err error) {
	if c.err != nil {
		return c.err
	}

	if c.availabilityId == nil || c.tutorId == nil {
		return databaseError.ErrNotEnoughParameters
	}

	availabilities, err := databaseService.CurrentDatabaseConnector.GetAvailabilityByID(*c.tutorId)
	if err != nil {
		return
	}
	for _, availability := range availabilities {
		if int(availability.TutorID) == *c.tutorId {
			return databaseService.CurrentDatabaseConnector.DeleteTutorAvailabilityByID(*c.availabilityId)
		}
	}
	return httpError.ErrUnauthorizedAccess
}

func (c *availabilityOptions) GetAll() (times []databaseModel.Availability, err error) {
	if c.err != nil {
		return nil, c.err
	}
	if c.tutorId == nil {
		return nil, databaseError.ErrNotEnoughParameters
	}
	if !c.fromTime.IsZero() && !c.toTime.IsZero() {
		if c.fromTime.After(c.toTime) {
			return nil, databaseError.ErrInvalidTimes
		}
		return databaseService.CurrentDatabaseConnector.GetAvailabilityByIDFromTo(*c.tutorId, c.fromTime, c.toTime)
	}
	if !c.fromTime.IsZero() {
		return databaseService.CurrentDatabaseConnector.GetAvailabilityByIDFrom(*c.tutorId, c.fromTime)
	}
	if !c.toTime.IsZero() {
		return databaseService.CurrentDatabaseConnector.GetAvailabilityByIDTo(*c.tutorId, c.toTime)
	}
	return databaseService.CurrentDatabaseConnector.GetAvailabilityByID(*c.tutorId)
}
