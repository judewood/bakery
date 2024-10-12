package main

import (
	"log/slog"

	"github.com/judewood/bakery/config"
	"github.com/judewood/bakery/internal/products"
	"github.com/judewood/bakery/internal/router"
	"github.com/judewood/bakery/logger"
)

func main() {
	config := config.New("./environments")
	logger.InitLogger(config.GetStringSetting("logs.level"))

	productStore := new(products.ProductStore)
	productService := products.NewProductService(productStore)
	productController := products.NewProductController(productService)

	server := router.SetupRouter()
	server = router.GetProducts(server, productController.GetProducts)
	server = router.AddProduct(server, productController.Add)
	server.Run(":8080")
	slog.Debug("Server started")
}
