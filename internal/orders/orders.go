package orders

import (
	"fmt"
	"strings"

	"github.com/judewood/bakery/internal/products"
	"github.com/judewood/bakery/random"
)

// ProductQuantity represents a quantity of a product in an order
type ProductQuantity struct {
	ProductID string `json:"productId"`
	RecipeID  string `json:"recipeId"`
	Quantity  int    `json:"quantity"`
}

// Order is an order for one of more products
type Order struct {
	productStorer products.ProductStorer
	Random        random.RandomProvider
	Items         []ProductQuantity
}

// NewOrder returns a  pointer to an Order
func NewOrder(productStorer products.ProductStorer, random random.RandomProvider) *Order {
	return &Order{
		productStorer: productStorer,
		Random:        random,
	}
}

// RandomOrder generates a random order
func (o *Order) RandomOrder() (Order, error) {
	availableProducts, err := o.productStorer.GetAvailableProducts()
	if err != nil {
		return Order{}, err
	}
	for _, product := range availableProducts {
		productQuantity := ProductQuantity{
			ProductID: product.Name,
			RecipeID:  product.RecipeID,
			Quantity:  o.Random.GetRandom(5),
		}
		if productQuantity.Quantity > 0 {
			o.Items = append(o.Items[:], productQuantity)
		}
	}
	return *o, nil
}

// FormatOrder formats an order for display
func (o *Order) FormatOrder() string {
	var b strings.Builder
	fmt.Print(&b, "\nOrder received for:")
	for _, v := range o.Items {
		fmt.Fprintf(&b, "\n%v of %v", v.Quantity, v.ProductID)
	}
	return b.String()
}
