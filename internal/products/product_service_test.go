package products

import (
	"errors"
	"reflect"
	"testing"

	"github.com/judewood/bakery/myfmt"
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
			myfmt.Errorf(t, "Failed TestGetAvailableProducts.\nWanted: %v \nGot %v", testCase.want, gotProducts)
		}

		if testCase.err == nil && gotError != nil {
			myfmt.Errorf(t, "Failed TestGetAvailableProducts.\nGot Error %v", gotError.Error())
		}

		if testCase.err != nil && gotError == nil {
			myfmt.Errorf(t, "Failed TestGetAvailableProducts.\nExpected Error %v", testCase.err.Error())
		}

	}
}

func TestFormatProducts(t *testing.T) {
	want := "We have available:\n Vanilla cake\n plain cookie\n Doughnut\n"

	got := FormatProducts(sampleProducts)

	if got != want {
		myfmt.Errorf(t, "Failed TestFormatProducts. \nWanted:\n *%v*\nGot:\n *%v*", want, got)
	}
}
