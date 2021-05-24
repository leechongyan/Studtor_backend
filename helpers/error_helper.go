package helpers

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/leechongyan/Studtor_backend/constants"
)

type RequestError struct {
	StatusCode int

	Err error
}

func (r *RequestError) Error() string {
	return fmt.Sprintf("status %d: %v", r.StatusCode, r.Err)
}

func RaiseInitFailure() *RequestError {
	return &RequestError{
		StatusCode: http.StatusInternalServerError,
		Err:        errors.New(constants.INIT_FAILURE),
	}
}

func RaiseUnauthorizedAccess() *RequestError {
	return &RequestError{
		StatusCode: http.StatusUnauthorized,
		Err:        errors.New(constants.UNAUTHORIZED_ACCESS),
	}
}

func RaiseNonExistentAccount() *RequestError {
	return &RequestError{
		StatusCode: http.StatusUnauthorized,
		Err:        errors.New(constants.NONEXISTENT_ACCOUNT),
	}
}

func RaiseExistentAccount() *RequestError {
	return &RequestError{
		StatusCode: http.StatusBadRequest,
		Err:        errors.New(constants.EXISTENT_ACCOUNT),
	}
}

func RaiseWrongValidation() *RequestError {
	return &RequestError{
		StatusCode: http.StatusUnauthorized,
		Err:        errors.New(constants.WRONG_VALIDATION),
	}
}

func RaiseWrongLoginCredentials() *RequestError {
	return &RequestError{
		StatusCode: http.StatusUnauthorized,
		Err:        errors.New(constants.WRONG_LOGIN_CREDENTIALS),
	}
}

func RaiseFailureEmailSend() *RequestError {
	return &RequestError{
		StatusCode: http.StatusInternalServerError,
		Err:        errors.New(constants.FAILURE_EMAILSEND),
	}
}

func RaiseInvalidTokenFormat() *RequestError {
	return &RequestError{
		StatusCode: http.StatusUnauthorized,
		Err:        errors.New(constants.INVALID_TOKEN_FORMAT),
	}
}

func RaiseNoAuthorizationHeader() *RequestError {
	return &RequestError{
		StatusCode: http.StatusUnauthorized,
		Err:        errors.New(constants.No_AUTHORIZATION_HEADER),
	}
}

func RaiseInvalidAuthorizationMethod() *RequestError {
	return &RequestError{
		StatusCode: http.StatusUnauthorized,
		Err:        errors.New(constants.INVALID_AUTHORIZATION_METHOD),
	}
}

func RaiseInvalidToken() *RequestError {
	return &RequestError{
		StatusCode: http.StatusUnauthorized,
		Err:        errors.New(constants.INVALID_TOKEN),
	}
}

func RaiseExpiredToken() *RequestError {
	return &RequestError{
		StatusCode: http.StatusUnauthorized,
		Err:        errors.New(constants.EXPIRED_TOKEN),
	}
}

func RaiseFailureGenerateClaim() *RequestError {
	return &RequestError{
		StatusCode: http.StatusInternalServerError,
		Err:        errors.New(constants.CLAIMS_GENERATE_FAILURE),
	}
}

func RaiseCannotParseClaims() *RequestError {
	return &RequestError{
		StatusCode: http.StatusInternalServerError,
		Err:        errors.New(constants.CLAIMS_PARSE_FAILURE),
	}
}

func RaiseCannotParseJson() *RequestError {
	return &RequestError{
		StatusCode: http.StatusBadRequest,
		Err:        errors.New(constants.FAILURE_PARSE_JSON),
	}
}

func RaiseValidationErrorJson() *RequestError {
	return &RequestError{
		StatusCode: http.StatusBadRequest,
		Err:        errors.New(constants.VALIDATION_JSON_ERROR),
	}
}

func RaiseInvalidEmail() *RequestError {
	return &RequestError{
		StatusCode: http.StatusBadRequest,
		Err:        errors.New(constants.EMAIL_NOT_VALID),
	}
}

func RaiseNotVerified() *RequestError {
	return &RequestError{
		StatusCode: http.StatusUnauthorized,
		Err:        errors.New(constants.NOT_VERIFIED),
	}
}

func RaiseUserNotInDatabase() *RequestError {
	return &RequestError{
		StatusCode: http.StatusInternalServerError,
		Err:        errors.New(constants.USER_NOT_IN_DATABASE),
	}
}

func RaiseCannotSaveUserInDatabase() *RequestError {
	return &RequestError{
		StatusCode: http.StatusInternalServerError,
		Err:        errors.New(constants.CANNOT_SAVE_USER_IN_DATABASE),
	}
}

func RaiseDatabaseError() *RequestError {
	return &RequestError{
		StatusCode: http.StatusInternalServerError,
		Err:        errors.New(constants.DATABASE_ERROR),
	}
}

func RaiseLoginExpired() *RequestError {
	return &RequestError{
		StatusCode: http.StatusUnauthorized,
		Err:        errors.New(constants.LOGIN_EXPIRED),
	}
}

func RaiseCannotParseRequest() *RequestError {
	return &RequestError{
		StatusCode: http.StatusBadRequest,
		Err:        errors.New(constants.CANNOT_PARSE_REQUEST),
	}
}

func RaiseCannotParseFile() *RequestError {
	return &RequestError{
		StatusCode: http.StatusBadRequest,
		Err:        errors.New(constants.CANNOT_PARSE_FILE),
	}
}

func RaiseStorageFailure() *RequestError {
	return &RequestError{
		StatusCode: http.StatusInternalServerError,
		Err:        errors.New(constants.STORAGE_ERROR),
	}
}

func RaiseNoRefreshToken() *RequestError {
	return &RequestError{
		StatusCode: http.StatusInternalServerError,
		Err:        errors.New(constants.NO_REFRESH_TOKEN),
	}
}
