package products

import (
	"fmt"

	"github.com/stretchr/testify/mock"
)

type MockProductStore struct {
	mock.Mock 
}

func (p *MockProductStore) GetAvailableProducts() ([]Product, error) {
	fmt.Println("Mocked GetAvailableProducts")
	args := p.Called()
	return args.Get(0).([]Product), args.Error(1)
}

func (p *MockProductStore) AddProduct(Product)(Product, error) {
	fmt.Println("Mocked AddProduct")
	args := p.Called()
	return args.Get(0).(Product), args.Error(1)
}
