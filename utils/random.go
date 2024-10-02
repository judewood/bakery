package utils

import "math/rand"

type IRandom interface {
	GetRandom(max int) int
}

type Random struct {
	RandomValue int
}

func NewRandom() *Random {
	return &Random{}
}

// GetRandom returns a random int between 0 and given max
func (r *Random) GetRandom(max int) int {
	return rand.Intn(max + 1)
}
