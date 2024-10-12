package products

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ProductController struct {
	productServer ProductServer
}

func NewProductController(productServer ProductServer) *ProductController {
	return &ProductController{
		productServer: productServer,
	}
}

func Ping(c *gin.Context) {
	c.String(200, "pong")
}

func (p *ProductController) GetProducts(c *gin.Context) {
	v, err := p.productServer.GetAvailableProducts()
	if err != nil {
		fmt.Printf("Failed to get available products %v", err)
		c.IndentedJSON(http.StatusBadRequest, nil)
		return
	}
	c.IndentedJSON(http.StatusOK, v)
}

func (p *ProductController) Add(ctx *gin.Context){
	product := Product{}
	err := ctx.BindJSON(&product)
	if err != nil {
		fmt.Errorf("\ngin says %v", err)
		return
	}
	added, err := p.productServer.Add(product)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, nil)
		return	
	}
	ctx.JSON(http.StatusCreated, added)	
}
