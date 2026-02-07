package boundaryscanner

// BoundaryScanner represents the different boundary scanner types for highlighting in Elasticsearch.
//
// BoundaryScanner is a string type used to specify how highlighted fragments are bounded.
// It determines the strategy for finding the boundaries of highlighted snippets.
//
// Example usage:
//
//	var bs BoundaryScanner = Sentence
//
//	// Use bs in a highlight configuration
//
// Constants:
//   - Chars: The chars boundary scanner breaks highlighted fragments at characters specified by boundary_chars.
//   - Sentence: The sentence boundary scanner breaks highlighted fragments at the next sentence boundary.
//   - Word: The word boundary scanner breaks highlighted fragments at the next word boundary.
type BoundaryScanner string

const (
	// Chars indicates that the boundary scanner should break at specified boundary characters.
	Chars BoundaryScanner = "chars"

	// Sentence indicates that the boundary scanner should break at sentence boundaries.
	Sentence BoundaryScanner = "sentence"

	// Word indicates that the boundary scanner should break at word boundaries.
	Word BoundaryScanner = "word"
)

func (boundaryScanner BoundaryScanner) String() string {
	return string(boundaryScanner)
}
