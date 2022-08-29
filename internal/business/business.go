// Package business is the business layer, contains everything related to business logic like business rules and business process
package business

import (
	"github.com/yael-castro/agrak/internal/model"
)

// ProductManager defines the common actions for product manager
type ProductManager interface {
	// CreateProduct saves a new model.Product into the storage
	CreateProduct(product *model.Product) error
	// ObtainProduct returns the record identified by model.SKU from the store
	ObtainProduct(sku model.SKU) (model.Product, error)
	// UpdateProduct updates the record for model.Product identified by model.SKU
	UpdateProduct(product model.Product) error
	// DeleteProduct removes a record of model.Product identified by model.SKU from the store
	DeleteProduct(sku model.SKU) error
	// ListProducts returns all records of model.Product into the store
	ListProducts() (model.Products, error)
}
