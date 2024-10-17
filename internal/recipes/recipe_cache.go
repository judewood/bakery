package recipes

type RecipeCacher interface {
	GetRecipe(id string) (Recipe, error)
}

type RecipeCache struct {
	recipes     map[string]Recipe
	recipeStore RecipeStorer
}

func NewRecipeCache(recipeStore RecipeStorer) *RecipeCache {
	return &RecipeCache{
		recipes:     make(map[string]Recipe),
		recipeStore: recipeStore,
	}
}

func (r *RecipeCache) GetRecipe(id string) (Recipe, error) {
	if recipe, ok := r.recipes[id]; ok {
		return recipe, nil
	}
	recipe, err := r.recipeStore.GetRecipeFromS3(id)
	if err != nil {
		return Recipe{}, err
	}
	return recipe, nil
}
