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

func (m *MockRecipeStore) GetRecipe(id string) (Recipe, error) {
	fmt.Printf("GetRecipe called with id: %v\n", id)
	args := m.Called(id)
	return args.Get(0).(Recipe), args.Error(1)
}
