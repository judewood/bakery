package products

import (
	"fmt"
	"log/slog"
	"strings"

	"github.com/judewood/bakery/errorutils"
)

// ProductServer provides a product service
type ProductServer interface {
	GetAll() ([]Product, error)
	Get(id string) (Product, error)
	Add(product Product) (Product, error)
	Update(product Product) (Product, error)
	Delete(id string) (Product, error)
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

// GetAll returns a slice of all the bakery's  products
func (p *ProductService) GetAll() ([]Product, error) {
	availableProducts, err := p.productStore.GetAll()
	if err != nil {
		return nil, err
	}
	slog.Debug("Got products", "products", availableProducts)
	return availableProducts, nil

}

// Get returns product with given id
func (p *ProductService) Get(id string) (Product, error) {
	slog.Debug("Get product by id", "id", id)
	if len(id) < 2 {
		return Product{}, errorutils.ErrorMissingID
	}
	return p.productStore.Get(id)
}

// Add adds new product to store
func (p *ProductService) Add(product Product) (Product, error) {
	if isInvalid, err := isInvalid(product); isInvalid {
		slog.Warn("failed to add product", "Product", product, "Error", err)
		return Product{}, err
	}
	added, err := p.productStore.Add(product)
	slog.Info("Added product", "Product", added)
	return added, err
}

// Update overwrites existing product with given product
func (p *ProductService) Update(product Product) (Product, error) {
	if isInvalid, err := isInvalid(product); isInvalid {
		return Product{}, err
	}
	product, err := p.productStore.Update(product)
	if err != nil {
		return Product{}, err
	}
	return product, nil
}

// Delete deletes an existing product
func (p *ProductService) Delete(id string) (Product, error) {
	if len(id) < 2 {
		return Product{}, errorutils.ErrorMissingID
	}
	return p.productStore.Delete(id)
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
		return true, errorutils.ErrorMissingRequired
	}
	return false, nil
}
