package errors

import "errors"

// Define the database errors

// ErrDatabaseInternalError is the catch all for all database errors excluding records not found
// capture other gorm errors for internal error
var ErrInvalidAvailability = errors.New("availability is occupied")
var ErrNoRecordFound = errors.New("no records found")
var ErrUnauthorizedAccess = errors.New("unauthorized access")
var ErrInvalidBooking = errors.New("Tutor does not teach this course")

// Database connector errors
// validation error
var ErrNotEnoughParameters = errors.New("not enough parameters")
var ErrMethodNotImplemented = errors.New("method is not implemented")
var ErrInvalidSizeParameter = errors.New("size cannot be 0 or negative")
var ErrInvalidParameters = errors.New("input parameters are not valid")
