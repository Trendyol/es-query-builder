package es

import (
	Operator "github.com/Trendyol/es-query-builder/es/enums/operator"
)

type simpleQueryStringType Object

// SimpleQueryString creates a new simpleQueryStringType object with the specified query string.
//
// This function initializes a simpleQueryStringType object with a simple query string, which
// is typically used to perform simple text search queries in Elasticsearch. The query string
// can contain multiple terms and operators, allowing for basic search expressions.
//
// Example usage:
//
//	q := es.SimpleQueryString("Foo + Bar")
//	// q now contains a simpleQueryStringType object with a simple query string query.
//
// Parameters:
//   - query: The query string to be used in the search. The type is generic and can be
//     any type that represents a query string.
//
// Returns:
//
//	A simpleQueryStringType object containing the specified query string.
func SimpleQueryString[T any](query T) simpleQueryStringType {
	return simpleQueryStringType{
		"simple_query_string": Object{
			"query": query,
		},
	}
}

// Fields sets the fields to be searched within the simpleQueryStringType object.
//
// This method specifies a list of fields that the query string should search.
// If multiple fields are provided, the query will search across all of them,
// allowing for more flexible and comprehensive search queries.
//
// Example usage:
//
//	q := es.SimpleQueryString("Foo Bar").Fields([]string{"title", "content"})
//	// q now searches within the "title" and "content" fields.
//
// Parameters:
//   - value: A slice of strings representing the field names to be searched.
//
// Returns:
//
// The updated simpleQueryStringType object with the "fields" option set.
func (q simpleQueryStringType) Fields(value []string) simpleQueryStringType {
	return q.putInTheField("fields", value)
}

// Analyzer sets the analyzer to be used for the simpleQueryStringType object.
//
// This method specifies the analyzer that should be applied to the query string.
// Analyzers are used to process the text, such as tokenizing and normalizing it,
// allowing for more refined and accurate search queries based on the specified analyzer.
//
// Example usage:
//
//	q := es.SimpleQueryString("Foo Bar").Analyzer("standard")
//	// q now uses the "standard" analyzer for processing the query string.
//
// Parameters:
//   - value: A string representing the name of the analyzer to be used.
//
// Returns:
//
// The updated simpleQueryStringType object with the "analyzer" set.
func (q simpleQueryStringType) Analyzer(value string) simpleQueryStringType {
	return q.putInTheField("analyzer", value)
}

// DefaultOperator sets the default operator for the simpleQueryStringType object.
//
// This method specifies the default operator to be used between terms in the query string
// when no explicit operator is provided. The default operator can be operator.And or operator.Or,
// determining whether all terms (and) or any term (or) must be matched in the search results.
//
// Example usage:
//
//	q := es.SimpleQueryString("Foo Bar").
//	DefaultOperator(operator.Or)
//
// q now uses "or" as the default operator, meaning any term can match in the query.
//
// Parameters:
//   - operator: A operator.Operator representing the default operator to be used ("and" or "or").
//
// Returns:
//
// The updated simpleQueryStringType object with the "default_operator" set.
func (q simpleQueryStringType) DefaultOperator(operator Operator.Operator) simpleQueryStringType {
	return q.putInTheField("default_operator", operator)
}

// MinimumShouldMatch sets the minimum number of clauses that must match for the simpleQueryStringType object.
//
// This method specifies the minimum number of clauses that must match in order
// for a document to be considered a match. This can be expressed as an absolute number or a percentage,
// allowing for fine-tuned control over query matching behavior.
//
// Example usage:
//
//	q := es.SimpleQueryString("Foo Bar Baz").MinimumShouldMatch("2")
//	// q now requires that at least 2 of the terms match for a document to be considered a match.
//
// Parameters:
//   - value: A string representing the minimum number or percentage of clauses that must match.
//
// Returns:
//
// The updated simpleQueryStringType object with the "minimum_should_match" option set.
func (q simpleQueryStringType) MinimumShouldMatch(value string) simpleQueryStringType {
	return q.putInTheField("minimum_should_match", value)
}

// FuzzyMaxExpansions sets the maximum number of expansions for fuzzy matching in the simpleQueryStringType object.
//
// This method specifies the maximum number of terms that the query will expand to
// when performing fuzzy matching. This setting controls the number of variations
// of the search term that will be considered in the query, affecting both performance
// and accuracy of fuzzy searches.
//
// Example usage:
//
//	q := es.SimpleQueryString("Foo~").FuzzyMaxExpansions(50)
//	// q now allows up to 50 expansions for fuzzy matching.
//
// Parameters:
//   - value: An integer representing the maximum number of term expansions for fuzzy matching.
//
// Returns:
//
// The updated simpleQueryStringType object with the "fuzzy_max_expansions" option set.
func (q simpleQueryStringType) FuzzyMaxExpansions(value int64) simpleQueryStringType {
	return q.putInTheField("fuzzy_max_expansions", value)
}

// FuzzyPrefixLength sets the prefix length for fuzzy matching in the simpleQueryStringType object.
//
// This method specifies the length of the initial characters that must match exactly
// before applying any fuzziness in the query. Increasing the prefix length can improve
// performance by reducing the number of potential matches, while still allowing for
// approximate matching beyond the prefix.
//
// Example usage:
//
//	q := es.SimpleQueryString("Foo~").FuzzyPrefixLength(2)
//	// q now requires the first 2 characters to match exactly before applying fuzziness.
//
// Parameters:
//   - value: An integer representing the number of initial characters that must match exactly.
//
// Returns:
//
// The updated simpleQueryStringType object with the "fuzzy_prefix_length" option set.
func (q simpleQueryStringType) FuzzyPrefixLength(value int64) simpleQueryStringType {
	return q.putInTheField("fuzzy_prefix_length", value)
}

// FuzzyTranspositions sets the option to allow transpositions in fuzzy matching for the simpleQueryStringType object.
//
// This method enables or disables the allowance of transpositions (swapping of adjacent characters)
// in fuzzy matching. When set to true, terms that are similar but have transposed characters
// (e.g., "ab" vs. "ba") will still be considered a match, increasing the flexibility of the search.
//
// Example usage:
//
//	q := es.SimpleQueryString("Foo~").FuzzyTranspositions(true)
//	// q now allows transpositions in fuzzy matching.
//
// Parameters:
//   - value: A boolean indicating whether transpositions are allowed (true) or not (false).
//
// Returns:
//
// The updated simpleQueryStringType object with the "fuzzy_transpositions" option set.
func (q simpleQueryStringType) FuzzyTranspositions(value bool) simpleQueryStringType {
	return q.putInTheField("fuzzy_transpositions", value)
}

// AnalyzeWildcard sets the option to analyze wildcard terms in the simpleQueryStringType object.
//
// This method determines whether wildcard terms in the query string should be analyzed.
// When set to true, wildcard terms (* and ?) will be analyzed by the analyzer defined
// for the field, allowing for more accurate searches when using wildcards.
//
// Example usage:
//
//	q := es.SimpleQueryString("Fo*").AnalyzeWildcard(true)
//	// q now analyzes wildcard terms in the query string.
//
// Parameters:
//   - value: A boolean indicating whether wildcard terms should be analyzed (true) or not (false).
//
// Returns:
//
// The updated simpleQueryStringType object with the "analyze_wildcard" option set.
func (q simpleQueryStringType) AnalyzeWildcard(value bool) simpleQueryStringType {
	return q.putInTheField("analyze_wildcard", value)
}

// AutoGenerateSynonymsPhraseQuery sets the option to automatically generate phrase queries for synonyms
// in the simpleQueryStringType object.
//
// This method enables or disables the automatic generation of phrase queries for synonyms in the query string.
// When set to true, Elasticsearch will automatically create phrase queries for terms that have synonyms,
// which can improve search accuracy when working with synonym filters.
//
// Example usage:
//
//	q := es.SimpleQueryString("quick brown fox").AutoGenerateSynonymsPhraseQuery(true)
//	// q now automatically generates phrase queries for synonyms in the query string.
//
// Parameters:
//   - value: A boolean indicating whether to automatically generate phrase queries for synonyms (true) or not (false).
//
// Returns:
//
// The updated simpleQueryStringType object with the "auto_generate_synonyms_phrase_query" option set.
func (q simpleQueryStringType) AutoGenerateSynonymsPhraseQuery(value bool) simpleQueryStringType {
	return q.putInTheField("auto_generate_synonyms_phrase_query", value)
}

// Flags sets the flags for the simpleQueryStringType object.
//
// This method specifies which features of the simple_query_string query should be enabled.
// It allows fine-grained control over the query's behavior by enabling or disabling specific features.
//
// Example usage:
//
//	q := es.SimpleQueryString("Foo Bar").Flags("AND|OR|PREFIX")
//	// q now enables AND, OR, and PREFIX features for the query.
//
// Parameters:
//   - value: A string representing the enabled features, separated by '|'.
//     Possible values include: ALL, NONE, AND, OR, NOT, PREFIX, PHRASE, PRECEDENCE, ESCAPE, WHITESPACE, FUZZY, NEAR, SLOP.
//
// Returns:
//
// The updated simpleQueryStringType object with the "flags" option set.
func (q simpleQueryStringType) Flags(value string) simpleQueryStringType {
	return q.putInTheField("flags", value)
}

// Lenient sets the leniency option for the simpleQueryStringType object.
//
// This method determines whether the query should be lenient when encountering
// errors, such as analyzing incompatible fields. When set to true, the query will
// ignore such errors, allowing for more flexible and fault-tolerant searches,
// especially in cases where the data types may not match perfectly.
//
// Example usage:
//
//	q := es.SimpleQueryString("Foo Bar").Lenient(true)
//	// q is now lenient, allowing it to tolerate errors during the query.
//
// Parameters:
//   - value: A boolean indicating whether leniency is enabled (true) or disabled (false).
//
// Returns:
//
// The updated simpleQueryStringType object with the "lenient" option set.
func (q simpleQueryStringType) Lenient(value bool) simpleQueryStringType {
	return q.putInTheField("lenient", value)
}

// QuoteFieldSuffix sets the field suffix to be used for quoted text in the simpleQueryStringType object.
//
// This method specifies a suffix to be appended to the field names when analyzing quoted text in the query string.
// This is useful for applying different analyzers or field mappings to quoted phrases compared to unquoted terms.
//
// Example usage:
//
//	q := es.SimpleQueryString("Foo \"Bar Baz\"").QuoteFieldSuffix("_phrase")
//	// q now appends "_phrase" to the field names when processing quoted text in the query string.
//
// Parameters:
//   - value: A string representing the suffix to be appended to field names for quoted text.
//
// Returns:
//
// The updated simpleQueryStringType object with the "quote_field_suffix" option set.
func (q simpleQueryStringType) QuoteFieldSuffix(value string) simpleQueryStringType {
	return q.putInTheField("quote_field_suffix", value)
}

func (q simpleQueryStringType) putInTheField(key string, value any) simpleQueryStringType {
	for _, fieldObj := range q {
		if fieldObject, ok := fieldObj.(Object); ok {
			fieldObject[key] = value
			break
		}
	}
	return q
}
