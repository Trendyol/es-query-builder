package es_test

import (
	"testing"

	ScoreMode "github.com/Trendyol/es-query-builder/es/enums/score-mode"

	"github.com/Trendyol/es-query-builder/es"
	"github.com/Trendyol/es-query-builder/test/assert"
)

////    Nested    ////

func Test_Nested_should_exist_on_es_package(t *testing.T) {
	t.Parallel()
	// Given When Then
	assert.NotNil(t, es.Nested[any])
}

func Test_Nested_method_should_create_nestedType(t *testing.T) {
	t.Parallel()
	// Given
	n := es.Nested("path", es.Object{})

	// Then
	assert.NotNil(t, n)
	assert.IsTypeString(t, "es.nestedType", n)
}

func Test_Nested_should_create_query_json_with_nested_field_inside(t *testing.T) {
	t.Parallel()
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
	t.Parallel()
	// Given
	n := es.Nested("path", es.Object{})

	// When Then
	assert.NotNil(t, n.InnerHits)
}

func Test_InnerHits_should_add_inner_hits_field_into_Nested(t *testing.T) {
	t.Parallel()
	// Given
	query := es.NewQuery(
		es.Nested("nested.path", es.Object{}).
			InnerHits(
				es.InnerHits().
					Size(1_000).
					From(5_000),
			),
	)

	// When Then
	assert.NotNil(t, query)
	bodyJSON := assert.MarshalWithoutError(t, query)
	assert.Equal(t, "{\"query\":{\"nested\":{\"inner_hits\":{\"from\":5000,\"size\":1000},\"path\":\"nested.path\",\"query\":{}}}}", bodyJSON)
}

func Test_Nested_should_have_ScoreMode_method(t *testing.T) {
	t.Parallel()
	// Given
	n := es.Nested("path", es.Object{})

	// When Then
	assert.NotNil(t, n.ScoreMode)
}

func Test_ScoreMod_should_add_score_mode_field_into_Nested(t *testing.T) {
	t.Parallel()
	// Given
	query := es.NewQuery(
		es.Nested("nested.path", es.Object{}).ScoreMode(ScoreMode.Sum),
	)

	// When Then
	assert.NotNil(t, query)
	bodyJSON := assert.MarshalWithoutError(t, query)
	assert.Equal(t, "{\"query\":{\"nested\":{\"path\":\"nested.path\",\"query\":{},\"score_mode\":\"sum\"}}}", bodyJSON)
}

func Test_Nested_should_have_Boost_method(t *testing.T) {
	t.Parallel()
	// Given
	n := es.Nested("path", es.Object{})

	// When Then
	assert.NotNil(t, n.Boost)
}

func Test_Boost_should_add_boost_field_into_Nested(t *testing.T) {
	t.Parallel()
	// Given
	query := es.NewQuery(
		es.Nested("nested.path", es.Object{}).Boost(5.56),
	)

	// When Then
	assert.NotNil(t, query)
	bodyJSON := assert.MarshalWithoutError(t, query)
	assert.Equal(t, "{\"query\":{\"nested\":{\"boost\":5.56,\"path\":\"nested.path\",\"query\":{}}}}", bodyJSON)
}

func Test_Nested_should_have_IgnoreUnmapped_method(t *testing.T) {
	t.Parallel()
	// Given
	n := es.Nested("path", es.Object{})

	// When Then
	assert.NotNil(t, n.IgnoreUnmapped)
}

func Test_IgnoreUnmapped_should_add_ignore_unmapped_field_into_Nested(t *testing.T) {
	t.Parallel()
	// Given
	query := es.NewQuery(
		es.Nested("nested.path", es.Object{}).IgnoreUnmapped(true),
	)

	// When Then
	assert.NotNil(t, query)
	bodyJSON := assert.MarshalWithoutError(t, query)
	assert.Equal(t, "{\"query\":{\"nested\":{\"ignore_unmapped\":true,\"path\":\"nested.path\",\"query\":{}}}}", bodyJSON)
}
