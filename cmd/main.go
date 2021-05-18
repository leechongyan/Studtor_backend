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

	// require token
	home := v1.Group("/home")
	home.Use(middleware.Authentication())
	home.GET("/", authhandler.GetMain())
	home.GET("/getallcourses", tuthandler.GetAllCourses())
	home.GET("/getalltutors", tuthandler.GetAllTutors())
	// end of version v1

	router.Run(viper.GetString("port"))
}
