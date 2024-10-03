package models

const Flour = "flour"
const Sugar = "sugar"
const Eggs = "eggs"
const Butter = "butter"

type Ingredient struct {
	Name     string
	Quantity int
}

type Recipe struct {
	ID          string
	Ingredients []Ingredient
	BakeTime    int
}

type Product struct {
	Name     string
	RecipeID string
}

type ProductQuantity struct {
	ProductID string
	RecipeID  string
	Quantity  int
}
