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

	mockRandom := mocks.NewMockRandom()
	mockRandom.On("GetRandom", 5).Return(3)

	order := NewOrder(mockRandom).RandomOrder(availableProducts)

	if !reflect.DeepEqual(wantedItems, order.Items) {
		t.Errorf("Failed TestRandomOrder. \nWanted %v\nGot %v", wantedItems, order.Items)
	}
}

func TestFormatOrder(t *testing.T) {
	mockRandom := mocks.NewMockRandom()
	mockRandom.On("GetRandom", 5).Return(3)

	order := NewOrder(mockRandom).RandomOrder(availableProducts)
	want := "\n3 of Vanilla cake\n3 of plain cookie\n3 of Doughnut"
	got := order.FormatOrder()

	if got != want {
		t.Errorf("Failed TestFormatOrder. \nWanted:\n *%v*\nGot:\n *%v*", want, got)
	}

}
