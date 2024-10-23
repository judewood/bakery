package testutils

import (
	"testing"
)

const ThumbsDown = "\U0001f44e"
const ThumbsUp = "\U0001f44d"

// Passed prints a yellow thumbs up emoji to the test output for improved readability
func Passed(t *testing.T) {
	t.Log(ThumbsUp)
}

// Failed prints a yellow thumbs down emoji followed by the given actual result 
// to the test output for improved readability
func Failed(t *testing.T, got any) {
	t.Errorf("\n%s Got: %v", ThumbsDown, got)
}

// FailedToDeserialise prints a yellow thumbs down emoji followed by 
// the fail reason and the given input that could not be deserialised
// and prints it to the test output for improved readability
func FailedToDeserialise(t *testing.T, body any) {
	t.Errorf("\n%s testutils.Failed to deserialise the response: %v", ThumbsDown, body)
}

// FailedToDeserialise returns prefixed with a yellow thumbs down emoji followed by 
// the fail reason and the given input stream that could not be read
// and prints it to the test output for improved readability
func FailedToReadResponse(t *testing.T, err error) {
	t.Errorf("\n%s Failed to read the response body: %v", ThumbsDown, err)
}
