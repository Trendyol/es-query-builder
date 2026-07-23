package es_test

import (
	"testing"

	Order "github.com/Trendyol/es-query-builder/es/enums/sort/order"

	"github.com/Trendyol/es-query-builder/es"
	"github.com/Trendyol/es-query-builder/test/assert"
)

func Test_TopHitsAgg_should_exist_on_es_package(t *testing.T) {
	t.Parallel()
	// Given When Then
	assert.NotNil(t, es.TopHitsAgg)
}

func Test_TopHitsAgg_should_return_type_of_topHitsAggType(t *testing.T) {
	t.Parallel()
	// Given
	aggsQuery := es.TopHitsAgg()

	// When Then
	assert.NotNil(t, aggsQuery)
	assert.IsTypeString(t, "es.topHitsAggType", aggsQuery)
	assert.MarshalWithoutError(t, aggsQuery)
}

func Test_TopHitsAgg_should_create_json_with_top_hits_field(t *testing.T) {
	t.Parallel()
	// Given
	a := es.TopHitsAgg()

	// When Then
	assert.NotNil(t, a)
	bodyJSON := assert.MarshalWithoutError(t, a)
	assert.Equal(t, "{\"top_hits\":{}}", bodyJSON)
}

func Test_TopHitsAgg_should_have_Size_method(t *testing.T) {
	t.Parallel()
	// Given
	a := es.TopHitsAgg()

	// When Then
	assert.NotNil(t, a.Size)
}

func Test_Size_should_add_size_field_into_TopHitsAgg(t *testing.T) {
	t.Parallel()
	// Given
	a := es.TopHitsAgg().Size(3)

	// When Then
	assert.NotNil(t, a)
	bodyJSON := assert.MarshalWithoutError(t, a)
	assert.Equal(t, "{\"top_hits\":{\"size\":3}}", bodyJSON)
}

func Test_TopHitsAgg_should_have_From_method(t *testing.T) {
	t.Parallel()
	// Given
	a := es.TopHitsAgg()

	// When Then
	assert.NotNil(t, a.From)
}

func Test_From_should_add_from_field_into_TopHitsAgg(t *testing.T) {
	t.Parallel()
	// Given
	a := es.TopHitsAgg().From(10)

	// When Then
	assert.NotNil(t, a)
	bodyJSON := assert.MarshalWithoutError(t, a)
	assert.Equal(t, "{\"top_hits\":{\"from\":10}}", bodyJSON)
}

func Test_TopHitsAgg_should_have_Sort_method(t *testing.T) {
	t.Parallel()
	// Given
	a := es.TopHitsAgg()

	// When Then
	assert.NotNil(t, a.Sort)
}

func Test_Sort_should_add_sort_field_into_TopHitsAgg(t *testing.T) {
	t.Parallel()
	// Given
	a := es.TopHitsAgg().Sort(es.Sort("date").Order(Order.Desc))

	// When Then
	assert.NotNil(t, a)
	bodyJSON := assert.MarshalWithoutError(t, a)
	assert.Equal(t, "{\"top_hits\":{\"sort\":[{\"date\":{\"order\":\"desc\"}}]}}", bodyJSON)
}

func Test_TopHitsAgg_should_have_SourceFalse_method(t *testing.T) {
	t.Parallel()
	// Given
	a := es.TopHitsAgg()

	// When Then
	assert.NotNil(t, a.SourceFalse)
}

func Test_SourceFalse_should_add_source_false_into_TopHitsAgg(t *testing.T) {
	t.Parallel()
	// Given
	a := es.TopHitsAgg().SourceFalse()

	// When Then
	assert.NotNil(t, a)
	bodyJSON := assert.MarshalWithoutError(t, a)
	assert.Equal(t, "{\"top_hits\":{\"_source\":false}}", bodyJSON)
}

func Test_TopHitsAgg_should_have_SourceIncludes_method(t *testing.T) {
	t.Parallel()
	// Given
	a := es.TopHitsAgg()

	// When Then
	assert.NotNil(t, a.SourceIncludes)
}

func Test_SourceIncludes_should_add_source_includes_into_TopHitsAgg(t *testing.T) {
	t.Parallel()
	// Given
	a := es.TopHitsAgg().SourceIncludes("title", "price")

	// When Then
	assert.NotNil(t, a)
	bodyJSON := assert.MarshalWithoutError(t, a)
	assert.Equal(t, "{\"top_hits\":{\"_source\":{\"includes\":[\"title\",\"price\"]}}}", bodyJSON)
}

func Test_TopHitsAgg_should_have_SourceExcludes_method(t *testing.T) {
	t.Parallel()
	// Given
	a := es.TopHitsAgg()

	// When Then
	assert.NotNil(t, a.SourceExcludes)
}

func Test_SourceExcludes_should_add_source_excludes_into_TopHitsAgg(t *testing.T) {
	t.Parallel()
	// Given
	a := es.TopHitsAgg().SourceExcludes("description")

	// When Then
	assert.NotNil(t, a)
	bodyJSON := assert.MarshalWithoutError(t, a)
	assert.Equal(t, "{\"top_hits\":{\"_source\":{\"excludes\":[\"description\"]}}}", bodyJSON)
}

func Test_TopHitsAgg_should_have_Highlight_method(t *testing.T) {
	t.Parallel()
	// Given
	a := es.TopHitsAgg()

	// When Then
	assert.NotNil(t, a.Highlight)
}

func Test_Highlight_should_add_highlight_field_into_TopHitsAgg(t *testing.T) {
	t.Parallel()
	// Given
	a := es.TopHitsAgg().Highlight(es.Highlight().PreTags("<em>"))

	// When Then
	assert.NotNil(t, a)
	bodyJSON := assert.MarshalWithoutError(t, a)
	assert.Equal(t, "{\"top_hits\":{\"highlight\":{\"pre_tags\":[\"\\u003cem\\u003e\"]}}}", bodyJSON)
}

func Test_TopHitsAgg_should_have_Explain_method(t *testing.T) {
	t.Parallel()
	// Given
	a := es.TopHitsAgg()

	// When Then
	assert.NotNil(t, a.Explain)
}

func Test_Explain_should_add_explain_field_into_TopHitsAgg(t *testing.T) {
	t.Parallel()
	// Given
	a := es.TopHitsAgg().Explain(true)

	// When Then
	assert.NotNil(t, a)
	bodyJSON := assert.MarshalWithoutError(t, a)
	assert.Equal(t, "{\"top_hits\":{\"explain\":true}}", bodyJSON)
}

func Test_TopHitsAgg_should_have_Version_method(t *testing.T) {
	t.Parallel()
	// Given
	a := es.TopHitsAgg()

	// When Then
	assert.NotNil(t, a.Version)
}

func Test_Version_should_add_version_field_into_TopHitsAgg(t *testing.T) {
	t.Parallel()
	// Given
	a := es.TopHitsAgg().Version(true)

	// When Then
	assert.NotNil(t, a)
	bodyJSON := assert.MarshalWithoutError(t, a)
	assert.Equal(t, "{\"top_hits\":{\"version\":true}}", bodyJSON)
}

func Test_TopHitsAgg_should_have_SeqNoPrimaryTerm_method(t *testing.T) {
	t.Parallel()
	// Given
	a := es.TopHitsAgg()

	// When Then
	assert.NotNil(t, a.SeqNoPrimaryTerm)
}

func Test_SeqNoPrimaryTerm_should_add_seq_no_primary_term_field_into_TopHitsAgg(t *testing.T) {
	t.Parallel()
	// Given
	a := es.TopHitsAgg().SeqNoPrimaryTerm(true)

	// When Then
	assert.NotNil(t, a)
	bodyJSON := assert.MarshalWithoutError(t, a)
	assert.Equal(t, "{\"top_hits\":{\"seq_no_primary_term\":true}}", bodyJSON)
}

func Test_TopHitsAgg_should_have_TrackScores_method(t *testing.T) {
	t.Parallel()
	// Given
	a := es.TopHitsAgg()

	// When Then
	assert.NotNil(t, a.TrackScores)
}

func Test_TrackScores_should_add_track_scores_field_into_TopHitsAgg(t *testing.T) {
	t.Parallel()
	// Given
	a := es.TopHitsAgg().TrackScores(true)

	// When Then
	assert.NotNil(t, a)
	bodyJSON := assert.MarshalWithoutError(t, a)
	assert.Equal(t, "{\"top_hits\":{\"track_scores\":true}}", bodyJSON)
}

func Test_TopHitsAgg_should_have_Meta_method(t *testing.T) {
	t.Parallel()
	// Given
	a := es.TopHitsAgg()

	// When Then
	assert.NotNil(t, a.Meta)
}

func Test_Meta_should_add_meta_field_into_TopHitsAgg(t *testing.T) {
	t.Parallel()
	// Given
	a := es.TopHitsAgg().
		Meta("k1", "v1").
		Meta("k2", "v2")

	// When Then
	assert.NotNil(t, a)
	bodyJSON := assert.MarshalWithoutError(t, a)
	assert.Equal(t, "{\"meta\":{\"k1\":\"v1\",\"k2\":\"v2\"},\"top_hits\":{}}", bodyJSON)
}

func Test_TopHitsAgg_inside_TermsAgg_should_create_correct_json(t *testing.T) {
	t.Parallel()
	// Given
	a := es.TermsAgg("brand").Size(10).Aggs(
		es.Agg("sample", es.TopHitsAgg().Size(3).Sort(es.Sort("date").Order(Order.Desc))),
	)

	// When Then
	assert.NotNil(t, a)
	bodyJSON := assert.MarshalWithoutError(t, a)
	// nolint:golint,lll
	assert.Equal(t, "{\"aggs\":{\"sample\":{\"top_hits\":{\"size\":3,\"sort\":[{\"date\":{\"order\":\"desc\"}}]}}},\"terms\":{\"field\":\"brand\",\"size\":10}}", bodyJSON)
}

func Test_TopHitsAgg_should_create_json_with_all_fields(t *testing.T) {
	t.Parallel()
	// Given
	a := es.TopHitsAgg().
		Size(3).
		From(0).
		Sort(es.Sort("date").Order(Order.Desc)).
		SourceIncludes("title", "price").
		Explain(false).
		Version(false).
		SeqNoPrimaryTerm(false).
		TrackScores(true)

	// When Then
	assert.NotNil(t, a)
	bodyJSON := assert.MarshalWithoutError(t, a)
	// nolint:golint,lll
	assert.Equal(t, "{\"top_hits\":{\"_source\":{\"includes\":[\"title\",\"price\"]},\"explain\":false,\"from\":0,\"seq_no_primary_term\":false,\"size\":3,\"sort\":[{\"date\":{\"order\":\"desc\"}}],\"track_scores\":true,\"version\":false}}", bodyJSON)
}
