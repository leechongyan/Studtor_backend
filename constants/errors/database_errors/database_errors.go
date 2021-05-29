package database_errors

import "errors"

// Define the database errors
var ErrNoRecordFound = errors.New("no record found")

// ErrDatabaseInternalError is the catch all for all database errors excluding records not found
var ErrDatabaseInternalError = errors.New("database internal error")

// Database connector errors
var ErrNotEnoughParameters = errors.New("not enough parameters")
var ErrInvalidTimes = errors.New("from time cannot be greater than to time")
var ErrMethodNotImplemented = errors.New("method is not implemented")
var ErrInvalidSizeParameter = errors.New("size cannot be 0 or negative")
var ErrInvalidParameters = errors.New("input parameters are not valid")
var ErrRecordToBeUpdatedNotFound = errors.New("record to be updated not found")
var ErrCreateRecordFailed = errors.New("failed to create record")
var ErrUpdateRecordFailed = errors.New("failed to update record")
var ErrDeleteRecordFailed = errors.New("failed to delete record")
var ErrRecordAlreadyExists = errors.New("record already exists")
