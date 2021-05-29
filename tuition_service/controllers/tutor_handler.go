package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	authModel "github.com/leechongyan/Studtor_backend/authentication_service/models"
	httpError "github.com/leechongyan/Studtor_backend/constants/errors/http_errors"
	tutorConnector "github.com/leechongyan/Studtor_backend/database_service/connector/tutor_connector"
	httpHelper "github.com/leechongyan/Studtor_backend/helpers/http_helpers"
	typeHelper "github.com/leechongyan/Studtor_backend/helpers/type_conversion"
	"github.com/leechongyan/Studtor_backend/tuition_service/models"
)

// Get all the tutors who taught the course with pagination
func GetTutorsForCourse() gin.HandlerFunc {
	return func(c *gin.Context) {
		var query models.TutorPaginatedQuery

		err := httpHelper.ExtractGetRequestBody(c, &query)
		if err != nil {
			c.JSON(http.StatusBadRequest, err.Error())
			return
		}

		courseId, err := strconv.Atoi(c.Param("course_id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, httpError.ErrParamParsingFailure.Error())
			return
		}

		tutorConnector := tutorConnector.Init()

		tutorConnector.SetCourseId(courseId)

		if query.FromId != nil {
			tutorConnector.SetTutorId(*query.FromId)
		}

		if query.Size != nil {
			tutorConnector.SetSize(*query.Size)
		}

		tutors, err := tutorConnector.GetAll()

		if err != nil {
			c.JSON(http.StatusInternalServerError, err.Error())
			return
		}

		// convert tutors to hide confidential info
		tutorProfiles := make([]authModel.Userprofile, len(tutors))
		for i, tutor := range tutors {
			tutorProfiles[i] = typeHelper.ConvertFromDatabaseUserToUserProfile(tutor)
		}

		c.JSON(http.StatusOK, tutorProfiles)
	}
}

// Register to teach a course for a tutor
func RegisterCourse() gin.HandlerFunc {
	return func(c *gin.Context) {

		// check tutid param whether is same as the token id
		if c.GetString("id") != c.Param("tutor_id") {
			c.JSON(http.StatusUnauthorized, httpError.ErrUnauthorizedAccess.Error())
			return
		}

		tutorId, _ := strconv.Atoi(c.GetString("id"))

		courseId, err := strconv.Atoi(c.Param("course_id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, httpError.ErrParamParsingFailure.Error())
			return
		}

		tutorConnector := tutorConnector.Init()
		err = tutorConnector.SetTutorId(tutorId).SetCourseId(courseId).Add()
		if err != nil {
			c.JSON(http.StatusInternalServerError, err.Error())
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
			c.JSON(http.StatusUnauthorized, httpError.ErrUnauthorizedAccess.Error())
			return
		}

		tutorId, _ := strconv.Atoi(c.GetString("id"))

		courseId, err := strconv.Atoi(c.Param("course_id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, httpError.ErrParamParsingFailure.Error())
			return
		}

		tutorConnector := tutorConnector.Init()
		err = tutorConnector.SetTutorId(tutorId).SetCourseId(courseId).Delete()
		if err != nil {
			c.JSON(http.StatusInternalServerError, err.Error())
			return
		}
		c.JSON(http.StatusOK, "Success")
	}
}
