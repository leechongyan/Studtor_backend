package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	availability_connector "github.com/leechongyan/Studtor_backend/database_service/connector/availability_connector"
	booking_connector "github.com/leechongyan/Studtor_backend/database_service/connector/booking_connector"
	course_connector "github.com/leechongyan/Studtor_backend/database_service/connector/course_connector"
	tutor_connector "github.com/leechongyan/Studtor_backend/database_service/connector/tutor_connector"
	"github.com/leechongyan/Studtor_backend/helpers"
	"github.com/leechongyan/Studtor_backend/tuition_service/models"
)

func GetAllCourses() gin.HandlerFunc {
	return func(c *gin.Context) {
		var query models.CoursePaginated_query

		err := helpers.ExtractGetRequestBody(c, &query)
		if err != nil {
			c.JSON(err.StatusCode, err.Error())
			return
		}
		get_course_connector := course_connector.Init()

		if query.Size != nil {
			get_course_connector.SetSize(*query.Size)
		}
		if query.From_id != nil {
			get_course_connector.SetCourseCode(*query.From_id)
		}
		courses, e := get_course_connector.GetAll()

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
		var query models.TutorPaginated_query

		err := helpers.ExtractGetRequestBody(c, &query)
		if err != nil {
			c.JSON(err.StatusCode, err.Error())
			return
		}

		get_tutor_connector := tutor_connector.Init()

		if query.From_id != nil {
			get_tutor_connector.SetTutorId(*query.From_id)
		}

		if query.Size != nil {
			get_tutor_connector.SetSize(*query.Size)
		}

		// whether you are getting the tutor for a certain course
		course := c.Param("course")
		if course != "" {
			course_id, _ := strconv.Atoi(course)
			get_tutor_connector.SetCourse(course_id)
		}

		tutors, e := get_tutor_connector.GetAll()

		if e != nil {
			err := helpers.RaiseDatabaseError()
			c.JSON(err.StatusCode, err.Error())
			return
		}
		c.JSON(http.StatusOK, tutors)
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
		get_time_connector := availability_connector.Init()

		if tutor_id != "" {
			tut_id, _ := strconv.Atoi(tutor_id)
			get_time_connector.SetTutorId(tut_id)
		}
		e := get_time_connector.SetFromTime(slot_query.From).SetToTime(slot_query.To).Add()

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
		var slot_query models.Availability

		err := helpers.ExtractPostRequestBody(c, &slot_query)
		if err != nil {
			c.JSON(err.StatusCode, err.Error())
			return
		}
		get_time_connector := availability_connector.Init()

		e := get_time_connector.SetAvailabilityId(*slot_query.Availability_id).Delete()

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

		get_time_connector := availability_connector.Init()

		if tutor_id != "" {
			tut_id, _ := strconv.Atoi(tutor_id)
			get_time_connector.SetTutorId(tut_id)
		}

		times, e := get_time_connector.SetFromTime(slot_query.From).SetToTime(slot_query.To).Get()

		if e != nil {
			err := helpers.RaiseDatabaseError()
			c.JSON(err.StatusCode, err.Error())
			return
		}

		c.JSON(http.StatusOK, times)
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

		get_time_connector := booking_connector.Init()

		id, _ := strconv.Atoi(c.GetString("id"))

		e := get_time_connector.SetCourseId(*slot_query.Course).SetStudentId(id).SetAvailabilityId(*slot_query.Availability_id).Add()

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

		get_time_connector := booking_connector.Init()

		id, _ := strconv.Atoi(c.GetString("id"))

		e := get_time_connector.SetCourseId(*slot_query.Course).SetStudentId(id).SetAvailabilityId(*slot_query.Availability_id).Delete()

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

		user_id := c.Param("user")

		get_time_connector := booking_connector.Init()
		var id int
		if user_id != "" {
			id, _ = strconv.Atoi(user_id)
		} else {
			id, _ = strconv.Atoi(c.GetString("id"))
		}

		times, e := get_time_connector.SetUserId(id).SetFromTime(slot_query.From).SetToTime(slot_query.To).Get()

		if e != nil {
			err := helpers.RaiseDatabaseError()
			c.JSON(err.StatusCode, err.Error())
			return
		}

		c.JSON(http.StatusOK, times)
	}
}
