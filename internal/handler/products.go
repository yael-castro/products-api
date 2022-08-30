package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/yael-castro/products-api/internal/business"
	"github.com/yael-castro/products-api/internal/model"
	"net/http"
)

// _ "implements" constraint for ProductStore
var _ ProductManager = ProductStore{}

// ProductStore contains the group of gin.HandlerFunc for handle requests related to management of product storage
type ProductStore struct {
	business.ProductManager
}

// CreateProduct gin.HandlerFunc to handle http requests made to add a product into the storage
func (p ProductStore) CreateProduct(c *gin.Context) {
	product := model.Product{}

	c.Header("Content-Type", "application/json")
	err := c.BindJSON(&product)
	if err != nil {
		handleError(c, err)
		return
	}

	err = p.ProductManager.CreateProduct(&product)
	if err != nil {
		handleError(c, err)
		return
	}

	c.JSON(http.StatusCreated, product)
}

// ObtainProduct gin.HandlerFunc to handle http requests made to obtain a product from the storage
func (p ProductStore) ObtainProduct(c *gin.Context) {
	sku := c.Param("id")

	product, err := p.ProductManager.ObtainProduct(model.SKU(sku))
	if err != nil {
		handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, product)
}

// UpdateProduct gin.HandlerFunc to handle http requests made to update existing product in the storage
func (p ProductStore) UpdateProduct(c *gin.Context) {
	product := model.Product{}

	c.Header("Content-Type", "application/json")
	err := c.BindJSON(&product)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"message": err.Error()})
		return
	}

	err = p.ProductManager.UpdateProduct(product)
	if err != nil {
		handleError(c, err)
		return
	}

	c.JSON(http.StatusCreated, product)
}

// DeleteProduct gin.HandlerFunc to handle http requests made to remove a product from the storage
func (p ProductStore) DeleteProduct(c *gin.Context) {
	sku := c.Param("id")

	err := p.ProductManager.DeleteProduct(model.SKU(sku))
	if err != nil {
		handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "ok"})
}

// ObtainProducts gin.HandlerFunc to handle http requests made to remove a product from the storage
// TODO: pagination
func (p ProductStore) ObtainProducts(c *gin.Context) {
	products, err := p.ProductManager.ListProducts()
	if err != nil {
		handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, products)
}
