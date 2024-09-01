package es_test

import (
	"github.com/GokselKUCUKSAHIN/es-query-builder/es"
	"github.com/GokselKUCUKSAHIN/es-query-builder/test/assert"
	"testing"
)

func Test_QueryString_should_exist_on_es_package(t *testing.T) {
	// Given When Then
	assert.NotNil(t, es.QueryString[any])
}

func Test_QueryString_method_should_create_queryStringType(t *testing.T) {
	// Given
	b := es.QueryString("value")

	// Then
	assert.NotNil(t, b)
	assert.IsTypeString(t, "es.queryStringType", b)
}

func Test_QueryString_method_should_create_query_string_with_required_query(t *testing.T) {
	// Given
	b := es.NewQuery(
		es.QueryString("value"),
	)

	// Then
	assert.NotNil(t, b)
	bodyJSON := assert.MarshalWithoutError(t, b)
	assert.Equal(t, "{\"query\":{\"query_string\":{\"query\":\"value\"}}}", bodyJSON)
}

func Test_QueryString_method_should_create_query_string_with_default_field(t *testing.T) {
	// Given
	b := es.NewQuery(
		es.QueryString("value").DefaultField("defaultField"),
	)

	// Then
	assert.NotNil(t, b)
	bodyJSON := assert.MarshalWithoutError(t, b)
	assert.Equal(t, "{\"query\":{\"query_string\":{\"default_field\":\"defaultField\",\"query\":\"value\"}}}", bodyJSON)
}

func Test_QueryString_method_should_create_query_string_with_allow_leading_wildcard(t *testing.T) {
	// Given
	b := es.NewQuery(
		es.QueryString("value").AllowLeadingWildcard(false),
	)

	// Then
	assert.NotNil(t, b)
	bodyJSON := assert.MarshalWithoutError(t, b)
	assert.Equal(t, "{\"query\":{\"query_string\":{\"allow_leading_wildcard\":false,\"query\":\"value\"}}}", bodyJSON)
}

func Test_QueryString_method_should_create_query_string_with_analyze_wildcard(t *testing.T) {
	// Given
	b := es.NewQuery(
		es.QueryString("value").AnalyzeWildcard(true),
	)

	// Then
	assert.NotNil(t, b)
	bodyJSON := assert.MarshalWithoutError(t, b)
	assert.Equal(t, "{\"query\":{\"query_string\":{\"analyze_wildcard\":true,\"query\":\"value\"}}}", bodyJSON)
}

func Test_QueryString_method_should_create_query_string_with_analyzer(t *testing.T) {
	// Given
	b := es.NewQuery(
		es.QueryString("value").Analyzer("value"),
	)

	// Then
	assert.NotNil(t, b)
	bodyJSON := assert.MarshalWithoutError(t, b)
	assert.Equal(t, "{\"query\":{\"query_string\":{\"analyzer\":\"value\",\"query\":\"value\"}}}", bodyJSON)
}

func Test_QueryString_method_should_create_query_string_with_auto_generate_synonyms_phrase_query(t *testing.T) {
	// Given
	b := es.NewQuery(
		es.QueryString("value").AutoGenerateSynonymsPhraseQuery(false),
	)

	// Then
	assert.NotNil(t, b)
	bodyJSON := assert.MarshalWithoutError(t, b)
	assert.Equal(t, "{\"query\":{\"query_string\":{\"auto_generate_synonyms_phrase_query\":false,\"query\":\"value\"}}}", bodyJSON)
}

func Test_QueryString_method_should_create_query_string_with_boost(t *testing.T) {
	// Given
	b := es.NewQuery(
		es.QueryString("value").Boost(2.5),
	)

	// Then
	assert.NotNil(t, b)
	bodyJSON := assert.MarshalWithoutError(t, b)
	assert.Equal(t, "{\"query\":{\"query_string\":{\"boost\":2.5,\"query\":\"value\"}}}", bodyJSON)
}

func Test_QueryString_method_should_create_query_string_with_default_operator(t *testing.T) {
	// Given
	b := es.NewQuery(
		es.QueryString("value").DefaultOperator("value"),
	)

	// Then
	assert.NotNil(t, b)
	bodyJSON := assert.MarshalWithoutError(t, b)
	assert.Equal(t, "{\"query\":{\"query_string\":{\"default_operator\":\"value\",\"query\":\"value\"}}}", bodyJSON)
}

func Test_QueryString_method_should_create_query_string_with_enable_position_increments(t *testing.T) {
	// Given
	b := es.NewQuery(
		es.QueryString("value").EnablePositionIncrements(false),
	)

	// Then
	assert.NotNil(t, b)
	bodyJSON := assert.MarshalWithoutError(t, b)
	assert.Equal(t, "{\"query\":{\"query_string\":{\"enable_position_increments\":false,\"query\":\"value\"}}}", bodyJSON)
}

func Test_QueryString_method_should_create_query_string_with_fields(t *testing.T) {
	// Given
	b := es.NewQuery(
		es.QueryString("value").Fields([]string{"field1", "field2"}),
	)

	// Then
	assert.NotNil(t, b)
	bodyJSON := assert.MarshalWithoutError(t, b)
	assert.Equal(t, "{\"query\":{\"query_string\":{\"fields\":[\"field1\",\"field2\"],\"query\":\"value\"}}}", bodyJSON)
}

func Test_QueryString_method_should_create_query_string_with_fuzziness(t *testing.T) {
	// Given
	b := es.NewQuery(
		es.QueryString("value").Fuzziness("value"),
	)

	// Then
	assert.NotNil(t, b)
	bodyJSON := assert.MarshalWithoutError(t, b)
	assert.Equal(t, "{\"query\":{\"query_string\":{\"fuzziness\":\"value\",\"query\":\"value\"}}}", bodyJSON)
}

func Test_QueryString_method_should_create_query_string_with_fuzzy_max_expansions(t *testing.T) {
	// Given
	b := es.NewQuery(
		es.QueryString("value").FuzzyMaxExpansions(50),
	)

	// Then
	assert.NotNil(t, b)
	bodyJSON := assert.MarshalWithoutError(t, b)
	assert.Equal(t, "{\"query\":{\"query_string\":{\"fuzzy_max_expansions\":50,\"query\":\"value\"}}}", bodyJSON)
}

func Test_QueryString_method_should_create_query_string_with_fuzzy_prefix_length(t *testing.T) {
	// Given
	b := es.NewQuery(
		es.QueryString("value").FuzzyPrefixLength(1),
	)

	// Then
	assert.NotNil(t, b)
	bodyJSON := assert.MarshalWithoutError(t, b)
	assert.Equal(t, "{\"query\":{\"query_string\":{\"fuzzy_prefix_length\":1,\"query\":\"value\"}}}", bodyJSON)
}

func Test_QueryString_method_should_create_query_string_with_fuzzy_transpositions(t *testing.T) {
	// Given
	b := es.NewQuery(
		es.QueryString("value").FuzzyTranspositions(false),
	)

	// Then
	assert.NotNil(t, b)
	bodyJSON := assert.MarshalWithoutError(t, b)
	assert.Equal(t, "{\"query\":{\"query_string\":{\"fuzzy_transpositions\":false,\"query\":\"value\"}}}", bodyJSON)
}

func Test_QueryString_method_should_create_query_string_with_lenient(t *testing.T) {
	// Given
	b := es.NewQuery(
		es.QueryString("value").Lenient(true),
	)

	// Then
	assert.NotNil(t, b)
	bodyJSON := assert.MarshalWithoutError(t, b)
	assert.Equal(t, "{\"query\":{\"query_string\":{\"lenient\":true,\"query\":\"value\"}}}", bodyJSON)
}

func Test_QueryString_method_should_create_query_string_with_max_determinized_states(t *testing.T) {
	// Given
	b := es.NewQuery(
		es.QueryString("value").MaxDeterminizedStates(5000),
	)

	// Then
	assert.NotNil(t, b)
	bodyJSON := assert.MarshalWithoutError(t, b)
	assert.Equal(t, "{\"query\":{\"query_string\":{\"max_determinized_states\":5000,\"query\":\"value\"}}}", bodyJSON)
}

func Test_QueryString_method_should_create_query_string_with_minimum_should_match(t *testing.T) {
	// Given
	b := es.NewQuery(
		es.QueryString("value").MinimumShouldMatch("1"),
	)

	// Then
	assert.NotNil(t, b)
	bodyJSON := assert.MarshalWithoutError(t, b)
	assert.Equal(t, "{\"query\":{\"query_string\":{\"minimum_should_match\":\"1\",\"query\":\"value\"}}}", bodyJSON)
}

func Test_QueryString_method_should_create_query_string_with_quote_analyzer(t *testing.T) {
	// Given
	b := es.NewQuery(
		es.QueryString("value").QuoteAnalyzer("value"),
	)

	// Then
	assert.NotNil(t, b)
	bodyJSON := assert.MarshalWithoutError(t, b)
	assert.Equal(t, "{\"query\":{\"query_string\":{\"query\":\"value\",\"quote_analyzer\":\"value\"}}}", bodyJSON)
}

func Test_QueryString_method_should_create_query_string_with_phrase_slop(t *testing.T) {
	// Given
	b := es.NewQuery(
		es.QueryString("value").PhraseSlop(0),
	)

	// Then
	assert.NotNil(t, b)
	bodyJSON := assert.MarshalWithoutError(t, b)
	assert.Equal(t, "{\"query\":{\"query_string\":{\"phrase_slop\":0,\"query\":\"value\"}}}", bodyJSON)
}

func Test_QueryString_method_should_create_query_string_with_quote_field_suffix(t *testing.T) {
	// Given
	b := es.NewQuery(
		es.QueryString("value").QuoteFieldSuffix("value"),
	)

	// Then
	assert.NotNil(t, b)
	bodyJSON := assert.MarshalWithoutError(t, b)
	assert.Equal(t, "{\"query\":{\"query_string\":{\"query\":\"value\",\"quote_field_suffix\":\"value\"}}}", bodyJSON)
}

func Test_QueryString_method_should_create_query_string_with_rewrite(t *testing.T) {
	// Given
	b := es.NewQuery(
		es.QueryString("value").Rewrite("value"),
	)

	// Then
	assert.NotNil(t, b)
	bodyJSON := assert.MarshalWithoutError(t, b)
	assert.Equal(t, "{\"query\":{\"query_string\":{\"query\":\"value\",\"rewrite\":\"value\"}}}", bodyJSON)
}

func Test_QueryString_method_should_create_query_string_with_time_zone(t *testing.T) {
	// Given
	b := es.NewQuery(
		es.QueryString("value").TimeZone("value"),
	)

	// Then
	assert.NotNil(t, b)
	bodyJSON := assert.MarshalWithoutError(t, b)
	assert.Equal(t, "{\"query\":{\"query_string\":{\"query\":\"value\",\"time_zone\":\"value\"}}}", bodyJSON)
}

func Test_QueryString_method_should_create_query_string_with_all_parameters(t *testing.T) {
	// Given
	b := es.NewQuery(
		es.QueryString("value").
			DefaultField("value").
			AllowLeadingWildcard(false).
			AnalyzeWildcard(true).
			Analyzer("value").
			AutoGenerateSynonymsPhraseQuery(true).
			Boost(2.5).
			DefaultOperator("value").
			EnablePositionIncrements(true).
			Fields([]string{"field1", "field2"}).
			Fuzziness("value").
			FuzzyMaxExpansions(2).
			FuzzyPrefixLength(1.0).
			FuzzyTranspositions(true).
			Lenient(true).
			MaxDeterminizedStates(2).MinimumShouldMatch("value").
			QuoteAnalyzer("value").
			PhraseSlop(2).
			QuoteFieldSuffix("value").
			Rewrite("value").
			TimeZone("value"),
	)

	// Then
	assert.NotNil(t, b)
	bodyJSON := assert.MarshalWithoutError(t, b)
	assert.Equal(t, "{\"query\":{\"query_string\":{\"allow_leading_wildcard\":false,\"analyze_wildcard\":true,"+
		"\"analyzer\":\"value\",\"auto_generate_synonyms_phrase_query\":true,"+
		"\"boost\":2.5,\"default_field\":\"value\",\"default_operator\":\"value\","+
		"\"enable_position_increments\":true,\"fields\":[\"field1\",\"field2\"],"+
		"\"fuzziness\":\"value\",\"fuzzy_max_expansions\":2,\"fuzzy_prefix_length\":1,"+
		"\"fuzzy_transpositions\":true,\"lenient\":true,\"max_determinized_states\":2,"+
		"\"minimum_should_match\":\"value\",\"phrase_slop\":2,\"query\":\"value\",\"quote_analyzer\":\"value\","+
		"\"quote_field_suffix\":\"value\",\"rewrite\":\"value\",\"time_zone\":\"value\"}}}", bodyJSON)
}
