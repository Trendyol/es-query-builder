package es_test

import (
	"testing"

	"github.com/Trendyol/es-query-builder/es"
	"github.com/Trendyol/es-query-builder/test/assert"
)

////   Terms   ////

func Test_Terms_should_exist_on_es_package(t *testing.T) {
	// Given When Then
	assert.NotNil(t, es.Terms)
}

func Test_Terms_should_create_json_with_terms_field_inside_query(t *testing.T) {
	// Given
	query := es.NewQuery(
		es.Terms("key", "value1", "value2", "value3"),
	)

	// When Then
	assert.NotNil(t, query)
	bodyJSON := assert.MarshalWithoutError(t, query)
	assert.Equal(t, "{\"query\":{\"terms\":{\"key\":[\"value1\",\"value2\",\"value3\"]}}}", bodyJSON)
}

func Test_Terms_method_should_create_termsType(t *testing.T) {
	// Given
	b := es.Terms("key", "value1", "value2", "value3")

	// Then
	assert.NotNil(t, b)
	assert.IsTypeString(t, "es.termsType", b)
}

////   TermsArray   ////

func Test_TermsArray_should_exist_on_es_package(t *testing.T) {
	// Given When Then
	assert.NotNil(t, es.TermsArray[string])
}

func Test_TermsArray_should_create_json_with_terms_field_inside_query(t *testing.T) {
	// Given
	query := es.NewQuery(
		es.TermsArray("key", []any{"value1", "value2", "value3"}),
	)

	// When Then
	assert.NotNil(t, query)
	bodyJSON := assert.MarshalWithoutError(t, query)
	assert.Equal(t, "{\"query\":{\"terms\":{\"key\":[\"value1\",\"value2\",\"value3\"]}}}", bodyJSON)
}

func Test_TermsArray_method_should_create_termsType(t *testing.T) {
	// Given
	b := es.TermsArray("key", []any{"value1", "value2", "value3"})

	// Then
	assert.NotNil(t, b)
	assert.IsTypeString(t, "es.termsType", b)
}

////   TermsFunc   ////

func Test_TermsFunc_should_exist_on_es_package(t *testing.T) {
	// Given When Then
	assert.NotNil(t, es.TermsFunc[string])
}

func Test_TermsFunc_should_create_json_with_terms_field_inside_query(t *testing.T) {
	// Given
	query := es.NewQuery(
		es.TermsFunc("key", []string{"a", "b", "c"}, func(key string, values []string) bool {
			return true
		}),
	)

	// When Then
	assert.NotNil(t, query)
	bodyJSON := assert.MarshalWithoutError(t, query)
	assert.Equal(t, "{\"query\":{\"terms\":{\"key\":[\"a\",\"b\",\"c\"]}}}", bodyJSON)
}

func Test_TermsFunc_should_not_add_terms_field_inside_query_when_callback_result_is_false(t *testing.T) {
	// Given
	query := es.NewQuery(
		es.TermsFunc("key", []string{"a", "b", "c"}, func(key string, value []string) bool {
			return false
		}),
	)

	// When Then
	assert.NotNil(t, query)
	bodyJSON := assert.MarshalWithoutError(t, query)
	assert.Equal(t, "{\"query\":{}}", bodyJSON)
}

func Test_TermsFunc_should_add_only_terms_fields_inside_the_query_when_callback_result_is_true(t *testing.T) {
	// Given
	query := es.NewQuery(
		es.Bool().
			Filter(
				es.TermsFunc("a", []string{"10", "11", "12"}, func(key string, value []string) bool {
					return false
				}),
				es.TermsFunc("c", []string{"20", "21", "22"}, func(key string, value []string) bool {
					return false
				}),
				es.TermsFunc("e", []string{"30", "31", "32"}, func(key string, value []string) bool {
					return true
				}),
			),
	)

	// When Then
	assert.NotNil(t, query)
	bodyJSON := assert.MarshalWithoutError(t, query)
	assert.Equal(t, "{\"query\":{\"bool\":{\"filter\":[{\"terms\":{\"e\":[\"30\",\"31\",\"32\"]}}]}}}", bodyJSON)
}

func Test_TermsFunc_method_should_create_termType(t *testing.T) {
	// Given
	b := es.TermsFunc("key", []string{"a", "b", "c"}, func(key string, value []string) bool {
		return true
	})

	// Then
	assert.NotNil(t, b)
	assert.IsTypeString(t, "es.termsType", b)
}

////   TermsIf   ////

func Test_TermsIf_should_exist_on_es_package(t *testing.T) {
	// Given When Then
	assert.NotNil(t, es.TermsIf[string])
}

func Test_TermsIf_should_create_json_with_terms_field_inside_query(t *testing.T) {
	// Given
	query := es.NewQuery(
		es.TermsIf("key", []string{"a", "b", "c"}, true),
	)

	// When Then
	assert.NotNil(t, query)
	bodyJSON := assert.MarshalWithoutError(t, query)
	assert.Equal(t, "{\"query\":{\"terms\":{\"key\":[\"a\",\"b\",\"c\"]}}}", bodyJSON)
}

func Test_TermsIf_should_not_add_terms_field_inside_query_when_condition_is_false(t *testing.T) {
	// Given
	query := es.NewQuery(
		es.TermsIf("key", []string{"a", "b", "c"}, false),
	)

	// When Then
	assert.NotNil(t, query)
	bodyJSON := assert.MarshalWithoutError(t, query)
	assert.Equal(t, "{\"query\":{}}", bodyJSON)
}

func Test_TermsIf_should_add_only_terms_fields_inside_the_query_when_condition_is_true(t *testing.T) {
	// Given
	query := es.NewQuery(
		es.Bool().
			Filter(
				es.TermsIf("a", []string{"10", "11", "12"}, false),
				es.TermsIf("c", []string{"20", "21", "22"}, false),
				es.TermsIf("e", []string{"30", "31", "32"}, true),
			),
	)

	// When Then
	assert.NotNil(t, query)
	bodyJSON := assert.MarshalWithoutError(t, query)
	assert.Equal(t, "{\"query\":{\"bool\":{\"filter\":[{\"terms\":{\"e\":[\"30\",\"31\",\"32\"]}}]}}}", bodyJSON)
}

func Test_TermsIf_method_should_create_termType(t *testing.T) {
	// Given
	b := es.TermsIf("key", []string{"a", "b", "c"}, true)

	// Then
	assert.NotNil(t, b)
	assert.IsTypeString(t, "es.termsType", b)
}
