package service

import (
	"errors"
	"reflect"
	"testing"

	"github.com/judewood/bakery/mocks"
	"github.com/judewood/bakery/models"
)

func TestGetAvailableProducts(t *testing.T) {

	type testCase struct {
		want []models.Product
		err  error
	}

	sampleProducts := []models.Product{
		{Name: "Vanilla cake", RecipeID: "1"},
		{Name: "plain cookie", RecipeID: "2"},
		{Name: "Doughnut", RecipeID: "3"},
	}

	mockError := errors.New("Mocked error")
	testCases := []testCase{
		{want: nil, err: mockError},
		{want: []models.Product{}, err: nil},
		{want: sampleProducts, err: nil},
	}

	for _, testCase := range testCases {
		mockProductStore := new(mocks.MockProductStore)
		mockProductStore.On("GetAvailableProducts").Return(testCase.want, testCase.err)

		productService := NewProductService(mockProductStore)

		gotProducts, gotError := productService.GetAvailableProducts()

		if !reflect.DeepEqual(gotProducts, testCase.want) {
			t.Errorf("Failed TestGetAvailableProducts.\nWanted: %v \nGot %v", testCase.want, gotProducts)
		}

		if testCase.err == nil && gotError != nil {
			t.Errorf("Failed TestGetAvailableProducts.\nGot Error %v", gotError.Error())
		}

		if testCase.err != nil && gotError == nil {
			t.Errorf("Failed TestGetAvailableProducts.\nExpected Error %v", testCase.err.Error())
		}

	}
}
