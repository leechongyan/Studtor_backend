package internal

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"net/http"
	"time"
	"strconv"
)

var jwtKey = []byte(viper.GetString("jwtKey"))

// mock users
var users = map[string]string{
	"user1": "password1",
	"user2": "password2",
}

type Credentials struct {
	Password string `json:"password"`
	Username string `json:"username"`
}

//
type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func LoginHandler(c *gin.Context) {
	var creds Credentials

	err := c.BindJSON(&creds)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg":"Bad Request"})
		return
	}

	expectedPassword, ok := users[creds.Username]

	if !ok || expectedPassword != creds.Password {
		c.JSON(http.StatusUnauthorized, gin.H{"msg":"Status Unauthorized"})
		return
	}

	e, _ := strconv.Atoi(viper.GetString("expirationTime"))
	expirationTime := time.Now().Add(time.Duration(e)*time.Minute)

	// Create the JWT claims, which includes the username and expiry time
	claims := &Claims{
		Username: creds.Username,
		StandardClaims: jwt.StandardClaims{
			// In JWT, the expiry time is expressed as unix milliseconds
			ExpiresAt: expirationTime.Unix(),
		},
	}

	// Declare the token with the algorithm used for signing, and the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// Create the JWT string
	tokenString, err := token.SignedString(jwtKey)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg":"Internal Error"})
		return
	}

	// Set the client cookie for "token" as the JWT we just generated
	http.SetCookie(c.Writer, &http.Cookie{
		Name:    "token",
		Value:   tokenString,
		Expires: expirationTime,
	})
}

func ValidateHandler(c *gin.Context) {
	cookie, err := c.Request.Cookie("token")
	if err != nil {
		if err == http.ErrNoCookie {
			// If the cookie is not set, return an unauthorized status
			c.JSON(http.StatusUnauthorized, gin.H{"msg":"Status Unauthorized"})
			return
		}
		c.JSON(http.StatusBadRequest, gin.H{"msg":"Bad Request"})
		return
	}

	// Get the JWT string from the cookie
	tknStr := cookie.Value

	claims := &Claims{}

	// Parse the JWT string and store the result in `claims`.
	// Note that we are passing the key in this method as well. This method will return an error
	// if the token is invalid (if it has expired according to the expiry time we set on sign in),
	// or if the signature does not match
	tkn, err := jwt.ParseWithClaims(tknStr, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			c.JSON(http.StatusUnauthorized, gin.H{"msg":"Status Unauthorized"})
			return
		}
		c.JSON(http.StatusBadRequest, gin.H{"msg":"Bad Request"})
		return
	}
	if !tkn.Valid {
		c.JSON(http.StatusUnauthorized, gin.H{"msg":"Status Unauthorized"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"msg":fmt.Sprintf("Welcome %s!", claims.Username)})
}

func RefreshHandler(c *gin.Context) {
	cookie, err := c.Request.Cookie("token")
	if err != nil {
		if err == http.ErrNoCookie {
			c.JSON(http.StatusUnauthorized, gin.H{"msg":"Status Unauthorized"})
			return
		}
		c.JSON(http.StatusBadRequest, gin.H{"msg":"Bad Request"})
		return
	}

	tknStr := cookie.Value
	claims := &Claims{}

	tkn, err := jwt.ParseWithClaims(tknStr, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if !tkn.Valid {
		c.JSON(http.StatusUnauthorized, gin.H{"msg":"Status Unauthorized"})
		return
	}
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			c.JSON(http.StatusUnauthorized, gin.H{"msg":"Status Unauthorized"})
			return
		}
		c.JSON(http.StatusBadRequest, gin.H{"msg":"Bad Request"})
		return
	}
	// We ensure that a new token is not issued until enough time has elapsed
	// In this case, a new token will only be issued if the old token is within
	// 30 seconds of expiry. Otherwise, return a bad request status
	if time.Unix(claims.ExpiresAt, 0).Sub(time.Now()) > 30*time.Second {
		c.JSON(http.StatusBadRequest, gin.H{"msg":"Bad Request"})
		return
	}

	// Create a new token for the current use, with a renewed expiration time
	e, _ := strconv.Atoi(viper.GetString("expirationTime"))
	expirationTime := time.Now().Add(time.Duration(e)*time.Minute)

	claims.ExpiresAt = expirationTime.Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg":"Internal Error"})
		return
	}

	// Set the new token as the cookie
	http.SetCookie(c.Writer, &http.Cookie{
		Name:    "token",
		Value:   tokenString,
		Expires: expirationTime,
	})
}