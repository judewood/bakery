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
	if args.Get(1) == nil { //no error
		return args.Get(0).([]models.Product), nil
	}
	return args.Get(0).([]models.Product), args.Get(1).(error)
}
