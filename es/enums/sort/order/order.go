package order

// Order represents the sorting order for queries or operations.
//
// Order is a string type used to specify the direction of sorting or ordering
// for query results or data operations. It provides options for ascending or
// descending order, as well as a default sorting option.
//
// Example usage:
//
//	var order Order = Asc
//
//	// Use order in a query or operation configuration
//
// Constants:
//   - Asc: Ascending order for sorting or ordering.
//   - Desc: Descending order for sorting or ordering.
//   - Default: The default sorting order, typically used as a fallback.
type Order string

const (
	// Asc indicates that the results should be sorted in ascending order.
	Asc Order = "asc"

	// Desc indicates that the results should be sorted in descending order.
	Desc Order = "desc"

	// Default indicates the default sorting order, used as a fallback or if no specific order is provided.
	Default Order = "_default"
)

func (order Order) String() string {
	return string(order)
}
