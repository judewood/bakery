package store

import "github.com/judewood/bakery/models"

var products = []models.Product{
	{Name: "Vanilla cakes", RecipeID: "1"},
	{Name: "Plain cookies", RecipeID: "2"},
	{Name: "Doughnuts", RecipeID: "3"},
}

type IProductStore interface {
	GetAvailableProducts() ([]models.Product, error)
}

// RecipeStore provides crud operations on the persistent store of product recipes
type ProductStore struct {
	AvailableProducts []models.Product
}

// New returns pointer to RecipeStore
func NewProductStore() *ProductStore {
	return &ProductStore{
		AvailableProducts: products,
	}
}

func (p *ProductStore) GetAvailableProducts() ([]models.Product, error) {
	return products, nil
}
