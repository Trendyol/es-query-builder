package es_test

import (
	"testing"

	"github.com/GokselKUCUKSAHIN/es-query-builder/es"
	"github.com/GokselKUCUKSAHIN/es-query-builder/test/assert"

	Operator "github.com/GokselKUCUKSAHIN/es-query-builder/es/enums/match/operator"
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

func Test_Match_should_create_json_with_match_field_inside_query(t *testing.T) {
	// Given
	query := es.NewQuery(
		es.Match("message", "this is a test").
			Boost(2.14).
			Operator(Operator.Or),
	)

	// When Then
	assert.NotNil(t, query)
	bodyJSON := assert.MarshalWithoutError(t, query)
	// nolint:golint,lll
	assert.Equal(t, "{\"query\":{\"match\":{\"message\":{\"boost\":2.14,\"operator\":\"or\",\"query\":\"this is a test\"}}}}", bodyJSON)
}
