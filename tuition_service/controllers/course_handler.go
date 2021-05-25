package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	courseConnector "github.com/leechongyan/Studtor_backend/database_service/connector/course_connector"
	errorHelper "github.com/leechongyan/Studtor_backend/helpers/error_helpers"
)

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

func GetSingleCourse() gin.HandlerFunc {
	return func(c *gin.Context) {
		courseConnector := courseConnector.Init()

		courseId := c.Param("course")

		// get single course
		i, e := strconv.Atoi(courseId)
		if e != nil {
			err := errorHelper.RaiseCannotParseRequest()
			c.JSON(err.StatusCode, err.Error())
			return
		}

		course, e := courseConnector.SetCourseId(i).GetSingle()
		if e != nil {
			err := errorHelper.RaiseDatabaseError()
			c.JSON(err.StatusCode, err.Error())
			return
		}
		c.JSON(http.StatusOK, course)
	}
}

func GetCoursesOfTutor() gin.HandlerFunc {
	return func(c *gin.Context) {
		courseConnector := courseConnector.Init()

		tutorId := c.Param("tutor_id")

		// get single course
		i, e := strconv.Atoi(tutorId)
		if e != nil {
			err := errorHelper.RaiseCannotParseRequest()
			c.JSON(err.StatusCode, err.Error())
			return
		}

		courses, e := courseConnector.SetTutorId(i).GetAll()
		if e != nil {
			err := errorHelper.RaiseDatabaseError()
			c.JSON(err.StatusCode, err.Error())
			return
		}
		c.JSON(http.StatusOK, courses)
	}
}
