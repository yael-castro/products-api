// Package dependency contains every related to dependency injection and dependency management
package dependency

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/yael-castro/products-api/internal/business"
	"github.com/yael-castro/products-api/internal/handler"
	"github.com/yael-castro/products-api/internal/repository"
	"os"
)

// Profile defines options of dependency injection
type Profile uint

// Supported profiles for dependency injection
const (
	// Default defines the production profile
	Default Profile = iota
	// Testing defines the testing profile used to make a unit and integration tests
	Testing
)

// Injector defines a dependency injector
type Injector interface {
	// Inject takes any data type and fill of required dependencies (dependency injection)
	Inject(any) error
}

// InjectorFunc function that implements the Injector interface
type InjectorFunc func(any) error

// Inject executes f to inject dependencies to i
func (f InjectorFunc) Inject(a any) error {
	return f(a)
}

// NewInjector is an abstract factory to Injector, it builds an instance of Injector interface based on the Profile based as parameter
//
// Supported profiles: Default and Testing
//
// If pass a parameter an invalid profile it panics
func NewInjector(p Profile) Injector {
	switch p {
	case Default:
		return InjectorFunc(handlerDefault)
	}

	panic(fmt.Sprintf(`invalid profile: "%d" is not supported`, p))
}

// handlerDefault InjectorFunc for *handler.GinHandlers that uses a Default Profile
func handlerDefault(a any) error {
	engine, ok := a.(*gin.Engine)
	if !ok {
		return fmt.Errorf(`an instance of "%T" is required not "%T"`, engine, a)
	}

	gormDSN := os.Getenv("GORM_DSN")
	if gormDSN == "" {
		return errors.New("missing environment variable GORM_DSN")
	}

	db, err := repository.NewGormDB(gormDSN)
	if err != nil {
		return err
	}

	handlers := handler.GinHandlers{}

	handlers.ProductManager = handler.ProductStore{
		ProductManager: business.ProductStore{
			StorageManager: repository.ProductStore{
				DB: db,
			},
		},
	}

	*engine = *handler.NewGinEngine(handlers)
	return nil
}
