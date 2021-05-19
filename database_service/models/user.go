package models

import (
	"database/sql"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID            uint
	First_name    string
	Last_name     string
	Password      string
	Email         string
	Token         sql.NullString
	User_type     string
	Refresh_token sql.NullString
	V_key         sql.NullString
	Verified      int
	Created_at    sql.NullTime
	Updated_at    sql.NullTime
}
