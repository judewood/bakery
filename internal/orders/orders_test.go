package orders

import (
	"reflect"
	"testing"

	"github.com/judewood/bakery/internal/products"
	"github.com/judewood/bakery/myfmt"
	"github.com/judewood/bakery/random"
)

var availableProducts = []products.Product{
	{Name: "Vanilla cake", RecipeID: "1"},
	{Name: "plain cookie", RecipeID: "2"},
	{Name: "Doughnut", RecipeID: "3"},
}

func TestRandomOrder(t *testing.T) {

	wantedItems := []ProductQuantity{
		{ProductID: "Vanilla cake", RecipeID: "1", Quantity: 3},
		{ProductID: "plain cookie", RecipeID: "2", Quantity: 3},
		{ProductID: "Doughnut", RecipeID: "3", Quantity: 3},
	}

	type testCase struct {
		wantItems        []ProductQuantity
		wantError        error
		availProducts    []products.Product
		availProductsErr error
	}

	testCases := []testCase{
		{
			wantItems:        wantedItems,
			wantError:        nil,
			availProducts:    availableProducts,
			availProductsErr: nil,
		},
	}

	for _, testCase := range testCases {
		mockProductStore := new(products.MockProductStore)
		mockProductStore.On("GetAll").Return(testCase.availProducts, testCase.availProductsErr)

		mockRandom := random.NewMockRandom()
		mockRandom.On("GetRandom", 5).Return(3)

		order, gotError := NewOrder(mockProductStore, mockRandom).RandomOrder()

		if !reflect.DeepEqual(testCase.wantItems, order.Items) {
			myfmt.Errorf(t, "Failed TestRandomOrder. \nWanted %v\nGot %v", wantedItems, order.Items)
		}
		if testCase.wantError == nil && gotError != nil {
			myfmt.Errorf(t, "Failed TestRandomOrder.\nGot Error %v", gotError.Error())
		}
		if testCase.wantError != nil && gotError == nil {
			myfmt.Errorf(t, "Failed TestRandomOrder.\nExpected Error %v", testCase.wantError)
		}
		t.Log()
	}
}

func TestFormatOrder(t *testing.T) {
	mockRandom := random.NewMockRandom()
	mockRandom.On("GetRandom", 5).Return(3)

	mockProductStore := new(products.MockProductStore)
	mockProductStore.On("GetAll").Return(availableProducts, nil)

	order, _ := NewOrder(mockProductStore, mockRandom).RandomOrder()
	want := "\n3 of Vanilla cake\n3 of plain cookie\n3 of Doughnut"
	got := order.FormatOrder()

	if got != want {
		t.Errorf("Failed TestFormatOrder. \nWanted:\n *%v*\nGot:\n *%v*", want, got)
	}

}
