package condition_test

import (
	"testing"

	"github.com/Trendyol/es-query-builder/es"
	"github.com/Trendyol/es-query-builder/es/condition"
	"github.com/Trendyol/es-query-builder/test/assert"
)

func Test_IfElse_should_return_item_when_condition_is_true(t *testing.T) {
	t.Parallel()
	// Given
	x := 15

	// When
	query := es.NewQuery(
		es.Bool().
			Filter(
				condition.IfElse(x > 10, es.Term("foo", "bar")),
				es.Exists("brandId"),
			),
	)

	// Then
	assert.NotNil(t, query)
	bodyJSON := assert.MarshalWithoutError(t, query)
	// nolint:golint,lll
	assert.Equal(t, "{\"query\":{\"bool\":{\"filter\":[{\"term\":{\"foo\":{\"value\":\"bar\"}}},{\"exists\":{\"field\":\"brandId\"}}]}}}", bodyJSON)
}

func Test_IfElse_should_return_nil_when_condition_is_false(t *testing.T) {
	t.Parallel()
	// Given
	x := 5

	// When
	query := es.NewQuery(
		es.Bool().
			Filter(
				condition.IfElse(x > 10, es.Term("foo", "bar")),
				es.Exists("brandId"),
			),
	)

	// Then
	assert.NotNil(t, query)
	bodyJSON := assert.MarshalWithoutError(t, query)
	assert.Equal(t, "{\"query\":{\"bool\":{\"filter\":[{\"exists\":{\"field\":\"brandId\"}}]}}}", bodyJSON)
}

func Test_IfElse_ElseIf_should_return_elseIf_item_when_first_is_false_and_second_is_true(t *testing.T) {
	t.Parallel()
	// Given
	x := 5
	y := 10

	// When
	query := es.NewQuery(
		es.Bool().
			Filter(
				condition.IfElse(x > 10, es.Term("foo", "bar"),
					condition.ElseIf(y < 20, es.Exists("fizz")),
				),
			),
	)

	// Then
	assert.NotNil(t, query)
	bodyJSON := assert.MarshalWithoutError(t, query)
	assert.Equal(t, "{\"query\":{\"bool\":{\"filter\":[{\"exists\":{\"field\":\"fizz\"}}]}}}", bodyJSON)
}

func Test_IfElse_ElseIf_should_return_first_item_when_both_conditions_are_true(t *testing.T) {
	t.Parallel()
	// Given
	x := 15
	y := 10

	// When
	query := es.NewQuery(
		es.Bool().
			Filter(
				condition.IfElse(x > 10, es.Term("foo", "bar"),
					condition.ElseIf(y < 20, es.Exists("fizz")),
				),
			),
	)

	// Then
	assert.NotNil(t, query)
	bodyJSON := assert.MarshalWithoutError(t, query)
	assert.Equal(t, "{\"query\":{\"bool\":{\"filter\":[{\"term\":{\"foo\":{\"value\":\"bar\"}}}]}}}", bodyJSON)
}

func Test_IfElse_ElseIf_should_return_nil_when_all_conditions_are_false(t *testing.T) {
	t.Parallel()
	// Given
	x := 5
	y := 25

	// When
	query := es.NewQuery(
		es.Bool().
			Filter(
				condition.IfElse(x > 10, es.Term("foo", "bar"),
					condition.ElseIf(y < 20, es.Exists("fizz")),
				),
				es.Exists("brandId"),
			),
	)

	// Then
	assert.NotNil(t, query)
	bodyJSON := assert.MarshalWithoutError(t, query)
	assert.Equal(t, "{\"query\":{\"bool\":{\"filter\":[{\"exists\":{\"field\":\"brandId\"}}]}}}", bodyJSON)
}

func Test_IfElse_Else_should_return_default_when_all_conditions_are_false(t *testing.T) {
	t.Parallel()
	// Given
	x := 5
	y := 25

	// When
	query := es.NewQuery(
		es.Bool().
			Filter(
				condition.IfElse(x > 10, es.Term("foo", "bar"),
					condition.ElseIf(y < 20, es.Exists("fizz")),
					condition.Else(es.Term("default", "value")),
				),
			),
	)

	// Then
	assert.NotNil(t, query)
	bodyJSON := assert.MarshalWithoutError(t, query)
	assert.Equal(t, "{\"query\":{\"bool\":{\"filter\":[{\"term\":{\"default\":{\"value\":\"value\"}}}]}}}", bodyJSON)
}

func Test_IfElse_Else_should_return_first_item_when_condition_is_true(t *testing.T) {
	t.Parallel()
	// Given
	x := 15

	// When
	query := es.NewQuery(
		es.Bool().
			Filter(
				condition.IfElse(x > 10, es.Term("foo", "bar"),
					condition.Else(es.Term("default", "value")),
				),
			),
	)

	// Then
	assert.NotNil(t, query)
	bodyJSON := assert.MarshalWithoutError(t, query)
	assert.Equal(t, "{\"query\":{\"bool\":{\"filter\":[{\"term\":{\"foo\":{\"value\":\"bar\"}}}]}}}", bodyJSON)
}

func Test_IfElse_Else_should_return_elseIf_item_when_first_false_second_true(t *testing.T) {
	t.Parallel()
	// Given
	x := 5
	y := 10

	// When
	query := es.NewQuery(
		es.Bool().
			Filter(
				condition.IfElse(x > 10, es.Term("foo", "bar"),
					condition.ElseIf(y < 20, es.Exists("fizz")),
					condition.Else(es.Term("default", "value")),
				),
			),
	)

	// Then
	assert.NotNil(t, query)
	bodyJSON := assert.MarshalWithoutError(t, query)
	assert.Equal(t, "{\"query\":{\"bool\":{\"filter\":[{\"exists\":{\"field\":\"fizz\"}}]}}}", bodyJSON)
}

func Test_IfElse_multiple_ElseIf_should_return_correct_item(t *testing.T) {
	t.Parallel()
	// Given
	x := 5
	y := 25
	z := 3

	// When
	query := es.NewQuery(
		es.Bool().
			Filter(
				condition.IfElse(x > 10, es.Term("first", "value"),
					condition.ElseIf(y < 20, es.Term("second", "value")),
					condition.ElseIf(z < 5, es.Term("third", "value")),
				),
			),
	)

	// Then
	assert.NotNil(t, query)
	bodyJSON := assert.MarshalWithoutError(t, query)
	assert.Equal(t, "{\"query\":{\"bool\":{\"filter\":[{\"term\":{\"third\":{\"value\":\"value\"}}}]}}}", bodyJSON)
}

func Test_IfElse_should_work_with_slices(t *testing.T) {
	t.Parallel()
	// Given
	cond := true

	// When
	result := condition.IfElse(cond, es.Array{1, 2, 3})

	// Then
	assert.NotNil(t, result)
	bodyJSON := assert.MarshalWithoutError(t, result)
	assert.Equal(t, "[1,2,3]", bodyJSON)
}

func Test_IfElse_Else_should_work_with_slices(t *testing.T) {
	t.Parallel()
	// Given
	cond := false

	// When
	result := condition.IfElse(cond, es.Array{1, 2, 3},
		condition.Else(es.Array{4, 5, 6}),
	)

	// Then
	assert.NotNil(t, result)
	bodyJSON := assert.MarshalWithoutError(t, result)
	assert.Equal(t, "[4,5,6]", bodyJSON)
}
