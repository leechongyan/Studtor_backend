package account

import (
	"github.com/gin-gonic/gin"
	httpError "github.com/leechongyan/Studtor_backend/constants/errors/http_errors"
)

func CheckUserType(c *gin.Context, role string) (err error) {
	userType := c.GetString("user_type")
	if userType != role {
		return httpError.ErrUnauthorizedAccess
	}
	return
}

func MatchUserTypeToUid(c *gin.Context, userId string) (err error) {
	userType := c.GetString("user_type")
	uid := c.GetString("uid")

	if userType == "USER" && uid != userId {
		return httpError.ErrUnauthorizedAccess
	}
	return CheckUserType(c, userType)
}
