package products

import (
	"fmt"
	"log/slog"
	"net/http"
	"strings"

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
	v, err := p.productService.GetAll()
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

func (p *ProductController) Get(ctx *gin.Context) {
	slog.Info(fmt.Sprintln("context params", ctx.Params))
	id := ctx.Param("id")
	slog.Debug("Get by id request", "id", strings.ToLower(id))

	product, err := p.productService.Get(id)

	if err != nil {
		if err.Error() == MissingID {
			ctx.JSON(http.StatusBadRequest, "")
			return
		}
		if err.Error() == fmt.Sprintf(NotFound, id) {
			ctx.JSON(http.StatusNoContent, "")
			return
		}
		ctx.JSON(http.StatusInternalServerError, "")
		return
	}
	ctx.IndentedJSON(http.StatusOK, product)
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

func (p *ProductController) Update(ctx *gin.Context) {
	product := Product{}
	err := ctx.BindJSON(&product)
	if err != nil {
		fmt.Printf("\ngin says %v", err)
		return
	}
	added, err := p.productService.Update(product)
	if err != nil {
		if err.Error() == MissingRequired {
			ctx.JSON(http.StatusBadRequest, "")
			return
		}
		if err.Error() == fmt.Sprintf(NotFound, product.Name) {
			ctx.JSON(http.StatusNoContent, "")
			return
		}
		ctx.JSON(http.StatusInternalServerError, "")
		return
	}
	ctx.JSON(http.StatusOK, added)
}

func (p *ProductController) Delete(ctx *gin.Context) {
	slog.Info(fmt.Sprintln("context params", ctx.Params))
	id := ctx.Param("id")
	slog.Debug("Delete by id request", "id", strings.ToLower(id))

	product, err := p.productService.Delete(id)

	if err != nil {
		if err.Error() == MissingID {
			ctx.JSON(http.StatusBadRequest, "")
			return
		}
		if err.Error() == fmt.Sprintf(NotFound, id) {
			ctx.JSON(http.StatusNoContent, "")
			return
		}
		ctx.JSON(http.StatusInternalServerError, "")
		return
	}
	ctx.IndentedJSON(http.StatusOK, product)
}
