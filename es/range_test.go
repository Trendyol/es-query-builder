package es_test

import (
	"testing"

	RangeRelation "github.com/Trendyol/es-query-builder/es/enums/range-relation"

	"github.com/Trendyol/es-query-builder/es"
	"github.com/Trendyol/es-query-builder/test/assert"
)

////   Range   ////

func Test_Range_should_exist_on_es_package(t *testing.T) {
	t.Parallel()
	// Given When Then
	assert.NotNil(t, es.Range)
}

func Test_Range_should_add_range_field_when_inside_query(t *testing.T) {
	t.Parallel()
	// Given
	query := es.NewQuery(
		es.Bool().
			Must(
				es.Range("age").
					GreaterThanOrEqual(10).
					LessThanOrEqual(20),
				es.Term("language", "tr"),
			),
	)

	// When Then
	assert.NotNil(t, query)
	bodyJSON := assert.MarshalWithoutError(t, query)
	// nolint:golint,lll
	assert.Equal(t, "{\"query\":{\"bool\":{\"must\":[{\"range\":{\"age\":{\"gte\":10,\"lte\":20}}},{\"term\":{\"language\":{\"value\":\"tr\"}}}]}}}", bodyJSON)
}

func Test_Range_method_should_create_rangeType(t *testing.T) {
	t.Parallel()
	// Given
	b := es.Range("age")
	query := es.NewQuery(b)

	// Then
	assert.NotNil(t, query)
	assert.NotNil(t, b)
	assert.IsTypeString(t, "es.rangeType", b)
}

func Test_Range_should_create_json_with_range_field_inside_query(t *testing.T) {
	t.Parallel()
	// Given
	query := es.NewQuery(
		es.Range("age").
			GreaterThanOrEqual(10).
			LessThanOrEqual(20),
	)

	// When Then
	assert.NotNil(t, query)
	bodyJSON := assert.MarshalWithoutError(t, query)
	assert.Equal(t, "{\"query\":{\"range\":{\"age\":{\"gte\":10,\"lte\":20}}}}", bodyJSON)
}

func Test_Range_should_have_LessThan_method(t *testing.T) {
	t.Parallel()
	// Given
	r := es.Range("age")

	// When Then
	assert.NotNil(t, r)
	assert.NotNil(t, r.LessThan)
}

func Test_Range_should_have_LessThanIf_method(t *testing.T) {
	t.Parallel()
	// Given
	r := es.Range("age")

	// When Then
	assert.NotNil(t, r)
	assert.NotNil(t, r.LessThanIf)
}

func Test_Range_should_create_json_with_lt_field_inside_query_when_condition_is_true(t *testing.T) {
	t.Parallel()

	// Given
	variable := 20
	query := es.NewQuery(
		es.Range("age").
			GreaterThanOrEqual(10).
			LessThanIf(variable, variable > 15),
	)

	// When Then
	assert.NotNil(t, query)
	bodyJSON := assert.MarshalWithoutError(t, query)
	assert.Equal(t, "{\"query\":{\"range\":{\"age\":{\"gte\":10,\"lt\":20}}}}", bodyJSON)
}

func Test_Range_should_not_create_json_with_lt_field_inside_query_when_condition_is_false(t *testing.T) {
	t.Parallel()

	// Given
	variable := 20
	query := es.NewQuery(
		es.Range("age").
			GreaterThanOrEqual(10).
			LessThanIf(variable, variable < 15),
	)

	// When Then
	assert.NotNil(t, query)
	bodyJSON := assert.MarshalWithoutError(t, query)
	assert.Equal(t, "{\"query\":{\"range\":{\"age\":{\"gte\":10}}}}", bodyJSON)
}

func Test_Range_should_have_LessThanOrEqual_method(t *testing.T) {
	t.Parallel()
	// Given
	r := es.Range("age")

	// When Then
	assert.NotNil(t, r)
	assert.NotNil(t, r.LessThanOrEqual)
}

func Test_Range_should_have_LessThanOrEqualIf_method(t *testing.T) {
	t.Parallel()
	// Given
	r := es.Range("age")

	// When Then
	assert.NotNil(t, r)
	assert.NotNil(t, r.LessThanOrEqualIf)
}

func Test_Range_should_create_json_with_lte_field_inside_query_when_condition_is_true(t *testing.T) {
	t.Parallel()

	// Given
	variable := 20
	query := es.NewQuery(
		es.Range("age").
			GreaterThanOrEqual(10).
			LessThanOrEqualIf(variable, variable > 15),
	)

	// When Then
	assert.NotNil(t, query)
	bodyJSON := assert.MarshalWithoutError(t, query)
	assert.Equal(t, "{\"query\":{\"range\":{\"age\":{\"gte\":10,\"lte\":20}}}}", bodyJSON)
}

func Test_Range_should_not_create_json_with_lte_field_inside_query_when_condition_is_false(t *testing.T) {
	t.Parallel()

	// Given
	variable := 20
	query := es.NewQuery(
		es.Range("age").
			GreaterThanOrEqual(10).
			LessThanOrEqualIf(variable, variable < 15),
	)

	// When Then
	assert.NotNil(t, query)
	bodyJSON := assert.MarshalWithoutError(t, query)
	assert.Equal(t, "{\"query\":{\"range\":{\"age\":{\"gte\":10}}}}", bodyJSON)
}

func Test_Range_should_have_GreaterThan_method(t *testing.T) {
	t.Parallel()
	// Given
	r := es.Range("age")

	// When Then
	assert.NotNil(t, r)
	assert.NotNil(t, r.GreaterThan)
}

func Test_Range_should_have_GreaterThanIf_method(t *testing.T) {
	t.Parallel()
	// Given
	r := es.Range("age")

	// When Then
	assert.NotNil(t, r)
	assert.NotNil(t, r.GreaterThanIf)
}

func Test_Range_should_create_json_with_gt_field_inside_query_when_condition_is_true(t *testing.T) {
	t.Parallel()

	// Given
	variable := 10
	query := es.NewQuery(
		es.Range("age").
			GreaterThanIf(variable, variable < 15).
			LessThanOrEqual(20),
	)

	// When Then
	assert.NotNil(t, query)
	bodyJSON := assert.MarshalWithoutError(t, query)
	assert.Equal(t, "{\"query\":{\"range\":{\"age\":{\"gt\":10,\"lte\":20}}}}", bodyJSON)
}

func Test_Range_should_not_create_json_with_gt_field_inside_query_when_condition_is_false(t *testing.T) {
	t.Parallel()

	// Given
	variable := 10
	query := es.NewQuery(
		es.Range("age").
			GreaterThanIf(variable, variable > 15).
			LessThanOrEqual(20),
	)

	// When Then
	assert.NotNil(t, query)
	bodyJSON := assert.MarshalWithoutError(t, query)
	assert.Equal(t, "{\"query\":{\"range\":{\"age\":{\"lte\":20}}}}", bodyJSON)
}

func Test_Range_should_have_GreaterThanOrEqual_method(t *testing.T) {
	t.Parallel()
	// Given
	r := es.Range("age")

	// When Then
	assert.NotNil(t, r)
	assert.NotNil(t, r.GreaterThanOrEqual)
}

func Test_Range_should_have_GreaterThanOrEqualIf_method(t *testing.T) {
	t.Parallel()
	// Given
	r := es.Range("age")

	// When Then
	assert.NotNil(t, r)
	assert.NotNil(t, r.GreaterThanOrEqualIf)
}

func Test_Range_should_create_json_with_gte_field_inside_query_when_condition_is_true(t *testing.T) {
	t.Parallel()

	// Given
	variable := 10
	query := es.NewQuery(
		es.Range("age").
			GreaterThanOrEqualIf(variable, variable < 15).
			LessThanOrEqual(20),
	)

	// When Then
	assert.NotNil(t, query)
	bodyJSON := assert.MarshalWithoutError(t, query)
	assert.Equal(t, "{\"query\":{\"range\":{\"age\":{\"gte\":10,\"lte\":20}}}}", bodyJSON)
}

func Test_Range_should_not_create_json_with_gte_field_inside_query_when_condition_is_false(t *testing.T) {
	t.Parallel()

	// Given
	variable := 10
	query := es.NewQuery(
		es.Range("age").
			GreaterThanOrEqualIf(variable, variable > 15).
			LessThanOrEqual(20),
	)

	// When Then
	assert.NotNil(t, query)
	bodyJSON := assert.MarshalWithoutError(t, query)
	assert.Equal(t, "{\"query\":{\"range\":{\"age\":{\"lte\":20}}}}", bodyJSON)
}

func Test_Range_gte_should_override_gt_and_vise_versa(t *testing.T) {
	t.Parallel()
	// Given
	query := es.NewQuery(
		es.Range("age").
			GreaterThanOrEqual(10).
			GreaterThan(20),
	)

	// When Then
	assert.NotNil(t, query)
	bodyJSON := assert.MarshalWithoutError(t, query)
	assert.Equal(t, "{\"query\":{\"range\":{\"age\":{\"gt\":20}}}}", bodyJSON)
}

func Test_Range_lte_should_override_lt_and_vise_versa(t *testing.T) {
	t.Parallel()
	// Given
	query := es.NewQuery(
		es.Range("age").
			LessThan(11).
			LessThanOrEqual(23),
	)

	// When Then
	assert.NotNil(t, query)
	bodyJSON := assert.MarshalWithoutError(t, query)
	assert.Equal(t, "{\"query\":{\"range\":{\"age\":{\"lte\":23}}}}", bodyJSON)
}

func Test_Range_should_have_Format_method(t *testing.T) {
	t.Parallel()
	// Given
	r := es.Range("age")

	// When Then
	assert.NotNil(t, r)
	assert.NotNil(t, r.Format)
}

func Test_Range_Format_should_create_json_with_range_field_inside_query(t *testing.T) {
	t.Parallel()
	// Given
	query := es.NewQuery(
		es.Range("birth-date").
			GreaterThanOrEqual("1990-01-01").
			LessThanOrEqual("2024-04-12").
			Format("yyyy-MM-dd"),
	)

	// When Then
	assert.NotNil(t, query)
	bodyJSON := assert.MarshalWithoutError(t, query)
	// nolint:golint,lll
	assert.Equal(t, "{\"query\":{\"range\":{\"birth-date\":{\"format\":\"yyyy-MM-dd\",\"gte\":\"1990-01-01\",\"lte\":\"2024-04-12\"}}}}", bodyJSON)
}

func Test_Range_should_have_Boost_method(t *testing.T) {
	t.Parallel()
	// Given
	r := es.Range("age")

	// When Then
	assert.NotNil(t, r)
	assert.NotNil(t, r.Boost)
}

func Test_Range_Boost_should_create_json_with_range_field_inside_query(t *testing.T) {
	t.Parallel()
	// Given
	query := es.NewQuery(
		es.Range("partition").
			GreaterThan(112).
			LessThan(765).
			Boost(3.2),
	)

	// When Then
	assert.NotNil(t, query)
	bodyJSON := assert.MarshalWithoutError(t, query)
	// nolint:golint,lll
	assert.Equal(t, "{\"query\":{\"range\":{\"partition\":{\"boost\":3.2,\"gt\":112,\"lt\":765}}}}", bodyJSON)
}

func Test_Range_should_have_From_method(t *testing.T) {
	t.Parallel()
	// Given
	r := es.Range("age")

	// When Then
	assert.NotNil(t, r)
	assert.NotNil(t, r.From)
}

func Test_Range_From_should_create_json_with_range_field_inside_query(t *testing.T) {
	t.Parallel()
	// Given
	query := es.NewQuery(
		es.Range("partition").
			From(512),
	)

	// When Then
	assert.NotNil(t, query)
	bodyJSON := assert.MarshalWithoutError(t, query)
	// nolint:golint,lll
	assert.Equal(t, "{\"query\":{\"range\":{\"partition\":{\"from\":512}}}}", bodyJSON)
}

func Test_Range_should_have_FromIf_method(t *testing.T) {
	t.Parallel()

	// Given
	r := es.Range("age")

	// When Then
	assert.NotNil(t, r)
	assert.NotNil(t, r.FromIf)
}

func Test_Range_FromIf_should_create_json_with_range_field_inside_query_when_condition_is_true(t *testing.T) {
	t.Parallel()

	// Given
	variable := 512
	query := es.NewQuery(
		es.Range("partition").
			FromIf(variable, variable > 128).
			To(1024),
	)

	// When Then
	assert.NotNil(t, query)
	bodyJSON := assert.MarshalWithoutError(t, query)
	// nolint:golint,lll
	assert.Equal(t, "{\"query\":{\"range\":{\"partition\":{\"from\":512,\"to\":1024}}}}", bodyJSON)
}

func Test_Range_FromIf_should_not_create_json_with_range_field_inside_query_when_condition_is_false(t *testing.T) {
	t.Parallel()

	// Given
	variable := 512
	query := es.NewQuery(
		es.Range("partition").
			FromIf(variable, variable < 128).
			To(1024),
	)

	// When Then
	assert.NotNil(t, query)
	bodyJSON := assert.MarshalWithoutError(t, query)
	// nolint:golint,lll
	assert.Equal(t, "{\"query\":{\"range\":{\"partition\":{\"to\":1024}}}}", bodyJSON)
}

func Test_Range_should_have_To_method(t *testing.T) {
	t.Parallel()
	// Given
	r := es.Range("age")

	// When Then
	assert.NotNil(t, r)
	assert.NotNil(t, r.To)
}

func Test_Range_To_should_create_json_with_range_field_inside_query(t *testing.T) {
	t.Parallel()
	// Given
	query := es.NewQuery(
		es.Range("partition").
			To(1024),
	)

	// When Then
	assert.NotNil(t, query)
	bodyJSON := assert.MarshalWithoutError(t, query)
	// nolint:golint,lll
	assert.Equal(t, "{\"query\":{\"range\":{\"partition\":{\"to\":1024}}}}", bodyJSON)
}

func Test_Range_should_have_ToIf_method(t *testing.T) {
	t.Parallel()
	// Given
	r := es.Range("age")

	// When Then
	assert.NotNil(t, r)
	assert.NotNil(t, r.ToIf)
}

func Test_Range_ToIf_should_create_json_with_range_field_inside_query_when_condition_is_true(t *testing.T) {
	t.Parallel()

	// Given
	variable := 1024
	query := es.NewQuery(
		es.Range("partition").
			From(128).
			ToIf(variable, variable > 512),
	)

	// When Then
	assert.NotNil(t, query)
	bodyJSON := assert.MarshalWithoutError(t, query)
	// nolint:golint,lll
	assert.Equal(t, "{\"query\":{\"range\":{\"partition\":{\"from\":128,\"to\":1024}}}}", bodyJSON)
}

func Test_Range_ToIf_should_not_create_json_with_range_field_inside_query_when_condition_is_false(t *testing.T) {
	t.Parallel()

	// Given
	variable := 1024
	query := es.NewQuery(
		es.Range("partition").
			From(128).
			ToIf(variable, variable < 512),
	)

	// When Then
	assert.NotNil(t, query)
	bodyJSON := assert.MarshalWithoutError(t, query)
	// nolint:golint,lll
	assert.Equal(t, "{\"query\":{\"range\":{\"partition\":{\"from\":128}}}}", bodyJSON)
}

func Test_Range_should_have_Relation_method(t *testing.T) {
	t.Parallel()
	// Given
	r := es.Range("age")

	// When Then
	assert.NotNil(t, r)
	assert.NotNil(t, r.Relation)
}

func Test_Range_Relation_should_create_json_with_range_field_inside_query(t *testing.T) {
	t.Parallel()
	// Given
	query := es.NewQuery(
		es.Range("partition").
			From(512).
			To(1024).
			Relation(RangeRelation.Within),
	)

	// When Then
	assert.NotNil(t, query)
	bodyJSON := assert.MarshalWithoutError(t, query)
	// nolint:golint,lll
	assert.Equal(t, "{\"query\":{\"range\":{\"partition\":{\"from\":512,\"relation\":\"within\",\"to\":1024}}}}", bodyJSON)
}
