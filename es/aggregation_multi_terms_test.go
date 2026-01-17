package es_test

import (
	"testing"

	CollectMode "github.com/Trendyol/es-query-builder/es/enums/collect-mode"
	ExecutionHint "github.com/Trendyol/es-query-builder/es/enums/execution-hint"
	ScriptLanguage "github.com/Trendyol/es-query-builder/es/enums/script-language"
	Order "github.com/Trendyol/es-query-builder/es/enums/sort/order"

	"github.com/Trendyol/es-query-builder/es"
	"github.com/Trendyol/es-query-builder/test/assert"
)

func Test_MultiTermsAgg_should_exist_on_es_package(t *testing.T) {
	t.Parallel()
	// Given When Then
	assert.NotNil(t, es.MultiTermsAgg)
}

func Test_MultiTermsAgg_should_return_type_of_multi_termsAggType(t *testing.T) {
	t.Parallel()
	// Given
	aggsQuery := es.MultiTermsAgg(
		es.TermAgg("price"),
		es.TermAgg("stock"),
	)

	// When Then
	assert.NotNil(t, aggsQuery)
	assert.IsTypeString(t, "es.multiTermsAggType", aggsQuery)
	assert.MarshalWithoutError(t, aggsQuery)
}

func Test_MultiTermsAgg_should_create_json_with_terms_field_inside(t *testing.T) {
	t.Parallel()
	// Given
	a := es.MultiTermsAgg(
		es.TermAgg("price"),
		es.TermAgg("stock"),
	)

	// When Then
	assert.NotNil(t, a)
	bodyJSON := assert.MarshalWithoutError(t, a)
	assert.Equal(t, "{\"multi_terms\":{\"terms\":[{\"field\":\"price\"},{\"field\":\"stock\"}]}}", bodyJSON)
}

func Test_MultiTermsAgg_should_have_Missing_method(t *testing.T) {
	t.Parallel()
	// Given
	a := es.MultiTermsAgg(
		es.TermAgg("price"),
		es.TermAgg("stock"),
	)

	// When Then
	assert.NotNil(t, a.Missing)
}

func Test_Missing_should_add_missing_field_into_MultiTermsAgg(t *testing.T) {
	t.Parallel()
	// Given
	a := es.MultiTermsAgg(
		es.TermAgg("price"),
		es.TermAgg("stock"),
	).Missing("missing_name")

	// When Then
	assert.NotNil(t, a)
	bodyJSON := assert.MarshalWithoutError(t, a)
	assert.Equal(t, "{\"multi_terms\":{\"missing\":\"missing_name\",\"terms\":[{\"field\":\"price\"},{\"field\":\"stock\"}]}}", bodyJSON)
}

func Test_MultiTermsAgg_should_have_Script_method(t *testing.T) {
	t.Parallel()
	// Given
	a := es.MultiTermsAgg(
		es.TermAgg("price"),
		es.TermAgg("stock"),
	)

	// When Then
	assert.NotNil(t, a.Script)
}

func Test_MultiTermsAgg_should_have_Meta_method(t *testing.T) {
	t.Parallel()
	// Given
	a := es.MultiTermsAgg(
		es.TermAgg("price"),
		es.TermAgg("stock"),
	)

	// When Then
	assert.NotNil(t, a.Meta)
}

func Test_MultiTermsAgg_add_meta_field_into_MultiTermAgg(t *testing.T) {
	t.Parallel()
	// Given
	a := es.MultiTermsAgg(
		es.TermAgg("price"),
		es.TermAgg("stock"),
	).
		Meta("k1", "v1").
		Meta("k2", "v2")

	// When Then
	assert.NotNil(t, a)
	bodyJSON := assert.MarshalWithoutError(t, a)
	assert.Equal(t, "{\"meta\":{\"k1\":\"v1\",\"k2\":\"v2\"},\"multi_terms\":{\"terms\":[{\"field\":\"price\"},{\"field\":\"stock\"}]}}", bodyJSON)
}

func Test_Script_should_add_script_field_into_MultiTermsAgg(t *testing.T) {
	t.Parallel()
	// Given
	a := es.MultiTermsAgg(
		es.TermAgg("price"),
		es.TermAgg("stock"),
	).Script(es.ScriptID("id_12345", ScriptLanguage.Painless))

	// When Then
	assert.NotNil(t, a)
	bodyJSON := assert.MarshalWithoutError(t, a)
	// nolint:golint,lll
	assert.Equal(t, "{\"multi_terms\":{\"script\":{\"id\":\"id_12345\",\"lang\":\"painless\"},\"terms\":[{\"field\":\"price\"},{\"field\":\"stock\"}]}}", bodyJSON)
}

func Test_MultiTermsAgg_should_have_Size_method(t *testing.T) {
	t.Parallel()
	// Given
	a := es.MultiTermsAgg(
		es.TermAgg("price"),
		es.TermAgg("stock"),
	)

	// When Then
	assert.NotNil(t, a.Size)
}

func Test_Size_should_add_size_field_into_MultiTermsAgg(t *testing.T) {
	t.Parallel()
	// Given
	a := es.MultiTermsAgg(
		es.TermAgg("price"),
		es.TermAgg("stock"),
	).Size(100)

	// When Then
	assert.NotNil(t, a)
	bodyJSON := assert.MarshalWithoutError(t, a)
	assert.Equal(t, "{\"multi_terms\":{\"size\":100,\"terms\":[{\"field\":\"price\"},{\"field\":\"stock\"}]}}", bodyJSON)
}

func Test_MultiTermsAgg_should_have_IgnoreUnmapped_method(t *testing.T) {
	t.Parallel()
	// Given
	a := es.MultiTermsAgg(
		es.TermAgg("price"),
		es.TermAgg("stock"),
	)

	// When Then
	assert.NotNil(t, a.IgnoreUnmapped)
}

func Test_IgnoreUnmapped_should_add_ignore_unmapped_field_into_MultiTermsAgg(t *testing.T) {
	t.Parallel()
	// Given
	a := es.MultiTermsAgg(
		es.TermAgg("price"),
		es.TermAgg("stock"),
	).IgnoreUnmapped(true)

	// When Then
	assert.NotNil(t, a)
	bodyJSON := assert.MarshalWithoutError(t, a)
	assert.Equal(t, "{\"multi_terms\":{\"ignore_unmapped\":true,\"terms\":[{\"field\":\"price\"},{\"field\":\"stock\"}]}}", bodyJSON)
}

func Test_MultiTermsAgg_should_have_ShardSize_method(t *testing.T) {
	t.Parallel()
	// Given
	a := es.MultiTermsAgg(
		es.TermAgg("price"),
		es.TermAgg("stock"),
	)

	// When Then
	assert.NotNil(t, a.ShardSize)
}

func Test_ShardSize_should_add_shard_size_field_into_MultiTermsAgg(t *testing.T) {
	t.Parallel()
	// Given
	a := es.MultiTermsAgg(
		es.TermAgg("price"),
		es.TermAgg("stock"),
	).ShardSize(100)

	// When Then
	assert.NotNil(t, a)
	bodyJSON := assert.MarshalWithoutError(t, a)
	assert.Equal(t, "{\"multi_terms\":{\"shard_size\":100,\"terms\":[{\"field\":\"price\"},{\"field\":\"stock\"}]}}", bodyJSON)
}

func Test_MultiTermsAgg_should_have_Include_method(t *testing.T) {
	t.Parallel()
	// Given
	a := es.MultiTermsAgg(
		es.TermAgg("price"),
		es.TermAgg("stock"),
	)

	// When Then
	assert.NotNil(t, a.Include)
}

func Test_Include_should_add_include_field_into_MultiTermsAgg(t *testing.T) {
	t.Parallel()
	// Given
	a := es.MultiTermsAgg(
		es.TermAgg("price"),
		es.TermAgg("stock"),
	).Include("hello", "world")

	// When Then
	assert.NotNil(t, a)
	bodyJSON := assert.MarshalWithoutError(t, a)
	assert.Equal(t, "{\"multi_terms\":{\"include\":[\"hello\",\"world\"],\"terms\":[{\"field\":\"price\"},{\"field\":\"stock\"}]}}", bodyJSON)
}

func Test_MultiTermsAgg_should_have_Exclude_method(t *testing.T) {
	t.Parallel()
	// Given
	a := es.MultiTermsAgg(
		es.TermAgg("price"),
		es.TermAgg("stock"),
	)

	// When Then
	assert.NotNil(t, a.Exclude)
}

func Test_Exclude_should_add_exclude_field_into_MultiTermsAgg(t *testing.T) {
	t.Parallel()
	// Given
	a := es.MultiTermsAgg(
		es.TermAgg("price"),
		es.TermAgg("stock"),
	).Exclude("hello", "world")

	// When Then
	assert.NotNil(t, a)
	bodyJSON := assert.MarshalWithoutError(t, a)
	assert.Equal(t, "{\"multi_terms\":{\"exclude\":[\"hello\",\"world\"],\"terms\":[{\"field\":\"price\"},{\"field\":\"stock\"}]}}", bodyJSON)
}

func Test_MultiTermsAgg_should_have_MinDocCount_method(t *testing.T) {
	t.Parallel()
	// Given
	a := es.MultiTermsAgg(
		es.TermAgg("price"),
		es.TermAgg("stock"),
	)

	// When Then
	assert.NotNil(t, a.MinDocCount)
}

func Test_MinDocCount_should_add_min_doc_count_field_into_MultiTermsAgg(t *testing.T) {
	t.Parallel()
	// Given
	a := es.MultiTermsAgg(
		es.TermAgg("price"),
		es.TermAgg("stock"),
	).MinDocCount(250)

	// When Then
	assert.NotNil(t, a)
	bodyJSON := assert.MarshalWithoutError(t, a)
	assert.Equal(t, "{\"multi_terms\":{\"min_doc_count\":250,\"terms\":[{\"field\":\"price\"},{\"field\":\"stock\"}]}}", bodyJSON)
}

func Test_MultiTermsAgg_should_have_ExecutionHint_method(t *testing.T) {
	t.Parallel()
	// Given
	a := es.MultiTermsAgg(
		es.TermAgg("price"),
		es.TermAgg("stock"),
	)

	// When Then
	assert.NotNil(t, a.ExecutionHint)
}

func Test_ExecutionHint_should_add_execution_hint_field_into_MultiTermsAgg(t *testing.T) {
	t.Parallel()
	// Given
	a := es.MultiTermsAgg(
		es.TermAgg("price"),
		es.TermAgg("stock"),
	).ExecutionHint(ExecutionHint.Map)

	// When Then
	assert.NotNil(t, a)
	bodyJSON := assert.MarshalWithoutError(t, a)
	assert.Equal(t, "{\"multi_terms\":{\"execution_hint\":\"map\",\"terms\":[{\"field\":\"price\"},{\"field\":\"stock\"}]}}", bodyJSON)
}

func Test_MultiTermsAgg_should_have_CollectMode_method(t *testing.T) {
	t.Parallel()
	// Given
	a := es.MultiTermsAgg(
		es.TermAgg("price"),
		es.TermAgg("stock"),
	)

	// When Then
	assert.NotNil(t, a.CollectMode)
}

func Test_CollectMode_should_add_collect_mode_field_into_MultiTermsAgg(t *testing.T) {
	t.Parallel()
	// Given
	a := es.MultiTermsAgg(
		es.TermAgg("price"),
		es.TermAgg("stock"),
	).CollectMode(CollectMode.BreadthFirst)

	// When Then
	assert.NotNil(t, a)
	bodyJSON := assert.MarshalWithoutError(t, a)
	assert.Equal(t, "{\"multi_terms\":{\"collect_mode\":\"breadth_first\",\"terms\":[{\"field\":\"price\"},{\"field\":\"stock\"}]}}", bodyJSON)
}

func Test_MultiTermsAgg_should_have_Order_method(t *testing.T) {
	t.Parallel()
	// Given
	a := es.MultiTermsAgg(
		es.TermAgg("price"),
		es.TermAgg("stock"),
	)

	// When Then
	assert.NotNil(t, a.Order)
}

func Test_Order_should_add_order_field_into_MultiTermsAgg_when_not_empty(t *testing.T) {
	t.Parallel()
	// Given
	a := es.MultiTermsAgg(
		es.TermAgg("price"),
		es.TermAgg("stock"),
	).Order(
		es.AggOrder("price", Order.Desc),
		es.AggOrder("stock", Order.Asc),
	)

	// When Then
	assert.NotNil(t, a)
	bodyJSON := assert.MarshalWithoutError(t, a)
	// nolint:golint,lll
	assert.Equal(t, "{\"multi_terms\":{\"order\":[{\"price\":\"desc\"},{\"stock\":\"asc\"}],\"terms\":[{\"field\":\"price\"},{\"field\":\"stock\"}]}}", bodyJSON)
}

func Test_Order_should_not_add_order_field_into_MultiTermsAgg_when_it_is_empty(t *testing.T) {
	t.Parallel()
	// Given
	a := es.MultiTermsAgg(
		es.TermAgg("price"),
		es.TermAgg("stock"),
	).Order(nil)

	// When Then
	assert.NotNil(t, a)
	bodyJSON := assert.MarshalWithoutError(t, a)
	assert.Equal(t, "{\"multi_terms\":{\"terms\":[{\"field\":\"price\"},{\"field\":\"stock\"}]}}", bodyJSON)
}

func Test_MultiTermsAgg_should_have_Aggs_method(t *testing.T) {
	t.Parallel()
	// Given
	a := es.MultiTermsAgg(
		es.TermAgg("price"),
		es.TermAgg("stock"),
	)

	// When Then
	assert.NotNil(t, a.Aggs)
}

func Test_Aggs_should_add_aggs_field_into_MultiTermsAgg_when_not_empty(t *testing.T) {
	t.Parallel()
	// Given
	a := es.MultiTermsAgg(
		es.TermAgg("price"),
		es.TermAgg("stock"),
	).Aggs(
		es.Agg("multi_terms_stock",
			es.MultiTermsAgg(es.TermAgg("listing")),
		),
	)

	// When Then
	assert.NotNil(t, a)
	bodyJSON := assert.MarshalWithoutError(t, a)
	// nolint:golint,lll
	assert.Equal(t, "{\"aggs\":{\"multi_terms_stock\":{\"multi_terms\":{\"terms\":[{\"field\":\"listing\"}]}}},\"multi_terms\":{\"terms\":[{\"field\":\"price\"},{\"field\":\"stock\"}]}}", bodyJSON)
}

func Test_Aggs_should_not_add_aggs_field_into_MultiTermsAgg_when_it_is_empty(t *testing.T) {
	t.Parallel()
	// Given
	a := es.MultiTermsAgg(
		es.TermAgg("price"),
		es.TermAgg("stock"),
	).Aggs(nil)

	// When Then
	assert.NotNil(t, a)
	bodyJSON := assert.MarshalWithoutError(t, a)
	assert.Equal(t, "{\"multi_terms\":{\"terms\":[{\"field\":\"price\"},{\"field\":\"stock\"}]}}", bodyJSON)
}
