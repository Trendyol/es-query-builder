package es_test

import (
	"testing"

	"github.com/GokselKUCUKSAHIN/es-query-builder/es"
	"github.com/GokselKUCUKSAHIN/es-query-builder/test/assert"

	ScoreMode "github.com/GokselKUCUKSAHIN/es-query-builder/es/enums/nested/score-mode"
)

////    Nested    ////

func Test_Nested_should_exist_on_es_package(t *testing.T) {
	// Given When Then
	assert.NotNil(t, es.Nested[any])
}

func Test_Nested_method_should_create_nestedType(t *testing.T) {
	// Given
	n := es.Nested("path", es.Object{})

	// Then
	assert.NotNil(t, n)
	assert.IsTypeString(t, "es.nestedType", n)
}

func Test_Nested_should_create_query_json_with_nested_field_inside(t *testing.T) {
	// Given
	query := es.NewQuery(
		es.Nested("nested.path",
			es.Object{},
		),
	)

	// When Then
	assert.NotNil(t, query)
	bodyJSON := assert.MarshalWithoutError(t, query)
	assert.Equal(t, "{\"query\":{\"nested\":{\"path\":\"nested.path\",\"query\":{}}}}", bodyJSON)
}

func Test_Nested_should_have_InnerHits_method(t *testing.T) {
	// Given
	n := es.Nested("path", es.Object{})

	// When Then
	assert.NotNil(t, n.InnerHits)
}

func Test_InnerHits_should_add_inner_hits_field_into_Nested(t *testing.T) {
	// Given
	query := es.NewQuery(
		es.Nested("nested.path", es.Object{}).InnerHits(es.Object{"inner": "hits"}),
	)

	// When Then
	assert.NotNil(t, query)
	bodyJSON := assert.MarshalWithoutError(t, query)
	assert.Equal(t, "{\"query\":{\"nested\":{\"inner_hits\":{\"inner\":\"hits\"},\"path\":\"nested.path\",\"query\":{}}}}", bodyJSON)
}

func Test_Nested_should_have_ScoreMode_method(t *testing.T) {
	// Given
	n := es.Nested("path", es.Object{})

	// When Then
	assert.NotNil(t, n.ScoreMode)
}

func Test_ScoreMod_should_add_score_mode_field_into_Nested(t *testing.T) {
	// Given
	query := es.NewQuery(
		es.Nested("nested.path", es.Object{}).ScoreMode(ScoreMode.Sum),
	)

	// When Then
	assert.NotNil(t, query)
	bodyJSON := assert.MarshalWithoutError(t, query)
	assert.Equal(t, "{\"query\":{\"nested\":{\"path\":\"nested.path\",\"query\":{},\"score_mode\":\"sum\"}}}", bodyJSON)
}
