package es

import ZeroTermsQuery "github.com/Trendyol/es-query-builder/es/enums/zero-terms-query"

type matchPhraseType Object

// MatchPhrase creates a new es.matchPhraseType object with the specified field and query.
//
// This function initializes an es.matchPhraseType object for a match phrase query, where the key
// is the field name and query is the value to search for in that field. This is used
// to construct queries that match the specified value in the given field.
//
// Example usage:
//
//	m := es.MatchPhrase("title", "es-query-builder")
//	// m now contains an es.matchPhraseType object that matches the query "es-query-builder" in the "title" field.
//
// Parameters:
//   - key: A string representing the field name for the match phrase query.
//   - query: The value to be matched in the specified field. The type is generic.
//
// Returns:
//
//	An es.matchPhraseType object containing the specified match phrase query.
func MatchPhrase[T any](key string, query T) matchPhraseType {
	return matchPhraseType{
		"match_phrase": Object{
			key: Object{
				"query": query,
			},
		},
	}
}

// Analyzer sets the "analyzer" field in the match phrase query.
//
// This method specifies the analyzer to use for the match phrase query, which determines
// how the input text is processed during analysis (e.g., tokenization and normalization).
// Custom analyzers can be used to tailor the query behavior to specific requirements.
//
// Example usage:
//
//	m := es.MatchPhrase("title", "es-query-builder").Analyzer("custom_analyzer")
//	// m now has an "analyzer" field set to "custom_analyzer" in the match phrase query object.
//
// Parameters:
//   - value: A string representing the name of the analyzer to use.
//
// Returns:
//
//	The updated es.matchPhraseType object with the "analyzer" field set to the specified value.
func (m matchPhraseType) Analyzer(value string) matchPhraseType {
	return m.putInTheField("analyzer", value)
}

// Boost sets the "boost" field in the match phrase query.
//
// This method configures the match phrase query to use a specified boost factor, which influences
// the relevance scoring of the matched documents.
//
// Example usage:
//
//	m := es.MatchPhrase("title", "es-query-builder").Boost(1.5)
//	// m now has a "boost" field set to 1.5 in the match phrase query object.
//
// Parameters:
//   - boost: A float64 value representing the boost factor to be applied to the match phrase query.
//
// Returns:
//
//	The updated es.matchPhraseType object with the "boost" field set to the specified value.
func (m matchPhraseType) Boost(boost float64) matchPhraseType {
	return m.putInTheField("boost", boost)
}

// ZeroTermsQuery sets the "zero_terms_query" field in the match phrase query.
//
// This method configures the behavior of the match phrase query when no terms remain after analysis
// (for example, if all terms are stop words). The specified zero_terms_query value determines
// how to handle this scenario, with options like "all" to match all documents or "none" to
// match none.
//
// Example usage:
//
//	m := es.MatchPhrase("title", "es-query-builder").ZeroTermsQuery(zerotermsquery.All)
//	// m now has a "zero_terms_query" field set to "all" in the match phrase query object.
//
// Parameters:
//   - zeroTermsQuery: A zerotermsquery.ZeroTermsQuery value that specifies the behavior for zero-term queries.
//
// Returns:
//
//	The updated es.matchPhraseType object with the "zero_terms_query" field set to the specified value.
func (m matchPhraseType) ZeroTermsQuery(zeroTermsQuery ZeroTermsQuery.ZeroTermsQuery) matchPhraseType {
	return m.putInTheField("zero_terms_query", zeroTermsQuery)
}

// Slop sets the "slop" field in the match phrase query.
//
// This method specifies the allowed distance between terms in a phrase query, enabling more
// flexibility in matching phrases that may have slight variations in word order or spacing.
// A higher slop value allows more variation, while a slop of 0 requires exact matching.
//
// Example usage:
//
//	m := es.MatchPhrase("title", "es-query-builder").Slop(2)
//	// m now has a "slop" field set to 2 in the match phrase query object.
//
// Parameters:
//   - slop: An integer representing the maximum allowed distance between terms.
//
// Returns:
//
//	The updated es.matchPhraseType object with the "slop" field set to the specified value.
func (m matchPhraseType) Slop(slop int) matchPhraseType {
	return m.putInTheField("slop", slop)
}

func (m matchPhraseType) putInTheField(key string, value any) matchPhraseType {
	return genericPutInTheFieldOfFirstChild(m, "match_phrase", key, value)
}
