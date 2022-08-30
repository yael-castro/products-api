package business

import (
	"github.com/yael-castro/products-api/internal/model"
	error2 "github.com/yael-castro/products-api/internal/model/error"
	"github.com/yael-castro/products-api/internal/repository"
)

// _ "implements" constraint for ProductStore
var _ ProductManager = ProductStore{}

// ProductStore manage the products management
type ProductStore struct {
	repository.StorageManager[model.SKU, model.Product]
}

// validateProductData validates if the model.Product received is valid, if the model.Product is not valid returns an error
func (s ProductStore) validateProductData(product model.Product) error {
	if err := product.SKU.IsValid(); err != nil {
		return error2.Validation(err.Error())
	}

	switch {
	case product.Name == "":
		return error2.Validation("product name must not be blank")
	case len(product.Name) < 3:
		return error2.Validation("product name is too short")
	case len(product.Name) > 50:
		return error2.Validation("product name is too large")

	case product.Brand == "":
		return error2.Validation("product brand must not be blank")
	case len(product.Brand) < 3:
		return error2.Validation("product brand is too short")
	case len(product.Brand) > 50:
		return error2.Validation("product brand is too large")

	case product.Size != nil && *product.Size == "":
		return error2.Validation("product size must not be blank")

	case product.Price < 1.0 || product.Price > 99_999_999.00:
		return error2.Validation("invalid product price")

	case product.PrincipalImage == nil:
		return error2.Validation("principal image for product is required")
	case product.PrincipalImage.URL == nil:
		return error2.Validation("principal image for product is required")
	}

	return nil
}

// CreateProduct validates the model.Product and if it is valid, a record is created in the storage
func (s ProductStore) CreateProduct(product *model.Product) error {
	err := s.validateProductData(*product)
	if err != nil {
		return err
	}

	return s.Create(product)
}

// ObtainProduct if the model.SKU received as parameter is valid, search into storage a record identifier by the model.SKU
//
// The model.SKU is validated before to search the model.Product into storage to avoid
func (s ProductStore) ObtainProduct(sku model.SKU) (model.Product, error) {
	if err := sku.IsValid(); err != nil {
		return model.Product{}, error2.Validation(err.Error())
	}

	return s.Obtain(sku)
}

// UpdateProduct updates the record of model.Product identified by the model.SKU
//
// The model.SKU is validated before updating the record to avoid unnecessary and wasted requests to storage
func (s ProductStore) UpdateProduct(product model.Product) error {
	err := s.validateProductData(product)
	if err != nil {
		return err
	}

	return s.Update(product.SKU, product)
}

// DeleteProduct deletes the record of model.Product identified by the model.SKU received
//
// The model.SKU is validated before de-registration to avoid unnecessary and wasted storage requests
func (s ProductStore) DeleteProduct(sku model.SKU) error {
	if err := sku.IsValid(); err != nil {
		return error2.Validation(err.Error())
	}

	return s.Delete(sku)
}

// ListProducts returns all records of model.Product from the storage
func (s ProductStore) ListProducts() (model.Products, error) {
	return s.List()
}
