package fragmenter

// Fragmenter represents the different fragmenter types for highlighting in Elasticsearch.
//
// Fragmenter is a string type used to specify how text should be broken up into
// highlight fragments. It determines the strategy for creating text fragments
// that contain highlighted terms.
//
// Example usage:
//
//	var f Fragmenter = Span
//
//	// Use f in a highlight configuration
//
// Constants:
//   - Simple: The simple fragmenter breaks text into same-sized fragments.
//   - Span: The span fragmenter breaks text into same-sized fragments, but tries to avoid breaking up highlighted terms.
type Fragmenter string

const (
	// Simple indicates that text should be broken into same-sized fragments.
	Simple Fragmenter = "simple"

	// Span indicates that text should be broken into same-sized fragments while trying to avoid breaking highlighted terms.
	Span Fragmenter = "span"
)

func (fragmenter Fragmenter) String() string {
	return string(fragmenter)
}
