package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	availabilityConnector "github.com/leechongyan/Studtor_backend/database_service/connector/availability_connector"
	bookingConnector "github.com/leechongyan/Studtor_backend/database_service/connector/booking_connector"
	errorHelper "github.com/leechongyan/Studtor_backend/helpers/error_helpers"
	httpHelper "github.com/leechongyan/Studtor_backend/helpers/http_helpers"
	"github.com/leechongyan/Studtor_backend/tuition_service/models"
)

func PutAvailableTimeTutor() gin.HandlerFunc {
	return func(c *gin.Context) {
		var slotQuery models.TimeSlot

		err := httpHelper.ExtractPostRequestBody(c, &slotQuery)
		if err != nil {
			c.JSON(err.StatusCode, err.Error())
			return
		}
		tutorId, _ := strconv.Atoi(c.GetString("id"))

		availabilityConnector := availabilityConnector.Init()
		e := availabilityConnector.SetTutorId(tutorId).SetFromTime(slotQuery.From).SetToTime(slotQuery.To).Add()

		if e != nil {
			err := errorHelper.RaiseDatabaseError()
			c.JSON(err.StatusCode, err.Error())
			return
		}

		c.JSON(http.StatusOK, "Success")
	}
}

func DeleteAvailableTimeTutor() gin.HandlerFunc {
	return func(c *gin.Context) {
		var slotQuery models.Availability

		err := httpHelper.ExtractPostRequestBody(c, &slotQuery)
		if err != nil {
			c.JSON(err.StatusCode, err.Error())
			return
		}
		availabilityConnector := availabilityConnector.Init()

		e := availabilityConnector.SetAvailabilityId(*slotQuery.AvailabilityId).Delete()

		if e != nil {
			err := errorHelper.RaiseDatabaseError()
			c.JSON(err.StatusCode, err.Error())
			return
		}

		c.JSON(http.StatusOK, "Success")
	}
}

// TODO:
func GetAvailableTimeTutor() gin.HandlerFunc {
	return func(c *gin.Context) {
		var slotQuery models.TimePaginatedQuery

		err := httpHelper.ExtractPostRequestBody(c, &slotQuery)
		if err != nil {
			c.JSON(err.StatusCode, err.Error())
			return
		}

		tutorId := c.Param("tutor_id")
		tutId, e := strconv.Atoi(tutorId)
		if e != nil {
			err := errorHelper.RaiseCannotParseRequest()
			c.JSON(err.StatusCode, err.Error())
			return
		}

		availabilityConnector := availabilityConnector.Init()
		availabilityConnector.SetTutorId(tutId)

		times, e := availabilityConnector.SetFromTime(slotQuery.From).SetToTime(slotQuery.To).Get()

		if e != nil {
			err := errorHelper.RaiseDatabaseError()
			c.JSON(err.StatusCode, err.Error())
			return
		}

		c.JSON(http.StatusOK, times)
	}
}

func BookTimeTutor() gin.HandlerFunc {
	return func(c *gin.Context) {
		var slotQuery models.BookSlot

		err := httpHelper.ExtractPostRequestBody(c, &slotQuery)
		if err != nil {
			c.JSON(err.StatusCode, err.Error())
			return
		}

		bookingConnector := bookingConnector.Init()

		id, _ := strconv.Atoi(c.GetString("id"))

		e := bookingConnector.SetCourseId(*slotQuery.Course).SetStudentId(id).SetAvailabilityId(*slotQuery.AvailabilityId).Add()

		if e != nil {
			err := errorHelper.RaiseDatabaseError()
			c.JSON(err.StatusCode, err.Error())
			return
		}

		c.JSON(http.StatusOK, "Success")
	}
}

func UnbookTimeTutor() gin.HandlerFunc {
	return func(c *gin.Context) {
		var slotQuery models.BookSlot

		err := httpHelper.ExtractPostRequestBody(c, &slotQuery)
		if err != nil {
			c.JSON(err.StatusCode, err.Error())
			return
		}

		bookingConnector := bookingConnector.Init()

		id, _ := strconv.Atoi(c.GetString("id"))

		e := bookingConnector.SetCourseId(*slotQuery.Course).SetStudentId(id).SetAvailabilityId(*slotQuery.AvailabilityId).Delete()

		if e != nil {
			err := errorHelper.RaiseDatabaseError()
			c.JSON(err.StatusCode, err.Error())
			return
		}

		c.JSON(http.StatusOK, "Success")
	}
}

func GetAllBookedTime() gin.HandlerFunc {
	return func(c *gin.Context) {
		var slotQuery models.TimePaginatedQuery

		err := httpHelper.ExtractPostRequestBody(c, &slotQuery)
		if err != nil {
			c.JSON(err.StatusCode, err.Error())
			return
		}

		userId := c.Param("user_id")
		id, e := strconv.Atoi(userId)
		if e != nil {
			err := errorHelper.RaiseCannotParseRequest()
			c.JSON(err.StatusCode, err.Error())
			return
		}

		bookingConnector := bookingConnector.Init()

		times, e := bookingConnector.SetUserId(id).SetFromTime(slotQuery.From).SetToTime(slotQuery.To).Get()

		if e != nil {
			err := errorHelper.RaiseDatabaseError()
			c.JSON(err.StatusCode, err.Error())
			return
		}

		c.JSON(http.StatusOK, times)
	}
}
