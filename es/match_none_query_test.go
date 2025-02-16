package es_test

import (
	"testing"

	"github.com/Trendyol/es-query-builder/es"
	"github.com/Trendyol/es-query-builder/test/assert"
)

////   MatchNone   ////

func Test_MatchNone_should_exist_on_es_package(t *testing.T) {
	t.Parallel()
	// Given When Then
	assert.NotNil(t, es.MatchNone[any])
}

func Test_MatchNone_method_should_create_matchNoneType(t *testing.T) {
	t.Parallel()
	// Given
	b := es.MatchNone("key", "value")

	// Then
	assert.NotNil(t, b)
	assert.IsTypeString(t, "es.matchNoneType", b)
}

func Test_MatchNone_should_create_json_with_match_none_field_inside_query(t *testing.T) {
	t.Parallel()
	// Given
	query := es.NewQuery(
		es.MatchNone("fooBar", "lorem ipsum").
			Boost(6.19),
	)

	// When Then
	assert.NotNil(t, query)
	bodyJSON := assert.MarshalWithoutError(t, query)
	// nolint:golint,lll
	assert.Equal(t, "{\"query\":{\"match_none\":{\"fooBar\":{\"boost\":6.19,\"query\":\"lorem ipsum\"}}}}", bodyJSON)
}
