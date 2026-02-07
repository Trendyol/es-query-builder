package es_test

import (
	"testing"

	"github.com/Trendyol/es-query-builder/es"
	"github.com/Trendyol/es-query-builder/test/assert"
)

////   ReverseNestedAgg   ////

func Test_ReverseNestedAgg_should_exist_on_es_package(t *testing.T) {
	t.Parallel()
	// Given When Then
	assert.NotNil(t, es.ReverseNestedAgg)
}

func Test_ReverseNestedAgg_method_should_create_reverseNestedAggType(t *testing.T) {
	t.Parallel()
	// Given
	agg := es.ReverseNestedAgg()

	// Then
	assert.NotNil(t, agg)
	assert.IsTypeString(t, "es.reverseNestedAggType", agg)
}

func Test_ReverseNestedAgg_should_create_json_with_reverse_nested(t *testing.T) {
	t.Parallel()
	// Given
	agg := es.ReverseNestedAgg()

	// When Then
	assert.NotNil(t, agg)
	bodyJSON := assert.MarshalWithoutError(t, agg)
	assert.Equal(t, "{\"reverse_nested\":{}}", bodyJSON)
}

////   ReverseNestedAgg Path   ////

func Test_ReverseNestedAgg_should_have_Path_method(t *testing.T) {
	t.Parallel()
	// Given
	agg := es.ReverseNestedAgg()

	// When Then
	assert.NotNil(t, agg.Path)
}

func Test_ReverseNestedAgg_Path_should_create_json_with_path_field(t *testing.T) {
	t.Parallel()
	// Given
	agg := es.ReverseNestedAgg().Path("comments")

	// When Then
	assert.NotNil(t, agg)
	bodyJSON := assert.MarshalWithoutError(t, agg)
	assert.Equal(t, "{\"reverse_nested\":{\"path\":\"comments\"}}", bodyJSON)
}

////   ReverseNestedAgg Aggs   ////

func Test_ReverseNestedAgg_should_have_Aggs_method(t *testing.T) {
	t.Parallel()
	// Given
	agg := es.ReverseNestedAgg()

	// When Then
	assert.NotNil(t, agg.Aggs)
}

func Test_ReverseNestedAgg_Aggs_should_create_json_with_sub_aggregations(t *testing.T) {
	t.Parallel()
	// Given
	agg := es.ReverseNestedAgg().
		Aggs(es.Agg("top_tags", es.TermsAgg("tags")))

	// When Then
	assert.NotNil(t, agg)
	bodyJSON := assert.MarshalWithoutError(t, agg)
	assert.Equal(t, "{\"aggs\":{\"top_tags\":{\"terms\":{\"field\":\"tags\"}}},\"reverse_nested\":{}}", bodyJSON)
}

func Test_ReverseNestedAgg_inside_nested_agg_should_create_correct_json(t *testing.T) {
	t.Parallel()
	// Given
	query := es.NewQuery(nil).
		Aggs(es.Agg("comments",
			es.NestedAgg("comments").
				Aggs(
					es.Agg("top_usernames", es.TermsAgg("comments.username")),
					es.Agg("comment_to_issue",
						es.ReverseNestedAgg().
							Aggs(es.Agg("top_tags_per_comment", es.TermsAgg("tags"))),
					),
				),
		))

	// When Then
	assert.NotNil(t, query)
	bodyJSON := assert.MarshalWithoutError(t, query)
	// nolint:golint,lll
	assert.Equal(t, "{\"aggs\":{\"comments\":{\"aggs\":{\"comment_to_issue\":{\"aggs\":{\"top_tags_per_comment\":{\"terms\":{\"field\":\"tags\"}}},\"reverse_nested\":{}},\"top_usernames\":{\"terms\":{\"field\":\"comments.username\"}}},\"nested\":{\"path\":\"comments\"}}},\"query\":{}}", bodyJSON)
}

////   ReverseNestedAgg Meta   ////

func Test_ReverseNestedAgg_should_have_Meta_method(t *testing.T) {
	t.Parallel()
	// Given
	agg := es.ReverseNestedAgg()

	// When Then
	assert.NotNil(t, agg.Meta)
}

func Test_ReverseNestedAgg_Meta_should_create_json_with_meta_field(t *testing.T) {
	t.Parallel()
	// Given
	agg := es.ReverseNestedAgg().Meta("color", "blue")

	// When Then
	assert.NotNil(t, agg)
	bodyJSON := assert.MarshalWithoutError(t, agg)
	assert.Equal(t, "{\"meta\":{\"color\":\"blue\"},\"reverse_nested\":{}}", bodyJSON)
}
