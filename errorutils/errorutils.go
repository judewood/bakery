package errorutils

import (
	"errors"
	"fmt"
)

const NotFound string = "id: %s not found"

func NotFoundError(s string) error {
	return fmt.Errorf(NotFound, s)
}

const MissingRequired string = "missing required field"
const MissingID = "missing id in request"

var ErrorMissingID = errors.New(MissingID)
var ErrorMissingRequired = errors.New(MissingRequired)
