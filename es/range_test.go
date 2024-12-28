package es_test

import (
	range_relation "github.com/Trendyol/es-query-builder/es/enums/range-relation"
	"testing"

	"github.com/Trendyol/es-query-builder/es"
	"github.com/Trendyol/es-query-builder/test/assert"
)

////   Range   ////

func Test_Range_should_exist_on_es_package(t *testing.T) {
	// Given When Then
	assert.NotNil(t, es.Range)
}

func Test_Range_should_add_range_field_when_inside_query(t *testing.T) {
	// Given
	query := es.NewQuery(
		es.Bool().
			Must(
				es.Range("age").
					GreaterThanOrEqual(10).
					LesserThanOrEqual(20),
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
	// Given
	b := es.Range("age")
	query := es.NewQuery(b)

	// Then
	assert.NotNil(t, query)
	assert.NotNil(t, b)
	assert.IsTypeString(t, "es.rangeType", b)
}

func Test_Range_should_create_json_with_range_field_inside_query(t *testing.T) {
	// Given
	query := es.NewQuery(
		es.Range("age").
			GreaterThanOrEqual(10).
			LesserThanOrEqual(20),
	)

	// When Then
	assert.NotNil(t, query)
	bodyJSON := assert.MarshalWithoutError(t, query)
	assert.Equal(t, "{\"query\":{\"range\":{\"age\":{\"gte\":10,\"lte\":20}}}}", bodyJSON)
}

func Test_Range_should_have_LesserThan_method(t *testing.T) {
	// Given
	r := es.Range("age")

	// When Then
	assert.NotNil(t, r)
	assert.NotNil(t, r.LesserThan)
}

func Test_Range_should_have_LesserThanOrEqual_method(t *testing.T) {
	// Given
	r := es.Range("age")

	// When Then
	assert.NotNil(t, r)
	assert.NotNil(t, r.LesserThanOrEqual)
}

func Test_Range_should_have_GreaterThan_method(t *testing.T) {
	// Given
	r := es.Range("age")

	// When Then
	assert.NotNil(t, r)
	assert.NotNil(t, r.GreaterThan)
}

func Test_Range_should_have_GreaterThanOrEqual_method(t *testing.T) {
	// Given
	r := es.Range("age")

	// When Then
	assert.NotNil(t, r)
	assert.NotNil(t, r.GreaterThanOrEqual)
}

func Test_Range_gte_should_override_gt_and_vise_versa(t *testing.T) {
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
	// Given
	query := es.NewQuery(
		es.Range("age").
			LesserThan(11).
			LesserThanOrEqual(23),
	)

	// When Then
	assert.NotNil(t, query)
	bodyJSON := assert.MarshalWithoutError(t, query)
	assert.Equal(t, "{\"query\":{\"range\":{\"age\":{\"lte\":23}}}}", bodyJSON)
}

func Test_Range_should_have_Format_method(t *testing.T) {
	// Given
	r := es.Range("age")

	// When Then
	assert.NotNil(t, r)
	assert.NotNil(t, r.Format)
}

func Test_Range_Format_should_create_json_with_range_field_inside_query(t *testing.T) {
	// Given
	query := es.NewQuery(
		es.Range("birth-date").
			GreaterThanOrEqual("1990-01-01").
			LesserThanOrEqual("2024-04-12").
			Format("yyyy-MM-dd"),
	)

	// When Then
	assert.NotNil(t, query)
	bodyJSON := assert.MarshalWithoutError(t, query)
	// nolint:golint,lll
	assert.Equal(t, "{\"query\":{\"range\":{\"birth-date\":{\"format\":\"yyyy-MM-dd\",\"gte\":\"1990-01-01\",\"lte\":\"2024-04-12\"}}}}", bodyJSON)
}

func Test_Range_should_have_Boost_method(t *testing.T) {
	// Given
	r := es.Range("age")

	// When Then
	assert.NotNil(t, r)
	assert.NotNil(t, r.Boost)
}

func Test_Range_Boost_should_create_json_with_range_field_inside_query(t *testing.T) {
	// Given
	query := es.NewQuery(
		es.Range("partition").
			GreaterThan(112).
			LesserThan(765).
			Boost(3.2),
	)

	// When Then
	assert.NotNil(t, query)
	bodyJSON := assert.MarshalWithoutError(t, query)
	// nolint:golint,lll
	assert.Equal(t, "{\"query\":{\"range\":{\"partition\":{\"boost\":3.2,\"gt\":112,\"lt\":765}}}}", bodyJSON)
}

func Test_Range_should_have_From_method(t *testing.T) {
	// Given
	r := es.Range("age")

	// When Then
	assert.NotNil(t, r)
	assert.NotNil(t, r.From)
}

func Test_Range_From_should_create_json_with_range_field_inside_query(t *testing.T) {
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

func Test_Range_should_have_To_method(t *testing.T) {
	// Given
	r := es.Range("age")

	// When Then
	assert.NotNil(t, r)
	assert.NotNil(t, r.To)
}

func Test_Range_To_should_create_json_with_range_field_inside_query(t *testing.T) {
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

func Test_Range_should_have_Relation_method(t *testing.T) {
	// Given
	r := es.Range("age")

	// When Then
	assert.NotNil(t, r)
	assert.NotNil(t, r.Relation)
}

func Test_Range_Relation_should_create_json_with_range_field_inside_query(t *testing.T) {
	// Given
	query := es.NewQuery(
		es.Range("partition").
			From(512).
			To(1024).
			Relation(range_relation.Within),
	)

	// When Then
	assert.NotNil(t, query)
	bodyJSON := assert.MarshalWithoutError(t, query)
	// nolint:golint,lll
	assert.Equal(t, "{\"query\":{\"range\":{\"partition\":{\"from\":512,\"relation\":\"within\",\"to\":1024}}}}", bodyJSON)
}
