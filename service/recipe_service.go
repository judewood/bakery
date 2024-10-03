package service

import (
	"fmt"

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

func (r *RecipeService) GetRecipe(id string) (models.Recipe, error) {
	recipe, err := r.recipeStore.GetRecipe(id)
	if err != nil {
		return models.Recipe{}, err
	}
	fmt.Printf("Got recipe for recipe Id: %v", id)
	return recipe, nil

}
