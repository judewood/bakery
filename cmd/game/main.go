package main

import (
	"fmt"
	"log"
	"log/slog"

	"github.com/judewood/bakery/config"
	"github.com/judewood/bakery/internal/bakers"
	"github.com/judewood/bakery/internal/orders"
	"github.com/judewood/bakery/internal/products"
	"github.com/judewood/bakery/internal/recipes"
	"github.com/judewood/bakery/internal/s3client"
	"github.com/judewood/bakery/logger"
	"github.com/judewood/bakery/random"
)

func main() {
	config := config.New("./environments")
	logger.InitLogger(config.GetStringSetting("logs.level"))

	slog.Debug("Bakery is open for business")
	productStore := new(products.ProductStore)
	recipeStore := recipes.New(config.GetStringSetting("s3.recipesUrl"), s3client.New())
	recipeCache := recipes.NewRecipeCache(recipeStore)
	//recipeStore.GetRecipeFromS3(config.GetStringSetting("s3.recipesUrl"))
	random := new(random.Random)

	bakerService := bakers.NewCakeBaker(recipeCache)

	order, err := orders.NewOrder(productStore, random, orders.WithCustomerCollect(true))
	if err != nil {
		log.Fatal(err)
	}
	order.CreateOrder()

	fmt.Print(order.FormatOrder())

	ch := make(chan orders.ProductQuantity, len(order.Items))
	go func() {
		for _, v := range order.Items {
			ch <- v
		}
		close(ch)
	}()

	for product := range ch {
		err := bakerService.Bake(product)
		if err != nil {
			fmt.Println(err.Error())
		}
	}

	order.Ready()
}
