package es_test

import (
	"testing"

	"github.com/Trendyol/es-query-builder/es"
	"github.com/Trendyol/es-query-builder/test/assert"
)

////   RangeAgg   ////

func Test_RangeAgg_should_exist_on_es_package(t *testing.T) {
	t.Parallel()
	// Given When Then
	assert.NotNil(t, es.RangeAgg)
}

func Test_RangeAgg_method_should_create_rangeAggType(t *testing.T) {
	t.Parallel()
	// Given
	agg := es.RangeAgg("price")

	// Then
	assert.NotNil(t, agg)
	assert.IsTypeString(t, "es.rangeAggType", agg)
}

func Test_RangeAgg_should_create_json_with_field(t *testing.T) {
	t.Parallel()
	// Given
	agg := es.RangeAgg("price")

	// When Then
	assert.NotNil(t, agg)
	bodyJSON := assert.MarshalWithoutError(t, agg)
	assert.Equal(t, "{\"range\":{\"field\":\"price\"}}", bodyJSON)
}

////   RangeAgg Range   ////

func Test_RangeAgg_should_have_Range_method(t *testing.T) {
	t.Parallel()
	// Given
	agg := es.RangeAgg("price")

	// When Then
	assert.NotNil(t, agg.Range)
}

func Test_RangeAgg_Range_should_create_json_with_ranges(t *testing.T) {
	t.Parallel()
	// Given
	query := es.NewQuery(nil).
		Aggs(es.Agg("price_ranges",
			es.RangeAgg("price").
				Range(es.RangeEntry().To(50)).
				Range(es.RangeEntry().From(50).To(100)).
				Range(es.RangeEntry().From(100)),
		))

	// When Then
	assert.NotNil(t, query)
	bodyJSON := assert.MarshalWithoutError(t, query)
	// nolint:golint,lll
	assert.Equal(t, "{\"aggs\":{\"price_ranges\":{\"range\":{\"field\":\"price\",\"ranges\":[{\"to\":50},{\"from\":50,\"to\":100},{\"from\":100}]}}},\"query\":{}}", bodyJSON)
}

func Test_RangeAgg_Range_should_create_json_with_keyed_ranges(t *testing.T) {
	t.Parallel()
	// Given
	agg := es.RangeAgg("price").
		Range(es.RangeEntry().Key("cheap").To(50)).
		Range(es.RangeEntry().Key("average").From(50).To(100)).
		Range(es.RangeEntry().Key("expensive").From(100))

	// When Then
	assert.NotNil(t, agg)
	bodyJSON := assert.MarshalWithoutError(t, agg)
	// nolint:golint,lll
	assert.Equal(t, "{\"range\":{\"field\":\"price\",\"ranges\":[{\"key\":\"cheap\",\"to\":50},{\"from\":50,\"key\":\"average\",\"to\":100},{\"from\":100,\"key\":\"expensive\"}]}}", bodyJSON)
}

////   RangeAgg Keyed   ////

func Test_RangeAgg_should_have_Keyed_method(t *testing.T) {
	t.Parallel()
	// Given
	agg := es.RangeAgg("price")

	// When Then
	assert.NotNil(t, agg.Keyed)
}

func Test_RangeAgg_Keyed_should_create_json_with_keyed_field(t *testing.T) {
	t.Parallel()
	// Given
	agg := es.RangeAgg("price").Keyed(true)

	// When Then
	assert.NotNil(t, agg)
	bodyJSON := assert.MarshalWithoutError(t, agg)
	assert.Equal(t, "{\"range\":{\"field\":\"price\",\"keyed\":true}}", bodyJSON)
}

////   RangeAgg Missing   ////

func Test_RangeAgg_should_have_Missing_method(t *testing.T) {
	t.Parallel()
	// Given
	agg := es.RangeAgg("price")

	// When Then
	assert.NotNil(t, agg.Missing)
}

func Test_RangeAgg_Missing_should_create_json_with_missing_field(t *testing.T) {
	t.Parallel()
	// Given
	agg := es.RangeAgg("price").Missing(0)

	// When Then
	assert.NotNil(t, agg)
	bodyJSON := assert.MarshalWithoutError(t, agg)
	assert.Equal(t, "{\"range\":{\"field\":\"price\",\"missing\":0}}", bodyJSON)
}

////   RangeAgg Aggs   ////

func Test_RangeAgg_should_have_Aggs_method(t *testing.T) {
	t.Parallel()
	// Given
	agg := es.RangeAgg("price")

	// When Then
	assert.NotNil(t, agg.Aggs)
}

func Test_RangeAgg_Aggs_should_create_json_with_sub_aggregations(t *testing.T) {
	t.Parallel()
	// Given
	query := es.NewQuery(nil).
		Aggs(es.Agg("price_ranges",
			es.RangeAgg("price").
				Range(es.RangeEntry().To(50)).
				Range(es.RangeEntry().From(50)).
				Aggs(es.Agg("avg_score", es.AvgAgg("score"))),
		))

	// When Then
	assert.NotNil(t, query)
	bodyJSON := assert.MarshalWithoutError(t, query)
	// nolint:golint,lll
	assert.Equal(t, "{\"aggs\":{\"price_ranges\":{\"aggs\":{\"avg_score\":{\"avg\":{\"field\":\"score\"}}},\"range\":{\"field\":\"price\",\"ranges\":[{\"to\":50},{\"from\":50}]}}},\"query\":{}}", bodyJSON)
}

////   RangeAgg Meta   ////

func Test_RangeAgg_should_have_Meta_method(t *testing.T) {
	t.Parallel()
	// Given
	agg := es.RangeAgg("price")

	// When Then
	assert.NotNil(t, agg.Meta)
}

func Test_RangeAgg_Meta_should_create_json_with_meta_field(t *testing.T) {
	t.Parallel()
	// Given
	agg := es.RangeAgg("price").Meta("color", "blue")

	// When Then
	assert.NotNil(t, agg)
	bodyJSON := assert.MarshalWithoutError(t, agg)
	assert.Equal(t, "{\"meta\":{\"color\":\"blue\"},\"range\":{\"field\":\"price\"}}", bodyJSON)
}

////   RangeEntry   ////

func Test_RangeEntry_should_exist_on_es_package(t *testing.T) {
	t.Parallel()
	// Given When Then
	assert.NotNil(t, es.RangeEntry)
}

func Test_RangeEntry_method_should_create_rangeAggEntry(t *testing.T) {
	t.Parallel()
	// Given
	entry := es.RangeEntry()

	// Then
	assert.NotNil(t, entry)
	assert.IsTypeString(t, "es.rangeAggEntry", entry)
}

func Test_RangeEntry_From_should_create_json_with_from_field(t *testing.T) {
	t.Parallel()
	// Given
	entry := es.RangeEntry().From(50)

	// When Then
	assert.NotNil(t, entry)
	bodyJSON := assert.MarshalWithoutError(t, entry)
	assert.Equal(t, "{\"from\":50}", bodyJSON)
}

func Test_RangeEntry_To_should_create_json_with_to_field(t *testing.T) {
	t.Parallel()
	// Given
	entry := es.RangeEntry().To(100)

	// When Then
	assert.NotNil(t, entry)
	bodyJSON := assert.MarshalWithoutError(t, entry)
	assert.Equal(t, "{\"to\":100}", bodyJSON)
}

func Test_RangeEntry_Key_should_create_json_with_key_field(t *testing.T) {
	t.Parallel()
	// Given
	entry := es.RangeEntry().Key("cheap").To(50)

	// When Then
	assert.NotNil(t, entry)
	bodyJSON := assert.MarshalWithoutError(t, entry)
	assert.Equal(t, "{\"key\":\"cheap\",\"to\":50}", bodyJSON)
}
