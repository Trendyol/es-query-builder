package es_test

import (
	"testing"

	ScriptLanguage "github.com/Trendyol/es-query-builder/es/enums/script-language"

	"github.com/Trendyol/es-query-builder/es"
	"github.com/Trendyol/es-query-builder/test/assert"
)

func Test_SumAgg_should_exist_on_es_package(t *testing.T) {
	t.Parallel()
	// Given When Then
	assert.NotNil(t, es.SumAgg)
}

func Test_SumAgg_should_return_type_of_sumAggType(t *testing.T) {
	t.Parallel()
	// Given
	aggsQuery := es.SumAgg("price")

	// When Then
	assert.NotNil(t, aggsQuery)
	assert.IsTypeString(t, "es.sumAggType", aggsQuery)
	assert.MarshalWithoutError(t, aggsQuery)
}

func Test_SumAgg_should_create_json_with_terms_field_inside(t *testing.T) {
	t.Parallel()
	// Given
	a := es.SumAgg("price")

	// When Then
	assert.NotNil(t, a)
	bodyJSON := assert.MarshalWithoutError(t, a)
	assert.Equal(t, "{\"sum\":{\"field\":\"price\"}}", bodyJSON)
}

func Test_SumAgg_should_have_Missing_method(t *testing.T) {
	t.Parallel()
	// Given
	a := es.SumAgg("price")

	// When Then
	assert.NotNil(t, a.Missing)
}

func Test_Missing_should_add_missing_field_into_SumAgg(t *testing.T) {
	t.Parallel()
	// Given
	a := es.SumAgg("price").Missing("missing_name")

	// When Then
	assert.NotNil(t, a)
	bodyJSON := assert.MarshalWithoutError(t, a)
	assert.Equal(t, "{\"sum\":{\"field\":\"price\",\"missing\":\"missing_name\"}}", bodyJSON)
}

func Test_SumAgg_should_have_Script_method(t *testing.T) {
	t.Parallel()
	// Given
	a := es.SumAgg("price")

	// When Then
	assert.NotNil(t, a.Script)
}

func Test_Script_should_add_script_field_into_SumAgg(t *testing.T) {
	t.Parallel()
	// Given
	a := es.SumAgg("price").Script(es.ScriptID("id_12345", ScriptLanguage.Painless))

	// When Then
	assert.NotNil(t, a)
	bodyJSON := assert.MarshalWithoutError(t, a)
	assert.Equal(t, "{\"sum\":{\"field\":\"price\",\"script\":{\"id\":\"id_12345\",\"lang\":\"painless\"}}}", bodyJSON)
}

func Test_SumAgg_should_have_Format_method(t *testing.T) {
	t.Parallel()
	// Given
	a := es.SumAgg("price")

	// When Then
	assert.NotNil(t, a.Format)
}

func Test_Format_should_add_format_field_into_SumAgg(t *testing.T) {
	t.Parallel()
	// Given
	a := es.SumAgg("price").Format("#.00")

	// When Then
	assert.NotNil(t, a)
	bodyJSON := assert.MarshalWithoutError(t, a)
	assert.Equal(t, "{\"sum\":{\"field\":\"price\",\"format\":\"#.00\"}}", bodyJSON)
}

func Test_SumAgg_should_have_Meta_method(t *testing.T) {
	t.Parallel()
	// Given
	a := es.SumAgg("price")

	// When Then
	assert.NotNil(t, a.Meta)
}

func Test_Meta_should_add_meta_field_into_SumAgg(t *testing.T) {
	t.Parallel()
	// Given
	a := es.SumAgg("price").
		Meta("k1", "v1").
		Meta("k2", "v2")

	// When Then
	assert.NotNil(t, a)
	bodyJSON := assert.MarshalWithoutError(t, a)
	assert.Equal(t, "{\"sum\":{\"field\":\"price\",\"meta\":{\"k1\":\"v1\",\"k2\":\"v2\"}}}", bodyJSON)
}

func Test_SumAgg_should_have_Aggs_method(t *testing.T) {
	t.Parallel()
	// Given
	a := es.SumAgg("price")

	// When Then
	assert.NotNil(t, a.Aggs)
}

func Test_Aggs_should_add_aggs_field_into_SumAgg_when_not_empty(t *testing.T) {
	t.Parallel()
	// Given
	a := es.SumAgg("price").Aggs(es.Agg("sum_stock", es.SumAgg("stock")))

	// When Then
	assert.NotNil(t, a)
	bodyJSON := assert.MarshalWithoutError(t, a)
	assert.Equal(t, "{\"aggs\":{\"sum_stock\":{\"sum\":{\"field\":\"stock\"}}},\"sum\":{\"field\":\"price\"}}", bodyJSON)
}

func Test_Aggs_should_not_add_aggs_field_into_SumAgg_when_it_is_empty(t *testing.T) {
	t.Parallel()
	// Given
	a := es.SumAgg("price").Aggs(nil)

	// When Then
	assert.NotNil(t, a)
	bodyJSON := assert.MarshalWithoutError(t, a)
	assert.Equal(t, "{\"sum\":{\"field\":\"price\"}}", bodyJSON)
}
