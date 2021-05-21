package main

import (
	"github.com/gin-gonic/gin"
	AuthHandler "github.com/leechongyan/Studtor_backend/authentication_service/controllers"
	"github.com/leechongyan/Studtor_backend/database_service"
	"github.com/leechongyan/Studtor_backend/helpers"
	TuitionHandler "github.com/leechongyan/Studtor_backend/tuition_service/controllers"
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
	v1 := router.Group("/api/v1")

	// does not require token
	AuthHandler.InitAuthRouter(v1)
	TuitionHandler.InitTuitionRouter(v1)
	router.Run(viper.GetString("port"))
}
