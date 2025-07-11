package utilities

// Returns the value pointed to by the given pointer if it is not nil; otherwise returns the provided fallback value.
//
// # Parameters
//
//	pointer *T
//
// A pointer to a value of type T.
//
//	fallback T
//
// The fallback value to return if pointer is nil.
//
// # Returns
//
//	value T
//
// The value pointed to by pointer, or the fallback value if pointer is nil.
func GetValueOrFallback[T any](pointer *T, fallback T) (value T) {
	if pointer != nil {
		return *pointer
	}
	return fallback
}
