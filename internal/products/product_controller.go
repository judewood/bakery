package products

import (
	"github.com/gin-gonic/gin"
)

type ProductController interface {
	GetProducts() []Product
	Add(ctx *gin.Context) (Product, error)
}

type ProductCon struct {
	productServer ProductServer
}

func NewProductController(productServer ProductServer) *ProductCon {
	return &ProductCon{
		productServer: productServer,
	}
}

func (p *ProductCon) GetProducts() []Product {
	v, _ := p.productServer.GetAvailableProducts()
	return v
}

func (p *ProductCon) Add(ctx *gin.Context) (Product, error) {
	var product Product
	ctx.BindJSON(&product)
	return product, nil
}
