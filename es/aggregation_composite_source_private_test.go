package es

import (
	"testing"

	"github.com/Trendyol/es-query-builder/test/assert"
)

func Test_putInSource_should_return_unchanged_when_source_is_empty(t *testing.T) {
	t.Parallel()

	// Given
	source := compositeSourceType{}

	// When
	result := source.putInSource("order", "asc")

	// Then
	assert.NotNil(t, result)
	bodyJSON := assert.MarshalWithoutError(t, result)
	assert.Equal(t, "{}", bodyJSON)
}

func Test_putInSource_should_return_unchanged_when_value_is_not_object(t *testing.T) {
	t.Parallel()

	// Given
	source := compositeSourceType{"brand": "not-object"}

	// When
	result := source.putInSource("order", "asc")

	// Then
	assert.NotNil(t, result)
	bodyJSON := assert.MarshalWithoutError(t, result)
	assert.Equal(t, "{\"brand\":\"not-object\"}", bodyJSON)
}

func Test_putInSource_should_return_unchanged_when_nested_value_is_not_object(t *testing.T) {
	t.Parallel()

	// Given
	source := compositeSourceType{"brand": Object{"terms": "bad"}}

	// When
	result := source.putInSource("order", "asc")

	// Then
	assert.NotNil(t, result)
	bodyJSON := assert.MarshalWithoutError(t, result)
	assert.Equal(t, "{\"brand\":{\"terms\":\"bad\"}}", bodyJSON)
}
