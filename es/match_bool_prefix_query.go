package es

import Operator "github.com/Trendyol/es-query-builder/es/enums/operator"

type matchBoolPrefixType Object

// MatchBoolPrefix creates a new es.matchBoolPrefix object with the specified field and query.
//
// This function initializes an es.matchBoolPrefix object for a match bool prefix query, where the key
// is the field name and query is the value to search for in that field. This is used
// to construct queries that match the specified value in the given field.
//
// Example usage:
//
//	m := es.MatchBoolPrefix("title", "es-query-builder")
//	// m now contains an es.matchBoolPrefix object that matches the query "es-query-builder" in the "title" field.
//
// Parameters:
//   - key: A string representing the field name for the match bool prefix query.
//   - query: The value to be matched in the specified field. The type is generic.
//
// Returns:
//
//	An es.matchBoolPrefix object containing the specified match bool prefix query.
func MatchBoolPrefix[T any](key string, query T) matchBoolPrefixType {
	return matchBoolPrefixType{
		"match_bool_prefix": Object{
			key: Object{
				"query": query,
			},
		},
	}
}

// Analyzer sets the analyzer to be used for the es.matchBoolPrefixType object.
//
// This method specifies the analyzer that should be applied to the query string.
// Analyzers are used to process the text, such as tokenizing and normalizing it,
// allowing for more refined and accurate search queries based on the specified analyzer.
//
// Example usage:
//
//	q := es.MatchBoolPrefix("title", "es-query-builder").Analyzer("standard")
//	// q now uses the "standard" analyzer for processing the query string.
//
// Parameters:
//   - value: A string representing the name of the analyzer to be used.
//
// Returns:
//
// The updated es.matchBoolPrefix object with the "analyzer" set.
func (m matchBoolPrefixType) Analyzer(value string) matchBoolPrefixType {
	return m.putInTheField("analyzer", value)
}

// MinimumShouldMatch sets the "minimum_should_match" field in the match bool prefix query.
//
// This method defines the minimum number of "should" clauses that must match for a document
// to be included in the results. The value can be an absolute number, a percentage, or a
// complex expression (e.g., "2", "75%", or "3<75%").
//
// Example usage:
//
//	m := es.MatchBoolPrefix("title", "es-query-builder").MinimumShouldMatch("75%")
//	// m now has a "minimum_should_match" field set to "75%" in the match bool prefix query object.
//
// Parameters:
//   - minimumShouldMatch: A value of any type specifying the minimum number of "should" clauses
//     required to match.
//
// Returns:
//
//	The updated es.matchBoolPrefixType object with the "minimum_should_match" field set to the specified value.
func (m matchBoolPrefixType) MinimumShouldMatch(minimumShouldMatch any) matchBoolPrefixType {
	return m.putInTheField("minimum_should_match", minimumShouldMatch)
}

// Operator sets the "operator" field in the match bool prefix query.
//
// This method configures the match bool prefix query to use a specified operator (e.g., "AND" or "OR")
// for the matching process.
//
// Example usage:
//
//	m := es.MatchBoolPrefix("title", "es-query-builder").Operator("AND")
//	// m now has an "operator" field set to "AND" in the match bool prefix query object.
//
// Parameters:
//   - operator: An Operator.Operator value representing the logical operator to be used in the match bool prefix query.
//
// Returns:
//
//	The updated es.matchBoolPrefix object with the "operator" field set to the specified value.
func (m matchBoolPrefixType) Operator(operator Operator.Operator) matchBoolPrefixType {
	return m.putInTheField("operator", operator)
}

// Boost sets the "boost" field in the match bool prefix query.
//
// This method configures the match bool prefix query to use a specified boost factor, which influences
// the relevance scoring of the matched documents.
//
// Example usage:
//
//	m := es.MatchBoolPrefix("title", "es-query-builder").Boost(1.5)
//	// m now has a "boost" field set to 1.5 in the match bool prefix query object.
//
// Parameters:
//   - boost: A float64 value representing the boost factor to be applied to the match bool prefix query.
//
// Returns:
//
//	The updated es.matchBoolPrefix object with the "boost" field set to the specified value.
func (m matchBoolPrefixType) Boost(boost float64) matchBoolPrefixType {
	return m.putInTheField("boost", boost)
}

// Fuzziness sets the "fuzziness" field in the match bool prefix query.
//
// This method configures the match bool prefix query to use a specified fuzziness level, which determines
// the allowed edit distance (e.g., number of character changes) for a term to be considered a match.
// Common values include "AUTO", or integers representing the number of edits (e.g., 1 or 2).
//
// Example usage:
//
//	m := es.MatchBoolPrefix("title", "es-query-builder").Fuzziness("AUTO")
//	// m now has a "fuzziness" field set to "AUTO" in the match bool prefix query object.
//
// Parameters:
//   - fuzziness: A value of any type (typically a string or integer) representing the fuzziness
//     level to be applied to the match bool prefix query.
//
// Returns:
//
//	The updated es.matchBoolPrefix object with the "fuzziness" field set to the specified value.
func (m matchBoolPrefixType) Fuzziness(fuzziness any) matchBoolPrefixType {
	return m.putInTheField("fuzziness", fuzziness)
}

// FuzzyRewrite sets the "fuzzy_rewrite" field in the match bool prefix query.
//
// This method configures the match bool prefix query to use a specified fuzzy rewrite method,
// which controls how multi-term queries are rewritten. Common values include "constant_score",
// "scoring_boolean", and other rewrite options that influence the scoring and performance of
// fuzzy matching.
//
// Example usage:
//
//	m := es.MatchBoolPrefix("title", "es-query-builder").FuzzyRewrite("constant_score")
//	// m now has a "fuzzy_rewrite" field set to "constant_score" in the match bool prefix query object.
//
// Parameters:
//   - fuzzyRewrite: A string value representing the rewrite method to be used for fuzzy matching in the match bool prefix query.
//
// Returns:
//
//	The updated es.matchBoolPrefix object with the "fuzzy_rewrite" field set to the specified value.
func (m matchBoolPrefixType) FuzzyRewrite(fuzzyRewrite string) matchBoolPrefixType {
	return m.putInTheField("fuzzy_rewrite", fuzzyRewrite)
}

// FuzzyTranspositions sets the "fuzzy_transpositions" field in the match bool prefix query.
//
// This method configures the match bool prefix query to allow or disallow transpositions (e.g., swapping two adjacent characters)
// when performing fuzzy matching. Setting this field to true enables transpositions, which can increase the match rate
// for terms with common typos or character swaps.
//
// Example usage:
//
//	m := es.MatchBoolPrefix("title", "es-query-builder").FuzzyTranspositions(true)
//	// m now has a "fuzzy_transpositions" field set to true in the match bool prefix query object.
//
// Parameters:
//   - fuzzyTranspositions: A boolean value indicating whether transpositions should be allowed in fuzzy matching.
//
// Returns:
//
//	The updated es.matchBoolPrefix object with the "fuzzy_transpositions" field set to the specified value.
func (m matchBoolPrefixType) FuzzyTranspositions(fuzzyTranspositions bool) matchBoolPrefixType {
	return m.putInTheField("fuzzy_transpositions", fuzzyTranspositions)
}

// MaxExpansions sets the "max_expansions" field in the match bool prefix query.
//
// This method configures the match bool prefix query to limit the maximum number of terms that can be expanded
// for multi-term queries, such as those involving fuzzy matching. Higher values allow more terms to
// be considered, but may impact performance.
//
// Example usage:
//
//	m := es.MatchBoolPrefix("title", "es-query-builder").MaxExpansions(50)
//	// m now has a "max_expansions" field set to 50 in the match bool prefix query object.
//
// Parameters:
//   - maxExpansions: An integer representing the maximum number of term expansions to be allowed in the match bool prefix query.
//
// Returns:
//
//	The updated es.matchBoolPrefix object with the "max_expansions" field set to the specified value.
func (m matchBoolPrefixType) MaxExpansions(maxExpansions int) matchBoolPrefixType {
	return m.putInTheField("max_expansions", maxExpansions)
}

// PrefixLength sets the "prefix_length" field in the match bool prefix query.
//
// This method configures the match bool prefix query to specify a minimum prefix length for fuzzy matching,
// which defines the number of initial characters in a term that must match exactly before
// considering fuzziness. Increasing this value can improve performance by reducing the number
// of fuzzy matches, but may also limit the flexibility of the query.
//
// Example usage:
//
//	m := es.MatchBoolPrefix("title", "es-query-builder").PrefixLength(2)
//	// m now has a "prefix_length" field set to 2 in the match bool prefix query object.
//
// Parameters:
//   - prefixLength: An integer representing the number of initial characters that must match exactly in fuzzy matching.
//
// Returns:
//
//	The updated es.matchBoolPrefix object with the "prefix_length" field set to the specified value.
func (m matchBoolPrefixType) PrefixLength(prefixLength int) matchBoolPrefixType {
	return m.putInTheField("prefix_length", prefixLength)
}

func (m matchBoolPrefixType) putInTheField(key string, value any) matchBoolPrefixType {
	putInTheNestedField(Object(m), "match_bool_prefix", key, value)
	return m
}
