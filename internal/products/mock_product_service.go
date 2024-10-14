package products

import (
	"github.com/stretchr/testify/mock"
)

type MockProductService struct {
	mock.Mock
}

func NewMockProductService() *MockProductService {
	return &MockProductService{}
}

func (m *MockProductService) GetAll() ([]Product, error) {
	args := m.Called()
	return args.Get(0).([]Product), args.Error(1)
}

func (p *MockProductService) Get(id string) (Product, error) {
	args := p.Called()
	return args.Get(0).(Product), args.Error(1)
}

func (m *MockProductService) Add(product Product) (Product, error) {
	args := m.Called()
	return args.Get(0).(Product), args.Error(1)
}

func (p *MockProductService) Update(product Product) (Product, error) {
	args := p.Called()
	return args.Get(0).(Product), args.Error(1)
}

func (p *MockProductService) Delete(id string) (Product, error) {
	args := p.Called()
	return args.Get(0).(Product), args.Error(1)
}

