package es_test

import (
	"testing"

	BoundaryScanner "github.com/Trendyol/es-query-builder/es/enums/boundary-scanner"
	Fragmenter "github.com/Trendyol/es-query-builder/es/enums/fragmenter"
	HighlighterType "github.com/Trendyol/es-query-builder/es/enums/highlighter-type"

	"github.com/Trendyol/es-query-builder/es"
	"github.com/Trendyol/es-query-builder/test/assert"
)

////   Highlight   ////

func Test_Highlight_should_exist_on_es_package(t *testing.T) {
	t.Parallel()
	// Given When Then
	assert.NotNil(t, es.Highlight)
}

func Test_Highlight_method_should_create_highlightType(t *testing.T) {
	t.Parallel()
	// Given
	h := es.Highlight()

	// Then
	assert.NotNil(t, h)
	assert.IsTypeString(t, "es.highlightType", h)
}

func Test_Highlight_should_create_json_with_highlight_field_inside_query(t *testing.T) {
	t.Parallel()
	// Given
	query := es.NewQuery(es.MatchAll()).
		Highlight(es.Highlight())

	// When Then
	assert.NotNil(t, query)
	bodyJSON := assert.MarshalWithoutError(t, query)
	assert.Equal(t, "{\"highlight\":{},\"query\":{\"match_all\":{}}}", bodyJSON)
}

////   Object Highlight method   ////

func Test_Object_should_have_Highlight_method(t *testing.T) {
	t.Parallel()
	// Given
	b := es.NewQuery(nil)

	// When Then
	assert.NotNil(t, b.Highlight)
}

////   PreTags   ////

func Test_Highlight_should_have_PreTags_method(t *testing.T) {
	t.Parallel()
	// Given
	h := es.Highlight()

	// When Then
	assert.NotNil(t, h.PreTags)
}

func Test_Highlight_PreTags_should_create_json_with_pre_tags_field(t *testing.T) {
	t.Parallel()
	// Given
	query := es.NewQuery(es.MatchAll()).
		Highlight(es.Highlight().PreTags("<em>"))

	// When Then
	assert.NotNil(t, query)
	bodyJSON := assert.MarshalWithoutError(t, query)
	assert.Equal(t, "{\"highlight\":{\"pre_tags\":[\"\\u003cem\\u003e\"]},\"query\":{\"match_all\":{}}}", bodyJSON)
}

func Test_Highlight_PreTags_should_create_json_with_multiple_pre_tags(t *testing.T) {
	t.Parallel()
	// Given
	query := es.NewQuery(es.MatchAll()).
		Highlight(es.Highlight().PreTags("<em>", "<strong>"))

	// When Then
	assert.NotNil(t, query)
	bodyJSON := assert.MarshalWithoutError(t, query)
	assert.Equal(t, "{\"highlight\":{\"pre_tags\":[\"\\u003cem\\u003e\",\"\\u003cstrong\\u003e\"]},\"query\":{\"match_all\":{}}}", bodyJSON)
}

////   PostTags   ////

func Test_Highlight_should_have_PostTags_method(t *testing.T) {
	t.Parallel()
	// Given
	h := es.Highlight()

	// When Then
	assert.NotNil(t, h.PostTags)
}

func Test_Highlight_PostTags_should_create_json_with_post_tags_field(t *testing.T) {
	t.Parallel()
	// Given
	query := es.NewQuery(es.MatchAll()).
		Highlight(es.Highlight().PostTags("</em>"))

	// When Then
	assert.NotNil(t, query)
	bodyJSON := assert.MarshalWithoutError(t, query)
	assert.Equal(t, "{\"highlight\":{\"post_tags\":[\"\\u003c/em\\u003e\"]},\"query\":{\"match_all\":{}}}", bodyJSON)
}

////   Field   ////

func Test_Highlight_should_have_Field_method(t *testing.T) {
	t.Parallel()
	// Given
	h := es.Highlight()

	// When Then
	assert.NotNil(t, h.Field)
}

func Test_Highlight_Field_should_create_json_with_fields_object(t *testing.T) {
	t.Parallel()
	// Given
	query := es.NewQuery(es.MatchAll()).
		Highlight(
			es.Highlight().
				Field(es.HighlightField("title")),
		)

	// When Then
	assert.NotNil(t, query)
	bodyJSON := assert.MarshalWithoutError(t, query)
	assert.Equal(t, "{\"highlight\":{\"fields\":{\"title\":{}}},\"query\":{\"match_all\":{}}}", bodyJSON)
}

func Test_Highlight_Field_should_create_json_with_multiple_fields(t *testing.T) {
	t.Parallel()
	// Given
	query := es.NewQuery(es.MatchAll()).
		Highlight(
			es.Highlight().
				Field(es.HighlightField("title")).
				Field(es.HighlightField("content")),
		)

	// When Then
	assert.NotNil(t, query)
	bodyJSON := assert.MarshalWithoutError(t, query)
	assert.Equal(t, "{\"highlight\":{\"fields\":{\"content\":{},\"title\":{}}},\"query\":{\"match_all\":{}}}", bodyJSON)
}

func Test_Highlight_Field_should_create_json_with_multiple_fields_in_single_call(t *testing.T) {
	t.Parallel()
	// Given
	query := es.NewQuery(es.MatchAll()).
		Highlight(
			es.Highlight().
				Field(
					es.HighlightField("title"),
					es.HighlightField("content"),
				),
		)

	// When Then
	assert.NotNil(t, query)
	bodyJSON := assert.MarshalWithoutError(t, query)
	assert.Equal(t, "{\"highlight\":{\"fields\":{\"content\":{},\"title\":{}}},\"query\":{\"match_all\":{}}}", bodyJSON)
}

////   Type   ////

func Test_Highlight_should_have_Type_method(t *testing.T) {
	t.Parallel()
	// Given
	h := es.Highlight()

	// When Then
	assert.NotNil(t, h.Type)
}

func Test_Highlight_Type_should_create_json_with_type_field(t *testing.T) {
	t.Parallel()
	// Given
	query := es.NewQuery(es.MatchAll()).
		Highlight(es.Highlight().Type(HighlighterType.Unified))

	// When Then
	assert.NotNil(t, query)
	bodyJSON := assert.MarshalWithoutError(t, query)
	assert.Equal(t, "{\"highlight\":{\"type\":\"unified\"},\"query\":{\"match_all\":{}}}", bodyJSON)
}

////   Order   ////

func Test_Highlight_should_have_Order_method(t *testing.T) {
	t.Parallel()
	// Given
	h := es.Highlight()

	// When Then
	assert.NotNil(t, h.Order)
}

func Test_Highlight_Order_should_create_json_with_order_field(t *testing.T) {
	t.Parallel()
	// Given
	query := es.NewQuery(es.MatchAll()).
		Highlight(es.Highlight().Order("score"))

	// When Then
	assert.NotNil(t, query)
	bodyJSON := assert.MarshalWithoutError(t, query)
	assert.Equal(t, "{\"highlight\":{\"order\":\"score\"},\"query\":{\"match_all\":{}}}", bodyJSON)
}

////   Encoder   ////

func Test_Highlight_should_have_Encoder_method(t *testing.T) {
	t.Parallel()
	// Given
	h := es.Highlight()

	// When Then
	assert.NotNil(t, h.Encoder)
}

func Test_Highlight_Encoder_should_create_json_with_encoder_field(t *testing.T) {
	t.Parallel()
	// Given
	query := es.NewQuery(es.MatchAll()).
		Highlight(es.Highlight().Encoder("html"))

	// When Then
	assert.NotNil(t, query)
	bodyJSON := assert.MarshalWithoutError(t, query)
	assert.Equal(t, "{\"highlight\":{\"encoder\":\"html\"},\"query\":{\"match_all\":{}}}", bodyJSON)
}

////   RequireFieldMatch   ////

func Test_Highlight_should_have_RequireFieldMatch_method(t *testing.T) {
	t.Parallel()
	// Given
	h := es.Highlight()

	// When Then
	assert.NotNil(t, h.RequireFieldMatch)
}

func Test_Highlight_RequireFieldMatch_should_create_json_with_require_field_match_field(t *testing.T) {
	t.Parallel()
	// Given
	query := es.NewQuery(es.MatchAll()).
		Highlight(es.Highlight().RequireFieldMatch(false))

	// When Then
	assert.NotNil(t, query)
	bodyJSON := assert.MarshalWithoutError(t, query)
	assert.Equal(t, "{\"highlight\":{\"require_field_match\":false},\"query\":{\"match_all\":{}}}", bodyJSON)
}

////   FragmentSize   ////

func Test_Highlight_should_have_FragmentSize_method(t *testing.T) {
	t.Parallel()
	// Given
	h := es.Highlight()

	// When Then
	assert.NotNil(t, h.FragmentSize)
}

func Test_Highlight_FragmentSize_should_create_json_with_fragment_size_field(t *testing.T) {
	t.Parallel()
	// Given
	query := es.NewQuery(es.MatchAll()).
		Highlight(es.Highlight().FragmentSize(150))

	// When Then
	assert.NotNil(t, query)
	bodyJSON := assert.MarshalWithoutError(t, query)
	assert.Equal(t, "{\"highlight\":{\"fragment_size\":150},\"query\":{\"match_all\":{}}}", bodyJSON)
}

////   NumberOfFragments   ////

func Test_Highlight_should_have_NumberOfFragments_method(t *testing.T) {
	t.Parallel()
	// Given
	h := es.Highlight()

	// When Then
	assert.NotNil(t, h.NumberOfFragments)
}

func Test_Highlight_NumberOfFragments_should_create_json_with_number_of_fragments_field(t *testing.T) {
	t.Parallel()
	// Given
	query := es.NewQuery(es.MatchAll()).
		Highlight(es.Highlight().NumberOfFragments(5))

	// When Then
	assert.NotNil(t, query)
	bodyJSON := assert.MarshalWithoutError(t, query)
	assert.Equal(t, "{\"highlight\":{\"number_of_fragments\":5},\"query\":{\"match_all\":{}}}", bodyJSON)
}

////   NoMatchSize   ////

func Test_Highlight_should_have_NoMatchSize_method(t *testing.T) {
	t.Parallel()
	// Given
	h := es.Highlight()

	// When Then
	assert.NotNil(t, h.NoMatchSize)
}

func Test_Highlight_NoMatchSize_should_create_json_with_no_match_size_field(t *testing.T) {
	t.Parallel()
	// Given
	query := es.NewQuery(es.MatchAll()).
		Highlight(es.Highlight().NoMatchSize(150))

	// When Then
	assert.NotNil(t, query)
	bodyJSON := assert.MarshalWithoutError(t, query)
	assert.Equal(t, "{\"highlight\":{\"no_match_size\":150},\"query\":{\"match_all\":{}}}", bodyJSON)
}

////   BoundaryScanner   ////

func Test_Highlight_should_have_BoundaryScanner_method(t *testing.T) {
	t.Parallel()
	// Given
	h := es.Highlight()

	// When Then
	assert.NotNil(t, h.BoundaryScanner)
}

func Test_Highlight_BoundaryScanner_should_create_json_with_boundary_scanner_field(t *testing.T) {
	t.Parallel()
	// Given
	query := es.NewQuery(es.MatchAll()).
		Highlight(es.Highlight().BoundaryScanner(BoundaryScanner.Sentence))

	// When Then
	assert.NotNil(t, query)
	bodyJSON := assert.MarshalWithoutError(t, query)
	assert.Equal(t, "{\"highlight\":{\"boundary_scanner\":\"sentence\"},\"query\":{\"match_all\":{}}}", bodyJSON)
}

////   BoundaryChars   ////

func Test_Highlight_should_have_BoundaryChars_method(t *testing.T) {
	t.Parallel()
	// Given
	h := es.Highlight()

	// When Then
	assert.NotNil(t, h.BoundaryChars)
}

func Test_Highlight_BoundaryChars_should_create_json_with_boundary_chars_field(t *testing.T) {
	t.Parallel()
	// Given
	query := es.NewQuery(es.MatchAll()).
		Highlight(es.Highlight().BoundaryChars(".,!?"))

	// When Then
	assert.NotNil(t, query)
	bodyJSON := assert.MarshalWithoutError(t, query)
	assert.Equal(t, "{\"highlight\":{\"boundary_chars\":\".,!?\"},\"query\":{\"match_all\":{}}}", bodyJSON)
}

////   BoundaryMaxScan   ////

func Test_Highlight_should_have_BoundaryMaxScan_method(t *testing.T) {
	t.Parallel()
	// Given
	h := es.Highlight()

	// When Then
	assert.NotNil(t, h.BoundaryMaxScan)
}

func Test_Highlight_BoundaryMaxScan_should_create_json_with_boundary_max_scan_field(t *testing.T) {
	t.Parallel()
	// Given
	query := es.NewQuery(es.MatchAll()).
		Highlight(es.Highlight().BoundaryMaxScan(20))

	// When Then
	assert.NotNil(t, query)
	bodyJSON := assert.MarshalWithoutError(t, query)
	assert.Equal(t, "{\"highlight\":{\"boundary_max_scan\":20},\"query\":{\"match_all\":{}}}", bodyJSON)
}

////   BoundaryScannerLocale   ////

func Test_Highlight_should_have_BoundaryScannerLocale_method(t *testing.T) {
	t.Parallel()
	// Given
	h := es.Highlight()

	// When Then
	assert.NotNil(t, h.BoundaryScannerLocale)
}

func Test_Highlight_BoundaryScannerLocale_should_create_json_with_boundary_scanner_locale_field(t *testing.T) {
	t.Parallel()
	// Given
	query := es.NewQuery(es.MatchAll()).
		Highlight(es.Highlight().BoundaryScannerLocale("en-US"))

	// When Then
	assert.NotNil(t, query)
	bodyJSON := assert.MarshalWithoutError(t, query)
	assert.Equal(t, "{\"highlight\":{\"boundary_scanner_locale\":\"en-US\"},\"query\":{\"match_all\":{}}}", bodyJSON)
}

////   Fragmenter   ////

func Test_Highlight_should_have_Fragmenter_method(t *testing.T) {
	t.Parallel()
	// Given
	h := es.Highlight()

	// When Then
	assert.NotNil(t, h.Fragmenter)
}

func Test_Highlight_Fragmenter_should_create_json_with_fragmenter_field(t *testing.T) {
	t.Parallel()
	// Given
	query := es.NewQuery(es.MatchAll()).
		Highlight(es.Highlight().Fragmenter(Fragmenter.Span))

	// When Then
	assert.NotNil(t, query)
	bodyJSON := assert.MarshalWithoutError(t, query)
	assert.Equal(t, "{\"highlight\":{\"fragmenter\":\"span\"},\"query\":{\"match_all\":{}}}", bodyJSON)
}

////   FragmentOffset   ////

func Test_Highlight_should_have_FragmentOffset_method(t *testing.T) {
	t.Parallel()
	// Given
	h := es.Highlight()

	// When Then
	assert.NotNil(t, h.FragmentOffset)
}

func Test_Highlight_FragmentOffset_should_create_json_with_fragment_offset_field(t *testing.T) {
	t.Parallel()
	// Given
	query := es.NewQuery(es.MatchAll()).
		Highlight(es.Highlight().FragmentOffset(10))

	// When Then
	assert.NotNil(t, query)
	bodyJSON := assert.MarshalWithoutError(t, query)
	assert.Equal(t, "{\"highlight\":{\"fragment_offset\":10},\"query\":{\"match_all\":{}}}", bodyJSON)
}

////   MaxFragmentLength   ////

func Test_Highlight_should_have_MaxFragmentLength_method(t *testing.T) {
	t.Parallel()
	// Given
	h := es.Highlight()

	// When Then
	assert.NotNil(t, h.MaxFragmentLength)
}

func Test_Highlight_MaxFragmentLength_should_create_json_with_max_fragment_length_field(t *testing.T) {
	t.Parallel()
	// Given
	query := es.NewQuery(es.MatchAll()).
		Highlight(es.Highlight().MaxFragmentLength(200))

	// When Then
	assert.NotNil(t, query)
	bodyJSON := assert.MarshalWithoutError(t, query)
	assert.Equal(t, "{\"highlight\":{\"max_fragment_length\":200},\"query\":{\"match_all\":{}}}", bodyJSON)
}

////   MaxAnalyzedOffset   ////

func Test_Highlight_should_have_MaxAnalyzedOffset_method(t *testing.T) {
	t.Parallel()
	// Given
	h := es.Highlight()

	// When Then
	assert.NotNil(t, h.MaxAnalyzedOffset)
}

func Test_Highlight_MaxAnalyzedOffset_should_create_json_with_max_analyzed_offset_field(t *testing.T) {
	t.Parallel()
	// Given
	query := es.NewQuery(es.MatchAll()).
		Highlight(es.Highlight().MaxAnalyzedOffset(1000000))

	// When Then
	assert.NotNil(t, query)
	bodyJSON := assert.MarshalWithoutError(t, query)
	assert.Equal(t, "{\"highlight\":{\"max_analyzed_offset\":1000000},\"query\":{\"match_all\":{}}}", bodyJSON)
}

////   HighlightQuery   ////

func Test_Highlight_should_have_HighlightQuery_method(t *testing.T) {
	t.Parallel()
	// Given
	h := es.Highlight()

	// When Then
	assert.NotNil(t, h.HighlightQuery)
}

func Test_Highlight_HighlightQuery_should_create_json_with_highlight_query_field(t *testing.T) {
	t.Parallel()
	// Given
	query := es.NewQuery(es.MatchAll()).
		Highlight(
			es.Highlight().
				HighlightQuery(es.Match("content", "elasticsearch")),
		)

	// When Then
	assert.NotNil(t, query)
	bodyJSON := assert.MarshalWithoutError(t, query)
	// nolint:golint,lll
	assert.Equal(t, "{\"highlight\":{\"highlight_query\":{\"match\":{\"content\":{\"query\":\"elasticsearch\"}}}},\"query\":{\"match_all\":{}}}", bodyJSON)
}

func Test_Highlight_HighlightQuery_should_handle_bool_query(t *testing.T) {
	t.Parallel()
	// Given
	query := es.NewQuery(es.MatchAll()).
		Highlight(
			es.Highlight().
				HighlightQuery(es.Bool().Must(es.Term("status", "active"))),
		)

	// When Then
	assert.NotNil(t, query)
	bodyJSON := assert.MarshalWithoutError(t, query)
	// nolint:golint,lll
	assert.Equal(t, "{\"highlight\":{\"highlight_query\":{\"bool\":{\"must\":[{\"term\":{\"status\":{\"value\":\"active\"}}}]}}},\"query\":{\"match_all\":{}}}", bodyJSON)
}

////   TagsSchema   ////

func Test_Highlight_should_have_TagsSchema_method(t *testing.T) {
	t.Parallel()
	// Given
	h := es.Highlight()

	// When Then
	assert.NotNil(t, h.TagsSchema)
}

func Test_Highlight_TagsSchema_should_create_json_with_tags_schema_field(t *testing.T) {
	t.Parallel()
	// Given
	query := es.NewQuery(es.MatchAll()).
		Highlight(es.Highlight().TagsSchema("styled"))

	// When Then
	assert.NotNil(t, query)
	bodyJSON := assert.MarshalWithoutError(t, query)
	assert.Equal(t, "{\"highlight\":{\"tags_schema\":\"styled\"},\"query\":{\"match_all\":{}}}", bodyJSON)
}

////   HighlightField   ////

func Test_HighlightField_should_exist_on_es_package(t *testing.T) {
	t.Parallel()
	// Given When Then
	assert.NotNil(t, es.HighlightField)
}

func Test_HighlightField_method_should_create_highlightFieldType(t *testing.T) {
	t.Parallel()
	// Given
	hf := es.HighlightField("title")

	// Then
	assert.NotNil(t, hf)
	assert.IsTypeString(t, "es.highlightFieldType", hf)
}

func Test_HighlightField_should_create_json_with_field(t *testing.T) {
	t.Parallel()
	// Given
	hf := es.HighlightField("title")

	// When Then
	assert.NotNil(t, hf)
	bodyJSON := assert.MarshalWithoutError(t, hf)
	assert.Equal(t, "{\"title\":{}}", bodyJSON)
}

////   HighlightField FragmentSize   ////

func Test_HighlightField_should_have_FragmentSize_method(t *testing.T) {
	t.Parallel()
	// Given
	hf := es.HighlightField("title")

	// When Then
	assert.NotNil(t, hf.FragmentSize)
}

func Test_HighlightField_FragmentSize_should_create_json_with_fragment_size_field(t *testing.T) {
	t.Parallel()
	// Given
	hf := es.HighlightField("content").FragmentSize(150)

	// When Then
	assert.NotNil(t, hf)
	bodyJSON := assert.MarshalWithoutError(t, hf)
	assert.Equal(t, "{\"content\":{\"fragment_size\":150}}", bodyJSON)
}

////   HighlightField NumberOfFragments   ////

func Test_HighlightField_should_have_NumberOfFragments_method(t *testing.T) {
	t.Parallel()
	// Given
	hf := es.HighlightField("title")

	// When Then
	assert.NotNil(t, hf.NumberOfFragments)
}

func Test_HighlightField_NumberOfFragments_should_create_json_with_number_of_fragments_field(t *testing.T) {
	t.Parallel()
	// Given
	hf := es.HighlightField("title").NumberOfFragments(0)

	// When Then
	assert.NotNil(t, hf)
	bodyJSON := assert.MarshalWithoutError(t, hf)
	assert.Equal(t, "{\"title\":{\"number_of_fragments\":0}}", bodyJSON)
}

////   HighlightField NoMatchSize   ////

func Test_HighlightField_should_have_NoMatchSize_method(t *testing.T) {
	t.Parallel()
	// Given
	hf := es.HighlightField("title")

	// When Then
	assert.NotNil(t, hf.NoMatchSize)
}

func Test_HighlightField_NoMatchSize_should_create_json_with_no_match_size_field(t *testing.T) {
	t.Parallel()
	// Given
	hf := es.HighlightField("content").NoMatchSize(150)

	// When Then
	assert.NotNil(t, hf)
	bodyJSON := assert.MarshalWithoutError(t, hf)
	assert.Equal(t, "{\"content\":{\"no_match_size\":150}}", bodyJSON)
}

////   HighlightField PreTags   ////

func Test_HighlightField_should_have_PreTags_method(t *testing.T) {
	t.Parallel()
	// Given
	hf := es.HighlightField("title")

	// When Then
	assert.NotNil(t, hf.PreTags)
}

func Test_HighlightField_PreTags_should_create_json_with_pre_tags_field(t *testing.T) {
	t.Parallel()
	// Given
	hf := es.HighlightField("title").PreTags("<b>")

	// When Then
	assert.NotNil(t, hf)
	bodyJSON := assert.MarshalWithoutError(t, hf)
	assert.Equal(t, "{\"title\":{\"pre_tags\":[\"\\u003cb\\u003e\"]}}", bodyJSON)
}

////   HighlightField PostTags   ////

func Test_HighlightField_should_have_PostTags_method(t *testing.T) {
	t.Parallel()
	// Given
	hf := es.HighlightField("title")

	// When Then
	assert.NotNil(t, hf.PostTags)
}

func Test_HighlightField_PostTags_should_create_json_with_post_tags_field(t *testing.T) {
	t.Parallel()
	// Given
	hf := es.HighlightField("title").PostTags("</b>")

	// When Then
	assert.NotNil(t, hf)
	bodyJSON := assert.MarshalWithoutError(t, hf)
	assert.Equal(t, "{\"title\":{\"post_tags\":[\"\\u003c/b\\u003e\"]}}", bodyJSON)
}

////   HighlightField Type   ////

func Test_HighlightField_should_have_Type_method(t *testing.T) {
	t.Parallel()
	// Given
	hf := es.HighlightField("title")

	// When Then
	assert.NotNil(t, hf.Type)
}

func Test_HighlightField_Type_should_create_json_with_type_field(t *testing.T) {
	t.Parallel()
	// Given
	hf := es.HighlightField("title").Type(HighlighterType.Plain)

	// When Then
	assert.NotNil(t, hf)
	bodyJSON := assert.MarshalWithoutError(t, hf)
	assert.Equal(t, "{\"title\":{\"type\":\"plain\"}}", bodyJSON)
}

////   HighlightField Order   ////

func Test_HighlightField_should_have_Order_method(t *testing.T) {
	t.Parallel()
	// Given
	hf := es.HighlightField("title")

	// When Then
	assert.NotNil(t, hf.Order)
}

func Test_HighlightField_Order_should_create_json_with_order_field(t *testing.T) {
	t.Parallel()
	// Given
	hf := es.HighlightField("content").Order("score")

	// When Then
	assert.NotNil(t, hf)
	bodyJSON := assert.MarshalWithoutError(t, hf)
	assert.Equal(t, "{\"content\":{\"order\":\"score\"}}", bodyJSON)
}

////   HighlightField HighlightQuery   ////

func Test_HighlightField_should_have_HighlightQuery_method(t *testing.T) {
	t.Parallel()
	// Given
	hf := es.HighlightField("title")

	// When Then
	assert.NotNil(t, hf.HighlightQuery)
}

func Test_HighlightField_HighlightQuery_should_not_set_highlight_query_when_query_is_nil(t *testing.T) {
	t.Parallel()
	// Given
	hf := es.HighlightField("content").HighlightQuery(nil)

	// When Then
	assert.NotNil(t, hf)
	bodyJSON := assert.MarshalWithoutError(t, hf)
	assert.Equal(t, "{\"content\":{}}", bodyJSON)
}

func Test_HighlightField_HighlightQuery_should_create_json_with_highlight_query_field(t *testing.T) {
	t.Parallel()
	// Given
	hf := es.HighlightField("content").HighlightQuery(es.Match("content", "elasticsearch"))

	// When Then
	assert.NotNil(t, hf)
	bodyJSON := assert.MarshalWithoutError(t, hf)
	assert.Equal(t, "{\"content\":{\"highlight_query\":{\"match\":{\"content\":{\"query\":\"elasticsearch\"}}}}}", bodyJSON)
}

////   HighlightField RequireFieldMatch   ////

func Test_HighlightField_should_have_RequireFieldMatch_method(t *testing.T) {
	t.Parallel()
	// Given
	hf := es.HighlightField("title")

	// When Then
	assert.NotNil(t, hf.RequireFieldMatch)
}

func Test_HighlightField_RequireFieldMatch_should_create_json_with_require_field_match_field(t *testing.T) {
	t.Parallel()
	// Given
	hf := es.HighlightField("content").RequireFieldMatch(false)

	// When Then
	assert.NotNil(t, hf)
	bodyJSON := assert.MarshalWithoutError(t, hf)
	assert.Equal(t, "{\"content\":{\"require_field_match\":false}}", bodyJSON)
}

////   HighlightField Fragmenter   ////

func Test_HighlightField_should_have_Fragmenter_method(t *testing.T) {
	t.Parallel()
	// Given
	hf := es.HighlightField("title")

	// When Then
	assert.NotNil(t, hf.Fragmenter)
}

func Test_HighlightField_Fragmenter_should_create_json_with_fragmenter_field(t *testing.T) {
	t.Parallel()
	// Given
	hf := es.HighlightField("content").Fragmenter(Fragmenter.Span)

	// When Then
	assert.NotNil(t, hf)
	bodyJSON := assert.MarshalWithoutError(t, hf)
	assert.Equal(t, "{\"content\":{\"fragmenter\":\"span\"}}", bodyJSON)
}

////   HighlightField MatchedFields   ////

func Test_HighlightField_should_have_MatchedFields_method(t *testing.T) {
	t.Parallel()
	// Given
	hf := es.HighlightField("title")

	// When Then
	assert.NotNil(t, hf.MatchedFields)
}

func Test_HighlightField_MatchedFields_should_create_json_with_matched_fields(t *testing.T) {
	t.Parallel()
	// Given
	hf := es.HighlightField("content").MatchedFields("content", "content.plain")

	// When Then
	assert.NotNil(t, hf)
	bodyJSON := assert.MarshalWithoutError(t, hf)
	assert.Equal(t, "{\"content\":{\"matched_fields\":[\"content\",\"content.plain\"]}}", bodyJSON)
}

////   Complex Highlight Query   ////

func Test_Highlight_should_create_complex_json_with_all_parameters(t *testing.T) {
	t.Parallel()
	// Given
	query := es.NewQuery(
		es.Match("content", "elasticsearch"),
	).
		Highlight(
			es.Highlight().
				PreTags("<em>").
				PostTags("</em>").
				Type(HighlighterType.Unified).
				Order("score").
				Encoder("html").
				RequireFieldMatch(true).
				FragmentSize(150).
				NumberOfFragments(5).
				BoundaryScanner(BoundaryScanner.Sentence).
				BoundaryScannerLocale("en-US").
				Field(
					es.HighlightField("title").NumberOfFragments(0),
					es.HighlightField("content").FragmentSize(200).NumberOfFragments(3),
				),
		)

	// When Then
	assert.NotNil(t, query)
	bodyJSON := assert.MarshalWithoutError(t, query)
	// nolint:golint,lll
	assert.Equal(t, "{\"highlight\":{\"boundary_scanner\":\"sentence\",\"boundary_scanner_locale\":\"en-US\",\"encoder\":\"html\",\"fields\":{\"content\":{\"fragment_size\":200,\"number_of_fragments\":3},\"title\":{\"number_of_fragments\":0}},\"fragment_size\":150,\"number_of_fragments\":5,\"order\":\"score\",\"post_tags\":[\"\\u003c/em\\u003e\"],\"pre_tags\":[\"\\u003cem\\u003e\"],\"require_field_match\":true,\"type\":\"unified\"},\"query\":{\"match\":{\"content\":{\"query\":\"elasticsearch\"}}}}", bodyJSON)
}

func Test_Highlight_with_field_specific_highlight_query(t *testing.T) {
	t.Parallel()
	// Given
	query := es.NewQuery(es.MatchAll()).
		Highlight(
			es.Highlight().
				Field(
					es.HighlightField("content").
						FragmentSize(150).
						NumberOfFragments(3).
						HighlightQuery(es.Match("content", "search")),
				),
		)

	// When Then
	assert.NotNil(t, query)
	bodyJSON := assert.MarshalWithoutError(t, query)
	// nolint:golint,lll
	assert.Equal(t, "{\"highlight\":{\"fields\":{\"content\":{\"fragment_size\":150,\"highlight_query\":{\"match\":{\"content\":{\"query\":\"search\"}}},\"number_of_fragments\":3}}},\"query\":{\"match_all\":{}}}", bodyJSON)
}
