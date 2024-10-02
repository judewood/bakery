package service

import (
	"fmt"
	"time"

	"github.com/judewood/bakery/models"
	"github.com/judewood/bakery/store"
)

type RecipeService struct {
	recipeStore store.RecipeStore
}

func New(recipeStore store.RecipeStore) *RecipeService {
	return &RecipeService{
		recipeStore: recipeStore,
	}
}

func (r *RecipeService) GetRecipe(id string) (*models.Recipe, error) {
	recipe, err := r.recipeStore.GetRecipe(id)
	if err != nil {
		return nil, err
	}
	fmt.Printf("Got recipe for recipe Id: %v", id)
	return recipe, nil

}

// MixIngredients currently just prints out. Will utilise available mixers later
func (r *RecipeService) MixIngredients(id string) {
	fmt.Printf("Mixed ingredients for recipe Id: %v", id)
}

// BakeProduct simulates baking products by sleeping for the bake time
func (r *RecipeService) BakeProduct(bakeTime int) {
	// Our bakery uses super fast ovens so bake time is in seconds
	time.Sleep(time.Duration(bakeTime) * time.Second)
	fmt.Printf("Baked recipe for %v seconds", bakeTime)
}
