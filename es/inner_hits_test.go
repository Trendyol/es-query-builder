package es_test

import (
	"github.com/Trendyol/es-query-builder/es"
	"github.com/Trendyol/es-query-builder/es/enums/sort/order"
	"github.com/Trendyol/es-query-builder/test/assert"
	"testing"
)

func Test_InnerHits_should_exist_on_es_package(t *testing.T) {
	// Given When Then
	assert.NotNil(t, es.InnerHits)
}

func Test_InnerHits_method_should_create_innerHitsType(t *testing.T) {
	// Given
	ih := es.InnerHits()

	// Then
	assert.NotNil(t, ih)
	assert.IsTypeString(t, "es.innerHitsType", ih)
}

func Test_FieldCollapse_should_exist_on_es_package(t *testing.T) {
	// Given When Then
	assert.NotNil(t, es.FieldCollapse)
}

func Test_FieldCollapse_method_should_create_fieldCollapseType(t *testing.T) {
	// Given
	ih := es.FieldCollapse("name")

	// Then
	assert.NotNil(t, ih)
	assert.IsTypeString(t, "es.fieldCollapseType", ih)
}

func Test_FieldAndFormat_should_exist_on_es_package(t *testing.T) {
	// Given When Then
	assert.NotNil(t, es.FieldAndFormat)
}

func Test_FieldAndFormat_method_should_create_fieldAndFormatType(t *testing.T) {
	// Given
	ih := es.FieldAndFormat("name")

	// Then
	assert.NotNil(t, ih)
	assert.IsTypeString(t, "es.fieldAndFormatType", ih)
}

func Test_InnerHits_should_have_Explain_method(t *testing.T) {
	// Given
	ih := es.InnerHits()

	// When Then
	assert.NotNil(t, ih.Explain)
}

func Test_InnerHits_Explain_should_create_json_with_explain_field_inside_inner_hits(t *testing.T) {
	// Given
	ih := es.InnerHits().Explain(true)
	// When Then
	assert.NotNil(t, ih)
	bodyJSON := assert.MarshalWithoutError(t, ih)
	assert.Equal(t, "{\"explain\":true}", bodyJSON)
}

func Test_InnerHits_should_have_Fields_method(t *testing.T) {
	// Given
	ih := es.InnerHits()

	// When Then
	assert.NotNil(t, ih.Fields)
}

func Test_InnerHits_Fields_should_create_json_with_fields_field_inside_inner_hits(t *testing.T) {
	// Given
	ih := es.InnerHits().Fields("f1", "g2", "h3")
	// When Then
	assert.NotNil(t, ih)
	bodyJSON := assert.MarshalWithoutError(t, ih)
	assert.Equal(t, "{\"fields\":[\"f1\",\"g2\",\"h3\"]}", bodyJSON)
}

func Test_InnerHits_should_have_Size_method(t *testing.T) {
	// Given
	ih := es.InnerHits()

	// When Then
	assert.NotNil(t, ih.Size)
}

func Test_InnerHits_Size_should_create_json_with_size_field_inside_inner_hits(t *testing.T) {
	// Given
	ih := es.InnerHits().Size(100)
	// When Then
	assert.NotNil(t, ih)
	bodyJSON := assert.MarshalWithoutError(t, ih)
	assert.Equal(t, "{\"size\":100}", bodyJSON)
}

func Test_InnerHits_should_have_From_method(t *testing.T) {
	// Given
	ih := es.InnerHits()

	// When Then
	assert.NotNil(t, ih.From)
}

func Test_InnerHits_From_should_create_json_with_from_field_inside_inner_hits(t *testing.T) {
	// Given
	ih := es.InnerHits().From(5_000)
	// When Then
	assert.NotNil(t, ih)
	bodyJSON := assert.MarshalWithoutError(t, ih)
	assert.Equal(t, "{\"from\":5000}", bodyJSON)
}

func Test_InnerHits_should_have_Name_method(t *testing.T) {
	// Given
	ih := es.InnerHits()

	// When Then
	assert.NotNil(t, ih.Name)
}

func Test_InnerHits_Name_should_create_json_with_name_field_inside_inner_hits(t *testing.T) {
	// Given
	ih := es.InnerHits().Name("Göksel")
	// When Then
	assert.NotNil(t, ih)
	bodyJSON := assert.MarshalWithoutError(t, ih)
	assert.Equal(t, "{\"name\":\"Göksel\"}", bodyJSON)
}

func Test_InnerHits_should_have_SourceFalse_method(t *testing.T) {
	// Given
	ih := es.InnerHits()

	// When Then
	assert.NotNil(t, ih.SourceFalse)
}

func Test_InnerHits_SourceFalse_should_create_json_with_source_field_inside_inner_hits_with_false(t *testing.T) {
	// Given
	ih := es.InnerHits().SourceFalse()
	// When Then
	assert.NotNil(t, ih)
	bodyJSON := assert.MarshalWithoutError(t, ih)
	assert.Equal(t, "{\"_source\":false}", bodyJSON)
}

func Test_InnerHits_should_have_SourceIncludes_method(t *testing.T) {
	// Given
	ih := es.InnerHits()

	// When Then
	assert.NotNil(t, ih.SourceIncludes)
}

func Test_InnerHits_SourceIncludes_should_create_json_with_source_field_inside_inner_hits(t *testing.T) {
	// Given
	ih := es.InnerHits().SourceIncludes("id", "name", "description")
	// When Then
	assert.NotNil(t, ih)
	bodyJSON := assert.MarshalWithoutError(t, ih)
	assert.Equal(t, "{\"_source\":{\"includes\":[\"id\",\"name\",\"description\"]}}", bodyJSON)
}

func Test_InnerHits_should_have_SourceExcludes_method(t *testing.T) {
	// Given
	ih := es.InnerHits()

	// When Then
	assert.NotNil(t, ih.SourceExcludes)
}

func Test_InnerHits_SourceExcludes_should_create_json_with_source_field_inside_inner_hits(t *testing.T) {
	// Given
	ih := es.InnerHits().SourceExcludes("secret", "key", "partition")
	// When Then
	assert.NotNil(t, ih)
	bodyJSON := assert.MarshalWithoutError(t, ih)
	assert.Equal(t, "{\"_source\":{\"excludes\":[\"secret\",\"key\",\"partition\"]}}", bodyJSON)
}

func Test_InnerHits_should_have_StoredFields_method(t *testing.T) {
	// Given
	ih := es.InnerHits()

	// When Then
	assert.NotNil(t, ih.StoredFields)
}

func Test_InnerHits_StoredFields_should_create_json_with_stored_fields_field_inside_inner_hits(t *testing.T) {
	// Given
	ih := es.InnerHits().StoredFields("x1", "y2", "z3")
	// When Then
	assert.NotNil(t, ih)
	bodyJSON := assert.MarshalWithoutError(t, ih)
	assert.Equal(t, "{\"stored_fields\":[\"x1\",\"y2\",\"z3\"]}", bodyJSON)
}

func Test_InnerHits_should_have_TrackScores_method(t *testing.T) {
	// Given
	ih := es.InnerHits()

	// When Then
	assert.NotNil(t, ih.TrackScores)
}

func Test_InnerHits_TrackScores_should_create_json_with_track_scores_field_inside_inner_hits(t *testing.T) {
	// Given
	ih := es.InnerHits().TrackScores(false)
	// When Then
	assert.NotNil(t, ih)
	bodyJSON := assert.MarshalWithoutError(t, ih)
	assert.Equal(t, "{\"track_scores\":false}", bodyJSON)
}

func Test_InnerHits_should_have_Version_method(t *testing.T) {
	// Given
	ih := es.InnerHits()

	// When Then
	assert.NotNil(t, ih.Version)
}

func Test_InnerHits_Version_should_create_json_with_version_field_inside_inner_hits(t *testing.T) {
	// Given
	ih := es.InnerHits().Version(true)
	// When Then
	assert.NotNil(t, ih)
	bodyJSON := assert.MarshalWithoutError(t, ih)
	assert.Equal(t, "{\"version\":true}", bodyJSON)
}

func Test_InnerHits_should_have_IgnoreUnmapped_method(t *testing.T) {
	// Given
	ih := es.InnerHits()

	// When Then
	assert.NotNil(t, ih.IgnoreUnmapped)
}

func Test_InnerHits_IgnoreUnmapped_should_create_json_with_ignore_unmapped_field_inside_inner_hits(t *testing.T) {
	// Given
	ih := es.InnerHits().IgnoreUnmapped(true)
	// When Then
	assert.NotNil(t, ih)
	bodyJSON := assert.MarshalWithoutError(t, ih)
	assert.Equal(t, "{\"ignore_unmapped\":true}", bodyJSON)
}

func Test_InnerHits_should_have_SeqNoPrimaryTerm_method(t *testing.T) {
	// Given
	ih := es.InnerHits()

	// When Then
	assert.NotNil(t, ih.SeqNoPrimaryTerm)
}

func Test_InnerHits_SeqNoPrimaryTerm_should_create_json_with_seq_no_primary_term_field_inside_inner_hits(t *testing.T) {
	// Given
	ih := es.InnerHits().SeqNoPrimaryTerm(false)
	// When Then
	assert.NotNil(t, ih)
	bodyJSON := assert.MarshalWithoutError(t, ih)
	assert.Equal(t, "{\"seq_no_primary_term\":false}", bodyJSON)
}

func Test_InnerHits_should_have_Sort_method(t *testing.T) {
	// Given
	ih := es.InnerHits()

	// When Then
	assert.NotNil(t, ih.Sort)
}

func Test_InnerHits_Sort_should_create_json_with_sort_field_inside_inner_hits(t *testing.T) {
	// Given
	ih := es.InnerHits().Sort(es.Sort("indexedAt").Order(order.Desc))
	// When Then
	assert.NotNil(t, ih)
	bodyJSON := assert.MarshalWithoutError(t, ih)
	assert.Equal(t, "{\"sort\":[{\"indexedAt\":{\"order\":\"desc\"}}]}", bodyJSON)
}
