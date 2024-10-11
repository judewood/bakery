package router

import (
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})
	return r
}

func AddRouteGetProducts(r *gin.Engine, handler gin.HandlerFunc) *gin.Engine {
	r.GET("/products", handler)
	return r
}
