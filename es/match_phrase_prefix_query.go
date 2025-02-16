package es

import ZeroTermsQuery "github.com/Trendyol/es-query-builder/es/enums/zero-terms-query"

type matchPhrasePrefixType Object

// MatchPhrasePrefix creates a new es.matchPhrasePrefixType object with the specified field and query.
//
// This function initializes an es.matchPhrasePrefixType object for a match phrase prefix query, where the key
// is the field name and query is the value to search for in that field. This is used
// to construct queries that match the specified value in the given field.
//
// Example usage:
//
//	m := es.MatchPhrasePrefix("title", "es-query-builder")
//	// m now contains an es.matchPhrasePrefixType object that matches the query "es-query-builder" in the "title" field.
//
// Parameters:
//   - key: A string representing the field name for the match phrase prefix query.
//   - query: The value to be matched in the specified field. The type is generic.
//
// Returns:
//
//	An es.matchPhrasePrefixType object containing the specified match phrase prefix query.
func MatchPhrasePrefix[T any](key string, query T) matchPhrasePrefixType {
	return matchPhrasePrefixType{
		"match_phrase_prefix": Object{
			key: Object{
				"query": query,
			},
		},
	}
}

// Analyzer sets the "analyzer" field in the match phrase prefix query.
//
// This method specifies the analyzer to use for the match phrase prefix query, which determines
// how the input text is processed during analysis (e.g., tokenization and normalization).
// Custom analyzers can be used to tailor the query behavior to specific requirements.
//
// Example usage:
//
//	m := es.MatchPhrasePrefix("title", "es-query-builder").Analyzer("custom_analyzer")
//	// m now has an "analyzer" field set to "custom_analyzer" in the match phrase prefix query object.
//
// Parameters:
//   - value: A string representing the name of the analyzer to use.
//
// Returns:
//
//	The updated es.matchPhrasePrefixType object with the "analyzer" field set to the specified value.
func (m matchPhrasePrefixType) Analyzer(value string) matchPhrasePrefixType {
	return m.putInTheField("analyzer", value)
}

// Boost sets the "boost" field in the match phrase prefix query.
//
// This method configures the match phrase prefix query to use a specified boost factor, which influences
// the relevance scoring of the matched documents.
//
// Example usage:
//
//	m := es.MatchPhrasePrefix("title", "es-query-builder").Boost(1.5)
//	// m now has a "boost" field set to 1.5 in the match phrase prefix query object.
//
// Parameters:
//   - boost: A float64 value representing the boost factor to be applied to the match phrase prefix query.
//
// Returns:
//
//	The updated es.matchPhrasePrefixType object with the "boost" field set to the specified value.
func (m matchPhrasePrefixType) Boost(boost float64) matchPhrasePrefixType {
	return m.putInTheField("boost", boost)
}

// MaxExpansions sets the "max_expansions" field in the match phrase prefix query.
//
// This method configures the match phrase prefix query to limit the maximum number of terms that can be expanded
// for multi-term queries, such as those involving fuzzy matching. Higher values allow more terms to
// be considered, but may impact performance.
//
// Example usage:
//
//	m := es.MatchPhrasePrefix("title", "es-query-builder").MaxExpansions(50)
//	// m now has a "max_expansions" field set to 50 in the match phrase prefix query object.
//
// Parameters:
//   - maxExpansions: An integer representing the maximum number of term expansions to be allowed in the match phrase prefix query.
//
// Returns:
//
//	The updated es.matchPhrasePrefixType object with the "max_expansions" field set to the specified value.
func (m matchPhrasePrefixType) MaxExpansions(maxExpansions int) matchPhrasePrefixType {
	return m.putInTheField("max_expansions", maxExpansions)
}

// ZeroTermsQuery sets the "zero_terms_query" field in the match phrase prefix query.
//
// This method configures the behavior of the match phrase prefix query when no terms remain after analysis
// (for example, if all terms are stop words). The specified zero_terms_query value determines
// how to handle this scenario, with options like "all" to match all documents or "none" to
// match none.
//
// Example usage:
//
//	m := es.MatchPhrasePrefix("title", "es-query-builder").ZeroTermsQuery(zerotermsquery.All)
//	// m now has a "zero_terms_query" field set to "all" in the match phrase prefix query object.
//
// Parameters:
//   - zeroTermsQuery: A zerotermsquery.ZeroTermsQuery value that specifies the behavior for zero-term queries.
//
// Returns:
//
//	The updated es.matchPhrasePrefixType object with the "zero_terms_query" field set to the specified value.
func (m matchPhrasePrefixType) ZeroTermsQuery(zeroTermsQuery ZeroTermsQuery.ZeroTermsQuery) matchPhrasePrefixType {
	return m.putInTheField("zero_terms_query", zeroTermsQuery)
}

// Slop sets the "slop" field in the match phrase prefix query.
//
// This method specifies the allowed distance between terms in a phrase query, enabling more
// flexibility in matching phrases that may have slight variations in word order or spacing.
// A higher slop value allows more variation, while a slop of 0 requires exact matching.
//
// Example usage:
//
//	m := es.MatchPhrasePrefix("title", "es-query-builder").Slop(2)
//	// m now has a "slop" field set to 2 in the match phrase prefix query object.
//
// Parameters:
//   - slop: An integer representing the maximum allowed distance between terms.
//
// Returns:
//
//	The updated es.matchPhrasePrefixType object with the "slop" field set to the specified value.
func (m matchPhrasePrefixType) Slop(slop int) matchPhrasePrefixType {
	return m.putInTheField("slop", slop)
}

func (m matchPhrasePrefixType) putInTheField(key string, value any) matchPhrasePrefixType {
	if matchPhrasePrefix, ok := m["match_phrase_prefix"].(Object); ok {
		for _, fieldObj := range matchPhrasePrefix {
			if fieldObject, foOk := fieldObj.(Object); foOk {
				fieldObject[key] = value
				break
			}
		}
	}
	return m
}
