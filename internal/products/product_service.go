package products

import (
	"fmt"
	"strings"
)

// ProductServer provides a product service
type ProductServer interface{
GetAvailableProducts() ([]Product, error)
}

// ProductService applies business logic to bakery products
type ProductService struct {
	productStore ProductStorer
}

// NewProductService returns a pointer to a ProductService
func NewProductService(productStore ProductStorer) *ProductService {
	return &ProductService{
		productStore: productStore,
	}
}
// GetAvailableProducts returns a slice of all the bakery's  products 
func (p *ProductService) GetAvailableProducts() ([]Product, error) {

	availableProducts, err := p.productStore.GetAvailableProducts()
	if err != nil {
		return nil, err
	}
	return availableProducts, nil

}

func (p *ProductService) Add(Product) (Product, error) {
	return Product{}, nil
}

// FormatProducts formats a slice of products for display
func FormatProducts(products []Product) string {
	var sb strings.Builder
	fmt.Fprint(&sb, "We have available:")
	for _, v := range products {
		fmt.Fprintf(&sb, "\n %v", v.Name)
	}
	fmt.Fprintf(&sb, "%v", "\n")
	return sb.String()
}

