package database_models

import (
	"time"

	"gorm.io/gorm"
)

// Follow conventions here: https://gorm.io/docs/models.html#Conventions

// Availability is the model for tutor availability ORM
type Availability struct {
	gorm.Model
	ID       uint `gorm:"primaryKey"`
	TutorID  uint
	Date     time.Time
	TimeSlot int
}
