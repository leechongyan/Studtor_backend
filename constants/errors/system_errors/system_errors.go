package system_errors

import "errors"

var ErrInitFailure = errors.New("fail to parse config file")
var ErrEmailSendingFailure = errors.New("fail to send email")
var ErrClaimsGenerateFailure = errors.New("fail to generate claims")
var ErrClaimsParseFailure = errors.New("fail to parse claims")
