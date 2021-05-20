package database_service

import "time"

type DB_options struct {
	Course    *string
	Size      *int    `form:"size"`
	From_id   *string `form:"from_id"`
	To_id     *string
	From_time time.Time
	To_time   time.Time
}
