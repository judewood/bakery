package store

import (
	"fmt"

	"github.com/judewood/bakery/models"
)

// Recipes is in memory store of recipe for each product that the bakery sells
var Recipes = map[string]models.Recipe{
	"1": {
		ID: "1", //"Vanilla cake"
		Ingredients: []models.Ingredient{
			{Name: "flour", Quantity: 400},
			{Name: "eggs", Quantity: 4},
			{Name: "sugar", Quantity: 400},
		},
		BakeTime: 3,
	},
	"2": {
		ID: "2", //"plain cookie"
		Ingredients: []models.Ingredient{
			{Name: "flour", Quantity: 300},
			{Name: "butter", Quantity: 200},
			{Name: "sugar", Quantity: 200},
		},
		BakeTime: 1,
	},
	"3": {
		ID: "3", //"Doughnut"
		Ingredients: []models.Ingredient{
			{Name: "flour", Quantity: 500},
			{Name: "sugar", Quantity: 300},
		},
		BakeTime: 2,
	},
}

// RecipeStorer contains CRUD methods for recipes
type RecipeStorer interface {
	GetRecipe(id string) (models.Recipe, error)
}

// RecipeStore implements crud operations on recipes
type RecipeStore struct {
}

// GetRecipe returns recipe with given if if it exists. Otherwise nil
func (r *RecipeStore) GetRecipe(id string) (models.Recipe, error) {
	if v, ok := Recipes[id]; ok {
		return v, nil
	}
	return models.Recipe{}, fmt.Errorf("recipe Id: %v is not available", id)
}
