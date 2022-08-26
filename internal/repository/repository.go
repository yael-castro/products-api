// Package repository is the data access layer, contains everything related to data persistence mechanisms
package repository

// Type defines the available storage types
type Type uint

// Repository types
const (
	// SQL indicates a SQL store like Postgresql or MySQL
	SQL Type = iota
	// NoSQL indicates a NoSQL store like MongoDB
	NoSQL
	// Memory indicates that the data will persist in a memory store like Redis
	Memory
)
