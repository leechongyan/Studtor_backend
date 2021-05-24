package account

import (
	"log"
	"strconv"
	"strings"
	"time"

	userConnector "github.com/leechongyan/Studtor_backend/database_service/connector/user_connector"
	errorHelper "github.com/leechongyan/Studtor_backend/helpers/error_helpers"
	"github.com/spf13/viper"

	jwt "github.com/dgrijalva/jwt-go"
)

type SignedDetails struct {
	Email     string
	ID        int
	FirstName string
	LastName  string
	UserType  string
	jwt.StandardClaims
}

var SECRET_KEY string = viper.GetString("jwtKey")

func GenerateAllTokens(id int, email string, firstName string, lastName string, userType string) (signedToken string, signedRefreshToken string, err *errorHelper.RequestError) {
	accessDuration, _ := strconv.Atoi(viper.GetString("accessExpirationTime"))
	refreshDuration, _ := strconv.Atoi(viper.GetString("refreshExpirationTime"))

	claims := &SignedDetails{
		Email:     email,
		ID:        id,
		FirstName: firstName,
		LastName:  lastName,
		UserType:  userType,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(accessDuration)).Unix(),
		},
	}

	refreshClaims := &SignedDetails{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(refreshDuration)).Unix(),
		},
	}

	token, e := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(SECRET_KEY))

	if e != nil {
		log.Panic(err)
		err = errorHelper.RaiseFailureGenerateClaim()
		return
	}

	refreshToken, e := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims).SignedString([]byte(SECRET_KEY))

	if e != nil {
		log.Panic(err)
		err = errorHelper.RaiseFailureGenerateClaim()
		return
	}

	return token, refreshToken, err
}

func ValidateToken(signedToken string) (claims *SignedDetails, err *errorHelper.RequestError) {
	token, e := jwt.ParseWithClaims(
		signedToken,
		&SignedDetails{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(SECRET_KEY), nil
		},
	)

	if e != nil {
		err = errorHelper.RaiseCannotParseClaims()
		return nil, err
	}

	claims, ok := token.Claims.(*SignedDetails)
	if !ok {
		err = errorHelper.RaiseInvalidToken()
		return nil, err
	}

	if claims.ExpiresAt < time.Now().Local().Unix() {
		err = errorHelper.RaiseExpiredToken()
		return nil, err
	}

	return claims, nil
}

func UpdateAllTokens(signedToken string, signedRefreshToken string, userEmail string) (err *errorHelper.RequestError) {
	userConnector := userConnector.Init()
	oldUser, e := userConnector.SetUserEmail(userEmail).Get()
	if e != nil {
		err = errorHelper.RaiseUserNotInDatabase()
		return err
	}

	oldUser.Token.Valid = true
	oldUser.Token.String = signedToken
	oldUser.RefreshToken.Valid = true
	oldUser.RefreshToken.String = signedRefreshToken

	// if is a new creation
	if oldUser.UserCreatedAt.IsZero() {
		createdAt, _ := time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		oldUser.UserCreatedAt = createdAt
	}

	updatedAt, _ := time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	oldUser.UpdatedAt = updatedAt

	// updating database
	e = userConnector.SetUser(oldUser).Add()
	if e != nil {
		err = errorHelper.RaiseCannotSaveUserInDatabase()
		return err
	}
	return
}

func ExtractTokenFromHeader(header string) (token string, err *errorHelper.RequestError) {
	splitToken := strings.Split(header, " ")

	if len(splitToken) != 2 {
		err = errorHelper.RaiseInvalidTokenFormat()
		return "", err
	}

	if splitToken[0] != "Bearer" {
		err = errorHelper.RaiseInvalidAuthorizationMethod()
		return "", err
	}

	return splitToken[1], nil
}
