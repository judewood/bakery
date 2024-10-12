package router

import (
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/ping", Ping)
	return r
}

func GetProducts(r *gin.Engine, handler gin.HandlerFunc) *gin.Engine {
	r.GET("/products", handler)
	return r
}

func AddProduct(r *gin.Engine, handler gin.HandlerFunc) *gin.Engine {
	r.POST("/products", handler)
	return r
}

func Ping(c *gin.Context) {
	c.String(200, "pong")
}
