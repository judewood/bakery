package main

import (
	"fmt"
	"log"

	"github.com/judewood/bakery/internal/bakers"
	"github.com/judewood/bakery/internal/orders"
	"github.com/judewood/bakery/models"
	"github.com/judewood/bakery/store"
	"github.com/judewood/bakery/utils"
)

func main() {
	productStore := new(store.ProductStore)
	recipeStore := new(store.RecipeStore)

	random := new(utils.Random)

	bakerService := bakers.NewCakeBaker(recipeStore)

	order, err := orders.NewOrder(productStore, random).RandomOrder()
	if err != nil {
		log.Panic(err)
	}

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
