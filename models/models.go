package models

// Ingredient is a food ingredient for a product
type Ingredient struct {
	Name     string
	Quantity int
}

// Recipe is the ingredients and instructions for creating a product
type Recipe struct {
	ID          string
	Ingredients []Ingredient
	BakeTime    int
}

// Product is a saleable food item
type Product struct {
	Name     string
	RecipeID string
}

// ProductQuantity represents a quantity of a product in an order
type ProductQuantity struct {
	ProductID string
	RecipeID  string
	Quantity  int
}
