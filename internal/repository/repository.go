// Package repository is the data access layer, contains everything related to data persistence mechanisms
package repository

import (
	"fmt"
	error2 "github.com/yael-castro/products-api/internal/model/error"
)

// StorageManager defines the common methods for storage management
type StorageManager[K comparable, V any] interface {
	// Create saves a new record into the storage
	Create(*V) error
	// Obtain returns the record identified by K from the store
	Obtain(K) (V, error)
	// Update updates a record identified by K
	Update(K, V) error
	// Delete removes a record identified by K from the store
	Delete(K) error
	// List returns all records into the store
	List() ([]V, error)
}

// MockStorage simulates data persistence to test some features more easy, also can be used as a memory repository
type MockStorage[K comparable, V any] map[K]V

// Create missing implementation
func (m *MockStorage[K, V]) Create(v *V) error {
	return nil
}

// Obtain returns the record associate to the key received as parameter
func (m MockStorage[K, V]) Obtain(k K) (V, error) {
	v, ok := m[k]
	if !ok {
		return v, error2.NotFound(fmt.Sprintf(`product identified by sku '%v' does not exist`, k))
	}

	return v, nil
}

// Update replaces the record associate to the key received as parameter
func (m *MockStorage[K, V]) Update(k K, v V) error {
	(*m)[k] = v
	return nil
}

// Delete removes the record associate to the key received as parameter
func (m *MockStorage[K, V]) Delete(k K) error {
	delete(*m, k)
	return nil
}

// List returns all records from the hash map
func (m MockStorage[K, V]) List() ([]V, error) {
	list := make([]V, 0)

	for _, v := range m {
		list = append(list, v)
	}

	return list, nil
}
