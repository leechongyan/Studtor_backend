package database_errors

import "errors"

// Define the database errors, errors are thrown in 2 scenarios, first is when there is no entry, second is when database update fails
var ErrNoEntry = errors.New("no entry")
var ErrDatabaseInternalError = errors.New("database internal error")
