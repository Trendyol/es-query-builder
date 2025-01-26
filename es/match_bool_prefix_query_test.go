package es_test

import (
	"testing"

	Operator "github.com/Trendyol/es-query-builder/es/enums/operator"

	"github.com/Trendyol/es-query-builder/es"
	"github.com/Trendyol/es-query-builder/test/assert"
)

////   Match Bool Prefix   ////

func Test_Match_Bool_Prefix_should_exist_on_es_package(t *testing.T) {
	// Given When Then
	assert.NotNil(t, es.MatchBoolPrefix[any])
}

func Test_Match_Bool_Prefix_method_should_create_matchBoolPrefixType(t *testing.T) {
	// Given
	b := es.MatchBoolPrefix("key", "value")

	// Then
	assert.NotNil(t, b)
	assert.IsTypeString(t, "es.matchBoolPrefixType", b)
}

func Test_Match_Bool_Prefix_should_have_Boost_method(t *testing.T) {
	// Given
	match := es.MatchBoolPrefix("key", "value")

	// When Then
	assert.NotNil(t, match.Boost)
}

func Test_Match_Bool_Prefix_Boost_should_create_json_with_boost_field_inside_match(t *testing.T) {
	// Given
	query := es.NewQuery(
		es.MatchBoolPrefix("type", "Folder").
			Boost(3.14),
	)

	// When Then
	assert.NotNil(t, query)
	bodyJSON := assert.MarshalWithoutError(t, query)
	assert.Equal(t, "{\"query\":{\"match_bool_prefix\":{\"type\":{\"boost\":3.14,\"query\":\"Folder\"}}}}", bodyJSON)
}

func Test_Match_Bool_Prefix_should_have_Analyzer_method(t *testing.T) {
	// Given
	match := es.MatchBoolPrefix("key", "value")

	// When Then
	assert.NotNil(t, match.Analyzer)
}

func Test_Match_Bool_Prefix_Analyzer_should_create_json_with_analyzer_field_inside_match(t *testing.T) {
	// Given
	query := es.NewQuery(
		es.MatchBoolPrefix("type", "Folder").
			Analyzer("standart"),
	)

	// When Then
	assert.NotNil(t, query)
	bodyJSON := assert.MarshalWithoutError(t, query)
	assert.Equal(t, "{\"query\":{\"match_bool_prefix\":{\"type\":{\"analyzer\":\"standart\",\"query\":\"Folder\"}}}}", bodyJSON)
}

func Test_Match_Bool_Prefix_should_have_MinimumShouldMatch_method(t *testing.T) {
	// Given
	match := es.MatchBoolPrefix("key", "value")

	// When Then
	assert.NotNil(t, match.MinimumShouldMatch)
}

func Test_Match_Bool_Prefix_MinimumShouldMatch_should_create_json_with_minimum_should_match_field_inside_match(t *testing.T) {
	// Given
	query := es.NewQuery(
		es.MatchBoolPrefix("type", "Folder").
			MinimumShouldMatch(65),
	)

	// When Then
	assert.NotNil(t, query)
	bodyJSON := assert.MarshalWithoutError(t, query)
	assert.Equal(t, "{\"query\":{\"match_bool_prefix\":{\"type\":{\"minimum_should_match\":65,\"query\":\"Folder\"}}}}", bodyJSON)
}

func Test_Match_Bool_Prefix_should_have_Operator_method(t *testing.T) {
	// Given
	match := es.MatchBoolPrefix("key", "value")

	// When Then
	assert.NotNil(t, match.Operator)
}

func Test_Match_Bool_Prefix_Operator_should_create_json_with_operator_field_inside_match(t *testing.T) {
	// Given
	query := es.NewQuery(
		es.MatchBoolPrefix("type", "Folder").
			Operator(Operator.And),
	)

	// When Then
	assert.NotNil(t, query)
	bodyJSON := assert.MarshalWithoutError(t, query)
	assert.Equal(t, "{\"query\":{\"match_bool_prefix\":{\"type\":{\"operator\":\"and\",\"query\":\"Folder\"}}}}", bodyJSON)
}

func Test_Match_Bool_Prefix_should_have_Fuzziness_method(t *testing.T) {
	// Given
	match := es.MatchBoolPrefix("key", "value")

	// When Then
	assert.NotNil(t, match.Fuzziness)
}

func Test_Match_Bool_Prefix_Fuzziness_should_create_json_with_fuzziness_field_inside_match(t *testing.T) {
	// Given
	query := es.NewQuery(
		es.MatchBoolPrefix("type", "Folder").
			Fuzziness("AUTO"),
	)

	// When Then
	assert.NotNil(t, query)
	bodyJSON := assert.MarshalWithoutError(t, query)
	assert.Equal(t, "{\"query\":{\"match_bool_prefix\":{\"type\":{\"fuzziness\":\"AUTO\",\"query\":\"Folder\"}}}}", bodyJSON)
}

func Test_Match_Bool_Prefix_should_have_FuzzyRewrite_method(t *testing.T) {
	// Given
	match := es.MatchBoolPrefix("key", "value")

	// When Then
	assert.NotNil(t, match.FuzzyRewrite)
}

func Test_Match_Bool_Prefix_FuzzyRewrite_should_create_json_with_fuzzy_rewrite_field_inside_match(t *testing.T) {
	// Given
	query := es.NewQuery(
		es.MatchBoolPrefix("type", "Folder").
			FuzzyRewrite("constant_score"),
	)

	// When Then
	assert.NotNil(t, query)
	bodyJSON := assert.MarshalWithoutError(t, query)
	assert.Equal(t, "{\"query\":{\"match_bool_prefix\":{\"type\":{\"fuzzy_rewrite\":\"constant_score\",\"query\":\"Folder\"}}}}", bodyJSON)
}

func Test_Match_Bool_Prefix_should_have_FuzzyTranspositions_method(t *testing.T) {
	// Given
	match := es.MatchBoolPrefix("key", "value")

	// When Then
	assert.NotNil(t, match.FuzzyTranspositions)
}

func Test_Match_Bool_Prefix_FuzzyTranspositions_should_create_json_with_fuzzy_transpositions_field_inside_match(t *testing.T) {
	// Given
	query := es.NewQuery(
		es.MatchBoolPrefix("type", "Directory").
			FuzzyTranspositions(false),
	)

	// When Then
	assert.NotNil(t, query)
	bodyJSON := assert.MarshalWithoutError(t, query)
	assert.Equal(t, "{\"query\":{\"match_bool_prefix\":{\"type\":{\"fuzzy_transpositions\":false,\"query\":\"Directory\"}}}}", bodyJSON)
}

func Test_Match_Bool_Prefix_should_have_MaxExpansions_method(t *testing.T) {
	// Given
	match := es.MatchBoolPrefix("key", "value")

	// When Then
	assert.NotNil(t, match.MaxExpansions)
}

func Test_Match_Bool_Prefix_MaxExpansions_should_create_json_with_max_expansions_field_inside_match(t *testing.T) {
	// Given
	query := es.NewQuery(
		es.MatchBoolPrefix("type", "Folder").
			MaxExpansions(99),
	)

	// When Then
	assert.NotNil(t, query)
	bodyJSON := assert.MarshalWithoutError(t, query)
	assert.Equal(t, "{\"query\":{\"match_bool_prefix\":{\"type\":{\"max_expansions\":99,\"query\":\"Folder\"}}}}", bodyJSON)
}

func Test_Match_Bool_Prefix_should_have_PrefixLength_method(t *testing.T) {
	// Given
	match := es.MatchBoolPrefix("key", "value")

	// When Then
	assert.NotNil(t, match.PrefixLength)
}

func Test_Match_Bool_Prefix_PrefixLength_should_create_json_with_prefix_length_field_inside_match(t *testing.T) {
	// Given
	query := es.NewQuery(
		es.MatchBoolPrefix("type", "Folder").
			PrefixLength(40),
	)

	// When Then
	assert.NotNil(t, query)
	bodyJSON := assert.MarshalWithoutError(t, query)
	assert.Equal(t, "{\"query\":{\"match_bool_prefix\":{\"type\":{\"prefix_length\":40,\"query\":\"Folder\"}}}}", bodyJSON)
}

func Test_Match_Bool_Prefix_should_create_json_with_match_field_inside_query(t *testing.T) {
	// Given
	query := es.NewQuery(
		es.MatchBoolPrefix("message", "this is a test").
			Analyzer("standart").
			Boost(2.14).
			MinimumShouldMatch("99%").
			Operator(Operator.Or).
			Fuzziness("AUTO").
			FuzzyRewrite("constant_score").
			FuzzyTranspositions(true).
			MaxExpansions(50).
			PrefixLength(2),
	)

	// When Then
	assert.NotNil(t, query)
	bodyJSON := assert.MarshalWithoutError(t, query)
	// nolint:golint,lll
	assert.Equal(t, "{\"query\":{\"match_bool_prefix\":{\"message\":{\"analyzer\":\"standart\",\"boost\":2.14,\"fuzziness\":\"AUTO\",\"fuzzy_rewrite\":\"constant_score\",\"fuzzy_transpositions\":true,\"max_expansions\":50,\"minimum_should_match\":\"99%\",\"operator\":\"or\",\"prefix_length\":2,\"query\":\"this is a test\"}}}}", bodyJSON)
}
