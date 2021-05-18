package models

type Pagination struct {
	From string `form:"from" validate:"required"`
	Size int    `form:"size" validate:"required"`
}
