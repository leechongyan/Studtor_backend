package database_errors

import "errors"

// Define the database errors, errors are thrown in 2 scenarios, first is when there is no entry, second is when database update fails
var ErrNoEntry = errors.New("no entry")
var ErrDatabaseInternalError = errors.New("database internal error")

// Database connector errors
var ErrNotEnoughParameters = errors.New("not enough parameters")
var ErrInvalidTimes = errors.New("from time cannot be greater than to time")
var ErrMethodNotImplemented = errors.New("method is not implemented")
var ErrInvalidSizeParameter = errors.New("size cannot be 0 or negative")
