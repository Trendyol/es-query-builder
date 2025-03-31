package es_test

import (
	"testing"

	ScriptLanguage "github.com/Trendyol/es-query-builder/es/enums/script-language"

	"github.com/Trendyol/es-query-builder/es"
	"github.com/Trendyol/es-query-builder/test/assert"
)

func Test_CardinalityAgg_should_exist_on_es_package(t *testing.T) {
	t.Parallel()
	// Given When Then
	assert.NotNil(t, es.CardinalityAgg)
}

func Test_CardinalityAgg_should_return_type_of_cardinalityAggType(t *testing.T) {
	t.Parallel()
	// Given
	aggsQuery := es.CardinalityAgg("price")

	// When Then
	assert.NotNil(t, aggsQuery)
	assert.IsTypeString(t, "es.cardinalityAggType", aggsQuery)
	assert.MarshalWithoutError(t, aggsQuery)
}

func Test_CardinalityAgg_should_create_json_with_terms_field_inside(t *testing.T) {
	t.Parallel()
	// Given
	a := es.CardinalityAgg("price")

	// When Then
	assert.NotNil(t, a)
	bodyJSON := assert.MarshalWithoutError(t, a)
	assert.Equal(t, "{\"cardinality\":{\"field\":\"price\"}}", bodyJSON)
}

func Test_CardinalityAgg_should_have_Missing_method(t *testing.T) {
	t.Parallel()
	// Given
	a := es.CardinalityAgg("price")

	// When Then
	assert.NotNil(t, a.Missing)
}

func Test_Missing_should_add_missing_field_into_CardinalityAgg(t *testing.T) {
	t.Parallel()
	// Given
	a := es.CardinalityAgg("price").Missing("missing_name")

	// When Then
	assert.NotNil(t, a)
	bodyJSON := assert.MarshalWithoutError(t, a)
	assert.Equal(t, "{\"cardinality\":{\"field\":\"price\",\"missing\":\"missing_name\"}}", bodyJSON)
}

func Test_CardinalityAgg_should_have_Script_method(t *testing.T) {
	t.Parallel()
	// Given
	a := es.CardinalityAgg("price")

	// When Then
	assert.NotNil(t, a.Script)
}

func Test_Script_should_add_script_field_into_CardinalityAgg(t *testing.T) {
	t.Parallel()
	// Given
	a := es.CardinalityAgg("price").Script(es.ScriptID("id_12345", ScriptLanguage.Painless))

	// When Then
	assert.NotNil(t, a)
	bodyJSON := assert.MarshalWithoutError(t, a)
	assert.Equal(t, "{\"cardinality\":{\"field\":\"price\",\"script\":{\"id\":\"id_12345\",\"lang\":\"painless\"}}}", bodyJSON)
}

func Test_CardinalityAgg_should_have_PrecisionThreshold_method(t *testing.T) {
	t.Parallel()
	// Given
	a := es.CardinalityAgg("price")

	// When Then
	assert.NotNil(t, a.PrecisionThreshold)
}

func Test_PrecisionThreshold_should_add_precision_threshold_field_into_CardinalityAgg(t *testing.T) {
	t.Parallel()
	// Given
	a := es.CardinalityAgg("price").PrecisionThreshold(100)

	// When Then
	assert.NotNil(t, a)
	bodyJSON := assert.MarshalWithoutError(t, a)
	assert.Equal(t, "{\"cardinality\":{\"field\":\"price\",\"precision_threshold\":100}}", bodyJSON)
}

func Test_CardinalityAgg_should_have_Meta_method(t *testing.T) {
	t.Parallel()
	// Given
	a := es.CardinalityAgg("price")

	// When Then
	assert.NotNil(t, a.Meta)
}

func Test_Meta_should_add_meta_field_into_CardinalityAgg(t *testing.T) {
	t.Parallel()
	// Given
	a := es.CardinalityAgg("price").
		Meta("k1", "v1").
		Meta("k2", "v2")

	// When Then
	assert.NotNil(t, a)
	bodyJSON := assert.MarshalWithoutError(t, a)
	assert.Equal(t, "{\"cardinality\":{\"field\":\"price\",\"meta\":{\"k1\":\"v1\",\"k2\":\"v2\"}}}", bodyJSON)
}

func Test_CardinalityAgg_should_have_Aggs_method(t *testing.T) {
	t.Parallel()
	// Given
	a := es.CardinalityAgg("price")

	// When Then
	assert.NotNil(t, a.Aggs)
}

func Test_Aggs_should_add_aggs_field_into_CardinalityAgg_when_not_empty(t *testing.T) {
	t.Parallel()
	// Given
	a := es.CardinalityAgg("price").Aggs(es.Agg("cardinality_stock", es.CardinalityAgg("stock")))

	// When Then
	assert.NotNil(t, a)
	bodyJSON := assert.MarshalWithoutError(t, a)
	// nolint:golint,lll
	assert.Equal(t, "{\"aggs\":{\"cardinality_stock\":{\"cardinality\":{\"field\":\"stock\"}}},\"cardinality\":{\"field\":\"price\"}}", bodyJSON)
}

func Test_Aggs_should_not_add_aggs_field_into_CardinalityAgg_when_it_is_empty(t *testing.T) {
	t.Parallel()
	// Given
	a := es.CardinalityAgg("price").Aggs(nil)

	// When Then
	assert.NotNil(t, a)
	bodyJSON := assert.MarshalWithoutError(t, a)
	assert.Equal(t, "{\"cardinality\":{\"field\":\"price\"}}", bodyJSON)
}
