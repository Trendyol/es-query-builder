package es_test

import (
	"testing"

	"github.com/Trendyol/es-query-builder/es"
	"github.com/Trendyol/es-query-builder/test/assert"
)

func Test_ConstantScore_should_exist_on_es_package(t *testing.T) {
	t.Parallel()
	// Given When Then
	assert.NotNil(t, es.ConstantScore)
}

func Test_ConstantScore_should_create_json_with_constant_score_field_inside_query(t *testing.T) {
	t.Parallel()
	// Given
	query := es.NewQuery(
		es.ConstantScore(es.Term("name", "Göksel")),
	)

	// When Then
	assert.NotNil(t, query)
	bodyJSON := assert.MarshalWithoutError(t, query)
	assert.Equal(t, "{\"query\":{\"constant_score\":{\"filter\":{\"term\":{\"name\":{\"value\":\"Göksel\"}}}}}}", bodyJSON)
}

func Test_ConstantScore_method_should_create_constantScoreType(t *testing.T) {
	t.Parallel()
	// Given
	constantScore := es.ConstantScore(es.Term("pi", 3.1415926535897))

	// Then
	assert.NotNil(t, constantScore)
	assert.IsTypeString(t, "es.constantScoreType", constantScore)
}

func Test_ConstantScore_should_have_Boost_method(t *testing.T) {
	t.Parallel()
	// Given
	constantScore := es.ConstantScore(nil)

	// When Then
	assert.NotNil(t, constantScore.Boost)
}

func Test_ConstantScore_Boost_should_create_json_with_boost_field_inside_constant_score(t *testing.T) {
	t.Parallel()
	// Given
	query := es.NewQuery(
		es.ConstantScore(
			es.Range("timestamp").
				GreaterThan("2020"),
		).Boost(3.14),
	)

	// When Then
	assert.NotNil(t, query)
	bodyJSON := assert.MarshalWithoutError(t, query)
	assert.Equal(t, "{\"query\":{\"constant_score\":{\"boost\":3.14,\"filter\":{\"range\":{\"timestamp\":{\"gt\":\"2020\"}}}}}}", bodyJSON)
}

func Test_ConstantScore_should_have_Name_method(t *testing.T) {
	t.Parallel()
	// Given
	constantScore := es.ConstantScore(nil)

	// When Then
	assert.NotNil(t, constantScore.Boost)
}

func Test_ConstantScore_Name_should_create_json_with__name_field_inside_constant_score(t *testing.T) {
	t.Parallel()
	// Given
	query := es.NewQuery(
		es.ConstantScore(
			es.Range("timestamp").
				GreaterThan("2020"),
		).Name("query_name"),
	)

	// When Then
	assert.NotNil(t, query)
	bodyJSON := assert.MarshalWithoutError(t, query)
	// nolint:golint,lll
	assert.Equal(t, "{\"query\":{\"constant_score\":{\"_name\":\"query_name\",\"filter\":{\"range\":{\"timestamp\":{\"gt\":\"2020\"}}}}}}", bodyJSON)
}
