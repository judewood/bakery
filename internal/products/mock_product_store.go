package products

import (
	"github.com/stretchr/testify/mock"
)

type MockProductStore struct {
	mock.Mock
}

func (p *MockProductStore) GetAll() ([]Product, error) {
	args := p.Called()
	return args.Get(0).([]Product), args.Error(1)
}

func (p *MockProductStore) Get(id string) (Product, error) {
	args := p.Called()
	return args.Get(0).(Product), args.Error(1)
}

func (p *MockProductStore) Add(Product) (Product, error) {
	args := p.Called()
	return args.Get(0).(Product), args.Error(1)
}

func (p *MockProductStore) Update(product Product) (Product, error) {
	args := p.Called()
	return args.Get(0).(Product), args.Error(1)
}

func (p *MockProductStore) Delete(id string) (Product, error) {
	args := p.Called()
	return args.Get(0).(Product), args.Error(1)
}
