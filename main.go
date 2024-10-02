package main

import (
	"fmt"

	"github.com/judewood/bakery/models"
	"github.com/judewood/bakery/service"
	"github.com/judewood/bakery/store"
)

func main() {
	productStore := store.NewProductStore()
	productService := service.NewProductService(*productStore)
	g := productService.GetAvailableProducts()
	displayAvailableProducts(g)
}

func displayAvailableProducts(products []models.Product) {
	fmt.Print("Creating random order for:")
	for _, v := range products {
		fmt.Printf("\n %v ", v.Name)
	}
}
