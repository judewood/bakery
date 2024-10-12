package products

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ProductController struct {
	productService ProductServer
}

func NewProductController(productServer ProductServer) *ProductController {
	return &ProductController{
		productService: productServer,
	}
}

func Ping(c *gin.Context) {
	c.String(200, "pong")
}

func (p *ProductController) GetProducts(c *gin.Context) {
	v, err := p.productService.GetAvailableProducts()
	if err != nil {
		fmt.Printf("Failed to get available products %v", err)
		c.IndentedJSON(http.StatusBadRequest, nil)
		return
	}
	if len(v) == 0 {
		c.IndentedJSON(http.StatusNoContent, v)
		return
	}
	c.IndentedJSON(http.StatusOK, v)
}

func (p *ProductController) Add(ctx *gin.Context) {
	product := Product{}
	err := ctx.BindJSON(&product)
	if err != nil {
		fmt.Printf("\ngin says %v", err)
		return
	}
	added, err := p.productService.Add(product)
	if err != nil {
		if err.Error() == MissingRequired {
			ctx.JSON(http.StatusBadRequest, "")
			return
		}
		ctx.JSON(http.StatusInternalServerError, "")
		return
	}
	ctx.JSON(http.StatusCreated, added)
}
