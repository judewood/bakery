package orders

import (
	"fmt"
	"strings"

	"github.com/judewood/bakery/models"
	"github.com/judewood/bakery/utils"
)

type Order struct {
	Random utils.IRandom
	Items  []models.ProductQuantity
}

func NewOrder(random utils.IRandom) *Order {
	return &Order{
		Random: random,
	}
}

func (o *Order) RandomOrder(availableProducts []models.Product) *Order {

	for _, product := range availableProducts {
		productQuantity := models.ProductQuantity{
			ProductID:  product.Name,
			RecipeID: product.RecipeID,
			Quantity: o.Random.GetRandom(5),
		}
		if productQuantity.Quantity > 0 {
			o.Items = append(o.Items[:], productQuantity)
		}
	}
	return o
}

func (o *Order) FormatOrder() string {
	var b strings.Builder
	fmt.Print(&b, "\nOrder received for:")
	for _, v := range o.Items {
		fmt.Fprintf(&b, "\n%v of %v", v.Quantity, v.ProductID)
	}
	return b.String()
}
