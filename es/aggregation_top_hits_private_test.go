package es

import (
	"testing"

	Order "github.com/Trendyol/es-query-builder/es/enums/sort/order"

	"github.com/Trendyol/es-query-builder/test/assert"
)

func Test_TopHitsAgg_Sort_should_return_unchanged_when_top_hits_is_not_object(t *testing.T) {
	t.Parallel()

	// Given
	agg := topHitsAggType{"top_hits": "invalid"}

	// When
	result := agg.Sort(Sort("date").Order(Order.Desc))

	// Then
	assert.NotNil(t, result)
	bodyJSON := assert.MarshalWithoutError(t, result)
	assert.Equal(t, "{\"top_hits\":\"invalid\"}", bodyJSON)
}

func Test_TopHitsAgg_SourceIncludes_should_return_unchanged_when_top_hits_is_not_object(t *testing.T) {
	t.Parallel()

	// Given
	agg := topHitsAggType{"top_hits": "invalid"}

	// When
	result := agg.SourceIncludes("title")

	// Then
	assert.NotNil(t, result)
	bodyJSON := assert.MarshalWithoutError(t, result)
	assert.Equal(t, "{\"top_hits\":\"invalid\"}", bodyJSON)
}

func Test_TopHitsAgg_SourceExcludes_should_return_unchanged_when_top_hits_is_not_object(t *testing.T) {
	t.Parallel()

	// Given
	agg := topHitsAggType{"top_hits": "invalid"}

	// When
	result := agg.SourceExcludes("description")

	// Then
	assert.NotNil(t, result)
	bodyJSON := assert.MarshalWithoutError(t, result)
	assert.Equal(t, "{\"top_hits\":\"invalid\"}", bodyJSON)
}
