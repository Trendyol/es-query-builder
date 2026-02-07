package es_test

import (
	"testing"

	"github.com/Trendyol/es-query-builder/es"
	"github.com/Trendyol/es-query-builder/test/assert"
)

////   DateHistogramAgg   ////

func Test_DateHistogramAgg_should_exist_on_es_package(t *testing.T) {
	t.Parallel()
	// Given When Then
	assert.NotNil(t, es.DateHistogramAgg)
}

func Test_DateHistogramAgg_method_should_create_dateHistogramAggType(t *testing.T) {
	t.Parallel()
	// Given
	agg := es.DateHistogramAgg("timestamp")

	// Then
	assert.NotNil(t, agg)
	assert.IsTypeString(t, "es.dateHistogramAggType", agg)
}

func Test_DateHistogramAgg_should_create_json_with_field(t *testing.T) {
	t.Parallel()
	// Given
	agg := es.DateHistogramAgg("timestamp")

	// When Then
	assert.NotNil(t, agg)
	bodyJSON := assert.MarshalWithoutError(t, agg)
	assert.Equal(t, "{\"date_histogram\":{\"field\":\"timestamp\"}}", bodyJSON)
}

////   DateHistogramAgg CalendarInterval   ////

func Test_DateHistogramAgg_should_have_CalendarInterval_method(t *testing.T) {
	t.Parallel()
	// Given
	agg := es.DateHistogramAgg("timestamp")

	// When Then
	assert.NotNil(t, agg.CalendarInterval)
}

func Test_DateHistogramAgg_CalendarInterval_should_create_json_with_calendar_interval_field(t *testing.T) {
	t.Parallel()
	// Given
	query := es.NewQuery(nil).
		Aggs(es.Agg("sales_over_time",
			es.DateHistogramAgg("date").CalendarInterval("month"),
		))

	// When Then
	assert.NotNil(t, query)
	bodyJSON := assert.MarshalWithoutError(t, query)
	// nolint:golint,lll
	assert.Equal(t, "{\"aggs\":{\"sales_over_time\":{\"date_histogram\":{\"calendar_interval\":\"month\",\"field\":\"date\"}}},\"query\":{}}", bodyJSON)
}

////   DateHistogramAgg FixedInterval   ////

func Test_DateHistogramAgg_should_have_FixedInterval_method(t *testing.T) {
	t.Parallel()
	// Given
	agg := es.DateHistogramAgg("timestamp")

	// When Then
	assert.NotNil(t, agg.FixedInterval)
}

func Test_DateHistogramAgg_FixedInterval_should_create_json_with_fixed_interval_field(t *testing.T) {
	t.Parallel()
	// Given
	agg := es.DateHistogramAgg("timestamp").FixedInterval("30d")

	// When Then
	assert.NotNil(t, agg)
	bodyJSON := assert.MarshalWithoutError(t, agg)
	assert.Equal(t, "{\"date_histogram\":{\"field\":\"timestamp\",\"fixed_interval\":\"30d\"}}", bodyJSON)
}

////   DateHistogramAgg Format   ////

func Test_DateHistogramAgg_should_have_Format_method(t *testing.T) {
	t.Parallel()
	// Given
	agg := es.DateHistogramAgg("timestamp")

	// When Then
	assert.NotNil(t, agg.Format)
}

func Test_DateHistogramAgg_Format_should_create_json_with_format_field(t *testing.T) {
	t.Parallel()
	// Given
	agg := es.DateHistogramAgg("timestamp").CalendarInterval("month").Format("yyyy-MM-dd")

	// When Then
	assert.NotNil(t, agg)
	bodyJSON := assert.MarshalWithoutError(t, agg)
	// nolint:golint,lll
	assert.Equal(t, "{\"date_histogram\":{\"calendar_interval\":\"month\",\"field\":\"timestamp\",\"format\":\"yyyy-MM-dd\"}}", bodyJSON)
}

////   DateHistogramAgg TimeZone   ////

func Test_DateHistogramAgg_should_have_TimeZone_method(t *testing.T) {
	t.Parallel()
	// Given
	agg := es.DateHistogramAgg("timestamp")

	// When Then
	assert.NotNil(t, agg.TimeZone)
}

func Test_DateHistogramAgg_TimeZone_should_create_json_with_time_zone_field(t *testing.T) {
	t.Parallel()
	// Given
	agg := es.DateHistogramAgg("timestamp").CalendarInterval("day").TimeZone("-01:00")

	// When Then
	assert.NotNil(t, agg)
	bodyJSON := assert.MarshalWithoutError(t, agg)
	// nolint:golint,lll
	assert.Equal(t, "{\"date_histogram\":{\"calendar_interval\":\"day\",\"field\":\"timestamp\",\"time_zone\":\"-01:00\"}}", bodyJSON)
}

////   DateHistogramAgg Offset   ////

func Test_DateHistogramAgg_should_have_Offset_method(t *testing.T) {
	t.Parallel()
	// Given
	agg := es.DateHistogramAgg("timestamp")

	// When Then
	assert.NotNil(t, agg.Offset)
}

func Test_DateHistogramAgg_Offset_should_create_json_with_offset_field(t *testing.T) {
	t.Parallel()
	// Given
	agg := es.DateHistogramAgg("timestamp").CalendarInterval("day").Offset("+6h")

	// When Then
	assert.NotNil(t, agg)
	bodyJSON := assert.MarshalWithoutError(t, agg)
	assert.Equal(t, "{\"date_histogram\":{\"calendar_interval\":\"day\",\"field\":\"timestamp\",\"offset\":\"+6h\"}}", bodyJSON)
}

////   DateHistogramAgg MinDocCount   ////

func Test_DateHistogramAgg_should_have_MinDocCount_method(t *testing.T) {
	t.Parallel()
	// Given
	agg := es.DateHistogramAgg("timestamp")

	// When Then
	assert.NotNil(t, agg.MinDocCount)
}

func Test_DateHistogramAgg_MinDocCount_should_create_json_with_min_doc_count_field(t *testing.T) {
	t.Parallel()
	// Given
	agg := es.DateHistogramAgg("timestamp").CalendarInterval("month").MinDocCount(1)

	// When Then
	assert.NotNil(t, agg)
	bodyJSON := assert.MarshalWithoutError(t, agg)
	// nolint:golint,lll
	assert.Equal(t, "{\"date_histogram\":{\"calendar_interval\":\"month\",\"field\":\"timestamp\",\"min_doc_count\":1}}", bodyJSON)
}

////   DateHistogramAgg ExtendedBounds   ////

func Test_DateHistogramAgg_should_have_ExtendedBounds_method(t *testing.T) {
	t.Parallel()
	// Given
	agg := es.DateHistogramAgg("timestamp")

	// When Then
	assert.NotNil(t, agg.ExtendedBounds)
}

func Test_DateHistogramAgg_ExtendedBounds_should_create_json_with_extended_bounds_field(t *testing.T) {
	t.Parallel()
	// Given
	agg := es.DateHistogramAgg("timestamp").CalendarInterval("month").
		ExtendedBounds("2020-01-01", "2020-12-31")

	// When Then
	assert.NotNil(t, agg)
	bodyJSON := assert.MarshalWithoutError(t, agg)
	// nolint:golint,lll
	assert.Equal(t, "{\"date_histogram\":{\"calendar_interval\":\"month\",\"extended_bounds\":{\"max\":\"2020-12-31\",\"min\":\"2020-01-01\"},\"field\":\"timestamp\"}}", bodyJSON)
}

////   DateHistogramAgg HardBounds   ////

func Test_DateHistogramAgg_should_have_HardBounds_method(t *testing.T) {
	t.Parallel()
	// Given
	agg := es.DateHistogramAgg("timestamp")

	// When Then
	assert.NotNil(t, agg.HardBounds)
}

func Test_DateHistogramAgg_HardBounds_should_create_json_with_hard_bounds_field(t *testing.T) {
	t.Parallel()
	// Given
	agg := es.DateHistogramAgg("timestamp").CalendarInterval("month").
		HardBounds("2020-01-01", "2020-12-31")

	// When Then
	assert.NotNil(t, agg)
	bodyJSON := assert.MarshalWithoutError(t, agg)
	// nolint:golint,lll
	assert.Equal(t, "{\"date_histogram\":{\"calendar_interval\":\"month\",\"field\":\"timestamp\",\"hard_bounds\":{\"max\":\"2020-12-31\",\"min\":\"2020-01-01\"}}}", bodyJSON)
}

////   DateHistogramAgg Keyed   ////

func Test_DateHistogramAgg_should_have_Keyed_method(t *testing.T) {
	t.Parallel()
	// Given
	agg := es.DateHistogramAgg("timestamp")

	// When Then
	assert.NotNil(t, agg.Keyed)
}

func Test_DateHistogramAgg_Keyed_should_create_json_with_keyed_field(t *testing.T) {
	t.Parallel()
	// Given
	agg := es.DateHistogramAgg("timestamp").CalendarInterval("month").Keyed(true)

	// When Then
	assert.NotNil(t, agg)
	bodyJSON := assert.MarshalWithoutError(t, agg)
	// nolint:golint,lll
	assert.Equal(t, "{\"date_histogram\":{\"calendar_interval\":\"month\",\"field\":\"timestamp\",\"keyed\":true}}", bodyJSON)
}

////   DateHistogramAgg Missing   ////

func Test_DateHistogramAgg_should_have_Missing_method(t *testing.T) {
	t.Parallel()
	// Given
	agg := es.DateHistogramAgg("timestamp")

	// When Then
	assert.NotNil(t, agg.Missing)
}

func Test_DateHistogramAgg_Missing_should_create_json_with_missing_field(t *testing.T) {
	t.Parallel()
	// Given
	agg := es.DateHistogramAgg("timestamp").CalendarInterval("month").Missing("2000-01-01")

	// When Then
	assert.NotNil(t, agg)
	bodyJSON := assert.MarshalWithoutError(t, agg)
	// nolint:golint,lll
	assert.Equal(t, "{\"date_histogram\":{\"calendar_interval\":\"month\",\"field\":\"timestamp\",\"missing\":\"2000-01-01\"}}", bodyJSON)
}

////   DateHistogramAgg Order   ////

func Test_DateHistogramAgg_should_have_Order_method(t *testing.T) {
	t.Parallel()
	// Given
	agg := es.DateHistogramAgg("timestamp")

	// When Then
	assert.NotNil(t, agg.Order)
}

func Test_DateHistogramAgg_Order_should_handle_nil_order(t *testing.T) {
	t.Parallel()
	// Given
	agg := es.DateHistogramAgg("timestamp").CalendarInterval("month").Order(nil)

	// When Then
	assert.NotNil(t, agg)
	bodyJSON := assert.MarshalWithoutError(t, agg)
	// nolint:golint,lll
	assert.Equal(t, "{\"date_histogram\":{\"calendar_interval\":\"month\",\"field\":\"timestamp\"}}", bodyJSON)
}

////   DateHistogramAgg Aggs   ////

func Test_DateHistogramAgg_should_have_Aggs_method(t *testing.T) {
	t.Parallel()
	// Given
	agg := es.DateHistogramAgg("timestamp")

	// When Then
	assert.NotNil(t, agg.Aggs)
}

func Test_DateHistogramAgg_Aggs_should_create_json_with_sub_aggregations(t *testing.T) {
	t.Parallel()
	// Given
	query := es.NewQuery(nil).
		Aggs(es.Agg("sales_over_time",
			es.DateHistogramAgg("date").CalendarInterval("month").
				Aggs(es.Agg("total_sales", es.SumAgg("price"))),
		))

	// When Then
	assert.NotNil(t, query)
	bodyJSON := assert.MarshalWithoutError(t, query)
	// nolint:golint,lll
	assert.Equal(t, "{\"aggs\":{\"sales_over_time\":{\"aggs\":{\"total_sales\":{\"sum\":{\"field\":\"price\"}}},\"date_histogram\":{\"calendar_interval\":\"month\",\"field\":\"date\"}}},\"query\":{}}", bodyJSON)
}

////   DateHistogramAgg Meta   ////

func Test_DateHistogramAgg_should_have_Meta_method(t *testing.T) {
	t.Parallel()
	// Given
	agg := es.DateHistogramAgg("timestamp")

	// When Then
	assert.NotNil(t, agg.Meta)
}

func Test_DateHistogramAgg_Meta_should_create_json_with_meta_field(t *testing.T) {
	t.Parallel()
	// Given
	agg := es.DateHistogramAgg("timestamp").Meta("color", "blue")

	// When Then
	assert.NotNil(t, agg)
	bodyJSON := assert.MarshalWithoutError(t, agg)
	assert.Equal(t, "{\"date_histogram\":{\"field\":\"timestamp\"},\"meta\":{\"color\":\"blue\"}}", bodyJSON)
}

////   Complex DateHistogramAgg   ////

func Test_DateHistogramAgg_should_create_complex_json_with_all_parameters(t *testing.T) {
	t.Parallel()
	// Given
	query := es.NewQuery(nil).
		Aggs(es.Agg("sales_over_time",
			es.DateHistogramAgg("date").
				CalendarInterval("month").
				Format("yyyy-MM-dd").
				TimeZone("CET").
				Offset("+6h").
				MinDocCount(1).
				Keyed(true).
				Aggs(es.Agg("total_sales", es.SumAgg("price"))),
		))

	// When Then
	assert.NotNil(t, query)
	bodyJSON := assert.MarshalWithoutError(t, query)
	// nolint:golint,lll
	assert.Equal(t, "{\"aggs\":{\"sales_over_time\":{\"aggs\":{\"total_sales\":{\"sum\":{\"field\":\"price\"}}},\"date_histogram\":{\"calendar_interval\":\"month\",\"field\":\"date\",\"format\":\"yyyy-MM-dd\",\"keyed\":true,\"min_doc_count\":1,\"offset\":\"+6h\",\"time_zone\":\"CET\"}}},\"query\":{}}", bodyJSON)
}
