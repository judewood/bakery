package main

import (
	"fmt"

	"github.com/judewood/bakery/orders"
	"github.com/judewood/bakery/service"
	"github.com/judewood/bakery/store"
	"github.com/judewood/bakery/utils"
)

func main() {
	productStore := store.NewProductStore()
	productService := service.NewProductService(productStore)
	random := utils.NewRandom()

	availableProducts, _ := productService.GetAvailableProducts()
	fmt.Print(productService.FormatProducts(availableProducts))

	order := orders.NewOrder(random).RandomOrder(availableProducts)
	fmt.Print(order.FormatOrder())
}
