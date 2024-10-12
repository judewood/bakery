package myfmt

import "testing"

const ThumbsDown = "\U0001f44e"
const ThumbsUp = "\U0001f44d"

func Errorf(t *testing.T, s string , args... any) {
	t.Errorf(ThumbsDown + s, args)
}

func Fatalf(t *testing.T, s string , args... any) {
	t.Fatalf(ThumbsDown + s, args)
}