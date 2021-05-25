package models

import (
	"time"
)

//User is the model that governs all notes objects retrived or inserted into the DB
type User struct {
	Id             *int    `json:"id"`
	FirstName      *string `json:"first_name" validate:"required,min=2,max=100"`
	LastName       *string `json:"last_name" validate:"required,min=2,max=100"`
	Password       *string `json:"password" validate:"required,min=6"`
	Email          *string `json:"email" validate:"email,required"`
	Token          *string `json:"token"`
	UserType       *string `json:"user_type" validate:"required,eq=ADMIN|eq=USER"`
	RefreshToken   *string `json:"refresh_token"`
	VKey           *string
	ProfilePicture *string
	Verified       bool
	UserCreatedAt  time.Time `json:"created_at"`
	UserUpdatedAt  time.Time `json:"updated_at"`
}

type Verification struct {
	Email *string `json:"email" validate:"email,required"`
	VKey  *string `json:"verification_key" validate:"required"`
}

type Login struct {
	Email    *string `json:"email" validate:"email,required"`
	Password *string `json:"password" validate:"required,min=6"`
}

type Refresh struct {
	Email *string `json:"email" validate:"email,required"`
}
