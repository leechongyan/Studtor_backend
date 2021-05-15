package controllers

import (
	"fmt"
	"log"

	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"

	helper "github.com/leechongyan/Studtor_backend/authentication_service/helpers/account"
	"github.com/leechongyan/Studtor_backend/authentication_service/models"
	"github.com/leechongyan/Studtor_backend/authentication_service/database"

	"golang.org/x/crypto/bcrypt"
)





var validate = validator.New()

//HashPassword is used to encrypt the password before it is stored in the DB
func HashPassword(password string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)

	if err != nil {
		log.Panic(err)
	}

	return string(bytes)
}

//VerifyPassword checks the input password while verifying it with the passward in the DB.
func VerifyPassword(userPassword string, providedPassword string) (bool, string) {
	err := bcrypt.CompareHashAndPassword([]byte(providedPassword), []byte(userPassword))

	if err != nil {
		return false, fmt.Sprintf("login or passowrd is incorrect")
	}

	return true, ""
}

func SignUp() gin.HandlerFunc {
	return func(c *gin.Context) {
		var user models.User

		if err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		validationErr := validate.Struct(user)
		if validationErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
			return
		}

		// check whether this email exist 
		_, ok := database.UserCollection[*user.Email]

		if ok {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "email already exists"})
			return
		}

		password := HashPassword(*user.Password)
		user.Password = &password

		user.Created_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		user.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		token, refreshToken, _ := helper.GenerateAllTokens(*user.Email, *user.First_name, *user.Last_name, *user.User_type)
		user.Token = &token
		user.Refresh_token = &refreshToken

		database.UserCollection[*user.Email] = user

		c.JSON(http.StatusOK, gin.H{"Success": "Successful Sign Up"})
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

		if !ok {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "email or password is incorrect"})
			return
		}

		passwordIsValid, msg := VerifyPassword(*user.Password, *foundUser.Password)
		if passwordIsValid != true {
			c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
			return
		}

		token, refreshToken, _ := helper.GenerateAllTokens(*foundUser.Email, *foundUser.First_name, *foundUser.Last_name, *foundUser.User_type)

		helper.UpdateAllTokens(token, refreshToken, *foundUser.Email)

		c.JSON(http.StatusOK, foundUser)

	}
}

func GetMain() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Print("Successful")
		c.JSON(http.StatusOK, gin.H{"Success": "Successful Sign Up"})
	}
}

// func GetUsers() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		if err := helper.CheckUserType(c, "ADMIN"); err != nil {
// 			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 			return
// 		}
// 		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

// 		// recordPerPage := 10
// 		recordPerPage, err := strconv.Atoi(c.Query("recordPerPage"))
// 		if err != nil || recordPerPage < 1 {
// 			recordPerPage = 10
// 		}

// 		page, err1 := strconv.Atoi(c.Query("page"))
// 		if err1 != nil || page < 1 {
// 			page = 1
// 		}

// 		startIndex := (page - 1) * recordPerPage
// 		startIndex, err = strconv.Atoi(c.Query("startIndex"))

// 		matchStage := bson.D{{"$match", bson.D{{}}}}
// 		groupStage := bson.D{{"$group", bson.D{{"_id", bson.D{{"_id", "null"}}}, {"total_count", bson.D{{"$sum", 1}}}, {"data", bson.D{{"$push", "$$ROOT"}}}}}}
// 		projectStage := bson.D{
// 			{"$project", bson.D{
// 				{"_id", 0},
// 				{"total_count", 1},
// 				{"user_items", bson.D{{"$slice", []interface{}{"$data", startIndex, recordPerPage}}}},
// 			}}}

// 		result, err := userCollection.Aggregate(ctx, mongo.Pipeline{
// 			matchStage, groupStage, projectStage})
// 		defer cancel()
// 		if err != nil {
// 			c.JSON(http.StatusInternalServerError, gin.H{"error": "error occured while listing user items"})
// 		}
// 		var allusers []bson.M
// 		if err = result.All(ctx, &allusers); err != nil {
// 			log.Fatal(err)
// 		}
// 		c.JSON(http.StatusOK, allusers[0])

// 	}
// }

//GetUser is the api used to tget a single user
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