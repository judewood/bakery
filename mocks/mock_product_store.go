package mocks

import (
	"fmt"

	"github.com/judewood/bakery/models"
	"github.com/stretchr/testify/mock"
)

type MockProductStore struct {
	mock.Mock 
}

func (p *MockProductStore) GetAvailableProducts() ([]models.Product, error) {
	fmt.Println("Mocked GetAvailableProducts")
	args := p.Called()
	return args.Get(0).([]models.Product), args.Error(1)
}
