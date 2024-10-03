package mocks

import (
	"fmt"

	"github.com/judewood/bakery/models"
	"github.com/stretchr/testify/mock"
)

type MockRecipeStore struct {
	mock.Mock
}

func NewMockRecipeStore() *MockRecipeStore {
	return &MockRecipeStore{}
}

func (m *MockRecipeStore) GetRecipe(id string) (models.Recipe, error) {
	fmt.Printf("GetRecipe called with id: %v\n", id)
	args := m.Called(id)
	return args.Get(0).(models.Recipe), args.Error(1)
}
