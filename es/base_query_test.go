package es_test

import (
	"reflect"
	"testing"

	"github.com/GokselKUCUKSAHIN/es-query-builder/es"
	"github.com/GokselKUCUKSAHIN/es-query-builder/test/assert"

	Mode "github.com/GokselKUCUKSAHIN/es-query-builder/es/enums/sort/mode"
	Order "github.com/GokselKUCUKSAHIN/es-query-builder/es/enums/sort/order"
)

////   NewQuery   ////

func Test_NewQuery_should_exist_on_es_package(t *testing.T) {
	// Given When Then
	assert.NotNil(t, es.NewQuery)
}

func Test_NewQuery_should_create_a_new_Object(t *testing.T) {
	// Given When
	bodyA := es.NewQuery(nil)
	bodyB := es.NewQuery(nil)

	// Then
	assert.NotNil(t, bodyA)
	assert.NotNil(t, bodyB)
	assert.Equal(t, bodyA, bodyB)
	assert.NotEqualReference(t, bodyA, bodyB)
	assert.MarshalWithoutError(t, bodyA)
	assert.MarshalWithoutError(t, bodyB)
}

func Test_NewQuery_should_return_type_of_Object(t *testing.T) {
	// Given
	query := es.NewQuery(nil)

	// When
	bodyType := reflect.TypeOf(query).String()

	// Then
	assert.NotNil(t, query)
	assert.Equal(t, "es.Object", bodyType)
	assert.MarshalWithoutError(t, query)
}

func Test_NewQuery_should_add_query_field_into_Object(t *testing.T) {
	// Given
	query := es.NewQuery(nil)

	// When
	q, exists := query["query"]

	// Then
	assert.True(t, exists)
	assert.NotNil(t, q)
}

func Test_NewQuery_should_create_json_with_query_field(t *testing.T) {
	// Given
	query := es.NewQuery(nil)

	// When Then
	assert.NotNil(t, query)
	bodyJSON := assert.MarshalWithoutError(t, query)
	assert.Equal(t, "{\"query\":{}}", bodyJSON)
}

func Test_NewQuery_Bool_should_create_json_with_bool_field_inside_query(t *testing.T) {
	// Given
	query := es.NewQuery(
		es.Bool(),
	)

	// When Then
	assert.NotNil(t, query)
	bodyJSON := assert.MarshalWithoutError(t, query)
	assert.Equal(t, "{\"query\":{\"bool\":{}}}", bodyJSON)
}

////   Object   ////

func Test_Object_should_have_TrackTotalHits_method(t *testing.T) {
	// Given
	b := es.NewQuery(nil)

	// When Then
	assert.NotNil(t, b.TrackTotalHits)
}

func Test_TrackTotalHits_should_add_track_total_hits_field_into_Object(t *testing.T) {
	// Given
	query := es.NewQuery(nil)

	// When
	_, beforeExists := query["track_total_hits"]
	object := query.TrackTotalHits(true)
	trackTotalHits, afterExists := query["track_total_hits"]

	// Then
	assert.NotNil(t, object)
	assert.False(t, beforeExists)
	assert.True(t, afterExists)
	assert.True(t, trackTotalHits.(bool))
}

func Test_Object_should_have_Size_method(t *testing.T) {
	// Given
	b := es.NewQuery(nil)

	// When Then
	assert.NotNil(t, b.Size)
}

func Test_Size_should_add_size_field_into_Object(t *testing.T) {
	// Given
	query := es.NewQuery(nil)

	// When
	_, beforeExists := query["size"]
	object := query.Size(123)
	size, afterExists := query["size"]

	// Then
	assert.NotNil(t, object)
	assert.False(t, beforeExists)
	assert.True(t, afterExists)
	assert.Equal(t, 123, size.(int))
}

func Test_Object_should_have_From_method(t *testing.T) {
	// Given
	b := es.NewQuery(nil)

	// When Then
	assert.NotNil(t, b.From)
}

func Test_From_should_add_from_field_into_Object(t *testing.T) {
	// Given
	query := es.NewQuery(nil)

	// When
	_, beforeExists := query["from"]
	object := query.From(1500)
	from, afterExists := query["from"]

	// Then
	assert.NotNil(t, object)
	assert.False(t, beforeExists)
	assert.True(t, afterExists)
	assert.Equal(t, 1500, from.(int))
}

func Test_Object_should_have_Sort_method(t *testing.T) {
	// Given
	b := es.NewQuery(nil)

	// When Then
	assert.NotNil(t, b.Sort)
}

func Test_Sort_should_exist_on_es_package(t *testing.T) {
	// Given When Then
	assert.NotNil(t, es.Sort)
}

func Test_Sort_should_return_sortType_with_order(t *testing.T) {
	// Given
	sort := es.Sort("name").Order(Order.Asc)

	// When
	bodyType := reflect.TypeOf(sort).String()

	// Then
	assert.NotNil(t, sort)
	assert.Equal(t, "es.sortType", bodyType)
	bodyJSON := assert.MarshalWithoutError(t, sort)
	assert.Equal(t, "{\"name\":{\"order\":\"asc\"}}", bodyJSON)
}

func Test_Sort_should_return_sortType_with_mode(t *testing.T) {
	// Given
	sort := es.Sort("age").Mode(Mode.Median)

	// When
	bodyType := reflect.TypeOf(sort).String()

	// Then
	assert.NotNil(t, sort)
	assert.Equal(t, "es.sortType", bodyType)
	bodyJSON := assert.MarshalWithoutError(t, sort)
	assert.Equal(t, "{\"age\":{\"mode\":\"median\"}}", bodyJSON)
}

func Test_Sort_should_return_sortType_with_order_and_mode(t *testing.T) {
	// Given
	sort := es.Sort("salary").Order(Order.Desc).Mode(Mode.Sum)

	// When
	bodyType := reflect.TypeOf(sort).String()

	// Then
	assert.NotNil(t, sort)
	assert.Equal(t, "es.sortType", bodyType)
	bodyJSON := assert.MarshalWithoutError(t, sort)
	assert.Equal(t, "{\"salary\":{\"mode\":\"sum\",\"order\":\"desc\"}}", bodyJSON)
}

func Test_Sort_should_add_sort_field_into_Object(t *testing.T) {
	// Given
	query := es.NewQuery(nil)

	// When
	_, beforeExists := query["sort"]
	query.Sort(es.Sort("name").Order(Order.Desc))
	sort, afterExists := query["sort"]

	// Then
	assert.NotNil(t, sort)
	assert.False(t, beforeExists)
	assert.True(t, afterExists)
	assert.IsTypeString(t, "[]es.sortType", sort)
	bodyJSON := assert.MarshalWithoutError(t, query)
	assert.Equal(t, "{\"query\":{},\"sort\":[{\"name\":{\"order\":\"desc\"}}]}", bodyJSON)
}

func Test_Object_should_have_Source_method(t *testing.T) {
	// Given
	b := es.NewQuery(nil)

	// When Then
	assert.NotNil(t, b.Source)
}

func Test_Source_should_add_source_field_into_Object(t *testing.T) {
	// Given
	query := es.NewQuery(nil)

	// When
	_, beforeExists := query["_source"]
	query.Source()
	source, afterExists := query["_source"]

	// Then
	assert.NotNil(t, source)
	assert.False(t, beforeExists)
	assert.True(t, afterExists)
	assert.IsTypeString(t, "es.sourceType", source)
	bodyJSON := assert.MarshalWithoutError(t, query)
	assert.Equal(t, "{\"_source\":{},\"query\":{}}", bodyJSON)
}

func Test_Source_should_have_Includes_method(t *testing.T) {
	// Given
	query := es.NewQuery(nil)

	// When
	source := query.Source()

	// Then
	assert.NotNil(t, source)
	assert.IsTypeString(t, "es.sourceType", source)
	assert.NotNil(t, source.Includes)
}

func Test_Source_should_have_Excludes_method(t *testing.T) {
	// Given
	query := es.NewQuery(nil)

	// When
	source := query.Source()

	// Then
	assert.NotNil(t, source)
	assert.IsTypeString(t, "es.sourceType", source)
	assert.NotNil(t, source.Excludes)
}

func Test_Source_should_create_json_with_source_field(t *testing.T) {
	// Given
	query := es.NewQuery(nil)

	// When
	query.Source().
		Includes("hello", "world").
		Excludes("Lorem", "Ipsum")

	bodyJSON := assert.MarshalWithoutError(t, query)
	assert.Equal(t, "{\"_source\":{\"excludes\":[\"Lorem\",\"Ipsum\"],\"includes\":[\"hello\",\"world\"]},\"query\":{}}", bodyJSON)
}

func Test_Source_should_append_existing_fields(t *testing.T) {
	// Given
	query := es.NewQuery(nil)

	// When
	query.Source().
		Includes("hello", "world").
		Excludes("Lorem", "Ipsum").
		Includes("golang", "gopher").
		Excludes("Metallica", "Iron Maiden")

	bodyJSON := assert.MarshalWithoutError(t, query)
	// nolint:golint,lll
	assert.Equal(t, "{\"_source\":{\"excludes\":[\"Lorem\",\"Ipsum\",\"Metallica\",\"Iron Maiden\"],\"includes\":[\"hello\",\"world\",\"golang\",\"gopher\"]},\"query\":{}}", bodyJSON)
}

func Test_Object_should_have_SourceFalse_method(t *testing.T) {
	// Given
	b := es.NewQuery(nil)

	// When Then
	assert.NotNil(t, b.SourceFalse)
}

func Test_SourceFalse_should_set_source_field_as_false(t *testing.T) {
	// Given
	query := es.NewQuery(nil)

	// When
	_, beforeExists := query["_source"]
	query.SourceFalse()
	source, afterExists := query["_source"]

	// Then
	assert.NotNil(t, source)
	assert.False(t, beforeExists)
	assert.True(t, afterExists)
	assert.False(t, source.(bool))
	bodyJSON := assert.MarshalWithoutError(t, query)
	assert.Equal(t, "{\"_source\":false,\"query\":{}}", bodyJSON)
}

func Test_Object_should_have_Aggs_method(t *testing.T) {
	// Given
	b := es.NewQuery(nil)

	// When Then
	assert.NotNil(t, b.Aggs)
}
