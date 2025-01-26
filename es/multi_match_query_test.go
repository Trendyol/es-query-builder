package es_test

import (
	"testing"

	Operator "github.com/Trendyol/es-query-builder/es/enums/operator"
	TextQueryType "github.com/Trendyol/es-query-builder/es/enums/text-query-type"
	ZeroTermsQuery "github.com/Trendyol/es-query-builder/es/enums/zero-terms-query"

	"github.com/Trendyol/es-query-builder/es"
	"github.com/Trendyol/es-query-builder/test/assert"
)

////   Multi Match   ////

func Test_Multi_Match_should_exist_on_es_package(t *testing.T) {
	// Given When Then
	assert.NotNil(t, es.MultiMatch[any])
}

func Test_Multi_Match_method_should_create_multiMatchType(t *testing.T) {
	// Given
	b := es.MultiMatch("value")

	// Then
	assert.NotNil(t, b)
	assert.IsTypeString(t, "es.multiMatchType", b)
}

func Test_Multi_Match_should_have_Boost_method(t *testing.T) {
	// Given
	match := es.MultiMatch("value")

	// When Then
	assert.NotNil(t, match.Boost)
}

func Test_Multi_Match_Boost_should_create_json_with_boost_field_inside_match(t *testing.T) {
	// Given
	query := es.NewQuery(
		es.MultiMatch("Folder").
			Boost(3.14),
	)

	// When Then
	assert.NotNil(t, query)
	bodyJSON := assert.MarshalWithoutError(t, query)
	assert.Equal(t, "{\"query\":{\"multi_match\":{\"boost\":3.14,\"query\":\"Folder\"}}}", bodyJSON)
}

func Test_Multi_Match_should_have_Analyzer_method(t *testing.T) {
	// Given
	match := es.MultiMatch("value")

	// When Then
	assert.NotNil(t, match.Analyzer)
}

func Test_Multi_Match_Analyzer_should_create_json_with_analyzer_field_inside_match(t *testing.T) {
	// Given
	query := es.NewQuery(
		es.MultiMatch("Folder").
			Analyzer("standard"),
	)

	// When Then
	assert.NotNil(t, query)
	bodyJSON := assert.MarshalWithoutError(t, query)
	assert.Equal(t, "{\"query\":{\"multi_match\":{\"analyzer\":\"standard\",\"query\":\"Folder\"}}}", bodyJSON)
}

func Test_Multi_Match_should_have_Operator_method(t *testing.T) {
	// Given
	match := es.MultiMatch("value")

	// When Then
	assert.NotNil(t, match.Operator)
}

func Test_Multi_Match_Operator_should_create_json_with_operator_field_inside_match(t *testing.T) {
	// Given
	query := es.NewQuery(
		es.MultiMatch("Folder").
			Operator(Operator.And),
	)

	// When Then
	assert.NotNil(t, query)
	bodyJSON := assert.MarshalWithoutError(t, query)
	assert.Equal(t, "{\"query\":{\"multi_match\":{\"operator\":\"and\",\"query\":\"Folder\"}}}", bodyJSON)
}

func Test_Multi_Match_should_have_CutoffFrequency_method(t *testing.T) {
	// Given
	match := es.MultiMatch("value")

	// When Then
	assert.NotNil(t, match.CutoffFrequency)
}

func Test_Multi_Match_CutoffFrequency_should_create_json_with_cutoff_frequency_field_inside_match(t *testing.T) {
	// Given
	query := es.NewQuery(
		es.MultiMatch("Folder").
			CutoffFrequency(0.0001),
	)

	// When Then
	assert.NotNil(t, query)
	bodyJSON := assert.MarshalWithoutError(t, query)
	assert.Equal(t, "{\"query\":{\"multi_match\":{\"cutoff_frequency\":0.0001,\"query\":\"Folder\"}}}", bodyJSON)
}

func Test_Multi_Match_should_have_Fields_method(t *testing.T) {
	// Given
	match := es.MultiMatch("value")

	// When Then
	assert.NotNil(t, match.Fields)
}

func Test_Multi_Match_Fields_should_create_json_with_fields_field_inside_match(t *testing.T) {
	// Given
	query := es.NewQuery(
		es.MultiMatch("Folder").
			Fields("title", "description"),
	)

	// When Then
	assert.NotNil(t, query)
	bodyJSON := assert.MarshalWithoutError(t, query)
	assert.Equal(t, "{\"query\":{\"multi_match\":{\"fields\":[\"title\",\"description\"],\"query\":\"Folder\"}}}", bodyJSON)
}

func Test_Multi_Match_should_have_Fuzziness_method(t *testing.T) {
	// Given
	match := es.MultiMatch("value")

	// When Then
	assert.NotNil(t, match.Fuzziness)
}

func Test_Multi_Match_Fuzziness_should_create_json_with_fuzziness_field_inside_match(t *testing.T) {
	// Given
	query := es.NewQuery(
		es.MultiMatch("Folder").
			Fuzziness("AUTO"),
	)

	// When Then
	assert.NotNil(t, query)
	bodyJSON := assert.MarshalWithoutError(t, query)
	assert.Equal(t, "{\"query\":{\"multi_match\":{\"fuzziness\":\"AUTO\",\"query\":\"Folder\"}}}", bodyJSON)
}

func Test_Multi_Match_should_have_Slop_method(t *testing.T) {
	// Given
	match := es.MultiMatch("value")

	// When Then
	assert.NotNil(t, match.Slop)
}

func Test_Multi_Match_Slop_should_create_json_with_slop_field_inside_match(t *testing.T) {
	// Given
	query := es.NewQuery(
		es.MultiMatch("Folder").
			Slop(5),
	)

	// When Then
	assert.NotNil(t, query)
	bodyJSON := assert.MarshalWithoutError(t, query)
	assert.Equal(t, "{\"query\":{\"multi_match\":{\"query\":\"Folder\",\"slop\":5}}}", bodyJSON)
}

func Test_Multi_Match_should_have_MinimumShouldMatch_method(t *testing.T) {
	// Given
	match := es.MultiMatch("value")

	// When Then
	assert.NotNil(t, match.MinimumShouldMatch)
}

func Test_Multi_Match_MinimumShouldMatch_should_create_json_with_minimum_should_match_field_inside_match(t *testing.T) {
	// Given
	query := es.NewQuery(
		es.MultiMatch("Folder").
			MinimumShouldMatch(3),
	)

	// When Then
	assert.NotNil(t, query)
	bodyJSON := assert.MarshalWithoutError(t, query)
	assert.Equal(t, "{\"query\":{\"multi_match\":{\"minimum_should_match\":3,\"query\":\"Folder\"}}}", bodyJSON)
}

func Test_Multi_Match_should_have_FuzzyRewrite_method(t *testing.T) {
	// Given
	match := es.MultiMatch("value")

	// When Then
	assert.NotNil(t, match.FuzzyRewrite)
}

func Test_Multi_Match_FuzzyRewrite_should_create_json_with_fuzzy_rewrite_field_inside_match(t *testing.T) {
	// Given
	query := es.NewQuery(
		es.MultiMatch("Folder").
			FuzzyRewrite("constant_score"),
	)

	// When Then
	assert.NotNil(t, query)
	bodyJSON := assert.MarshalWithoutError(t, query)
	assert.Equal(t, "{\"query\":{\"multi_match\":{\"fuzzy_rewrite\":\"constant_score\",\"query\":\"Folder\"}}}", bodyJSON)
}

func Test_Multi_Match_should_have_FuzzyTranspositions_method(t *testing.T) {
	// Given
	match := es.MultiMatch("value")

	// When Then
	assert.NotNil(t, match.FuzzyTranspositions)
}

func Test_Multi_Match_FuzzyTranspositions_should_create_json_with_fuzzy_transpositions_field_inside_match(t *testing.T) {
	// Given
	query := es.NewQuery(
		es.MultiMatch("Directory").
			FuzzyTranspositions(false),
	)

	// When Then
	assert.NotNil(t, query)
	bodyJSON := assert.MarshalWithoutError(t, query)
	assert.Equal(t, "{\"query\":{\"multi_match\":{\"fuzzy_transpositions\":false,\"query\":\"Directory\"}}}", bodyJSON)
}

func Test_Multi_Match_should_have_Lenient_method(t *testing.T) {
	// Given
	match := es.MultiMatch("value")

	// When Then
	assert.NotNil(t, match.Lenient)
}

func Test_Multi_Match_Lenient_should_create_json_with_lenient_field_inside_match(t *testing.T) {
	// Given
	query := es.NewQuery(
		es.MultiMatch("Folder").
			Lenient(true),
	)

	// When Then
	assert.NotNil(t, query)
	bodyJSON := assert.MarshalWithoutError(t, query)
	assert.Equal(t, "{\"query\":{\"multi_match\":{\"lenient\":true,\"query\":\"Folder\"}}}", bodyJSON)
}

func Test_Multi_Match_should_have_MaxExpansions_method(t *testing.T) {
	// Given
	match := es.MultiMatch("value")

	// When Then
	assert.NotNil(t, match.MaxExpansions)
}

func Test_Multi_Match_MaxExpansions_should_create_json_with_max_expansions_field_inside_match(t *testing.T) {
	// Given
	query := es.NewQuery(
		es.MultiMatch("Folder").
			MaxExpansions(99),
	)

	// When Then
	assert.NotNil(t, query)
	bodyJSON := assert.MarshalWithoutError(t, query)
	assert.Equal(t, "{\"query\":{\"multi_match\":{\"max_expansions\":99,\"query\":\"Folder\"}}}", bodyJSON)
}

func Test_Multi_Match_should_have_PrefixLength_method(t *testing.T) {
	// Given
	match := es.MultiMatch("value")

	// When Then
	assert.NotNil(t, match.PrefixLength)
}

func Test_Multi_Match_PrefixLength_should_create_json_with_prefix_length_field_inside_match(t *testing.T) {
	// Given
	query := es.NewQuery(
		es.MultiMatch("Folder").
			PrefixLength(40),
	)

	// When Then
	assert.NotNil(t, query)
	bodyJSON := assert.MarshalWithoutError(t, query)
	assert.Equal(t, "{\"query\":{\"multi_match\":{\"prefix_length\":40,\"query\":\"Folder\"}}}", bodyJSON)
}

func Test_Multi_Match_should_have_AutoGenerateSynonymsPhraseQuery_method(t *testing.T) {
	// Given
	match := es.MultiMatch("value")

	// When Then
	assert.NotNil(t, match.AutoGenerateSynonymsPhraseQuery)
}

// nolint:golint,lll
func Test_Multi_Match_AutoGenerateSynonymsPhraseQuery_should_create_json_auto_generate_synonyms_phrase_query_field_inside_match(t *testing.T) {
	// Given
	query := es.NewQuery(
		es.MultiMatch("Folder").
			AutoGenerateSynonymsPhraseQuery(false),
	)

	// When Then
	assert.NotNil(t, query)
	bodyJSON := assert.MarshalWithoutError(t, query)
	assert.Equal(t, "{\"query\":{\"multi_match\":{\"auto_generate_synonyms_phrase_query\":false,\"query\":\"Folder\"}}}", bodyJSON)
}

func Test_Multi_Match_should_have_ZeroTermsQuery_method(t *testing.T) {
	// Given
	match := es.MultiMatch("value")

	// When Then
	assert.NotNil(t, match.ZeroTermsQuery)
}

func Test_Multi_Match_ZeroTermsQuery_should_create_json_with_zero_terms_query_field_inside_match(t *testing.T) {
	// Given
	query := es.NewQuery(
		es.MultiMatch("Folder").
			ZeroTermsQuery(ZeroTermsQuery.All),
	)

	// When Then
	assert.NotNil(t, query)
	bodyJSON := assert.MarshalWithoutError(t, query)
	assert.Equal(t, "{\"query\":{\"multi_match\":{\"query\":\"Folder\",\"zero_terms_query\":\"all\"}}}", bodyJSON)
}

func Test_Multi_Match_should_have_TieBreaker_method(t *testing.T) {
	// Given
	match := es.MultiMatch("value")

	// When Then
	assert.NotNil(t, match.TieBreaker)
}

func Test_Multi_Match_TieBreaker_should_create_json_with_tie_breaker_field_inside_match(t *testing.T) {
	// Given
	query := es.NewQuery(
		es.MultiMatch("Folder").
			TieBreaker(8.8),
	)

	// When Then
	assert.NotNil(t, query)
	bodyJSON := assert.MarshalWithoutError(t, query)
	assert.Equal(t, "{\"query\":{\"multi_match\":{\"query\":\"Folder\",\"tie_breaker\":8.8}}}", bodyJSON)
}

func Test_Multi_Match_should_have_Type_method(t *testing.T) {
	// Given
	match := es.MultiMatch("value")

	// When Then
	assert.NotNil(t, match.Type)
}

func Test_Multi_Match_Type_should_create_json_with_type_field_inside_match(t *testing.T) {
	// Given
	query := es.NewQuery(
		es.MultiMatch("Folder").
			Type(TextQueryType.Phrase),
	)

	// When Then
	assert.NotNil(t, query)
	bodyJSON := assert.MarshalWithoutError(t, query)
	assert.Equal(t, "{\"query\":{\"multi_match\":{\"query\":\"Folder\",\"type\":\"phrase\"}}}", bodyJSON)
}

func Test_Multi_Match_should_create_json_with_match_field_inside_query(t *testing.T) {
	// Given
	query := es.NewQuery(
		es.MultiMatch("this is a test").
			Analyzer("standart").
			Boost(2.14).
			Operator(Operator.Or).
			CutoffFrequency(0.241).
			Fuzziness("AUTO").
			FuzzyRewrite("constant_score").
			FuzzyTranspositions(true).
			Lenient(true).
			MaxExpansions(50).
			MinimumShouldMatch("50%").
			PrefixLength(2).
			Slop(7).
			TieBreaker(4.35).
			AutoGenerateSynonymsPhraseQuery(true).
			Type(TextQueryType.Crossfields).
			ZeroTermsQuery(ZeroTermsQuery.None),
	)

	// When Then
	assert.NotNil(t, query)
	bodyJSON := assert.MarshalWithoutError(t, query)
	// nolint:golint,lll
	assert.Equal(t, "{\"query\":{\"multi_match\":{\"analyzer\":\"standart\",\"auto_generate_synonyms_phrase_query\":true,\"boost\":2.14,\"cutoff_frequency\":0.241,\"fuzziness\":\"AUTO\",\"fuzzy_rewrite\":\"constant_score\",\"fuzzy_transpositions\":true,\"lenient\":true,\"max_expansions\":50,\"minimum_should_match\":\"50%\",\"operator\":\"or\",\"prefix_length\":2,\"query\":\"this is a test\",\"slop\":7,\"tie_breaker\":4.35,\"type\":\"cross_fields\",\"zero_terms_query\":\"none\"}}}", bodyJSON)
}
