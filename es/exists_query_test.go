package es_test

import (
	"testing"

	"github.com/Trendyol/es-query-builder/es"
	"github.com/Trendyol/es-query-builder/test/assert"
)

////   Exists   ////

func Test_Exists_should_exist_on_es_package(t *testing.T) {
	// Given When Then
	assert.NotNil(t, es.Exists)
}

func Test_Exists_should_create_json_with_exists_field_inside_query(t *testing.T) {
	// Given
	query := es.NewQuery(
		es.Exists("key"),
	)

	// When Then
	assert.NotNil(t, query)
	bodyJSON := assert.MarshalWithoutError(t, query)
	assert.Equal(t, "{\"query\":{\"exists\":{\"field\":\"key\"}}}", bodyJSON)
}

func Test_Exists_method_should_create_existsType(t *testing.T) {
	// Given
	b := es.Exists("key")

	// Then
	assert.NotNil(t, b)
	assert.IsTypeString(t, "es.existsType", b)
}

////   ExistsFunc   ////

func Test_ExistsFunc_should_exist_on_es_package(t *testing.T) {
	// Given When Then
	assert.NotNil(t, es.ExistsFunc)
}

func Test_ExistsFunc_should_create_json_with_exists_field_inside_query(t *testing.T) {
	// Given
	query := es.NewQuery(
		es.ExistsFunc("key", func(key string) bool {
			return true
		}),
	)

	// When Then
	assert.NotNil(t, query)
	bodyJSON := assert.MarshalWithoutError(t, query)
	assert.Equal(t, "{\"query\":{\"exists\":{\"field\":\"key\"}}}", bodyJSON)
}

func Test_ExistsFunc_should_not_add_exists_field_inside_query_when_callback_result_is_false(t *testing.T) {
	// Given
	query := es.NewQuery(
		es.ExistsFunc("key", func(key string) bool {
			return false
		}),
	)

	// When Then
	assert.NotNil(t, query)
	bodyJSON := assert.MarshalWithoutError(t, query)
	assert.Equal(t, "{\"query\":{}}", bodyJSON)
}

func Test_ExistsFunc_should_add_only_exists_fields_inside_the_query_when_callback_result_is_true(t *testing.T) {
	// Given
	query := es.NewQuery(
		es.Bool().
			Filter(
				es.ExistsFunc("a", func(key string) bool {
					return false
				}),
				es.ExistsFunc("c", func(key string) bool {
					return true
				}),
				es.ExistsFunc("e", func(key string) bool {
					return true
				}),
			),
	)

	// When Then
	assert.NotNil(t, query)
	bodyJSON := assert.MarshalWithoutError(t, query)
	assert.Equal(t, "{\"query\":{\"bool\":{\"filter\":[{\"exists\":{\"field\":\"c\"}},{\"exists\":{\"field\":\"e\"}}]}}}", bodyJSON)
}

func Test_ExistsFunc_method_should_create_existsType(t *testing.T) {
	// Given
	b := es.ExistsFunc("key", func(key string) bool {
		return true
	})

	// Then
	assert.NotNil(t, b)
	assert.IsTypeString(t, "es.existsType", b)
}

////   ExistsIF   ////

func Test_ExistsIf_should_exist_on_es_package(t *testing.T) {
	// Given When Then
	assert.NotNil(t, es.ExistsIf)
}

func Test_ExistsIf_should_create_json_with_exists_field_inside_query(t *testing.T) {
	// Given
	query := es.NewQuery(
		es.ExistsIf("key", true),
	)

	// When Then
	assert.NotNil(t, query)
	bodyJSON := assert.MarshalWithoutError(t, query)
	assert.Equal(t, "{\"query\":{\"exists\":{\"field\":\"key\"}}}", bodyJSON)
}

func Test_ExistsIf_should_not_add_exists_field_inside_query_when_condition_is_false(t *testing.T) {
	// Given
	query := es.NewQuery(
		es.ExistsIf("key", false),
	)

	// When Then
	assert.NotNil(t, query)
	bodyJSON := assert.MarshalWithoutError(t, query)
	assert.Equal(t, "{\"query\":{}}", bodyJSON)
}

func Test_ExistsIf_should_add_only_exists_fields_inside_the_query_when_condition_is_true(t *testing.T) {
	// Given
	query := es.NewQuery(
		es.Bool().
			Filter(
				es.ExistsIf("a", false),
				es.ExistsIf("c", true),
				es.ExistsIf("e", true),
			),
	)

	// When Then
	assert.NotNil(t, query)
	bodyJSON := assert.MarshalWithoutError(t, query)
	assert.Equal(t, "{\"query\":{\"bool\":{\"filter\":[{\"exists\":{\"field\":\"c\"}},{\"exists\":{\"field\":\"e\"}}]}}}", bodyJSON)
}

func Test_ExistsIf_method_should_create_existsType(t *testing.T) {
	// Given
	b := es.ExistsFunc("key", func(key string) bool {
		return true
	})

	// Then
	assert.NotNil(t, b)
	assert.IsTypeString(t, "es.existsType", b)
}
