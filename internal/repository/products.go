package repository

import (
	"fmt"
	"github.com/yael-castro/products-api/internal/model"
	error2 "github.com/yael-castro/products-api/internal/model/error"
	"gorm.io/gorm"
)

// "implement" constraint for ProductStore
var _ StorageManager[model.SKU, model.Product] = ProductStore{}

// ProductStore has the common methods to manage the storage of model.Product
type ProductStore struct {
	*gorm.DB
}

// Create inserts into the database a new record using the *model.Product received as parameter
func (p ProductStore) Create(product *model.Product) error {
	return p.DB.Create(product).Error
}

// Obtain finds the record for model.Product identified by model.SKU
func (p ProductStore) Obtain(sku model.SKU) (product model.Product, err error) {
	db := p.DB.Where("sku = ?", sku).Find(&product)

	if err = db.Error; err != nil {
		return
	}

	if db.RowsAffected < 1 {
		err = error2.NotFound(fmt.Sprintf(`product identified by sku '%s' does not exist`, sku))
	}

	return
}

// Update using the instance of model.SKU and model.Product updates the record into database identified by model.SKU
func (p ProductStore) Update(sku model.SKU, product model.Product) error {
	return p.DB.Where("sku = ?", sku).Updates(product).Error
}

// Delete removes the record identified by model.SKU from the database
func (p ProductStore) Delete(sku model.SKU) error {
	rowsAffected := p.DB.Delete(model.Product{SKU: sku}).RowsAffected

	if rowsAffected < 1 {
		return error2.NotFound(fmt.Sprintf(`product identified by sku '%s' does not exist`, sku))
	}

	return nil
}

// List lists all model.Products from the database
func (p ProductStore) List() (products model.Products, err error) {
	products = model.Products{}
	err = p.DB.Find(&products).Error
	return
}
