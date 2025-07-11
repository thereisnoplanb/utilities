package utilities

// Returns the value pointed to by the given pointer if it is not nil; otherwise returns the default value.
//
// # Parameters
//
//	pointer *T
//
// A pointer to a value of type T.
//
// # Returns
//
//	value T
//
// The value pointed to by pointer, or the default value if pointer is nil.
func GetValueOrDefault[T any](pointer *T) (result T) {
	if pointer != nil {
		result = *pointer
	}
	return result
}
