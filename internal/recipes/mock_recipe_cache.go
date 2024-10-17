package recipes

import (
	"fmt"

	"github.com/stretchr/testify/mock"
)

type MockRecipeCache struct {
	mock.Mock
}

func NewMockRecipeCache() *MockRecipeCache {
	return &MockRecipeCache{}
}

func (m *MockRecipeCache) GetRecipe(id string) (Recipe, error) {
	fmt.Printf("GetRecipe called with id: %v\n", id)
	args := m.Called(id)
	return args.Get(0).(Recipe), args.Error(1)
}
