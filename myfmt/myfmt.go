package myfmt

import "testing"

const ThumbsDown = "\U0001f44e"

func Errorf(t *testing.T, s string , args... any) {
	t.Errorf(ThumbsDown + s, args)
}