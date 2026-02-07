package es_test

import (
	"testing"

	"github.com/Trendyol/es-query-builder/es"
	"github.com/Trendyol/es-query-builder/test/assert"
)

////   FiltersAgg   ////

func Test_FiltersAgg_should_exist_on_es_package(t *testing.T) {
	t.Parallel()
	// Given When Then
	assert.NotNil(t, es.FiltersAgg)
}

func Test_FiltersAgg_method_should_create_filtersAggType(t *testing.T) {
	t.Parallel()
	// Given
	agg := es.FiltersAgg()

	// Then
	assert.NotNil(t, agg)
	assert.IsTypeString(t, "es.filtersAggType", agg)
}

func Test_FiltersAgg_should_create_json_with_empty_filters(t *testing.T) {
	t.Parallel()
	// Given
	agg := es.FiltersAgg()

	// When Then
	assert.NotNil(t, agg)
	bodyJSON := assert.MarshalWithoutError(t, agg)
	assert.Equal(t, "{\"filters\":{\"filters\":{}}}", bodyJSON)
}

////   FiltersAgg Filter   ////

func Test_FiltersAgg_should_have_Filter_method(t *testing.T) {
	t.Parallel()
	// Given
	agg := es.FiltersAgg()

	// When Then
	assert.NotNil(t, agg.Filter)
}

func Test_FiltersAgg_Filter_should_create_json_with_named_filter(t *testing.T) {
	t.Parallel()
	// Given
	agg := es.FiltersAgg().
		Filter("errors", es.Term("status", "error"))

	// When Then
	assert.NotNil(t, agg)
	bodyJSON := assert.MarshalWithoutError(t, agg)
	assert.Equal(t, "{\"filters\":{\"filters\":{\"errors\":{\"term\":{\"status\":{\"value\":\"error\"}}}}}}", bodyJSON)
}

func Test_FiltersAgg_Filter_should_create_json_with_multiple_named_filters(t *testing.T) {
	t.Parallel()
	// Given
	query := es.NewQuery(nil).
		Aggs(es.Agg("messages",
			es.FiltersAgg().
				Filter("errors", es.Match("body", "error")).
				Filter("warnings", es.Match("body", "warning")),
		))

	// When Then
	assert.NotNil(t, query)
	bodyJSON := assert.MarshalWithoutError(t, query)
	// nolint:golint,lll
	assert.Equal(t, "{\"aggs\":{\"messages\":{\"filters\":{\"filters\":{\"errors\":{\"match\":{\"body\":{\"query\":\"error\"}}},\"warnings\":{\"match\":{\"body\":{\"query\":\"warning\"}}}}}}},\"query\":{}}", bodyJSON)
}

func Test_FiltersAgg_Filter_should_handle_bool_filter(t *testing.T) {
	t.Parallel()
	// Given
	agg := es.FiltersAgg().
		Filter("active", es.Bool().Must(es.Term("status", "active")))

	// When Then
	assert.NotNil(t, agg)
	bodyJSON := assert.MarshalWithoutError(t, agg)
	// nolint:golint,lll
	assert.Equal(t, "{\"filters\":{\"filters\":{\"active\":{\"bool\":{\"must\":[{\"term\":{\"status\":{\"value\":\"active\"}}}]}}}}}", bodyJSON)
}

////   FiltersAgg OtherBucket   ////

func Test_FiltersAgg_should_have_OtherBucket_method(t *testing.T) {
	t.Parallel()
	// Given
	agg := es.FiltersAgg()

	// When Then
	assert.NotNil(t, agg.OtherBucket)
}

func Test_FiltersAgg_OtherBucket_should_create_json_with_other_bucket_field(t *testing.T) {
	t.Parallel()
	// Given
	agg := es.FiltersAgg().
		Filter("errors", es.Term("status", "error")).
		OtherBucket(true)

	// When Then
	assert.NotNil(t, agg)
	bodyJSON := assert.MarshalWithoutError(t, agg)
	// nolint:golint,lll
	assert.Equal(t, "{\"filters\":{\"filters\":{\"errors\":{\"term\":{\"status\":{\"value\":\"error\"}}}},\"other_bucket\":true}}", bodyJSON)
}

////   FiltersAgg OtherBucketKey   ////

func Test_FiltersAgg_should_have_OtherBucketKey_method(t *testing.T) {
	t.Parallel()
	// Given
	agg := es.FiltersAgg()

	// When Then
	assert.NotNil(t, agg.OtherBucketKey)
}

func Test_FiltersAgg_OtherBucketKey_should_create_json_with_other_bucket_key_field(t *testing.T) {
	t.Parallel()
	// Given
	agg := es.FiltersAgg().
		Filter("errors", es.Term("status", "error")).
		OtherBucket(true).
		OtherBucketKey("remaining")

	// When Then
	assert.NotNil(t, agg)
	bodyJSON := assert.MarshalWithoutError(t, agg)
	// nolint:golint,lll
	assert.Equal(t, "{\"filters\":{\"filters\":{\"errors\":{\"term\":{\"status\":{\"value\":\"error\"}}}},\"other_bucket\":true,\"other_bucket_key\":\"remaining\"}}", bodyJSON)
}

////   FiltersAgg Aggs   ////

func Test_FiltersAgg_should_have_Aggs_method(t *testing.T) {
	t.Parallel()
	// Given
	agg := es.FiltersAgg()

	// When Then
	assert.NotNil(t, agg.Aggs)
}

func Test_FiltersAgg_Aggs_should_create_json_with_sub_aggregations(t *testing.T) {
	t.Parallel()
	// Given
	query := es.NewQuery(nil).
		Aggs(es.Agg("messages",
			es.FiltersAgg().
				Filter("errors", es.Match("body", "error")).
				Aggs(es.Agg("avg_price", es.AvgAgg("price"))),
		))

	// When Then
	assert.NotNil(t, query)
	bodyJSON := assert.MarshalWithoutError(t, query)
	// nolint:golint,lll
	assert.Equal(t, "{\"aggs\":{\"messages\":{\"aggs\":{\"avg_price\":{\"avg\":{\"field\":\"price\"}}},\"filters\":{\"filters\":{\"errors\":{\"match\":{\"body\":{\"query\":\"error\"}}}}}}},\"query\":{}}", bodyJSON)
}

////   FiltersAgg Meta   ////

func Test_FiltersAgg_should_have_Meta_method(t *testing.T) {
	t.Parallel()
	// Given
	agg := es.FiltersAgg()

	// When Then
	assert.NotNil(t, agg.Meta)
}

func Test_FiltersAgg_Meta_should_create_json_with_meta_field(t *testing.T) {
	t.Parallel()
	// Given
	agg := es.FiltersAgg().Meta("color", "blue")

	// When Then
	assert.NotNil(t, agg)
	bodyJSON := assert.MarshalWithoutError(t, agg)
	assert.Equal(t, "{\"filters\":{\"filters\":{}},\"meta\":{\"color\":\"blue\"}}", bodyJSON)
}
