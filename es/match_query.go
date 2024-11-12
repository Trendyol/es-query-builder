package es

import (
	Operator "github.com/Trendyol/es-query-builder/es/enums/operator"
	ZeroTermsQuery "github.com/Trendyol/es-query-builder/es/enums/zero-terms-query"
)

type matchType Object

// Match creates a new matchType object with the specified field and query.
//
// This function initializes a matchType object for a match query, where the key
// is the field name and query is the value to search for in that field. This is used
// to construct queries that match the specified value in the given field.
//
// Example usage:
//
//	m := es.Match("title", "es-query-builder")
//	// m now contains a matchType object that matches the query "es-query-builder" in the "title" field.
//
// Parameters:
//   - key: A string representing the field name for the match query.
//   - query: The value to be matched in the specified field. The type is generic.
//
// Returns:
//
//	A matchType object containing the specified match query.
func Match[T any](key string, query T) matchType {
	return matchType{
		"match": Object{
			key: Object{
				"query": query,
			},
		},
	}
}

func (m matchType) putInTheField(key string, value any) matchType {
	if match, ok := m["match"].(Object); ok {
		for _, fieldObj := range match {
			if fieldObject, foOk := fieldObj.(Object); foOk {
				fieldObject[key] = value
				break
			}
		}
	}
	return m
}

// Operator sets the "operator" field in the match query.
//
// This method configures the match query to use a specified operator (e.g., "AND" or "OR")
// for the matching process.
//
// Example usage:
//
//	m := es.Match("title", "es-query-builder").Operator("AND")
//	// m now has an "operator" field set to "AND" in the match query object.
//
// Parameters:
//   - operator: An Operator.Operator value representing the logical operator to be used in the match query.
//
// Returns:
//
//	The updated matchType object with the "operator" field set to the specified value.
func (m matchType) Operator(operator Operator.Operator) matchType {
	return m.putInTheField("operator", operator)
}

// Boost sets the "boost" field in the match query.
//
// This method configures the match query to use a specified boost factor, which influences
// the relevance scoring of the matched documents.
//
// Example usage:
//
//	m := es.Match("title", "es-query-builder").Boost(1.5)
//	// m now has a "boost" field set to 1.5 in the match query object.
//
// Parameters:
//   - boost: A float64 value representing the boost factor to be applied to the match query.
//
// Returns:
//
//	The updated matchType object with the "boost" field set to the specified value.
func (m matchType) Boost(boost float64) matchType {
	return m.putInTheField("boost", boost)
}

// CutoffFrequency sets the "cutoff_frequency" field in the match query.
//
// This method configures the match query to use a specified cutoff frequency, which is useful
// for controlling how often terms should appear in the document for it to be considered a match.
// A lower cutoff frequency increases precision, while a higher one allows more terms to be matched.
//
// Example usage:
//
//	m := es.Match("title", "es-query-builder").CutoffFrequency(0.001)
//	// m now has a "cutoff_frequency" field set to 0.001 in the match query object.
//
// Parameters:
//   - cutoffFrequency: A float64 value representing the cutoff frequency threshold to be used in the match query.
//
// Returns:
//
//	The updated matchType object with the "cutoff_frequency" field set to the specified value.
func (m matchType) CutoffFrequency(cutoffFrequency float64) matchType {
	return m.putInTheField("cutoff_frequency", cutoffFrequency)
}

// Fuzziness sets the "fuzziness" field in the match query.
//
// This method configures the match query to use a specified fuzziness level, which determines
// the allowed edit distance (e.g., number of character changes) for a term to be considered a match.
// Common values include "AUTO", or integers representing the number of edits (e.g., 1 or 2).
//
// Example usage:
//
//	m := es.Match("title", "es-query-builder").Fuzziness("AUTO")
//	// m now has a "fuzziness" field set to "AUTO" in the match query object.
//
// Parameters:
//   - fuzziness: A value of any type (typically a string or integer) representing the fuzziness level to be applied to the match query.
//
// Returns:
//
//	The updated matchType object with the "fuzziness" field set to the specified value.
func (m matchType) Fuzziness(fuzziness any) matchType {
	return m.putInTheField("fuzziness", fuzziness)
}

// FuzzyRewrite sets the "fuzzy_rewrite" field in the match query.
//
// This method configures the match query to use a specified fuzzy rewrite method,
// which controls how multi-term queries are rewritten. Common values include "constant_score",
// "scoring_boolean", and other rewrite options that influence the scoring and performance of
// fuzzy matching.
//
// Example usage:
//
//	m := es.Match("title", "es-query-builder").FuzzyRewrite("constant_score")
//	// m now has a "fuzzy_rewrite" field set to "constant_score" in the match query object.
//
// Parameters:
//   - fuzzyRewrite: A string value representing the rewrite method to be used for fuzzy matching in the match query.
//
// Returns:
//
//	The updated matchType object with the "fuzzy_rewrite" field set to the specified value.
func (m matchType) FuzzyRewrite(fuzzyRewrite string) matchType {
	return m.putInTheField("fuzzy_rewrite", fuzzyRewrite)
}

// FuzzyTranspositions sets the "fuzzy_transpositions" field in the match query.
//
// This method configures the match query to allow or disallow transpositions (e.g., swapping two adjacent characters)
// when performing fuzzy matching. Setting this field to true enables transpositions, which can increase the match rate
// for terms with common typos or character swaps.
//
// Example usage:
//
//	m := es.Match("title", "es-query-builder").FuzzyTranspositions(true)
//	// m now has a "fuzzy_transpositions" field set to true in the match query object.
//
// Parameters:
//   - fuzzyTranspositions: A boolean value indicating whether transpositions should be allowed in fuzzy matching.
//
// Returns:
//
//	The updated matchType object with the "fuzzy_transpositions" field set to the specified value.
func (m matchType) FuzzyTranspositions(fuzzyTranspositions bool) matchType {
	return m.putInTheField("fuzzy_transpositions", fuzzyTranspositions)
}

// Lenient sets the "lenient" field in the match query.
//
// This method configures the match query to use lenient parsing, allowing it to skip errors
// for data type mismatches. When set to true, documents that contain mismatched data types
// (e.g., text in a numeric field) will not cause errors and will be ignored instead.
//
// Example usage:
//
//	m := es.Match("title", "es-query-builder").Lenient(true)
//	// m now has a "lenient" field set to true in the match query object.
//
// Parameters:
//   - lenient: A boolean value indicating whether lenient parsing should be enabled.
//
// Returns:
//
//	The updated matchType object with the "lenient" field set to the specified value.
func (m matchType) Lenient(lenient bool) matchType {
	return m.putInTheField("lenient", lenient)
}

// MaxExpansions sets the "max_expansions" field in the match query.
//
// This method configures the match query to limit the maximum number of terms that can be expanded
// for multi-term queries, such as those involving fuzzy matching. Higher values allow more terms to
// be considered, but may impact performance.
//
// Example usage:
//
//	m := es.Match("title", "es-query-builder").MaxExpansions(50)
//	// m now has a "max_expansions" field set to 50 in the match query object.
//
// Parameters:
//   - maxExpansions: An integer representing the maximum number of term expansions to be allowed in the match query.
//
// Returns:
//
//	The updated matchType object with the "max_expansions" field set to the specified value.
func (m matchType) MaxExpansions(maxExpansions int) matchType {
	return m.putInTheField("max_expansions", maxExpansions)
}

// PrefixLength sets the "prefix_length" field in the match query.
//
// This method configures the match query to specify a minimum prefix length for fuzzy matching,
// which defines the number of initial characters in a term that must match exactly before
// considering fuzziness. Increasing this value can improve performance by reducing the number
// of fuzzy matches, but may also limit the flexibility of the query.
//
// Example usage:
//
//	m := es.Match("title", "es-query-builder").PrefixLength(2)
//	// m now has a "prefix_length" field set to 2 in the match query object.
//
// Parameters:
//   - prefixLength: An integer representing the number of initial characters that must match exactly in fuzzy matching.
//
// Returns:
//
//	The updated matchType object with the "prefix_length" field set to the specified value.
func (m matchType) PrefixLength(prefixLength int) matchType {
	return m.putInTheField("prefix_length", prefixLength)
}

// AutoGenerateSynonymsPhraseQuery sets the "auto_generate_synonyms_phrase_query" field in the match query.
//
// This method enables or disables automatic generation of phrase queries for synonyms in the match query.
// When enabled, the query engine automatically expands the search to include synonym phrases, which can
// enhance results by including terms related to the search term. However, enabling this may increase
// the complexity of the query and affect performance, especially with a large synonym dictionary.
//
// Example usage:
//
//	m := es.Match("title", "es-query-builder").AutoGenerateSynonymsPhraseQuery(true)
//	// m now has the "auto_generate_synonyms_phrase_query" field set to true in the match query object.
//
// Parameters:
//   - autoGenerateSynonymsPhraseQuery: A boolean indicating whether to automatically generate phrase queries for synonyms.
//
// Returns:
//
//	The updated matchType object with the "auto_generate_synonyms_phrase_query" field set to the specified value.
func (m matchType) AutoGenerateSynonymsPhraseQuery(autoGenerateSynonymsPhraseQuery bool) matchType {
	return m.putInTheField("auto_generate_synonyms_phrase_query", autoGenerateSynonymsPhraseQuery)
}

// ZeroTermsQuery sets the "zero_terms_query" field in the match query.
//
// This method configures the behavior of the match query when no terms remain after analysis
// (for example, if all terms are stop words). The specified zero_terms_query value determines
// how to handle this scenario, with options like "all" to match all documents or "none" to
// match none.
//
// Example usage:
//
//	m := es.Match("title", "es-query-builder").ZeroTermsQuery(zerotermsquery.All)
//	// m now has a "zero_terms_query" field set to "all" in the match query object.
//
// Parameters:
//   - zeroTermsQuery: A zerotermsquery.ZeroTermsQuery value that specifies the behavior for zero-term queries.
//
// Returns:
//
//	The updated matchType object with the "zero_terms_query" field set to the specified value.
func (m matchType) ZeroTermsQuery(zeroTermsQuery ZeroTermsQuery.ZeroTermsQuery) matchType {
	return m.putInTheField("zero_terms_query", zeroTermsQuery)
}
