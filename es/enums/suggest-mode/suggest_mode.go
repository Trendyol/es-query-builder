package suggestmode

// SuggestMode controls which suggestions are included from a term suggester.
//
// Example usage:
//
//	suggester := es.TermSuggester("message").SuggestMode(SuggestMode.Popular)
type SuggestMode string

const (
	// Missing only provides suggestions for terms not in the index.
	Missing SuggestMode = "missing"

	// Popular only suggests terms that appear more frequently than the input term.
	Popular SuggestMode = "popular"

	// Always suggests any matching suggestion based on edit distance.
	Always SuggestMode = "always"
)

func (suggestMode SuggestMode) String() string {
	return string(suggestMode)
}
