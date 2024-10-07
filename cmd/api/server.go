package main

import (
	"github.com/gin-gonic/gin"
	"github.com/judewood/bakery/controller"
	"github.com/judewood/bakery/internal/products"
	"github.com/judewood/bakery/store"
)

func main() {
	productStore := new(store.ProductStore)

	productService := products.NewProductService(productStore)
	productController := controller.NewProductController(productService)

	server := gin.Default()

	server.GET("/products", func(ctx *gin.Context) {
		ctx.JSON(200, productController.GetProducts())
	})

	server.Run(":8080")
}
