package models

import (
	"time"
)

//User is the model that governs all notes objects retrived or inserted into the DB
type Slot_query struct {
	Tutor_email *string   `json:"email" validate:"email,required"`
	From        time.Time `json:"from"`
	To          time.Time `json:"to"`
}
