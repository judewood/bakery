package products

import (
	"fmt"
	"strings"

	"github.com/judewood/bakery/models"
	"github.com/judewood/bakery/store"
)

type ProductServer interface{
GetAvailableProducts() ([]models.Product, error)
}

// ProductService applies business logic to bakery products
type ProductService struct {
	productStore store.ProductStorer
}

// NewProductService returns a pointer to a ProductService
func NewProductService(productStore store.ProductStorer) *ProductService {
	return &ProductService{
		productStore: productStore,
	}
}
// GetAvailableProducts returns a slice of all the bakery's  products 
func (p *ProductService) GetAvailableProducts() ([]models.Product, error) {

	availableProducts, err := p.productStore.GetAvailableProducts()
	if err != nil {
		return nil, err
	}
	return availableProducts, nil

}

func (p *ProductService) Add(models.Product) (models.Product, error) {
	return models.Product{}, nil
}

// FormatProducts formats a slice of products for display
func FormatProducts(products []models.Product) string {
	var sb strings.Builder
	fmt.Fprint(&sb, "We have available:")
	for _, v := range products {
		fmt.Fprintf(&sb, "\n %v", v.Name)
	}
	fmt.Fprintf(&sb, "%v", "\n")
	return sb.String()
}

