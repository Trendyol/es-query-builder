package es_test

import (
	"testing"

	"github.com/Trendyol/es-query-builder/es"
	"github.com/Trendyol/es-query-builder/test/assert"
)

////   DateRangeAgg   ////

func Test_DateRangeAgg_should_exist_on_es_package(t *testing.T) {
	t.Parallel()
	// Given When Then
	assert.NotNil(t, es.DateRangeAgg)
}

func Test_DateRangeAgg_method_should_create_dateRangeAggType(t *testing.T) {
	t.Parallel()
	// Given
	agg := es.DateRangeAgg("date")

	// Then
	assert.NotNil(t, agg)
	assert.IsTypeString(t, "es.dateRangeAggType", agg)
}

func Test_DateRangeAgg_should_create_json_with_field(t *testing.T) {
	t.Parallel()
	// Given
	agg := es.DateRangeAgg("date")

	// When Then
	assert.NotNil(t, agg)
	bodyJSON := assert.MarshalWithoutError(t, agg)
	assert.Equal(t, "{\"date_range\":{\"field\":\"date\"}}", bodyJSON)
}

////   DateRangeAgg Range   ////

func Test_DateRangeAgg_should_have_Range_method(t *testing.T) {
	t.Parallel()
	// Given
	agg := es.DateRangeAgg("date")

	// When Then
	assert.NotNil(t, agg.Range)
}

func Test_DateRangeAgg_Range_should_create_json_with_ranges(t *testing.T) {
	t.Parallel()
	// Given
	query := es.NewQuery(nil).
		Aggs(es.Agg("date_ranges",
			es.DateRangeAgg("date").
				Range(es.DateRangeEntry().To("now-10M/M")).
				Range(es.DateRangeEntry().From("now-10M/M")),
		))

	// When Then
	assert.NotNil(t, query)
	bodyJSON := assert.MarshalWithoutError(t, query)
	// nolint:golint,lll
	assert.Equal(t, "{\"aggs\":{\"date_ranges\":{\"date_range\":{\"field\":\"date\",\"ranges\":[{\"to\":\"now-10M/M\"},{\"from\":\"now-10M/M\"}]}}},\"query\":{}}", bodyJSON)
}

func Test_DateRangeAgg_Range_should_create_json_with_keyed_ranges(t *testing.T) {
	t.Parallel()
	// Given
	agg := es.DateRangeAgg("date").
		Range(es.DateRangeEntry().Key("last_year").From("now-1y/y").To("now/y")).
		Range(es.DateRangeEntry().Key("this_year").From("now/y"))

	// When Then
	assert.NotNil(t, agg)
	bodyJSON := assert.MarshalWithoutError(t, agg)
	// nolint:golint,lll
	assert.Equal(t, "{\"date_range\":{\"field\":\"date\",\"ranges\":[{\"from\":\"now-1y/y\",\"key\":\"last_year\",\"to\":\"now/y\"},{\"from\":\"now/y\",\"key\":\"this_year\"}]}}", bodyJSON)
}

////   DateRangeAgg Format   ////

func Test_DateRangeAgg_should_have_Format_method(t *testing.T) {
	t.Parallel()
	// Given
	agg := es.DateRangeAgg("date")

	// When Then
	assert.NotNil(t, agg.Format)
}

func Test_DateRangeAgg_Format_should_create_json_with_format_field(t *testing.T) {
	t.Parallel()
	// Given
	agg := es.DateRangeAgg("date").Format("MM-yyyy")

	// When Then
	assert.NotNil(t, agg)
	bodyJSON := assert.MarshalWithoutError(t, agg)
	assert.Equal(t, "{\"date_range\":{\"field\":\"date\",\"format\":\"MM-yyyy\"}}", bodyJSON)
}

////   DateRangeAgg Keyed   ////

func Test_DateRangeAgg_should_have_Keyed_method(t *testing.T) {
	t.Parallel()
	// Given
	agg := es.DateRangeAgg("date")

	// When Then
	assert.NotNil(t, agg.Keyed)
}

func Test_DateRangeAgg_Keyed_should_create_json_with_keyed_field(t *testing.T) {
	t.Parallel()
	// Given
	agg := es.DateRangeAgg("date").Keyed(true)

	// When Then
	assert.NotNil(t, agg)
	bodyJSON := assert.MarshalWithoutError(t, agg)
	assert.Equal(t, "{\"date_range\":{\"field\":\"date\",\"keyed\":true}}", bodyJSON)
}

////   DateRangeAgg Missing   ////

func Test_DateRangeAgg_should_have_Missing_method(t *testing.T) {
	t.Parallel()
	// Given
	agg := es.DateRangeAgg("date")

	// When Then
	assert.NotNil(t, agg.Missing)
}

func Test_DateRangeAgg_Missing_should_create_json_with_missing_field(t *testing.T) {
	t.Parallel()
	// Given
	agg := es.DateRangeAgg("date").Missing("1970-01-01")

	// When Then
	assert.NotNil(t, agg)
	bodyJSON := assert.MarshalWithoutError(t, agg)
	assert.Equal(t, "{\"date_range\":{\"field\":\"date\",\"missing\":\"1970-01-01\"}}", bodyJSON)
}

////   DateRangeAgg TimeZone   ////

func Test_DateRangeAgg_should_have_TimeZone_method(t *testing.T) {
	t.Parallel()
	// Given
	agg := es.DateRangeAgg("date")

	// When Then
	assert.NotNil(t, agg.TimeZone)
}

func Test_DateRangeAgg_TimeZone_should_create_json_with_time_zone_field(t *testing.T) {
	t.Parallel()
	// Given
	agg := es.DateRangeAgg("date").TimeZone("CET")

	// When Then
	assert.NotNil(t, agg)
	bodyJSON := assert.MarshalWithoutError(t, agg)
	assert.Equal(t, "{\"date_range\":{\"field\":\"date\",\"time_zone\":\"CET\"}}", bodyJSON)
}

////   DateRangeAgg Aggs   ////

func Test_DateRangeAgg_should_have_Aggs_method(t *testing.T) {
	t.Parallel()
	// Given
	agg := es.DateRangeAgg("date")

	// When Then
	assert.NotNil(t, agg.Aggs)
}

func Test_DateRangeAgg_Aggs_should_create_json_with_sub_aggregations(t *testing.T) {
	t.Parallel()
	// Given
	query := es.NewQuery(nil).
		Aggs(es.Agg("date_ranges",
			es.DateRangeAgg("date").
				Range(es.DateRangeEntry().To("now-10M/M")).
				Aggs(es.Agg("avg_price", es.AvgAgg("price"))),
		))

	// When Then
	assert.NotNil(t, query)
	bodyJSON := assert.MarshalWithoutError(t, query)
	// nolint:golint,lll
	assert.Equal(t, "{\"aggs\":{\"date_ranges\":{\"aggs\":{\"avg_price\":{\"avg\":{\"field\":\"price\"}}},\"date_range\":{\"field\":\"date\",\"ranges\":[{\"to\":\"now-10M/M\"}]}}},\"query\":{}}", bodyJSON)
}

////   DateRangeAgg Meta   ////

func Test_DateRangeAgg_should_have_Meta_method(t *testing.T) {
	t.Parallel()
	// Given
	agg := es.DateRangeAgg("date")

	// When Then
	assert.NotNil(t, agg.Meta)
}

func Test_DateRangeAgg_Meta_should_create_json_with_meta_field(t *testing.T) {
	t.Parallel()
	// Given
	agg := es.DateRangeAgg("date").Meta("color", "blue")

	// When Then
	assert.NotNil(t, agg)
	bodyJSON := assert.MarshalWithoutError(t, agg)
	assert.Equal(t, "{\"date_range\":{\"field\":\"date\"},\"meta\":{\"color\":\"blue\"}}", bodyJSON)
}

////   DateRangeEntry   ////

func Test_DateRangeEntry_should_exist_on_es_package(t *testing.T) {
	t.Parallel()
	// Given When Then
	assert.NotNil(t, es.DateRangeEntry)
}

func Test_DateRangeEntry_method_should_create_dateRangeAggEntry(t *testing.T) {
	t.Parallel()
	// Given
	entry := es.DateRangeEntry()

	// Then
	assert.NotNil(t, entry)
	assert.IsTypeString(t, "es.dateRangeAggEntry", entry)
}

func Test_DateRangeEntry_From_should_create_json_with_from_field(t *testing.T) {
	t.Parallel()
	// Given
	entry := es.DateRangeEntry().From("now-10M/M")

	// When Then
	assert.NotNil(t, entry)
	bodyJSON := assert.MarshalWithoutError(t, entry)
	assert.Equal(t, "{\"from\":\"now-10M/M\"}", bodyJSON)
}

func Test_DateRangeEntry_To_should_create_json_with_to_field(t *testing.T) {
	t.Parallel()
	// Given
	entry := es.DateRangeEntry().To("now-10M/M")

	// When Then
	assert.NotNil(t, entry)
	bodyJSON := assert.MarshalWithoutError(t, entry)
	assert.Equal(t, "{\"to\":\"now-10M/M\"}", bodyJSON)
}

func Test_DateRangeEntry_Key_should_create_json_with_key_field(t *testing.T) {
	t.Parallel()
	// Given
	entry := es.DateRangeEntry().Key("last_year").From("now-1y/y").To("now/y")

	// When Then
	assert.NotNil(t, entry)
	bodyJSON := assert.MarshalWithoutError(t, entry)
	assert.Equal(t, "{\"from\":\"now-1y/y\",\"key\":\"last_year\",\"to\":\"now/y\"}", bodyJSON)
}

////   Complex DateRangeAgg   ////

func Test_DateRangeAgg_should_create_complex_json_with_all_parameters(t *testing.T) {
	t.Parallel()
	// Given
	agg := es.DateRangeAgg("date").
		Format("MM-yyyy").
		TimeZone("CET").
		Keyed(true).
		Range(es.DateRangeEntry().To("now-10M/M")).
		Range(es.DateRangeEntry().From("now-10M/M"))

	// When Then
	assert.NotNil(t, agg)
	bodyJSON := assert.MarshalWithoutError(t, agg)
	// nolint:golint,lll
	assert.Equal(t, "{\"date_range\":{\"field\":\"date\",\"format\":\"MM-yyyy\",\"keyed\":true,\"ranges\":[{\"to\":\"now-10M/M\"},{\"from\":\"now-10M/M\"}],\"time_zone\":\"CET\"}}", bodyJSON)
}
