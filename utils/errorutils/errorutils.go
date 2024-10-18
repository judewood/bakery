package errorutils

import (
	"errors"
	"fmt"
)

const NotFound string = "id: %s not found"

// NotFoundError returns a consistently formatted
// error for when an entity cannot be found 
func NotFoundError(s string) error {
	return fmt.Errorf(NotFound, s)
}

// MissingRequired is consistent text for failed validation due to missing required field
const MissingRequired string = "missing required field"

// MissingID is consistent text for failed validation due to missing or badly formatted identifier in input
const MissingID = "missing id in request"

// ErrorMissingID returns a consistently formatted
// error for a missing or badly formatted input identifier 
var ErrorMissingID = errors.New(MissingID)

// ErrorMissingRequired returns a consistently formatted
// error for a missing required field 
var ErrorMissingRequired = errors.New(MissingRequired)
