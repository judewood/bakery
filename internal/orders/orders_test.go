package orders

import (
	"reflect"
	"testing"

	"github.com/judewood/bakery/internal/products"
	"github.com/judewood/bakery/random"
	"github.com/judewood/bakery/utils/testutils"
)

var availableProducts = []products.Product{
	{Name: "Vanilla cake", RecipeID: "cake_base"},
	{Name: "plain cookie", RecipeID: "cookie_base"},
	{Name: "Doughnut", RecipeID: "doughnut_base"},
}

func TestRandomOrder(t *testing.T) {

	wantedItems := []ProductQuantity{
		{ProductID: "Vanilla cake", RecipeID: "cake_base", Quantity: 3},
		{ProductID: "plain cookie", RecipeID: "cookie_base", Quantity: 3},
		{ProductID: "Doughnut", RecipeID: "doughnut_base", Quantity: 3},
	}

	type testCase struct {
		wantItems        []ProductQuantity
		availProducts    []products.Product
		availProductsErr error
	}

	testCases := []testCase{
		{wantItems: wantedItems, availProducts: availableProducts, availProductsErr: nil}}

	for _, testCase := range testCases {
		t.Log("Given I want to create a order containing random products")
		{
			mockProductStore := new(products.MockProductStore)
			mockProductStore.On("GetAll").Return(testCase.availProducts, testCase.availProductsErr)

			mockRandom := random.NewMockRandom()
			mockRandom.On("GetRandom", 5).Return(3)

			t.Log("When I request an order")
			{
				order, gotError := NewOrder(mockProductStore, mockRandom).RandomOrder()
				t.Logf("Then I should get expected order :%v", wantedItems)
				{
					if reflect.DeepEqual(testCase.wantItems, order.Items) {
						testutils.Passed(t)
					} else {
						testutils.Failed(t, order.Items)
					}
				}
				t.Log("and no error should be returned")
				{
					if gotError == nil {
						testutils.Passed(t)
					} else {
						testutils.Failed(t, gotError)
					}
				}
			}
		}
	}
}

func TestFormatOrder(t *testing.T) {
	t.Log("Given I have an order")
	{
		mockRandom := random.NewMockRandom()
		mockRandom.On("GetRandom", 5).Return(3)

		mockProductStore := new(products.MockProductStore)
		mockProductStore.On("GetAll").Return(availableProducts, nil)

		order, _ := NewOrder(mockProductStore, mockRandom).RandomOrder()
		want := "\n3 of Vanilla cake\n3 of plain cookie\n3 of Doughnut"

		t.Log("When the order is formatted")
		{
			got := order.FormatOrder()
			t.Logf("The formatted order should be :%s", want)
			{
				if got == want {
					testutils.Passed(t)
				} else {
					testutils.Failed(t, got)
				}
			}
		}
	}
}
