package http_helpers

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	httpError "github.com/leechongyan/Studtor_backend/constants/errors/http_errors"
)

var validate = validator.New()

func ExtractGetRequestBody(c *gin.Context, req interface{}) (err error) {
	e := c.ShouldBind(req)
	if e != nil {
		return httpError.ErrJsonParsingFailure
	}
	validationErr := validate.Struct(req)
	if validationErr != nil {
		return httpError.ErrJsonValidationError
	}
	return
}

func ExtractPostRequestBody(c *gin.Context, req interface{}) (err error) {
	e := c.BindJSON(req)
	if e != nil {
		return httpError.ErrJsonParsingFailure
	}
	validationErr := validate.Struct(req)
	if validationErr != nil {
		return httpError.ErrJsonValidationError
	}
	return
}
