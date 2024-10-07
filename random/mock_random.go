package random

import (
	"github.com/stretchr/testify/mock"
)

type MockRandom struct {
	mock.Mock
}

func NewMockRandom() *MockRandom {
	return &MockRandom{}
}

// GetRandom returns the value specified by the test
func (m *MockRandom) GetRandom(max int) int {
	args := m.Called(max)
	return args.Get(0).(int)
}
