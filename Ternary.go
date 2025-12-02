package utilities

// Evaluates a bool expression and returns the result of one of the two expressions, depending on whether the bool expression evaluates to true or false.
//
// # Parameters
//
//	condition bool
//
// The condition expression must evaluate to true or false.
//
//	consequent T
//
// A value returned when the condition evaluates to true.
//
//	alternative T
//
// A value returned when the condition evaluates to false.
//
// # Returns
// If condition evaluates to true, the consequent value is returned, the alternative value otherwise.
//
// # Example
//
//	result := Ternary(5 > 1, 100, 0)
//
//	// result = 100
func Ternary[T any](condition bool, consequent, alternative T) T {
	if condition {
		return consequent
	}
	return alternative
}

// Evaluates a bool expression and returns the result of one of the two expressions, depending on whether the bool expression evaluates to true or false.
//
// # Parameters
//
//	condition bool
//
// A condition expression that must evaluate to true or false.
//
//	consequent T
//
// A consequent expression that is evaluated when condition is true, and its result becomes the result of the operation.
//
//	alternative T
//
// An alternative expression that is evaluated when condition is false, and its result becomes the result of the operation.
//
// # Returns
//
// If condition evaluates to true, the consequent expression is evaluated, and its result becomes the result of the operation.
// If condition evaluates to false, the alternative expression is evaluated, and its result becomes the result of the operation.
// Only consequent or alternative expression is evaluated.
//
// # Example
//
//	slice := []int{}
//	result := TernaryFunc(len(slice) > 0, func() int { return slice[0] }, func() { return -1 })
//
//	// result = -1
//
//	slice := []int{1,2,3}
//	result := TernaryFunc(len(slice) > 0, func() int { return slice[0] }, func() { return -1 })
//
//	// result = 1
func TernaryFunc[T any](condition bool, consequent, alternative func() T) T {
	if condition {
		return consequent()
	}
	return alternative()
}
