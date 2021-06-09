package controllers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	httpError "github.com/leechongyan/Studtor_backend/constants/errors/http_errors"
	clientModel "github.com/leechongyan/Studtor_backend/database_service/client_models"
	userModel "github.com/leechongyan/Studtor_backend/database_service/client_models"
	availabilityConnector "github.com/leechongyan/Studtor_backend/database_service/connector/availability_connector"
	bookingConnector "github.com/leechongyan/Studtor_backend/database_service/connector/booking_connector"
	courseConnector "github.com/leechongyan/Studtor_backend/database_service/connector/course_connector"
	userConnector "github.com/leechongyan/Studtor_backend/database_service/connector/user_connector"
	databaseModel "github.com/leechongyan/Studtor_backend/database_service/database_models"
	httpHelper "github.com/leechongyan/Studtor_backend/helpers/http_helpers"
	timeslotHelper "github.com/leechongyan/Studtor_backend/helpers/timeslot_helpers"
	mailService "github.com/leechongyan/Studtor_backend/mail_service"
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

		_, err = availabilityConnector.Init().SetTutorId(tutorId).SetDate(slotQuery.Date).SetTimeId(*slotQuery.TimeId).Add()
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

		times, err := availabilityConnector.Init().SetTutorId(tutId).SetDays(*slotQuery.Days).SetDate(slotQuery.Date).GetAll()
		if err != nil {
			c.JSON(http.StatusInternalServerError, err.Error())
			return
		}

		c.JSON(http.StatusOK, times)
	}
}

func getAvailabilityFromID(availabilityId int) (availability databaseModel.Availability, err error) {
	return availabilityConnector.Init().SetAvailabilityId(availabilityId).GetSingle()
}

func getUserFromID(userId int) (user userModel.User, err error) {
	return userConnector.Init().SetUserId(userId).GetUser()
}

func getCourseFromID(courseId int) (course userModel.CourseWithSize, err error) {
	return courseConnector.Init().SetCourseId(courseId).GetSingle()
}

func notifyBooking(booking clientModel.BookingDetails) (err error) {
	return mailService.CurrentMailService.SendBookingConfirmation(booking.StudentName(), booking.StudentEmail(), booking.TutorName(), booking.TutorEmail(), booking.CourseName(), booking.Date(), timeslotHelper.ConvertSlotIdToTimeString(booking.TimeSlot()))
}

func notifyUnbooking(booking clientModel.BookingDetails) (err error) {
	return mailService.CurrentMailService.SendBookingCancellation(booking.StudentName(), booking.StudentEmail(), booking.TutorName(), booking.TutorEmail(), booking.CourseName(), booking.Date(), timeslotHelper.ConvertSlotIdToTimeString(booking.TimeSlot()))
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

		uid, _ := strconv.Atoi(c.GetString("id"))

		id, err := bookingConnector.Init().SetCourseId(courseId).SetUserId(uid).SetAvailabilityId(availabilityId).Add()
		if err != nil {
			c.JSON(http.StatusInternalServerError, err.Error())
			return
		}

		// get the full booking
		book, err := bookingConnector.Init().SetBookingId(id).GetSingle()
		if err != nil {
			c.JSON(http.StatusInternalServerError, err.Error())
			return
		}

		// notify booking
		err = notifyBooking(book)

		log.Printf("%v", err)

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

		// get the booking details first
		booking, err := bookingConnector.Init().SetBookingId(bookingId).GetSingle()
		if err != nil {
			c.JSON(http.StatusInternalServerError, err.Error())
			return
		}

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

		// if successful cancellation then notify
		err = notifyUnbooking(booking)

		log.Printf("%v", err)

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

		times, err := bookingConnector.Init().SetUserId(userId).SetDate(slotQuery.Date).SetDays(*slotQuery.Days).GetAll()
		if err != nil {
			c.JSON(http.StatusInternalServerError, err.Error())
			return
		}

		c.JSON(http.StatusOK, times)
	}
}
