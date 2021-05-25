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

func GetAllTutors() gin.HandlerFunc {
	return func(c *gin.Context) {
		var query models.TutorPaginatedQuery

		err := httpHelper.ExtractGetRequestBody(c, &query)
		if err != nil {
			c.JSON(err.StatusCode, err.Error())
			return
		}

		tutorConnector := tutorConnector.Init()

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
	}
}

func GetSingleTutor() gin.HandlerFunc {
	return func(c *gin.Context) {
		var query models.TutorPaginatedQuery

		err := httpHelper.ExtractGetRequestBody(c, &query)
		if err != nil {
			c.JSON(err.StatusCode, err.Error())
			return
		}

		tutorConnector := tutorConnector.Init()

		tutId := c.Param("tutor_id")

		// return a single tutor
		tutorId, e := strconv.Atoi(tutId)
		if e != nil {
			err := errorHelper.RaiseCannotParseRequest()
			c.JSON(err.StatusCode, err.Error())
			return
		}
		tutor, e := tutorConnector.SetTutorId(tutorId).GetSingle()
		if e != nil {
			err := errorHelper.RaiseDatabaseError()
			c.JSON(err.StatusCode, err.Error())
			return
		}
		c.JSON(http.StatusOK, tutor)
	}
}

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

func RegisterCourse() gin.HandlerFunc {
	return func(c *gin.Context) {

		tutId := c.Param("tutor_id")
		tutorId, e := strconv.Atoi(tutId)
		if e != nil {
			err := errorHelper.RaiseCannotParseRequest()
			c.JSON(err.StatusCode, err.Error())
			return
		}

		cId := c.Param("course_id")
		courseId, e := strconv.Atoi(cId)
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

func DeregisterCourse() gin.HandlerFunc {
	return func(c *gin.Context) {

		tutId := c.Param("tutor_id")
		tutorId, e := strconv.Atoi(tutId)
		if e != nil {
			err := errorHelper.RaiseCannotParseRequest()
			c.JSON(err.StatusCode, err.Error())
			return
		}

		cId := c.Param("course_id")
		courseId, e := strconv.Atoi(cId)
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
