package es_test

import (
	"testing"

	Order "github.com/Trendyol/es-query-builder/es/enums/sort/order"

	"github.com/Trendyol/es-query-builder/es"
	"github.com/Trendyol/es-query-builder/test/assert"
)

func Test_CompositeAgg_should_exist_on_es_package(t *testing.T) {
	t.Parallel()
	// Given When Then
	assert.NotNil(t, es.CompositeAgg)
}

func Test_CompositeAgg_should_return_type_of_compositeAggType(t *testing.T) {
	t.Parallel()
	// Given
	aggsQuery := es.CompositeAgg(es.CompositeTerms("brand", "brand.keyword"))

	// When Then
	assert.NotNil(t, aggsQuery)
	assert.IsTypeString(t, "es.compositeAggType", aggsQuery)
	assert.MarshalWithoutError(t, aggsQuery)
}

func Test_CompositeAgg_should_create_json_with_composite_field_inside(t *testing.T) {
	t.Parallel()
	// Given
	a := es.CompositeAgg(es.CompositeTerms("brand", "brand.keyword"))

	// When Then
	assert.NotNil(t, a)
	bodyJSON := assert.MarshalWithoutError(t, a)
	assert.Equal(t, "{\"composite\":{\"sources\":[{\"brand\":{\"terms\":{\"field\":\"brand.keyword\"}}}]}}", bodyJSON)
}

func Test_CompositeAgg_should_have_Size_method(t *testing.T) {
	t.Parallel()
	// Given
	a := es.CompositeAgg(es.CompositeTerms("brand", "brand.keyword"))

	// When Then
	assert.NotNil(t, a.Size)
}

func Test_Size_should_add_size_field_into_CompositeAgg(t *testing.T) {
	t.Parallel()
	// Given
	a := es.CompositeAgg(es.CompositeTerms("brand", "brand.keyword")).Size(100)

	// When Then
	assert.NotNil(t, a)
	bodyJSON := assert.MarshalWithoutError(t, a)
	assert.Equal(t, "{\"composite\":{\"size\":100,\"sources\":[{\"brand\":{\"terms\":{\"field\":\"brand.keyword\"}}}]}}", bodyJSON)
}

func Test_CompositeAgg_should_have_After_method(t *testing.T) {
	t.Parallel()
	// Given
	a := es.CompositeAgg(es.CompositeTerms("brand", "brand.keyword"))

	// When Then
	assert.NotNil(t, a.After)
}

func Test_After_should_add_after_field_into_CompositeAgg(t *testing.T) {
	t.Parallel()
	// Given
	a := es.CompositeAgg(es.CompositeTerms("brand", "brand.keyword")).
		After(es.Object{"brand": "nike"})

	// When Then
	assert.NotNil(t, a)
	bodyJSON := assert.MarshalWithoutError(t, a)
	// nolint:golint,lll
	assert.Equal(t, "{\"composite\":{\"after\":{\"brand\":\"nike\"},\"sources\":[{\"brand\":{\"terms\":{\"field\":\"brand.keyword\"}}}]}}", bodyJSON)
}

func Test_CompositeAgg_should_have_Aggs_method(t *testing.T) {
	t.Parallel()
	// Given
	a := es.CompositeAgg(es.CompositeTerms("brand", "brand.keyword"))

	// When Then
	assert.NotNil(t, a.Aggs)
}

func Test_Aggs_should_add_aggs_field_into_CompositeAgg(t *testing.T) {
	t.Parallel()
	// Given
	a := es.CompositeAgg(es.CompositeTerms("brand", "brand.keyword")).
		Aggs(es.Agg("avg_score", es.AvgAgg("score")))

	// When Then
	assert.NotNil(t, a)
	bodyJSON := assert.MarshalWithoutError(t, a)
	// nolint:golint,lll
	assert.Equal(t, "{\"aggs\":{\"avg_score\":{\"avg\":{\"field\":\"score\"}}},\"composite\":{\"sources\":[{\"brand\":{\"terms\":{\"field\":\"brand.keyword\"}}}]}}", bodyJSON)
}

func Test_Aggs_should_not_add_aggs_when_nil_into_CompositeAgg(t *testing.T) {
	t.Parallel()
	// Given
	a := es.CompositeAgg(es.CompositeTerms("brand", "brand.keyword")).Aggs(nil)

	// When Then
	assert.NotNil(t, a)
	bodyJSON := assert.MarshalWithoutError(t, a)
	assert.Equal(t, "{\"composite\":{\"sources\":[{\"brand\":{\"terms\":{\"field\":\"brand.keyword\"}}}]}}", bodyJSON)
}

func Test_CompositeAgg_should_have_Meta_method(t *testing.T) {
	t.Parallel()
	// Given
	a := es.CompositeAgg(es.CompositeTerms("brand", "brand.keyword"))

	// When Then
	assert.NotNil(t, a.Meta)
}

func Test_Meta_should_add_meta_field_into_CompositeAgg(t *testing.T) {
	t.Parallel()
	// Given
	a := es.CompositeAgg(es.CompositeTerms("brand", "brand.keyword")).
		Meta("description", "Brand pages")

	// When Then
	assert.NotNil(t, a)
	bodyJSON := assert.MarshalWithoutError(t, a)
	// nolint:golint,lll
	assert.Equal(t, "{\"composite\":{\"sources\":[{\"brand\":{\"terms\":{\"field\":\"brand.keyword\"}}}]},\"meta\":{\"description\":\"Brand pages\"}}", bodyJSON)
}

func Test_CompositeAgg_with_multiple_sources_should_create_correct_json(t *testing.T) {
	t.Parallel()
	// Given
	a := es.CompositeAgg(
		es.CompositeTerms("brand", "brand.keyword"),
		es.CompositeHistogram("price", "price", 50),
	).Size(100).After(es.Object{"brand": "nike", "price": 100}).
		Aggs(es.Agg("avg_score", es.AvgAgg("score")))

	// When Then
	assert.NotNil(t, a)
	bodyJSON := assert.MarshalWithoutError(t, a)
	// nolint:golint,lll
	assert.Equal(t, "{\"aggs\":{\"avg_score\":{\"avg\":{\"field\":\"score\"}}},\"composite\":{\"after\":{\"brand\":\"nike\",\"price\":100},\"size\":100,\"sources\":[{\"brand\":{\"terms\":{\"field\":\"brand.keyword\"}}},{\"price\":{\"histogram\":{\"field\":\"price\",\"interval\":50}}}]}}", bodyJSON)
}

func Test_CompositeTerms_should_exist_on_es_package(t *testing.T) {
	t.Parallel()
	// Given When Then
	assert.NotNil(t, es.CompositeTerms)
}

func Test_CompositeTerms_should_return_type_of_compositeSourceType(t *testing.T) {
	t.Parallel()
	// Given
	source := es.CompositeTerms("brand", "brand.keyword")

	// When Then
	assert.NotNil(t, source)
	assert.IsTypeString(t, "es.compositeSourceType", source)
	assert.MarshalWithoutError(t, source)
}

func Test_CompositeTerms_should_create_json_with_terms_source(t *testing.T) {
	t.Parallel()
	// Given
	source := es.CompositeTerms("brand", "brand.keyword")

	// When Then
	assert.NotNil(t, source)
	bodyJSON := assert.MarshalWithoutError(t, source)
	assert.Equal(t, "{\"brand\":{\"terms\":{\"field\":\"brand.keyword\"}}}", bodyJSON)
}

func Test_CompositeTerms_MissingBucket_should_add_missing_bucket_field(t *testing.T) {
	t.Parallel()
	// Given
	source := es.CompositeTerms("brand", "brand.keyword").MissingBucket(true)

	// When Then
	assert.NotNil(t, source)
	bodyJSON := assert.MarshalWithoutError(t, source)
	assert.Equal(t, "{\"brand\":{\"terms\":{\"field\":\"brand.keyword\",\"missing_bucket\":true}}}", bodyJSON)
}

func Test_CompositeTerms_Order_should_add_order_field(t *testing.T) {
	t.Parallel()
	// Given
	source := es.CompositeTerms("brand", "brand.keyword").Order(Order.Desc)

	// When Then
	assert.NotNil(t, source)
	bodyJSON := assert.MarshalWithoutError(t, source)
	assert.Equal(t, "{\"brand\":{\"terms\":{\"field\":\"brand.keyword\",\"order\":\"desc\"}}}", bodyJSON)
}

func Test_CompositeHistogram_should_exist_on_es_package(t *testing.T) {
	t.Parallel()
	// Given When Then
	assert.NotNil(t, es.CompositeHistogram)
}

func Test_CompositeHistogram_should_create_json_with_histogram_source(t *testing.T) {
	t.Parallel()
	// Given
	source := es.CompositeHistogram("price", "price", 50)

	// When Then
	assert.NotNil(t, source)
	assert.IsTypeString(t, "es.compositeSourceType", source)
	bodyJSON := assert.MarshalWithoutError(t, source)
	assert.Equal(t, "{\"price\":{\"histogram\":{\"field\":\"price\",\"interval\":50}}}", bodyJSON)
}

func Test_CompositeHistogram_Offset_should_add_offset_field(t *testing.T) {
	t.Parallel()
	// Given
	source := es.CompositeHistogram("price", "price", 10).Offset(5)

	// When Then
	assert.NotNil(t, source)
	bodyJSON := assert.MarshalWithoutError(t, source)
	assert.Equal(t, "{\"price\":{\"histogram\":{\"field\":\"price\",\"interval\":10,\"offset\":5}}}", bodyJSON)
}

func Test_CompositeHistogram_MissingBucket_and_Order_should_create_correct_json(t *testing.T) {
	t.Parallel()
	// Given
	source := es.CompositeHistogram("price", "price", 50).
		MissingBucket(true).
		Order(Order.Asc)

	// When Then
	assert.NotNil(t, source)
	bodyJSON := assert.MarshalWithoutError(t, source)
	// nolint:golint,lll
	assert.Equal(t, "{\"price\":{\"histogram\":{\"field\":\"price\",\"interval\":50,\"missing_bucket\":true,\"order\":\"asc\"}}}", bodyJSON)
}

func Test_CompositeDateHistogram_should_exist_on_es_package(t *testing.T) {
	t.Parallel()
	// Given When Then
	assert.NotNil(t, es.CompositeDateHistogram)
}

func Test_CompositeDateHistogram_should_create_json_with_date_histogram_source(t *testing.T) {
	t.Parallel()
	// Given
	source := es.CompositeDateHistogram("by_day", "timestamp").CalendarInterval("day")

	// When Then
	assert.NotNil(t, source)
	assert.IsTypeString(t, "es.compositeSourceType", source)
	bodyJSON := assert.MarshalWithoutError(t, source)
	assert.Equal(t, "{\"by_day\":{\"date_histogram\":{\"calendar_interval\":\"day\",\"field\":\"timestamp\"}}}", bodyJSON)
}

func Test_CompositeDateHistogram_FixedInterval_should_add_fixed_interval_field(t *testing.T) {
	t.Parallel()
	// Given
	source := es.CompositeDateHistogram("by_day", "timestamp").FixedInterval("1d")

	// When Then
	assert.NotNil(t, source)
	bodyJSON := assert.MarshalWithoutError(t, source)
	assert.Equal(t, "{\"by_day\":{\"date_histogram\":{\"field\":\"timestamp\",\"fixed_interval\":\"1d\"}}}", bodyJSON)
}

func Test_CompositeDateHistogram_Format_and_TimeZone_should_create_correct_json(t *testing.T) {
	t.Parallel()
	// Given
	source := es.CompositeDateHistogram("by_day", "timestamp").
		CalendarInterval("day").
		Format("yyyy-MM-dd").
		TimeZone("+03:00").
		MissingBucket(true).
		Order(Order.Asc)

	// When Then
	assert.NotNil(t, source)
	bodyJSON := assert.MarshalWithoutError(t, source)
	// nolint:golint,lll
	assert.Equal(t, "{\"by_day\":{\"date_histogram\":{\"calendar_interval\":\"day\",\"field\":\"timestamp\",\"format\":\"yyyy-MM-dd\",\"missing_bucket\":true,\"order\":\"asc\",\"time_zone\":\"+03:00\"}}}", bodyJSON)
}

func Test_CompositeAgg_inside_NewQuery_should_create_correct_json(t *testing.T) {
	t.Parallel()
	// Given
	query := es.NewQuery(es.MatchAll()).Aggs(
		es.Agg("pages",
			es.CompositeAgg(
				es.CompositeTerms("brand", "brand.keyword"),
				es.CompositeHistogram("price", "price", 50),
			).Size(100).After(es.Object{"brand": "nike", "price": 100}).
				Aggs(es.Agg("avg_score", es.AvgAgg("score"))),
		),
	)

	// When Then
	assert.NotNil(t, query)
	bodyJSON := assert.MarshalWithoutError(t, query)
	// nolint:golint,lll
	assert.Equal(t, "{\"aggs\":{\"pages\":{\"aggs\":{\"avg_score\":{\"avg\":{\"field\":\"score\"}}},\"composite\":{\"after\":{\"brand\":\"nike\",\"price\":100},\"size\":100,\"sources\":[{\"brand\":{\"terms\":{\"field\":\"brand.keyword\"}}},{\"price\":{\"histogram\":{\"field\":\"price\",\"interval\":50}}}]}}},\"query\":{\"match_all\":{}}}", bodyJSON)
}
