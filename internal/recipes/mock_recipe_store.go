package recipes

import (
	"fmt"

	"github.com/stretchr/testify/mock"
)

type MockRecipeStore struct {
	mock.Mock
}

func NewMockRecipeStore() *MockRecipeStore {
	return &MockRecipeStore{}
}

func (m *MockRecipeStore) GetRecipeFromS3(id string) (Recipe, error) {
	fmt.Printf("GetRecipeFromS3 called with id: %v\n", id)
	args := m.Called(id)
	return args.Get(0).(Recipe), args.Error(1)
}
