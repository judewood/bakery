package main

import (
	"fmt"
	"sync"

	"github.com/judewood/bakery/models"
	"github.com/judewood/bakery/orders"
	"github.com/judewood/bakery/service"
	"github.com/judewood/bakery/store"
	"github.com/judewood/bakery/utils"
)

var wg sync.WaitGroup

func main() {
	productStore := store.NewProductStore()
	recipeStore := store.NewRecipeStore()

	productService := service.NewProductService(productStore)
	random := utils.NewRandom()

	bakerService := service.NewCakeBaker(recipeStore)

	availableProducts, _ := productService.GetAvailableProducts()
	fmt.Print(productService.FormatProducts(availableProducts))

	order := orders.NewOrder(random).RandomOrder(availableProducts)
	fmt.Print(order.FormatOrder())

	ch := make(chan models.ProductQuantity)
	wg.Add(2)
	go func() {
		defer wg.Done()
		for _, v := range order.Items {
			ch <- v
		}
		close(ch)
	}()

	go func() {
		defer wg.Done()
		for v := range ch {
			err := bakerService.Bake(v)
			if err != nil {
				fmt.Println(err.Error())
			}
		}
	}()
	wg.Wait()

	bakerService.Package()

}
