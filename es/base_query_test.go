package es_test

import (
	"reflect"
	"testing"

	Mode "github.com/Trendyol/es-query-builder/es/enums/sort/mode"
	Order "github.com/Trendyol/es-query-builder/es/enums/sort/order"

	"github.com/Trendyol/es-query-builder/es"
	"github.com/Trendyol/es-query-builder/test/assert"
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

////   TrackTotalHits   ////

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

////   Size   ////

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

////   From   ////

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

////   Sort   ////

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

////   Source   ////

func Test_Object_should_have_SourceIncludes_method(t *testing.T) {
	// Given
	b := es.NewQuery(nil)

	// When Then
	assert.NotNil(t, b.SourceIncludes())
}

func Test_SourceIncludes_should_add_source_field_with_includes_to_Object(t *testing.T) {
	// Given
	query := es.NewQuery(nil)

	// When
	_, beforeExists := query["_source"]
	query.SourceIncludes("test")
	source, afterExists := query["_source"]

	// Then
	assert.False(t, beforeExists)
	assert.True(t, afterExists)
	assert.NotNil(t, source)
	assert.IsTypeString(t, "es.Object", source)
	bodyJSON := assert.MarshalWithoutError(t, query)
	assert.Equal(t, "{\"_source\":{\"includes\":[\"test\"]},\"query\":{}}", bodyJSON)
}

func Test_SourceIncludes_should_apped_includes_when_it_already_exists_in_the_source(t *testing.T) {
	// Given
	query := es.NewQuery(nil).SourceIncludes("first", "second")

	// When
	_, beforeExists := query["_source"]
	query.
		SourceIncludes("third").
		SourceIncludes("fourth", "fifth")
	source, afterExists := query["_source"]

	// Then
	assert.True(t, beforeExists)
	assert.True(t, afterExists)
	assert.NotNil(t, source)
	assert.IsTypeString(t, "es.Object", source)
	bodyJSON := assert.MarshalWithoutError(t, query)
	// nolint:golint,lll
	assert.Equal(t, "{\"_source\":{\"includes\":[\"first\",\"second\",\"third\",\"fourth\",\"fifth\"]},\"query\":{}}", bodyJSON)
}

func Test_SourceIncludes_should_not_add_includes_to_Object_when_fields_are_empty(t *testing.T) {
	// Given
	query := es.NewQuery(nil)

	// When
	_, beforeExists := query["_source"]
	query.SourceIncludes() // empty
	source, afterExists := query["_source"]

	// Then
	assert.False(t, beforeExists)
	assert.False(t, afterExists)
	assert.Nil(t, source)
	bodyJSON := assert.MarshalWithoutError(t, query)
	assert.Equal(t, "{\"query\":{}}", bodyJSON)
}

func Test_Object_should_have_SourceExcludes_method(t *testing.T) {
	// Given
	b := es.NewQuery(nil)

	// When Then
	assert.NotNil(t, b.SourceExcludes())
}

func Test_SourceExcludes_should_add_source_field_with_excludes_to_Object(t *testing.T) {
	// Given
	query := es.NewQuery(nil)

	// When
	_, beforeExists := query["_source"]
	query.SourceExcludes("test")
	source, afterExists := query["_source"]

	// Then
	assert.False(t, beforeExists)
	assert.True(t, afterExists)
	assert.NotNil(t, source)
	assert.IsTypeString(t, "es.Object", source)
	bodyJSON := assert.MarshalWithoutError(t, query)
	assert.Equal(t, "{\"_source\":{\"excludes\":[\"test\"]},\"query\":{}}", bodyJSON)
}

func Test_SourceExcludes_should_apped_excludes_when_it_already_exists_in_the_source(t *testing.T) {
	// Given
	query := es.NewQuery(nil).SourceExcludes("first", "second")

	// When
	_, beforeExists := query["_source"]
	query.
		SourceExcludes("third").
		SourceExcludes("fourth", "fifth")
	source, afterExists := query["_source"]

	// Then
	assert.True(t, beforeExists)
	assert.True(t, afterExists)
	assert.NotNil(t, source)
	assert.IsTypeString(t, "es.Object", source)
	bodyJSON := assert.MarshalWithoutError(t, query)
	// nolint:golint,lll
	assert.Equal(t, "{\"_source\":{\"excludes\":[\"first\",\"second\",\"third\",\"fourth\",\"fifth\"]},\"query\":{}}", bodyJSON)
}

func Test_SourceExcludes_should_not_add_excludes_to_Object_when_fields_are_empty(t *testing.T) {
	// Given
	query := es.NewQuery(nil)

	// When
	_, beforeExists := query["_source"]
	query.SourceExcludes() // empty
	source, afterExists := query["_source"]

	// Then
	assert.False(t, beforeExists)
	assert.False(t, afterExists)
	assert.Nil(t, source)
	bodyJSON := assert.MarshalWithoutError(t, query)
	assert.Equal(t, "{\"query\":{}}", bodyJSON)
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

////   Aggs   ////

func Test_Object_should_have_Aggs_method(t *testing.T) {
	// Given
	b := es.NewQuery(nil)

	// When Then
	assert.NotNil(t, b.Aggs)
}
