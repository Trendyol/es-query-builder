package es_test

import (
	"testing"

	ZeroTermsQuery "github.com/Trendyol/es-query-builder/es/enums/zero-terms-query"

	"github.com/Trendyol/es-query-builder/es"
	"github.com/Trendyol/es-query-builder/test/assert"
)

////   Match Phrase   ////

func Test_MatchPhrase_should_exist_on_es_package(t *testing.T) {
	// Given When Then
	assert.NotNil(t, es.MatchPhrase[any])
}

func Test_MatchPhrase_method_should_create_matchPhraseType(t *testing.T) {
	// Given
	b := es.MatchPhrase("key", "value")

	// Then
	assert.NotNil(t, b)
	assert.IsTypeString(t, "es.matchPhraseType", b)
}

func Test_MatchPhrase_should_have_Analyzer_method(t *testing.T) {
	// Given
	match := es.MatchPhrase("key", "value")

	// When Then
	assert.NotNil(t, match.Analyzer)
}

func Test_MatchPhrase_Analyzer_should_create_json_with_analyzer_field_inside_match_phrase(t *testing.T) {
	// Given
	query := es.NewQuery(
		es.MatchPhrase("type", "Folder").
			Analyzer("standart"),
	)

	// When Then
	assert.NotNil(t, query)
	bodyJSON := assert.MarshalWithoutError(t, query)
	assert.Equal(t, "{\"query\":{\"match_phrase\":{\"type\":{\"analyzer\":\"standart\",\"query\":\"Folder\"}}}}", bodyJSON)
}

func Test_MatchPhrase_should_have_Boost_method(t *testing.T) {
	// Given
	match := es.MatchPhrase("key", "value")

	// When Then
	assert.NotNil(t, match.Boost)
}

func Test_MatchPhrase_Boost_should_create_json_with_boost_field_inside_match_phrase(t *testing.T) {
	// Given
	query := es.NewQuery(
		es.MatchPhrase("type", "Folder").
			Boost(3.14),
	)

	// When Then
	assert.NotNil(t, query)
	bodyJSON := assert.MarshalWithoutError(t, query)
	assert.Equal(t, "{\"query\":{\"match_phrase\":{\"type\":{\"boost\":3.14,\"query\":\"Folder\"}}}}", bodyJSON)
}

func Test_MatchPhrase_should_have_Slop_method(t *testing.T) {
	// Given
	match := es.MatchPhrase("key", "value")

	// When Then
	assert.NotNil(t, match.Slop)
}

func Test_MatchPhrase_Slop_should_create_json_with_slop_field_inside_match_phrase(t *testing.T) {
	// Given
	query := es.NewQuery(
		es.MatchPhrase("type", "Folder").
			Slop(3),
	)

	// When Then
	assert.NotNil(t, query)
	bodyJSON := assert.MarshalWithoutError(t, query)
	assert.Equal(t, "{\"query\":{\"match_phrase\":{\"type\":{\"query\":\"Folder\",\"slop\":3}}}}", bodyJSON)
}

func Test_MatchPhrase_should_have_ZeroTermsQuery_method(t *testing.T) {
	// Given
	match := es.MatchPhrase("key", "value")

	// When Then
	assert.NotNil(t, match.ZeroTermsQuery)
}

func Test_MatchPhrase_ZeroTermsQuery_should_create_json_with_zero_terms_query_field_inside_match_phrase(t *testing.T) {
	// Given
	query := es.NewQuery(
		es.MatchPhrase("type", "Folder").
			ZeroTermsQuery(ZeroTermsQuery.All),
	)

	// When Then
	assert.NotNil(t, query)
	bodyJSON := assert.MarshalWithoutError(t, query)
	assert.Equal(t, "{\"query\":{\"match_phrase\":{\"type\":{\"query\":\"Folder\",\"zero_terms_query\":\"all\"}}}}", bodyJSON)
}

func Test_MatchPhrase_should_create_json_with_match_phrase_field_inside_query(t *testing.T) {
	// Given
	query := es.NewQuery(
		es.MatchPhrase("message", "this is a test").
			Analyzer("standart").
			Boost(2.14).
			Slop(9).
			ZeroTermsQuery(ZeroTermsQuery.None),
	)

	// When Then
	assert.NotNil(t, query)
	bodyJSON := assert.MarshalWithoutError(t, query)
	// nolint:golint,lll
	assert.Equal(t, "{\"query\":{\"match_phrase\":{\"message\":{\"analyzer\":\"standart\",\"boost\":2.14,\"query\":\"this is a test\",\"slop\":9,\"zero_terms_query\":\"none\"}}}}", bodyJSON)
}
