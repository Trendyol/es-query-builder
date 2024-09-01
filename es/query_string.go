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
//	q := QueryString("Go AND library")
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
// DefaultField("defaultField"),
// )
//
//	q := QueryString("Go library").DefaultField("title")
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
// AllowLeadingWildcard(true),
// )
//
//	q := QueryString("Go* library").AllowLeadingWildcard(true)
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
// AnalyzeWildcard(true),
// )
//
//	q := QueryString("Go* library").AnalyzeWildcard(true)
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
// Analyzer("custom_analyzer"),
// )
//
//	q := QueryString("Go library").Analyzer("standard")
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
// AutoGenerateSynonymsPhraseQuery(true),
// )
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
// Boost(2.0),
// )
//
//	q := QueryString("Go library").Boost(1.5)
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
// DefaultOperator("AND"),
// )
//
//	q := QueryString("Go library").DefaultOperator("OR")
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

func (q queryStringType) EnablePositionIncrements(value bool) queryStringType {
	q.putInTheField("enable_position_increments", value)
	return q
}

func (q queryStringType) Fields(value []string) queryStringType {
	q.putInTheField("fields", value)
	return q
}

func (q queryStringType) Fuzziness(value string) queryStringType {
	q.putInTheField("fuzziness", value)
	return q
}

func (q queryStringType) FuzzyMaxExpansions(value int64) queryStringType {
	q.putInTheField("fuzzy_max_expansions", value)
	return q
}

func (q queryStringType) FuzzyPrefixLength(value int64) queryStringType {
	q.putInTheField("fuzzy_prefix_length", value)
	return q
}

func (q queryStringType) FuzzyTranspositions(value bool) queryStringType {
	q.putInTheField("fuzzy_transpositions", value)
	return q
}

func (q queryStringType) Lenient(value bool) queryStringType {
	q.putInTheField("lenient", value)
	return q
}

func (q queryStringType) MaxDeterminizedStates(value int64) queryStringType {
	q.putInTheField("max_determinized_states", value)
	return q
}

func (q queryStringType) MinimumShouldMatch(value string) queryStringType {
	q.putInTheField("minimum_should_match", value)
	return q
}

func (q queryStringType) QuoteAnalyzer(value string) queryStringType {
	q.putInTheField("quote_analyzer", value)
	return q
}

func (q queryStringType) PhraseSlop(value int64) queryStringType {
	q.putInTheField("phrase_slop", value)
	return q
}

func (q queryStringType) QuoteFieldSuffix(value string) queryStringType {
	q.putInTheField("quote_field_suffix", value)
	return q
}

func (q queryStringType) Rewrite(value string) queryStringType {
	q.putInTheField("rewrite", value)
	return q
}

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
