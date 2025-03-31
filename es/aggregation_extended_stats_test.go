package es_test

import (
	"testing"

	ScriptLanguage "github.com/Trendyol/es-query-builder/es/enums/script-language"

	"github.com/Trendyol/es-query-builder/es"
	"github.com/Trendyol/es-query-builder/test/assert"
)

func Test_ExtendedStatsAgg_should_exist_on_es_package(t *testing.T) {
	t.Parallel()
	// Given When Then
	assert.NotNil(t, es.ExtendedStatsAgg)
}

func Test_ExtendedStatsAgg_should_return_type_of_extended_statsAggType(t *testing.T) {
	t.Parallel()
	// Given
	aggsQuery := es.ExtendedStatsAgg("price")

	// When Then
	assert.NotNil(t, aggsQuery)
	assert.IsTypeString(t, "es.extendedStatsAggType", aggsQuery)
	assert.MarshalWithoutError(t, aggsQuery)
}

func Test_ExtendedStatsAgg_should_create_json_with_terms_field_inside(t *testing.T) {
	t.Parallel()
	// Given
	a := es.ExtendedStatsAgg("price")

	// When Then
	assert.NotNil(t, a)
	bodyJSON := assert.MarshalWithoutError(t, a)
	assert.Equal(t, "{\"extended_stats\":{\"field\":\"price\"}}", bodyJSON)
}

func Test_ExtendedStatsAgg_should_have_Missing_method(t *testing.T) {
	t.Parallel()
	// Given
	a := es.ExtendedStatsAgg("price")

	// When Then
	assert.NotNil(t, a.Missing)
}

func Test_Missing_should_add_missing_field_into_ExtendedStatsAgg(t *testing.T) {
	t.Parallel()
	// Given
	a := es.ExtendedStatsAgg("price").Missing("missing_name")

	// When Then
	assert.NotNil(t, a)
	bodyJSON := assert.MarshalWithoutError(t, a)
	assert.Equal(t, "{\"extended_stats\":{\"field\":\"price\",\"missing\":\"missing_name\"}}", bodyJSON)
}

func Test_ExtendedStatsAgg_should_have_Script_method(t *testing.T) {
	t.Parallel()
	// Given
	a := es.ExtendedStatsAgg("price")

	// When Then
	assert.NotNil(t, a.Script)
}

func Test_Script_should_add_script_field_into_ExtendedStatsAgg(t *testing.T) {
	t.Parallel()
	// Given
	a := es.ExtendedStatsAgg("price").Script(es.ScriptID("id_12345", ScriptLanguage.Painless))

	// When Then
	assert.NotNil(t, a)
	bodyJSON := assert.MarshalWithoutError(t, a)
	assert.Equal(t, "{\"extended_stats\":{\"field\":\"price\",\"script\":{\"id\":\"id_12345\",\"lang\":\"painless\"}}}", bodyJSON)
}

func Test_ExtendedStatsAgg_should_have_Format_method(t *testing.T) {
	t.Parallel()
	// Given
	a := es.ExtendedStatsAgg("price")

	// When Then
	assert.NotNil(t, a.Format)
}

func Test_Format_should_add_format_field_into_ExtendedStatsAgg(t *testing.T) {
	t.Parallel()
	// Given
	a := es.ExtendedStatsAgg("price").Format("#.00")

	// When Then
	assert.NotNil(t, a)
	bodyJSON := assert.MarshalWithoutError(t, a)
	assert.Equal(t, "{\"extended_stats\":{\"field\":\"price\",\"format\":\"#.00\"}}", bodyJSON)
}

func Test_ExtendedStatsAgg_should_have_Meta_method(t *testing.T) {
	t.Parallel()
	// Given
	a := es.ExtendedStatsAgg("price")

	// When Then
	assert.NotNil(t, a.Meta)
}

func Test_Meta_should_add_meta_field_into_ExtendedStatsAgg(t *testing.T) {
	t.Parallel()
	// Given
	a := es.ExtendedStatsAgg("price").
		Meta("k1", "v1").
		Meta("k2", "v2")

	// When Then
	assert.NotNil(t, a)
	bodyJSON := assert.MarshalWithoutError(t, a)
	assert.Equal(t, "{\"extended_stats\":{\"field\":\"price\",\"meta\":{\"k1\":\"v1\",\"k2\":\"v2\"}}}", bodyJSON)
}

func Test_ExtendedStatsAgg_should_have_Aggs_method(t *testing.T) {
	t.Parallel()
	// Given
	a := es.ExtendedStatsAgg("price")

	// When Then
	assert.NotNil(t, a.Aggs)
}

func Test_Aggs_should_add_aggs_field_into_ExtendedStatsAgg_when_not_empty(t *testing.T) {
	t.Parallel()
	// Given
	a := es.ExtendedStatsAgg("price").Aggs(es.Agg("extended_stats_stock", es.ExtendedStatsAgg("stock")))

	// When Then
	assert.NotNil(t, a)
	bodyJSON := assert.MarshalWithoutError(t, a)
	assert.Equal(t, "{\"aggs\":{\"extended_stats_stock\":{\"extended_stats\":{\"field\":\"stock\"}}},\"extended_stats\":{\"field\":\"price\"}}", bodyJSON)
}

func Test_Aggs_should_not_add_aggs_field_into_ExtendedStatsAgg_when_it_is_empty(t *testing.T) {
	t.Parallel()
	// Given
	a := es.ExtendedStatsAgg("price").Aggs(nil)

	// When Then
	assert.NotNil(t, a)
	bodyJSON := assert.MarshalWithoutError(t, a)
	assert.Equal(t, "{\"extended_stats\":{\"field\":\"price\"}}", bodyJSON)
}
