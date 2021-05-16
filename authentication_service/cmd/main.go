package main

import (
	"fmt"
	routes "github.com/leechongyan/Studtor_backend/authentication_service/routes"
	"github.com/spf13/viper"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.New()
	router.Use(gin.Logger())
	InitializeViper()

	authorized  := router.Group("/")
	routes.AuthRoutes(authorized)

	home := router.Group("/home")
	routes.UserRoutes(home)

	router.Run(viper.GetString("port"))
}

func InitializeViper() {
	// Set the file name of the configurations file
	viper.SetConfigName("config")

	// Set the path to look for the configurations file
	viper.AddConfigPath("../")

	// Enable VIPER to read Environment Variables
	viper.AutomaticEnv()

	viper.SetConfigType("yml")
	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file, %s", err)
	}
}