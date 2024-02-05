package es_test

import (
	"github.com/GokselKUCUKSAHIN/es-query-builder/es"
	"github.com/GokselKUCUKSAHIN/es-query-builder/es/test/assert"
	"reflect"
	"testing"
)

func Test_New_should_creates_a_new_Object(t *testing.T) {
	// Given When
	bodyA := es.New()
	bodyB := es.New()

	// Then
	assert.NotNil(t, bodyA)
	assert.NotNil(t, bodyB)
	assert.Equal(t, bodyA, bodyB)
	assert.NotEqualReference(t, bodyA, bodyB)
	assert.MarshalWithoutError(t, bodyA)
	assert.MarshalWithoutError(t, bodyB)
}

func Test_New_should_return_type_of_Object(t *testing.T) {
	// Given
	body := es.New()

	// When
	bodyType := reflect.TypeOf(body).String()

	// Then
	assert.NotNil(t, body)
	assert.Equal(t, "es.Object", bodyType)
	assert.MarshalWithoutError(t, body)
}

func Test_New_should_creates_empty_Object(t *testing.T) {
	// Given
	body := es.New()

	// When Then
	assert.NotNil(t, body)
	bodyJSON := assert.MarshalWithoutError(t, body)
	assert.Equal(t, bodyJSON, "{}")
}

func Test_Object_should_has_Query_method(t *testing.T) {
	// Given
	body := es.New()

	// When Then
	assert.NotNil(t, body.Query)
}

func Test_Query_method_should_create_queryType(t *testing.T) {
	// Given
	body := es.New()

	// When
	query := body.Query()

	//Then
	assert.NotNil(t, query)
	assert.IsTypeString(t, "es.queryType", query)
}

func Test_Query_should_add_query_field_onto_Object(t *testing.T) {
	// Given
	body := es.New()

	// When
	query := body.Query()

	//Then
	q, exists := body["query"]
	assert.True(t, exists)
	assert.NotNil(t, q)
	assert.EqualReference(t, query, q)
}

func Test_Query_should_replace_existing_query_when_it_called_twice(t *testing.T) {
	// Given
	body := es.New()

	// When
	q1 := body.Query()
	q2 := body.Query()

	//Then
	assert.IsType(t, q1, q2)
	q, exists := body["query"]
	assert.True(t, exists)
	assert.NotNil(t, q)
	assert.NotEqualReference(t, q1, q)
	assert.EqualReference(t, q2, q)
}

func Test_Query_should_has_Bool_method(t *testing.T) {
	// Given
	body := es.New()
	query := body.Query()

	// When Then
	assert.NotNil(t, query.Bool)
}

func Test_Bool_method_should_create_boolType(t *testing.T) {
	// Given
	body := es.New()

	// When
	query := body.Query()
	b := query.Bool()

	//Then
	assert.NotNil(t, b)
	assert.IsTypeString(t, "es.boolType", b)
}

func Test_Bool_should_add_bool_field_onto_Query(t *testing.T) {
	// Given
	body := es.New()
	query := body.Query()

	// When
	b := query.Bool()

	//Then
	b1, exists := query["bool"]
	assert.True(t, exists)
	assert.NotNil(t, b1)
	assert.EqualReference(t, b, b1)
}

func Test_Bool_should_replace_existing_bool_when_it_called_twice(t *testing.T) {
	// Given
	body := es.New()
	query := body.Query()

	// When
	b1 := query.Bool()
	b2 := query.Bool()

	//Then
	assert.IsType(t, b1, b2)
	b, exists := query["bool"]
	assert.True(t, exists)
	assert.NotNil(t, b)
	assert.NotEqualReference(t, b1, b)
	assert.EqualReference(t, b2, b)
}
