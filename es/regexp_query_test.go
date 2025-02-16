package es_test

import (
	"testing"

	"github.com/Trendyol/es-query-builder/es"
	"github.com/Trendyol/es-query-builder/test/assert"
)

////   Regexp   ////

func Test_Regexp_should_exist_on_es_package(t *testing.T) {
	t.Parallel()
	// Given When Then
	assert.NotNil(t, es.Regexp)
}

func Test_Regexp_should_have_Boost_method(t *testing.T) {
	t.Parallel()
	// Given
	regex := es.Regexp("key", "[a-z0-9]")

	// When Then
	assert.NotNil(t, regex.Boost)
}

func Test_Regexp_should_have_Flags_method(t *testing.T) {
	t.Parallel()
	// Given
	regex := es.Regexp("key", "[a-z0-9]")

	// When Then
	assert.NotNil(t, regex.Flags)
}

func Test_Regexp_should_have_Rewrite_method(t *testing.T) {
	t.Parallel()
	// Given
	regex := es.Regexp("key", "[a-z0-9]")

	// When Then
	assert.NotNil(t, regex.Rewrite)
}

func Test_Regexp_should_have_CaseInsensitive_method(t *testing.T) {
	t.Parallel()
	// Given
	regex := es.Regexp("key", "[a-z0-9]")

	// When Then
	assert.NotNil(t, regex.CaseInsensitive)
}

func Test_Regexp_should_have_MaxDeterminizedStates_method(t *testing.T) {
	t.Parallel()
	// Given
	regex := es.Regexp("key", "[a-z0-9]")

	// When Then
	assert.NotNil(t, regex.MaxDeterminizedStates)
}

func Test_Regexp_should_create_json_with_regexp_field_inside_query(t *testing.T) {
	t.Parallel()
	// Given
	query := es.NewQuery(
		es.Regexp("endpoint", "/books/.*"),
	)

	// When Then
	assert.NotNil(t, query)
	bodyJSON := assert.MarshalWithoutError(t, query)
	assert.Equal(t, "{\"query\":{\"regexp\":{\"endpoint\":{\"value\":\"/books/.*\"}}}}", bodyJSON)
}

func Test_Regexp_method_should_create_regexpType(t *testing.T) {
	t.Parallel()
	// Given
	b := es.Regexp("key", "value")

	// Then
	assert.NotNil(t, b)
	assert.IsTypeString(t, "es.regexpType", b)
}

func Test_Regexp_should_create_json_with_match_all_field_inside_caseinsensitive_query(t *testing.T) {
	t.Parallel()
	// Given
	query := es.NewQuery(
		es.Regexp("key", "value1").
			CaseInsensitive(false),
	)

	// When Then
	assert.NotNil(t, query)
	bodyJSON := assert.MarshalWithoutError(t, query)
	assert.Equal(t, "{\"query\":{\"regexp\":{\"key\":{\"case_insensitive\":false,\"value\":\"value1\"}}}}", bodyJSON)
}

func Test_Regexp_should_create_json_with_match_all_field_inside_maxdeterminizedstates_query(t *testing.T) {
	t.Parallel()
	// Given
	query := es.NewQuery(
		es.Regexp("key", "value1").
			MaxDeterminizedStates(1000),
	)

	// When Then
	assert.NotNil(t, query)
	bodyJSON := assert.MarshalWithoutError(t, query)
	assert.Equal(t, "{\"query\":{\"regexp\":{\"key\":{\"max_determinized_states\":1000,\"value\":\"value1\"}}}}", bodyJSON)
}

func Test_Regexp_should_create_json_with_match_all_field_inside_rewrite_query(t *testing.T) {
	t.Parallel()
	// Given
	query := es.NewQuery(
		es.Regexp("key", "value1").
			Rewrite("a"),
	)

	// When Then
	assert.NotNil(t, query)
	bodyJSON := assert.MarshalWithoutError(t, query)
	assert.Equal(t, "{\"query\":{\"regexp\":{\"key\":{\"rewrite\":\"a\",\"value\":\"value1\"}}}}", bodyJSON)
}

func Test_Regexp_should_create_json_with_match_all_field_inside_flags_query(t *testing.T) {
	t.Parallel()
	// Given
	query := es.NewQuery(
		es.Regexp("key", "value1").
			Flags("ALL"),
	)

	// When Then
	assert.NotNil(t, query)
	bodyJSON := assert.MarshalWithoutError(t, query)
	assert.Equal(t, "{\"query\":{\"regexp\":{\"key\":{\"flags\":\"ALL\",\"value\":\"value1\"}}}}", bodyJSON)
}

func Test_Regexp_Boost_should_create_json_with_boost_field_inside_regexp(t *testing.T) {
	t.Parallel()
	// Given
	query := es.NewQuery(
		es.Regexp("name", "[bst]ake").
			Boost(2.718),
	)

	// When Then
	assert.NotNil(t, query)
	bodyJSON := assert.MarshalWithoutError(t, query)
	assert.Equal(t, "{\"query\":{\"regexp\":{\"name\":{\"boost\":2.718,\"value\":\"[bst]ake\"}}}}", bodyJSON)
}
