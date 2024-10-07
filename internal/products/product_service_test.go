package products

import (
	"errors"
	"reflect"
	"testing"
)

var sampleProducts = []Product{
	{Name: "Vanilla cake", RecipeID: "1"},
	{Name: "plain cookie", RecipeID: "2"},
	{Name: "Doughnut", RecipeID: "3"},
}

func TestGetAvailableProducts(t *testing.T) {
	type testCase struct {
		want []Product
		err  error
	}

	mockError := errors.New("Mocked error")
	testCases := []testCase{
		{want: nil, err: mockError},
		{want: []Product{}, err: nil},
		{want: sampleProducts, err: nil},
	}

	for _, testCase := range testCases {
		mockProductStore := new(MockProductStore)
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

func TestFormatProducts(t *testing.T) {
	want := "We have available:\n Vanilla cake\n plain cookie\n Doughnut\n"

	got := FormatProducts(sampleProducts)

	if got != want {
		t.Errorf("Failed TestFormatProducts. \nWanted:\n *%v*\nGot:\n *%v*", want, got)
	}
}
