package zerotermsquery

// ZeroTermsQuery sets the "zero_terms_query" field in the matchType object.
//
// This method specifies the behavior of the match query when no terms remain after analysis,
// such as when all terms are stop words. It updates the matchType object to include the provided
// zeroTermsQuery value, which determines how the query should respond in this scenario.
//
// Example usage:
//
//	match := es.Match("field", "value").ZeroTermsQuery(ZeroTermsQuery.All)
//	// match now has the "zero_terms_query" field set to "all" in the matchType object.
//
// Parameters:
//   - zeroTermsQuery: A ZeroTermsQuery value indicating the behavior for zero-term queries.
//     It can be either ZeroTermsQuery.All or ZeroTermsQuery.None.
//
// Returns:
//
//	The updated matchType object with the "zero_terms_query" field set to the specified value.
type ZeroTermsQuery string

const (
	// All indicates that all documents should be matched when no terms remain.
	All ZeroTermsQuery = "all"

	// None indicates that no documents should be matched when no terms remain.
	None ZeroTermsQuery = "none"
)

func (zeroTermsQuery ZeroTermsQuery) String() string {
	return string(zeroTermsQuery)
}
