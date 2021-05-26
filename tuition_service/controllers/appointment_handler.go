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

// Put tutor available time
func PutAvailableTimeTutor() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.GetString("id") != c.Param("tutor_id") {
			err := errorHelper.RaiseUnauthorizedAccess()
			c.JSON(err.StatusCode, err.Error())
			return
		}

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

// Remove tutor available time
func DeleteAvailableTimeTutor() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.GetString("id") != c.Param("tutor_id") {
			err := errorHelper.RaiseUnauthorizedAccess()
			c.JSON(err.StatusCode, err.Error())
			return
		}

		availabilityId, e := strconv.Atoi(c.Param("availability_id"))
		if e != nil {
			err := errorHelper.RaiseCannotParseRequest()
			c.JSON(err.StatusCode, err.Error())
			return
		}

		tutorId, _ := strconv.Atoi(c.GetString("id"))
		availabilityConnector := availabilityConnector.Init()

		// tutor id is needed to check whether the availabilityid belongs to the tutor id
		e = availabilityConnector.SetTutorId(tutorId).SetAvailabilityId(availabilityId).Delete()

		if e != nil {
			err := errorHelper.RaiseDatabaseError()
			c.JSON(err.StatusCode, err.Error())
			return
		}

		c.JSON(http.StatusOK, "Success")
	}
}

// Get a tutor available times with pagination
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

		times, e := availabilityConnector.SetTutorId(tutId).SetFromTime(slotQuery.From).SetToTime(slotQuery.To).GetAll()

		if e != nil {
			err := errorHelper.RaiseDatabaseError()
			c.JSON(err.StatusCode, err.Error())
			return
		}

		c.JSON(http.StatusOK, times)
	}
}

// Book an available timeslot with the tutor with the course id
func BookTimeTutor() gin.HandlerFunc {
	return func(c *gin.Context) {
		availabilityId, e := strconv.Atoi(c.Param("availability_id"))
		if e != nil {
			err := errorHelper.RaiseCannotParseRequest()
			c.JSON(err.StatusCode, err.Error())
			return
		}

		courseId, e := strconv.Atoi(c.Param("course_id"))
		if e != nil {
			err := errorHelper.RaiseCannotParseRequest()
			c.JSON(err.StatusCode, err.Error())
			return
		}

		bookingConnector := bookingConnector.Init()

		id, _ := strconv.Atoi(c.GetString("id"))

		e = bookingConnector.SetCourseId(courseId).SetUserId(id).SetAvailabilityId(availabilityId).Add()

		if e != nil {
			err := errorHelper.RaiseDatabaseError()
			c.JSON(err.StatusCode, err.Error())
			return
		}

		c.JSON(http.StatusOK, "Success")
	}
}

// Remove a booking
func UnbookTimeTutor() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.GetString("id") != c.Param("user_id") {
			err := errorHelper.RaiseUnauthorizedAccess()
			c.JSON(err.StatusCode, err.Error())
			return
		}

		bookingId, e := strconv.Atoi(c.Param("booking_id"))
		if e != nil {
			err := errorHelper.RaiseCannotParseRequest()
			c.JSON(err.StatusCode, err.Error())
			return
		}
		userId, _ := strconv.Atoi(c.GetString("id"))

		bookingConnector := bookingConnector.Init()

		// user id is needed to check whether the bookingid involves the user
		e = bookingConnector.SetUserId(userId).SetBookingId(bookingId).Delete()
		if e == nil {
			c.JSON(http.StatusOK, "Success")
			return
		}

		c.JSON(http.StatusOK, "Success")
	}
}

// Get all the booked time of a user with pagination
func GetAllBookedTime() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.GetString("id") != c.Param("user_id") {
			err := errorHelper.RaiseUnauthorizedAccess()
			c.JSON(err.StatusCode, err.Error())
			return
		}

		var slotQuery models.TimePaginatedQuery
		err := httpHelper.ExtractGetRequestBody(c, &slotQuery)
		if err != nil {
			c.JSON(err.StatusCode, err.Error())
			return
		}

		userId, _ := strconv.Atoi(c.GetString("id"))

		bookingConnector := bookingConnector.Init()

		times, e := bookingConnector.SetUserId(userId).SetFromTime(slotQuery.From).SetToTime(slotQuery.To).GetAll()

		if e != nil {
			err := errorHelper.RaiseDatabaseError()
			c.JSON(err.StatusCode, err.Error())
			return
		}

		c.JSON(http.StatusOK, times)
	}
}
