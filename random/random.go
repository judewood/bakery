package random

import "math/rand"

// RandomProvider has method set of random generator functions for testing
type RandomProvider interface {
	GetRandom(max int) int
}

// Random has methods that wrap math/rand package functions
// to allow them to be mocked for testing
type Random struct {
}

// GetRandom returns a random int between 0 and given max
func (r *Random) GetRandom(max int) int {
	return rand.Intn(max + 1)
}
