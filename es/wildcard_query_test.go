package es_test

import (
	"testing"

	"github.com/Trendyol/es-query-builder/es"
	"github.com/Trendyol/es-query-builder/test/assert"
)

////   Wildcard   ////

func Test_Wildcard_should_exist_on_es_package(t *testing.T) {
	t.Parallel()
	// Given When Then
	assert.NotNil(t, es.Wildcard)
}

func Test_Wildcard_should_create_json_with_wildcard_field_inside_query(t *testing.T) {
	t.Parallel()
	// Given
	query := es.NewQuery(
		es.Wildcard("user.id", "ki*y"),
	)

	// When Then
	assert.NotNil(t, query)
	bodyJSON := assert.MarshalWithoutError(t, query)
	assert.Equal(t, "{\"query\":{\"wildcard\":{\"user.id\":{\"value\":\"ki*y\"}}}}", bodyJSON)
}

func Test_Wildcard_method_should_create_wildcardType(t *testing.T) {
	t.Parallel()
	// Given
	w := es.Wildcard("user.id", "ki*y")

	// Then
	assert.NotNil(t, w)
	assert.IsTypeString(t, "es.wildcardType", w)
}

func Test_Wildcard_should_have_CaseInsensitive_method(t *testing.T) {
	t.Parallel()
	// Given
	wildcard := es.Wildcard("user.id", "ki*y")

	// When Then
	assert.NotNil(t, wildcard.CaseInsensitive)
}

func Test_Wildcard_CaseInsensitive_should_create_json_with_case_insensitive_field_inside_wildcard(t *testing.T) {
	t.Parallel()
	// Given
	query := es.NewQuery(
		es.Wildcard("user.id", "Ki*Y").
			CaseInsensitive(true),
	)

	// When Then
	assert.NotNil(t, query)
	bodyJSON := assert.MarshalWithoutError(t, query)
	assert.Equal(t, "{\"query\":{\"wildcard\":{\"user.id\":{\"case_insensitive\":true,\"value\":\"Ki*Y\"}}}}", bodyJSON)
}

func Test_Wildcard_should_have_Rewrite_method(t *testing.T) {
	t.Parallel()
	// Given
	wildcard := es.Wildcard("user.id", "ki*y")

	// When Then
	assert.NotNil(t, wildcard.Rewrite)
}

func Test_Wildcard_Rewrite_should_create_json_with_rewrite_field_inside_wildcard(t *testing.T) {
	t.Parallel()
	// Given
	query := es.NewQuery(
		es.Wildcard("user.id", "ki*y").
			Rewrite("constant_score"),
	)

	// When Then
	assert.NotNil(t, query)
	bodyJSON := assert.MarshalWithoutError(t, query)
	assert.Equal(t, "{\"query\":{\"wildcard\":{\"user.id\":{\"rewrite\":\"constant_score\",\"value\":\"ki*y\"}}}}", bodyJSON)
}

func Test_Wildcard_should_have_Boost_method(t *testing.T) {
	t.Parallel()
	// Given
	wildcard := es.Wildcard("user.id", "ki*y")

	// When Then
	assert.NotNil(t, wildcard.Boost)
}

func Test_Wildcard_Boost_should_create_json_with_boost_field_inside_wildcard(t *testing.T) {
	t.Parallel()
	// Given
	query := es.NewQuery(
		es.Wildcard("user.id", "ki*y").
			Boost(1.5),
	)

	// When Then
	assert.NotNil(t, query)
	bodyJSON := assert.MarshalWithoutError(t, query)
	assert.Equal(t, "{\"query\":{\"wildcard\":{\"user.id\":{\"boost\":1.5,\"value\":\"ki*y\"}}}}", bodyJSON)
}

func Test_Wildcard_should_create_json_with_all_fields(t *testing.T) {
	t.Parallel()
	// Given
	query := es.NewQuery(
		es.Wildcard("user.id", "ki*y").
			CaseInsensitive(true).
			Rewrite("constant_score").
			Boost(2.0),
	)

	// When Then
	assert.NotNil(t, query)
	bodyJSON := assert.MarshalWithoutError(t, query)
	assert.Equal(t, "{\"query\":{\"wildcard\":{\"user.id\":{\"boost\":2,\"case_insensitive\":true,\"rewrite\":\"constant_score\",\"value\":\"ki*y\"}}}}", bodyJSON)
}

////   WildcardFunc   ////

func Test_WildcardFunc_should_exist_on_es_package(t *testing.T) {
	t.Parallel()
	// Given When Then
	assert.NotNil(t, es.WildcardFunc)
}

func Test_WildcardFunc_should_create_json_with_wildcard_field_inside_query(t *testing.T) {
	t.Parallel()
	// Given
	query := es.NewQuery(
		es.WildcardFunc("user.id", "ki*y", func(key string, value string) bool {
			return true
		}),
	)

	// When Then
	assert.NotNil(t, query)
	bodyJSON := assert.MarshalWithoutError(t, query)
	assert.Equal(t, "{\"query\":{\"wildcard\":{\"user.id\":{\"value\":\"ki*y\"}}}}", bodyJSON)
}

func Test_WildcardFunc_should_not_add_wildcard_field_inside_query_when_callback_result_is_false(t *testing.T) {
	t.Parallel()
	// Given
	query := es.NewQuery(
		es.WildcardFunc("user.id", "ki*y", func(key string, value string) bool {
			return false
		}),
	)

	// When Then
	assert.NotNil(t, query)
	bodyJSON := assert.MarshalWithoutError(t, query)
	assert.Equal(t, "{\"query\":{}}", bodyJSON)
}

func Test_WildcardFunc_method_should_create_wildcardType(t *testing.T) {
	t.Parallel()
	// Given
	w := es.WildcardFunc("user.id", "ki*y", func(key string, value string) bool {
		return true
	})

	// Then
	assert.NotNil(t, w)
	assert.IsTypeString(t, "es.wildcardType", w)
}

////   WildcardIf   ////

func Test_WildcardIf_should_exist_on_es_package(t *testing.T) {
	t.Parallel()
	// Given When Then
	assert.NotNil(t, es.WildcardIf)
}

func Test_WildcardIf_should_create_json_with_wildcard_field_inside_query(t *testing.T) {
	t.Parallel()
	// Given
	query := es.NewQuery(
		es.WildcardIf("user.id", "ki*y", true),
	)

	// When Then
	assert.NotNil(t, query)
	bodyJSON := assert.MarshalWithoutError(t, query)
	assert.Equal(t, "{\"query\":{\"wildcard\":{\"user.id\":{\"value\":\"ki*y\"}}}}", bodyJSON)
}

func Test_WildcardIf_should_not_add_wildcard_field_inside_query_when_condition_is_false(t *testing.T) {
	t.Parallel()
	// Given
	query := es.NewQuery(
		es.WildcardIf("user.id", "ki*y", false),
	)

	// When Then
	assert.NotNil(t, query)
	bodyJSON := assert.MarshalWithoutError(t, query)
	assert.Equal(t, "{\"query\":{}}", bodyJSON)
}

func Test_WildcardIf_method_should_create_wildcardType(t *testing.T) {
	t.Parallel()
	// Given
	w := es.WildcardIf("user.id", "ki*y", true)

	// Then
	assert.NotNil(t, w)
	assert.IsTypeString(t, "es.wildcardType", w)
}
