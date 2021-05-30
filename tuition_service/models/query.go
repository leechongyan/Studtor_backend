package models

import (
	"time"
)

// when querying, from_id and size are optional field
// default start from start
// default return all

type TutorPaginatedQuery struct {
	FromId *int `form:"from_id"`
	Size   *int `form:"size"`
}

// default start from start
// default end at end
// need state just the date, time no need to specify
type TimePaginatedQuery struct {
	From time.Time `form:"from"`
	To   time.Time `form:"to"`
}

type TimeSlot struct {
	// TimeId *int `json:"time_id" validate:"required"`
	From time.Time `json:"from"`
	To   time.Time `json:"to"`
}

type Availability struct {
	AvailabilityId *int `json:"availability_id" validate:"required"`
}
