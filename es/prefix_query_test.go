package es_test

import (
	"testing"

	"github.com/Trendyol/es-query-builder/es"
	"github.com/Trendyol/es-query-builder/test/assert"
)

////   Prefix   ////

func Test_Prefix_should_exist_on_es_package(t *testing.T) {
	t.Parallel()
	// Given When Then
	assert.NotNil(t, es.Prefix)
}

func Test_Prefix_should_create_json_with_prefix_field_inside_query(t *testing.T) {
	t.Parallel()
	// Given
	query := es.NewQuery(
		es.Prefix("user.id", "ki"),
	)

	// When Then
	assert.NotNil(t, query)
	bodyJSON := assert.MarshalWithoutError(t, query)
	assert.Equal(t, "{\"query\":{\"prefix\":{\"user.id\":{\"value\":\"ki\"}}}}", bodyJSON)
}

func Test_Prefix_method_should_create_prefixType(t *testing.T) {
	t.Parallel()
	// Given
	p := es.Prefix("user.id", "ki")

	// Then
	assert.NotNil(t, p)
	assert.IsTypeString(t, "es.prefixType", p)
}

func Test_Prefix_should_have_CaseInsensitive_method(t *testing.T) {
	t.Parallel()
	// Given
	prefix := es.Prefix("user.id", "ki")

	// When Then
	assert.NotNil(t, prefix.CaseInsensitive)
}

func Test_Prefix_CaseInsensitive_should_create_json_with_case_insensitive_field_inside_prefix(t *testing.T) {
	t.Parallel()
	// Given
	query := es.NewQuery(
		es.Prefix("user.id", "Ki").
			CaseInsensitive(true),
	)

	// When Then
	assert.NotNil(t, query)
	bodyJSON := assert.MarshalWithoutError(t, query)
	assert.Equal(t, "{\"query\":{\"prefix\":{\"user.id\":{\"case_insensitive\":true,\"value\":\"Ki\"}}}}", bodyJSON)
}

func Test_Prefix_should_have_Rewrite_method(t *testing.T) {
	t.Parallel()
	// Given
	prefix := es.Prefix("user.id", "ki")

	// When Then
	assert.NotNil(t, prefix.Rewrite)
}

func Test_Prefix_Rewrite_should_create_json_with_rewrite_field_inside_prefix(t *testing.T) {
	t.Parallel()
	// Given
	query := es.NewQuery(
		es.Prefix("user.id", "ki").
			Rewrite("constant_score"),
	)

	// When Then
	assert.NotNil(t, query)
	bodyJSON := assert.MarshalWithoutError(t, query)
	assert.Equal(t, "{\"query\":{\"prefix\":{\"user.id\":{\"rewrite\":\"constant_score\",\"value\":\"ki\"}}}}", bodyJSON)
}

func Test_Prefix_should_have_Boost_method(t *testing.T) {
	t.Parallel()
	// Given
	prefix := es.Prefix("user.id", "ki")

	// When Then
	assert.NotNil(t, prefix.Boost)
}

func Test_Prefix_Boost_should_create_json_with_boost_field_inside_prefix(t *testing.T) {
	t.Parallel()
	// Given
	query := es.NewQuery(
		es.Prefix("user.id", "ki").
			Boost(1.5),
	)

	// When Then
	assert.NotNil(t, query)
	bodyJSON := assert.MarshalWithoutError(t, query)
	assert.Equal(t, "{\"query\":{\"prefix\":{\"user.id\":{\"boost\":1.5,\"value\":\"ki\"}}}}", bodyJSON)
}

func Test_Prefix_should_create_json_with_all_fields(t *testing.T) {
	t.Parallel()
	// Given
	query := es.NewQuery(
		es.Prefix("user.id", "ki").
			CaseInsensitive(true).
			Rewrite("constant_score").
			Boost(2.0),
	)

	// When Then
	assert.NotNil(t, query)
	bodyJSON := assert.MarshalWithoutError(t, query)
	// nolint:golint,lll
	assert.Equal(t, "{\"query\":{\"prefix\":{\"user.id\":{\"boost\":2,\"case_insensitive\":true,\"rewrite\":\"constant_score\",\"value\":\"ki\"}}}}", bodyJSON)
}

////   PrefixFunc   ////

func Test_PrefixFunc_should_exist_on_es_package(t *testing.T) {
	t.Parallel()
	// Given When Then
	assert.NotNil(t, es.PrefixFunc)
}

func Test_PrefixFunc_should_create_json_with_prefix_field_inside_query(t *testing.T) {
	t.Parallel()
	// Given
	query := es.NewQuery(
		es.PrefixFunc("user.id", "ki", func(key string, value string) bool {
			return true
		}),
	)

	// When Then
	assert.NotNil(t, query)
	bodyJSON := assert.MarshalWithoutError(t, query)
	assert.Equal(t, "{\"query\":{\"prefix\":{\"user.id\":{\"value\":\"ki\"}}}}", bodyJSON)
}

func Test_PrefixFunc_should_not_add_prefix_field_inside_query_when_callback_result_is_false(t *testing.T) {
	t.Parallel()
	// Given
	query := es.NewQuery(
		es.PrefixFunc("user.id", "ki", func(key string, value string) bool {
			return false
		}),
	)

	// When Then
	assert.NotNil(t, query)
	bodyJSON := assert.MarshalWithoutError(t, query)
	assert.Equal(t, "{\"query\":{}}", bodyJSON)
}

func Test_PrefixFunc_method_should_create_prefixType(t *testing.T) {
	t.Parallel()
	// Given
	p := es.PrefixFunc("user.id", "ki", func(key string, value string) bool {
		return true
	})

	// Then
	assert.NotNil(t, p)
	assert.IsTypeString(t, "es.prefixType", p)
}

////   PrefixIf   ////

func Test_PrefixIf_should_exist_on_es_package(t *testing.T) {
	t.Parallel()
	// Given When Then
	assert.NotNil(t, es.PrefixIf)
}

func Test_PrefixIf_should_create_json_with_prefix_field_inside_query(t *testing.T) {
	t.Parallel()
	// Given
	query := es.NewQuery(
		es.PrefixIf("user.id", "ki", true),
	)

	// When Then
	assert.NotNil(t, query)
	bodyJSON := assert.MarshalWithoutError(t, query)
	assert.Equal(t, "{\"query\":{\"prefix\":{\"user.id\":{\"value\":\"ki\"}}}}", bodyJSON)
}

func Test_PrefixIf_should_not_add_prefix_field_inside_query_when_condition_is_false(t *testing.T) {
	t.Parallel()
	// Given
	query := es.NewQuery(
		es.PrefixIf("user.id", "ki", false),
	)

	// When Then
	assert.NotNil(t, query)
	bodyJSON := assert.MarshalWithoutError(t, query)
	assert.Equal(t, "{\"query\":{}}", bodyJSON)
}

func Test_PrefixIf_method_should_create_prefixType(t *testing.T) {
	t.Parallel()
	// Given
	p := es.PrefixIf("user.id", "ki", true)

	// Then
	assert.NotNil(t, p)
	assert.IsTypeString(t, "es.prefixType", p)
}
