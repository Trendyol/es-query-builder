package es

type queryStringType Object

// QueryString creates a new queryStringType object with the specified query string.
//
// This function initializes a queryStringType object with a query string, which
// is typically used to perform full-text search queries in Elasticsearch. The query string
// can contain multiple terms and operators, allowing for complex search expressions.
//
// Example usage:
//
//	q := QueryString("Foo AND Bar")
//	// q now contains a queryStringType object with a query string query.
//
// Parameters:
//   - query: The query string to be used in the search. The type is generic and can be
//     any type that represents a query string.
//
// Returns:
//
//	A queryStringType object containing the specified query string.
func QueryString[T any](query T) queryStringType {
	return queryStringType{
		"query_string": Object{
			"query": query,
		},
	}
}

// DefaultField sets the default field for the queryStringType object.
//
// This method specifies the default field to search within if no field is explicitly mentioned
// in the query string. It is useful when you want to perform a query across a single field
// by default when no field is specified in the query string.
//
// Example usage:
//
// b := es.QueryString("value").
// DefaultField("defaultField")
//
//	q := QueryString("Foo Bar").DefaultField("title")
//	// q now contains a queryStringType object where the default field for the query is "title".
//
// Parameters:
//   - value: A string representing the field name to be used as the default field in the query.
//
// Returns:
//
// The updated queryStringType object with the new "default_field".
func (q queryStringType) DefaultField(value string) queryStringType {
	q.putInTheField("default_field", value)
	return q
}

// AllowLeadingWildcard sets the option to allow leading wildcards in the queryStringType object.
//
// This method enables or disables the use of leading wildcards in the query string.
// When set to true, wildcard queries can begin with a wildcard character (* or ?),
// which can be useful for matching terms that share a common suffix.
//
// Example usage:
//
// b := es.QueryString("value").
// AllowLeadingWildcard(true)
//
//	q := QueryString("Fo* bar").AllowLeadingWildcard(true)
//	// q now allows leading wildcards in the query string.
//
// Parameters:
//   - value: A boolean indicating whether leading wildcards are allowed (true) or not (false).
//
// Returns:
//
// The updated queryStringType object with the "allow_leading_wildcard" option set.
func (q queryStringType) AllowLeadingWildcard(value bool) queryStringType {
	q.putInTheField("allow_leading_wildcard", value)
	return q
}

// AnalyzeWildcard sets the option to analyze wildcard terms in the queryStringType object.
//
// This method determines whether wildcard terms in the query string should be analyzed.
// When set to true, wildcard terms (* and ?) will be analyzed by the analyzer defined
// for the field, allowing for more accurate searches when using wildcards.
//
// Example usage:
//
// b := es.QueryString("value").
// AnalyzeWildcard(true)
//
//	q := QueryString("Fo* bar").AnalyzeWildcard(true)
//	// q now analyzes wildcard terms in the query string.
//
// Parameters:
//   - value: A boolean indicating whether wildcard terms should be analyzed (true) or not (false).
//
// Returns:
//
// The updated queryStringType object with the "analyze_wildcard" option set.
func (q queryStringType) AnalyzeWildcard(value bool) queryStringType {
	q.putInTheField("analyze_wildcard", value)
	return q
}

// Analyzer sets the analyzer to be used for the queryStringType object.
//
// This method specifies the analyzer that should be applied to the query string.
// Analyzers are used to process the text, such as tokenizing and normalizing it,
// allowing for more refined and accurate search queries based on the specified analyzer.
//
// Example usage:
//
// b := es.QueryString("value").
// Analyzer("custom_analyzer")
//
//	q := QueryString("Foo Bar").Analyzer("standard")
//	// q now uses the "standard" analyzer for processing the query string.
//
// Parameters:
//   - value: A string representing the name of the analyzer to be used.
//
// Returns:
//
// The updated queryStringType object with the "analyzer" set.
func (q queryStringType) Analyzer(value string) queryStringType {
	q.putInTheField("analyzer", value)
	return q
}

// AutoGenerateSynonymsPhraseQuery sets the option to automatically generate phrase queries for synonyms in the queryStringType object.
//
// This method enables or disables the automatic generation of phrase queries for synonyms in the query string.
// When set to true, Elasticsearch will automatically create phrase queries for terms that have synonyms,
// which can improve search accuracy when working with synonym filters.
//
// Example usage:
//
// b := es.QueryString("value").
// AutoGenerateSynonymsPhraseQuery(true)
//
//	q := QueryString("quick brown fox").AutoGenerateSynonymsPhraseQuery(true)
//	// q now automatically generates phrase queries for synonyms in the query string.
//
// Parameters:
//   - value: A boolean indicating whether to automatically generate phrase queries for synonyms (true) or not (false).
//
// Returns:
//
// The updated queryStringType object with the "auto_generate_synonyms_phrase_query" option set.
func (q queryStringType) AutoGenerateSynonymsPhraseQuery(value bool) queryStringType {
	q.putInTheField("auto_generate_synonyms_phrase_query", value)
	return q
}

// Boost sets the boost factor for the queryStringType object.
//
// This method specifies the boost value to increase or decrease the relevance of the query.
// A higher boost value increases the relevance score of the query, making it more likely
// to appear higher in the search results, while a lower value decreases its importance.
//
// Example usage:
//
// b := es.QueryString("value").
// Boost(2.0)
//
//	q := QueryString("Foo Bar").Boost(1.5)
//	// q now has a boost factor of 1.5, increasing its relevance in the search results.
//
// Parameters:
//   - value: A float64 representing the boost factor to be applied to the query.
//
// Returns:
//
// The updated queryStringType object with the "boost" value set.
func (q queryStringType) Boost(value float64) queryStringType {
	q.putInTheField("boost", value)
	return q
}

// DefaultOperator sets the default operator for the queryStringType object.
//
// This method specifies the default operator to be used between terms in the query string
// when no explicit operator is provided. The default operator can be "AND" or "OR",
// determining whether all terms (AND) or any term (OR) must be matched in the search results.
//
// Example usage:
//
// b := es.QueryString("value").
// DefaultOperator("AND")
//
//	q := QueryString("Foo Bar").DefaultOperator("OR")
//	// q now uses "OR" as the default operator, meaning any term can match in the query.
//
// Parameters:
//   - value: A string representing the default operator to be used ("AND" or "OR").
//
// Returns:
//
// The updated queryStringType object with the "default_operator" set.
func (q queryStringType) DefaultOperator(value string) queryStringType {
	q.putInTheField("default_operator", value)
	return q
}

// EnablePositionIncrements sets the option to enable or disable position increments in the queryStringType object.
//
// This method determines whether to account for position increments when analyzing the query string.
// When set to true, position increments are taken into account, which can improve the accuracy of
// phrase and proximity queries by considering gaps caused by stopwords or other factors.
//
// Example usage:
//
// b := es.QueryString("value").
// EnablePositionIncrements(true)
//
// Parameters:
//   - value: A boolean indicating whether to enable (true) or disable (false) position increments.
//
// Returns:
//
// The updated queryStringType object with the "enable_position_increments" option set.
func (q queryStringType) EnablePositionIncrements(value bool) queryStringType {
	q.putInTheField("enable_position_increments", value)
	return q
}

// Fields sets the fields to be searched within the queryStringType object.
//
// This method specifies a list of fields that the query string should search.
// If multiple fields are provided, the query will search across all of them,
// allowing for more flexible and comprehensive search queries.
//
// Example usage:
//
// b := es.QueryString("value").
// Fields([]string{"title", "description"})
//
//	q := QueryString("Foo Bar").Fields([]string{"title", "content"})
//	// q now searches within the "title" and "content" fields.
//
// Parameters:
//   - value: A slice of strings representing the field names to be searched.
//
// Returns:
//
// The updated queryStringType object with the "fields" option set.
func (q queryStringType) Fields(value []string) queryStringType {
	q.putInTheField("fields", value)
	return q
}

// Fuzziness sets the fuzziness level for the queryStringType object.
//
// This method specifies the fuzziness level for the query, allowing for approximate
// matching of terms. Fuzziness is particularly useful for handling misspellings or
// variations in search terms. The fuzziness value can be "AUTO", a number indicating
// the maximum allowed edits, or a specific fuzziness setting.
//
// Example usage:
//
// b := es.QueryString("value").
// Fuzziness("AUTO")
//
//	q := QueryString("Foo Bar").Fuzziness("2")
//	// q now uses a fuzziness level of "2" to allow for approximate matching.
//
// Parameters:
//   - value: A string representing the fuzziness level (e.g., "AUTO", "1", "2").
//
// Returns:
//
// The updated queryStringType object with the "fuzziness" option set.
func (q queryStringType) Fuzziness(value string) queryStringType {
	q.putInTheField("fuzziness", value)
	return q
}

// FuzzyMaxExpansions sets the maximum number of expansions for fuzzy matching in the queryStringType object.
//
// This method specifies the maximum number of terms that the query will expand to
// when performing fuzzy matching. This setting controls the number of variations
// of the search term that will be considered in the query, affecting both performance
// and accuracy of fuzzy searches.
//
// Example usage:
//
// b := es.QueryString("value").
// FuzzyMaxExpansions(50)
//
//	q := QueryString("Foo Bar").FuzzyMaxExpansions(100)
//	// q now allows up to 100 expansions for fuzzy matching.
//
// Parameters:
//   - value: An integer representing the maximum number of term expansions for fuzzy matching.
//
// Returns:
//
// The updated queryStringType object with the "fuzzy_max_expansions" option set.
func (q queryStringType) FuzzyMaxExpansions(value int64) queryStringType {
	q.putInTheField("fuzzy_max_expansions", value)
	return q
}

// FuzzyPrefixLength sets the prefix length for fuzzy matching in the queryStringType object.
//
// This method specifies the length of the initial characters that must match exactly
// before applying any fuzziness in the query. Increasing the prefix length can improve
// performance by reducing the number of potential matches, while still allowing for
// approximate matching beyond the prefix.
//
// Example usage:
//
// b := es.QueryString("value").
// FuzzyPrefixLength(2)
//
//	q := QueryString("Foo Bar").FuzzyPrefixLength(3)
//	// q now requires the first 3 characters to match exactly before applying fuzziness.
//
// Parameters:
//   - value: An integer representing the number of initial characters that must match exactly.
//
// Returns:
//
// The updated queryStringType object with the "fuzzy_prefix_length" option set.
func (q queryStringType) FuzzyPrefixLength(value int64) queryStringType {
	q.putInTheField("fuzzy_prefix_length", value)
	return q
}

// FuzzyTranspositions sets the option to allow transpositions in fuzzy matching for the queryStringType object.
//
// This method enables or disables the allowance of transpositions (swapping of adjacent characters)
// in fuzzy matching. When set to true, terms that are similar but have transposed characters
// (e.g., "ab" vs. "ba") will still be considered a match, increasing the flexibility of the search.
//
// Example usage:
//
// b := es.QueryString("value").
// FuzzyTranspositions(true)
//
//	q := QueryString("Foo Bar").FuzzyTranspositions(true)
//	// q now allows transpositions in fuzzy matching.
//
// Parameters:
//   - value: A boolean indicating whether transpositions are allowed (true) or not (false).
//
// Returns:
//
// The updated queryStringType object with the "fuzzy_transpositions" option set.
func (q queryStringType) FuzzyTranspositions(value bool) queryStringType {
	q.putInTheField("fuzzy_transpositions", value)
	return q
}

// Lenient sets the leniency option for the queryStringType object.
//
// This method determines whether the query should be lenient when encountering
// errors, such as analyzing incompatible fields. When set to true, the query will
// ignore such errors, allowing for more flexible and fault-tolerant searches,
// especially in cases where the data types may not match perfectly.
//
// Example usage:
//
// b := es.QueryString("value").
// Lenient(true)
//
//	q := QueryString("Foo Bar").Lenient(true)
//	// q is now lenient, allowing it to tolerate errors during the query.
//
// Parameters:
//   - value: A boolean indicating whether leniency is enabled (true) or disabled (false).
//
// Returns:
//
// The updated queryStringType object with the "lenient" option set.
func (q queryStringType) Lenient(value bool) queryStringType {
	q.putInTheField("lenient", value)
	return q
}

// MaxDeterminizedStates sets the maximum number of determinized states for the queryStringType object.
//
// This method specifies the maximum number of states that can be determinized when expanding
// wildcard, prefix, and other complex queries into a finite automaton. Limiting this number
// helps control the potential complexity and resource usage of the query, preventing excessive
// expansion that could impact performance.
//
// Example usage:
//
// b := es.QueryString("value").
// MaxDeterminizedStates(10000)
//
//	q := QueryString("Foo*").MaxDeterminizedStates(5000)
//	// q now limits the determinized states to 5000 to control query complexity.
//
// Parameters:
//   - value: An integer representing the maximum number of determinized states allowed.
//
// Returns:
//
// The updated queryStringType object with the "max_determinized_states" option set.
func (q queryStringType) MaxDeterminizedStates(value int64) queryStringType {
	q.putInTheField("max_determinized_states", value)
	return q
}

// MinimumShouldMatch sets the minimum number of "should" clauses that must match for the queryStringType object.
//
// This method specifies the minimum number of optional ("should") clauses that must match in order
// for a document to be considered a match. This can be expressed as an absolute number or a percentage,
// allowing for fine-tuned control over query matching behavior, especially in complex boolean queries.
//
// Example usage:
//
// b := es.QueryString("value").
// MinimumShouldMatch("2")
//
//	q := QueryString("Foo Bar").MinimumShouldMatch("2")
//	// q now requires that at least 2 of the "should" clauses match for a document to be considered a match.
//
// Parameters:
//   - value: A string representing the minimum number or percentage of "should" clauses that must match.
//
// Returns:
//
// The updated queryStringType object with the "minimum_should_match" option set.
func (q queryStringType) MinimumShouldMatch(value string) queryStringType {
	q.putInTheField("minimum_should_match", value)
	return q
}

// QuoteAnalyzer sets the analyzer to be used for quoted text in the queryStringType object.
//
// This method specifies the analyzer that should be applied to terms within quotes in the query string.
// When a query contains quoted text, this analyzer will be used to process that portion of the query,
// allowing for customized analysis of phrases or exact matches within quotes.
//
// Example usage:
//
// b := es.QueryString("value").
// QuoteAnalyzer("custom_phrase_analyzer")
//
//	q := QueryString("Foo Bar").QuoteAnalyzer("standard")
//	// q now uses the "standard" analyzer for quoted text in the query string.
//
// Parameters:
//   - value: A string representing the name of the analyzer to be used for quoted text.
//
// Returns:
//
// The updated queryStringType object with the "quote_analyzer" option set.
func (q queryStringType) QuoteAnalyzer(value string) queryStringType {
	q.putInTheField("quote_analyzer", value)
	return q
}

// PhraseSlop sets the slop factor for phrase queries in the queryStringType object.
//
// This method specifies the allowed number of positions (or "slop") that terms in a phrase query can be
// moved around while still being considered a match. A higher slop value allows for more flexibility
// in the order and proximity of terms, making it useful for handling variations in phrase structure.
//
// Example usage:
//
// b := es.QueryString("value").
// PhraseSlop(3)
//
//	q := QueryString("Foo Bar").PhraseSlop(2)
//	// q now allows a slop of 2 positions for the phrase match, accommodating slight variations in term order.
//
// Parameters:
//   - value: An integer representing the maximum number of position increments allowed for the phrase query.
//
// Returns:
//
// The updated queryStringType object with the "phrase_slop" option set.
func (q queryStringType) PhraseSlop(value int64) queryStringType {
	q.putInTheField("phrase_slop", value)
	return q
}

// QuoteFieldSuffix sets the field suffix to be used for quoted text in the queryStringType object.
//
// This method specifies a suffix to be appended to the field names when analyzing quoted text in the query string.
// This is useful for applying different analyzers or field mappings to quoted phrases compared to unquoted terms.
//
// Example usage:
//
// b := es.QueryString("value").
// QuoteFieldSuffix("_quoted")
//
//	q := QueryString("Foo Bar").QuoteFieldSuffix("_phrase")
//	// q now appends "_phrase" to the field names when processing quoted text in the query string.
//
// Parameters:
//   - value: A string representing the suffix to be appended to field names for quoted text.
//
// Returns:
//
// The updated queryStringType object with the "quote_field_suffix" option set.
func (q queryStringType) QuoteFieldSuffix(value string) queryStringType {
	q.putInTheField("quote_field_suffix", value)
	return q
}

// Rewrite sets the rewrite method for the queryStringType object.
//
// This method specifies the rewrite method to be used for rewriting the query string. Rewrite methods
// are used to transform the query into a more optimized form for execution, which can affect both
// performance and the accuracy of the search results. Common rewrite methods include "constant_score",
// "scoring_boolean", and others, depending on the specific use case. This parameter is for expert users only.
// Changing the value of this parameter can impact search performance and relevance.
//
// Example usage:
//
// b := es.QueryString("value").
// Rewrite("constant_score")
//
//	q := QueryString("Foo Bar").Rewrite("scoring_boolean")
//	// q now uses the "scoring_boolean" rewrite method for optimizing the query execution.
//
// Parameters:
//   - value: A string representing the rewrite method to be applied to the query string.
//
// Returns:
//
// The updated queryStringType object with the "rewrite" option set.
func (q queryStringType) Rewrite(value string) queryStringType {
	q.putInTheField("rewrite", value)
	return q
}

// TimeZone sets the time zone for date and time fields in the queryStringType object.
//
// This method specifies the time zone to be applied when parsing and interpreting date and time values
// in the query string. Setting the correct time zone ensures accurate date range queries and comparisons,
// especially when dealing with data from multiple time zones.
//
// Example usage:
//
// b := es.QueryString("value").
// TimeZone("UTC")
//
//	q := QueryString("timestamp:[2024-01-01 TO 2024-12-31]").TimeZone("America/New_York")
//	// q now applies the "America/New_York" time zone to date and time fields in the query string.
//
// Parameters:
//   - value: A string representing the time zone to be used for date and time fields, such as "UTC" or "America/New_York".
//
// Returns:
//
// The updated queryStringType object with the "time_zone" option set.
func (q queryStringType) TimeZone(value string) queryStringType {
	q.putInTheField("time_zone", value)
	return q
}

func (q queryStringType) putInTheField(key string, value any) queryStringType {
	for field := range q {
		if fieldObject, ok := q[field].(Object); ok {
			fieldObject[key] = value
		}
	}
	return q
}
