package model

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type (
	Product struct {
		// SKU internal stock-keeping unit. It is the candidate identifier of a product
		SKU SKU `json:"sku" gorm:"type:varchar;primaryKey"`
		// Name short description of the product
		Name string `json:"name" gorm:"type:varchar;not null"`
		// Brand name of the brand
		Brand string `json:"brand" gorm:"type:varchar;not null"`
		// Size product size
		Size *string `json:"size" gorm:"type:varchar;not null"`
		// Price sell price
		Price float64 `json:"price" gorm:"type:decimal;not null"`
		// PrincipalImage URL of the principal image of the product, which is used in catalogs
		// and is the first image that is showed to customers when access product detail page
		PrincipalImage *URL `json:"principalImage" gorm:"varchar;not null"`
		// OtherImages list of images of the product
		OtherImages URLs `json:"otherImages" gorm:"[]varchar;not null"`
	}

	// Products alias for []Product
	Products = []Product
)

// SKU stock-keeping unit
type SKU string

// IsValid check if the SKU are valid, if it does not valid returns an error
func (s SKU) IsValid() error {
	if s == "" || !strings.HasPrefix(string(s), "FAL-") {
		return errors.New("missing prefix 'FAL-'")
	}

	suffix, err := strconv.ParseInt(string(s[4:]), 10, 64)
	if err != nil {
		return errors.New("suffix is not a number")
	}

	if suffix < 1_000_000 {
		return fmt.Errorf("invalid suffix '%d'", suffix)
	}

	if suffix > 99_999_999 {
		return fmt.Errorf("invalid suffix '%d'", suffix)
	}

	return nil
}
