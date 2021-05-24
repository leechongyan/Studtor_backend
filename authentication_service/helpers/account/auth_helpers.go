package account

import (
	"github.com/gin-gonic/gin"
	errorHelper "github.com/leechongyan/Studtor_backend/helpers/error_helpers"
)

func CheckUserType(c *gin.Context, role string) (err *errorHelper.RequestError) {
	userType := c.GetString("user_type")
	if userType != role {
		err = errorHelper.RaiseUnauthorizedAccess()
		return err
	}

	return err
}

func MatchUserTypeToUid(c *gin.Context, userId string) (err *errorHelper.RequestError) {
	userType := c.GetString("user_type")
	uid := c.GetString("uid")

	if userType == "USER" && uid != userId {
		err = errorHelper.RaiseUnauthorizedAccess()
		return err
	}
	err = CheckUserType(c, userType)

	return err
}
