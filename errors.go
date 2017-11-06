package outputter

import (
	"errors"
)

// ErrorOutputAddRowNoHeaders is an error for not having yet called SetHeaders()
var ErrorOutputAddRowNoHeaders = errors.New("Cannot AddRow with before calling SetHeaders")

// ErrorOutputAddRowTooFewHeaders is an error for having fewer keys than values
var ErrorOutputAddRowTooFewHeaders = errors.New("Cannot AddRow with more values than headers")

// ErrorUnknownOutputter is an error for specifying an unknown ErrorUnknownFormatter
var ErrorUnknownOutputter = errors.New("Unknown formatter specified")

// ErrorInvalidOutputter is an erro for specifying that a registered output isn't an Ouputter
var ErrorInvalidOutputter = errors.New("Specified outputter is invalid")
