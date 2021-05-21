package middleware

import (
	"strconv"

	"github.com/gin-gonic/gin"
	helper "github.com/leechongyan/Studtor_backend/authentication_service/helpers/account"
	"github.com/leechongyan/Studtor_backend/helpers"
)

func Authentication(c *gin.Context) {
	clientToken := c.Request.Header.Get("Authorization")
	if clientToken == "" {
		err := helpers.RaiseNoAuthorizationHeader()
		c.JSON(err.StatusCode, err.Error())
		c.Abort()
		return
	}

	clientToken, err := helper.ExtractTokenFromHeader(clientToken)
	if err != nil {
		c.JSON(err.StatusCode, err.Error())
		c.Abort()
		return
	}

	claims, err := helper.ValidateToken(clientToken)
	if err != nil {
		c.JSON(err.StatusCode, err.Error())
		c.Abort()
		return
	}

	c.Set("email", claims.Email)
	c.Set("id", strconv.Itoa(claims.ID))
	c.Set("first_name", claims.First_name)
	c.Set("last_name", claims.Last_name)
	c.Set("user_type", claims.User_type)

	c.Next()
}
