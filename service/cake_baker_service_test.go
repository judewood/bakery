package service

import (
	"errors"
	"testing"

	"github.com/judewood/bakery/mocks"
	"github.com/judewood/bakery/models"
)

func TestBake(t *testing.T) {
	mockError := errors.New("Mocked error")
	sampleRecipe := models.Recipe{
		ID: "2", //"plain cookie"
		Ingredients: []models.Ingredient{
			{Name: "flour", Quantity: 300},
			{Name: "butter", Quantity: 200},
			{Name: "sugar", Quantity: 200},
		},
		BakeTime: 0,
	}

	sampleProductQuantity := []models.ProductQuantity{
		{ProductID: "Vanilla Cake",
			RecipeID: "1",
			Quantity: 1,
		},
		{ProductID: "plain cookie",
			RecipeID: "2",
			Quantity: 2,
		},
	}

	type TestCase struct {
		Input models.ProductQuantity
		err   error
	}

	testCases := []TestCase{
		{
			Input: sampleProductQuantity[0],
			err:   mockError,
		},
		{
			Input: sampleProductQuantity[1],
			err:   nil,
		},
	}
	mockRecipeStore := mocks.NewMockRecipeStore()
	for _, testCase := range testCases {
		mockRecipeStore.On("GetRecipe", sampleProductQuantity[0].RecipeID).Return(models.Recipe{}, errors.New("Mocked error"))
		mockRecipeStore.On("GetRecipe", sampleProductQuantity[1].RecipeID).Return(sampleRecipe, nil)

		cakeBaker := NewCakeBaker(mockRecipeStore)
		gotError := cakeBaker.Bake(testCase.Input)

		if testCase.err == nil && gotError != nil {
			t.Errorf("Failed TestBake.\nGot Error %v", gotError.Error())
		}
		if testCase.err != nil && gotError == nil {
			t.Errorf("Failed TestBake.\nExpected Error %v", testCase.err)
		}
	}

}
