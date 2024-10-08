package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/judewood/bakery/config"
	"github.com/judewood/bakery/internal/products"
)

func main() {
	config := config.New("./environments")
	logLevel := config.GetStringSetting("logs.level")
	fmt.Printf("\nLogging at level: %s", logLevel)
	fmt.Println()
	productStore := new(products.ProductStore)

	productService := products.NewProductService(productStore)
	productController := products.NewProductController(productService)

	server := gin.Default()

	server.GET("/products", func(ctx *gin.Context) {
		ctx.JSON(200, productController.GetProducts())
	})

	server.Run(":8080")
}
