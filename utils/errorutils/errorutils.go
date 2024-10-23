package errorutils

import (
	"errors"
)

var ErrorNotFound = errors.New("not found")

// ErrorMissingID returns a consistently formatted
// error for a missing or badly formatted input identifier
var ErrorMissingID = errors.New("missing id in request")

// ErrorMissingRequired returns a consistently formatted
// error for a missing required field
var ErrorMissingRequired = errors.New("missing required field")
