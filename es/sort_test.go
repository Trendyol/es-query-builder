package es_test

import (
	"reflect"
	"testing"

	Mode "github.com/Trendyol/es-query-builder/es/enums/sort/mode"
	Order "github.com/Trendyol/es-query-builder/es/enums/sort/order"

	"github.com/Trendyol/es-query-builder/es"
	"github.com/Trendyol/es-query-builder/test/assert"
)

func Test_Sort_should_exist_on_es_package(t *testing.T) {
	t.Parallel()
	// Given When Then
	assert.NotNil(t, es.Sort)
}

func Test_Sort_method_should_create_sortType(t *testing.T) {
	t.Parallel()
	// Given
	sort := es.Sort("date")

	// Then
	assert.NotNil(t, sort)
	assert.IsTypeString(t, "es.sortType", sort)
}

func Test_Sort_should_have_Mode_method(t *testing.T) {
	t.Parallel()
	// Given
	sort := es.Sort("date")

	// When Then
	assert.NotNil(t, sort.Mode)
}

func Test_Mode_should_add_mode_field_into_Sort(t *testing.T) {
	t.Parallel()
	// Given
	sort := es.Sort("date").Mode(Mode.Median)

	// When Then
	assert.NotNil(t, sort)
	bodyJSON := assert.MarshalWithoutError(t, sort)
	assert.Equal(t, "{\"date\":{\"mode\":\"median\"}}", bodyJSON)
}

func Test_Sort_should_have_Order_method(t *testing.T) {
	t.Parallel()
	// Given
	sort := es.Sort("date")

	// When Then
	assert.NotNil(t, sort.Order)
}

func Test_Order_should_add_order_field_into_Sort(t *testing.T) {
	t.Parallel()
	// Given
	sort := es.Sort("date").Order(Order.Desc)

	// When Then
	assert.NotNil(t, sort)
	bodyJSON := assert.MarshalWithoutError(t, sort)
	assert.Equal(t, "{\"date\":{\"order\":\"desc\"}}", bodyJSON)
}

func Test_Sort_should_have_Nested_method(t *testing.T) {
	t.Parallel()
	// Given
	sort := es.Sort("date")

	// When Then
	assert.NotNil(t, sort.Nested)
}

func Test_Nested_should_add_nested_field_into_Sort(t *testing.T) {
	t.Parallel()
	// Given
	sort := es.Sort("date").Nested(es.NestedSort("timestamp"))

	// When Then
	assert.NotNil(t, sort)
	bodyJSON := assert.MarshalWithoutError(t, sort)
	assert.Equal(t, "{\"date\":{\"nested\":{\"path\":\"timestamp\"}}}", bodyJSON)
}

func Test_Sort_should_return_sortType_with_order(t *testing.T) {
	t.Parallel()
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
	t.Parallel()
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
	t.Parallel()
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
	t.Parallel()
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
