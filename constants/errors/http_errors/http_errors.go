package http_errors

import "errors"

var ErrUnauthorizedAccess = errors.New("unauthorized access")
var ErrInvalidEmail = errors.New("invalid email")
var ErrExistentAccount = errors.New("account exists")
var ErrNonExistentAccount = errors.New("account not exists")
var ErrWrongValidation = errors.New("wrong validation code")
var ErrWrongCredential = errors.New("incorrect email or password")
var ErrNotVerified = errors.New("account not verified")
var ErrExpiredRefreshToken = errors.New("expired refresh token")
var ErrInvalidTokenFormat = errors.New("invalid token format")
var ErrNoAuthorizationHeader = errors.New("no authorization header")
var ErrInvalidAuthorizationMethod = errors.New("invalid authorization method")
var ErrInvalidToken = errors.New("invalid token")
var ErrExpiredToken = errors.New("expired token")
var ErrParamParsingFailure = errors.New("param parsing failure")
var ErrJsonParsingFailure = errors.New("json parsing failure")
var ErrJsonValidationError = errors.New("json validation error")
var ErrFileParsingFailure = errors.New("file parsing failure")
