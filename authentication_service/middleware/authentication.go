package middleware

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	authHelper "github.com/leechongyan/Studtor_backend/authentication_service/helpers/account"
	httpError "github.com/leechongyan/Studtor_backend/constants/errors/http_errors"
	systemError "github.com/leechongyan/Studtor_backend/constants/errors/system_errors"
)

func Authentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		clientToken := c.Request.Header.Get("Authorization")
		if clientToken == "" {
			c.JSON(http.StatusBadRequest, httpError.ErrNoAuthorizationHeader.Error())
			c.Abort()
			return
		}

		clientToken, err := authHelper.ExtractTokenFromHeader(clientToken)
		if err != nil {
			c.JSON(http.StatusBadRequest, err.Error())
			c.Abort()
			return
		}

		claims, err := authHelper.ValidateToken(clientToken)
		if err != nil {
			if err == systemError.ErrClaimsParseFailure {
				c.JSON(http.StatusInternalServerError, err.Error())
			} else {
				c.JSON(http.StatusUnauthorized, err.Error())
			}
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
