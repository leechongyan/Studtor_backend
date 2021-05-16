package account

import (
	"log"
	"time"
	"errors"
	"strings"

	"github.com/leechongyan/Studtor_backend/authentication_service/database"
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

func GenerateAllTokens(email string, firstName string, lastName string, userType string) (signedToken string, signedRefreshToken string, err error) {
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

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(SECRET_KEY))

	if err != nil {
		log.Panic(err)
		return
	}

	refreshToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims).SignedString([]byte(SECRET_KEY))

	if err != nil {
		log.Panic(err)
		return
	}

	return token, refreshToken, err
}

func ValidateToken(signedToken string) (claims *SignedDetails, err error) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&SignedDetails{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(SECRET_KEY), nil
		},
	)

	if err != nil {
		return
	}

	claims, ok := token.Claims.(*SignedDetails)
	if !ok {
		return nil, errors.New("Invalid Token")
	}

	if claims.ExpiresAt < time.Now().Local().Unix() {
		return nil, errors.New("Token Expired")
	}

	return claims, nil
}

func UpdateAllTokens(signedToken string, signedRefreshToken string, userEmail string) (err error) {

	oldUser := database.UserCollection[userEmail]

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
	database.UserCollection[userEmail] = oldUser
	// TODO: Connector to the database not mock object

	return
}

func ExtractTokenFromHeader(header string) (token string, err error) {
	splitToken := strings.Split(header, " ")

	if len(splitToken) != 2 {
		return "", errors.New("Invalid Token Format")
	}

	if splitToken[0] != "Bearer" {
		return "", errors.New("Invalid Authorization Method provided")
	}

	return splitToken[1], nil
}