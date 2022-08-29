package business

import (
	"errors"
	"github.com/yael-castro/agrak/internal/model"
	error2 "github.com/yael-castro/agrak/internal/model/error"
	"github.com/yael-castro/agrak/internal/repository"
	"net/url"
	"reflect"
	"strconv"
	"testing"
)

func TestProductStore_ObtainProduct(t *testing.T) {
	tdt := []struct {
		sku             model.SKU
		expectedProduct model.Product
		expectedErr     error
	}{
		{
			sku:         "invalid",
			expectedErr: error2.Validation("missing prefix 'FAL-'"),
		},
		{
			sku:         "FAL-0",
			expectedErr: error2.Validation("invalid suffix '0'"),
		},
		{
			sku:         "FAL-100000000000",
			expectedErr: error2.Validation("invalid suffix '100000000000'"),
		},
		{
			sku:             model.SKU("FAL-" + strconv.Itoa(99_999_999)),
			expectedProduct: model.Product{SKU: "FAL-99999999"},
		},
	}

	store := ProductStore{
		StorageManager: &repository.MockStorage[model.SKU, model.Product]{
			"FAL-99999999": model.Product{SKU: "FAL-99999999"},
		},
	}

	for i, v := range tdt {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			product, err := store.ObtainProduct(v.sku)

			if !errors.Is(err, v.expectedErr) {
				t.Fatalf("expected error '%v' unexpected error '%v'", v.expectedErr, err)
			}

			if err != nil {
				t.Skip(err)
			}

			if !reflect.DeepEqual(v.expectedProduct, product) {
				t.Fatalf("expected product '%v' unexpected product '%v'", v.expectedProduct, product)
			}

			t.Log(product)
		})
	}
}

func TestProductStore_DeleteProduct(t *testing.T) {
	tdt := []struct {
		sku         model.SKU
		expectedErr error
	}{
		{
			sku:         "invalid",
			expectedErr: error2.Validation("missing prefix 'FAL-'"),
		},
		{
			sku:         "FAL-0",
			expectedErr: error2.Validation("invalid suffix '0'"),
		},
		{
			sku:         "FAL-100000000000",
			expectedErr: error2.Validation("invalid suffix '100000000000'"),
		},
		{
			sku: model.SKU("FAL-" + strconv.Itoa(99_999_999)),
		},
	}

	store := ProductStore{
		StorageManager: &repository.MockStorage[model.SKU, model.Product]{},
	}

	for i, v := range tdt {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			err := store.DeleteProduct(v.sku)

			if !errors.Is(err, v.expectedErr) {
				t.Fatalf("expected error '%v' unexpected error '%v'", v.expectedErr, err)
			}

			if err != nil {
				t.Skip(err)
			}

			t.Log("SUCCESS")
		})
	}
}

func TestProductStore_CreateProduct(t *testing.T) {
	tdt := []struct {
		product     model.Product
		expectedErr error
	}{
		{
			product: model.Product{
				SKU:            "FAL-1234567",
				Price:          10,
				Brand:          "Nike",
				Name:           "Shoes",
				PrincipalImage: &model.URL{URL: &url.URL{}},
			},
		},
		{
			product: model.Product{
				SKU: "FAL-1",
			},
			expectedErr: error2.Validation("invalid suffix '1'"),
		},
		{
			product: model.Product{
				SKU: "FAL-1234567",
			},
			expectedErr: error2.Validation("product name must not be blank"),
		},
		{
			product: model.Product{
				SKU:  "FAL-1234567",
				Name: "A",
			},
			expectedErr: error2.Validation("product name is too short"),
		},
		{
			product: model.Product{
				SKU:  "FAL-1234567",
				Name: string(make([]byte, 51)),
			},
			expectedErr: error2.Validation("product name is too large"),
		},
		{
			product: model.Product{
				SKU:  "FAL-1234567",
				Name: "AAA",
			},
			expectedErr: error2.Validation("product brand must not be blank"),
		},
		{
			product: model.Product{
				SKU:   "FAL-1234567",
				Name:  "AAA",
				Brand: "AA",
			},
			expectedErr: error2.Validation("product brand is too short"),
		},
		{
			product: model.Product{
				SKU:   "FAL-1234567",
				Name:  "AAA",
				Brand: string(make([]byte, 51)),
			},
			expectedErr: error2.Validation("product brand is too large"),
		},
	}

	store := ProductStore{
		StorageManager: &repository.MockStorage[model.SKU, model.Product]{},
	}

	for i, v := range tdt {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			err := store.CreateProduct(&v.product)
			if !errors.Is(err, v.expectedErr) {
				t.Fatalf("expected error '%v' unexpected error '%v'", v.expectedErr, err)
			}

			if err != nil {
				t.Skip(err)
			}

			t.Log(v.product)
		})
	}
}

func TestProductStore_UpdateProduct(t *testing.T) {
	tdt := []struct {
		product     model.Product
		expectedErr error
	}{
		{
			product: model.Product{
				SKU:            "FAL-1234567",
				Price:          10,
				Brand:          "Nike",
				Name:           "Shoes",
				PrincipalImage: &model.URL{URL: &url.URL{}},
			},
		},
		{
			product: model.Product{
				SKU: "FAL-1",
			},
			expectedErr: error2.Validation("invalid suffix '1'"),
		},
		{
			product: model.Product{
				SKU: "FAL-1234567",
			},
			expectedErr: error2.Validation("product name must not be blank"),
		},
		{
			product: model.Product{
				SKU:  "FAL-1234567",
				Name: "A",
			},
			expectedErr: error2.Validation("product name is too short"),
		},
		{
			product: model.Product{
				SKU:  "FAL-1234567",
				Name: string(make([]byte, 51)),
			},
			expectedErr: error2.Validation("product name is too large"),
		},
		{
			product: model.Product{
				SKU:  "FAL-1234567",
				Name: "AAA",
			},
			expectedErr: error2.Validation("product brand must not be blank"),
		},
		{
			product: model.Product{
				SKU:   "FAL-1234567",
				Name:  "AAA",
				Brand: "AA",
			},
			expectedErr: error2.Validation("product brand is too short"),
		},
		{
			product: model.Product{
				SKU:   "FAL-1234567",
				Name:  "AAA",
				Brand: string(make([]byte, 51)),
			},
			expectedErr: error2.Validation("product brand is too large"),
		},
	}

	store := ProductStore{
		StorageManager: &repository.MockStorage[model.SKU, model.Product]{},
	}

	for i, v := range tdt {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			err := store.UpdateProduct(v.product)
			if !errors.Is(err, v.expectedErr) {
				t.Fatalf("expected error '%v' unexpected error '%v'", v.expectedErr, err)
			}

			if err != nil {
				t.Skip(err)
			}

			t.Log(v.product)
		})
	}
}
