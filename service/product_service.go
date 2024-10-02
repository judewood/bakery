package service

import (
	"github.com/judewood/bakery/models"
	"github.com/judewood/bakery/store"
)

type ProductService struct {
	productStore store.ProductStore
}

func NewProductService(productStore store.ProductStore) *ProductService {
	return &ProductService{
		productStore: productStore,
	}
}

func (p *ProductService) GetAvailableProducts() []models.Product {
	return p.productStore.GetAvailableProducts()
}
