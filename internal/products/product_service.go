package products

import (
	"errors"
	"fmt"
	"log/slog"
	"strings"
)

const MissingRequired string = "missing required field"

// ProductServer provides a product service
type ProductServer interface {
	GetAvailableProducts() ([]Product, error)
	Add(product Product) (Product, error)
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
	slog.Debug("Got products", "products", availableProducts)
	return availableProducts, nil

}

// Add adds new product to store
func (p *ProductService) Add(product Product) (Product, error) {
	if isInvalid, err := isInvalid(product); isInvalid {
		return Product{}, err
	}
	added, err := p.productStore.AddProduct(product)
	slog.Info("Added product", "Product", added)
	return added, err
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

// isInvalid returns whether this is a valid product and error if it is not valid
func isInvalid(product Product) (bool, error) {
	if product.Name == "" || product.RecipeID == "" {
		return true, errors.New(MissingRequired)
	}
	return false, nil
}
