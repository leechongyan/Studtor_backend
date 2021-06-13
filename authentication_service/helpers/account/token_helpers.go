package account

import (
	"strings"
	"time"

	httpError "github.com/leechongyan/Studtor_backend/constants/errors/http_errors"
	systemError "github.com/leechongyan/Studtor_backend/constants/errors/system_errors"
	userConnector "github.com/leechongyan/Studtor_backend/database_service/connector/user_connector"
	databaseError "github.com/leechongyan/Studtor_backend/database_service/errors"

	jwt "github.com/dgrijalva/jwt-go"
)

var jwtKey string
var accessExpirationTime int
var refreshExpirationTime int

type SignedDetails struct {
	Email    string
	ID       int
	Name     string
	UserType string
	jwt.StandardClaims
}

func InitJWT(jKey string, accessExpiration int, refreshExpiration int) {
	jwtKey = jKey
	accessExpirationTime = accessExpiration
	refreshExpirationTime = refreshExpiration
}

func GenerateAllTokens(id int, email string, name string, userType string) (signedToken string, signedRefreshToken string, err error) {
	claims := &SignedDetails{
		Email:    email,
		ID:       id,
		Name:     name,
		UserType: userType,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(accessExpirationTime)).Unix(),
		},
	}

	refreshClaims := &SignedDetails{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(refreshExpirationTime)).Unix(),
		},
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(jwtKey))

	if err != nil {
		return "", "", systemError.ErrClaimsGenerateFailure
	}

	refreshToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims).SignedString([]byte(jwtKey))

	if err != nil {
		return "", "", systemError.ErrClaimsGenerateFailure
	}

	return token, refreshToken, err
}

func ValidateToken(signedToken string) (claims *SignedDetails, err error) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&SignedDetails{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(jwtKey), nil
		},
	)

	if err != nil {
		return nil, systemError.ErrClaimsParseFailure
	}

	claims, ok := token.Claims.(*SignedDetails)
	if !ok {
		return nil, httpError.ErrInvalidToken
	}

	if claims.ExpiresAt < time.Now().Local().Unix() {
		return nil, httpError.ErrExpiredToken
	}

	return claims, nil
}

func UpdateAllTokens(signedToken string, signedRefreshToken string, userEmail string) (err error) {
	userConnector := userConnector.Init()
	oldUser, err := userConnector.SetUserEmail(userEmail).GetUser()
	if err != nil {
		if err == databaseError.ErrNoRecordFound {
			return httpError.ErrNonExistentAccount
		}
		return err
	}

	oldUser.SetToken(signedToken)
	oldUser.SetRefreshToken(signedRefreshToken)

	// if is a new creation
	if oldUser.UserCreatedAt().IsZero() {
		createdAt, _ := time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		oldUser.SetUserCreatedAt(createdAt)
	}

	updatedAt, _ := time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	oldUser.SetUserUpdatedAt(updatedAt)

	// updating database
	_, err = userConnector.SetUser(oldUser).Add()
	return err
}

func ExtractTokenFromHeader(header string) (token string, err error) {
	splitToken := strings.Split(header, " ")

	if len(splitToken) != 2 {
		return "", httpError.ErrInvalidTokenFormat
	}

	if splitToken[0] != "Bearer" {
		return "", httpError.ErrInvalidAuthorizationMethod
	}

	return splitToken[1], nil
}
