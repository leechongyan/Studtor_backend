package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	httpError "github.com/leechongyan/Studtor_backend/constants/errors/http_errors"
	availabilityConnector "github.com/leechongyan/Studtor_backend/database_service/connector/availability_connector"
	bookingConnector "github.com/leechongyan/Studtor_backend/database_service/connector/booking_connector"
	httpHelper "github.com/leechongyan/Studtor_backend/helpers/http_helpers"
	"github.com/leechongyan/Studtor_backend/tuition_service/models"
)

// Put tutor available time
func PutAvailableTimeTutor() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.GetString("id") != c.Param("tutor_id") {
			c.JSON(http.StatusUnauthorized, httpError.ErrUnauthorizedAccess.Error())
			return
		}

		var slotQuery models.TimeSlot

		err := httpHelper.ExtractPostRequestBody(c, &slotQuery)
		if err != nil {
			c.JSON(http.StatusBadRequest, err.Error())
			return
		}

		tutorId, _ := strconv.Atoi(c.GetString("id"))

		err = availabilityConnector.Init().SetTutorId(tutorId).SetFromTime(slotQuery.From).SetToTime(slotQuery.To).Add()
		if err != nil {
			c.JSON(http.StatusInternalServerError, err.Error())
			return
		}

		c.JSON(http.StatusOK, "Success")
	}
}

// Remove tutor available time
func DeleteAvailableTimeTutor() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.GetString("id") != c.Param("tutor_id") {
			c.JSON(http.StatusUnauthorized, httpError.ErrUnauthorizedAccess.Error())
			return
		}

		availabilityId, err := strconv.Atoi(c.Param("availability_id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, httpError.ErrParamParsingFailure.Error())
			return
		}

		tutorId, _ := strconv.Atoi(c.GetString("id"))

		// tutor id is needed to check whether the availabilityid belongs to the tutor id
		err = availabilityConnector.Init().SetTutorId(tutorId).SetAvailabilityId(availabilityId).Delete()
		if err != nil {
			if err == httpError.ErrUnauthorizedAccess {
				c.JSON(http.StatusUnauthorized, err.Error())
				return
			}
			c.JSON(http.StatusInternalServerError, err.Error())
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
			c.JSON(http.StatusBadRequest, err.Error())
			return
		}

		tutId, err := strconv.Atoi(c.Param("tutor_id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, httpError.ErrParamParsingFailure.Error())
			return
		}

		times, err := availabilityConnector.Init().SetTutorId(tutId).SetFromTime(slotQuery.From).SetToTime(slotQuery.To).GetAll()
		if err != nil {
			c.JSON(http.StatusInternalServerError, err.Error())
			return
		}

		c.JSON(http.StatusOK, times)
	}
}

// Book an available timeslot with the tutor with the course id
func BookTimeTutor() gin.HandlerFunc {
	return func(c *gin.Context) {
		availabilityId, err := strconv.Atoi(c.Param("availability_id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, httpError.ErrParamParsingFailure.Error())
			return
		}

		courseId, err := strconv.Atoi(c.Param("course_id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, httpError.ErrParamParsingFailure.Error())
			return
		}

		id, _ := strconv.Atoi(c.GetString("id"))

		err = bookingConnector.Init().SetCourseId(courseId).SetUserId(id).SetAvailabilityId(availabilityId).Add()
		if err != nil {
			c.JSON(http.StatusInternalServerError, err.Error())
			return
		}

		c.JSON(http.StatusOK, "Success")
	}
}

// Remove a booking
func UnbookTimeTutor() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.GetString("id") != c.Param("user_id") {
			c.JSON(http.StatusUnauthorized, httpError.ErrUnauthorizedAccess.Error())
			return
		}

		bookingId, err := strconv.Atoi(c.Param("booking_id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, httpError.ErrParamParsingFailure.Error())
			return
		}

		userId, _ := strconv.Atoi(c.GetString("id"))

		// user id is needed to check whether the bookingid involves the user
		err = bookingConnector.Init().SetUserId(userId).SetBookingId(bookingId).Delete()
		if err != nil {
			if err == httpError.ErrUnauthorizedAccess {
				c.JSON(http.StatusUnauthorized, err.Error())
				return
			}
			c.JSON(http.StatusInternalServerError, err.Error())
			return
		}

		c.JSON(http.StatusOK, "Success")
	}
}

// Get all the booked time of a user with pagination
func GetAllBookedTime() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.GetString("id") != c.Param("user_id") {
			c.JSON(http.StatusUnauthorized, httpError.ErrUnauthorizedAccess.Error())
			return
		}

		var slotQuery models.TimePaginatedQuery
		err := httpHelper.ExtractGetRequestBody(c, &slotQuery)
		if err != nil {
			c.JSON(http.StatusBadRequest, err.Error())
			return
		}

		userId, _ := strconv.Atoi(c.GetString("id"))

		times, err := bookingConnector.Init().SetUserId(userId).SetFromTime(slotQuery.From).SetToTime(slotQuery.To).GetAll()
		if err != nil {
			c.JSON(http.StatusInternalServerError, err.Error())
			return
		}

		c.JSON(http.StatusOK, times)
	}
}
