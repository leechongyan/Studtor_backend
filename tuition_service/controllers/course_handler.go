package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	courseConnector "github.com/leechongyan/Studtor_backend/database_service/connector/course_connector"
	errorHelper "github.com/leechongyan/Studtor_backend/helpers/error_helpers"
)

// Get all the courses
func GetCourses() gin.HandlerFunc {
	return func(c *gin.Context) {
		courseConnector := courseConnector.Init()
		courses, e := courseConnector.GetAll()

		if e != nil {
			err := errorHelper.RaiseDatabaseError()
			c.JSON(err.StatusCode, err.Error())
			return
		}
		c.JSON(http.StatusOK, courses)
	}
}

// Get a single course
func GetSingleCourse() gin.HandlerFunc {
	return func(c *gin.Context) {
		courseConnector := courseConnector.Init()

		// get single course
		courseId, e := strconv.Atoi(c.Param("course_id"))
		if e != nil {
			err := errorHelper.RaiseCannotParseRequest()
			c.JSON(err.StatusCode, err.Error())
			return
		}

		course, e := courseConnector.SetCourseId(courseId).GetSingle()
		if e != nil {
			err := errorHelper.RaiseDatabaseError()
			c.JSON(err.StatusCode, err.Error())
			return
		}
		c.JSON(http.StatusOK, course)
	}
}

// Get all the courses taught by a tutor
func GetCoursesOfTutor() gin.HandlerFunc {
	return func(c *gin.Context) {
		courseConnector := courseConnector.Init()

		tutorId, e := strconv.Atoi(c.Param("tutor_id"))
		if e != nil {
			err := errorHelper.RaiseCannotParseRequest()
			c.JSON(err.StatusCode, err.Error())
			return
		}

		courses, e := courseConnector.SetTutorId(tutorId).GetAll()
		if e != nil {
			err := errorHelper.RaiseDatabaseError()
			c.JSON(err.StatusCode, err.Error())
			return
		}
		c.JSON(http.StatusOK, courses)
	}
}
