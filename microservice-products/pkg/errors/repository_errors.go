package errors

import "errors"

// There is a collection of repository error variables defined here.
// It should be used to represent specific error conditions in the repository layer.
var (
	ErrProductNotFound = errors.New("product not found")
)
