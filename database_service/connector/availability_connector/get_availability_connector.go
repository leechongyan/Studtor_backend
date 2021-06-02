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
	timeId         *int
	date           time.Time
	days           *int
	err            error
}

type AvailabilityConnector interface {
	SetTutorId(tutorId int) *availabilityOptions
	SetAvailabilityId(availabilityId int) *availabilityOptions
	SetTimeId(timeId int) *availabilityOptions
	SetDate(date time.Time) *availabilityOptions
	SetDays(days int) *availabilityOptions
	Add() (err error)
	Delete() (err error)
	GetAll() (times []databaseModel.Availability, err error)
	GetSingle() (time databaseModel.Availability, err error)
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

func (c *availabilityOptions) SetTimeId(timeId int) *availabilityOptions {
	c.timeId = &timeId
	return c
}

func (c *availabilityOptions) SetDate(date time.Time) *availabilityOptions {
	c.date = date
	return c
}

func (c *availabilityOptions) SetDays(days int) *availabilityOptions {
	c.days = &days
	return c
}

func (c *availabilityOptions) Add() (err error) {
	if c.err != nil {
		return c.err
	}
	if c.date.IsZero() || c.timeId == nil || c.tutorId == nil {
		return databaseError.ErrNotEnoughParameters
	}
	return databaseService.CurrentDatabaseConnector.CreateTutorAvailability(*c.tutorId, c.date, *c.timeId)
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
	if c.tutorId == nil || c.date.IsZero() || c.days == nil {
		return nil, databaseError.ErrNotEnoughParameters
	}
	return databaseService.CurrentDatabaseConnector.GetAvailabilityByIDFromDateForSize(*c.tutorId, c.date, *c.days)
}

func (c *availabilityOptions) GetSingle() (time databaseModel.Availability, err error) {
	if c.err != nil {
		return databaseModel.Availability{}, c.err
	}
	if c.availabilityId == nil {
		return databaseModel.Availability{}, databaseError.ErrNotEnoughParameters
	}
	return databaseService.CurrentDatabaseConnector.GetSingleAvailability(*c.availabilityId)
}
