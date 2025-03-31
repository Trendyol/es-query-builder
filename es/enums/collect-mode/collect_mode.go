package collectmode

// CollectMode represents the mode in which terms aggregation collects terms.
//
// This type specifies how terms are collected in the terms aggregation.
// It determines whether Elasticsearch uses breadth-first or depth-first traversal
// when gathering terms for aggregation.
//
// Example usage:
//
//	collectMode := collect_mode.BreadthFirst
//	// collectMode now holds the value "breadth_first" for traversal type.
//
// Parameters:
//   - CollectMode: A type representing the collection mode for terms aggregation.
//     It can be either BreadthFirst or DepthFirst.
//
// Returns:
//
//	The string representation of the collect mode type.
type CollectMode string

const (
	// BreadthFirst indicates that terms should be collected using breadth-first traversal.
	// This method collects terms in a broader search space first and then refines the search.
	// It can be more memory efficient for certain cases.
	BreadthFirst CollectMode = "breadth_first"

	// DepthFirst indicates that terms should be collected using depth-first traversal.
	// This method collects terms in a more granular, depth-first manner, often leading to deeper analysis
	// on fewer terms at a time.
	// It may provide better precision in certain circumstances, but is more memory intensive.
	DepthFirst CollectMode = "depth_first"
)

// String returns the string representation of the CollectMode.
func (collectMode CollectMode) String() string {
	return string(collectMode)
}
