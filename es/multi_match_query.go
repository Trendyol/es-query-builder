package es

import (
	Operator "github.com/Trendyol/es-query-builder/es/enums/operator"
	TextQueryType "github.com/Trendyol/es-query-builder/es/enums/text-query-type"
	ZeroTermsQuery "github.com/Trendyol/es-query-builder/es/enums/zero-terms-query"
)

type multiMatchType Object

// MultiMatch creates a new es.multiMatchType object with the specified query.
//
// This function initializes an es.multiMatchType object for a multi-match query,
// where the "query" is the value to search for across multiple fields. Multi-match
// queries are used when the search needs to target multiple fields simultaneously.
//
// Example usage:
//
//	m := es.MultiMatch("es-query-builder")
//	// m now contains an es.multiMatchType object that matches the query
//	// "es-query-builder" across multiple fields.
//
// Parameters:
//   - query: The value to be matched across multiple fields. The type is generic.
//
// Returns:
//
//	An es.multiMatchType object containing the specified multi-match query.
func MultiMatch[T any](query T) multiMatchType {
	return multiMatchType{
		"multi_match": Object{
			"query": query,
		},
	}
}

// Analyzer sets the "analyzer" field in the multi-match query.
//
// This method allows you to specify a custom analyzer to be used for the multi-match query.
// Analyzers are responsible for breaking text into tokens and can influence how terms
// are matched in the query.
//
// Example usage:
//
//	m := es.MultiMatch("es-query-builder").Analyzer("custom_analyzer")
//	// m now has an "analyzer" field set to "custom_analyzer" in the multi-match query object.
//
// Parameters:
//   - value: A string representing the name of the analyzer to use for the query.
//
// Returns:
//
//	The updated es.multiMatchType object with the "analyzer" field set to the specified value.
func (m multiMatchType) Analyzer(value string) multiMatchType {
	return m.putInTheField("analyzer", value)
}

// AutoGenerateSynonymsPhraseQuery sets the "auto_generate_synonyms_phrase_query" field
// in the multi-match query.
//
// This method enables or disables the automatic generation of phrase queries for synonyms.
// When enabled, the query can match documents containing synonyms for the search terms.
//
// Example usage:
//
//	m := es.MultiMatch("es-query-builder").AutoGenerateSynonymsPhraseQuery(true)
//	// m now has an "auto_generate_synonyms_phrase_query" field set to true in the multi-match query object.
//
// Parameters:
//   - autoGenerateSynonymsPhraseQuery: A boolean value indicating whether to enable
//     or disable automatic synonym phrase query generation.
//
// Returns:
//
//	The updated es.multiMatchType object with the "auto_generate_synonyms_phrase_query"
//	field set to the specified value.
func (m multiMatchType) AutoGenerateSynonymsPhraseQuery(autoGenerateSynonymsPhraseQuery bool) multiMatchType {
	return m.putInTheField("auto_generate_synonyms_phrase_query", autoGenerateSynonymsPhraseQuery)
}

// Boost sets the "boost" field in the multi-match query.
//
// This method allows you to specify a boost factor for the multi-match query,
// which influences the relevance scoring of documents that match the query.
// A higher boost value increases the importance of this query in the overall
// scoring.
//
// Example usage:
//
//	m := es.MultiMatch("es-query-builder").Boost(2.0)
//	// m now has a "boost" field set to 2.0 in the multi-match query object.
//
// Parameters:
//   - boost: A float64 value representing the boost factor to be applied to the multi-match query.
//
// Returns:
//
//	The updated es.multiMatchType object with the "boost" field set to the specified value.
func (m multiMatchType) Boost(boost float64) multiMatchType {
	return m.putInTheField("boost", boost)
}

// CutoffFrequency sets the "cutoff_frequency" field in the multi-match query.
//
// This method specifies the cutoff frequency for low-frequency terms in the query.
// Terms with a document frequency below this value are ignored, which can be useful
// for reducing the impact of noise in the query.
//
// Example usage:
//
//	m := es.MultiMatch("es-query-builder").CutoffFrequency(0.01)
//	// m now has a "cutoff_frequency" field set to 0.01 in the multi-match query object.
//
// Parameters:
//   - cutoffFrequency: A float64 value representing the cutoff frequency for low-frequency terms.
//
// Returns:
//
//	The updated es.multiMatchType object with the "cutoff_frequency" field set to the specified value.
func (m multiMatchType) CutoffFrequency(cutoffFrequency float64) multiMatchType {
	return m.putInTheField("cutoff_frequency", cutoffFrequency)
}

// Fields sets the "fields" field in the multi-match query.
//
// This method specifies the fields to be targeted by the multi-match query. It allows
// you to search across multiple fields simultaneously, providing flexibility in query
// construction.
//
// Example usage:
//
//	m := es.MultiMatch("es-query-builder").Fields("title", "description")
//	// m now has a "fields" field set to ["title", "description"] in the multi-match query object.
//
// Parameters:
//   - fields: A variadic list of strings representing the fields to include in the multi-match query.
//
// Returns:
//
//	The updated es.multiMatchType object with the "fields" field set to the specified values.
func (m multiMatchType) Fields(fields ...string) multiMatchType {
	return m.putInTheField("fields", fields)
}

// Fuzziness sets the "fuzziness" field in the multi-match query.
//
// This method specifies the fuzziness level for the query, which determines how many
// changes (insertions, deletions, or substitutions) are allowed for terms to match.
// It supports integer values, "AUTO", or other custom fuzziness settings.
//
// Example usage:
//
//	m := es.MultiMatch("es-query-builder").Fuzziness("AUTO")
//	// m now has a "fuzziness" field set to "AUTO" in the multi-match query object.
//
// Parameters:
//   - fuzziness: A value of any type representing the fuzziness level (e.g., an integer, "AUTO").
//
// Returns:
//
//	The updated es.multiMatchType object with the "fuzziness" field set to the specified value.
func (m multiMatchType) Fuzziness(fuzziness any) multiMatchType {
	return m.putInTheField("fuzziness", fuzziness)
}

// FuzzyRewrite sets the "fuzzy_rewrite" field in the multi-match query.
//
// This method specifies how the query should rewrite fuzzy terms. The value can be
// any valid Elasticsearch rewrite option, such as "constant_score", "scoring_boolean",
// or "top_terms_N" (e.g., "top_terms_50").
//
// Example usage:
//
//	m := es.MultiMatch("es-query-builder").FuzzyRewrite("constant_score")
//	// m now has a "fuzzy_rewrite" field set to "constant_score" in the multi-match query object.
//
// Parameters:
//   - fuzzyRewrite: A string representing the rewrite method for fuzzy terms.
//
// Returns:
//
//	The updated es.multiMatchType object with the "fuzzy_rewrite" field set to the specified value.
func (m multiMatchType) FuzzyRewrite(fuzzyRewrite string) multiMatchType {
	return m.putInTheField("fuzzy_rewrite", fuzzyRewrite)
}

// FuzzyTranspositions sets the "fuzzy_transpositions" field in the multi-match query.
//
// This method enables or disables transpositions in fuzzy matching. When enabled, the query
// considers transpositions (e.g., swapping adjacent characters) as a single edit operation.
//
// Example usage:
//
//	m := es.MultiMatch("es-query-builder").FuzzyTranspositions(true)
//	// m now has a "fuzzy_transpositions" field set to true in the multi-match query object.
//
// Parameters:
//   - fuzzyTranspositions: A boolean value indicating whether transpositions are allowed
//     in fuzzy matching.
//
// Returns:
//
//	The updated es.multiMatchType object with the "fuzzy_transpositions" field set to the specified value.
func (m multiMatchType) FuzzyTranspositions(fuzzyTranspositions bool) multiMatchType {
	return m.putInTheField("fuzzy_transpositions", fuzzyTranspositions)
}

// Lenient sets the "lenient" field in the multi-match query.
//
// This method determines whether the query should ignore exceptions caused by data type
// mismatches, such as querying a text field with a numeric value. When set to true, such
// errors are ignored, and the query execution continues.
//
// Example usage:
//
//	m := es.MultiMatch("es-query-builder").Lenient(true)
//	// m now has a "lenient" field set to true in the multi-match query object.
//
// Parameters:
//   - lenient: A boolean value indicating whether to enable lenient mode for the query.
//
// Returns:
//
//	The updated es.multiMatchType object with the "lenient" field set to the specified value.
func (m multiMatchType) Lenient(lenient bool) multiMatchType {
	return m.putInTheField("lenient", lenient)
}

// MaxExpansions sets the "max_expansions" field in the multi-match query.
//
// This method specifies the maximum number of terms that can be expanded for
// fuzzy matching. A higher value increases recall but may impact performance.
//
// Example usage:
//
//	m := es.MultiMatch("es-query-builder").MaxExpansions(50)
//	// m now has a "max_expansions" field set to 50 in the multi-match query object.
//
// Parameters:
//   - maxExpansions: An integer value representing the maximum number of terms to expand.
//
// Returns:
//
//	The updated es.multiMatchType object with the "max_expansions" field set to the specified value.
func (m multiMatchType) MaxExpansions(maxExpansions int) multiMatchType {
	return m.putInTheField("max_expansions", maxExpansions)
}

// MinimumShouldMatch sets the "minimum_should_match" field in the multi-match query.
//
// This method defines the minimum number of "should" clauses that must match for a document
// to be included in the results. The value can be an absolute number, a percentage, or a
// complex expression (e.g., "2", "75%", or "3<75%").
//
// Example usage:
//
//	m := es.MultiMatch("es-query-builder").MinimumShouldMatch("75%")
//	// m now has a "minimum_should_match" field set to "75%" in the multi-match query object.
//
// Parameters:
//   - minimumShouldMatch: A value of any type specifying the minimum number of "should" clauses
//     required to match.
//
// Returns:
//
//	The updated es.multiMatchType object with the "minimum_should_match" field set to the specified value.
func (m multiMatchType) MinimumShouldMatch(minimumShouldMatch any) multiMatchType {
	return m.putInTheField("minimum_should_match", minimumShouldMatch)
}

// Operator sets the "operator" field in the multi-match query.
//
// This method specifies how the clauses in the query should be combined. The operator can
// be "AND" or "OR", where "AND" requires all terms to match, and "OR" allows any term to match.
//
// Example usage:
//
//	m := es.MultiMatch("es-query-builder").Operator(Operator.AND)
//	// m now has an "operator" field set to "AND" in the multi-match query object.
//
// Parameters:
//   - operator: An Operator value (e.g., Operator.AND or Operator.OR) defining how query clauses
//     should be combined.
//
// Returns:
//
//	The updated es.multiMatchType object with the "operator" field set to the specified value.
func (m multiMatchType) Operator(operator Operator.Operator) multiMatchType {
	return m.putInTheField("operator", operator)
}

// PrefixLength sets the "prefix_length" field in the multi-match query.
//
// This method specifies the number of initial characters in a term that must match exactly
// when performing a fuzzy search. Increasing the prefix length can reduce the number of
// terms considered for matching, improving performance.
//
// Example usage:
//
//	m := es.MultiMatch("es-query-builder").PrefixLength(3)
//	// m now has a "prefix_length" field set to 3 in the multi-match query object.
//
// Parameters:
//   - prefixLength: An integer value representing the number of characters required to match
//     at the beginning of a term.
//
// Returns:
//
//	The updated es.multiMatchType object with the "prefix_length" field set to the specified value.
func (m multiMatchType) PrefixLength(prefixLength int) multiMatchType {
	return m.putInTheField("prefix_length", prefixLength)
}

// Slop sets the "slop" field in the multi-match query.
//
// This method specifies the allowable distance (in terms of word positions) between terms
// for a phrase match. A higher slop value increases the flexibility of matching phrases
// but reduces precision.
//
// Example usage:
//
//	m := es.MultiMatch("es-query-builder").Slop(2)
//	// m now has a "slop" field set to 2 in the multi-match query object.
//
// Parameters:
//   - slop: An integer value representing the maximum allowable distance between matched terms
//     in a phrase query.
//
// Returns:
//
//	The updated es.multiMatchType object with the "slop" field set to the specified value.
func (m multiMatchType) Slop(slop int) multiMatchType {
	return m.putInTheField("slop", slop)
}

// TieBreaker sets the "tie_breaker" field in the multi-match query.
//
// This method specifies the tie breaker value for multi-match queries using the "best_fields"
// or "most_fields" type. It determines how much the score of non-highest-scoring fields contributes
// to the overall score. A value of 0.0 ignores all non-highest scores, while a value of 1.0 adds them fully.
//
// Example usage:
//
//	m := es.MultiMatch("es-query-builder").TieBreaker(0.3)
//	// m now has a "tie_breaker" field set to 0.3 in the multi-match query object.
//
// Parameters:
//   - tieBreaker: A float64 value representing the tie breaker coefficient.
//
// Returns:
//
//	The updated es.multiMatchType object with the "tie_breaker" field set to the specified value.
func (m multiMatchType) TieBreaker(tieBreaker float64) multiMatchType {
	return m.putInTheField("tie_breaker", tieBreaker)
}

// Type sets the "type" field in the multi-match query.
//
// This method specifies the type of text query to use. The value can be one of several
// predefined query types, such as "best_fields", "most_fields", "phrase", or "phrase_prefix".
// Each type determines how the query processes and matches terms.
//
// Example usage:
//
//	m := es.MultiMatch("es-query-builder").Type(TextQueryType.Phrase)
//	// m now has a "type" field set to "phrase" in the multi-match query object.
//
// Parameters:
//   - textQueryType: A value of type TextQueryType, representing the desired text query type.
//
// Returns:
//
//	The updated es.multiMatchType object with the "type" field set to the specified value.
func (m multiMatchType) Type(textQueryType TextQueryType.TextQueryType) multiMatchType {
	return m.putInTheField("type", textQueryType)
}

// ZeroTermsQuery sets the "zero_terms_query" field in the multi-match query.
//
// This method specifies how the query should behave when the input terms resolve to an empty result set.
// Possible values are "none" (return no matches) or "all" (match all documents).
//
// Example usage:
//
//	m := es.MultiMatch("es-query-builder").ZeroTermsQuery(ZeroTermsQuery.All)
//	// m now has a "zero_terms_query" field set to "all" in the multi-match query object.
//
// Parameters:
//   - zeroTermsQuery: A value of type ZeroTermsQuery, representing the behavior for queries
//     with zero terms.
//
// Returns:
//
//	The updated es.multiMatchType object with the "zero_terms_query" field set to the specified value.
func (m multiMatchType) ZeroTermsQuery(zeroTermsQuery ZeroTermsQuery.ZeroTermsQuery) multiMatchType {
	return m.putInTheField("zero_terms_query", zeroTermsQuery)
}

func (m multiMatchType) putInTheField(key string, value any) multiMatchType {
	if multiMatch, ok := m["multi_match"].(Object); ok {
		multiMatch[key] = value
	}
	return m
}
