package repository

import (
	"errors"
	"flag"
	"github.com/yael-castro/agrak/internal/model"
	error2 "github.com/yael-castro/agrak/internal/model/error"
	"log"
	"net/url"
	"os"
	"reflect"
	"strconv"
	"testing"
)

// verbose cli flag for "V", indicates whether to show additional logs, such as SQL logs
var verbose = flag.Bool("V", false, "")

func TestProductStore_Create(t *testing.T) {
	tdt := []struct {
		product     model.Product
		expectedErr error
	}{
		{
			product: model.Product{
				SKU:   "1234",
				Name:  "...",
				Brand: "Nike",
				Size:  &[]string{"M"}[0],
				Price: 1_000,
				PrincipalImage: func() *model.URL {
					u, _ := url.Parse("https://example.com")
					return &model.URL{URL: u}
				}(),
				OtherImages: model.URLs{
					func() model.URL {
						u, _ := url.Parse("https://example.com")
						return model.URL{URL: u}
					}(),
				},
			},
		},
	}

	// gormDSN is the Data Source Name for GORM
	gormDSN := os.Getenv("GORM_DSN")
	if gormDSN == "" {
		t.Fatal(`missing environment variable "GORM_DSN"`)
	}

	db, err := NewGormDB(gormDSN)
	if err != nil {
		t.Fatal(err)
	}

	storage := ProductStore{DB: db}
	if *verbose {
		storage.DB = db.Debug()
	}

	for i, v := range tdt {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			t.Cleanup(func() {
				_ = storage.Delete(v.product.SKU)
			})

			err := storage.Create(&v.product)
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

func TestProductStore_Obtain(t *testing.T) {
	tdt := []struct {
		product      model.Product
		expectedErr  error
		skipCreation bool
	}{
		{
			product: model.Product{
				SKU:   "1234",
				Name:  "...",
				Brand: "Nike",
				Size:  &[]string{"M"}[0],
				Price: 1_000,
				PrincipalImage: func() *model.URL {
					u, _ := url.Parse("https://example.com")
					return &model.URL{URL: u}
				}(),
				OtherImages: model.URLs{
					func() model.URL {
						u, _ := url.Parse("https://example1.com")
						return model.URL{URL: u}
					}(),
				},
			},
		},
		{
			product:      model.Product{SKU: "12345"},
			skipCreation: true,
			expectedErr:  error2.NotFound("product identified by sku '12345' does not exist"),
		},
	}

	// gormDSN is the Data Source Name for GORM
	gormDSN := os.Getenv("GORM_DSN")
	if gormDSN == "" {
		t.Fatal(`missing environment variable "GORM_DSN"`)
	}

	db, err := NewGormDB(gormDSN)
	if err != nil {
		t.Fatal(err)
	}

	storage := ProductStore{DB: db}
	if *verbose {
		storage.DB = db.Debug()
	}

	for i, v := range tdt {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			if !v.skipCreation {
				_ = storage.Create(&v.product)
				t.Cleanup(func() {
					log.Println(storage.Delete(v.product.SKU))
				})
			}

			product, err := storage.Obtain(v.product.SKU)
			if !errors.Is(err, v.expectedErr) {
				t.Fatalf("expected error '%v' unexpected error '%v'", v.expectedErr, err)
			}

			if err != nil {
				t.Skip(err)
			}

			if !reflect.DeepEqual(v.product, product) {
				t.Fatalf("expected product '%v' unexpected product '%v'", v.product, product)
			}

			t.Logf("%+v", v.product)
		})
	}
}

func TestProductStore_Delete(t *testing.T) {
	tdt := []struct {
		product      model.Product
		expectedErr  error
		skipCreation bool
	}{
		{
			product: model.Product{
				SKU:   "1234",
				Name:  "...",
				Brand: "Nike",
				Size:  &[]string{"M"}[0],
				Price: 1_000,
				PrincipalImage: func() *model.URL {
					u, _ := url.Parse("https://example.com")
					return &model.URL{URL: u}
				}(),
				OtherImages: model.URLs{
					func() model.URL {
						u, _ := url.Parse("https://example.com")
						return model.URL{URL: u}
					}(),
				},
			},
		},
		{
			product:      model.Product{SKU: "12345"},
			skipCreation: true,
			expectedErr:  error2.NotFound(`product identified by sku '12345' does not exist`),
		},
	}

	// gormDSN is the Data Source Name for GORM
	gormDSN := os.Getenv("GORM_DSN")
	if gormDSN == "" {
		t.Fatal(`missing environment variable "GORM_DSN"`)
	}

	db, err := NewGormDB(gormDSN)
	if err != nil {
		t.Fatal(err)
	}

	storage := ProductStore{DB: db}
	if *verbose {
		storage.DB = db.Debug()
	}

	for i, v := range tdt {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			if !v.skipCreation {
				_ = storage.Create(&v.product)
			}

			err := storage.Delete(v.product.SKU)
			if !errors.Is(err, v.expectedErr) {
				t.Fatalf("expected error '%v' unexpected error '%v'", v.expectedErr, err)
			}

			if err != nil {
				t.Skip(err)
			}

			t.Logf("%+v", v.product)
		})
	}
}

func TestProductStore_Update(t *testing.T) {
	tdt := []struct {
		product      model.Product
		expectedErr  error
		skipCreation bool
	}{
		{
			product: model.Product{
				SKU:   "1234",
				Name:  "...",
				Brand: "Nike",
				Size:  &[]string{"M"}[0],
				Price: 1_000,
				PrincipalImage: func() *model.URL {
					u, _ := url.Parse("https://example.com")
					return &model.URL{URL: u}
				}(),
				OtherImages: model.URLs{
					func() model.URL {
						u, _ := url.Parse("https://example.com")
						return model.URL{URL: u}
					}(),
				},
			},
		},
		{
			product:      model.Product{SKU: "12345"},
			skipCreation: true,
		},
	}

	// gormDSN is the Data Source Name for GORM
	gormDSN := os.Getenv("GORM_DSN")
	if gormDSN == "" {
		t.Fatal(`missing environment variable "GORM_DSN"`)
	}

	db, err := NewGormDB(gormDSN)
	if err != nil {
		t.Fatal(err)
	}

	storage := ProductStore{DB: db}
	if *verbose {
		storage.DB = db.Debug()
	}

	for i, v := range tdt {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			if !v.skipCreation {
				_ = storage.Create(&v.product)
				t.Cleanup(func() {
					_ = storage.Delete(v.product.SKU)
				})
			}

			err := storage.Update(v.product.SKU, v.product)
			if !errors.Is(err, v.expectedErr) {
				t.Fatalf("expected error '%v' unexpected error '%v'", v.expectedErr, err)
			}

			if err != nil {
				t.Skip(err)
			}

			t.Logf("%+v", v.product)
		})
	}
}

func TestProductStore_List(t *testing.T) {
	tdt := []struct {
		products     []model.Product
		expectedErr  error
		skipCreation bool
	}{
		{
			products: model.Products{
				{
					SKU:   "1234",
					Name:  "...",
					Brand: "Nike",
					Size:  &[]string{"M"}[0],
					Price: 1_000,
					PrincipalImage: func() *model.URL {
						u, _ := url.Parse("https://example.com")
						return &model.URL{URL: u}
					}(),
					OtherImages: model.URLs{
						func() model.URL {
							u, _ := url.Parse("https://example.com")
							return model.URL{URL: u}
						}(),
					},
				},
			},
		},
		{
			products:     model.Products{},
			skipCreation: true,
		},
	}

	// gormDSN is the Data Source Name for GORM
	gormDSN := os.Getenv("GORM_DSN")
	if gormDSN == "" {
		t.Fatal(`missing environment variable "GORM_DSN"`)
	}

	db, err := NewGormDB(gormDSN)
	if err != nil {
		t.Fatal(err)
	}

	storage := ProductStore{DB: db}
	if *verbose {
		storage.DB = db.Debug()
	}

	for i, v := range tdt {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			if !v.skipCreation {
				for _, product := range v.products {
					_ = storage.Create(&product)

				}
				t.Cleanup(func() {
					for _, product := range v.products {
						_ = storage.Delete(product.SKU)
					}
				})
			}

			products, err := storage.List()
			if !errors.Is(err, v.expectedErr) {
				t.Fatalf("expected error '%v' unexpected error '%v'", v.expectedErr, err)
			}

			if err != nil {
				t.Skip(err)
			}

			if !reflect.DeepEqual(v.products, products) {
				t.Fatalf("expected products '%v' unexpected products '%v'", v.products, products)
			}

			t.Log(products)
		})
	}
}
