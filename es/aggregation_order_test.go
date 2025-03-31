package es_test

import (
	"testing"

	Order "github.com/Trendyol/es-query-builder/es/enums/sort/order"

	"github.com/Trendyol/es-query-builder/es"
	"github.com/Trendyol/es-query-builder/test/assert"
)

func Test_AggOrder_should_exist_on_es_package(t *testing.T) {
	t.Parallel()
	// Given When Then
	assert.NotNil(t, es.AggOrder)
}

func Test_AggOrder_should_return_type_of_avgAggType(t *testing.T) {
	t.Parallel()
	// Given
	aggsQuery := es.AggOrder("price", Order.Desc)

	// When Then
	assert.NotNil(t, aggsQuery)
	assert.IsTypeString(t, "es.aggOrder", aggsQuery)
	assert.MarshalWithoutError(t, aggsQuery)
}

func Test_AggOrder_should_create_json_with_terms_field_inside(t *testing.T) {
	t.Parallel()
	// Given
	a := es.AggOrder("price", Order.Asc)

	// When Then
	assert.NotNil(t, a)
	bodyJSON := assert.MarshalWithoutError(t, a)
	assert.Equal(t, "{\"price\":\"asc\"}", bodyJSON)
}
