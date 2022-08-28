package handler

import (
	"github.com/gin-gonic/gin"
)

// ProductManager defines the *gin.HandlerFunc group to manage the http requests related to product management
type ProductManager interface {
	// CreateProduct handle http requests to add a new product to the storage
	CreateProduct(*gin.Context)
	// ObtainProduct handle http requests to search a product from the storage
	ObtainProduct(*gin.Context)
	// DeleteProduct handle http requests to remove a product from the storage
	DeleteProduct(*gin.Context)
	// ObtainProducts handle http requests to list products
	ObtainProducts(*gin.Context)
}

// _ "implements" constraint for ProductStore
var _ ProductManager = ProductStore{}

type ProductStore struct {
}

func (p ProductStore) CreateProduct(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (p ProductStore) ObtainProduct(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (p ProductStore) DeleteProduct(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (p ProductStore) ObtainProducts(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}
