package main

import (
	"github.com/gin-gonic/gin"
	"github.com/judewood/bakery/internal/products"
)

func main() {
	productStore := new(products.ProductStore)

	productService := products.NewProductService(productStore)
	productController := products.NewProductController(productService)

	server := gin.Default()

	server.GET("/products", func(ctx *gin.Context) {
		ctx.JSON(200, productController.GetProducts())
	})

	server.Run(":8080")
}
