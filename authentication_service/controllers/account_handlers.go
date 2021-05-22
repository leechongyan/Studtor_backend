package controllers

import (
	"strconv"
	"strings"

	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"

	helper "github.com/leechongyan/Studtor_backend/authentication_service/helpers/account"
	"github.com/leechongyan/Studtor_backend/authentication_service/models"
	user_connector "github.com/leechongyan/Studtor_backend/database_service/connector/user_connector"
	"github.com/leechongyan/Studtor_backend/helpers"
	"github.com/leechongyan/Studtor_backend/mail_service"
)

func CheckEmailDomain(email string, domain string) bool {
	components := strings.Split(email, "@")
	_, dom := components[0], components[1]

	return strings.Contains(dom, domain)
}

func SignUp() gin.HandlerFunc {
	return func(c *gin.Context) {
		var user models.User

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

		// save user in db
		e = get_user_connector.SetUser(user).Add()
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
		var verification models.Verification

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

		v_k := *user.V_key
		// check whether verification code is correct
		if v_k != *verification.V_key {
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
		var user models.Login
		var foundUser models.User

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
		passwordIsValid := helper.VerifyPassword(*user.Password, *foundUser.Password)
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
		token, refreshToken, err := helper.GenerateAllTokens(*foundUser.Id, *foundUser.Email, *foundUser.First_name, *foundUser.Last_name, *foundUser.User_type)
		if err != nil {
			c.JSON(err.StatusCode, err.Error())
			return
		}

		err = helper.UpdateAllTokens(token, refreshToken, *foundUser.Email)
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

		refreshToken := foundUser.Refresh_token
		if refreshToken == nil {
			err := helpers.RaiseLoginExpired()
			c.JSON(err.StatusCode, err.Error())
			return
		}

		_, err := helper.ValidateToken(*refreshToken)
		if err != nil {
			c.JSON(err.StatusCode, err.Error())
			return
		}

		// if refresh is still valid then generate new token
		token, _, err := helper.GenerateAllTokens(*foundUser.Id, *foundUser.Email, *foundUser.First_name, *foundUser.Last_name, *foundUser.User_type)
		if err != nil {
			c.JSON(err.StatusCode, err.Error())
			return
		}

		err = helper.UpdateAllTokens(token, *refreshToken, email)
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
		foundUser.Refresh_token = nil

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

func GetUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		profile := make(map[string]interface{})
		user := c.Param("user")
		// get current user
		if user == "" {
			profile["id"] = c.GetString("id")
			profile["first_name"] = c.GetString("first_name")
			profile["last_name"] = c.GetString("last_name")
			c.JSON(http.StatusOK, profile)
			return
		}
		user_id, _ := strconv.Atoi(user)
		// get other user
		get_user_connector := user_connector.Init()
		foundUser, e := get_user_connector.SetUserId(user_id).Get()

		if e != nil {
			err := helpers.RaiseDatabaseError()
			c.JSON(err.StatusCode, err.Error())
			return
		}
		profile["id"] = foundUser.Id
		profile["first_name"] = foundUser.First_name
		profile["last_name"] = foundUser.Last_name
		c.JSON(http.StatusOK, profile)
	}
}
