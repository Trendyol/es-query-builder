package es_test

import (
	"testing"

	ZeroTermsQuery "github.com/Trendyol/es-query-builder/es/enums/zero-terms-query"

	"github.com/Trendyol/es-query-builder/es"
	"github.com/Trendyol/es-query-builder/test/assert"
)

////   Match Phrase Prefix   ////

func Test_MatchPhrasePrefix_should_exist_on_es_package(t *testing.T) {
	t.Parallel()
	// Given When Then
	assert.NotNil(t, es.MatchPhrasePrefix[any])
}

func Test_MatchPhrasePrefix_method_should_create_matchPhrasePrefixType(t *testing.T) {
	t.Parallel()
	// Given
	b := es.MatchPhrasePrefix("key", "value")

	// Then
	assert.NotNil(t, b)
	assert.IsTypeString(t, "es.matchPhrasePrefixType", b)
}

func Test_MatchPhrasePrefix_should_have_Analyzer_method(t *testing.T) {
	t.Parallel()
	// Given
	match := es.MatchPhrasePrefix("key", "value")

	// When Then
	assert.NotNil(t, match.Analyzer)
}

func Test_MatchPhrasePrefix_Analyzer_should_create_json_with_analyzer_field_inside_match_phrase_prefix(t *testing.T) {
	t.Parallel()
	// Given
	query := es.NewQuery(
		es.MatchPhrasePrefix("type", "Folder").
			Analyzer("standart"),
	)

	// When Then
	assert.NotNil(t, query)
	bodyJSON := assert.MarshalWithoutError(t, query)
	assert.Equal(t, "{\"query\":{\"match_phrase_prefix\":{\"type\":{\"analyzer\":\"standart\",\"query\":\"Folder\"}}}}", bodyJSON)
}

func Test_MatchPhrasePrefix_should_have_Boost_method(t *testing.T) {
	t.Parallel()
	// Given
	match := es.MatchPhrasePrefix("key", "value")

	// When Then
	assert.NotNil(t, match.Boost)
}

func Test_MatchPhrasePrefix_Boost_should_create_json_with_boost_field_inside_match_phrase_prefix(t *testing.T) {
	t.Parallel()
	// Given
	query := es.NewQuery(
		es.MatchPhrasePrefix("type", "Folder").
			Boost(3.14),
	)

	// When Then
	assert.NotNil(t, query)
	bodyJSON := assert.MarshalWithoutError(t, query)
	assert.Equal(t, "{\"query\":{\"match_phrase_prefix\":{\"type\":{\"boost\":3.14,\"query\":\"Folder\"}}}}", bodyJSON)
}

func Test_MatchPhrasePrefix_should_have_MaxExpansions_method(t *testing.T) {
	t.Parallel()
	// Given
	match := es.MatchPhrasePrefix("key", "value")

	// When Then
	assert.NotNil(t, match.MaxExpansions)
}

func Test_MatchPhrasePrefix_MaxExpansions_should_create_json_with_max_expansions_field_inside_match_phrase_prefix(t *testing.T) {
	t.Parallel()
	// Given
	query := es.NewQuery(
		es.MatchPhrasePrefix("type", "Folder").
			MaxExpansions(7),
	)

	// When Then
	assert.NotNil(t, query)
	bodyJSON := assert.MarshalWithoutError(t, query)
	assert.Equal(t, "{\"query\":{\"match_phrase_prefix\":{\"type\":{\"max_expansions\":7,\"query\":\"Folder\"}}}}", bodyJSON)
}

func Test_MatchPhrasePrefix_should_have_Slop_method(t *testing.T) {
	t.Parallel()
	// Given
	match := es.MatchPhrasePrefix("key", "value")

	// When Then
	assert.NotNil(t, match.Slop)
}

func Test_MatchPhrasePrefix_Slop_should_create_json_with_slop_field_inside_match_phrase_prefix(t *testing.T) {
	t.Parallel()
	// Given
	query := es.NewQuery(
		es.MatchPhrasePrefix("type", "Folder").
			Slop(3),
	)

	// When Then
	assert.NotNil(t, query)
	bodyJSON := assert.MarshalWithoutError(t, query)
	assert.Equal(t, "{\"query\":{\"match_phrase_prefix\":{\"type\":{\"query\":\"Folder\",\"slop\":3}}}}", bodyJSON)
}

func Test_MatchPhrasePrefix_should_have_ZeroTermsQuery_method(t *testing.T) {
	t.Parallel()
	// Given
	match := es.MatchPhrasePrefix("key", "value")

	// When Then
	assert.NotNil(t, match.ZeroTermsQuery)
}

func Test_MatchPhrasePrefix_ZeroTermsQuery_should_create_json_with_zero_terms_query_field_inside_match_phrase_prefix(t *testing.T) {
	t.Parallel()
	// Given
	query := es.NewQuery(
		es.MatchPhrasePrefix("type", "Folder").
			ZeroTermsQuery(ZeroTermsQuery.All),
	)

	// When Then
	assert.NotNil(t, query)
	bodyJSON := assert.MarshalWithoutError(t, query)
	assert.Equal(t, "{\"query\":{\"match_phrase_prefix\":{\"type\":{\"query\":\"Folder\",\"zero_terms_query\":\"all\"}}}}", bodyJSON)
}

func Test_MatchPhrasePrefix_should_create_json_with_match_phrase_prefix_field_inside_query(t *testing.T) {
	t.Parallel()
	// Given
	query := es.NewQuery(
		es.MatchPhrasePrefix("message", "this is a test").
			Analyzer("standart").
			Boost(2.14).
			MaxExpansions(50).
			Slop(9).
			ZeroTermsQuery(ZeroTermsQuery.None),
	)

	// When Then
	assert.NotNil(t, query)
	bodyJSON := assert.MarshalWithoutError(t, query)
	// nolint:golint,lll
	assert.Equal(t, "{\"query\":{\"match_phrase_prefix\":{\"message\":{\"analyzer\":\"standart\",\"boost\":2.14,\"max_expansions\":50,\"query\":\"this is a test\",\"slop\":9,\"zero_terms_query\":\"none\"}}}}", bodyJSON)
}
