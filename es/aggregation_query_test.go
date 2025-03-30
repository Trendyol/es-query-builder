package es_test

import (
	"reflect"
	"testing"

	"github.com/Trendyol/es-query-builder/es"
	"github.com/Trendyol/es-query-builder/test/assert"
)

func Test_NewAggs_should_exist_on_es_package(t *testing.T) {
	t.Parallel()
	// Given When Then
	assert.NotNil(t, es.NewAggs)
}

func Test_NewAggs_should_return_type_of_Object(t *testing.T) {
	t.Parallel()
	// Given
	aggsQuery := es.NewAggs()

	// When
	bodyType := reflect.TypeOf(aggsQuery).String()

	// Then
	assert.NotNil(t, aggsQuery)
	assert.Equal(t, "es.Object", bodyType)
	assert.MarshalWithoutError(t, aggsQuery)
}

func Test_NewAggs_should_add_aggs_field_into_Object(t *testing.T) {
	t.Parallel()
	// Given
	aggsQuery := es.NewAggs(nil)

	// When
	q, exists := aggsQuery["aggs"]

	// Then
	assert.True(t, exists)
	assert.NotNil(t, q)
}

func Test_Agg_should_exist_on_es_package(t *testing.T) {
	t.Parallel()
	// Given When Then
	assert.NotNil(t, es.Agg[es.Object])
}

func Test_Agg_should_return_type_of_aggsType(t *testing.T) {
	t.Parallel()
	// Given
	agg := es.Agg("agg_name", es.Object{})

	// When
	bodyType := reflect.TypeOf(agg).String()

	// Then
	assert.NotNil(t, agg)
	assert.Equal(t, "es.aggsType", bodyType)
	assert.MarshalWithoutError(t, agg)
}

func Test_Agg_should_add_given_name_field_into_aggsType(t *testing.T) {
	t.Parallel()
	// Given
	aggsQuery := es.Agg("agg_name", es.Object{})

	// When
	q, exists := aggsQuery["agg_name"]

	// Then
	assert.True(t, exists)
	assert.NotNil(t, q)
}

func Test_Object_should_have_Query_method(t *testing.T) {
	t.Parallel()
	// Given
	b := es.NewQuery(nil)

	// When Then
	assert.NotNil(t, b.Query)
}

func Test_Query_should_add_query_field_into_Object(t *testing.T) {
	t.Parallel()
	// Given
	query := es.NewAggs(nil)

	// When
	_, beforeExists := query["query"]
	object := query.Query(es.Term("hello", "world"))
	_, afterExists := query["query"]

	// Then
	assert.NotNil(t, object)
	assert.False(t, beforeExists)
	assert.True(t, afterExists)
}
