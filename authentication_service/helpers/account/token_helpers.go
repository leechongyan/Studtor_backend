package account

import (
	"strconv"
	"strings"
	"time"

	"github.com/leechongyan/Studtor_backend/constants/errors/database_errors"
	databaseError "github.com/leechongyan/Studtor_backend/constants/errors/database_errors"
	httpError "github.com/leechongyan/Studtor_backend/constants/errors/http_errors"
	systemError "github.com/leechongyan/Studtor_backend/constants/errors/system_errors"
	userConnector "github.com/leechongyan/Studtor_backend/database_service/connector/user_connector"

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

func GenerateAllTokens(id int, email string, firstName string, lastName string, userType string) (signedToken string, signedRefreshToken string, err error) {
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

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(SECRET_KEY))

	if err != nil {
		return "", "", systemError.ErrClaimsGenerateFailure
	}

	refreshToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims).SignedString([]byte(SECRET_KEY))

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
			return []byte(SECRET_KEY), nil
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
	oldUser, err := userConnector.SetUserEmail(userEmail).Get()
	if err != nil {
		if err == databaseError.ErrNoRecordFound {
			return httpError.ErrNonExistentAccount
		}
		return database_errors.ErrDatabaseInternalError
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
	err = userConnector.SetUser(oldUser).Add()
	if err != nil {
		return database_errors.ErrDatabaseInternalError
	}
	return
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
