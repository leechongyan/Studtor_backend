package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/leechongyan/Studtor_backend/database_service"
	"github.com/leechongyan/Studtor_backend/helpers"
	"github.com/leechongyan/Studtor_backend/tuition_service/models"
)

func GetAllCourses() gin.HandlerFunc {
	return func(c *gin.Context) {
		var db_options database_service.DB_options

		err := helpers.ExtractGetRequestBody(c, &db_options)
		if err != nil {
			c.JSON(err.StatusCode, err.Error())
			return
		}

		courses, e := database_service.CurrentDatabaseConnector.GetAllCourses(db_options)

		if e != nil {
			err := helpers.RaiseDatabaseError()
			c.JSON(err.StatusCode, err.Error())
			return
		}
		c.JSON(http.StatusOK, courses)
	}
}

func GetAllTutors() gin.HandlerFunc {
	return func(c *gin.Context) {
		var db_options database_service.DB_options
		err := helpers.ExtractGetRequestBody(c, &db_options)
		if err != nil {
			c.JSON(err.StatusCode, err.Error())
			return
		}

		course := c.Param("course")
		if course != "" {
			db_options.Course = &course
		}
		fmt.Print("HIT THIS")
		fmt.Print(*db_options.From_id)
		fmt.Print("WHATF")

		courses, e := database_service.CurrentDatabaseConnector.GetAllTutors(db_options)

		if e != nil {
			err := helpers.RaiseDatabaseError()
			c.JSON(err.StatusCode, err.Error())
			return
		}
		c.JSON(http.StatusOK, courses)
	}
}

func PutAvailableTimeTutor() gin.HandlerFunc {
	return func(c *gin.Context) {
		var slot_query models.TimeFrame_query

		err := helpers.ExtractPostRequestBody(c, &slot_query)
		if err != nil {
			c.JSON(err.StatusCode, err.Error())
			return
		}

		var db_options database_service.DB_options
		db_options.To_id = slot_query.Email
		db_options.From_time = slot_query.From
		db_options.To_time = slot_query.To

		e := database_service.CurrentDatabaseConnector.SaveTutorAvailableTimes(db_options)
		if e != nil {
			err := helpers.RaiseDatabaseError()
			c.JSON(err.StatusCode, err.Error())
			return
		}

		c.JSON(http.StatusOK, "Success")
	}
}

func DeleteAvailableTimeTutor() gin.HandlerFunc {
	return func(c *gin.Context) {
		var slot_query models.TimeFrame_query

		err := helpers.ExtractPostRequestBody(c, &slot_query)
		if err != nil {
			c.JSON(err.StatusCode, err.Error())
			return
		}

		var db_options database_service.DB_options
		db_options.To_id = slot_query.Email
		db_options.From_time = slot_query.From
		db_options.To_time = slot_query.To

		e := database_service.CurrentDatabaseConnector.DeleteTutorAvailableTimes(db_options)
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
		var db_options database_service.DB_options
		err := helpers.ExtractGetRequestBody(c, &db_options)
		if err != nil {
			c.JSON(err.StatusCode, err.Error())
			return
		}

		user := c.Param("tutor")

		db_options.To_id = &user

		availability, e := database_service.CurrentDatabaseConnector.GetTutorAvailableTimes(db_options)
		if e != nil {
			err := helpers.RaiseDatabaseError()
			c.JSON(err.StatusCode, err.Error())
			return
		}

		c.JSON(http.StatusOK, availability)
	}
}

func BookTimeTutor() gin.HandlerFunc {
	return func(c *gin.Context) {
		var slot_query models.TimeFrame_query

		err := helpers.ExtractPostRequestBody(c, &slot_query)
		if err != nil {
			c.JSON(err.StatusCode, err.Error())
			return
		}
		student_id := c.GetString("email")

		var db_options database_service.DB_options
		db_options.From_time = slot_query.From
		db_options.To_time = slot_query.To
		db_options.Course = slot_query.Course
		db_options.From_id = &student_id
		db_options.To_id = slot_query.Email

		e := database_service.CurrentDatabaseConnector.BookTutorTime(db_options)
		if e != nil {
			err := helpers.RaiseDatabaseError()
			c.JSON(err.StatusCode, err.Error())
			return
		}

		c.JSON(http.StatusOK, "Success")
	}
}

func UnbookTimeTutor() gin.HandlerFunc {
	return func(c *gin.Context) {
		var slot_query models.TimeFrame_query

		err := helpers.ExtractPostRequestBody(c, &slot_query)
		if err != nil {
			c.JSON(err.StatusCode, err.Error())
			return
		}
		student_id := c.GetString("email")

		var db_options database_service.DB_options
		db_options.From_time = slot_query.From
		db_options.To_time = slot_query.To
		db_options.Course = slot_query.Course
		db_options.From_id = &student_id
		db_options.To_id = slot_query.Email

		e := database_service.CurrentDatabaseConnector.UnBookTutorTime(db_options)
		if e != nil {
			err := helpers.RaiseDatabaseError()
			c.JSON(err.StatusCode, err.Error())
			return
		}

		c.JSON(http.StatusOK, "Success")
	}
}

func GetAllBookedTime() gin.HandlerFunc {
	return func(c *gin.Context) {
		var db_options database_service.DB_options
		err := helpers.ExtractGetRequestBody(c, &db_options)
		if err != nil {
			c.JSON(err.StatusCode, err.Error())
			return
		}

		user := c.Param("user")

		db_options.To_id = &user

		booked, e := database_service.CurrentDatabaseConnector.GetBookedTimes(db_options)
		if e != nil {
			err := helpers.RaiseDatabaseError()
			c.JSON(err.StatusCode, err.Error())
			return
		}

		c.JSON(http.StatusOK, booked)
	}
}
