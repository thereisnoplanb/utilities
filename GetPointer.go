package utilities

// Returns a pointer to a value.
//
// # Parameters
//
//	value *T
//
// The value to which the pointer is returned.
//
// # Returns
//
//	pointer *T
//
// A pointer to a value.
func GetPointer[T any](value T) (pointer *T) {
	return &value
}
