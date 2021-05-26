package availability_connector

import (
	"errors"
	"time"

	databaseService "github.com/leechongyan/Studtor_backend/database_service/controller"
	"github.com/leechongyan/Studtor_backend/database_service/models"
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
	Get() (times []models.Availability, err error)
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
		return errors.New("Two times have to be provided")
	}
	if c.tutorId == nil {
		return errors.New("Tutor id has to be provided")
	}
	return databaseService.CurrentDatabaseConnector.SaveTutorAvailability(*c.tutorId, c.fromTime, c.toTime)
}

func (c *availabilityOptions) Delete() (err error) {
	if c.err != nil {
		return c.err
	}

	if c.availabilityId == nil || c.tutorId == nil {
		return errors.New("Availability id and tutor id have to be provided")
	}

	availabilities, err := databaseService.CurrentDatabaseConnector.GetAvailabilityById(*c.tutorId)
	if err != nil {
		return
	}
	for _, availability := range availabilities {
		if int(availability.TutorID) == *c.tutorId {
			err = databaseService.CurrentDatabaseConnector.DeleteTutorAvailabilityById(*c.availabilityId)
			return
		}
	}
	return errors.New("Not authorised")
}

func (c *availabilityOptions) Get() (times []models.Availability, err error) {
	if c.err != nil {
		return nil, c.err
	}
	if c.tutorId == nil {
		return nil, errors.New("Tutor id has to be provided")
	}
	if !c.fromTime.IsZero() && !c.toTime.IsZero() {
		if c.fromTime.After(c.toTime) {
			return nil, errors.New("From Time cannot be greater than or equal to To Time")
		}
		return databaseService.CurrentDatabaseConnector.GetAvailabilityByIdFromTo(*c.tutorId, c.fromTime, c.toTime)
	}
	if !c.fromTime.IsZero() {
		return databaseService.CurrentDatabaseConnector.GetAvailabilityByIdFrom(*c.tutorId, c.fromTime)
	}
	if !c.toTime.IsZero() {
		return databaseService.CurrentDatabaseConnector.GetAvailabilityByIdTo(*c.tutorId, c.toTime)
	}
	return databaseService.CurrentDatabaseConnector.GetAvailabilityById(*c.tutorId)
}
