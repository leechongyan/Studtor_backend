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
		req := models.Paginated_query{}
		err := helpers.ExtractGetRequestBody(c, &req)
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
		req := models.Paginated_query{}
		err := helpers.ExtractGetRequestBody(c, &req)
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

func GetAllTutorsForACourse() gin.HandlerFunc {
	return func(c *gin.Context) {
		var course_query models.Course_query

		err := helpers.ExtractGetRequestBody(c, &course_query)
		if err != nil {
			c.JSON(err.StatusCode, err.Error())
			return
		}

		tutors, e := database_service.CurrentDatabaseConnector.GetAllTutorsForACourse(course_query.Course_code, course_query.From, course_query.Size)
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
		var slot_query models.TimeFrame_query

		err := helpers.ExtractPostRequestBody(c, &slot_query)
		if err != nil {
			c.JSON(err.StatusCode, err.Error())
			return
		}

		e := database_service.CurrentDatabaseConnector.SaveTutorAvailableTimes(slot_query)
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

		e := database_service.CurrentDatabaseConnector.DeleteTutorAvailableTimes(slot_query)
		if e != nil {
			err := helpers.RaiseDatabaseError()
			c.JSON(err.StatusCode, err.Error())
			return
		}

		c.JSON(http.StatusOK, "Success")
	}
}

func GetAllBookedTimeTutor() gin.HandlerFunc {
	return func(c *gin.Context) {
		var slot_query models.TimeFrame_query

		err := helpers.ExtractPostRequestBody(c, &slot_query)
		if err != nil {
			c.JSON(err.StatusCode, err.Error())
			return
		}

		booked, e := database_service.CurrentDatabaseConnector.GetTutorBookedTimes(slot_query)
		if e != nil {
			err := helpers.RaiseDatabaseError()
			c.JSON(err.StatusCode, err.Error())
			return
		}

		c.JSON(http.StatusOK, booked)
	}
}

func GetAvailableTimeTutor() gin.HandlerFunc {
	return func(c *gin.Context) {
		var slot_query models.TimeFrame_query

		err := helpers.ExtractPostRequestBody(c, &slot_query)
		if err != nil {
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

func BookTimeTutor() gin.HandlerFunc {
	return func(c *gin.Context) {
		var slot_query models.TimeFrame_query

		err := helpers.ExtractPostRequestBody(c, &slot_query)
		if err != nil {
			c.JSON(err.StatusCode, err.Error())
			return
		}

		e := database_service.CurrentDatabaseConnector.BookTutorTime(c.GetString("email"), slot_query)
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

		e := database_service.CurrentDatabaseConnector.UnBookTutorTime(c.GetString("email"), slot_query)
		if e != nil {
			err := helpers.RaiseDatabaseError()
			c.JSON(err.StatusCode, err.Error())
			return
		}

		c.JSON(http.StatusOK, "Success")
	}
}

func GetAllBookedTimeStudent() gin.HandlerFunc {
	return func(c *gin.Context) {
		var slot_query models.TimeFrame_query

		err := helpers.ExtractPostRequestBody(c, &slot_query)
		if err != nil {
			c.JSON(err.StatusCode, err.Error())
			return
		}

		booked, e := database_service.CurrentDatabaseConnector.GetStudentBookedTimes(slot_query)
		if e != nil {
			err := helpers.RaiseDatabaseError()
			c.JSON(err.StatusCode, err.Error())
			return
		}

		c.JSON(http.StatusOK, booked)
	}
}
