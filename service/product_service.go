package service

import (
	"fmt"
	"strings"

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

	availableProducts, err := p.productStore.GetAvailableProducts()
	if err != nil {
		return nil, err
	}
	return availableProducts, nil

}

func FormatProducts(products []models.Product) string {
	var sb strings.Builder
	fmt.Fprint(&sb, "We have available:")
	for _, v := range products {
		fmt.Fprintf(&sb, "\n %v", v.Name)
	}
	fmt.Fprintf(&sb, "%v", "\n")
	return sb.String()
}
