package service

import (
	"github.com/judewood/bakery/models"
	"github.com/judewood/bakery/store"
)

type ProductService struct {
	productStore store.IProductStore
}

func NewProductService(productStore store.IProductStore) *ProductService {
	return &ProductService{
		productStore: productStore,
	}
}

func (p *ProductService) GetAvailableProducts() ([]models.Product, error) {
	
	availableProducts, err:=  p.productStore.GetAvailableProducts()
	if err != nil {
		return nil, err
	}
	return availableProducts, nil

}
