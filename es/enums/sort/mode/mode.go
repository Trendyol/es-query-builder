package mode

// Mode represents various aggregation modes for queries or operations.
//
// Mode is a string type used to specify different ways to aggregate or process
// values in queries. It provides various options for combining or analyzing data.
//
// Example usage:
//
//	var mode Mode = Sum
//
//	// Use mode in a query or operation configuration
//
// Constants:
//   - Min: Minimum value from the set of values.
//   - Max: Maximum value from the set of values.
//   - Sum: Sum of all values in the set.
//   - Avg: Average of all values in the set.
//   - Median: Median value from the set of values.
//   - Default: The default aggregation mode, typically used as a fallback.
type Mode string

const (
	// Min indicates that the minimum value from the set should be used.
	Min Mode = "min"

	// Max indicates that the maximum value from the set should be used.
	Max Mode = "max"

	// Sum indicates that the sum of all values in the set should be used.
	Sum Mode = "sum"

	// Avg indicates that the average of all values in the set should be used.
	Avg Mode = "avg"

	// Median indicates that the median value from the set should be used.
	Median Mode = "median"

	// Default indicates the default aggregation mode, used as a fallback or if no specific mode is provided.
	Default Mode = "_default"
)

func (mode Mode) String() string {
	return string(mode)
}
