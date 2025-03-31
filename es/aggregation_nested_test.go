package es_test

import (
	"testing"

	"github.com/Trendyol/es-query-builder/es"
	"github.com/Trendyol/es-query-builder/test/assert"
)

func Test_NestedAgg_should_exist_on_es_package(t *testing.T) {
	t.Parallel()
	// Given When Then
	assert.NotNil(t, es.NestedAgg)
}

func Test_NestedAgg_should_return_type_of_nestedAggType(t *testing.T) {
	t.Parallel()
	// Given
	aggsQuery := es.NestedAgg("price")

	// When Then
	assert.NotNil(t, aggsQuery)
	assert.IsTypeString(t, "es.nestedAggType", aggsQuery)
	assert.MarshalWithoutError(t, aggsQuery)
}

func Test_NestedAgg_should_create_json_with_terms_field_inside(t *testing.T) {
	t.Parallel()
	// Given
	a := es.NestedAgg("price")

	// When Then
	assert.NotNil(t, a)
	bodyJSON := assert.MarshalWithoutError(t, a)
	assert.Equal(t, "{\"nested\":{\"path\":\"price\"}}", bodyJSON)
}

func Test_NestedAgg_should_have_Aggs_method(t *testing.T) {
	t.Parallel()
	// Given
	a := es.NestedAgg("price")

	// When Then
	assert.NotNil(t, a.Aggs)
}

func Test_Aggs_should_add_aggs_field_into_NestedAgg_when_not_empty(t *testing.T) {
	t.Parallel()
	// Given
	a := es.NestedAgg("price").Aggs(es.Agg("nested_stock", es.NestedAgg("stock")))

	// When Then
	assert.NotNil(t, a)
	bodyJSON := assert.MarshalWithoutError(t, a)
	assert.Equal(t, "{\"aggs\":{\"nested_stock\":{\"nested\":{\"path\":\"stock\"}}},\"nested\":{\"path\":\"price\"}}", bodyJSON)
}

func Test_Aggs_should_not_add_aggs_field_into_NestedAgg_when_it_is_empty(t *testing.T) {
	t.Parallel()
	// Given
	a := es.NestedAgg("price").Aggs(nil)

	// When Then
	assert.NotNil(t, a)
	bodyJSON := assert.MarshalWithoutError(t, a)
	assert.Equal(t, "{\"nested\":{\"path\":\"price\"}}", bodyJSON)
}
