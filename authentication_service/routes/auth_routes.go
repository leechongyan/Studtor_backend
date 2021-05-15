package routes

import (
	handler "github.com/leechongyan/Studtor_backend/authentication_service/controllers"

	"github.com/gin-gonic/gin"
)

//UserRoutes function
func AuthRoutes(incomingRoutes *gin.RouterGroup) {
	incomingRoutes.POST("/signup", handler.SignUp())
	// incomingRoutes.POST("/verify", handler.VerifyHandler())
	incomingRoutes.POST("/login", handler.Login())
}