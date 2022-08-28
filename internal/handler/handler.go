// Package handler is the presentation layer, contains everything needed to trigger the behavior of the program, both synchronous and asynchronous
package handler

import (
	"github.com/gin-gonic/gin"
)

// Handler defines the main handler that contains all *gin.HandlerFunc
type Handler interface {
	ProductManager
	HealthCheck(*gin.Context)
}

// _ "implements" constraint for GinHandlers
var _ Handler = GinHandlers{}

// GinHandlers is the collection of all *gin.HandlerFunc used to initialize the *gin.Engine
//
// In resume is the configuration to initialize the *gin.Engine
type GinHandlers struct {
	ProductManager
	healthCheck gin.HandlerFunc
}

// SetHealthCheck sets the gin.HandlerFunc that is used as health check handler
func (g *GinHandlers) SetHealthCheck(healthCheck gin.HandlerFunc) {
	g.healthCheck = healthCheck
}

// HealthCheck is the default *gin.HandlerFunc to know and monitoring the server status
func (g GinHandlers) HealthCheck(c *gin.Context) {
	g.healthCheck(c)
}

// HealthCheck is the default *gin.HandlerFunc used to know the health server and monitoring the server status
func HealthCheck(c *gin.Context) {
	c.Writer.Write(nil)
}

// NewGinEngine using an instance of Handler initializes the *gin.Engine
func NewGinEngine(h Handler) *gin.Engine {
	engine := gin.Default()

	engine.GET("/", h.HealthCheck)

	engine.POST("/products-api/v1/products/", h.CreateProduct)

	engine.GET("/products-api/v1/products/", h.ObtainProducts)
	engine.GET("/products-api/v1/products/:id", h.ObtainProduct)

	engine.PUT("/products-api/v1/products/:id", h.UpdateProduct)

	engine.DELETE("/products-api/v1/products/", h.DeleteProduct)

	return engine
}
