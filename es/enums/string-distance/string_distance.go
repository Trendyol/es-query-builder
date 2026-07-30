package stringdistance

// StringDistance controls the string distance algorithm used by a term suggester.
//
// Example usage:
//
//	suggester := es.TermSuggester("message").StringDistance(StringDistance.Levenshtein)
type StringDistance string

const (
	// Internal is the default based on damerau_levenshtein with optimizations.
	Internal StringDistance = "internal"

	// DamerauLevenshtein measures edit distance including transpositions.
	DamerauLevenshtein StringDistance = "damerau_levenshtein"

	// Levenshtein measures classic Levenshtein edit distance.
	Levenshtein StringDistance = "levenshtein"

	// JaroWinkler measures Jaro-Winkler string distance.
	JaroWinkler StringDistance = "jaro_winkler"

	// Ngram measures distance based on character n-grams.
	Ngram StringDistance = "ngram"
)

func (stringDistance StringDistance) String() string {
	return string(stringDistance)
}
