package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	httpError "github.com/leechongyan/Studtor_backend/constants/errors/http_errors"
	courseConnector "github.com/leechongyan/Studtor_backend/database_service/connector/course_connector"
)

// Get all the courses
func GetCourses() gin.HandlerFunc {
	return func(c *gin.Context) {
		courses, err := courseConnector.Init().GetAll()
		if err != nil {
			c.JSON(http.StatusInternalServerError, err.Error())
			return
		}
		c.JSON(http.StatusOK, courses)
	}
}

// Get a single course
func GetSingleCourse() gin.HandlerFunc {
	return func(c *gin.Context) {
		// get single course
		courseId, err := strconv.Atoi(c.Param("course_id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, httpError.ErrParamParsingFailure.Error())
			return
		}

		course, err := courseConnector.Init().SetCourseId(courseId).GetSingle()
		if err != nil {
			c.JSON(http.StatusInternalServerError, err.Error())
			return
		}
		c.JSON(http.StatusOK, course)
	}
}

// Get all the courses taught by a tutor
func GetCoursesOfTutor() gin.HandlerFunc {
	return func(c *gin.Context) {
		tutorId, err := strconv.Atoi(c.Param("tutor_id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, httpError.ErrParamParsingFailure.Error())
			return
		}

		courses, err := courseConnector.Init().SetTutorId(tutorId).GetAll()
		if err != nil {
			c.JSON(http.StatusInternalServerError, err.Error())
			return
		}
		c.JSON(http.StatusOK, courses)
	}
}
