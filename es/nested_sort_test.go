package es_test

import (
	"testing"

	"github.com/Trendyol/es-query-builder/es"
	"github.com/Trendyol/es-query-builder/test/assert"
)

func Test_NestedSort_should_exist_on_es_package(t *testing.T) {
	t.Parallel()
	// Given When Then
	assert.NotNil(t, es.NestedSort)
}

func Test_NestedSort_method_should_create_nestedSortType(t *testing.T) {
	t.Parallel()
	// Given
	n := es.NestedSort("timestamp")

	// Then
	assert.NotNil(t, n)
	assert.IsTypeString(t, "es.nestedSortType", n)
}

func Test_NestedSort_should_have_MaxChildren_method(t *testing.T) {
	t.Parallel()
	// Given
	nestedSort := es.NestedSort("timestamp")

	// When Then
	assert.NotNil(t, nestedSort.MaxChildren)
}

func Test_MaxChildren_should_add_max_children_field_into_NestedSort(t *testing.T) {
	t.Parallel()
	// Given
	nestedSort := es.NestedSort("timestamp").MaxChildren(42)

	// When Then
	assert.NotNil(t, nestedSort)
	bodyJSON := assert.MarshalWithoutError(t, nestedSort)
	assert.Equal(t, "{\"max_children\":42,\"path\":\"timestamp\"}", bodyJSON)
}

func Test_NestedSort_should_have_Filter_method(t *testing.T) {
	t.Parallel()
	// Given
	nestedSort := es.NestedSort("timestamp")

	// When Then
	assert.NotNil(t, nestedSort.Filter)
}

func Test_Filter_should_add_filter_field_into_NestedSort(t *testing.T) {
	t.Parallel()
	// Given
	nestedSort := es.NestedSort("timestamp").
		Filter(
			es.Bool().Must(
				es.Term("soldier", 76),
			),
		)

	// When Then
	assert.NotNil(t, nestedSort)
	bodyJSON := assert.MarshalWithoutError(t, nestedSort)
	assert.Equal(t, "{\"filter\":{\"bool\":{\"must\":[{\"term\":{\"soldier\":{\"value\":76}}}]}},\"path\":\"timestamp\"}", bodyJSON)
}

func Test_NestedSort_should_have_Nested_method(t *testing.T) {
	t.Parallel()
	// Given
	nestedSort := es.NestedSort("timestamp")

	// When Then
	assert.NotNil(t, nestedSort.Nested)
}

func Test_Nested_should_add_filter_field_into_NestedSort(t *testing.T) {
	t.Parallel()
	// Given
	nestedSort := es.NestedSort("timestamp").Nested(es.NestedSort("zone"))

	// When Then
	assert.NotNil(t, nestedSort)
	bodyJSON := assert.MarshalWithoutError(t, nestedSort)
	assert.Equal(t, "{\"nested\":{\"path\":\"zone\"},\"path\":\"timestamp\"}", bodyJSON)
}
