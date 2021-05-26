package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	tutorConnector "github.com/leechongyan/Studtor_backend/database_service/connector/tutor_connector"
	errorHelper "github.com/leechongyan/Studtor_backend/helpers/error_helpers"
	httpHelper "github.com/leechongyan/Studtor_backend/helpers/http_helpers"
	"github.com/leechongyan/Studtor_backend/tuition_service/models"
)

// Get all the tutors who taught the course with pagination
func GetTutorsForCourse() gin.HandlerFunc {
	return func(c *gin.Context) {
		var query models.TutorPaginatedQuery

		err := httpHelper.ExtractGetRequestBody(c, &query)
		if err != nil {
			c.JSON(err.StatusCode, err.Error())
			return
		}

		tutorConnector := tutorConnector.Init()

		course := c.Param("course")
		courseId, _ := strconv.Atoi(course)
		tutorConnector.SetCourseId(courseId)

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

// Register to teach a course for a tutor
func RegisterCourse() gin.HandlerFunc {
	return func(c *gin.Context) {

		// check tutid param whether is same as the token id
		if c.GetString("id") != c.Param("tutor_id") {
			err := errorHelper.RaiseUnauthorizedAccess()
			c.JSON(err.StatusCode, err.Error())
			return
		}

		tutorId, _ := strconv.Atoi(c.GetString("id"))
		courseId, e := strconv.Atoi(c.Param("course_id"))
		if e != nil {
			err := errorHelper.RaiseCannotParseRequest()
			c.JSON(err.StatusCode, err.Error())
			return
		}

		tutorConnector := tutorConnector.Init()
		e = tutorConnector.SetTutorId(tutorId).SetCourseId(courseId).Add()
		if e != nil {
			err := errorHelper.RaiseDatabaseError()
			c.JSON(err.StatusCode, err.Error())
			return
		}
		c.JSON(http.StatusOK, "Success")
	}
}

// Deregister to teach a course for a tutor
func DeregisterCourse() gin.HandlerFunc {
	return func(c *gin.Context) {

		// check tutid param whether is same as the token id
		if c.GetString("id") != c.Param("tutor_id") {
			err := errorHelper.RaiseUnauthorizedAccess()
			c.JSON(err.StatusCode, err.Error())
			return
		}

		tutorId, _ := strconv.Atoi(c.GetString("id"))
		courseId, e := strconv.Atoi(c.Param("course_id"))
		if e != nil {
			err := errorHelper.RaiseCannotParseRequest()
			c.JSON(err.StatusCode, err.Error())
			return
		}

		tutorConnector := tutorConnector.Init()
		e = tutorConnector.SetTutorId(tutorId).SetCourseId(courseId).Delete()
		if e != nil {
			err := errorHelper.RaiseDatabaseError()
			c.JSON(err.StatusCode, err.Error())
			return
		}
		c.JSON(http.StatusOK, "Success")
	}
}
