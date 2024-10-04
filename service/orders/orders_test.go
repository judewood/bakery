package orders

import (
	"reflect"
	"testing"

	"github.com/judewood/bakery/mocks"
	"github.com/judewood/bakery/models"
)

var availableProducts = []models.Product{
	{Name: "Vanilla cake", RecipeID: "1"},
	{Name: "plain cookie", RecipeID: "2"},
	{Name: "Doughnut", RecipeID: "3"},
}

func TestRandomOrder(t *testing.T) {

	wantedItems := []models.ProductQuantity{
		{ProductID: "Vanilla cake", RecipeID: "1", Quantity: 3},
		{ProductID: "plain cookie", RecipeID: "2", Quantity: 3},
		{ProductID: "Doughnut", RecipeID: "3", Quantity: 3},
	}

	type testCase struct {
		wantItems        []models.ProductQuantity
		wantError        error
		availProducts    []models.Product
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
		mockProductStore := new(mocks.MockProductStore)
		mockProductStore.On("GetAvailableProducts").Return(testCase.availProducts, testCase.availProductsErr)

		mockRandom := mocks.NewMockRandom()
		mockRandom.On("GetRandom", 5).Return(3)

		order, gotError := NewOrder(mockProductStore, mockRandom).RandomOrder()

		if !reflect.DeepEqual(testCase.wantItems, order.Items) {
			t.Errorf("Failed TestRandomOrder. \nWanted %v\nGot %v", wantedItems, order.Items)
		}
		if testCase.wantError == nil && gotError != nil {
			t.Errorf("Failed TestRandomOrder.\nGot Error %v", gotError.Error())
		}
		if testCase.wantError != nil && gotError == nil {
			t.Errorf("Failed TestRandomOrder.\nExpected Error %v", testCase.wantError)
		}
	}
}

func TestFormatOrder(t *testing.T) {
	mockRandom := mocks.NewMockRandom()
	mockRandom.On("GetRandom", 5).Return(3)

	mockProductStore := new(mocks.MockProductStore)
	mockProductStore.On("GetAvailableProducts").Return(availableProducts, nil)

	order, _ := NewOrder(mockProductStore, mockRandom).RandomOrder()
	want := "\n3 of Vanilla cake\n3 of plain cookie\n3 of Doughnut"
	got := order.FormatOrder()

	if got != want {
		t.Errorf("Failed TestFormatOrder. \nWanted:\n *%v*\nGot:\n *%v*", want, got)
	}

}
