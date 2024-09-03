package es_test

import (
	"testing"

	"github.com/Trendyol/es-query-builder/es"
	"github.com/Trendyol/es-query-builder/test/assert"

	Order "github.com/Trendyol/es-query-builder/es/enums/sort/order"
)

////    AGGS    ////

func Test_AggTerm_should_exist_on_es_package(t *testing.T) {
	// Given When Then
	assert.NotNil(t, es.AggTerm)
}

func Test_AggTerm_method_should_create_aggTermType(t *testing.T) {
	// Given
	a := es.AggTerm("path")

	// Then
	assert.NotNil(t, a)
	assert.IsTypeString(t, "es.aggTermType", a)
}

func Test_AggTerm_should_create_json_with_field_field_inside(t *testing.T) {
	// Given
	a := es.AggTerm("path")

	// When Then
	assert.NotNil(t, a)
	bodyJSON := assert.MarshalWithoutError(t, a)
	assert.Equal(t, "{\"field\":\"path\"}", bodyJSON)
}

func Test_AggTerm_should_have_Missing_method(t *testing.T) {
	// Given
	a := es.AggTerm("path")

	// When Then
	assert.NotNil(t, a.Missing)
}

func Test_Missing_should_add_missing_field_into_AggTerm(t *testing.T) {
	// Given
	a := es.AggTerm("path").Missing("missing_name")

	// When Then
	assert.NotNil(t, a)
	bodyJSON := assert.MarshalWithoutError(t, a)
	assert.Equal(t, "{\"field\":\"path\",\"missing\":\"missing_name\"}", bodyJSON)
}

func Test_AggTerms_should_exist_on_es_package(t *testing.T) {
	// Given When Then
	assert.NotNil(t, es.AggTerms)
}

func Test_AggTerms_method_should_create_aggType(t *testing.T) {
	// Given
	a := es.AggTerms()

	// Then
	assert.NotNil(t, a)
	assert.IsTypeString(t, "es.aggsType", a)
}

func Test_AggTerms_should_create_json_with_terms_field_inside(t *testing.T) {
	// Given
	a := es.AggTerms()

	// When Then
	assert.NotNil(t, a)
	bodyJSON := assert.MarshalWithoutError(t, a)
	assert.Equal(t, "{\"terms\":{}}", bodyJSON)
}

func Test_AggMultiTerms_should_exist_on_es_package(t *testing.T) {
	// Given When Then
	assert.NotNil(t, es.AggMultiTerms)
}

func Test_AggMultiTerms_method_should_create_aggType(t *testing.T) {
	// Given
	a := es.AggMultiTerms()

	// Then
	assert.NotNil(t, a)
	assert.IsTypeString(t, "es.aggsType", a)
}

func Test_AggMultiTerms_should_create_json_with_multi_terms_field_inside(t *testing.T) {
	// Given
	a := es.AggMultiTerms()

	// When Then
	assert.NotNil(t, a)
	bodyJSON := assert.MarshalWithoutError(t, a)
	assert.Equal(t, "{\"multi_terms\":{}}", bodyJSON)
}

func Test_AggNested_should_exist_on_es_package(t *testing.T) {
	// Given When Then
	assert.NotNil(t, es.AggNested)
}

func Test_AggNested_method_should_create_aggType(t *testing.T) {
	// Given
	a := es.AggNested()

	// Then
	assert.NotNil(t, a)
	assert.IsTypeString(t, "es.aggsType", a)
}

func Test_AggNested_should_create_json_with_nested_field_inside(t *testing.T) {
	// Given
	a := es.AggNested()

	// When Then
	assert.NotNil(t, a)
	bodyJSON := assert.MarshalWithoutError(t, a)
	assert.Equal(t, "{\"nested\":{}}", bodyJSON)
}

func Test_AggCustom_should_exist_on_es_package(t *testing.T) {
	// Given When Then
	assert.NotNil(t, es.AggCustom)
}

func Test_AggCustom_method_should_create_aggType(t *testing.T) {
	// Given
	a := es.AggCustom(nil)

	// Then
	assert.NotNil(t, a)
	assert.IsTypeString(t, "es.aggsType", a)
}

func Test_AggCustom_should_create_json(t *testing.T) {
	// Given
	a := es.AggCustom(es.Object{
		"custom": es.Object{
			"my_field": es.Array{1, 2, 3},
		},
	})

	// When Then
	assert.NotNil(t, a)
	bodyJSON := assert.MarshalWithoutError(t, a)
	assert.Equal(t, "{\"custom\":{\"my_field\":[1,2,3]}}", bodyJSON)
}

func Test_AggType_should_have_Field_method(t *testing.T) {
	// Given
	a := es.AggTerms()

	// When Then
	assert.NotNil(t, a.Field)
}

func Test_Field_should_add_field_field_into_AggTerm(t *testing.T) {
	// Given
	a := es.AggTerms().Field("path")

	// When Then
	assert.NotNil(t, a)
	bodyJSON := assert.MarshalWithoutError(t, a)
	assert.Equal(t, "{\"terms\":{\"field\":\"path\"}}", bodyJSON)
}

func Test_AggType_should_have_Path_method(t *testing.T) {
	// Given
	a := es.AggTerms()

	// When Then
	assert.NotNil(t, a.Path)
}

func Test_Path_should_add_path_field_into_AggNested(t *testing.T) {
	// Given
	a := es.AggNested().Path("review")

	// When Then
	assert.NotNil(t, a)
	bodyJSON := assert.MarshalWithoutError(t, a)
	assert.Equal(t, "{\"nested\":{\"path\":\"review\"}}", bodyJSON)
}

func Test_AggType_should_have_Size_method(t *testing.T) {
	// Given
	a := es.AggTerms()

	// When Then
	assert.NotNil(t, a.Size)
}

func Test_Size_should_add_size_field_into_AggType(t *testing.T) {
	// Given
	a := es.AggTerms().Size(333)

	// When Then
	assert.NotNil(t, a)
	bodyJSON := assert.MarshalWithoutError(t, a)
	assert.Equal(t, "{\"terms\":{\"size\":333}}", bodyJSON)
}

func Test_AggType_should_have_Order_method(t *testing.T) {
	// Given
	a := es.AggTerms()

	// When Then
	assert.NotNil(t, a.Field)
}

func Test_Order_should_add_order_field_into_AggType(t *testing.T) {
	// Given
	a := es.AggTerms().Order("path", Order.Desc)

	// When Then
	assert.NotNil(t, a)
	bodyJSON := assert.MarshalWithoutError(t, a)
	assert.Equal(t, "{\"terms\":{\"order\":{\"path\":\"desc\"}}}", bodyJSON)
}

func Test_AggType_should_have_Include_method(t *testing.T) {
	// Given
	a := es.AggTerms()

	// When Then
	assert.NotNil(t, a.Include)
}

func Test_Include_should_add_include_field_into_AggType(t *testing.T) {
	// Given
	a := es.AggTerms().Include("*.2024")

	// When Then
	assert.NotNil(t, a)
	bodyJSON := assert.MarshalWithoutError(t, a)
	assert.Equal(t, "{\"terms\":{\"include\":\"*.2024\"}}", bodyJSON)
}

func Test_AggType_should_have_Exclude_method(t *testing.T) {
	// Given
	a := es.AggTerms()

	// When Then
	assert.NotNil(t, a.Exclude)
}

func Test_Exclude_should_add_exclude_field_into_AggType(t *testing.T) {
	// Given
	a := es.AggTerms().Exclude("*.2021")

	// When Then
	assert.NotNil(t, a)
	bodyJSON := assert.MarshalWithoutError(t, a)
	assert.Equal(t, "{\"terms\":{\"exclude\":\"*.2021\"}}", bodyJSON)
}

func Test_AggType_should_have_Terms_method(t *testing.T) {
	// Given
	a := es.AggTerms()

	// When Then
	assert.NotNil(t, a.Terms)
}

func Test_AggMax_should_exist_on_es_package(t *testing.T) {
	// Given When Then
	assert.NotNil(t, es.AggMax)
}

func Test_AggMax_method_should_create_aggType(t *testing.T) {
	// Given
	a := es.AggMax()

	// Then
	assert.NotNil(t, a)
	assert.IsTypeString(t, "es.aggsType", a)
}

func Test_AggMax_should_create_json_with_terms_field_inside(t *testing.T) {
	// Given
	a := es.AggMax()

	// When Then
	assert.NotNil(t, a)
	bodyJSON := assert.MarshalWithoutError(t, a)
	assert.Equal(t, "{\"max\":{}}", bodyJSON)
}

////    Agg.Min    ////

func Test_AggMin_should_exist_on_es_package(t *testing.T) {
	// Given When Then
	assert.NotNil(t, es.AggMin)
}

func Test_AggMin_method_should_create_aggType(t *testing.T) {
	// Given
	a := es.AggMin()

	// Then
	assert.NotNil(t, a)
	assert.IsTypeString(t, "es.aggsType", a)
}

func Test_AggMin_should_create_json_with_terms_field_inside(t *testing.T) {
	// Given
	a := es.AggMin()

	// When Then
	assert.NotNil(t, a)
	bodyJSON := assert.MarshalWithoutError(t, a)
	assert.Equal(t, "{\"min\":{}}", bodyJSON)
}

////    Agg.Avg    ////

func Test_AggAvg_should_exist_on_es_package(t *testing.T) {
	// Given When Then
	assert.NotNil(t, es.AggAvg)
}

func Test_AggAvg_method_should_create_aggType(t *testing.T) {
	// Given
	a := es.AggAvg()

	// Then
	assert.NotNil(t, a)
	assert.IsTypeString(t, "es.aggsType", a)
}

func Test_AggAvg_should_create_json_with_terms_field_inside(t *testing.T) {
	// Given
	a := es.AggAvg()

	// When Then
	assert.NotNil(t, a)
	bodyJSON := assert.MarshalWithoutError(t, a)
	assert.Equal(t, "{\"avg\":{}}", bodyJSON)
}

func Test_Terms_should_add_terms_field_into_AggType(t *testing.T) {
	// Given
	a := es.AggMultiTerms().
		Terms(
			es.AggTerm("A1"),
			es.AggTerm("B2").
				Missing("Hell Divers"),
			es.AggTerm("C3"),
			es.AggTerm("D4"),
		)

	// When Then
	assert.NotNil(t, a)
	bodyJSON := assert.MarshalWithoutError(t, a)
	// nolint:golint,lll
	assert.Equal(t, "{\"multi_terms\":{\"terms\":[{\"field\":\"A1\"},{\"field\":\"B2\",\"missing\":\"Hell Divers\"},{\"field\":\"C3\"},{\"field\":\"D4\"}]}}", bodyJSON)
}

func Test_AggType_should_have_Aggs_method(t *testing.T) {
	// Given
	a := es.AggTerms()

	// When Then
	assert.NotNil(t, a.Aggs)
}

func Test_Aggs_should_add_aggs_field_into_AggType(t *testing.T) {
	// Given
	a := es.AggTerms().
		Field("path").
		Size(1_000).
		Order("_key", Order.Asc).
		Include("reduces").
		Aggs("test",
			es.AggMultiTerms().
				Terms(
					es.AggTerm("A1").
						Missing("a1"),
					es.AggTerm("B2"),
				),
		)

	// When Then
	assert.NotNil(t, a)
	bodyJSON := assert.MarshalWithoutError(t, a)
	// nolint:golint,lll
	assert.Equal(t, "{\"aggs\":{\"test\":{\"multi_terms\":{\"terms\":[{\"field\":\"A1\",\"missing\":\"a1\"},{\"field\":\"B2\"}]}}},\"terms\":{\"field\":\"path\",\"include\":\"reduces\",\"order\":{\"_key\":\"asc\"},\"size\":1000}}", bodyJSON)
}

func Test_Aggs_should_create_json_with_aggs_field_inside_query(t *testing.T) {
	// Given
	query := es.NewQuery(nil)
	query.Aggs("types",
		es.AggTerms().
			Field("type").
			Size(100),
	)

	// When Then
	assert.NotNil(t, query)
	assert.NotNil(t, query.Aggs)
	bodyJSON := assert.MarshalWithoutError(t, query)
	assert.Equal(t, "{\"aggs\":{\"types\":{\"terms\":{\"field\":\"type\",\"size\":100}}},\"query\":{}}", bodyJSON)
}

func Test_should_create_json_with_multiple_Aggs_inside_single_query(t *testing.T) {
	// Given
	query := es.NewQuery(nil).
		Aggs("types",
			es.AggTerms().
				Field("type").
				Size(100),
		).
		Aggs("average_review_score",
			es.AggAvg().
				Field("reviews.score"),
		)

	// When Then
	assert.NotNil(t, query)
	assert.NotNil(t, query.Aggs)
	bodyJSON := assert.MarshalWithoutError(t, query)
	// nolint:golint,lll
	assert.Equal(t, "{\"aggs\":{\"average_review_score\":{\"avg\":{\"field\":\"reviews.score\"}},\"types\":{\"terms\":{\"field\":\"type\",\"size\":100}}},\"query\":{}}", bodyJSON)
}
