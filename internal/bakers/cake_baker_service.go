package bakers

import (
	"fmt"
	"time"

	"github.com/judewood/bakery/internal/orders"
	"github.com/judewood/bakery/internal/recipes"
)

// CakeBaker bakes the cake using the recipe
type CakeBaker struct {
	recipeStore recipes.RecipeStorer
}

// NewCakeBaker returns address of CakeBaker struct
func NewCakeBaker(recipeStore recipes.RecipeStorer) *CakeBaker {
	return &CakeBaker{
		recipeStore: recipeStore,
	}
}

// Bake will receive Cake products from the rawProducts channel.
// It will Bake the products and then send the baked products onto a new channel for packaging
func (cb *CakeBaker) Bake(rawProduct orders.ProductQuantity) error {
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
	fmt.Println("\n\nOrder packaged and ready to go")
}
