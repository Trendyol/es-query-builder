package es

import (
	"testing"

	"github.com/Trendyol/es-query-builder/test/assert"
)

func Test_getObjectFromAggs_should_get_objects_from_aggs_when_exists(t *testing.T) {
	t.Parallel()
	// Given
	aggsQuery := AvgAgg("category").Meta("hello", "world")

	// When
	meta, ok := getObjectFromAggs(aggsQuery, "avg", "meta")

	// Then
	assert.True(t, ok)
	assert.NotNil(t, meta)
	bodyJSON := assert.MarshalWithoutError(t, meta)
	assert.Equal(t, "{\"hello\":\"world\"}", bodyJSON)
}

func Test_getObjectFromAggs_should_return_false_when_exists(t *testing.T) {
	t.Parallel()
	// Given
	aggsQuery := AvgAgg("category").Meta("hello", "world")

	// When
	_, ok := getObjectFromAggs(aggsQuery, "avg", "fooBar")

	// Then
	assert.False(t, ok)
}

func Test_getObjectFromAggs_should_return_false_when_invalid_aggs_type(t *testing.T) {
	t.Parallel()
	// Given
	aggsQuery := AvgAgg("category").Meta("hello", "world")

	// When
	_, ok := getObjectFromAggs(aggsQuery, "n/A", "meta")

	// Then
	assert.False(t, ok)
}
