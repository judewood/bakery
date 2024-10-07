package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/judewood/bakery/models"
	"github.com/judewood/bakery/internal/products"
)

type ProductController interface {
	GetProducts() []models.Product
	Add(ctx *gin.Context) (models.Product, error)
}

type ProductCon struct {
	productServer products.ProductServer
}

func NewProductController(productServer products.ProductServer) *ProductCon {
	return &ProductCon{
		productServer: productServer,
	}
}

func (p *ProductCon) GetProducts() []models.Product {
	v, _ := p.productServer.GetAvailableProducts()
	return v
}

func (p *ProductCon) Add(ctx *gin.Context) (models.Product, error) {
	var product models.Product
	ctx.BindJSON(&product)
	return product, nil
}
