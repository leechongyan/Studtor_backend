package controllers

import (
	"fmt"
	"strconv"
	"strings"

	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"

	helper "github.com/leechongyan/Studtor_backend/authentication_service/helpers/account"
	auth_model "github.com/leechongyan/Studtor_backend/authentication_service/models"
	user_connector "github.com/leechongyan/Studtor_backend/database_service/connector/user_connector"
	database_model "github.com/leechongyan/Studtor_backend/database_service/models"
	"github.com/leechongyan/Studtor_backend/helpers"
	"github.com/leechongyan/Studtor_backend/mail_service"
	"github.com/leechongyan/Studtor_backend/storage_service"
)

type userprofile struct {
	id         string
	first_name string
	last_name  string
}

func CheckEmailDomain(email string, domain string) bool {
	components := strings.Split(email, "@")
	_, dom := components[0], components[1]

	return strings.Contains(dom, domain)
}

// this is only used for initial sign up for loading the user credential into database user object
func convertUserType(authuser auth_model.User) (databaseuser database_model.User) {
	databaseuser = database_model.User{}
	databaseuser.UserCreatedAt = authuser.Created_at
	databaseuser.UserUpdatedAt = authuser.Updated_at
	if authuser.Id != nil {
		databaseuser.ID = uint(*authuser.Id)
	} else {
		databaseuser.ID = 0
	}
	databaseuser.FirstName = *authuser.First_name
	databaseuser.LastName = *authuser.Last_name
	databaseuser.Password = *authuser.Password
	databaseuser.Email = *authuser.Email
	if authuser.Token != nil {
		databaseuser.Token.String = *authuser.Token
		databaseuser.Token.Valid = true
	} else {
		databaseuser.Token.Valid = false
	}
	databaseuser.UserType = *authuser.User_type
	if authuser.Refresh_token != nil {
		databaseuser.RefreshToken.String = *authuser.Refresh_token
		databaseuser.RefreshToken.Valid = true
	} else {
		databaseuser.RefreshToken.Valid = false
	}
	if authuser.V_key != nil {
		databaseuser.VKey.String = *authuser.V_key
		databaseuser.VKey.Valid = true
	} else {
		databaseuser.VKey.Valid = false
	}
	return databaseuser
}

func SignUp() gin.HandlerFunc {
	return func(c *gin.Context) {
		var user auth_model.User

		err := helpers.ExtractPostRequestBody(c, &user)
		if err != nil {
			c.JSON(err.StatusCode, err.Error())
			return
		}

		// validate whether the email is valid with edu
		if !CheckEmailDomain(*user.Email, "edu") {
			err := helpers.RaiseInvalidEmail()
			c.JSON(err.StatusCode, err.Error())
			return
		}

		// check whether this email exist
		get_user_connector := user_connector.Init()
		_, e := get_user_connector.SetUserEmail(*user.Email).Get()
		if e == nil {
			err := helpers.RaiseExistentAccount()
			c.JSON(err.StatusCode, err.Error())
			return
		}
		password := helper.HashPassword(*user.Password)
		user.Password = &password

		// create a new verification code for user
		new_V_key := helper.EncodeToString(6)
		user.V_key = &new_V_key

		// need to convert auth user to database user
		database_user := convertUserType(user)

		// save user in db
		e = get_user_connector.SetUser(database_user).Add()
		if e != nil {
			err := helpers.RaiseCannotSaveUserInDatabase()
			c.JSON(err.StatusCode, err.Error())
			return
		}

		// send an email
		err = mail_service.SendVerificationCode(user, new_V_key)
		if err != nil {
			c.JSON(err.StatusCode, err.Error())
			return
		}

		c.JSON(http.StatusOK, "Success")
	}
}

func Verify() gin.HandlerFunc {
	return func(c *gin.Context) {
		var verification auth_model.Verification

		err := helpers.ExtractPostRequestBody(c, &verification)
		if err != nil {
			c.JSON(err.StatusCode, err.Error())
			return
		}

		get_user_connector := user_connector.Init()
		user, e := get_user_connector.SetUserEmail(*verification.Email).Get()
		if e != nil {
			err := helpers.RaiseWrongLoginCredentials()
			c.JSON(err.StatusCode, err.Error())
			return
		}

		if !user.VKey.Valid {
			err := helpers.RaiseWrongLoginCredentials()
			c.JSON(err.StatusCode, err.Error())
			return
		}
		// check whether verification code is correct
		if user.VKey.String != *verification.V_key {
			err := helpers.RaiseWrongValidation()
			c.JSON(err.StatusCode, err.Error())
			return
		}

		// if verification all pass and correct then create access token
		user.Verified = true
		e = get_user_connector.SetUser(user).Add()
		if e != nil {
			c.JSON(http.StatusInternalServerError, e.Error())
			return
		}

		c.JSON(http.StatusOK, "Success")
	}
}

func Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		var user auth_model.Login
		var foundUser database_model.User

		err := helpers.ExtractPostRequestBody(c, &user)
		if err != nil {
			c.JSON(err.StatusCode, err.Error())
			return
		}
		get_user_connector := user_connector.Init()
		foundUser, e := get_user_connector.SetUserEmail(*user.Email).Get()
		// check whether user exists
		if e != nil {
			err := helpers.RaiseWrongLoginCredentials()
			c.JSON(err.StatusCode, err.Error())
			return
		}

		// check whether password exists
		passwordIsValid := helper.VerifyPassword(*user.Password, foundUser.Password)
		if passwordIsValid != true {
			err := helpers.RaiseWrongLoginCredentials()
			c.JSON(err.StatusCode, err.Error())
			return
		}

		// check whether is verified or not
		if !foundUser.Verified {
			err := helpers.RaiseNotVerified()
			c.JSON(err.StatusCode, err.Error())
			return
		}

		// refresh token
		token, refreshToken, err := helper.GenerateAllTokens(int(foundUser.ID), foundUser.Email, foundUser.FirstName, foundUser.LastName, foundUser.UserType)
		if err != nil {
			c.JSON(err.StatusCode, err.Error())
			return
		}

		err = helper.UpdateAllTokens(token, refreshToken, foundUser.Email)
		if err != nil {
			c.JSON(err.StatusCode, err.Error())
			return
		}

		c.JSON(http.StatusOK, token)
	}
}

func RefreshToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		email_byte, e := ioutil.ReadAll(c.Request.Body)
		email := string(email_byte)
		if e != nil {
			err := helpers.RaiseCannotParseJson()
			c.JSON(err.StatusCode, err.Error())
			return
		}

		get_user_connector := user_connector.Init()
		foundUser, e := get_user_connector.SetUserEmail(email).Get()

		// check whether user exists
		if e != nil {
			err := helpers.RaiseWrongLoginCredentials()
			c.JSON(err.StatusCode, err.Error())
			return
		}

		if !foundUser.RefreshToken.Valid {
			err := helpers.RaiseLoginExpired()
			c.JSON(err.StatusCode, err.Error())
			return
		}

		_, err := helper.ValidateToken(foundUser.RefreshToken.String)
		if err != nil {
			c.JSON(err.StatusCode, err.Error())
			return
		}

		// if refresh is still valid then generate new token
		token, _, err := helper.GenerateAllTokens(int(foundUser.ID), foundUser.Email, foundUser.FirstName, foundUser.LastName, foundUser.UserType)
		if err != nil {
			c.JSON(err.StatusCode, err.Error())
			return
		}

		err = helper.UpdateAllTokens(token, foundUser.RefreshToken.String, email)
		if err != nil {
			c.JSON(err.StatusCode, err.Error())
			return
		}

		c.JSON(http.StatusOK, token)
	}
}

func Logout() gin.HandlerFunc {
	return func(c *gin.Context) {
		email_byte, e := ioutil.ReadAll(c.Request.Body)
		email := string(email_byte)
		if e != nil {
			err := helpers.RaiseCannotParseJson()
			c.JSON(err.StatusCode, err.Error())
			return
		}

		get_user_connector := user_connector.Init()
		foundUser, e := get_user_connector.SetUserEmail(email).Get()

		// check whether user exists
		if e != nil {
			err := helpers.RaiseWrongLoginCredentials()
			c.JSON(err.StatusCode, err.Error())
			return
		}

		// remove refresh token and force user to login again
		foundUser.Token.Valid = false
		foundUser.RefreshToken.Valid = false

		e = get_user_connector.SetUser(foundUser).Add()

		if e != nil {
			err := helpers.RaiseCannotSaveUserInDatabase()
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

func UploadUserProfilePicture() gin.HandlerFunc {
	return func(c *gin.Context) {
		file, fileheader, e := c.Request.FormFile("file")
		if e != nil {
			fmt.Print("Why is there error here!")
			fmt.Print(e)
			err := helpers.RaiseCannotParseFile()
			c.JSON(err.StatusCode, err.Error())
			return
		}
		fmt.Print("next!")
		defer file.Close()
		url, e := storage_service.CurrentStorageConnector.SaveUserProfilePicture(c.GetString("id"), file, *fileheader)
		if e != nil {
			err := helpers.RaiseStorageFailure()
			c.JSON(err.StatusCode, err.Error())
			return
		}
		c.JSON(http.StatusOK, url)
	}
}

func GetUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		profile := userprofile{}
		user := c.Param("user")
		user_id, _ := strconv.Atoi(user)
		// get other user
		get_user_connector := user_connector.Init()
		foundUser, e := get_user_connector.SetUserId(user_id).Get()

		if e != nil {
			err := helpers.RaiseDatabaseError()
			c.JSON(err.StatusCode, err.Error())
			return
		}
		profile.id = strconv.Itoa(int(foundUser.ID))
		profile.first_name = foundUser.FirstName
		profile.last_name = foundUser.LastName
		c.JSON(http.StatusOK, profile)
	}
}

func GetCurrentUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		profile := userprofile{}
		profile.id = c.GetString("id")
		profile.first_name = c.GetString("first_name")
		profile.last_name = c.GetString("last_name")
		c.JSON(http.StatusOK, profile)
		return

	}
}
