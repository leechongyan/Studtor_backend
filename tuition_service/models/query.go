package models

import (
	"time"
)

// type Course_query struct {
// 	Course_code string `form:"course_code" validate:"required"`
// 	From        string `form:"from" validate:"required"`
// 	Size        int    `form:"size" validate:"required"`
// }

// type Paginated_query struct {
// 	From string `form:"from" validate:"required"`
// 	Size int    `form:"size" validate:"required"`
// }

type TimeFrame_query struct {
	Email  *string   `json:"email" validate:"email,required"`
	Course *string   `json:"course" validate:"required"`
	From   time.Time `json:"from"`
	To     time.Time `json:"to"`
}
