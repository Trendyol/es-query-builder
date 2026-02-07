package highlightertype

// HighlighterType represents the different highlighter types available in Elasticsearch.
//
// HighlighterType is a string type used to specify which highlighter implementation
// to use for highlighting search results. Elasticsearch provides three highlighter
// implementations: unified, plain, and fvh (Fast Vector Highlighter).
//
// Example usage:
//
//	var ht HighlighterType = Unified
//
//	// Use ht in a highlight configuration
//
// Constants:
//   - Unified: The unified highlighter uses the Lucene Unified Highlighter (default).
//   - Plain: The plain highlighter uses the standard Lucene highlighter.
//   - Fvh: The fvh (Fast Vector Highlighter) highlighter uses the Lucene Fast Vector Highlighter.
type HighlighterType string

const (
	// Unified indicates the unified highlighter, which is the default highlighter in Elasticsearch.
	Unified HighlighterType = "unified"

	// Plain indicates the plain highlighter, which uses the standard Lucene highlighter.
	Plain HighlighterType = "plain"

	// Fvh indicates the Fast Vector Highlighter, which requires term_vector set to with_positions_offsets.
	Fvh HighlighterType = "fvh"
)

func (highlighterType HighlighterType) String() string {
	return string(highlighterType)
}
