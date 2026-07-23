package es_test

import (
	"testing"

	"github.com/Trendyol/es-query-builder/es"
	"github.com/Trendyol/es-query-builder/test/assert"
)

////   Fuzzy   ////

func Test_Fuzzy_should_exist_on_es_package(t *testing.T) {
	t.Parallel()
	// Given When Then
	assert.NotNil(t, es.Fuzzy)
}

func Test_Fuzzy_should_create_json_with_fuzzy_field_inside_query(t *testing.T) {
	t.Parallel()
	// Given
	query := es.NewQuery(
		es.Fuzzy("user", "ki"),
	)

	// When Then
	assert.NotNil(t, query)
	bodyJSON := assert.MarshalWithoutError(t, query)
	assert.Equal(t, "{\"query\":{\"fuzzy\":{\"user\":{\"value\":\"ki\"}}}}", bodyJSON)
}

func Test_Fuzzy_method_should_create_fuzzyType(t *testing.T) {
	t.Parallel()
	// Given
	f := es.Fuzzy("user", "ki")

	// Then
	assert.NotNil(t, f)
	assert.IsTypeString(t, "es.fuzzyType", f)
}

func Test_Fuzzy_should_have_Fuzziness_method(t *testing.T) {
	t.Parallel()
	// Given
	fuzzy := es.Fuzzy("user", "ki")

	// When Then
	assert.NotNil(t, fuzzy.Fuzziness)
}

func Test_Fuzzy_Fuzziness_should_create_json_with_string_fuzziness(t *testing.T) {
	t.Parallel()
	// Given
	query := es.NewQuery(
		es.Fuzzy("user", "ki").
			Fuzziness("AUTO"),
	)

	// When Then
	assert.NotNil(t, query)
	bodyJSON := assert.MarshalWithoutError(t, query)
	assert.Equal(t, "{\"query\":{\"fuzzy\":{\"user\":{\"fuzziness\":\"AUTO\",\"value\":\"ki\"}}}}", bodyJSON)
}

func Test_Fuzzy_Fuzziness_should_create_json_with_int_fuzziness(t *testing.T) {
	t.Parallel()
	// Given
	query := es.NewQuery(
		es.Fuzzy("user", "ki").
			Fuzziness(1),
	)

	// When Then
	assert.NotNil(t, query)
	bodyJSON := assert.MarshalWithoutError(t, query)
	assert.Equal(t, "{\"query\":{\"fuzzy\":{\"user\":{\"fuzziness\":1,\"value\":\"ki\"}}}}", bodyJSON)
}

func Test_Fuzzy_should_have_MaxExpansions_method(t *testing.T) {
	t.Parallel()
	// Given
	fuzzy := es.Fuzzy("user", "ki")

	// When Then
	assert.NotNil(t, fuzzy.MaxExpansions)
}

func Test_Fuzzy_MaxExpansions_should_create_json_with_max_expansions_field(t *testing.T) {
	t.Parallel()
	// Given
	query := es.NewQuery(
		es.Fuzzy("user", "ki").
			MaxExpansions(50),
	)

	// When Then
	assert.NotNil(t, query)
	bodyJSON := assert.MarshalWithoutError(t, query)
	assert.Equal(t, "{\"query\":{\"fuzzy\":{\"user\":{\"max_expansions\":50,\"value\":\"ki\"}}}}", bodyJSON)
}

func Test_Fuzzy_should_have_PrefixLength_method(t *testing.T) {
	t.Parallel()
	// Given
	fuzzy := es.Fuzzy("user", "ki")

	// When Then
	assert.NotNil(t, fuzzy.PrefixLength)
}

func Test_Fuzzy_PrefixLength_should_create_json_with_prefix_length_field(t *testing.T) {
	t.Parallel()
	// Given
	query := es.NewQuery(
		es.Fuzzy("user", "ki").
			PrefixLength(0),
	)

	// When Then
	assert.NotNil(t, query)
	bodyJSON := assert.MarshalWithoutError(t, query)
	assert.Equal(t, "{\"query\":{\"fuzzy\":{\"user\":{\"prefix_length\":0,\"value\":\"ki\"}}}}", bodyJSON)
}

func Test_Fuzzy_should_have_Transpositions_method(t *testing.T) {
	t.Parallel()
	// Given
	fuzzy := es.Fuzzy("user", "ki")

	// When Then
	assert.NotNil(t, fuzzy.Transpositions)
}

func Test_Fuzzy_Transpositions_should_create_json_with_transpositions_field(t *testing.T) {
	t.Parallel()
	// Given
	query := es.NewQuery(
		es.Fuzzy("user", "ki").
			Transpositions(true),
	)

	// When Then
	assert.NotNil(t, query)
	bodyJSON := assert.MarshalWithoutError(t, query)
	assert.Equal(t, "{\"query\":{\"fuzzy\":{\"user\":{\"transpositions\":true,\"value\":\"ki\"}}}}", bodyJSON)
}

func Test_Fuzzy_should_have_Rewrite_method(t *testing.T) {
	t.Parallel()
	// Given
	fuzzy := es.Fuzzy("user", "ki")

	// When Then
	assert.NotNil(t, fuzzy.Rewrite)
}

func Test_Fuzzy_Rewrite_should_create_json_with_rewrite_field(t *testing.T) {
	t.Parallel()
	// Given
	query := es.NewQuery(
		es.Fuzzy("user", "ki").
			Rewrite("constant_score"),
	)

	// When Then
	assert.NotNil(t, query)
	bodyJSON := assert.MarshalWithoutError(t, query)
	assert.Equal(t, "{\"query\":{\"fuzzy\":{\"user\":{\"rewrite\":\"constant_score\",\"value\":\"ki\"}}}}", bodyJSON)
}

func Test_Fuzzy_should_have_CaseInsensitive_method(t *testing.T) {
	t.Parallel()
	// Given
	fuzzy := es.Fuzzy("user", "ki")

	// When Then
	assert.NotNil(t, fuzzy.CaseInsensitive)
}

func Test_Fuzzy_CaseInsensitive_should_create_json_with_case_insensitive_field(t *testing.T) {
	t.Parallel()
	// Given
	query := es.NewQuery(
		es.Fuzzy("user", "Ki").
			CaseInsensitive(true),
	)

	// When Then
	assert.NotNil(t, query)
	bodyJSON := assert.MarshalWithoutError(t, query)
	assert.Equal(t, "{\"query\":{\"fuzzy\":{\"user\":{\"case_insensitive\":true,\"value\":\"Ki\"}}}}", bodyJSON)
}

func Test_Fuzzy_should_have_Boost_method(t *testing.T) {
	t.Parallel()
	// Given
	fuzzy := es.Fuzzy("user", "ki")

	// When Then
	assert.NotNil(t, fuzzy.Boost)
}

func Test_Fuzzy_Boost_should_create_json_with_boost_field(t *testing.T) {
	t.Parallel()
	// Given
	query := es.NewQuery(
		es.Fuzzy("user", "ki").
			Boost(1.5),
	)

	// When Then
	assert.NotNil(t, query)
	bodyJSON := assert.MarshalWithoutError(t, query)
	assert.Equal(t, "{\"query\":{\"fuzzy\":{\"user\":{\"boost\":1.5,\"value\":\"ki\"}}}}", bodyJSON)
}

func Test_Fuzzy_should_create_json_with_all_fields(t *testing.T) {
	t.Parallel()
	// Given
	query := es.NewQuery(
		es.Fuzzy("user", "ki").
			Fuzziness("AUTO").
			MaxExpansions(50).
			PrefixLength(0).
			Transpositions(true).
			Rewrite("constant_score").
			CaseInsensitive(true).
			Boost(1.0),
	)

	// When Then
	assert.NotNil(t, query)
	bodyJSON := assert.MarshalWithoutError(t, query)
	// nolint:golint,lll
	assert.Equal(t, "{\"query\":{\"fuzzy\":{\"user\":{\"boost\":1,\"case_insensitive\":true,\"fuzziness\":\"AUTO\",\"max_expansions\":50,\"prefix_length\":0,\"rewrite\":\"constant_score\",\"transpositions\":true,\"value\":\"ki\"}}}}", bodyJSON)
}

////   FuzzyFunc   ////

func Test_FuzzyFunc_should_exist_on_es_package(t *testing.T) {
	t.Parallel()
	// Given When Then
	assert.NotNil(t, es.FuzzyFunc)
}

func Test_FuzzyFunc_should_create_json_with_fuzzy_field_inside_query(t *testing.T) {
	t.Parallel()
	// Given
	query := es.NewQuery(
		es.FuzzyFunc("user", "ki", func(key string, value string) bool {
			return true
		}),
	)

	// When Then
	assert.NotNil(t, query)
	bodyJSON := assert.MarshalWithoutError(t, query)
	assert.Equal(t, "{\"query\":{\"fuzzy\":{\"user\":{\"value\":\"ki\"}}}}", bodyJSON)
}

func Test_FuzzyFunc_should_not_add_fuzzy_field_inside_query_when_callback_result_is_false(t *testing.T) {
	t.Parallel()
	// Given
	query := es.NewQuery(
		es.FuzzyFunc("user", "ki", func(key string, value string) bool {
			return false
		}),
	)

	// When Then
	assert.NotNil(t, query)
	bodyJSON := assert.MarshalWithoutError(t, query)
	assert.Equal(t, "{\"query\":{}}", bodyJSON)
}

func Test_FuzzyFunc_method_should_create_fuzzyType(t *testing.T) {
	t.Parallel()
	// Given
	f := es.FuzzyFunc("user", "ki", func(key string, value string) bool {
		return true
	})

	// Then
	assert.NotNil(t, f)
	assert.IsTypeString(t, "es.fuzzyType", f)
}

////   FuzzyIf   ////

func Test_FuzzyIf_should_exist_on_es_package(t *testing.T) {
	t.Parallel()
	// Given When Then
	assert.NotNil(t, es.FuzzyIf)
}

func Test_FuzzyIf_should_create_json_with_fuzzy_field_inside_query(t *testing.T) {
	t.Parallel()
	// Given
	query := es.NewQuery(
		es.FuzzyIf("user", "ki", true),
	)

	// When Then
	assert.NotNil(t, query)
	bodyJSON := assert.MarshalWithoutError(t, query)
	assert.Equal(t, "{\"query\":{\"fuzzy\":{\"user\":{\"value\":\"ki\"}}}}", bodyJSON)
}

func Test_FuzzyIf_should_not_add_fuzzy_field_inside_query_when_condition_is_false(t *testing.T) {
	t.Parallel()
	// Given
	query := es.NewQuery(
		es.FuzzyIf("user", "ki", false),
	)

	// When Then
	assert.NotNil(t, query)
	bodyJSON := assert.MarshalWithoutError(t, query)
	assert.Equal(t, "{\"query\":{}}", bodyJSON)
}

func Test_FuzzyIf_method_should_create_fuzzyType(t *testing.T) {
	t.Parallel()
	// Given
	f := es.FuzzyIf("user", "ki", true)

	// Then
	assert.NotNil(t, f)
	assert.IsTypeString(t, "es.fuzzyType", f)
}
