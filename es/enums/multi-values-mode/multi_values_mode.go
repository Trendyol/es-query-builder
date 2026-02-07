package multivaluesmode

// MultiValuesMode represents the different modes for handling multi-valued fields in decay functions.
//
// MultiValuesMode is a string type used to specify how a decay function should handle
// fields that contain multiple values. It determines which value is chosen for the
// distance calculation.
//
// Example usage:
//
//	var mode MultiValuesMode = Min
//
//	// Use mode in a decay function configuration
//
// Constants:
//   - Min: Use the minimum distance from the origin.
//   - Max: Use the maximum distance from the origin.
//   - Avg: Use the average distance from the origin.
//   - Sum: Use the sum of all distances from the origin.
type MultiValuesMode string

const (
	// Min indicates that the minimum distance should be used.
	Min MultiValuesMode = "min"

	// Max indicates that the maximum distance should be used.
	Max MultiValuesMode = "max"

	// Avg indicates that the average distance should be used.
	Avg MultiValuesMode = "avg"

	// Sum indicates that the sum of distances should be used.
	Sum MultiValuesMode = "sum"
)

func (multiValuesMode MultiValuesMode) String() string {
	return string(multiValuesMode)
}
