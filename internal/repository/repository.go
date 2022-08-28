// Package repository is the data access layer, contains everything related to data persistence mechanisms
package repository

// StorageManager defines the common methods for storage management
type StorageManager[K comparable, V any] interface {
	// Create saves a new record into the storage
	Create(*V) error
	// Obtain returns the record identified by K from the store
	Obtain(K) (K, V)
	// Update updates a record identified by K
	Update(K, V) error
	// Delete removes a record identified by K from the store
	Delete(K) error
	// List returns all records into the store
	List() error
}
