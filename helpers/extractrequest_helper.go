package helpers

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

func ExtractGetRequestBody(c *gin.Context, req interface{}) (err *RequestError) {
	e := c.ShouldBind(req)
	if e != nil {
		err = RaiseCannotParseJson()
		return
	}
	validationErr := validate.Struct(req)
	fmt.Print(validationErr)
	if validationErr != nil {
		err = RaiseValidationErrorJson()
		return
	}
	return nil
}

func ExtractPostRequestBody(c *gin.Context, req interface{}) (err *RequestError) {
	e := c.BindJSON(req)
	if e != nil {
		err = RaiseCannotParseJson()
		return
	}
	validationErr := validate.Struct(req)
	fmt.Print(validationErr)
	if validationErr != nil {
		err = RaiseValidationErrorJson()
		return
	}
	return nil
}
