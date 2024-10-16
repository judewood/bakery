package testutils

import (
	"testing"
)

const ThumbsDown = "\U0001f44e"
const ThumbsUp = "\U0001f44d"

func Passed(t *testing.T) {
	t.Log(ThumbsUp)
}

func Failed(t *testing.T, got any) {
	t.Errorf("\n%s Got: %v", ThumbsDown, got)
}

func FailedToDeserialise(t *testing.T, body any) {
	t.Errorf("\n%s testutils.Failed to deserialise the response: %v", ThumbsDown, body)
}

func FailedToReadResponse(t *testing.T, err error) {
	t.Errorf("\n%s Failed to read the response body: %v", ThumbsDown, err)
}
