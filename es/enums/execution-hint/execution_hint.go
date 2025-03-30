package executionhint

// ExecutionHint represents the hint for how terms aggregation should be executed.
//
// This type specifies the method Elasticsearch should use to compute the terms
// in a terms aggregation. It helps Elasticsearch optimize the aggregation
// process based on the data and query structure.
//
// Example usage:
//
//	executionHint := executionhint.GlobalOrdinals
//	// executionHint now holds the value "global_ordinals" for execution type.
//
// Parameters:
//   - ExecutionHint: A type representing the execution method for terms aggregation.
//     It can be either Map, GlobalOrdinals, or FieldData.
//
// Returns:
//
//	The string representation of the execution hint type.
type ExecutionHint string

const (
	// Map indicates that terms should be computed using a map-based approach.
	// This method processes the terms by mapping the values from the documents.
	// It is generally used for smaller data sets.
	Map ExecutionHint = "map"

	// GlobalOrdinals indicates that terms should be computed using global ordinals.
	// This method is more efficient for large datasets and high-cardinality fields.
	// It relies on a global ordinality to speed up computation and reduce memory usage.
	GlobalOrdinals ExecutionHint = "global_ordinals"

	// FieldData indicates that terms should be computed using fielddata.
	// This method uses fielddata for aggregations on text fields.
	// It is commonly used for text-based fields but can be memory-intensive.
	FieldData ExecutionHint = "fielddata"
)

// String returns the string representation of the ExecutionHint.
func (executionHint ExecutionHint) String() string {
	return string(executionHint)
}
