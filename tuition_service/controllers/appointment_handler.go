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
		aId := c.Param("availability_id")
		availabilityId, e := strconv.Atoi(aId)
		if e != nil {
			err := errorHelper.RaiseCannotParseRequest()
			c.JSON(err.StatusCode, err.Error())
			return
		}
		// TODO Check whether the availability id belongs to the current user

		availabilityConnector := availabilityConnector.Init()

		e := availabilityConnector.SetAvailabilityId(availabilityId).Delete()

		if e != nil {
			err := errorHelper.RaiseDatabaseError()
			c.JSON(err.StatusCode, err.Error())
			return
		}

		c.JSON(http.StatusOK, "Success")
	}
}

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
		aId := c.Param("availability_id")
		availabilityId, e := strconv.Atoi(aId)
		if e != nil {
			err := errorHelper.RaiseCannotParseRequest()
			c.JSON(err.StatusCode, err.Error())
			return
		}

		cId := c.Param("course_id")
		courseId, e := strconv.Atoi(cId)
		if e != nil {
			err := errorHelper.RaiseCannotParseRequest()
			c.JSON(err.StatusCode, err.Error())
			return
		}

		bookingConnector := bookingConnector.Init()

		id, _ := strconv.Atoi(c.GetString("id"))

		e = bookingConnector.SetCourseId(courseId).SetStudentId(id).SetAvailabilityId(availabilityId).Add()

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
		bId := c.Param("booking_id")
		bookingId, e := strconv.Atoi(bId)
		if e != nil {
			err := errorHelper.RaiseCannotParseRequest()
			c.JSON(err.StatusCode, err.Error())
			return
		}
		// TODO Check whether the booking id belongs to the tutor or the student

		bookingConnector := bookingConnector.Init()

		e = bookingConnector.SetBookingId(bookingId).Delete()

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

		err := httpHelper.ExtractGetRequestBody(c, &slotQuery)
		if err != nil {
			c.JSON(err.StatusCode, err.Error())
			return
		}

		bookingConnector := bookingConnector.Init()

		if slotQuery.IsStudent == nil || *slotQuery.IsStudent {
			// if is a student, use current student id
			// default is get user
			studentId, _ := strconv.Atoi(c.GetString("id"))
			bookingConnector.SetUserId(studentId)
		} else {
			// get the tutor
			tutorId, _ := strconv.Atoi(c.Param("tutor_id"))
			bookingConnector.SetUserId(tutorId)
		}

		times, e := bookingConnector.SetFromTime(slotQuery.From).SetToTime(slotQuery.To).Get()

		if e != nil {
			err := errorHelper.RaiseDatabaseError()
			c.JSON(err.StatusCode, err.Error())
			return
		}

		c.JSON(http.StatusOK, times)
	}
}
