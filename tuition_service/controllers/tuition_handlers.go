package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	availabilityConnector "github.com/leechongyan/Studtor_backend/database_service/connector/availability_connector"
	bookingConnector "github.com/leechongyan/Studtor_backend/database_service/connector/booking_connector"
	courseConnector "github.com/leechongyan/Studtor_backend/database_service/connector/course_connector"
	tutorConnector "github.com/leechongyan/Studtor_backend/database_service/connector/tutor_connector"
	errorHelper "github.com/leechongyan/Studtor_backend/helpers/error_helpers"
	httpHelper "github.com/leechongyan/Studtor_backend/helpers/http_helpers"
	"github.com/leechongyan/Studtor_backend/tuition_service/models"
)

func GetCourses() gin.HandlerFunc {
	return func(c *gin.Context) {
		var query models.CoursePaginatedQuery

		err := httpHelper.ExtractGetRequestBody(c, &query)
		if err != nil {
			c.JSON(err.StatusCode, err.Error())
			return
		}

		courseConnector := courseConnector.Init()

		if query.Size != nil {
			courseConnector.SetSize(*query.Size)
		}
		if query.FromId != nil {
			courseConnector.SetCourseCode(*query.FromId)
		}
		courses, e := courseConnector.GetAll()

		if e != nil {
			err := errorHelper.RaiseDatabaseError()
			c.JSON(err.StatusCode, err.Error())
			return
		}
		c.JSON(http.StatusOK, courses)
	}
}

func GetSingleCourse() gin.HandlerFunc {
	return func(c *gin.Context) {
		var query models.CoursePaginatedQuery

		err := httpHelper.ExtractGetRequestBody(c, &query)
		if err != nil {
			c.JSON(err.StatusCode, err.Error())
			return
		}

		courseConnector := courseConnector.Init()

		courseId := c.Param("course")

		// get single course
		i, _ := strconv.Atoi(courseId)
		course, e := courseConnector.SetCourseId(i).GetSingle()
		if e != nil {
			err := errorHelper.RaiseDatabaseError()
			c.JSON(err.StatusCode, err.Error())
			return
		}
		c.JSON(http.StatusOK, course)
		return
	}
}

func GetTutors() gin.HandlerFunc {
	return func(c *gin.Context) {
		var query models.TutorPaginatedQuery

		err := httpHelper.ExtractGetRequestBody(c, &query)
		if err != nil {
			c.JSON(err.StatusCode, err.Error())
			return
		}

		tutorConnector := tutorConnector.Init()

		tutId := c.Param("tutor_id")

		if tutId != "" {
			// return a single tutor
			tutorId, _ := strconv.Atoi(tutId)
			tutor, e := tutorConnector.SetTutorId(tutorId).GetSingle()
			if e != nil {
				err := errorHelper.RaiseDatabaseError()
				c.JSON(err.StatusCode, err.Error())
				return
			}
			c.JSON(http.StatusOK, tutor)
			return
		}
		// return a list of tutors
		if query.FromId != nil {
			tutorConnector.SetTutorId(*query.FromId)
		}

		if query.Size != nil {
			tutorConnector.SetSize(*query.Size)
		}

		tutors, e := tutorConnector.GetAll()
		if e != nil {
			err := errorHelper.RaiseDatabaseError()
			c.JSON(err.StatusCode, err.Error())
			return
		}
		c.JSON(http.StatusOK, tutors)
		return
	}
}

func GetCoursesTutors() gin.HandlerFunc {
	return func(c *gin.Context) {
		var query models.TutorPaginatedQuery

		err := httpHelper.ExtractGetRequestBody(c, &query)
		if err != nil {
			c.JSON(err.StatusCode, err.Error())
			return
		}

		tutorConnector := tutorConnector.Init()

		// whether you are getting the tutor for a certain course
		// getting a list of tutors
		course := c.Param("course")
		courseId, _ := strconv.Atoi(course)
		tutorConnector.SetCourse(courseId)

		if query.FromId != nil {
			tutorConnector.SetTutorId(*query.FromId)
		}

		if query.Size != nil {
			tutorConnector.SetSize(*query.Size)
		}

		tutors, e := tutorConnector.GetAll()

		if e != nil {
			err := errorHelper.RaiseDatabaseError()
			c.JSON(err.StatusCode, err.Error())
			return
		}
		c.JSON(http.StatusOK, tutors)
	}
}

func PutAvailableTimeTutor() gin.HandlerFunc {
	return func(c *gin.Context) {
		var slotQuery models.TimeSlot

		err := httpHelper.ExtractPostRequestBody(c, &slotQuery)
		if err != nil {
			c.JSON(err.StatusCode, err.Error())
			return
		}
		tutorId := c.GetString("id")
		availabilityConnector := availabilityConnector.Init()

		if tutorId != "" {
			tutId, _ := strconv.Atoi(tutorId)
			availabilityConnector.SetTutorId(tutId)
		}
		e := availabilityConnector.SetFromTime(slotQuery.From).SetToTime(slotQuery.To).Add()

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

		tutorId := c.Param("tutor")

		availabilityConnector := availabilityConnector.Init()

		if tutorId != "" {
			tutId, _ := strconv.Atoi(tutorId)
			availabilityConnector.SetTutorId(tutId)
		}

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

		userId := c.Param("user")

		bookingConnector := bookingConnector.Init()

		var id int
		if userId != "" {
			id, _ = strconv.Atoi(userId)
		} else {
			id, _ = strconv.Atoi(c.GetString("id"))
		}

		times, e := bookingConnector.SetUserId(id).SetFromTime(slotQuery.From).SetToTime(slotQuery.To).Get()

		if e != nil {
			err := errorHelper.RaiseDatabaseError()
			c.JSON(err.StatusCode, err.Error())
			return
		}

		c.JSON(http.StatusOK, times)
	}
}
