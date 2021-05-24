package http_helpers

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	errorHelper "github.com/leechongyan/Studtor_backend/helpers/error_helpers"
)

var validate = validator.New()

func ExtractGetRequestBody(c *gin.Context, req interface{}) (err *errorHelper.RequestError) {
	e := c.ShouldBind(req)
	if e != nil {
		err = errorHelper.RaiseCannotParseJson()
		return
	}
	validationErr := validate.Struct(req)
	if validationErr != nil {
		err = errorHelper.RaiseValidationErrorJson()
		return
	}
	return nil
}

func ExtractPostRequestBody(c *gin.Context, req interface{}) (err *errorHelper.RequestError) {
	e := c.BindJSON(req)
	if e != nil {
		err = errorHelper.RaiseCannotParseJson()
		return
	}
	validationErr := validate.Struct(req)
	if validationErr != nil {
		err = errorHelper.RaiseValidationErrorJson()
		return
	}
	return nil
}
