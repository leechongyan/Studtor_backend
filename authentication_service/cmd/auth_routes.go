package main

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	handler "github.com/leechongyan/Studtor_backend/authentication_service/internal"
)

func main() {
	router := gin.Default()
	handler.InitializeViper()

	router.POST("/login", handler.LoginHandler)
	router.POST("/refresh", handler.RefreshHandler)
	// router.GET("/", handler.ValidateHandler)

	router.Run(viper.GetString("port"))
}