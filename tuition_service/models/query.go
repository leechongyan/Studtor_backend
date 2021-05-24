package models

import (
	"time"
)

// when querying, from_id and size are optional field
// default start from start
// default return all
type CoursePaginated_query struct {
	From_id *string `form:"from_id"`
	Size    *int    `form:"size"`
}

type TutorPaginated_query struct {
	From_id *int `form:"from_id"`
	Size    *int `form:"size"`
}

// default start from start
// default end at end
type TimePaginated_query struct {
	From time.Time `form:"from"`
	To   time.Time `form:"to"`
}

type TimeSlot struct {
	From time.Time `json:"from" validate:"required"`
	To   time.Time `json:"to" validate:"required"`
}

// for booking a slot, all fields are required
type BookSlot struct {
	Course          *int `json:"course" validate:"required"`
	Availability_id *int `json:"availability_id" validate:"required"`
}

type Availability struct {
	Availability_id *int `json:"availability_id" validate:"required"`
}
