package main

import (
	"github.com/gin-gonic/gin"
	authhandler "github.com/leechongyan/Studtor_backend/authentication_service/controllers"
	"github.com/leechongyan/Studtor_backend/authentication_service/middleware"
	"github.com/leechongyan/Studtor_backend/database_service"
	"github.com/leechongyan/Studtor_backend/helpers"
	tuthandler "github.com/leechongyan/Studtor_backend/tuition_service/controllers"
	"github.com/spf13/viper"
)

func main() {
	router := gin.New()
	router.Use(gin.Logger())

	err := helpers.InitializeViper()
	database_service.InitDatabase()

	if err != nil {
		return
	}

	// current version is v1
	v1 := router.Group("/v1")

	// does not require token
	authorized := v1.Group("/")
	authorized.POST("/signup", authhandler.SignUp())
	authorized.POST("/verify", authhandler.Verify())
	authorized.POST("/login", authhandler.Login())
	authorized.POST("/refresh", authhandler.RefreshToken())
	authorized.POST("/logout", authhandler.Logout())

	// require token
	home := v1.Group("/home")
	home.Use(middleware.Authentication())

	// ping to validate authorized status
	home.GET("/", authhandler.GetMain())

	// for general usage
	home.GET("/getallcourses", tuthandler.GetAllCourses())

	// for tutor usage
	home.POST("/putavailabletimetutor", tuthandler.PutAvailableTimeTutor())
	home.POST("/deleteavailabletimetutor", tuthandler.DeleteAvailableTimeTutor())
	home.POST("/getallbookedtimetutor", tuthandler.GetAllBookedTimeTutor())

	// for student usage
	home.GET("/getalltutors", tuthandler.GetAllTutors())
	home.GET("/getalltutorsforacourse", tuthandler.GetAllTutorsForACourse())
	home.POST("/getavailabletimetutor", tuthandler.GetAvailableTimeTutor())
	home.POST("/booktimetutor", tuthandler.BookTimeTutor())
	home.POST("/unbooktimetutor", tuthandler.UnbookTimeTutor())
	home.POST("/getallbookedtimestudent", tuthandler.GetAllBookedTimeStudent())
	// end of version v1

	router.Run(viper.GetString("port"))
}
