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
	options       options
}

type options struct {
	customerCollection *bool //nil if not specified
}

type Option func(options *options) error

func WithCustomerCollect(customerCollect bool) Option {
	return func(options *options) error {
		options.customerCollection = &customerCollect
		return nil
	}
}

// NewOrder returns a  pointer to an Order
func NewOrder(productStorer products.ProductStorer, random random.RandomProvider, opts ...Option) (*Order, error) {
	options := options{}
	for _, opt := range opts {

		err := opt(&options)
		if err != nil {
			fmt.Println("TODO")
			return nil, err
		}
	}
	return &Order{
		productStorer: productStorer,
		Random:        random,
		options:       options,
	}, nil
}

// CreateOrder populates an order with random items
func (o *Order) CreateOrder() error {
	items, err := o.GetRandomItems()
	if err != nil {
		fmt.Println("TODO")
		return err
	}
	o.Items = items
	return nil
}

// GetRandomItems returns random selection of order items
func (o *Order) GetRandomItems() ([]ProductQuantity, error) {
	items := []ProductQuantity{}
	availableProducts, err := o.productStorer.GetAll()
	if err != nil {
		return items, err
	}
	for _, product := range availableProducts {
		productQuantity := ProductQuantity{
			ProductID: product.Name,
			RecipeID:  product.RecipeID,
			Quantity:  o.Random.GetRandom(5),
		}
		if productQuantity.Quantity > 0 {
			items = append(items, productQuantity)
		}
	}
	return items, nil
}

//Ready prints out that order is ready and how the order will be dispatched
func (o *Order) Ready() {
	dispatchType := "customer collection"
	isCustomerCollection := o.options.customerCollection
	if isCustomerCollection == nil || !*isCustomerCollection {
		dispatchType = "delivery"
	}
	fmt.Printf("\nOrder ready for %s", dispatchType)
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
