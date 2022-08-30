package handler

import (
	"bytes"
	"flag"
	"github.com/gin-gonic/gin"
	"github.com/yael-castro/products-api/internal/business"
	"github.com/yael-castro/products-api/internal/model"
	"github.com/yael-castro/products-api/internal/repository"
	"io"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
)

// verbose cli flag for "V", indicates whether to show additional logs, such as request logs
var verbose = flag.Bool("V", false, "")

func TestProductStore_CreateProduct(t *testing.T) {
	tdt := []struct {
		request      *http.Request
		expectedCode int
		// TODO: validate response data
	}{
		{
			request: func() *http.Request {
				request, _ := http.NewRequest(http.MethodPost, "/v1/products", bytes.NewBuffer([]byte(`{}`)))
				return request
			}(),
			expectedCode: http.StatusBadRequest,
		},
		{
			request: func() *http.Request {
				request, _ := http.NewRequest(http.MethodPost, "/v1/products", bytes.NewBuffer([]byte(`{
    				"sku": "FAL-12345677",
    				"name": "...",
    				"brand": "...",
    				"size": "M",
    				"price": 1.2,
					"principalImage": "https://example.com",
    				"otherImages": ["https://example.com"]
				}`)))

				request.Header.Set("Content-Type", "application/json")
				return request
			}(),
			expectedCode: http.StatusCreated,
		},
	}

	gin.SetMode(gin.TestMode)
	if *verbose {
		gin.SetMode(gin.DebugMode)
	}

	store := ProductStore{
		ProductManager: business.ProductStore{
			StorageManager: &repository.MockStorage[model.SKU, model.Product]{},
		},
	}

	// c := &gin.Context{Request: v.request, Writer: httptest.NewRecorder()}
	for i, v := range tdt {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)
			c.Request = v.request

			store.CreateProduct(c)

			if w.Code != v.expectedCode {
				t.Errorf(`expected code '%d' unexpected code '%d'`, v.expectedCode, w.Code)
			}

			data, err := io.ReadAll(w.Body)
			if err != nil {
				t.Fatal(err)
			}

			t.Log(string(data))
		})
	}
}

func TestProductStore_ObtainProduct(t *testing.T) {
	tdt := []struct {
		request      *http.Request
		expectedCode int
		// TODO: validate response data
	}{
		{
			request: func() *http.Request {
				request, _ := http.NewRequest(http.MethodGet, "/v1/products/FAL-12345678", nil)
				return request
			}(),
			expectedCode: http.StatusOK,
		},
		{
			request: func() *http.Request {
				request, _ := http.NewRequest(http.MethodGet, "/v1/products/FAL-12345679", nil)
				return request
			}(),
			expectedCode: http.StatusNotFound,
		},
		{
			request: func() *http.Request {
				request, _ := http.NewRequest(http.MethodGet, "/v1/products/1234", nil)
				return request
			}(),
			expectedCode: http.StatusBadRequest,
		},
	}

	gin.SetMode(gin.TestMode)
	if *verbose {
		gin.SetMode(gin.DebugMode)
	}

	store := ProductStore{
		ProductManager: business.ProductStore{
			StorageManager: &repository.MockStorage[model.SKU, model.Product]{
				"FAL-12345678": model.Product{},
			},
		},
	}

	engine := gin.New()
	engine.GET("/v1/products/:id", store.ObtainProduct)

	for i, v := range tdt {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			w := httptest.NewRecorder()

			engine.ServeHTTP(w, v.request)

			if w.Code != v.expectedCode {
				t.Errorf(`expected code '%d' unexpected code '%d'`, v.expectedCode, w.Code)
			}

			data, err := io.ReadAll(w.Body)
			if err != nil {
				t.Fatal(err)
			}

			t.Log(string(data))
		})
	}
}
