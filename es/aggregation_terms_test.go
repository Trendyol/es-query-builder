package es_test

import (
	"testing"

	CollectMode "github.com/Trendyol/es-query-builder/es/enums/collect-mode"
	ExecutionHint "github.com/Trendyol/es-query-builder/es/enums/execution-hint"
	Order "github.com/Trendyol/es-query-builder/es/enums/sort/order"
	ScriptLanguage "github.com/Trendyol/es-query-builder/es/enums/script-language"

	"github.com/Trendyol/es-query-builder/es"
	"github.com/Trendyol/es-query-builder/test/assert"
)

func Test_TermsAgg_should_exist_on_es_package(t *testing.T) {
	t.Parallel()
	// Given When Then
	assert.NotNil(t, es.TermsAgg)
}

func Test_TermsAgg_should_return_type_of_termsAggType(t *testing.T) {
	t.Parallel()
	// Given
	aggsQuery := es.TermsAgg("price")

	// When Then
	assert.NotNil(t, aggsQuery)
	assert.IsTypeString(t, "es.termsAggType", aggsQuery)
	assert.MarshalWithoutError(t, aggsQuery)
}

func Test_TermsAgg_should_create_json_with_terms_field_inside(t *testing.T) {
	t.Parallel()
	// Given
	a := es.TermsAgg("price")

	// When Then
	assert.NotNil(t, a)
	bodyJSON := assert.MarshalWithoutError(t, a)
	assert.Equal(t, "{\"terms\":{\"field\":\"price\"}}", bodyJSON)
}

func Test_TermsAgg_should_have_Missing_method(t *testing.T) {
	t.Parallel()
	// Given
	a := es.TermsAgg("price")

	// When Then
	assert.NotNil(t, a.Missing)
}

func Test_Missing_should_add_missing_field_into_TermsAgg(t *testing.T) {
	t.Parallel()
	// Given
	a := es.TermsAgg("price").
		Missing("missing_name")

	// When Then
	assert.NotNil(t, a)
	bodyJSON := assert.MarshalWithoutError(t, a)
	assert.Equal(t, "{\"terms\":{\"field\":\"price\",\"missing\":\"missing_name\"}}", bodyJSON)
}

func Test_TermsAgg_should_have_Script_method(t *testing.T) {
	t.Parallel()
	// Given
	a := es.TermsAgg("price")

	// When Then
	assert.NotNil(t, a.Script)
}

func Test_Script_should_add_script_field_into_ermsAgg(t *testing.T) {
	t.Parallel()
	// Given
	a := es.TermsAgg("price").
		Script(es.ScriptID("id_12345", ScriptLanguage.Painless))

	// When Then
	assert.NotNil(t, a)
	bodyJSON := assert.MarshalWithoutError(t, a)
	// nolint:golint,lll
	assert.Equal(t, "{\"terms\":{\"field\":\"price\",\"script\":{\"id\":\"id_12345\",\"lang\":\"painless\"}}}", bodyJSON)
}

func Test_TermsAgg_should_have_ShowTermDocCountError_method(t *testing.T) {
	t.Parallel()
	// Given
	a := es.TermsAgg("price")

	// When Then
	assert.NotNil(t, a.ShowTermDocCountError)
}

func Test_ShowTermDocCountError_should_add_show_term_doc_count_error_field_into_TermsAgg(t *testing.T) {
	t.Parallel()
	// Given
	a := es.TermsAgg("price").ShowTermDocCountError(false)

	// When Then
	assert.NotNil(t, a)
	bodyJSON := assert.MarshalWithoutError(t, a)
	assert.Equal(t, "{\"terms\":{\"field\":\"price\",\"show_term_doc_count_error\":false}}", bodyJSON)
}

func Test_TermsAgg_should_have_Size_method(t *testing.T) {
	t.Parallel()
	// Given
	a := es.TermsAgg("price")

	// When Then
	assert.NotNil(t, a.Size)
}

func Test_Size_should_add_size_field_into_TermsAgg(t *testing.T) {
	t.Parallel()
	// Given
	a := es.TermsAgg("price").Size(100)

	// When Then
	assert.NotNil(t, a)
	bodyJSON := assert.MarshalWithoutError(t, a)
	assert.Equal(t, "{\"terms\":{\"field\":\"price\",\"size\":100}}", bodyJSON)
}

func Test_TermsAgg_should_have_ShardSize_method(t *testing.T) {
	t.Parallel()
	// Given
	a := es.TermsAgg("price")

	// When Then
	assert.NotNil(t, a.ShardSize)
}

func Test_ShardSize_should_add_shard_size_field_into_TermsAgg(t *testing.T) {
	t.Parallel()
	// Given
	a := es.TermsAgg("price").ShardSize(100)

	// When Then
	assert.NotNil(t, a)
	bodyJSON := assert.MarshalWithoutError(t, a)
	assert.Equal(t, "{\"terms\":{\"field\":\"price\",\"shard_size\":100}}", bodyJSON)
}

func Test_TermsAgg_should_have_Include_method(t *testing.T) {
	t.Parallel()
	// Given
	a := es.TermsAgg("price")

	// When Then
	assert.NotNil(t, a.Include)
}

func Test_Include_should_add_include_field_into_TermsAgg(t *testing.T) {
	t.Parallel()
	// Given
	a := es.TermsAgg("price").Include("hello", "world")

	// When Then
	assert.NotNil(t, a)
	bodyJSON := assert.MarshalWithoutError(t, a)
	assert.Equal(t, "{\"terms\":{\"field\":\"price\",\"include\":[\"hello\",\"world\"]}}", bodyJSON)
}

func Test_TermsAgg_should_have_Exclude_method(t *testing.T) {
	t.Parallel()
	// Given
	a := es.TermsAgg("price")

	// When Then
	assert.NotNil(t, a.Exclude)
}

func Test_Exclude_should_add_exclude_field_into_TermsAgg(t *testing.T) {
	t.Parallel()
	// Given
	a := es.TermsAgg("price").Exclude("hello", "world")

	// When Then
	assert.NotNil(t, a)
	bodyJSON := assert.MarshalWithoutError(t, a)
	assert.Equal(t, "{\"terms\":{\"exclude\":[\"hello\",\"world\"],\"field\":\"price\"}}", bodyJSON)
}

func Test_TermsAgg_should_have_MinDocCount_method(t *testing.T) {
	t.Parallel()
	// Given
	a := es.TermsAgg("price")

	// When Then
	assert.NotNil(t, a.MinDocCount)
}

func Test_MinDocCount_should_add_min_doc_count_field_into_TermsAgg(t *testing.T) {
	t.Parallel()
	// Given
	a := es.TermsAgg("price").MinDocCount(250)

	// When Then
	assert.NotNil(t, a)
	bodyJSON := assert.MarshalWithoutError(t, a)
	assert.Equal(t, "{\"terms\":{\"field\":\"price\",\"min_doc_count\":250}}", bodyJSON)
}

func Test_TermsAgg_should_have_ExecutionHint_method(t *testing.T) {
	t.Parallel()
	// Given
	a := es.TermsAgg("price")

	// When Then
	assert.NotNil(t, a.ExecutionHint)
}

func Test_ExecutionHint_should_add_execution_hint_field_into_TermsAgg(t *testing.T) {
	t.Parallel()
	// Given
	a := es.TermsAgg("price").ExecutionHint(ExecutionHint.Map)

	// When Then
	assert.NotNil(t, a)
	bodyJSON := assert.MarshalWithoutError(t, a)
	assert.Equal(t, "{\"terms\":{\"execution_hint\":\"map\",\"field\":\"price\"}}", bodyJSON)
}

func Test_TermsAgg_should_have_CollectMode_method(t *testing.T) {
	t.Parallel()
	// Given
	a := es.TermsAgg("price")

	// When Then
	assert.NotNil(t, a.CollectMode)
}

func Test_CollectMode_should_add_collect_mode_field_into_TermsAgg(t *testing.T) {
	t.Parallel()
	// Given
	a := es.TermsAgg("price").CollectMode(CollectMode.DepthFirst)

	// When Then
	assert.NotNil(t, a)
	bodyJSON := assert.MarshalWithoutError(t, a)
	assert.Equal(t, "{\"terms\":{\"collect_mode\":\"depth_first\",\"field\":\"price\"}}", bodyJSON)
}

func Test_TermsAgg_should_have_Order_method(t *testing.T) {
	t.Parallel()
	// Given
	a := es.TermsAgg("price")

	// When Then
	assert.NotNil(t, a.Order)
}

func Test_Order_should_add_order_field_into_TermsAgg_when_not_empty(t *testing.T) {
	t.Parallel()
	// Given
	a := es.TermsAgg("price").Order(
		es.AggOrder("price", Order.Desc),
		es.AggOrder("stock", Order.Asc),
	)

	// When Then
	assert.NotNil(t, a)
	bodyJSON := assert.MarshalWithoutError(t, a)
	assert.Equal(t, "{\"terms\":{\"field\":\"price\",\"order\":[{\"price\":\"desc\"},{\"stock\":\"asc\"}]}}", bodyJSON)
}

func Test_Order_should_not_add_order_field_into_TermsAgg_when_it_is_empty(t *testing.T) {
	t.Parallel()
	// Given
	a := es.TermsAgg("price").Order(nil)

	// When Then
	assert.NotNil(t, a)
	bodyJSON := assert.MarshalWithoutError(t, a)
	assert.Equal(t, "{\"terms\":{\"field\":\"price\"}}", bodyJSON)
}

func Test_TermsAgg_should_have_Aggs_method(t *testing.T) {
	t.Parallel()
	// Given
	a := es.TermsAgg("price")

	// When Then
	assert.NotNil(t, a.Aggs)
}

func Test_Aggs_should_add_aggs_field_into_TermsAgg_when_not_empty(t *testing.T) {
	t.Parallel()
	// Given
	a := es.TermsAgg("price").Aggs(
		es.Agg("terms_stock", es.TermsAgg("listing")),
	)

	// When Then
	assert.NotNil(t, a)
	bodyJSON := assert.MarshalWithoutError(t, a)
	assert.Equal(t, "{\"aggs\":{\"terms_stock\":{\"terms\":{\"field\":\"listing\"}}},\"terms\":{\"field\":\"price\"}}", bodyJSON)
}

func Test_Aggs_should_not_add_aggs_field_into_TermsAgg_when_it_is_empty(t *testing.T) {
	t.Parallel()
	// Given
	a := es.TermsAgg("price").Aggs(nil)

	// When Then
	assert.NotNil(t, a)
	bodyJSON := assert.MarshalWithoutError(t, a)
	assert.Equal(t, "{\"terms\":{\"field\":\"price\"}}", bodyJSON)
}
