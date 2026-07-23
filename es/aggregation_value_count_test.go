package es_test

import (
	"testing"

	ScriptLanguage "github.com/Trendyol/es-query-builder/es/enums/script-language"

	"github.com/Trendyol/es-query-builder/es"
	"github.com/Trendyol/es-query-builder/test/assert"
)

func Test_ValueCountAgg_should_exist_on_es_package(t *testing.T) {
	t.Parallel()
	// Given When Then
	assert.NotNil(t, es.ValueCountAgg)
}

func Test_ValueCountAgg_should_return_type_of_valueCountAggType(t *testing.T) {
	t.Parallel()
	// Given
	aggsQuery := es.ValueCountAgg("price")

	// When Then
	assert.NotNil(t, aggsQuery)
	assert.IsTypeString(t, "es.valueCountAggType", aggsQuery)
	assert.MarshalWithoutError(t, aggsQuery)
}

func Test_ValueCountAgg_should_create_json_with_value_count_field(t *testing.T) {
	t.Parallel()
	// Given
	a := es.ValueCountAgg("price")

	// When Then
	assert.NotNil(t, a)
	bodyJSON := assert.MarshalWithoutError(t, a)
	assert.Equal(t, "{\"value_count\":{\"field\":\"price\"}}", bodyJSON)
}

func Test_ValueCountAgg_should_have_Missing_method(t *testing.T) {
	t.Parallel()
	// Given
	a := es.ValueCountAgg("price")

	// When Then
	assert.NotNil(t, a.Missing)
}

func Test_Missing_should_add_missing_field_into_ValueCountAgg(t *testing.T) {
	t.Parallel()
	// Given
	a := es.ValueCountAgg("price").Missing(0)

	// When Then
	assert.NotNil(t, a)
	bodyJSON := assert.MarshalWithoutError(t, a)
	assert.Equal(t, "{\"value_count\":{\"field\":\"price\",\"missing\":0}}", bodyJSON)
}

func Test_ValueCountAgg_should_have_Script_method(t *testing.T) {
	t.Parallel()
	// Given
	a := es.ValueCountAgg("price")

	// When Then
	assert.NotNil(t, a.Script)
}

func Test_Script_should_add_script_field_into_ValueCountAgg(t *testing.T) {
	t.Parallel()
	// Given
	a := es.ValueCountAgg("price").Script(es.ScriptID("id_12345", ScriptLanguage.Painless))

	// When Then
	assert.NotNil(t, a)
	bodyJSON := assert.MarshalWithoutError(t, a)
	assert.Equal(t, "{\"value_count\":{\"field\":\"price\",\"script\":{\"id\":\"id_12345\",\"lang\":\"painless\"}}}", bodyJSON)
}

func Test_ValueCountAgg_should_have_Meta_method(t *testing.T) {
	t.Parallel()
	// Given
	a := es.ValueCountAgg("price")

	// When Then
	assert.NotNil(t, a.Meta)
}

func Test_Meta_should_add_meta_field_into_ValueCountAgg(t *testing.T) {
	t.Parallel()
	// Given
	a := es.ValueCountAgg("price").
		Meta("k1", "v1").
		Meta("k2", "v2")

	// When Then
	assert.NotNil(t, a)
	bodyJSON := assert.MarshalWithoutError(t, a)
	assert.Equal(t, "{\"meta\":{\"k1\":\"v1\",\"k2\":\"v2\"},\"value_count\":{\"field\":\"price\"}}", bodyJSON)
}

func Test_ValueCountAgg_should_have_Aggs_method(t *testing.T) {
	t.Parallel()
	// Given
	a := es.ValueCountAgg("price")

	// When Then
	assert.NotNil(t, a.Aggs)
}

func Test_Aggs_should_add_aggs_field_into_ValueCountAgg_when_not_empty(t *testing.T) {
	t.Parallel()
	// Given
	a := es.ValueCountAgg("price").Aggs(es.Agg("by_category", es.TermsAgg("category")))

	// When Then
	assert.NotNil(t, a)
	bodyJSON := assert.MarshalWithoutError(t, a)
	assert.Equal(t, "{\"aggs\":{\"by_category\":{\"terms\":{\"field\":\"category\"}}},\"value_count\":{\"field\":\"price\"}}", bodyJSON)
}

func Test_Aggs_should_not_add_aggs_field_into_ValueCountAgg_when_it_is_empty(t *testing.T) {
	t.Parallel()
	// Given
	a := es.ValueCountAgg("price").Aggs(nil)

	// When Then
	assert.NotNil(t, a)
	bodyJSON := assert.MarshalWithoutError(t, a)
	assert.Equal(t, "{\"value_count\":{\"field\":\"price\"}}", bodyJSON)
}
