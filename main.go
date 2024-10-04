package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/judewood/bakery/controller"
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
	productController := controller.NewProductController(productService)

	random := new(utils.Random)

	bakerService := bakers.NewCakeBaker(recipeStore)

	server := gin.Default()

	server.GET("/products", func(ctx *gin.Context) {
		ctx.JSON(200, productController.GetProducts())
	})

	//server.Run(":8080")

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
