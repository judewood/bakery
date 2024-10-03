package main

import (
	"fmt"

	"github.com/judewood/bakery/models"
	"github.com/judewood/bakery/service/bakers"
	"github.com/judewood/bakery/service/orders"
	"github.com/judewood/bakery/service/products"
	"github.com/judewood/bakery/store"
	"github.com/judewood/bakery/utils"
)

func main() {
	productStore := new(store.ProductStore)
	recipeStore := new(store.RecipeStore)

	productService := products.NewProductService(productStore)
	random := new(utils.Random)

	bakerService := bakers.NewCakeBaker(recipeStore)

	availableProducts, _ := productService.GetAvailableProducts()
	fmt.Print(products.FormatProducts(availableProducts))

	order := orders.NewOrder(random).RandomOrder(availableProducts)
	fmt.Print(order.FormatOrder())

	ch := make(chan models.ProductQuantity)
	go func() {
		for _, v := range order.Items {
			ch <- v
		}
		close(ch)
	}()

	for v := range ch {
		err := bakerService.Bake(v)
		if err != nil {
			fmt.Println(err.Error())
		}
	}

	bakerService.Package()

}
