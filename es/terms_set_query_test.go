package es_test

import (
	ScriptLanguage "github.com/Trendyol/es-query-builder/es/enums/script-language"
	"testing"

	"github.com/Trendyol/es-query-builder/es"
	"github.com/Trendyol/es-query-builder/test/assert"
)

////   Terms Set   ////

func Test_TermsSet_should_exist_on_es_package(t *testing.T) {
	t.Parallel()
	// Given When Then
	assert.NotNil(t, es.TermsSet)
}

func Test_TermsSet_should_create_json_with_terms_set_field_inside_query(t *testing.T) {
	t.Parallel()
	// Given
	query := es.NewQuery(
		es.TermsSet("key", "value1", "value2", "value3"),
	)

	// When Then
	assert.NotNil(t, query)
	bodyJSON := assert.MarshalWithoutError(t, query)
	assert.Equal(t, "{\"query\":{\"terms_set\":{\"key\":{\"terms\":[\"value1\",\"value2\",\"value3\"]}}}}", bodyJSON)
}

func Test_TermsSet_method_should_create_termsSetType(t *testing.T) {
	t.Parallel()
	// Given
	b := es.TermsSet("key", "value1", "value2", "value3")

	// Then
	assert.NotNil(t, b)
	assert.IsTypeString(t, "es.termsSetType", b)
}

func Test_TermsSet_should_have_Boost_method(t *testing.T) {
	t.Parallel()
	// Given
	termsSet := es.TermsSet("key", "value1", "value2", "value3")

	// When Then
	assert.NotNil(t, termsSet.Boost)
}

func Test_TermsSet_Boost_should_create_json_with_boost_field_inside_terms_set(t *testing.T) {
	t.Parallel()
	// Given
	query := es.NewQuery(
		es.TermsSet("sector.name", "a1", "b2", "c3").
			Boost(2.718),
	)

	// When Then
	assert.NotNil(t, query)
	bodyJSON := assert.MarshalWithoutError(t, query)
	assert.Equal(t, "{\"query\":{\"terms_set\":{\"sector.name\":{\"boost\":2.718,\"terms\":[\"a1\",\"b2\",\"c3\"]}}}}", bodyJSON)
}

func Test_TermsSet_should_have_MinimumShouldMatchField_method(t *testing.T) {
	t.Parallel()
	// Given
	termsSet := es.TermsSet("key", "value1", "value2", "value3")

	// When Then
	assert.NotNil(t, termsSet.MinimumShouldMatchField)
}

func Test_TermsSet_MinimumShouldMatchField_should_create_json_with_minimum_should_match_field_field_inside_terms_set(t *testing.T) {
	t.Parallel()
	// Given
	query := es.NewQuery(
		es.TermsSet("sector.name", "a1", "b2", "c3").
			MinimumShouldMatchField("match_threshold"),
	)

	// When Then
	assert.NotNil(t, query)
	bodyJSON := assert.MarshalWithoutError(t, query)
	assert.Equal(t, "{\"query\":{\"terms_set\":{\"sector.name\":{\"minimum_should_match_field\":\"match_threshold\",\"terms\":[\"a1\",\"b2\",\"c3\"]}}}}", bodyJSON)
}

func Test_TermsSet_should_have_MinimumShouldMatchScript_method(t *testing.T) {
	t.Parallel()
	// Given
	termsSet := es.TermsSet("key", "value1", "value2", "value3")

	// When Then
	assert.NotNil(t, termsSet.MinimumShouldMatchScript)
}

func Test_TermsSet_MinimumShouldMatchScript_should_create_json_with_minimum_should_match_script_field_inside_terms_set(t *testing.T) {
	t.Parallel()
	// Given
	query := es.NewQuery(
		es.TermsSet("sector.name", "a1", "b2", "c3").
			MinimumShouldMatchScript(
				es.ScriptSource("Math.max(1, doc['match_threshold'].value - 1)", ScriptLanguage.Painless).
					Option("timeout", "10s").
					Option("retry", "5").
					Option("size", "500").
					Parameter("threshold", 2).
					Parameter("items", []int{1, 2, 3, 4}),
			),
	)

	// When Then
	assert.NotNil(t, query)
	bodyJSON := assert.MarshalWithoutError(t, query)
	// nolint:golint,lll
	assert.Equal(t, "{\"query\":{\"terms_set\":{\"sector.name\":{\"minimum_should_match_script\":{\"lang\":\"painless\",\"options\":{\"retry\":\"5\",\"size\":\"500\",\"timeout\":\"10s\"},\"params\":{\"items\":[1,2,3,4],\"threshold\":2},\"source\":\"Math.max(1, doc['match_threshold'].value - 1)\"},\"terms\":[\"a1\",\"b2\",\"c3\"]}}}}", bodyJSON)
}
