package recipes

import (
	"fmt"
)

// Ingredient is a food ingredient for a product
type Ingredient struct {
	Name     string `json:"name"`
	Quantity int    `json:"quantity"`
}

// Recipe is the ingredients and instructions for creating a product
type Recipe struct {
	ID          string       `json:"id"`
	Ingredients []Ingredient `json:"ingredients"`
	BakeTime    int          `json:"bakeTime"`
}


// Recipes is in memory store of recipe for each product that the bakery sells
var Recipes = map[string]Recipe{
	"1": {
		ID: "1", //"Vanilla cake"
		Ingredients: []Ingredient{
			{Name: "flour", Quantity: 400},
			{Name: "eggs", Quantity: 4},
			{Name: "sugar", Quantity: 400},
		},
		BakeTime: 3,
	},
	"2": {
		ID: "2", //"plain cookie"
		Ingredients: []Ingredient{
			{Name: "flour", Quantity: 300},
			{Name: "butter", Quantity: 200},
			{Name: "sugar", Quantity: 200},
		},
		BakeTime: 1,
	},
	"3": {
		ID: "3", //"Doughnut"
		Ingredients: []Ingredient{
			{Name: "flour", Quantity: 500},
			{Name: "sugar", Quantity: 300},
		},
		BakeTime: 2,
	},
}

// RecipeStorer contains CRUD methods for recipes
type RecipeStorer interface {
	GetRecipe(id string) (Recipe, error)
}

// RecipeStore implements crud operations on recipes
type RecipeStore struct {
}

// GetRecipe returns recipe with given if if it exists. Otherwise nil
func (r *RecipeStore) GetRecipe(id string) (Recipe, error) {
	if v, ok := Recipes[id]; ok {
		return v, nil
	}
	return Recipe{}, fmt.Errorf("recipe Id: %v is not available", id)
}
