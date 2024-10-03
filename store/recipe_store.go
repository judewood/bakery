package store

import (
	"fmt"

	"github.com/judewood/bakery/models"
)

// Recipes is he recipe for each product that the bakery sells
var Recipes = map[string]models.Recipe{
	"1": {
		ID: "1", //"Vanilla cake"
		Ingredients: []models.Ingredient{
			{Name: models.Flour, Quantity: 400},
			{Name: models.Eggs, Quantity: 4},
			{Name: models.Sugar, Quantity: 400},
		},
		BakeTime: 3,
	},
	"2": {
		ID: "2", //"plain cookie"
		Ingredients: []models.Ingredient{
			{Name: models.Flour, Quantity: 300},
			{Name: models.Butter, Quantity: 200},
			{Name: models.Sugar, Quantity: 200},
		},
		BakeTime: 1,
	},
	"3": {
		ID: "3", //"Doughnut"
		Ingredients: []models.Ingredient{
			{Name: models.Flour, Quantity: 500},
			{Name: models.Sugar, Quantity: 300},
		},
		BakeTime: 2,
	},
}

type IRecipeStore interface{
	 GetRecipe(id string) (models.Recipe, error)
}

// RecipeStore provides crud operations on the persistent store of product recipes
type RecipeStore struct {
	AvailableRecipes map[string]models.Recipe
}

// New returns pointer to RecipeStore
func NewRecipeStore() *RecipeStore {
	return &RecipeStore{
		AvailableRecipes: Recipes,
	}
}

// GetRecipe returns recipe with given if if it exists. Otherwise nil
func (r *RecipeStore) GetRecipe(id string) (models.Recipe, error) {
	if v, ok := r.AvailableRecipes[id]; ok {
		return v, nil
	}
	return models.Recipe{}, fmt.Errorf("recipe Id: %v is not available", id)
}
