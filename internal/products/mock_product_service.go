package products

import (
	"fmt"

	"github.com/stretchr/testify/mock"
)

type MockProductService struct {
	mock.Mock
}

func NewMockProductService() *MockProductService {
return &MockProductService{}
}

func(m *MockProductService) GetAvailableProducts() ([]Product, error){
     args := m.Called()
	 fmt.Println("called with", args.Get(0).([]Product))
	 return args.Get(0).([]Product), args.Error(1)
}
