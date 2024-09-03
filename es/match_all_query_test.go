package es_test

import (
	"testing"

	"github.com/Trendyol/es-query-builder/es"
	"github.com/Trendyol/es-query-builder/test/assert"
)

////    MatchAll    ////

func Test_MatchAll_should_exist_on_es_package(t *testing.T) {
	// Given When Then
	assert.NotNil(t, es.MatchAll)
}

func Test_MatchAll_method_should_create_matchAllType(t *testing.T) {
	// Given
	b := es.MatchAll()

	// Then
	assert.NotNil(t, b)
	assert.IsTypeString(t, "es.matchAllType", b)
}

func Test_MatchAll_should_create_json_with_match_all_field_inside_query(t *testing.T) {
	// Given
	query := es.NewQuery(
		es.MatchAll().
			Boost(2.14),
	)

	// When Then
	assert.NotNil(t, query)
	bodyJSON := assert.MarshalWithoutError(t, query)
	assert.Equal(t, "{\"query\":{\"match_all\":{\"boost\":2.14}}}", bodyJSON)
}
