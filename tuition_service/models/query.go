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
type TimePaginatedQuery struct {
	IsStudent *bool     `form:"is_student"`
	From      time.Time `form:"from"`
	To        time.Time `form:"to"`
}

type TimeSlot struct {
	From time.Time `json:"from" validate:"required"`
	To   time.Time `json:"to" validate:"required"`
}

type Availability struct {
	AvailabilityId *int `json:"availability_id" validate:"required"`
}
