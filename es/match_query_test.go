package es_test

import (
	"testing"

	"github.com/Trendyol/es-query-builder/es"
	"github.com/Trendyol/es-query-builder/test/assert"

	Operator "github.com/Trendyol/es-query-builder/es/enums/match/operator"
	ZeroTermsQuery "github.com/Trendyol/es-query-builder/es/enums/match/zero-terms-query"
)

////   Match   ////

func Test_Match_should_exist_on_es_package(t *testing.T) {
	// Given When Then
	assert.NotNil(t, es.Match[any])
}

func Test_Match_method_should_create_matchType(t *testing.T) {
	// Given
	b := es.Match("key", "value")

	// Then
	assert.NotNil(t, b)
	assert.IsTypeString(t, "es.matchType", b)
}

func Test_Match_should_have_Boost_method(t *testing.T) {
	// Given
	term := es.Match("key", "value")

	// When Then
	assert.NotNil(t, term.Boost)
}

func Test_Match_Boost_should_create_json_with_boost_field_inside_match(t *testing.T) {
	// Given
	query := es.NewQuery(
		es.Match("type", "Folder").
			Boost(3.14),
	)

	// When Then
	assert.NotNil(t, query)
	bodyJSON := assert.MarshalWithoutError(t, query)
	assert.Equal(t, "{\"query\":{\"match\":{\"type\":{\"boost\":3.14,\"query\":\"Folder\"}}}}", bodyJSON)
}

func Test_Match_should_have_Operator_method(t *testing.T) {
	// Given
	term := es.Match("key", "value")

	// When Then
	assert.NotNil(t, term.Operator)
}

func Test_Match_Operator_should_create_json_with_operator_field_inside_match(t *testing.T) {
	// Given
	query := es.NewQuery(
		es.Match("type", "Folder").
			Operator(Operator.And),
	)

	// When Then
	assert.NotNil(t, query)
	bodyJSON := assert.MarshalWithoutError(t, query)
	assert.Equal(t, "{\"query\":{\"match\":{\"type\":{\"operator\":\"and\",\"query\":\"Folder\"}}}}", bodyJSON)
}

func Test_Match_should_have_CutoffFrequency_method(t *testing.T) {
	// Given
	term := es.Match("key", "value")

	// When Then
	assert.NotNil(t, term.CutoffFrequency)
}

func Test_Match_CutoffFrequency_should_create_json_with_cutoff_frequency_field_inside_match(t *testing.T) {
	// Given
	query := es.NewQuery(
		es.Match("type", "Folder").
			CutoffFrequency(0.0001),
	)

	// When Then
	assert.NotNil(t, query)
	bodyJSON := assert.MarshalWithoutError(t, query)
	assert.Equal(t, "{\"query\":{\"match\":{\"type\":{\"cutoff_frequency\":0.0001,\"query\":\"Folder\"}}}}", bodyJSON)
}

func Test_Match_should_have_Fuzziness_method(t *testing.T) {
	// Given
	term := es.Match("key", "value")

	// When Then
	assert.NotNil(t, term.Fuzziness)
}

func Test_Match_Fuzziness_should_create_json_with_fuzziness_field_inside_match(t *testing.T) {
	// Given
	query := es.NewQuery(
		es.Match("type", "Folder").
			Fuzziness("AUTO"),
	)

	// When Then
	assert.NotNil(t, query)
	bodyJSON := assert.MarshalWithoutError(t, query)
	assert.Equal(t, "{\"query\":{\"match\":{\"type\":{\"fuzziness\":\"AUTO\",\"query\":\"Folder\"}}}}", bodyJSON)
}

func Test_Match_should_have_FuzzyRewrite_method(t *testing.T) {
	// Given
	term := es.Match("key", "value")

	// When Then
	assert.NotNil(t, term.FuzzyRewrite)
}

func Test_Match_FuzzyRewrite_should_create_json_with_fuzzy_rewrite_field_inside_match(t *testing.T) {
	// Given
	query := es.NewQuery(
		es.Match("type", "Folder").
			FuzzyRewrite("constant_score"),
	)

	// When Then
	assert.NotNil(t, query)
	bodyJSON := assert.MarshalWithoutError(t, query)
	assert.Equal(t, "{\"query\":{\"match\":{\"type\":{\"fuzzy_rewrite\":\"constant_score\",\"query\":\"Folder\"}}}}", bodyJSON)
}

func Test_Match_should_have_FuzzyTranspositions_method(t *testing.T) {
	// Given
	term := es.Match("key", "value")

	// When Then
	assert.NotNil(t, term.FuzzyTranspositions)
}

func Test_Match_FuzzyTranspositions_should_create_json_with_fuzzy_transpositions_field_inside_match(t *testing.T) {
	// Given
	query := es.NewQuery(
		es.Match("type", "Directory").
			FuzzyTranspositions(false),
	)

	// When Then
	assert.NotNil(t, query)
	bodyJSON := assert.MarshalWithoutError(t, query)
	assert.Equal(t, "{\"query\":{\"match\":{\"type\":{\"fuzzy_transpositions\":false,\"query\":\"Directory\"}}}}", bodyJSON)
}

func Test_Match_should_have_Lenient_method(t *testing.T) {
	// Given
	term := es.Match("key", "value")

	// When Then
	assert.NotNil(t, term.Lenient)
}

func Test_Match_Lenient_should_create_json_with_lenient_field_inside_match(t *testing.T) {
	// Given
	query := es.NewQuery(
		es.Match("type", "Folder").
			Lenient(true),
	)

	// When Then
	assert.NotNil(t, query)
	bodyJSON := assert.MarshalWithoutError(t, query)
	assert.Equal(t, "{\"query\":{\"match\":{\"type\":{\"lenient\":true,\"query\":\"Folder\"}}}}", bodyJSON)
}

func Test_Match_should_have_MaxExpansions_method(t *testing.T) {
	// Given
	term := es.Match("key", "value")

	// When Then
	assert.NotNil(t, term.MaxExpansions)
}

func Test_Match_MaxExpansions_should_create_json_with_max_expansions_field_inside_match(t *testing.T) {
	// Given
	query := es.NewQuery(
		es.Match("type", "Folder").
			MaxExpansions(99),
	)

	// When Then
	assert.NotNil(t, query)
	bodyJSON := assert.MarshalWithoutError(t, query)
	assert.Equal(t, "{\"query\":{\"match\":{\"type\":{\"max_expansions\":99,\"query\":\"Folder\"}}}}", bodyJSON)
}

func Test_Match_should_have_PrefixLength_method(t *testing.T) {
	// Given
	term := es.Match("key", "value")

	// When Then
	assert.NotNil(t, term.PrefixLength)
}

func Test_Match_PrefixLength_should_create_json_with_prefix_length_field_inside_match(t *testing.T) {
	// Given
	query := es.NewQuery(
		es.Match("type", "Folder").
			PrefixLength(40),
	)

	// When Then
	assert.NotNil(t, query)
	bodyJSON := assert.MarshalWithoutError(t, query)
	assert.Equal(t, "{\"query\":{\"match\":{\"type\":{\"prefix_length\":40,\"query\":\"Folder\"}}}}", bodyJSON)
}

func Test_Match_should_have_AutoGenerateSynonymsPhraseQuery_method(t *testing.T) {
	// Given
	term := es.Match("key", "value")

	// When Then
	assert.NotNil(t, term.AutoGenerateSynonymsPhraseQuery)
}

func Test_Match_AutoGenerateSynonymsPhraseQuery_should_create_json_with_auto_generate_synonyms_phrase_query_field_inside_match(t *testing.T) {
	// Given
	query := es.NewQuery(
		es.Match("type", "Folder").
			AutoGenerateSynonymsPhraseQuery(false),
	)

	// When Then
	assert.NotNil(t, query)
	bodyJSON := assert.MarshalWithoutError(t, query)
	assert.Equal(t, "{\"query\":{\"match\":{\"type\":{\"auto_generate_synonyms_phrase_query\":false,\"query\":\"Folder\"}}}}", bodyJSON)
}

func Test_Match_should_have_ZeroTermsQuery_method(t *testing.T) {
	// Given
	term := es.Match("key", "value")

	// When Then
	assert.NotNil(t, term.ZeroTermsQuery)
}

func Test_Match_ZeroTermsQuery_should_create_json_with_zero_terms_query_field_inside_match(t *testing.T) {
	// Given
	query := es.NewQuery(
		es.Match("type", "Folder").
			ZeroTermsQuery(ZeroTermsQuery.All),
	)

	// When Then
	assert.NotNil(t, query)
	bodyJSON := assert.MarshalWithoutError(t, query)
	assert.Equal(t, "{\"query\":{\"match\":{\"type\":{\"query\":\"Folder\",\"zero_terms_query\":\"all\"}}}}", bodyJSON)
}

func Test_Match_should_create_json_with_match_field_inside_query(t *testing.T) {
	// Given
	query := es.NewQuery(
		es.Match("message", "this is a test").
			Boost(2.14).
			Operator(Operator.Or).
			CutoffFrequency(0.241).
			Fuzziness("AUTO").
			FuzzyRewrite("constant_score").
			FuzzyTranspositions(true).
			Lenient(true).
			MaxExpansions(50).
			PrefixLength(2).
			AutoGenerateSynonymsPhraseQuery(true).
			ZeroTermsQuery(ZeroTermsQuery.None),
	)

	// When Then
	assert.NotNil(t, query)
	bodyJSON := assert.MarshalWithoutError(t, query)
	// nolint:golint,lll
	assert.Equal(t, "{\"query\":{\"match\":{\"message\":{\"auto_generate_synonyms_phrase_query\":true,\"boost\":2.14,\"cutoff_frequency\":0.241,\"fuzziness\":\"AUTO\",\"fuzzy_rewrite\":\"constant_score\",\"fuzzy_transpositions\":true,\"lenient\":true,\"max_expansions\":50,\"operator\":\"or\",\"prefix_length\":2,\"query\":\"this is a test\",\"zero_terms_query\":\"none\"}}}}", bodyJSON)
}
