package controllers

import (
	"strconv"
	"strings"

	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"

	authHelper "github.com/leechongyan/Studtor_backend/authentication_service/helpers/account"
	typeHelper "github.com/leechongyan/Studtor_backend/authentication_service/helpers/type_conversion"
	authModel "github.com/leechongyan/Studtor_backend/authentication_service/models"
	userConnector "github.com/leechongyan/Studtor_backend/database_service/connector/user_connector"
	databaseModel "github.com/leechongyan/Studtor_backend/database_service/models"
	errorHelper "github.com/leechongyan/Studtor_backend/helpers/error_helpers"
	httpHelper "github.com/leechongyan/Studtor_backend/helpers/http_helpers"
	mailService "github.com/leechongyan/Studtor_backend/mail_service"
	storageService "github.com/leechongyan/Studtor_backend/storage_service"
)

func CheckEmailDomain(email string, domain string) bool {
	components := strings.Split(email, "@")
	_, dom := components[0], components[1]

	return strings.Contains(dom, domain)
}

func SignUp() gin.HandlerFunc {
	return func(c *gin.Context) {
		var user authModel.User

		err := httpHelper.ExtractPostRequestBody(c, &user)
		if err != nil {
			c.JSON(err.StatusCode, err.Error())
			return
		}

		// validate whether the email is valid with edu
		if !CheckEmailDomain(*user.Email, "edu") {
			err := errorHelper.RaiseInvalidEmail()
			c.JSON(err.StatusCode, err.Error())
			return
		}

		// check whether this email exist
		userConnector := userConnector.Init()
		_, e := userConnector.SetUserEmail(*user.Email).Get()
		if e == nil {
			err := errorHelper.RaiseExistentAccount()
			c.JSON(err.StatusCode, err.Error())
			return
		}

		// hash the user password
		password := authHelper.HashPassword(*user.Password)
		user.Password = &password

		// save the user profile picture
		url, err := uploadUserProfilePicture(c)
		if err == nil {
			user.ProfilePicture = &url
		}

		// create a new verification code for user
		newVKey := authHelper.EncodeToString(6)
		user.VKey = &newVKey

		// need to convert auth user to database user
		database_user := typeHelper.ConvertFromAuthUserToDatabaseUser(user)

		// save user in db
		e = userConnector.SetUser(database_user).Add()
		if e != nil {
			err := errorHelper.RaiseCannotSaveUserInDatabase()
			c.JSON(err.StatusCode, err.Error())
			return
		}

		// send an email
		err = mailService.SendVerificationCode(user, newVKey)
		if err != nil {
			c.JSON(err.StatusCode, err.Error())
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
			c.JSON(err.StatusCode, err.Error())
			return
		}

		userConnector := userConnector.Init()
		user, e := userConnector.SetUserEmail(*verification.Email).Get()
		if e != nil {
			err := errorHelper.RaiseWrongLoginCredentials()
			c.JSON(err.StatusCode, err.Error())
			return
		}

		if !user.VKey.Valid {
			err := errorHelper.RaiseWrongLoginCredentials()
			c.JSON(err.StatusCode, err.Error())
			return
		}
		// check whether verification code is correct
		if user.VKey.String != *verification.VKey {
			err := errorHelper.RaiseWrongValidation()
			c.JSON(err.StatusCode, err.Error())
			return
		}

		// if verification all pass and correct then create access token
		user.Verified = true
		e = userConnector.SetUser(user).Add()
		if e != nil {
			c.JSON(http.StatusInternalServerError, e.Error())
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
			c.JSON(err.StatusCode, err.Error())
			return
		}
		userConnector := userConnector.Init()
		foundUser, e := userConnector.SetUserEmail(*user.Email).Get()
		// check whether user exists
		if e != nil {
			err := errorHelper.RaiseWrongLoginCredentials()
			c.JSON(err.StatusCode, err.Error())
			return
		}

		// check whether password exists
		passwordIsValid := authHelper.VerifyPassword(*user.Password, foundUser.Password)
		if !passwordIsValid {
			err := errorHelper.RaiseWrongLoginCredentials()
			c.JSON(err.StatusCode, err.Error())
			return
		}

		// check whether is verified or not
		if !foundUser.Verified {
			err := errorHelper.RaiseNotVerified()
			c.JSON(err.StatusCode, err.Error())
			return
		}

		// refresh token
		token, refreshToken, err := authHelper.GenerateAllTokens(int(foundUser.ID), foundUser.Email, foundUser.FirstName, foundUser.LastName, foundUser.UserType)
		if err != nil {
			c.JSON(err.StatusCode, err.Error())
			return
		}

		err = authHelper.UpdateAllTokens(token, refreshToken, foundUser.Email)
		if err != nil {
			c.JSON(err.StatusCode, err.Error())
			return
		}

		c.JSON(http.StatusOK, token)
	}
}

func RefreshToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		emailInByte, e := ioutil.ReadAll(c.Request.Body)
		email := string(emailInByte)
		if e != nil {
			err := errorHelper.RaiseCannotParseJson()
			c.JSON(err.StatusCode, err.Error())
			return
		}

		userConnector := userConnector.Init()
		foundUser, e := userConnector.SetUserEmail(email).Get()

		// check whether user exists
		if e != nil {
			err := errorHelper.RaiseWrongLoginCredentials()
			c.JSON(err.StatusCode, err.Error())
			return
		}

		if !foundUser.RefreshToken.Valid {
			err := errorHelper.RaiseLoginExpired()
			c.JSON(err.StatusCode, err.Error())
			return
		}

		_, err := authHelper.ValidateToken(foundUser.RefreshToken.String)
		if err != nil {
			c.JSON(err.StatusCode, err.Error())
			return
		}

		// if refresh is still valid then generate new token
		token, _, err := authHelper.GenerateAllTokens(int(foundUser.ID), foundUser.Email, foundUser.FirstName, foundUser.LastName, foundUser.UserType)
		if err != nil {
			c.JSON(err.StatusCode, err.Error())
			return
		}

		err = authHelper.UpdateAllTokens(token, foundUser.RefreshToken.String, email)
		if err != nil {
			c.JSON(err.StatusCode, err.Error())
			return
		}

		c.JSON(http.StatusOK, token)
	}
}

func Logout() gin.HandlerFunc {
	return func(c *gin.Context) {
		emailInByte, e := ioutil.ReadAll(c.Request.Body)
		email := string(emailInByte)
		if e != nil {
			err := errorHelper.RaiseCannotParseJson()
			c.JSON(err.StatusCode, err.Error())
			return
		}

		userConnector := userConnector.Init()
		foundUser, e := userConnector.SetUserEmail(email).Get()

		// check whether user exists
		if e != nil {
			err := errorHelper.RaiseWrongLoginCredentials()
			c.JSON(err.StatusCode, err.Error())
			return
		}

		// remove refresh token and force user to login again
		foundUser.Token.Valid = false
		foundUser.RefreshToken.Valid = false

		e = userConnector.SetUser(foundUser).Add()

		if e != nil {
			err := errorHelper.RaiseCannotSaveUserInDatabase()
			c.JSON(err.StatusCode, err.Error())
			return
		}

		c.JSON(http.StatusOK, "Success")
	}
}

func GetMain() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, "Success")
	}
}

func uploadUserProfilePicture(c *gin.Context) (url string, err *errorHelper.RequestError) {
	file, fileHeader, e := c.Request.FormFile("file")
	if e != nil {
		err = errorHelper.RaiseCannotParseFile()
		c.JSON(err.StatusCode, err.Error())
		return
	}
	defer file.Close()
	url, e = storageService.CurrentStorageConnector.SaveUserProfilePicture(c.GetString("id"), file, *fileHeader)
	if e != nil {
		err = errorHelper.RaiseStorageFailure()
		c.JSON(err.StatusCode, err.Error())
		return
	}
	return
}

func GetUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		user := c.Param("user")
		userId, _ := strconv.Atoi(user)
		userConnector := userConnector.Init()
		foundUser, e := userConnector.SetUserId(userId).Get()

		if e != nil {
			err := errorHelper.RaiseDatabaseError()
			c.JSON(err.StatusCode, err.Error())
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
		userConnector := userConnector.Init()
		foundUser, e := userConnector.SetUserId(userId).Get()
		if e != nil {
			err := errorHelper.RaiseDatabaseError()
			c.JSON(err.StatusCode, err.Error())
			return
		}
		profile := typeHelper.ConvertFromDatabaseUserToUserProfile(foundUser)
		c.JSON(http.StatusOK, profile)
	}
}
