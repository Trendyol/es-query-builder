package es_test

import (
	"testing"

	"github.com/Trendyol/es-query-builder/es"
	"github.com/Trendyol/es-query-builder/test/assert"
)

////   FilterAgg   ////

func Test_FilterAgg_should_exist_on_es_package(t *testing.T) {
	t.Parallel()
	// Given When Then
	assert.NotNil(t, es.FilterAgg)
}

func Test_FilterAgg_method_should_create_filterAggType(t *testing.T) {
	t.Parallel()
	// Given
	agg := es.FilterAgg(es.Term("status", "active"))

	// Then
	assert.NotNil(t, agg)
	assert.IsTypeString(t, "es.filterAggType", agg)
}

func Test_FilterAgg_should_create_json_with_filter(t *testing.T) {
	t.Parallel()
	// Given
	query := es.NewQuery(nil).
		Aggs(es.Agg("active_products", es.FilterAgg(es.Term("status", "active"))))

	// When Then
	assert.NotNil(t, query)
	bodyJSON := assert.MarshalWithoutError(t, query)
	// nolint:golint,lll
	assert.Equal(t, "{\"aggs\":{\"active_products\":{\"filter\":{\"term\":{\"status\":{\"value\":\"active\"}}}}},\"query\":{}}", bodyJSON)
}

func Test_FilterAgg_should_handle_nil_filter(t *testing.T) {
	t.Parallel()
	// Given
	agg := es.FilterAgg(nil)

	// When Then
	assert.NotNil(t, agg)
	bodyJSON := assert.MarshalWithoutError(t, agg)
	assert.Equal(t, "{\"filter\":{}}", bodyJSON)
}

func Test_FilterAgg_should_handle_bool_filter(t *testing.T) {
	t.Parallel()
	// Given
	query := es.NewQuery(nil).
		Aggs(es.Agg("filtered",
			es.FilterAgg(es.Bool().Must(es.Term("status", "active"))),
		))

	// When Then
	assert.NotNil(t, query)
	bodyJSON := assert.MarshalWithoutError(t, query)
	// nolint:golint,lll
	assert.Equal(t, "{\"aggs\":{\"filtered\":{\"filter\":{\"bool\":{\"must\":[{\"term\":{\"status\":{\"value\":\"active\"}}}]}}}},\"query\":{}}", bodyJSON)
}

////   FilterAgg Aggs   ////

func Test_FilterAgg_should_have_Aggs_method(t *testing.T) {
	t.Parallel()
	// Given
	agg := es.FilterAgg(es.Term("status", "active"))

	// When Then
	assert.NotNil(t, agg.Aggs)
}

func Test_FilterAgg_Aggs_should_create_json_with_sub_aggregations(t *testing.T) {
	t.Parallel()
	// Given
	query := es.NewQuery(nil).
		Aggs(es.Agg("active_products",
			es.FilterAgg(es.Term("status", "active")).
				Aggs(es.Agg("avg_price", es.AvgAgg("price"))),
		))

	// When Then
	assert.NotNil(t, query)
	bodyJSON := assert.MarshalWithoutError(t, query)
	// nolint:golint,lll
	assert.Equal(t, "{\"aggs\":{\"active_products\":{\"aggs\":{\"avg_price\":{\"avg\":{\"field\":\"price\"}}},\"filter\":{\"term\":{\"status\":{\"value\":\"active\"}}}}},\"query\":{}}", bodyJSON)
}

////   FilterAgg Meta   ////

func Test_FilterAgg_should_have_Meta_method(t *testing.T) {
	t.Parallel()
	// Given
	agg := es.FilterAgg(es.Term("status", "active"))

	// When Then
	assert.NotNil(t, agg.Meta)
}

func Test_FilterAgg_Meta_should_create_json_with_meta_field(t *testing.T) {
	t.Parallel()
	// Given
	agg := es.FilterAgg(es.Term("status", "active")).Meta("color", "blue")

	// When Then
	assert.NotNil(t, agg)
	bodyJSON := assert.MarshalWithoutError(t, agg)
	// nolint:golint,lll
	assert.Equal(t, "{\"filter\":{\"term\":{\"status\":{\"value\":\"active\"}}},\"meta\":{\"color\":\"blue\"}}", bodyJSON)
}
