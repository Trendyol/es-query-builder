package textquerytype

// TextQueryType represents the type of text query to use in query string queries.
//
// TextQueryType is a string type that defines how the query text is analyzed and matched
// against fields in the search index. It allows specifying various query strategies
// based on the desired behavior.
//
// Constants:
//   - Bestfields: Use the best matching fields for scoring.
//   - Mostfields: Combine scores from all matching fields.
//   - Crossfields: Treat matching fields as a single field.
//   - Phrase: Match terms as a phrase.
//   - Phraseprefix: Match terms as a phrase with a prefix.
//   - Boolprefix: Use a boolean query with prefix matching.
//
// Example usage:
//
//	var queryType TextQueryType = Bestfields
//
//	// Use queryType in a query string configuration.
type TextQueryType string

const (
	// Bestfields indicates that the best matching fields should be used for scoring.
	Bestfields TextQueryType = "best_fields"

	// Mostfields indicates that the scores from all matching fields should be combined.
	Mostfields TextQueryType = "most_fields"

	// Crossfields indicates that matching fields should be treated as a single field.
	Crossfields TextQueryType = "cross_fields"

	// Phrase indicates that terms should be matched as a phrase.
	Phrase TextQueryType = "phrase"

	// Phraseprefix indicates that terms should be matched as a phrase with a prefix.
	Phraseprefix TextQueryType = "phrase_prefix"

	// Boolprefix indicates that a boolean query with prefix matching should be used.
	Boolprefix TextQueryType = "bool_prefix"
)

func (textQueryType TextQueryType) String() string {
	return string(textQueryType)
}
