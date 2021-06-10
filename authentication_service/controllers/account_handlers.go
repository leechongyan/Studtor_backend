package controllers

import (
	"errors"
	"strconv"
	"strings"

	"net/http"

	"github.com/gin-gonic/gin"

	authHelper "github.com/leechongyan/Studtor_backend/authentication_service/helpers/account"
	query "github.com/leechongyan/Studtor_backend/authentication_service/models"
	userModel "github.com/leechongyan/Studtor_backend/database_service/client_models"
	userConnector "github.com/leechongyan/Studtor_backend/database_service/connector/user_connector"

	httpError "github.com/leechongyan/Studtor_backend/constants/errors/http_errors"
	databaseError "github.com/leechongyan/Studtor_backend/database_service/errors"
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

func getUserAccountWithEmail(email string) (user userModel.User, err error) {
	user, err = userConnector.Init().SetUserEmail(email).GetUser()
	return
}

func getUserAccountWithID(id int) (user userModel.User, err error) {
	user, err = userConnector.Init().SetUserId(id).GetUser()
	return
}

func getUserProfileWithID(id int) (user userModel.UserProfile, err error) {
	user, err = userConnector.Init().SetUserId(id).GetProfile()
	return
}

func saveUser(user userModel.User) (err error) {
	_, err = userConnector.Init().SetUser(user).Add()
	return
}

func sendVerificationToUser(user *userModel.User) (err error) {
	newVKey := authHelper.GenerateVerificationCode(6)
	user.SetVKey(newVKey)
	// send an email
	err = mailService.CurrentMailService.SendVerificationCode(*user, newVKey)
	return err
}

func preprocessUserSignUp(c *gin.Context, user *userModel.User) {
	password := authHelper.HashPassword(*user.Password())
	user.SetPassword(password)

	// save the user profile picture
	url, err := uploadUserProfilePicture(c)
	if err == nil {
		user.SetProfilePicture(url)
	}
}

func verifyUser(user *userModel.User, verificationCode string) (err error) {
	if *user.VKey() != verificationCode {
		return httpError.ErrWrongValidation
	}
	user.SetVerified(true)
	return
}

func generateTokens(generateRefresh bool, user userModel.User) (token string, err error) {
	// refresh token
	token, refreshToken, err := authHelper.GenerateAllTokens(*user.ID(), *user.Email(), *user.Name(), *user.UserType())
	if err != nil {
		return
	}
	if !generateRefresh {
		refreshToken = *user.RefreshToken()
	}
	err = authHelper.UpdateAllTokens(token, refreshToken, *user.Email())
	return
}

func isExpiredClient(user userModel.User) (expired bool) {
	if user.RefreshToken() == nil {
		return true
	}

	_, err := authHelper.ValidateToken(*user.RefreshToken())
	return err != nil
}

func clearAllTokens(user *userModel.User) {
	user.SetToken("")
	user.SetRefreshToken("")
}

// handle sign up request
func SignUp() gin.HandlerFunc {
	return func(c *gin.Context) {
		user, err := userModel.UnmarshalJson(c)
		if err != nil {
			c.JSON(http.StatusBadRequest, err.Error())
			return
		}

		// validate whether the email is valid with edu
		if !checkEduDomain(*user.Email(), "edu") {
			c.JSON(http.StatusBadRequest, httpError.ErrInvalidEmail.Error())
			return
		}

		// check whether this email exist
		_, err = getUserAccountWithEmail(*user.Email())
		if err == nil {
			c.JSON(http.StatusBadRequest, httpError.ErrExistentAccount.Error())
			return
		}
		if err != nil && !errors.Is(databaseError.ErrNoRecordFound, err) {
			c.JSON(http.StatusInternalServerError, err.Error())
			return
		}

		// hash the user password and upload user profile picture
		preprocessUserSignUp(c, &user)

		err = sendVerificationToUser(&user)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err.Error())
			return
		}

		// save user in db
		err = saveUser(user)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err.Error())
			return
		}

		c.JSON(http.StatusOK, "Success")
	}
}

func Verify() gin.HandlerFunc {
	return func(c *gin.Context) {
		var verification query.Verification

		err := httpHelper.ExtractPostRequestBody(c, &verification)
		if err != nil {
			c.JSON(http.StatusBadRequest, err.Error())
			return
		}

		user, err := getUserAccountWithEmail(*verification.Email)
		if err != nil {
			if errors.Is(databaseError.ErrNoRecordFound, err) {
				c.JSON(http.StatusUnauthorized, httpError.ErrNonExistentAccount.Error())
				return
			}
			c.JSON(http.StatusInternalServerError, err.Error())
			return
		}

		// check whether verification code is correct
		err = verifyUser(&user, *verification.VKey)
		if err != nil {
			c.JSON(http.StatusUnauthorized, err.Error())
			return
		}

		err = saveUser(user)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err.Error())
			return
		}

		c.JSON(http.StatusOK, "Success")
	}
}

func Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		var user query.Login

		err := httpHelper.ExtractPostRequestBody(c, &user)
		if err != nil {
			c.JSON(http.StatusBadRequest, err.Error())
			return
		}

		foundUser, err := getUserAccountWithEmail(*user.Email)
		// check whether user exists
		if err != nil {
			if errors.Is(databaseError.ErrNoRecordFound, err) {
				c.JSON(http.StatusUnauthorized, httpError.ErrNonExistentAccount.Error())
				return
			}
			c.JSON(http.StatusInternalServerError, err.Error())
			return
		}

		// check whether password exists
		passwordIsValid := authHelper.VerifyPassword(*user.Password, *foundUser.Password())
		if !passwordIsValid {
			c.JSON(http.StatusUnauthorized, httpError.ErrWrongCredential.Error())
			return
		}

		// check whether is verified or not
		if !foundUser.Verified() {
			c.JSON(http.StatusUnauthorized, httpError.ErrNotVerified.Error())
			return
		}

		token, err := generateTokens(true, foundUser)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err.Error())
			return
		}

		c.JSON(http.StatusOK, token)
	}
}

func RefreshToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		var refresh query.Refresh

		err := httpHelper.ExtractPostRequestBody(c, &refresh)
		if err != nil {
			c.JSON(http.StatusBadRequest, err.Error())
			return
		}

		// check whether user exists
		foundUser, err := getUserAccountWithEmail(*refresh.Email)
		// check whether user exists
		if err != nil {
			if errors.Is(databaseError.ErrNoRecordFound, err) {
				c.JSON(http.StatusUnauthorized, httpError.ErrNonExistentAccount.Error())
				return
			}
			c.JSON(http.StatusInternalServerError, err.Error())
			return
		}

		if isExpiredClient(foundUser) {
			c.JSON(http.StatusUnauthorized, httpError.ErrExpiredRefreshToken.Error())
			return
		}

		token, err := generateTokens(false, foundUser)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err.Error())
			return
		}

		c.JSON(http.StatusOK, token)
	}
}

func Logout() gin.HandlerFunc {
	return func(c *gin.Context) {
		userId, _ := strconv.Atoi(c.GetString("id"))

		// check whether user exists
		foundUser, err := getUserAccountWithID(userId)
		// check whether user exists
		if err != nil {
			if errors.Is(databaseError.ErrNoRecordFound, err) {
				c.JSON(http.StatusUnauthorized, httpError.ErrNonExistentAccount.Error())
				return
			}
			c.JSON(http.StatusInternalServerError, err.Error())
			return
		}

		// remove refresh token and force user to login again
		clearAllTokens(&foundUser)

		err = saveUser(foundUser)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err.Error())
			return
		}

		c.JSON(http.StatusOK, "Success")
	}
}

func uploadUserProfilePicture(c *gin.Context) (url string, err error) {
	file, _, err := c.Request.FormFile("file")
	if err != nil {
		return "", httpError.ErrFileParsingFailure
	}
	defer file.Close()
	url, err = storageService.CurrentStorageConnector.SaveUserProfilePicture(c.GetString("id"), file)
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
		foundUser, err := getUserProfileWithID(userId)
		if err != nil {
			if errors.Is(databaseError.ErrNoRecordFound, err) {
				c.JSON(http.StatusUnauthorized, httpError.ErrNonExistentAccount.Error())
				return
			}
			c.JSON(http.StatusInternalServerError, err.Error())
			return
		}
		c.JSON(http.StatusOK, foundUser)
	}
}

func GetCurrentUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.GetString("id")
		userId, _ := strconv.Atoi(id)

		// check whether user exists
		foundUser, err := getUserProfileWithID(userId)
		if err != nil {
			if errors.Is(databaseError.ErrNoRecordFound, err) {
				c.JSON(http.StatusUnauthorized, httpError.ErrNonExistentAccount.Error())
				return
			}
			c.JSON(http.StatusInternalServerError, err.Error())
			return
		}

		c.JSON(http.StatusOK, foundUser)
	}
}
