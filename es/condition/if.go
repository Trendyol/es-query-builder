package condition

// If returns the provided `item` if `condition` is true; otherwise, it returns nil.
// This function is generic and operates on maps with string keys or slices of any type.
//
// Parameters:
//   - item: The map or slice to potentially return if the condition is met.
//   - condition: A boolean that determines if `item` should be returned.
//
// Returns:
//   - The original `item` (map or slice) if `condition` is true, otherwise nil.
//
// Example usage:
//
//	result := condition.If(es.Object{"key": "value"}, true)
//	// result is es.Object{"key": "value"}
//
//	result = condition.If(es.Object{"key": "value"}, false)
//	// result is nil
func If[T ~map[string]any | ~[]any](item T, condition bool) T {
	if !condition {
		return nil
	}
	return item
}
