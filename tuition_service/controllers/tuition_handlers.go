package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/leechongyan/Studtor_backend/database_service"
	"github.com/leechongyan/Studtor_backend/helpers"
	"github.com/leechongyan/Studtor_backend/tuition_service/models"
)

func GetAllCourses() gin.HandlerFunc {
	return func(c *gin.Context) {
		var query models.ObjectPaginated_query

		err := helpers.ExtractGetRequestBody(c, &query)
		if err != nil {
			c.JSON(err.StatusCode, err.Error())
			return
		}

		var db_options database_service.DB_options
		db_options.From_id = query.From_id
		db_options.Size = query.Size

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
		var query models.ObjectPaginated_query

		err := helpers.ExtractGetRequestBody(c, &query)
		if err != nil {
			c.JSON(err.StatusCode, err.Error())
			return
		}

		var db_options database_service.DB_options
		db_options.From_id = query.From_id
		db_options.Size = query.Size

		// whether you are getting the tutor for a certain course
		course := c.Param("course")
		if course != "" {
			db_options.Course = &course
		}

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
		var slot_query models.TimeSlot

		err := helpers.ExtractPostRequestBody(c, &slot_query)
		if err != nil {
			c.JSON(err.StatusCode, err.Error())
			return
		}
		tutor_id := c.GetString("id")

		var db_options database_service.DB_options
		db_options.Tutor_id = &tutor_id
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
		var slot_query models.TimeSlot

		err := helpers.ExtractPostRequestBody(c, &slot_query)
		if err != nil {
			c.JSON(err.StatusCode, err.Error())
			return
		}
		tutor_id := c.GetString("id")

		var db_options database_service.DB_options
		db_options.Tutor_id = &tutor_id
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

// TODO:
func GetAvailableTimeTutor() gin.HandlerFunc {
	return func(c *gin.Context) {
		var slot_query models.TimePaginated_query

		err := helpers.ExtractPostRequestBody(c, &slot_query)
		if err != nil {
			c.JSON(err.StatusCode, err.Error())
			return
		}

		tutor_id := c.Param("tutor")

		var db_options database_service.DB_options
		db_options.Tutor_id = &tutor_id
		db_options.From_time = slot_query.From
		db_options.To_time = slot_query.To

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
		var slot_query models.BookSlot

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
		db_options.Student_id = &student_id
		db_options.Tutor_id = slot_query.Tutor

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
		var slot_query models.BookSlot

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
		db_options.Student_id = &student_id
		db_options.Tutor_id = slot_query.Tutor

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
		var slot_query models.TimePaginated_query

		err := helpers.ExtractPostRequestBody(c, &slot_query)
		if err != nil {
			c.JSON(err.StatusCode, err.Error())
			return
		}

		user := c.Param("user")

		var db_options database_service.DB_options
		db_options.User_id = &user
		db_options.From_time = slot_query.From
		db_options.To_time = slot_query.To

		booked, e := database_service.CurrentDatabaseConnector.GetBookedTimes(db_options)
		if e != nil {
			err := helpers.RaiseDatabaseError()
			c.JSON(err.StatusCode, err.Error())
			return
		}

		c.JSON(http.StatusOK, booked)
	}
}
