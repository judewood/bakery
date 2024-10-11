package main

import (
	"fmt"
	"log/slog"

	"github.com/judewood/bakery/config"
	"github.com/judewood/bakery/internal/products"
	"github.com/judewood/bakery/internal/router"
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
	server := router.SetupRouter()
	server = router.AddRouteGetProducts(server, productController.GetProducts)
	//server = server.GET("/products",  &productController)

	server.Run(":8080")
}
