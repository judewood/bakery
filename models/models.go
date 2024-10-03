package models

const Flour = "flour"
const Sugar = "sugar"
const Eggs = "eggs"
const Butter = "butter"

type Ingredient struct {
	Name     string
	Quantity int
}

type Baker interface {
	Bake(ch chan<- Product)
	Package(ch <-chan Product)
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
	RecipeID string
	Quantity  int
}
