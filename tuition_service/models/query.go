package models

import (
	"time"
)

// when querying, from_id and size are optional field
// default start from start
// default return all
type ObjectPaginated_query struct {
	From_id *string `form:"from_id"`
	Size    *int    `form:"size"`
}

// default start from start
// default end at end
type TimePaginated_query struct {
	From time.Time `form:"from_id"`
	To   time.Time `form:"to"`
}

// for booking a slot, all fields are required
type BookSlot struct {
	Course *string   `json:"course" validate:"required"`
	From   time.Time `json:"from_id" validate:"required"`
	To     time.Time `json:"to" validate:"required"`
	Tutor  *string   `json:"tutor" validate:"required"`
}
