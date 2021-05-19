package controllers

import (
	"strings"

	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"

	helper "github.com/leechongyan/Studtor_backend/authentication_service/helpers/account"
	"github.com/leechongyan/Studtor_backend/authentication_service/models"
	"github.com/leechongyan/Studtor_backend/database_service"
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
		_, e := database_service.CurrentDatabaseConnector.GetUser(*user.Email)
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
		e = database_service.CurrentDatabaseConnector.SaveUser(user)
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
		var verification models.Verifiation

		err := helpers.ExtractPostRequestBody(c, &verification)
		if err != nil {
			c.JSON(err.StatusCode, err.Error())
			return
		}

		user, e := database_service.CurrentDatabaseConnector.GetUser(*verification.Email)

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
		e = database_service.CurrentDatabaseConnector.SaveUser(user)
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

		foundUser, e := database_service.CurrentDatabaseConnector.GetUser(*user.Email)
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
		token, refreshToken, err := helper.GenerateAllTokens(*foundUser.Email, *foundUser.First_name, *foundUser.Last_name, *foundUser.User_type)
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

		foundUser, e := database_service.CurrentDatabaseConnector.GetUser(email)
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
		token, _, err := helper.GenerateAllTokens(*foundUser.Email, *foundUser.First_name, *foundUser.Last_name, *foundUser.User_type)
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

		foundUser, e := database_service.CurrentDatabaseConnector.GetUser(email)
		// check whether user exists
		if e != nil {
			err := helpers.RaiseWrongLoginCredentials()
			c.JSON(err.StatusCode, err.Error())
			return
		}

		// remove refresh token and force user to login again
		foundUser.Refresh_token = nil

		e = database_service.CurrentDatabaseConnector.SaveUser(foundUser)

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

//GetUser is the api used to get a single user
// func GetUser() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		userId := c.Param("user_id")

// 		if err := helper.MatchUserTypeToUid(c, userId); err != nil {
// 			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 			return
// 		}
// 		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

// 		var user models.User

// 		err := userCollection.FindOne(ctx, bson.M{"user_id": userId}).Decode(&user)
// 		defer cancel()
// 		if err != nil {
// 			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 			return
// 		}

// 		c.JSON(http.StatusOK, user)

// 	}
// }
