package router

import (
	"log/slog"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/ping", pingServer)
	return r
}

func GetProducts(r *gin.Engine, handler gin.HandlerFunc) *gin.Engine {
	r.GET("/products", handler)
	slog.Info("Added get all route for products")
	return r
}

func GetProduct(r *gin.Engine, handler gin.HandlerFunc) *gin.Engine {
	r.GET("/product/:id", handler)
	slog.Info("Added get route for product id")
	return r
}

func AddProduct(r *gin.Engine, handler gin.HandlerFunc) *gin.Engine {
	r.POST("/product", handler)
	slog.Info("Added POST route for products")
	return r
}

func UpdateProduct(r *gin.Engine, handler gin.HandlerFunc) *gin.Engine {
	r.PUT("/product", handler)
	slog.Info("Added PUT route for products")
	return r
}

func DeleteProduct(r *gin.Engine, handler gin.HandlerFunc) *gin.Engine {
	r.DELETE("/product/:id", handler)
	slog.Info("Added delete route for product id")
	return r
}

func pingServer(c *gin.Context) {
	slog.Info("Added GET ping route")
	c.String(200, "pong")
}
