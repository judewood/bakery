package store

import "github.com/judewood/bakery/models"

var products = []models.Product{
	{Name: "Vanilla cakes", RecipeID: "1"},
	{Name: "Plain cookies", RecipeID: "2"},
	{Name: "Doughnuts", RecipeID: "3"},
}

// ProductStorer contains set of Product CRUD operations
type ProductStorer interface {
	GetAvailableProducts() ([]models.Product, error)
}

// ProductStore provides crud operations for products
type ProductStore struct {
}

// GetAvailableProducts returns a slice of all the bakery's  products 
func (p *ProductStore) GetAvailableProducts() ([]models.Product, error) {
	return products, nil
}
