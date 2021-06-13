package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	schoolConnector "github.com/leechongyan/Studtor_backend/database_service/connector/school_connector"
)

// Get all the schools and their associated school courses
func GetSchools() gin.HandlerFunc {
	return func(c *gin.Context) {
		schools, err := schoolConnector.Init().GetAll()
		if err != nil {
			c.JSON(http.StatusInternalServerError, err.Error())
			return
		}
		c.JSON(http.StatusOK, schools)
	}
}
