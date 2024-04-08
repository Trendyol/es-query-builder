package es_test

import (
	"github.com/GokselKUCUKSAHIN/es-query-builder/es"
	"github.com/GokselKUCUKSAHIN/es-query-builder/es/test/assert"
	"reflect"
	"testing"
)

////   NewQuery   ////

func Test_NewQuery_should_exist_on_es_package(t *testing.T) {
	// Given When Then
	assert.NotNil(t, es.NewQuery)
}

func Test_NewQuery_should_creates_a_new_Object(t *testing.T) {
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
	body := es.NewQuery(nil)

	// When
	bodyType := reflect.TypeOf(body).String()

	// Then
	assert.NotNil(t, body)
	assert.Equal(t, "es.Object", bodyType)
	assert.MarshalWithoutError(t, body)
}

func Test_NewQuery_should_add_query_field_into_Object(t *testing.T) {
	// Given
	body := es.NewQuery(nil)

	// When
	q, exists := body["query"]

	// Then
	assert.True(t, exists)
	assert.NotNil(t, q)
}

func Test_NewQuery_should_creates_json_with_query_field(t *testing.T) {
	// Given
	body := es.NewQuery(nil)

	// When Then
	assert.NotNil(t, body)
	bodyJSON := assert.MarshalWithoutError(t, body)
	assert.Equal(t, bodyJSON, "{\"query\":{}}")
}

func Test_NewQuery_Bool_should_create_json_with_bool_field_inside_query(t *testing.T) {
	// Given
	body := es.NewQuery(
		es.Bool(),
	)

	// When Then
	assert.NotNil(t, body)
	bodyJSON := assert.MarshalWithoutError(t, body)
	assert.Equal(t, bodyJSON, "{\"query\":{\"bool\":{}}}")
}

////   Bool   ////

func Test_Bool_should_exist_on_es_package(t *testing.T) {
	// Given When Then
	assert.NotNil(t, es.Bool)
}

func Test_Bool_method_should_create_boolType(t *testing.T) {
	// Given
	b := es.Bool()

	// Then
	assert.NotNil(t, b)
	assert.IsTypeString(t, "es.boolType", b)
}

func Test_Bool_should_has_SetMinimumShouldMatch_method(t *testing.T) {
	// Given
	b := es.Bool()

	// When Then
	assert.NotNil(t, b.SetMinimumShouldMatch)
}

func Test_Bool_SetMinimumShouldMatch_should_create_json_with_minimum_should_match_field_inside_bool(t *testing.T) {
	// Given
	body := es.NewQuery(
		es.Bool().
			SetMinimumShouldMatch(7),
	)

	// When Then
	assert.NotNil(t, body)
	bodyJSON := assert.MarshalWithoutError(t, body)
	assert.Equal(t, bodyJSON, "{\"query\":{\"bool\":{\"minimum_should_match\":7}}}")
}

func Test_Bool_should_has_SetBoost_method(t *testing.T) {
	// Given
	b := es.Bool()

	// When Then
	assert.NotNil(t, b.SetBoost)
}

func Test_Bool_SetBoost_should_create_json_with_minimum_should_match_field_inside_bool(t *testing.T) {
	// Given
	body := es.NewQuery(
		es.Bool().
			SetBoost(3.1415),
	)

	// When Then
	assert.NotNil(t, body)
	bodyJSON := assert.MarshalWithoutError(t, body)
	assert.Equal(t, bodyJSON, "{\"query\":{\"bool\":{\"boost\":3.1415}}}")
}

func Test_Bool_should_have_Filter_method(t *testing.T) {
	// Given
	b := es.Bool()

	// When Then
	assert.NotNil(t, b.Filter)
}

func Test_Bool_should_have_Must_method(t *testing.T) {
	// Given
	b := es.Bool()

	// When Then
	assert.NotNil(t, b.Must)
}

func Test_Bool_should_have_MustNot_method(t *testing.T) {
	// Given
	b := es.Bool()

	// When Then
	assert.NotNil(t, b.MustNot)
}

func Test_Bool_should_have_Should_method(t *testing.T) {
	// Given
	b := es.Bool()

	// When Then
	assert.NotNil(t, b.Should)
}

////   Term   ////

func Test_Term_should_exist_on_es_package(t *testing.T) {
	// Given When Then
	assert.NotNil(t, es.Term)
}

func Test_Term_should_create_json_with_term_field_inside_query(t *testing.T) {
	// Given
	body := es.NewQuery(
		es.Term("key", "value"),
	)

	// When Then
	assert.NotNil(t, body)
	bodyJSON := assert.MarshalWithoutError(t, body)
	assert.Equal(t, bodyJSON, "{\"query\":{\"term\":{\"key\":\"value\"}}}")
}

func Test_Term_method_should_create_termType(t *testing.T) {
	// Given
	b := es.Term("key", "value")

	// Then
	assert.NotNil(t, b)
	assert.IsTypeString(t, "es.termType", b)
}

////   Terms   ////

func Test_Terms_should_exist_on_es_package(t *testing.T) {
	// Given When Then
	assert.NotNil(t, es.Terms)
}

func Test_Terms_should_create_json_with_terms_field_inside_query(t *testing.T) {
	// Given
	body := es.NewQuery(
		es.Terms("key", "value1", "value2", "value3"),
	)

	// When Then
	assert.NotNil(t, body)
	bodyJSON := assert.MarshalWithoutError(t, body)
	assert.Equal(t, bodyJSON, "{\"query\":{\"terms\":{\"key\":[\"value1\",\"value2\",\"value3\"]}}}")
}

func Test_Terms_method_should_create_termsType(t *testing.T) {
	// Given
	b := es.Terms("key", "value1", "value2", "value3")

	// Then
	assert.NotNil(t, b)
	assert.IsTypeString(t, "es.termsType", b)
}

////   TermsArray   ////

func Test_TermsArray_should_exist_on_es_package(t *testing.T) {
	// Given When Then
	assert.NotNil(t, es.TermsArray)
}

func Test_TermsArray_should_create_json_with_terms_field_inside_query(t *testing.T) {
	// Given
	body := es.NewQuery(
		es.TermsArray("key", []any{"value1", "value2", "value3"}),
	)

	// When Then
	assert.NotNil(t, body)
	bodyJSON := assert.MarshalWithoutError(t, body)
	assert.Equal(t, bodyJSON, "{\"query\":{\"terms\":{\"key\":[\"value1\",\"value2\",\"value3\"]}}}")
}

func Test_TermsArray_method_should_create_termsType(t *testing.T) {
	// Given
	b := es.TermsArray("key", []any{"value1", "value2", "value3"})

	// Then
	assert.NotNil(t, b)
	assert.IsTypeString(t, "es.termsType", b)
}

////   Exists   ////

func Test_Exists_should_exist_on_es_package(t *testing.T) {
	// Given When Then
	assert.NotNil(t, es.Exists)
}

func Test_Exists_should_create_json_with_exists_field_inside_query(t *testing.T) {
	// Given
	body := es.NewQuery(
		es.Exists("key"),
	)

	// When Then
	assert.NotNil(t, body)
	bodyJSON := assert.MarshalWithoutError(t, body)
	assert.Equal(t, bodyJSON, "{\"query\":{\"exists\":{\"field\":\"key\"}}}")
}

func Test_Exists_method_should_create_existsType(t *testing.T) {
	// Given
	b := es.Exists("key")

	// Then
	assert.NotNil(t, b)
	assert.IsTypeString(t, "es.existsType", b)
}

////   Range   ////

func Test_Range_should_exist_on_es_package(t *testing.T) {
	// Given When Then
	assert.NotNil(t, es.Range)
}

func Test_Range_should_create_json_with_range_field_inside_query(t *testing.T) {
	// Given
	body := es.NewQuery(
		es.Range("age", 20, 10),
	)

	// When Then
	assert.NotNil(t, body)
	bodyJSON := assert.MarshalWithoutError(t, body)
	assert.Equal(t, bodyJSON, "{\"query\":{\"range\":{\"age\":{\"gte\":10,\"lte\":20}}}}")
}

func Test_Range_method_should_create_rangeType(t *testing.T) {
	// Given
	b := es.Range("age", 20, 10)

	// Then
	assert.NotNil(t, b)
	assert.IsTypeString(t, "es.rangeType", b)
}

////   Bool.Filter   ////

func Test_Filter_method_should_return_boolType(t *testing.T) {
	// Given
	b := es.Bool()

	// When
	filter := b.Filter()

	// Then
	assert.NotNil(t, filter)
	assert.IsTypeString(t, "es.boolType", filter)
}

////   Bool.Must   ////

func Test_Must_method_should_return_boolType(t *testing.T) {
	// Given
	b := es.Bool()

	// When
	must := b.Must()

	// Then
	assert.NotNil(t, must)
	assert.IsTypeString(t, "es.boolType", must)
}

////   Bool.MustNot   ////

func Test_MustNot_method_should_return_boolType(t *testing.T) {
	// Given
	b := es.Bool()

	// When
	mustNot := b.MustNot()

	// Then
	assert.NotNil(t, mustNot)
	assert.IsTypeString(t, "es.boolType", mustNot)
}

////   Bool.Should   ////

func Test_Should_method_should_return_boolType(t *testing.T) {
	// Given
	b := es.Bool()

	// When
	should := b.Should()

	// Then
	assert.NotNil(t, should)
	assert.IsTypeString(t, "es.boolType", should)
}
