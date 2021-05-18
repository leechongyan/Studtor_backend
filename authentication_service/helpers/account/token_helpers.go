package account

import (
	"log"
	"strings"
	"time"

	"github.com/leechongyan/Studtor_backend/database_service"
	"github.com/leechongyan/Studtor_backend/helpers"
	"github.com/spf13/viper"

	jwt "github.com/dgrijalva/jwt-go"
)

type SignedDetails struct {
	Email      string
	First_name string
	Last_name  string
	User_type  string
	jwt.StandardClaims
}

var SECRET_KEY string = viper.GetString("jwtKey")

func GenerateAllTokens(email string, firstName string, lastName string, userType string) (signedToken string, signedRefreshToken string, err *helpers.RequestError) {
	claims := &SignedDetails{
		Email:      email,
		First_name: firstName,
		Last_name:  lastName,
		User_type:  userType,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(24)).Unix(),
		},
	}

	refreshClaims := &SignedDetails{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(168)).Unix(),
		},
	}

	token, e := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(SECRET_KEY))

	if e != nil {
		log.Panic(err)
		err = helpers.RaiseFailureGenerateClaim()
		return
	}

	refreshToken, e := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims).SignedString([]byte(SECRET_KEY))

	if e != nil {
		log.Panic(err)
		err = helpers.RaiseFailureGenerateClaim()
		return
	}

	return token, refreshToken, err
}

func ValidateToken(signedToken string) (claims *SignedDetails, err *helpers.RequestError) {
	token, e := jwt.ParseWithClaims(
		signedToken,
		&SignedDetails{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(SECRET_KEY), nil
		},
	)

	if e != nil {
		err = helpers.RaiseCannotParseClaims()
		return nil, err
	}

	claims, ok := token.Claims.(*SignedDetails)
	if !ok {
		err = helpers.RaiseInvalidToken()
		return nil, err
	}

	if claims.ExpiresAt < time.Now().Local().Unix() {
		err = helpers.RaiseExpiredToken()
		return nil, err
	}

	return claims, nil
}

func UpdateAllTokens(signedToken string, signedRefreshToken string, userEmail string) (err *helpers.RequestError) {
	oldUser, err := database_service.CurrentDatabaseConnector.GetUser(userEmail)

	if err != nil {
		return
	}

	oldUser.Token = &signedToken
	oldUser.Refresh_token = &signedRefreshToken

	// if is a new creation
	if oldUser.Created_at.IsZero() {
		Created_at, _ := time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		oldUser.Created_at = Created_at
	}

	Updated_at, _ := time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	oldUser.Updated_at = Updated_at

	// updating database
	err = database_service.CurrentDatabaseConnector.SaveUser(oldUser)
	// TODO: Connector to the database not mock object
	return
}

func ExtractTokenFromHeader(header string) (token string, err *helpers.RequestError) {
	splitToken := strings.Split(header, " ")

	if len(splitToken) != 2 {
		err = helpers.RaiseInvalidTokenFormat()
		return "", err
	}

	if splitToken[0] != "Bearer" {
		err = helpers.RaiseInvalidAuthorizationMethod()
		return "", err
	}

	return splitToken[1], nil
}
