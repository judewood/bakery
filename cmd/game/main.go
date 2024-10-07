package main

import (
	"fmt"
	"log"

	"github.com/judewood/bakery/internal/bakers"
	"github.com/judewood/bakery/internal/orders"
	"github.com/judewood/bakery/internal/products"
	"github.com/judewood/bakery/internal/recipes"
	"github.com/judewood/bakery/random"
)

func main() {
	productStore := new(products.ProductStore)
	recipeStore := new(recipes.RecipeStore)

	random := new(random.Random)

	bakerService := bakers.NewCakeBaker(recipeStore)

	order, err := orders.NewOrder(productStore, random).RandomOrder()
	if err != nil {
		log.Panic(err)
	}

	fmt.Print(order.FormatOrder())

	ch := make(chan orders.ProductQuantity)
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
