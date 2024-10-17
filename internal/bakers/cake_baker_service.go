package bakers

import (
	"log/slog"
	"time"

	"github.com/judewood/bakery/internal/orders"
	"github.com/judewood/bakery/internal/recipes"
)

// CakeBaker bakes the cake using the recipe
type CakeBaker struct {
	recipeCache recipes.RecipeCacher
}

// NewCakeBaker returns address of CakeBaker struct
func NewCakeBaker(recipeCache recipes.RecipeCacher) *CakeBaker {
	return &CakeBaker{
		recipeCache: recipeCache,
	}
}

// Bake will receive Cake products from the rawProducts channel.
// It will Bake the products and then send the baked products onto a new channel for packaging
func (cb *CakeBaker) Bake(rawProduct orders.ProductQuantity) error {
	recipe, err := cb.recipeCache.GetRecipe(rawProduct.RecipeID)
	if err != nil {
		return err
	}
	slog.Debug("Baking", "Count", rawProduct.Quantity, "Product Id", rawProduct.ProductID, "Duration", recipe.BakeTime)
	time.Sleep(time.Duration(recipe.BakeTime) * time.Second)
	return nil
}

// Package - still to be implemented
func (cb *CakeBaker) Package() {
	slog.Debug("Order packaged and ready to go")
}
