package es_test

import (
	"testing"

	"github.com/Trendyol/es-query-builder/es"
	"github.com/Trendyol/es-query-builder/test/assert"
)

func Test_Ids_should_exist_on_es_package(t *testing.T) {
	t.Parallel()
	// Given When Then
	assert.NotNil(t, es.Ids[string])
}

func Test_Ids_should_create_json_with_ids_field(t *testing.T) {
	t.Parallel()
	// Given
	query := es.NewQuery(
		es.Ids("1", "2", "3"),
	)

	// When Then
	assert.NotNil(t, query)
	bodyJSON := assert.MarshalWithoutError(t, query)
	assert.Equal(t, "{\"query\":{\"ids\":{\"values\":[\"1\",\"2\",\"3\"]}}}", bodyJSON)
}

func Test_IdsArray_should_create_json_with_ids_field(t *testing.T) {
	t.Parallel()
	// Given
	ids := []string{"a", "b"}
	query := es.NewQuery(
		es.IdsArray(ids),
	)

	// When Then
	assert.NotNil(t, query)
	bodyJSON := assert.MarshalWithoutError(t, query)
	assert.Equal(t, "{\"query\":{\"ids\":{\"values\":[\"a\",\"b\"]}}}", bodyJSON)
}

func Test_Ids_method_should_create_idsType(t *testing.T) {
	t.Parallel()
	// Given
	b := es.Ids("doc-1")

	// Then
	assert.NotNil(t, b)
	assert.IsTypeString(t, "es.idsType", b)
}

func Test_Ids_should_have_Boost_method(t *testing.T) {
	t.Parallel()
	// Given
	ids := es.Ids("doc-1")

	// When Then
	assert.NotNil(t, ids.Boost)
}

func Test_Ids_Boost_should_create_json_with_boost_field_inside_ids(t *testing.T) {
	t.Parallel()
	// Given
	query := es.NewQuery(
		es.Ids("doc-1", "doc-2").
			Boost(2.5),
	)

	// When Then
	assert.NotNil(t, query)
	bodyJSON := assert.MarshalWithoutError(t, query)
	assert.Equal(t, "{\"query\":{\"ids\":{\"boost\":2.5,\"values\":[\"doc-1\",\"doc-2\"]}}}", bodyJSON)
}

func Test_Ids_should_have_Name_method(t *testing.T) {
	t.Parallel()
	// Given
	ids := es.Ids("doc-1")

	// When Then
	assert.NotNil(t, ids.Name)
}

func Test_Ids_Name_should_create_json_with__name_field_inside_ids(t *testing.T) {
	t.Parallel()
	// Given
	query := es.NewQuery(
		es.Ids("doc-1", "doc-2").
			Name("ids-tag"),
	)

	// When Then
	assert.NotNil(t, query)
	bodyJSON := assert.MarshalWithoutError(t, query)
	assert.Equal(t, "{\"query\":{\"ids\":{\"_name\":\"ids-tag\",\"values\":[\"doc-1\",\"doc-2\"]}}}", bodyJSON)
}
