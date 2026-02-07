package es_test

import (
	"testing"

	"github.com/Trendyol/es-query-builder/es"
	"github.com/Trendyol/es-query-builder/test/assert"
)

////   HistogramAgg   ////

func Test_HistogramAgg_should_exist_on_es_package(t *testing.T) {
	t.Parallel()
	// Given When Then
	assert.NotNil(t, es.HistogramAgg)
}

func Test_HistogramAgg_method_should_create_histogramAggType(t *testing.T) {
	t.Parallel()
	// Given
	agg := es.HistogramAgg("price", 50)

	// Then
	assert.NotNil(t, agg)
	assert.IsTypeString(t, "es.histogramAggType", agg)
}

func Test_HistogramAgg_should_create_json_with_field_and_interval(t *testing.T) {
	t.Parallel()
	// Given
	query := es.NewQuery(nil).
		Aggs(es.Agg("prices", es.HistogramAgg("price", 50)))

	// When Then
	assert.NotNil(t, query)
	bodyJSON := assert.MarshalWithoutError(t, query)
	assert.Equal(t, "{\"aggs\":{\"prices\":{\"histogram\":{\"field\":\"price\",\"interval\":50}}},\"query\":{}}", bodyJSON)
}

////   HistogramAgg MinDocCount   ////

func Test_HistogramAgg_should_have_MinDocCount_method(t *testing.T) {
	t.Parallel()
	// Given
	agg := es.HistogramAgg("price", 50)

	// When Then
	assert.NotNil(t, agg.MinDocCount)
}

func Test_HistogramAgg_MinDocCount_should_create_json_with_min_doc_count_field(t *testing.T) {
	t.Parallel()
	// Given
	agg := es.HistogramAgg("price", 50).MinDocCount(1)

	// When Then
	assert.NotNil(t, agg)
	bodyJSON := assert.MarshalWithoutError(t, agg)
	assert.Equal(t, "{\"histogram\":{\"field\":\"price\",\"interval\":50,\"min_doc_count\":1}}", bodyJSON)
}

////   HistogramAgg ExtendedBounds   ////

func Test_HistogramAgg_should_have_ExtendedBounds_method(t *testing.T) {
	t.Parallel()
	// Given
	agg := es.HistogramAgg("price", 50)

	// When Then
	assert.NotNil(t, agg.ExtendedBounds)
}

func Test_HistogramAgg_ExtendedBounds_should_create_json_with_extended_bounds_field(t *testing.T) {
	t.Parallel()
	// Given
	agg := es.HistogramAgg("price", 50).ExtendedBounds(0, 500)

	// When Then
	assert.NotNil(t, agg)
	bodyJSON := assert.MarshalWithoutError(t, agg)
	assert.Equal(t, "{\"histogram\":{\"extended_bounds\":{\"max\":500,\"min\":0},\"field\":\"price\",\"interval\":50}}", bodyJSON)
}

////   HistogramAgg HardBounds   ////

func Test_HistogramAgg_should_have_HardBounds_method(t *testing.T) {
	t.Parallel()
	// Given
	agg := es.HistogramAgg("price", 50)

	// When Then
	assert.NotNil(t, agg.HardBounds)
}

func Test_HistogramAgg_HardBounds_should_create_json_with_hard_bounds_field(t *testing.T) {
	t.Parallel()
	// Given
	agg := es.HistogramAgg("price", 50).HardBounds(0, 1000)

	// When Then
	assert.NotNil(t, agg)
	bodyJSON := assert.MarshalWithoutError(t, agg)
	assert.Equal(t, "{\"histogram\":{\"field\":\"price\",\"hard_bounds\":{\"max\":1000,\"min\":0},\"interval\":50}}", bodyJSON)
}

////   HistogramAgg Offset   ////

func Test_HistogramAgg_should_have_Offset_method(t *testing.T) {
	t.Parallel()
	// Given
	agg := es.HistogramAgg("price", 10)

	// When Then
	assert.NotNil(t, agg.Offset)
}

func Test_HistogramAgg_Offset_should_create_json_with_offset_field(t *testing.T) {
	t.Parallel()
	// Given
	agg := es.HistogramAgg("price", 10).Offset(5)

	// When Then
	assert.NotNil(t, agg)
	bodyJSON := assert.MarshalWithoutError(t, agg)
	assert.Equal(t, "{\"histogram\":{\"field\":\"price\",\"interval\":10,\"offset\":5}}", bodyJSON)
}

////   HistogramAgg Keyed   ////

func Test_HistogramAgg_should_have_Keyed_method(t *testing.T) {
	t.Parallel()
	// Given
	agg := es.HistogramAgg("price", 50)

	// When Then
	assert.NotNil(t, agg.Keyed)
}

func Test_HistogramAgg_Keyed_should_create_json_with_keyed_field(t *testing.T) {
	t.Parallel()
	// Given
	agg := es.HistogramAgg("price", 50).Keyed(true)

	// When Then
	assert.NotNil(t, agg)
	bodyJSON := assert.MarshalWithoutError(t, agg)
	assert.Equal(t, "{\"histogram\":{\"field\":\"price\",\"interval\":50,\"keyed\":true}}", bodyJSON)
}

////   HistogramAgg Missing   ////

func Test_HistogramAgg_should_have_Missing_method(t *testing.T) {
	t.Parallel()
	// Given
	agg := es.HistogramAgg("price", 50)

	// When Then
	assert.NotNil(t, agg.Missing)
}

func Test_HistogramAgg_Missing_should_create_json_with_missing_field(t *testing.T) {
	t.Parallel()
	// Given
	agg := es.HistogramAgg("price", 50).Missing(0)

	// When Then
	assert.NotNil(t, agg)
	bodyJSON := assert.MarshalWithoutError(t, agg)
	assert.Equal(t, "{\"histogram\":{\"field\":\"price\",\"interval\":50,\"missing\":0}}", bodyJSON)
}

////   HistogramAgg Order   ////

func Test_HistogramAgg_should_have_Order_method(t *testing.T) {
	t.Parallel()
	// Given
	agg := es.HistogramAgg("price", 50)

	// When Then
	assert.NotNil(t, agg.Order)
}

func Test_HistogramAgg_Order_should_handle_nil_order(t *testing.T) {
	t.Parallel()
	// Given
	agg := es.HistogramAgg("price", 50).Order(nil)

	// When Then
	assert.NotNil(t, agg)
	bodyJSON := assert.MarshalWithoutError(t, agg)
	assert.Equal(t, "{\"histogram\":{\"field\":\"price\",\"interval\":50}}", bodyJSON)
}

////   HistogramAgg Aggs   ////

func Test_HistogramAgg_should_have_Aggs_method(t *testing.T) {
	t.Parallel()
	// Given
	agg := es.HistogramAgg("price", 50)

	// When Then
	assert.NotNil(t, agg.Aggs)
}

func Test_HistogramAgg_Aggs_should_create_json_with_sub_aggregations(t *testing.T) {
	t.Parallel()
	// Given
	query := es.NewQuery(nil).
		Aggs(es.Agg("prices",
			es.HistogramAgg("price", 50).
				Aggs(es.Agg("avg_score", es.AvgAgg("score"))),
		))

	// When Then
	assert.NotNil(t, query)
	bodyJSON := assert.MarshalWithoutError(t, query)
	// nolint:golint,lll
	assert.Equal(t, "{\"aggs\":{\"prices\":{\"aggs\":{\"avg_score\":{\"avg\":{\"field\":\"score\"}}},\"histogram\":{\"field\":\"price\",\"interval\":50}}},\"query\":{}}", bodyJSON)
}

////   HistogramAgg Meta   ////

func Test_HistogramAgg_should_have_Meta_method(t *testing.T) {
	t.Parallel()
	// Given
	agg := es.HistogramAgg("price", 50)

	// When Then
	assert.NotNil(t, agg.Meta)
}

func Test_HistogramAgg_Meta_should_create_json_with_meta_field(t *testing.T) {
	t.Parallel()
	// Given
	agg := es.HistogramAgg("price", 50).Meta("color", "blue")

	// When Then
	assert.NotNil(t, agg)
	bodyJSON := assert.MarshalWithoutError(t, agg)
	assert.Equal(t, "{\"histogram\":{\"field\":\"price\",\"interval\":50},\"meta\":{\"color\":\"blue\"}}", bodyJSON)
}
