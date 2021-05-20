package account

import (
	"github.com/gin-gonic/gin"
	"github.com/leechongyan/Studtor_backend/helpers"
)

func CheckUserType(c *gin.Context, role string) (err *helpers.RequestError) {
	userType := c.GetString("user_type")
	if userType != role {
		err = helpers.RaiseUnauthorizedAccess()
		return err
	}

	return err
}

func MatchUserTypeToUid(c *gin.Context, userId string) (err *helpers.RequestError) {
	userType := c.GetString("user_type")
	uid := c.GetString("uid")

	if userType == "USER" && uid != userId {
		err = helpers.RaiseUnauthorizedAccess()
		return err
	}
	err = CheckUserType(c, userType)

	return err
}
