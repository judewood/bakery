package main

import (
	"fmt"
	"log/slog"

	"github.com/gin-gonic/gin"
	"github.com/judewood/bakery/config"
	"github.com/judewood/bakery/internal/products"
	"github.com/judewood/bakery/logger"
)

var Logger *slog.Logger

func main() {
	config := config.New("./environments")
	Logger = logger.GetLogger(config.GetStringSetting("logs.level"))

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
