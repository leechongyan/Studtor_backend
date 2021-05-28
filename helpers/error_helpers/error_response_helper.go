package error_helpers

// import (
// 	"errors"
// 	"fmt"
// 	"net/http"

// 	databaseError "github.com/leechongyan/Studtor_backend/constants/error_messages/database_errors"
// 	httpError "github.com/leechongyan/Studtor_backend/constants/error_messages/http_errors"
// 	storageError "github.com/leechongyan/Studtor_backend/constants/error_messages/storage_errors"
// 	systemError "github.com/leechongyan/Studtor_backend/constants/error_messages/system_errors"
// )

// // Functions to phrase an error into a http response error

// type RequestError struct {
// 	StatusCode int

// 	Err error
// }

// func (r *RequestError) Error() string {
// 	return fmt.Sprintf("status %d: %v", r.StatusCode, r.Err)
// }

// // start of system errors

// func RaiseFailureEmailSend() *RequestError {
// 	return &RequestError{
// 		StatusCode: http.StatusInternalServerError,
// 		Err:        errors.New(systemError.FAILURE_EMAILSEND),
// 	}
// }

// func RaiseFailureGenerateClaim() *RequestError {
// 	return &RequestError{
// 		StatusCode: http.StatusInternalServerError,
// 		Err:        errors.New(systemError.CLAIMS_GENERATE_FAILURE),
// 	}
// }

// func RaiseCannotParseClaims() *RequestError {
// 	return &RequestError{
// 		StatusCode: http.StatusInternalServerError,
// 		Err:        errors.New(systemError.CLAIMS_PARSE_FAILURE),
// 	}
// }

// // end of system errors

// // start of http errors

// func RaiseUnauthorizedAccess() *RequestError {
// 	return &RequestError{
// 		StatusCode: http.StatusUnauthorized,
// 		Err:        errors.New(httpError.UNAUTHORIZED_ACCESS),
// 	}
// }

// func RaiseInvalidEmail() *RequestError {
// 	return &RequestError{
// 		StatusCode: http.StatusBadRequest,
// 		Err:        errors.New(httpError.EMAIL_NOT_VALID),
// 	}
// }

// func RaiseWrongValidation() *RequestError {
// 	return &RequestError{
// 		StatusCode: http.StatusUnauthorized,
// 		Err:        errors.New(httpError.WRONG_VALIDATION),
// 	}
// }

// func RaiseWrongLoginCredentials() *RequestError {
// 	return &RequestError{
// 		StatusCode: http.StatusUnauthorized,
// 		Err:        errors.New(httpError.WRONG_LOGIN_CREDENTIALS),
// 	}
// }

// func RaiseInvalidTokenFormat() *RequestError {
// 	return &RequestError{
// 		StatusCode: http.StatusUnauthorized,
// 		Err:        errors.New(httpError.INVALID_TOKEN_FORMAT),
// 	}
// }

// func RaiseNoAuthorizationHeader() *RequestError {
// 	return &RequestError{
// 		StatusCode: http.StatusUnauthorized,
// 		Err:        errors.New(httpError.No_AUTHORIZATION_HEADER),
// 	}
// }

// func RaiseInvalidAuthorizationMethod() *RequestError {
// 	return &RequestError{
// 		StatusCode: http.StatusUnauthorized,
// 		Err:        errors.New(httpError.INVALID_AUTHORIZATION_METHOD),
// 	}
// }

// func RaiseInvalidToken() *RequestError {
// 	return &RequestError{
// 		StatusCode: http.StatusUnauthorized,
// 		Err:        errors.New(httpError.INVALID_TOKEN),
// 	}
// }

// func RaiseExpiredToken() *RequestError {
// 	return &RequestError{
// 		StatusCode: http.StatusUnauthorized,
// 		Err:        errors.New(httpError.EXPIRED_TOKEN),
// 	}
// }

// func RaiseCannotParseJson() *RequestError {
// 	return &RequestError{
// 		StatusCode: http.StatusBadRequest,
// 		Err:        errors.New(httpError.FAILURE_PARSE_JSON),
// 	}
// }

// func RaiseValidationErrorJson() *RequestError {
// 	return &RequestError{
// 		StatusCode: http.StatusBadRequest,
// 		Err:        errors.New(httpError.VALIDATION_JSON_ERROR),
// 	}
// }

// func RaiseCannotParseFile() *RequestError {
// 	return &RequestError{
// 		StatusCode: http.StatusBadRequest,
// 		Err:        errors.New(httpError.CANNOT_PARSE_FILE),
// 	}
// }

// // end of http errors

// // start of database errors

// func RaiseNonExistentAccount() *RequestError {
// 	return &RequestError{
// 		StatusCode: http.StatusUnauthorized,
// 		Err:        errors.New(databaseError.NONEXISTENT_ACCOUNT),
// 	}
// }

// func RaiseExistentAccount() *RequestError {
// 	return &RequestError{
// 		StatusCode: http.StatusBadRequest,
// 		Err:        errors.New(databaseError.EXISTENT_ACCOUNT),
// 	}
// }

// func RaiseNotVerified() *RequestError {
// 	return &RequestError{
// 		StatusCode: http.StatusUnauthorized,
// 		Err:        errors.New(databaseError.NOT_VERIFIED),
// 	}
// }

// func RaiseUserNotInDatabase() *RequestError {
// 	return &RequestError{
// 		StatusCode: http.StatusInternalServerError,
// 		Err:        errors.New(databaseError.USER_NOT_IN_DATABASE),
// 	}
// }

// func RaiseCannotSaveUserInDatabase() *RequestError {
// 	return &RequestError{
// 		StatusCode: http.StatusInternalServerError,
// 		Err:        errors.New(databaseError.CANNOT_SAVE_USER_IN_DATABASE),
// 	}
// }

// func RaiseDatabaseError() *RequestError {
// 	return &RequestError{
// 		StatusCode: http.StatusInternalServerError,
// 		Err:        errors.New(databaseError.DATABASE_ERROR),
// 	}
// }

// func RaiseNoRefreshToken() *RequestError {
// 	return &RequestError{
// 		StatusCode: http.StatusInternalServerError,
// 		Err:        errors.New(databaseError.NO_REFRESH_TOKEN),
// 	}
// }

// // end of database errors

// // start of storage errors

// func RaiseStorageFailure() *RequestError {
// 	return &RequestError{
// 		StatusCode: http.StatusInternalServerError,
// 		Err:        errors.New(storageError.STORAGE_ERROR),
// 	}
// }

// end of storage errors
