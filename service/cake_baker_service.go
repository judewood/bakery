package service

import (
	"fmt"
	"time"

	"github.com/judewood/bakery/models"
	"github.com/judewood/bakery/store"
)

type CakeBaker struct {
	recipeStore store.IRecipeStore
}

func NewCakeBaker(recipeStore store.IRecipeStore) *CakeBaker {
	return &CakeBaker{
		recipeStore: recipeStore,
	}
}

// Bake will receive Cake products from the rawProducts channel.
// It will Bake the products and then send the baked products onto a new channel for packaging
func (cb *CakeBaker) Bake(rawProduct models.ProductQuantity) error {
	recipe, err := cb.recipeStore.GetRecipe(rawProduct.RecipeID)
	if err != nil {
		return err
	}
	fmt.Printf("\nBaking %v of %v for %v seconds...", rawProduct.Quantity, rawProduct.ProductID, recipe.BakeTime)
	time.Sleep(time.Duration(recipe.BakeTime) * time.Second)
	return nil
}

// Package - still to be implemented
func (cb *CakeBaker) Package() {
	fmt.Println("\n Order packaged")
}
