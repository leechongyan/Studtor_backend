package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"

	"github.com/leechongyan/Studtor_backend/database_service"
	"github.com/leechongyan/Studtor_backend/helpers"
	"github.com/leechongyan/Studtor_backend/tuition_service/models"
)

var validate = validator.New()

func ExtractPaginationFields(c *gin.Context) (req models.Pagination, err *helpers.RequestError) {
	e := c.ShouldBind(&req)
	if e != nil {
		err = helpers.RaiseCannotParseJson()
		return
	}
	validationErr := validate.Struct(req)
	if validationErr != nil {
		err = helpers.RaiseValidationErrorJson()
		return
	}
	return req, nil
}

func GetAllCourses() gin.HandlerFunc {
	return func(c *gin.Context) {

		req, err := ExtractPaginationFields(c)
		if err != nil {
			c.JSON(err.StatusCode, err.Error())
			return
		}

		courses, e := database_service.CurrentDatabaseConnector.GetAllCourses(req.From, req.Size)
		if e != nil {
			err = helpers.RaiseDatabaseError()
			c.JSON(err.StatusCode, err.Error())
			return
		}
		c.JSON(http.StatusOK, courses)
	}
}

func GetAllTutors() gin.HandlerFunc {
	return func(c *gin.Context) {

		req, err := ExtractPaginationFields(c)
		if err != nil {
			c.JSON(err.StatusCode, err.Error())
			return
		}

		tutors, e := database_service.CurrentDatabaseConnector.GetAllTutors(req.From, req.Size)
		if e != nil {
			err = helpers.RaiseDatabaseError()
			c.JSON(err.StatusCode, err.Error())
			return
		}
		c.JSON(http.StatusOK, tutors)
	}
}

func PutAvailableTimeTutor() gin.HandlerFunc {
	return func(c *gin.Context) {
		var slot_query models.Slot_query

		e := c.BindJSON(&slot_query)
		if e != nil {
			err := helpers.RaiseCannotParseJson()
			c.JSON(err.StatusCode, err.Error())
			return
		}

		e = database_service.CurrentDatabaseConnector.SaveTutorAvailableTimes(slot_query)
		if e != nil {
			err := helpers.RaiseDatabaseError()
			c.JSON(err.StatusCode, err.Error())
			return
		}

		c.JSON(http.StatusOK, "Success")
	}
}

func GetAvailableTimeTutor() gin.HandlerFunc {
	return func(c *gin.Context) {
		var slot_query models.Slot_query

		e := c.BindJSON(&slot_query)
		if e != nil {
			err := helpers.RaiseCannotParseJson()
			c.JSON(err.StatusCode, err.Error())
			return
		}

		availability, e := database_service.CurrentDatabaseConnector.GetTutorAvailableTimes(slot_query)
		if e != nil {
			err := helpers.RaiseDatabaseError()
			c.JSON(err.StatusCode, err.Error())
			return
		}

		c.JSON(http.StatusOK, availability)
	}
}
