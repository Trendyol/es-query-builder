package es_test

import (
	ScriptLanguage "github.com/Trendyol/es-query-builder/es/enums/script-language"
	"testing"

	"github.com/Trendyol/es-query-builder/es"
	"github.com/Trendyol/es-query-builder/test/assert"
)

func Test_StatsAgg_should_exist_on_es_package(t *testing.T) {
	t.Parallel()
	// Given When Then
	assert.NotNil(t, es.StatsAgg)
}

func Test_StatsAgg_should_return_type_of_statsAggType(t *testing.T) {
	t.Parallel()
	// Given
	aggsQuery := es.StatsAgg("price")

	// When Then
	assert.NotNil(t, aggsQuery)
	assert.IsTypeString(t, "es.statsAggType", aggsQuery)
	assert.MarshalWithoutError(t, aggsQuery)
}

func Test_StatsAgg_should_create_json_with_terms_field_inside(t *testing.T) {
	t.Parallel()
	// Given
	a := es.StatsAgg("price")

	// When Then
	assert.NotNil(t, a)
	bodyJSON := assert.MarshalWithoutError(t, a)
	assert.Equal(t, "{\"stats\":{\"field\":\"price\"}}", bodyJSON)
}

func Test_StatsAgg_should_have_Missing_method(t *testing.T) {
	t.Parallel()
	// Given
	a := es.StatsAgg("price")

	// When Then
	assert.NotNil(t, a.Missing)
}

func Test_Missing_should_add_missing_field_into_StatsAgg(t *testing.T) {
	t.Parallel()
	// Given
	a := es.StatsAgg("price").Missing("missing_name")

	// When Then
	assert.NotNil(t, a)
	bodyJSON := assert.MarshalWithoutError(t, a)
	assert.Equal(t, "{\"stats\":{\"field\":\"price\",\"missing\":\"missing_name\"}}", bodyJSON)
}

func Test_StatsAgg_should_have_Script_method(t *testing.T) {
	t.Parallel()
	// Given
	a := es.StatsAgg("price")

	// When Then
	assert.NotNil(t, a.Script)
}

func Test_Script_should_add_script_field_into_StatsAgg(t *testing.T) {
	t.Parallel()
	// Given
	a := es.StatsAgg("price").Script(es.ScriptID("id_12345", ScriptLanguage.Painless))

	// When Then
	assert.NotNil(t, a)
	bodyJSON := assert.MarshalWithoutError(t, a)
	assert.Equal(t, "{\"stats\":{\"field\":\"price\",\"script\":{\"id\":\"id_12345\",\"lang\":\"painless\"}}}", bodyJSON)
}

func Test_StatsAgg_should_have_Format_method(t *testing.T) {
	t.Parallel()
	// Given
	a := es.StatsAgg("price")

	// When Then
	assert.NotNil(t, a.Format)
}

func Test_Format_should_add_format_field_into_StatsAgg(t *testing.T) {
	t.Parallel()
	// Given
	a := es.StatsAgg("price").Format("#.00")

	// When Then
	assert.NotNil(t, a)
	bodyJSON := assert.MarshalWithoutError(t, a)
	assert.Equal(t, "{\"stats\":{\"field\":\"price\",\"format\":\"#.00\"}}", bodyJSON)
}

func Test_StatsAgg_should_have_Meta_method(t *testing.T) {
	t.Parallel()
	// Given
	a := es.StatsAgg("price")

	// When Then
	assert.NotNil(t, a.Meta)
}

func Test_Meta_should_add_meta_field_into_StatsAgg(t *testing.T) {
	t.Parallel()
	// Given
	a := es.StatsAgg("price").
		Meta("k1", "v1").
		Meta("k2", "v2")

	// When Then
	assert.NotNil(t, a)
	bodyJSON := assert.MarshalWithoutError(t, a)
	assert.Equal(t, "{\"stats\":{\"field\":\"price\",\"meta\":{\"k1\":\"v1\",\"k2\":\"v2\"}}}", bodyJSON)
}

func Test_StatsAgg_should_have_Aggs_method(t *testing.T) {
	t.Parallel()
	// Given
	a := es.StatsAgg("price")

	// When Then
	assert.NotNil(t, a.Aggs)
}

func Test_Aggs_should_add_aggs_field_into_StatsAgg_when_not_empty(t *testing.T) {
	t.Parallel()
	// Given
	a := es.StatsAgg("price").Aggs(es.Agg("stats_stock", es.StatsAgg("stock")))

	// When Then
	assert.NotNil(t, a)
	bodyJSON := assert.MarshalWithoutError(t, a)
	assert.Equal(t, "{\"aggs\":{\"stats_stock\":{\"stats\":{\"field\":\"stock\"}}},\"stats\":{\"field\":\"price\"}}", bodyJSON)
}

func Test_Aggs_should_not_add_aggs_field_into_StatsAgg_when_it_is_empty(t *testing.T) {
	t.Parallel()
	// Given
	a := es.StatsAgg("price").Aggs(nil)

	// When Then
	assert.NotNil(t, a)
	bodyJSON := assert.MarshalWithoutError(t, a)
	assert.Equal(t, "{\"stats\":{\"field\":\"price\"}}", bodyJSON)
}
