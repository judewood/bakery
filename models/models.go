package models

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

// Product is a saleable food item
type Product struct {
	Name     string `json:"name"`
	RecipeID string `json:"recipeId"`
}

// ProductQuantity represents a quantity of a product in an order
type ProductQuantity struct {
	ProductID string `json:"productId"`
	RecipeID  string `json:"recipeId"`
	Quantity  int    `json:"quantity"`
}
