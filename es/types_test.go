package es_test

import (
	"fmt"
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

	// Then
	assert.NotNil(t, query)
	assert.IsTypeString(t, "es.queryType", query)
}

func Test_Query_should_add_query_field_onto_Object(t *testing.T) {
	// Given
	body := es.New()

	// When
	query := body.Query()
	q, exists := body["query"]

	// Then
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
	q, exists := body["query"]

	// Then
	assert.IsType(t, q1, q2)
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

func Test_Query_should_creates_json_with_query_field(t *testing.T) {
	// Given
	body := es.New()
	query := body.Query()

	// When Then
	assert.NotNil(t, query)
	bodyJSON := assert.MarshalWithoutError(t, body)
	assert.Equal(t, bodyJSON, "{\"query\":{}}")
}

func Test_Bool_method_should_create_boolType(t *testing.T) {
	// Given
	body := es.New()

	// When
	query := body.Query()
	b := query.Bool()

	// Then
	assert.NotNil(t, b)
	assert.IsTypeString(t, "es.boolType", b)
}

func Test_Bool_should_add_bool_field_onto_Query(t *testing.T) {
	// Given
	body := es.New()
	query := body.Query()

	// When
	b := query.Bool()
	b1, exists := query["bool"]

	// Then
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
	b, exists := query["bool"]

	// Then
	assert.IsType(t, b1, b2)
	assert.True(t, exists)
	assert.NotNil(t, b)
	assert.NotEqualReference(t, b1, b)
	assert.EqualReference(t, b2, b)
}

func Test_Bool_should_creates_json_with_bool_field(t *testing.T) {
	// Given
	body := es.New()
	query := body.Query()
	b := query.Bool()

	// When Then
	assert.NotNil(t, b)
	bodyJSON := assert.MarshalWithoutError(t, body)
	assert.Equal(t, bodyJSON, "{\"query\":{\"bool\":{}}}")
}

func Test_Bool_should_have_Filter_method(t *testing.T) {
	// Given
	body := es.New()
	query := body.Query()
	_bool := query.Bool()

	// When Then
	assert.NotNil(t, _bool.Filter)
}

func Test_Filter_method_should_return_boolType(t *testing.T) {
	// Given
	body := es.New()
	query := body.Query()
	_bool := query.Bool()

	// When
	filter := _bool.Filter()

	// Then
	assert.NotNil(t, filter)
	assert.IsTypeString(t, "es.boolType", filter)
}

func Test_Bool_should_have_Must_method(t *testing.T) {
	// Given
	body := es.New()
	query := body.Query()
	_bool := query.Bool()

	// When Then
	assert.NotNil(t, _bool.Must)
}

func Test_Must_method_should_return_boolType(t *testing.T) {
	// Given
	body := es.New()
	query := body.Query()
	_bool := query.Bool()

	// When
	must := _bool.Must()

	// Then
	assert.NotNil(t, must)
	assert.IsTypeString(t, "es.boolType", must)
}

func Test_Bool_should_have_Should_method(t *testing.T) {
	// Given
	body := es.New()
	query := body.Query()
	_bool := query.Bool()

	// When Then
	assert.NotNil(t, _bool.Should)
}

func Test_Should_method_should_return_boolType(t *testing.T) {
	// Given
	body := es.New()
	query := body.Query()
	_bool := query.Bool()

	// When
	should := _bool.Should()

	// Then
	assert.NotNil(t, should)
	assert.IsTypeString(t, "es.boolType", should)
}

func Test_Bool_should_have_MustNot_method(t *testing.T) {
	// Given
	body := es.New()
	query := body.Query()
	_bool := query.Bool()

	// When Then
	assert.NotNil(t, _bool.MustNot)
}

func Test_MustNot_method_should_return_boolType(t *testing.T) {
	// Given
	body := es.New()
	query := body.Query()
	_bool := query.Bool()

	// When
	mustNot := _bool.MustNot()

	// Then
	assert.NotNil(t, mustNot)
	assert.IsTypeString(t, "es.boolType", mustNot)
}

func Test_Term_should_exist_on_es_package(t *testing.T) {
	// Given When Then
	assert.NotNil(t, es.Term)
}

func Test_Term_method_should_create_termType(t *testing.T) {
	// Given
	term := es.Term("test", 123)

	// When Then
	assert.NotNil(t, term)
	assert.IsTypeString(t, "es.termType", term)
}

func Test_Term_method_should_create_termType_with_term_Object(t *testing.T) {
	// Given
	key := "test"
	value := 123
	term := es.Term(key, value)

	// When
	termObject, exists := term["term"]

	// Then
	assert.True(t, exists)
	assert.NotNil(t, termObject)
	assert.IsTypeString(t, "es.Object", termObject)
}

func Test_Term_method_should_create_termType_with_given_key_value(t *testing.T) {
	// Given
	key := "test"
	value := 123
	term := es.Term(key, value)

	// When
	termObject, termExists := term["term"]
	termObject, ok := termObject.(es.Object)
	termValue, keyExists := termObject.(es.Object)[key]

	// Then
	assert.True(t, termExists)
	assert.True(t, ok)
	assert.NotNil(t, termObject)
	assert.True(t, keyExists)
	assert.NotNil(t, termValue)
	assert.Equal(t, value, termValue)
}

func Test_Term_should_add_into_Bool_with_no_errors(t *testing.T) {
	// Given
	body := es.New()
	b := body.
		Query().
		Bool()

	key := "test"
	value := 123
	term := es.Term(key, value)

	// When
	b2 := b.Filter(term)

	// Then
	assert.NotNil(t, b)
	assert.NotNil(t, b2)
	assert.EqualReference(t, b, b2)
	bodyJSON := assert.MarshalWithoutError(t, body)
	assert.Equal(t, bodyJSON, fmt.Sprintf("{\"query\":{\"bool\":{\"filter\":[{\"term\":{\"%s\":%d}}]}}}", key, value))
}

func Test_Term_should_creates_json_with_key_value(t *testing.T) {
	// Given
	key := "test"
	value := 123
	term := es.Term(key, value)

	// When Then
	assert.NotNil(t, term)
	termJSON := assert.MarshalWithoutError(t, term)
	assert.Equal(t, termJSON, "{\"term\":{\"test\":123}}")
}
