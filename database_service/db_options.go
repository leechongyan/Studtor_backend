package database_service

import "time"

type DB_options struct {
	Course *string

	// for getting paginated
	Size    *int
	From_id *string

	Tutor_id   *string
	Student_id *string

	User_id   *string
	From_time time.Time
	To_time   time.Time
}
