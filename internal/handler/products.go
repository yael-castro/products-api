package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/yael-castro/agrak/internal/business"
	"github.com/yael-castro/agrak/internal/model"
	"net/http"
)

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

// _ "implements" constraint for ProductStore
var _ ProductManager = ProductStore{}

type ProductStore struct {
	business.ProductManager
}

func (p ProductStore) CreateProduct(c *gin.Context) {
	product := model.Product{}

	err := c.Bind(&product)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	err = p.ProductManager.CreateProduct(&product)
	if err != nil {
		handleError(c, err)
		return
	}

	c.JSON(http.StatusCreated, product)
}

func (p ProductStore) ObtainProduct(c *gin.Context) {
	sku := c.Param("id")

	product, err := p.ProductManager.ObtainProduct(model.SKU(sku))
	if err != nil {
		handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, product)
}

func (p ProductStore) UpdateProduct(c *gin.Context) {
	product := model.Product{}

	err := c.Bind(&product)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	err = p.ProductManager.UpdateProduct(product)
	if err != nil {
		handleError(c, err)
		return
	}

	c.JSON(http.StatusCreated, product)
}

func (p ProductStore) DeleteProduct(c *gin.Context) {
	sku := c.Param("id")

	err := p.ProductManager.DeleteProduct(model.SKU(sku))
	if err != nil {
		handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "ok"})
}

func (p ProductStore) ObtainProducts(c *gin.Context) {
	products, err := p.ProductManager.ListProducts()
	if err != nil {
		handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, products)
}
