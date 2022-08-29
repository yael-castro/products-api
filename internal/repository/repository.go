// Package repository is the data access layer, contains everything related to data persistence mechanisms
package repository

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

type MockStorage[K comparable, V any] map[K]V

func (m *MockStorage[K, V]) Create(v *V) error {
	return nil
}

func (m MockStorage[K, V]) Obtain(k K) (V, error) {
	return m[k], nil
}

func (m *MockStorage[K, V]) Update(k K, v V) error {
	(*m)[k] = v
	return nil
}

func (m *MockStorage[K, V]) Delete(k K) error {
	delete(*m, k)
	return nil
}

func (m MockStorage[K, V]) List() ([]V, error) {
	list := make([]V, 0)

	for _, v := range m {
		list = append(list, v)
	}

	return list, nil
}
