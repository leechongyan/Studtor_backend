package controllers

import (
	"strings"

	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"

	helper "github.com/leechongyan/Studtor_backend/authentication_service/helpers/account"
	"github.com/leechongyan/Studtor_backend/authentication_service/models"
	"github.com/leechongyan/Studtor_backend/authentication_service/database"
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

		err := c.BindJSON(&user)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		validationErr := validate.Struct(user)
		if validationErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
			return
		}
		// validate whether the email is valid with edu
		if !CheckEmailDomain(*user.Email, "edu") {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "email is not valid"})
			return
		}

		// check whether this email exist 
		_, ok := database.UserCollection[*user.Email]
		if ok {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "email already exists"})
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
		err = mail_service.SendVerificationCode(user, new_V_key)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		}
		c.JSON(http.StatusOK, gin.H{"Success": "Successful Sign Up"})
	}
}

func Verify() gin.HandlerFunc {
	return func(c *gin.Context){
		var verification models.Verifiation

		if err := c.BindJSON(&verification); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		user, ok := database.UserCollection[*verification.Email]
		if !ok {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "account does not exist"})
			return
		}

		v_k := *user.V_key
		// check whether verification code is correct
		if v_k != *verification.V_key {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "wrong validation code"})
			return
		}

		// if verification all pass and correct then create access token
		token, refreshToken, err := helper.GenerateAllTokens(*user.Email, *user.First_name, *user.Last_name, *user.User_type)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}

		err = helper.UpdateAllTokens(token, refreshToken, *user.Email)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}

		c.JSON(http.StatusOK, gin.H{"Success": "Verified"})
	}
}

func Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		var user models.User
		var foundUser models.User

		if err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		foundUser, ok := database.UserCollection[*user.Email]
		// check whether user exists
		if !ok {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "email or password is incorrect"})
			return
		}

		// check whether password exists 
		passwordIsValid, msg := helper.VerifyPassword(*user.Password, *foundUser.Password)
		if passwordIsValid != true {
			c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
			return
		}

		// refresh token
		token, refreshToken, err := helper.GenerateAllTokens(*foundUser.Email, *foundUser.First_name, *foundUser.Last_name, *foundUser.User_type)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		
		err = helper.UpdateAllTokens(token, refreshToken, *foundUser.Email)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
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