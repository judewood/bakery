package orders

import (
	"reflect"
	"testing"

	"github.com/judewood/bakery/mocks"
	"github.com/judewood/bakery/models"
)

func TestRandomOrder(t *testing.T) {

	availableProducts := []models.Product{
		{Name: "Vanilla cake", RecipeID: "1"},
		{Name: "plain cookie", RecipeID: "2"},
		{Name: "Doughnut", RecipeID: "3"},
	}

	wantedItems := []models.ProductQuantity{
		{ProductID: "Vanilla cake", RecipeID: "1", Quantity: 3},
		{ProductID: "plain cookie", RecipeID: "2", Quantity: 3},
		{ProductID: "Doughnut", RecipeID: "3", Quantity: 3},
	}

	mockRandom := mocks.NewMockRandom()
	mockRandom.On("GetRandom", 5).Return(3)

	order := NewOrder(mockRandom).RandomOrder(availableProducts)

	if !reflect.DeepEqual(wantedItems, order.Items) {
		t.Errorf("Failed TestRandomOrder. \nWanted %v\nGot %v", wantedItems, order.Items)
	}
}
