package errors

import "errors"

var ErrStorage = errors.New("storage error")

var ErrStorageInitializationFailure = errors.New("cannot add storage client")
var ErrFileTransfer = errors.New("file transfer error")
var ErrObjectWriter = errors.New("cannot close object writer")
var ErrUrlFailure = errors.New("cannot generate url")
