package es_test

import (
	"testing"

	Operator "github.com/Trendyol/es-query-builder/es/enums/operator"

	"github.com/Trendyol/es-query-builder/es"
	"github.com/Trendyol/es-query-builder/test/assert"
)

////   Simple Query String   ////

func Test_SimpleQueryString_should_exist_on_es_package(t *testing.T) {
	// Given When
	assert.NotNil(t, es.SimpleQueryString[any])
}

func Test_SimpleQueryString_method_should_create_simpleQueryStringType(t *testing.T) {
	// Given When
	b := es.SimpleQueryString("value")

	// Then
	assert.NotNil(t, b)
	assert.IsTypeString(t, "es.simpleQueryStringType", b)
}

func Test_SimpleQueryString_method_should_create_simple_query_string_with_required_query(t *testing.T) {
	// Given When
	b := es.NewQuery(
		es.SimpleQueryString("value"),
	)

	// Then
	assert.NotNil(t, b)
	bodyJSON := assert.MarshalWithoutError(t, b)
	assert.Equal(t, "{\"query\":{\"simple_query_string\":{\"query\":\"value\"}}}", bodyJSON)
}

func Test_SimpleQueryString_method_should_create_simple_query_string_with_fields(t *testing.T) {
	// Given When
	b := es.NewQuery(
		es.SimpleQueryString("value").Fields([]string{"field1", "field2"}),
	)

	// Then
	assert.NotNil(t, b)
	bodyJSON := assert.MarshalWithoutError(t, b)
	assert.Equal(t, "{\"query\":{\"simple_query_string\":{\"fields\":[\"field1\",\"field2\"],\"query\":\"value\"}}}", bodyJSON)
}

func Test_SimpleQueryString_method_should_create_simple_query_string_with_analyzer(t *testing.T) {
	// Given When
	b := es.NewQuery(
		es.SimpleQueryString("value").Analyzer("standard"),
	)

	// Then
	assert.NotNil(t, b)
	bodyJSON := assert.MarshalWithoutError(t, b)
	assert.Equal(t, "{\"query\":{\"simple_query_string\":{\"analyzer\":\"standard\",\"query\":\"value\"}}}", bodyJSON)
}

func Test_SimpleQueryString_method_should_create_simple_query_string_with_default_operator(t *testing.T) {
	// Given When
	b := es.NewQuery(
		es.SimpleQueryString("value").DefaultOperator(Operator.And),
	)

	// Then
	assert.NotNil(t, b)
	bodyJSON := assert.MarshalWithoutError(t, b)
	assert.Equal(t, "{\"query\":{\"simple_query_string\":{\"default_operator\":\"and\",\"query\":\"value\"}}}", bodyJSON)
}

func Test_SimpleQueryString_method_should_create_simple_query_string_with_minimum_should_match(t *testing.T) {
	// Given When
	b := es.NewQuery(
		es.SimpleQueryString("value").MinimumShouldMatch("2"),
	)

	// Then
	assert.NotNil(t, b)
	bodyJSON := assert.MarshalWithoutError(t, b)
	assert.Equal(t, "{\"query\":{\"simple_query_string\":{\"minimum_should_match\":\"2\",\"query\":\"value\"}}}", bodyJSON)
}

func Test_SimpleQueryString_method_should_create_simple_query_string_with_fuzzy_max_expansions(t *testing.T) {
	// Given When
	b := es.NewQuery(
		es.SimpleQueryString("value").FuzzyMaxExpansions(50),
	)

	// Then
	assert.NotNil(t, b)
	bodyJSON := assert.MarshalWithoutError(t, b)
	assert.Equal(t, "{\"query\":{\"simple_query_string\":{\"fuzzy_max_expansions\":50,\"query\":\"value\"}}}", bodyJSON)
}

func Test_SimpleQueryString_method_should_create_simple_query_string_with_fuzzy_prefix_length(t *testing.T) {
	// Given When
	b := es.NewQuery(
		es.SimpleQueryString("value").FuzzyPrefixLength(2),
	)

	// Then
	assert.NotNil(t, b)
	bodyJSON := assert.MarshalWithoutError(t, b)
	assert.Equal(t, "{\"query\":{\"simple_query_string\":{\"fuzzy_prefix_length\":2,\"query\":\"value\"}}}", bodyJSON)
}

func Test_SimpleQueryString_method_should_create_simple_query_string_with_fuzzy_transpositions(t *testing.T) {
	// Given When
	b := es.NewQuery(
		es.SimpleQueryString("value").FuzzyTranspositions(true),
	)

	// Then
	assert.NotNil(t, b)
	bodyJSON := assert.MarshalWithoutError(t, b)
	assert.Equal(t, "{\"query\":{\"simple_query_string\":{\"fuzzy_transpositions\":true,\"query\":\"value\"}}}", bodyJSON)
}

func Test_SimpleQueryString_method_should_create_simple_query_string_with_analyze_wildcard(t *testing.T) {
	// Given When
	b := es.NewQuery(
		es.SimpleQueryString("value").AnalyzeWildcard(true),
	)

	// Then
	assert.NotNil(t, b)
	bodyJSON := assert.MarshalWithoutError(t, b)
	assert.Equal(t, "{\"query\":{\"simple_query_string\":{\"analyze_wildcard\":true,\"query\":\"value\"}}}", bodyJSON)
}

func Test_SimpleQueryString_method_should_create_simple_query_string_with_auto_generate_synonyms_phrase_query(t *testing.T) {
	// Given When
	b := es.NewQuery(
		es.SimpleQueryString("value").AutoGenerateSynonymsPhraseQuery(true),
	)

	// Then
	assert.NotNil(t, b)
	bodyJSON := assert.MarshalWithoutError(t, b)
	assert.Equal(t, "{\"query\":{\"simple_query_string\":{\"auto_generate_synonyms_phrase_query\":true,\"query\":\"value\"}}}", bodyJSON)
}

func Test_SimpleQueryString_method_should_create_simple_query_string_with_flags(t *testing.T) {
	// Given When
	b := es.NewQuery(
		es.SimpleQueryString("value").Flags("AND|OR|PREFIX"),
	)

	// Then
	assert.NotNil(t, b)
	bodyJSON := assert.MarshalWithoutError(t, b)
	assert.Equal(t, "{\"query\":{\"simple_query_string\":{\"flags\":\"AND|OR|PREFIX\",\"query\":\"value\"}}}", bodyJSON)
}

func Test_SimpleQueryString_method_should_create_simple_query_string_with_lenient(t *testing.T) {
	// Given When
	b := es.NewQuery(
		es.SimpleQueryString("value").Lenient(true),
	)

	// Then
	assert.NotNil(t, b)
	bodyJSON := assert.MarshalWithoutError(t, b)
	assert.Equal(t, "{\"query\":{\"simple_query_string\":{\"lenient\":true,\"query\":\"value\"}}}", bodyJSON)
}

func Test_SimpleQueryString_method_should_create_simple_query_string_with_quote_field_suffix(t *testing.T) {
	// Given When
	b := es.NewQuery(
		es.SimpleQueryString("value").QuoteFieldSuffix("_phrase"),
	)

	// Then
	assert.NotNil(t, b)
	bodyJSON := assert.MarshalWithoutError(t, b)
	assert.Equal(t, "{\"query\":{\"simple_query_string\":{\"query\":\"value\",\"quote_field_suffix\":\"_phrase\"}}}", bodyJSON)
}

func Test_SimpleQueryString_method_should_create_simple_query_string_with_boost(t *testing.T) {
	// Given When
	b := es.NewQuery(
		es.SimpleQueryString("value").Boost(3.12),
	)

	// Then
	assert.NotNil(t, b)
	bodyJSON := assert.MarshalWithoutError(t, b)
	assert.Equal(t, "{\"query\":{\"simple_query_string\":{\"boost\":3.12,\"query\":\"value\"}}}", bodyJSON)
}

func Test_SimpleQueryString_method_should_create_simple_query_string_with_all_parameters(t *testing.T) {
	// Given When
	b := es.NewQuery(
		es.SimpleQueryString("value").
			Fields([]string{"field1", "field2"}).
			Analyzer("standard").
			DefaultOperator(Operator.And).
			MinimumShouldMatch("2").
			FuzzyMaxExpansions(50).
			FuzzyPrefixLength(2).
			FuzzyTranspositions(true).
			AnalyzeWildcard(true).
			AutoGenerateSynonymsPhraseQuery(true).
			Flags("AND|OR|PREFIX").
			Lenient(true).
			QuoteFieldSuffix("_phrase").
			Boost(5.19),
	)

	// Then
	assert.NotNil(t, b)
	bodyJSON := assert.MarshalWithoutError(t, b)
	assert.Equal(t, "{\"query\":{\"simple_query_string\":{"+
		"\"analyze_wildcard\":true,"+
		"\"analyzer\":\"standard\","+
		"\"auto_generate_synonyms_phrase_query\":true,"+
		"\"boost\":5.19,"+
		"\"default_operator\":\"and\","+
		"\"fields\":[\"field1\",\"field2\"],"+
		"\"flags\":\"AND|OR|PREFIX\","+
		"\"fuzzy_max_expansions\":50,"+
		"\"fuzzy_prefix_length\":2,"+
		"\"fuzzy_transpositions\":true,"+
		"\"lenient\":true,"+
		"\"minimum_should_match\":\"2\","+
		"\"query\":\"value\","+
		"\"quote_field_suffix\":\"_phrase\""+
		"}}}", bodyJSON)
}
