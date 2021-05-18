package main

import (
	"github.com/gin-gonic/gin"
	handler "github.com/leechongyan/Studtor_backend/authentication_service/controllers"
	"github.com/leechongyan/Studtor_backend/authentication_service/middleware"
	"github.com/leechongyan/Studtor_backend/database_service"
	"github.com/leechongyan/Studtor_backend/helpers"
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
	authorized.POST("/signup", handler.SignUp())
	authorized.POST("/verify", handler.Verify())
	authorized.POST("/login", handler.Login())
	authorized.POST("/refresh", handler.RefreshToken())

	// require token
	home := v1.Group("/home")
	home.Use(middleware.Authentication())
	home.GET("/", handler.GetMain())
	// end of version v1

	router.Run(viper.GetString("port"))
}
