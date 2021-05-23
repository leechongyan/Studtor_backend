package main

import (
	"github.com/gin-gonic/gin"
	authhandler "github.com/leechongyan/Studtor_backend/authentication_service/controllers"
	"github.com/leechongyan/Studtor_backend/authentication_service/middleware"
	database_service "github.com/leechongyan/Studtor_backend/database_service/controller"
	"github.com/leechongyan/Studtor_backend/helpers"
	"github.com/leechongyan/Studtor_backend/storage_service"
	tuthandler "github.com/leechongyan/Studtor_backend/tuition_service/controllers"
	"github.com/spf13/viper"
)

func main() {
	router := gin.New()
	router.Use(gin.Logger())

	err := helpers.InitializeViper()
	err = database_service.InitDatabase()
	err = storage_service.InitStorage()

	if err != nil {
		return
	}

	// current version is v1
	v1 := router.Group("/v1")

	// does not require token
	authorized := v1.Group("/auth")
	authorized.POST("/signup", authhandler.SignUp())
	authorized.POST("/verify", authhandler.Verify())
	authorized.POST("/login", authhandler.Login())
	authorized.POST("/refresh", authhandler.RefreshToken())
	authorized.POST("/logout", authhandler.Logout())

	// require token
	home := v1.Group("/")
	home.Use(middleware.Authentication())

	// ping to validate authorized status
	home.GET("/", authhandler.GetMain())

	// for general usage
	home.GET("/user/*user", authhandler.GetUser())
	home.GET("/courses", tuthandler.GetAllCourses())

	// for tutor usage
	home.POST("/putavailabletime", tuthandler.PutAvailableTimeTutor())
	home.POST("/deleteavailabletime", tuthandler.DeleteAvailableTimeTutor())

	// for student usage
	home.GET("/tutors/*course", tuthandler.GetAllTutors())
	home.GET("/availabletime/:tutor", tuthandler.GetAvailableTimeTutor())
	home.POST("/book", tuthandler.BookTimeTutor())
	home.POST("/unbook", tuthandler.UnbookTimeTutor())
	home.GET("/bookedtime/*user", tuthandler.GetAllBookedTime())
	// end of version v1

	router.Run(viper.GetString("port"))
}
