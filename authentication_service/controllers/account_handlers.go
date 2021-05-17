package controllers

import (
	"strings"

	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"

	"github.com/leechongyan/Studtor_backend/authentication_service/database"
	helper "github.com/leechongyan/Studtor_backend/authentication_service/helpers/account"
	"github.com/leechongyan/Studtor_backend/authentication_service/models"
	"github.com/leechongyan/Studtor_backend/helpers"
	"github.com/leechongyan/Studtor_backend/mail_service"
)

var validate = validator.New()

func CheckEmailDomain(email string, domain string) bool {
	components := strings.Split(email, "@")
	_, dom := components[0], components[1]

	return strings.Contains(dom, domain)
}

func SignUp() gin.HandlerFunc {
	return func(c *gin.Context) {
		var user models.User

		e := c.BindJSON(&user)
		if e != nil {
			err := helpers.RaiseCannotParseJson()
			c.JSON(err.StatusCode, err.Error())
			return
		}
		validationErr := validate.Struct(user)
		if validationErr != nil {
			err := helpers.RaiseValidationErrorJson()
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
		_, ok := database.UserCollection[*user.Email]
		if ok {
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
		database.UserCollection[*user.Email] = user

		// send an email
		err := mail_service.SendVerificationCode(user, new_V_key)
		if err != nil {
			c.JSON(err.StatusCode, err.Error())
		}

		c.JSON(http.StatusOK, gin.H{"Success": "Successful Sign Up"})
	}
}

func Verify() gin.HandlerFunc {
	return func(c *gin.Context) {
		var verification models.Verifiation

		if err := c.BindJSON(&verification); err != nil {
			err := helpers.RaiseCannotParseJson()
			c.JSON(err.StatusCode, err.Error())
			return
		}

		user, ok := database.UserCollection[*verification.Email]
		if !ok {
			err := helpers.RaiseNonExistentAccount()
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
		token, refreshToken, err := helper.GenerateAllTokens(*user.Email, *user.First_name, *user.Last_name, *user.User_type)
		if err != nil {
			c.JSON(err.StatusCode, err.Error())
			return
		}

		err = helper.UpdateAllTokens(token, refreshToken, *user.Email)
		if err != nil {
			c.JSON(err.StatusCode, err.Error())
			return
		}

		c.JSON(http.StatusOK, gin.H{"Success": "Verified"})
	}
}

func Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		var user models.User
		var foundUser models.User

		if err := c.BindJSON(&user); err != nil {
			err := helpers.RaiseCannotParseJson()
			c.JSON(err.StatusCode, err.Error())
			return
		}

		foundUser, ok := database.UserCollection[*user.Email]
		// check whether user exists
		if !ok {
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

		c.JSON(http.StatusOK, foundUser)
	}
}

func GetMain() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"Success": "Successful Entry"})
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
