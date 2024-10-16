package products

import (
	"fmt"
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/judewood/bakery/errorutils"
)

type ProductController struct {
	productService ProductServer
}

func NewProductController(productServer ProductServer) *ProductController {
	return &ProductController{
		productService: productServer,
	}
}

func (p *ProductController) GetProducts(c *gin.Context) {
	slog.Debug("Get all products request")
	v, err := p.productService.GetAll()
	if err != nil {
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
	id := ctx.Param("id")
	slog.Debug("Get product by id request", "id", id)

	product, err := p.productService.Get(id)

	if err != nil {
		slog.Debug("fff"+ err.Error())
		if err.Error() == errorutils.MissingID {
			slog.Debug("miss"+ err.Error())
			ctx.JSON(http.StatusBadRequest, "")
			return
		}
		if err.Error() == fmt.Sprintf(errorutils.NotFound, id) {
			slog.Debug("notf" + err.Error())
			ctx.JSON(http.StatusNoContent, "")
			return
		}
		slog.Warn("oops" + err.Error())
		ctx.JSON(http.StatusInternalServerError, "")
		return
	}
	ctx.IndentedJSON(http.StatusOK, product)
}

func (p *ProductController) Add(ctx *gin.Context) {
	product := Product{}
	err := ctx.BindJSON(&product)
	if err != nil {
		slog.Warn("invalid product add", "error", err)
		ctx.JSON(http.StatusBadRequest, "")
		return
	}
	slog.Debug("Add product request", "product", product)
	added, err := p.productService.Add(product)
	if err != nil {
		if err.Error() == errorutils.MissingRequired {
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
		slog.Warn("invalid product update", "error", err)
		ctx.JSON(http.StatusBadRequest, "")
		return
	}
	slog.Debug("Update product request", "product", product)
	added, err := p.productService.Update(product)
	if err != nil {
		if err.Error() == errorutils.MissingRequired {
			ctx.JSON(http.StatusBadRequest, "")
			return
		}
		if err.Error() == errorutils.NotFoundError(product.Name).Error() {
			ctx.JSON(http.StatusNoContent, "")
			return
		}
		ctx.JSON(http.StatusInternalServerError, "")
		return
	}
	ctx.JSON(http.StatusOK, added)
}

func (p *ProductController) Delete(ctx *gin.Context) {
	id := ctx.Param("id")
	slog.Debug("Delete by id request", "id", id)

	product, err := p.productService.Delete(id)

	if err != nil {
		if err.Error() == errorutils.MissingID {
			ctx.JSON(http.StatusBadRequest, "")
			return
		}
		if err.Error() == errorutils.NotFoundError(id).Error() {
			ctx.JSON(http.StatusNoContent, "")
			return
		}
		ctx.JSON(http.StatusInternalServerError, "")
		return
	}
	ctx.IndentedJSON(http.StatusOK, product)
}
