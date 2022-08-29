// Package handler is the presentation layer, contains everything needed to trigger the behavior of the program, both synchronous and asynchronous
package handler

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	error2 "github.com/yael-castro/agrak/internal/model/error"
	"log"
	"net/http"
)

// Handler defines the main handler that contains all *gin.HandlerFunc
type Handler interface {
	ProductManager
}

// ProductManager defines the *gin.HandlerFunc group to manage the http requests related to product management
type ProductManager interface {
	// CreateProduct handle http requests to add a new product to the storage
	CreateProduct(*gin.Context)
	// ObtainProduct handle http requests to search a product from the storage
	ObtainProduct(*gin.Context)
	// UpdateProduct handle http requests to update a product from the storage
	UpdateProduct(*gin.Context)
	// DeleteProduct handle http requests to remove a product from the storage
	DeleteProduct(*gin.Context)
	// ObtainProducts handle http requests to list products
	ObtainProducts(*gin.Context)
}

// _ "implements" constraint for GinHandlers
var _ Handler = GinHandlers{}

// GinHandlers is the collection of all *gin.HandlerFunc used to initialize the *gin.Engine
//
// In resume is the configuration to initialize the *gin.Engine
type GinHandlers struct {
	ProductManager
}

// HealthCheck is the default *gin.HandlerFunc used to know the health server and monitoring the server status
func HealthCheck(c *gin.Context) {
	_, _ = c.Writer.Write(nil)
}

// NotFound is the default *gin.HandlerFunc used to handle http requests made to non exist paths
func NotFound(c *gin.Context) {
	c.JSON(http.StatusNotFound, gin.H{"message": fmt.Sprintf(`path '%s' does not exist`, c.Request.URL.Path)})
}

// NewGinEngine using an instance of Handler initializes the *gin.Engine
func NewGinEngine(h Handler) *gin.Engine {
	engine := gin.Default()

	// Default handlers
	engine.NoRoute(NotFound)
	engine.GET("/", HealthCheck)

	engine.POST("/v1/products/", h.CreateProduct)

	engine.GET("/v1/products/", h.ObtainProducts)
	engine.GET("/v1/products/:id", h.ObtainProduct)

	engine.PUT("/v1/products/:id", h.UpdateProduct)

	engine.DELETE("/v1/products/", h.DeleteProduct)

	return engine
}

// handleError handles errors and related it to http response codes
func handleError(c *gin.Context, err error) {
	switch err.(type) {
	case error2.Validation:
		c.JSON(http.StatusBadRequest, gin.H{"message": err})

	case error2.NotFound:
		c.JSON(http.StatusNotFound, gin.H{"message": err})

	case *json.MarshalerError:
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})

	case *json.SyntaxError:
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})

	case *error2.PG:
		switch err.(*error2.PG).Code {
		case "23505":
			c.JSON(http.StatusConflict, gin.H{"message": "duplicated record"})
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"message": "an unexpected error related to storage occurred"})
		}

	default:
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
	}

	log.Printf("ERROR: %T\n", err)
}
