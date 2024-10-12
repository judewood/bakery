package products

import (
	"fmt"
)

// Product is a saleable food item
type Product struct {
	Name     string `json:"name"`
	RecipeID string `json:"recipeId"`
}

var products = []Product{
	{Name: "Vanilla cakes", RecipeID: "1"},
	{Name: "Plain cookies", RecipeID: "2"},
	{Name: "Doughnuts", RecipeID: "3"},
}

// ProductStorer is something that can perform CRUD operations on stored products
type ProductStorer interface {
	GetAvailableProducts() ([]Product, error)
	AddProduct(product Product) (Product, error)
}

// ProductStore provides crud operations for products
type ProductStore struct {
}

// GetAvailableProducts returns a slice of all the bakery's  products
func (p *ProductStore) GetAvailableProducts() ([]Product, error) {
	return products, nil
}

func (p *ProductStore) AddProduct(product Product) (Product, error) {
	fmt.Printf("\nAdd new product %v", product)
	products = append(products, product)
	return product, nil
}
