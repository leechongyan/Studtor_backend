package database_service

import "time"

type options struct {
	db Mockdb
 	course string

	// for getting paginated
	size    int
	From_id *string

	Tutor_id   *string // === string because golang doesn't allow mutations of primitives
	Student_id *string

	User_id   *string
	From_time time.Time
	To_time   time.Time
}

type timeslots map[string]interface{}

// SaveTutorAvailableTimes(db_options options) (err error)
// DeleteTutorAvailableTimes(db_options options) (err error)
// GetTutorAvailableTimes(db_options options) (availableTimes Timeslots, err error)
// BookTutorTime(db_options options) (err error)
// UnBookTutorTime(db_options options) (err error)
// GetBookedTimes(db_options options) (bookedTimes Timeslots, err error)

// Lose validation -> errors cannot be returned
type Options interface {
	SetCourse(string) *options 
	SetSize(int) *options 
	GetBookedTimes() (bookedTimes Timeslots, err error)
}

func (options *options) SetCourse(course string) *options {
	options.course = course
	return options 
}

func (options *options) SetSize(size int) *options {
	options.size = size
	return options 
}

func (options *options) GetBookedTimes() (bookedTimes Timeslots, err error) {
	// should have course code for the booked time slots as well
	slots := make(map[string][]time.Time)
	slots["CZ1003"] = []time.Time{time.Now(), time.Now()}
	slots["CZ1004"] = []time.Time{time.Now(), time.Now()}
	bookedTimes = make(Timeslots)
	bookedTimes["first_name"] = "Jeff"
	bookedTimes["last_name"] = "Lee"
	bookedTimes["email"] = "clee051@e.ntu.edu.sg"
	bookedTimes["time_slots"] = slots
	return 
}

// Initializes options with sensible default
func InitOptions() options {
	return options{}
}