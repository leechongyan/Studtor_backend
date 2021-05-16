package middleware

import (
	"fmt"
	"net/http"
	"strings"

	helper "github.com/leechongyan/Studtor_backend/authentication_service/helpers/account"

	"github.com/gin-gonic/gin"
)

func Authentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		clientToken := c.Request.Header.Get("Authorization")
		if clientToken == "" {
			c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("No Authorization Header provided")})
			c.Abort()
			return
		}

		splitToken := strings.Split(clientToken, "Bearer ")
		clientToken = splitToken[1]

		claims, err := helper.ValidateToken(clientToken)
		if err != "" {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err})
			c.Abort()
			return
		}

		c.Set("email", claims.Email)
		c.Set("first_name", claims.First_name)
		c.Set("last_name", claims.Last_name)
		c.Set("user_type", claims.User_type)

		c.Next()
	}
}