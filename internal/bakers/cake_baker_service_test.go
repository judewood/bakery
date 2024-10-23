package bakers

import (
	"testing"

	"github.com/judewood/bakery/internal/orders"
	"github.com/judewood/bakery/internal/recipes"
	"github.com/judewood/bakery/utils/errorutils"
	"github.com/judewood/bakery/utils/testutils"
)

func TestBake(t *testing.T) {
	sampleRecipe := recipes.Recipe{
		ID: "2",
		Ingredients: []recipes.Ingredient{
			{Name: "flour", Quantity: 300},
			{Name: "butter", Quantity: 200},
			{Name: "sugar", Quantity: 200},
		},
		BakeTime: 0,
	}

	sampleProductQuantity := []orders.ProductQuantity{
		{ProductID: "Vanilla Cake", RecipeID: "cake_base", Quantity: 1},
		{ProductID: "plain cookie", RecipeID: "cookie_base", Quantity: 2},
	}

	type TestCase struct {
		name  string
		input orders.ProductQuantity
		want  recipes.Recipe
		err   error
	}

	testCases := []TestCase{
		{"found", sampleProductQuantity[0], sampleRecipe, nil},
		{"missing", sampleProductQuantity[1], recipes.Recipe{}, errorutils.ErrorNotFound},
	}
	mockRecipeCache := recipes.NewMockRecipeCache()
	for i, test := range testCases {
		tf := func(t *testing.T) {
			t.Log("Given that I need to bake a product")
			{
				t.Logf("test %d: When the ingredients are %s", i, test.name)
				{
					mockRecipeCache.On("GetRecipe", test.input.RecipeID).Return(recipes.Recipe{}, test.err)

					cakeBaker := NewCakeBaker(mockRecipeCache)
					gotError := cakeBaker.Bake(test.input)
					if test.err == nil {
						t.Log("Then Bake should return no error")
						{
							if gotError == nil {
								testutils.Passed(t)
							} else {
								testutils.Failed(t, gotError)
							}
						}
					} else {
						t.Logf("Then Bake should return error %s", test.err)
						{
							if gotError == nil || gotError.Error() != test.err.Error() {
								testutils.Failed(t, gotError)
							} else {
								testutils.Passed(t)
							}
						}
					}

				}
			}
		}
		t.Run(test.name, tf)
	}
}
