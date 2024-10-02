package store

import "github.com/judewood/bakery/models"

var Products = []models.Product{
	{Name: "Vanilla cake", RecipeID: "1"},
	{Name: "plain cookie", RecipeID: "2"},
	{Name: "Doughnut", RecipeID: "3"},
}

// RecipeStore provides crud operations on the persistent store of product recipes
type ProductStore struct {
	AvailableProducts []models.Product
}

// New returns pointer to RecipeStore
func NewProductStore() *ProductStore {
	return &ProductStore{
		AvailableProducts: Products,
	}
}

func (p * ProductStore) GetAvailableProducts() []models.Product {
	return Products
}
