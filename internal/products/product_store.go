package products

import (
	"strings"

	"github.com/judewood/bakery/utils/errorutils"
)

// Product is a saleable food item
type Product struct {
	Name     string `json:"name"`
	RecipeID string `json:"recipeId"`
}

// Persistent store of products - will move to AWS store
var products = []Product{
	{Name: "Vanilla cakes", RecipeID: "cake_base"},
	{Name: "Plain cookies", RecipeID: "cookie_base"},
	{Name: "Doughnuts", RecipeID: "doughnut_base"},
}

// ProductStorer is something that can perform CRUD operations on products
type ProductStorer interface {
	GetAll() ([]Product, error)
	Add(product Product) (Product, error)
	Get(id string) (Product, error)
	Delete(id string) (Product, error)
	Update(product Product) (Product, error)
}

// ProductStore provides crud operations for products
type ProductStore struct {
}

// GetAll returns a slice of all the bakery's  products
func (p *ProductStore) GetAll() ([]Product, error) {
	return products, nil
}

// Get returns matching product or empty product if not found
func (p *ProductStore) Get(id string) (Product, error) {
	for _, v := range products {
		if strings.ToLower(v.Name) == id {
			return v, nil
		}
	}
	return Product{}, errorutils.ErrorNotFound
}

// Add adds given product to the store if it is a valid product
func (p *ProductStore) Add(product Product) (Product, error) {
	products = append(products, product)
	return product, nil
}

// Get returns matching product or empty product if not found
func (p *ProductStore) Update(product Product) (Product, error) {
	for i, v := range products {
		if v.Name == product.Name {
			products[i] = product
			return product, nil
		}
	}
	return Product{}, errorutils.ErrorNotFound
}

// Delete deletes and returns matching product or empty product if not found
func (p *ProductStore) Delete(id string) (Product, error) {
	for i, v := range products {
		if strings.ToLower(v.Name) == id {
			if i < len(products)-1 {
				products = append(products[:i], products[i+1:]...)
				return v, nil
			}
			products = products[:i]
			return v, nil
		}

	}
	return Product{}, errorutils.ErrorNotFound
}
