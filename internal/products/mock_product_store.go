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
