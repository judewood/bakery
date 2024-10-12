package router

import (
	"log/slog"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/ping", Ping)
	return r
}

func GetProducts(r *gin.Engine, handler gin.HandlerFunc) *gin.Engine {
	r.GET("/products", handler)
	slog.Info("Added GET route for products")
	return r
}

func AddProduct(r *gin.Engine, handler gin.HandlerFunc) *gin.Engine {
	r.POST("/products", handler)
	slog.Info("Added POST route for products")
	return r
}

func Ping(c *gin.Context) {
	slog.Info("Added GET ping route")
	c.String(200, "pong")
}
