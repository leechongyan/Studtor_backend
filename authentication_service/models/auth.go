package models

import (
	"time"
)

//User is the model that governs all notes objects retrived or inserted into the DB
type User struct {
	First_name    *string `json:"first_name" validate:"required,min=2,max=100"`
	Last_name     *string `json:"last_name" validate:"required,min=2,max=100"`
	Password      *string `json:"password" validate:"required,min=6"`
	Email         *string `json:"email" validate:"email,required"`
	Token         *string `json:"token"`
	User_type     *string `json:"user_type" validate:"required,eq=ADMIN|eq=USER"`
	Refresh_token *string `json:"refresh_token"`
	V_key         *string
	Verified      bool
	Created_at    time.Time `json:"created_at"`
	Updated_at    time.Time `json:"updated_at"`
}

type Verifiation struct {
	Email *string `json:"email" validate:"email,required"`
	V_key *string `json:"verification_key" validate:"required"`
}

type Login struct {
	Email    *string `json:"email" validate:"email,required"`
	Password *string `json:"password" validate:"required,min=6"`
}
