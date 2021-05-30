package models

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
