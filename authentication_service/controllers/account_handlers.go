package controllers

import (
	"strconv"
	"strings"

	"net/http"

	"github.com/gin-gonic/gin"

	authHelper "github.com/leechongyan/Studtor_backend/authentication_service/helpers/account"
	typeHelper "github.com/leechongyan/Studtor_backend/authentication_service/helpers/type_conversion"
	authModel "github.com/leechongyan/Studtor_backend/authentication_service/models"
	userConnector "github.com/leechongyan/Studtor_backend/database_service/connector/user_connector"
	databaseModel "github.com/leechongyan/Studtor_backend/database_service/models"

	databaseError "github.com/leechongyan/Studtor_backend/constants/errors/database_errors"
	httpError "github.com/leechongyan/Studtor_backend/constants/errors/http_errors"
	httpHelper "github.com/leechongyan/Studtor_backend/helpers/http_helpers"
	mailService "github.com/leechongyan/Studtor_backend/mail_service"
	storageService "github.com/leechongyan/Studtor_backend/storage_service"
)

// checking that the email belongs to a given domain
func checkEduDomain(email string, domain string) bool {
	components := strings.Split(email, "@")
	_, dom := components[0], components[1]

	return strings.Contains(dom, domain)
}

func getAccountWithEmail(email string) (user databaseModel.User, err error) {
	userConnector := userConnector.Init()
	user, err = userConnector.SetUserEmail(email).Get()
	return
}

func getAccountWithID(id int) (user databaseModel.User, err error) {
	userConnector := userConnector.Init()
	user, err = userConnector.SetUserId(id).Get()
	return
}

// handle sign up request
func SignUp() gin.HandlerFunc {
	return func(c *gin.Context) {
		var user authModel.User

		err := httpHelper.ExtractPostRequestBody(c, &user)
		if err != nil {
			c.JSON(http.StatusBadRequest, err.Error())
			return
		}

		// validate whether the email is valid with edu
		if !checkEduDomain(*user.Email, "edu") {
			c.JSON(http.StatusBadRequest, httpError.ErrInvalidEmail.Error())
			return
		}

		// check whether this email exist
		_, err = getAccountWithEmail(*user.Email)
		if err == nil {
			c.JSON(http.StatusBadRequest, httpError.ErrExistentAccount.Error())
			return
		}
		if err != nil && err == databaseError.ErrDatabaseInternalError {
			c.JSON(http.StatusInternalServerError, err.Error())
			return
		}

		userConnector := userConnector.Init()

		// hash the user password
		password := authHelper.HashPassword(*user.Password)
		user.Password = &password

		// save the user profile picture
		url, err := uploadUserProfilePicture(c)
		if err == nil {
			user.ProfilePicture = &url
		}

		// create a new verification code for user
		newVKey := authHelper.GenerateVerificationCode(6)
		user.VKey = &newVKey

		// need to convert auth user to database user
		database_user := typeHelper.ConvertFromAuthUserToDatabaseUser(user)

		// save user in db
		err = userConnector.SetUser(database_user).Add()

		if err != nil {
			c.JSON(http.StatusInternalServerError, err.Error())
			return
		}

		// send an email
		err = mailService.SendVerificationCode(user, newVKey)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err.Error())
			return
		}

		c.JSON(http.StatusOK, "Success")
	}
}

func Verify() gin.HandlerFunc {
	return func(c *gin.Context) {
		var verification authModel.Verification

		err := httpHelper.ExtractPostRequestBody(c, &verification)
		if err != nil {
			c.JSON(http.StatusBadRequest, err.Error())
			return
		}

		userConnector := userConnector.Init()
		user, err := getAccountWithEmail(*verification.Email)
		if err != nil {
			if err == databaseError.ErrNoEntry {
				c.JSON(http.StatusUnauthorized, httpError.ErrNonExistentAccount.Error())
				return
			}
			c.JSON(http.StatusInternalServerError, err.Error())
			return
		}

		// check whether verification code is correct
		if user.VKey.String != *verification.VKey {
			c.JSON(http.StatusUnauthorized, httpError.ErrWrongValidation.Error())
			return
		}

		// if verification all pass and correct then create access token
		user.Verified = true
		err = userConnector.SetUser(user).Add()
		if err != nil {
			c.JSON(http.StatusInternalServerError, err.Error())
			return
		}

		c.JSON(http.StatusOK, "Success")
	}
}

func Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		var user authModel.Login
		var foundUser databaseModel.User

		err := httpHelper.ExtractPostRequestBody(c, &user)
		if err != nil {
			c.JSON(http.StatusBadRequest, err.Error())
			return
		}

		foundUser, err = getAccountWithEmail(*user.Email)
		// check whether user exists
		if err != nil {
			if err == databaseError.ErrNoEntry {
				c.JSON(http.StatusUnauthorized, httpError.ErrNonExistentAccount.Error())
				return
			}
			c.JSON(http.StatusInternalServerError, err.Error())
			return
		}

		// check whether password exists
		passwordIsValid := authHelper.VerifyPassword(*user.Password, foundUser.Password)
		if !passwordIsValid {
			c.JSON(http.StatusUnauthorized, httpError.ErrWrongCredential.Error())
			return
		}

		// check whether is verified or not
		if !foundUser.Verified {
			c.JSON(http.StatusUnauthorized, httpError.ErrNotVerified.Error())
			return
		}

		// refresh token
		token, refreshToken, err := authHelper.GenerateAllTokens(int(foundUser.ID), foundUser.Email, foundUser.FirstName, foundUser.LastName, foundUser.UserType)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err.Error())
			return
		}

		err = authHelper.UpdateAllTokens(token, refreshToken, foundUser.Email)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err.Error())
			return
		}

		c.JSON(http.StatusOK, token)
	}
}

func RefreshToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		var refresh authModel.Refresh

		err := httpHelper.ExtractPostRequestBody(c, &refresh)
		if err != nil {
			c.JSON(http.StatusBadRequest, err.Error())
			return
		}

		// check whether user exists
		foundUser, err := getAccountWithEmail(*refresh.Email)
		// check whether user exists
		if err != nil {
			if err == databaseError.ErrNoEntry {
				c.JSON(http.StatusUnauthorized, httpError.ErrNonExistentAccount.Error())
				return
			}
			c.JSON(http.StatusInternalServerError, err.Error())
			return
		}

		if !foundUser.RefreshToken.Valid {
			c.JSON(http.StatusUnauthorized, httpError.ErrExpiredRefreshToken.Error())
			return
		}

		_, err = authHelper.ValidateToken(foundUser.RefreshToken.String)
		if err != nil {
			c.JSON(http.StatusUnauthorized, err.Error())
			return
		}

		// if refresh is still valid then generate new token
		token, _, err := authHelper.GenerateAllTokens(int(foundUser.ID), foundUser.Email, foundUser.FirstName, foundUser.LastName, foundUser.UserType)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err.Error())
			return
		}

		err = authHelper.UpdateAllTokens(token, foundUser.RefreshToken.String, foundUser.Email)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err.Error())
			return
		}

		c.JSON(http.StatusOK, token)
	}
}

func Logout() gin.HandlerFunc {
	return func(c *gin.Context) {
		uid := c.GetString("id")
		userId, e := strconv.Atoi(uid)
		if e != nil {
			c.JSON(http.StatusBadRequest, httpError.ErrParamParsingFailure.Error())
			return
		}

		// check whether user exists
		foundUser, err := getAccountWithID(userId)
		// check whether user exists
		if err != nil {
			if err == databaseError.ErrNoEntry {
				c.JSON(http.StatusUnauthorized, httpError.ErrNonExistentAccount.Error())
				return
			}
			c.JSON(http.StatusInternalServerError, err.Error())
			return
		}

		// remove refresh token and force user to login again
		foundUser.Token.Valid = false
		foundUser.RefreshToken.Valid = false

		userConnector := userConnector.Init()
		err = userConnector.SetUser(foundUser).Add()

		if err != nil {
			c.JSON(http.StatusInternalServerError, err.Error())
			return
		}

		c.JSON(http.StatusOK, "Success")
	}
}

func uploadUserProfilePicture(c *gin.Context) (url string, err error) {
	file, fileHeader, err := c.Request.FormFile("file")
	if err != nil {
		return "", httpError.ErrFileParsingFailure
	}
	defer file.Close()
	url, err = storageService.CurrentStorageConnector.SaveUserProfilePicture(c.GetString("id"), file, *fileHeader)
	return
}

func GetUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		user := c.Param("user_id")
		userId, e := strconv.Atoi(user)
		if e != nil {
			c.JSON(http.StatusBadRequest, httpError.ErrParamParsingFailure.Error())
			return
		}

		// check whether user exists
		foundUser, err := getAccountWithID(userId)
		if err != nil {
			if err == databaseError.ErrNoEntry {
				c.JSON(http.StatusUnauthorized, httpError.ErrNonExistentAccount.Error())
				return
			}
			c.JSON(http.StatusInternalServerError, err.Error())
			return
		}

		profile := typeHelper.ConvertFromDatabaseUserToUserProfile(foundUser)
		c.JSON(http.StatusOK, profile)
	}
}

func GetCurrentUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.GetString("id")
		userId, _ := strconv.Atoi(id)

		// check whether user exists
		foundUser, err := getAccountWithID(userId)
		if err != nil {
			if err == databaseError.ErrNoEntry {
				c.JSON(http.StatusUnauthorized, httpError.ErrNonExistentAccount.Error())
				return
			}
			c.JSON(http.StatusInternalServerError, err.Error())
			return
		}

		profile := typeHelper.ConvertFromDatabaseUserToUserProfile(foundUser)
		c.JSON(http.StatusOK, profile)
	}
}
