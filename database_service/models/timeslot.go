package models

import (
	"time"

	"gorm.io/gorm"
)

type TimeSlot struct {
	gorm.Model
	Course_id    uint
	Tutor_name   string
	Student_name string
	From_time    time.Time
	To_time      time.Time
}
