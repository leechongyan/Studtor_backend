package middleware

import (
	"strconv"

	"github.com/gin-gonic/gin"
	authHelper "github.com/leechongyan/Studtor_backend/authentication_service/helpers/account"
	errorHelper "github.com/leechongyan/Studtor_backend/helpers/error_helpers"
)

func Authentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		clientToken := c.Request.Header.Get("Authorization")
		if clientToken == "" {
			err := errorHelper.RaiseNoAuthorizationHeader()
			c.JSON(err.StatusCode, err.Error())
			c.Abort()
			return
		}

		clientToken, err := authHelper.ExtractTokenFromHeader(clientToken)
		if err != nil {
			c.JSON(err.StatusCode, err.Error())
			c.Abort()
			return
		}

		claims, err := authHelper.ValidateToken(clientToken)
		if err != nil {
			c.JSON(err.StatusCode, err.Error())
			c.Abort()
			return
		}

		c.Set("email", claims.Email)
		c.Set("id", strconv.Itoa(claims.ID))
		c.Set("first_name", claims.FirstName)
		c.Set("last_name", claims.LastName)
		c.Set("user_type", claims.UserType)

		c.Next()
	}
}
