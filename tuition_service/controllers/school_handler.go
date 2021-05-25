package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	schoolConnector "github.com/leechongyan/Studtor_backend/database_service/connector/school_connector"
	errorHelper "github.com/leechongyan/Studtor_backend/helpers/error_helpers"
)

func GetSchools() gin.HandlerFunc {
	return func(c *gin.Context) {
		schoolConnector := schoolConnector.Init()

		schools, e := schoolConnector.GetAll()

		if e != nil {
			err := errorHelper.RaiseDatabaseError()
			c.JSON(err.StatusCode, err.Error())
			return
		}
		c.JSON(http.StatusOK, schools)
	}
}
