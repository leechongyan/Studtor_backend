package main

import (
	routes "github.com/leechongyan/Studtor_backend/routes"
	"github.com/leechongyan/Studtor_backend/helpers"
	"github.com/spf13/viper"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.New()
	router.Use(gin.Logger())
	
	err := helpers.InitializeViper()
	if err != nil {
		return 
	}

	// current version is v1
	v1 := router.Group("/v1")

	authorized  := v1.Group("/")
	routes.AuthRoutes(authorized)

	home := v1.Group("/home")
	routes.UserRoutes(home)

	router.Run(viper.GetString("port"))
}
