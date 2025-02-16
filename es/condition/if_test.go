package condition_test

import (
	"testing"

	"github.com/Trendyol/es-query-builder/es"
	"github.com/Trendyol/es-query-builder/es/condition"
	"github.com/Trendyol/es-query-builder/test/assert"
)

func Test_Condition_If_should_add_Term_When_condition_is_true(t *testing.T) {
	t.Parallel()
	// Given
	cond := true

	// When
	query := es.NewQuery(
		es.Bool().
			Filter(
				condition.If(es.Term("language", "en"), cond),
				es.Exists("brandId"),
			),
	)

	// Then
	assert.NotNil(t, query)
	bodyJSON := assert.MarshalWithoutError(t, query)
	// nolint:golint,lll
	assert.Equal(t, "{\"query\":{\"bool\":{\"filter\":[{\"term\":{\"language\":{\"value\":\"en\"}}},{\"exists\":{\"field\":\"brandId\"}}]}}}", bodyJSON)
}

func Test_Condition_If_should_not_add_Term_When_condition_is_not_true(t *testing.T) {
	t.Parallel()
	// Given
	cond := false

	// When
	query := es.NewQuery(
		es.Bool().
			Filter(
				condition.If(es.Term("language", "en"), cond),
				es.Exists("brandId"),
			),
	)

	// Then
	assert.NotNil(t, query)
	bodyJSON := assert.MarshalWithoutError(t, query)
	assert.Equal(t, "{\"query\":{\"bool\":{\"filter\":[{\"exists\":{\"field\":\"brandId\"}}]}}}", bodyJSON)
}

func Test_Condition_If_should_work_with_slices(t *testing.T) {
	t.Parallel()
	// Given
	cond := true

	// When
	result := condition.If(es.Array{1, 2, 3, 4}, cond)

	// Then
	assert.NotNil(t, result)
	bodyJSON := assert.MarshalWithoutError(t, result)
	assert.Equal(t, "[1,2,3,4]", bodyJSON)
}
