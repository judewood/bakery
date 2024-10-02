package models

const Flour = "flour"
const Sugar = "sugar"
const Eggs = "eggs"
const Butter = "butter"

type Ingredient struct {
	Name     string
	Quantity int
}

type Bake struct {
	Minutes int
}
type Recipe struct {
	ID          string
	Ingredients []Ingredient
	Method      Bake
}

type Product struct {
	Name     string
	RecipeID string
}

type ProductQuantity struct {
	ProductID string
	Quantity  int
}
