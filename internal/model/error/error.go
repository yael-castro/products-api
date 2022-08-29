// Package error contains everything related to error management and error handling
package error

// NotFound error caused by missing resource
type NotFound string

// Error returns the string value of NotFound
func (n NotFound) Error() string {
	return string(n)
}
