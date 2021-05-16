package routes  

import (
	handler "github.com/leechongyan/Studtor_backend/authentication_service/controllers"

	"github.com/gin-gonic/gin"
	"github.com/leechongyan/Studtor_backend/authentication_service/middleware"
)


// Logged In User Routes
func UserRoutes(incomingRoutes *gin.RouterGroup) {
	incomingRoutes.Use(middleware.Authentication())
	incomingRoutes.GET("/", handler.GetMain())
	// incomingRoutes.GET("/users/:user_id", handler.GetUser())
}